$script:EnvForgeRoot = Split-Path -Parent (Split-Path -Parent $PSScriptRoot)
$script:EnvForgeCatalogPath = Join-Path $script:EnvForgeRoot "src/catalog/catalog.json"

function Initialize-EnvForgeCore {
  param(
    [string]$Root,
    [string]$CatalogPath
  )

  if ($Root) {
    $script:EnvForgeRoot = $Root
  }
  if ($CatalogPath) {
    $script:EnvForgeCatalogPath = $CatalogPath
  }
}

function Read-EnvForgeCatalog {
  Get-Content -Raw -LiteralPath $script:EnvForgeCatalogPath | ConvertFrom-Json
}

function Get-EnvForgePlatformKey {
  if ($IsMacOS) { return "macos" }
  if ($IsLinux) { return "linux" }
  return "windows"
}

function Test-EnvForgeCommandAvailable {
  param([Parameter(Mandatory)][string]$Name)

  $commandName = ($Name -split "\s+")[0]
  return [bool](Get-Command $commandName -ErrorAction SilentlyContinue)
}

function Test-EnvForgeVerifyCommand {
  param([Parameter(Mandatory)][string]$CommandLine)

  try {
    $null = powershell -NoProfile -ExecutionPolicy Bypass -Command $CommandLine 2>&1 | Out-String
    return $LASTEXITCODE -eq 0
  } catch {
    return $false
  }
}

function Test-EnvForgeRecipeInstalled {
  param([Parameter(Mandatory)]$Recipe)

  if (-not (Test-EnvForgeCommandAvailable $Recipe.detect)) {
    return $false
  }
  if ($Recipe.verify -and $Recipe.verify.Count -gt 0) {
    return Test-EnvForgeVerifyCommand ([string]$Recipe.verify[0])
  }
  return $true
}

function Read-EnvForgeManifest {
  param([Parameter(Mandatory)][string]$Path)

  $fullPath = Join-Path $script:EnvForgeRoot $Path
  if (-not (Test-Path -LiteralPath $fullPath)) {
    $fullPath = $Path
  }
  if (-not (Test-Path -LiteralPath $fullPath)) {
    throw "Manifest not found: $Path"
  }

  $result = [ordered]@{
    path = $fullPath
    name = ""
    description = ""
    tools = @()
    stacks = @()
    env = [ordered]@{}
    verifyCommands = @()
  }

  $section = ""
  foreach ($line in Get-Content -LiteralPath $fullPath) {
    $trim = $line.Trim()
    if ($trim -eq "" -or $trim.StartsWith("#")) { continue }

    if ($trim -match "^name:\s*(.+)$" -and $section -eq "") {
      $result.name = $Matches[1].Trim('"')
      continue
    }
    if ($trim -match "^description:\s*(.+)$" -and $section -eq "") {
      $result.description = $Matches[1].Trim('"')
      continue
    }
    if ($trim -match "^(constraints|tools|stacks|env|path|verify):\s*$") {
      $section = $Matches[1]
      continue
    }
    if ($section -eq "verify" -and $trim -match "^commands:\s*$") {
      $section = "verify.commands"
      continue
    }
    if ($section -eq "tools" -and $trim -match "^- id:\s*(.+)$") {
      $result.tools += $Matches[1].Trim('"')
      continue
    }
    if ($section -eq "stacks" -and $trim -match "^- id:\s*(.+)$") {
      $result.stacks += $Matches[1].Trim('"')
      continue
    }
    if ($section -eq "env" -and $trim -match "^([A-Za-z_][A-Za-z0-9_]*):\s*(.*)$") {
      $result.env[$Matches[1]] = $Matches[2].Trim('"')
      continue
    }
    if ($section -eq "verify.commands" -and $trim -match "^- (.+)$") {
      $result.verifyCommands += $Matches[1].Trim()
      continue
    }
  }

  return [pscustomobject]$result
}

function Get-EnvForgeDetectResult {
  $checks = @("git", "node", "npm", "pnpm", "python", "go", "rustc", "cargo", "java", "cmake", "docker", "jupyter")
  $tools = foreach ($item in $checks) {
    [pscustomobject]@{
      id = $item
      installed = (Test-EnvForgeCommandAvailable $item)
    }
  }

  [pscustomobject]@{
    os = [System.Runtime.InteropServices.RuntimeInformation]::OSDescription
    architecture = [System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture.ToString()
    shell = "powershell"
    platformKey = (Get-EnvForgePlatformKey)
    tools = $tools
  }
}

function Resolve-EnvForgeRecipes {
  param(
    [Parameter(Mandatory)]$Manifest,
    [Parameter(Mandatory)]$Catalog
  )

  $recipeMap = @{}
  foreach ($recipe in $Catalog.recipes) { $recipeMap[$recipe.id] = $recipe }
  $stackMap = @{}
  foreach ($stack in $Catalog.stacks) { $stackMap[$stack.id] = $stack }

  $ids = New-Object System.Collections.Generic.List[string]
  foreach ($tool in $Manifest.tools) { $ids.Add($tool) }
  foreach ($stackId in $Manifest.stacks) {
    if ($stackMap.ContainsKey($stackId)) {
      foreach ($recipeId in $stackMap[$stackId].recipeIds) { $ids.Add($recipeId) }
    }
  }

  $unique = [ordered]@{}
  foreach ($id in $ids) {
    if ($recipeMap.ContainsKey($id)) { $unique[$id] = $recipeMap[$id] }
  }
  return @($unique.Values)
}

function Get-EnvForgePlan {
  param([Parameter(Mandatory)][string]$ManifestPath)

  $catalog = Read-EnvForgeCatalog
  $manifestData = Read-EnvForgeManifest $ManifestPath
  $platform = Get-EnvForgePlatformKey
  $recipes = Resolve-EnvForgeRecipes -Manifest $manifestData -Catalog $catalog

  $actions = foreach ($recipe in $recipes) {
    $installed = Test-EnvForgeRecipeInstalled $recipe
    $installCommand = $recipe.install.$platform
    [pscustomobject]@{
      id = $recipe.id
      name = $recipe.name
      category = $recipe.category
      installed = $installed
      action = $(if ($installed) { "skip" } else { "install" })
      command = $(if ($installed) { "" } else { $installCommand })
      verify = $recipe.verify
    }
  }

  $envActions = foreach ($key in $manifestData.env.Keys) {
    [pscustomobject]@{
      name = $key
      current = [Environment]::GetEnvironmentVariable($key, "User")
      desired = $manifestData.env[$key]
      action = "set-user-env"
    }
  }

  [pscustomobject]@{
    manifest = $manifestData
    platform = $platform
    actions = @($actions)
    env = @($envActions)
    verifyCommands = @($manifestData.verifyCommands)
  }
}

function Invoke-EnvForgeVerify {
  param([Parameter(Mandatory)][string]$ManifestPath)

  $plan = Get-EnvForgePlan $ManifestPath
  $commands = New-Object System.Collections.Generic.List[string]
  foreach ($action in $plan.actions) {
    foreach ($verify in $action.verify) { $commands.Add($verify) }
  }
  foreach ($verify in $plan.verifyCommands) { $commands.Add($verify) }

  foreach ($command in ($commands | Select-Object -Unique)) {
    $passed = $false
    $output = ""
    try {
      $output = powershell -NoProfile -ExecutionPolicy Bypass -Command $command 2>&1 | Out-String
      $passed = $LASTEXITCODE -eq 0
    } catch {
      $output = $_.Exception.Message
    }
    [pscustomobject]@{
      command = $command
      passed = $passed
      output = $output.Trim()
    }
  }
}

Export-ModuleMember -Function Initialize-EnvForgeCore, Read-EnvForgeCatalog, Get-EnvForgeDetectResult, Get-EnvForgePlan, Invoke-EnvForgeVerify
