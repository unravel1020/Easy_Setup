param(
  [Parameter(Position = 0)]
  [ValidateSet("help", "detect", "catalog", "plan", "verify", "ui")]
  [string]$Command = "help",

  [Parameter(Position = 1)]
  [string]$Manifest = "examples/fullstack-web.envforge.yaml",

  [switch]$Json,
  [switch]$Execute
)

$ErrorActionPreference = "Stop"
$RepoRoot = Split-Path -Parent (Split-Path -Parent $PSScriptRoot)
$ModulePath = Join-Path $RepoRoot "src/core/EnvForgeCore.psm1"
$CatalogPath = Join-Path $RepoRoot "src/catalog/catalog.json"

Import-Module $ModulePath -Force
Initialize-EnvForgeCore -Root $RepoRoot -CatalogPath $CatalogPath

function Write-EnvForgeObjectResult($Value) {
  if ($Json) {
    $Value | ConvertTo-Json -Depth 12
  } else {
    $Value | Format-List
  }
}

function Invoke-EnvForgePlanCommand([string]$ManifestPath) {
  $plan = Get-EnvForgePlan $ManifestPath
  if ($Json) {
    $plan | ConvertTo-Json -Depth 12
    return
  }

  Write-Host "EnvForge plan: $($plan.manifest.name)" -ForegroundColor Cyan
  Write-Host "Platform: $($plan.platform)"
  Write-Host ""
  Write-Host "Actions:"
  foreach ($action in $plan.actions) {
    $marker = if ($action.installed) { "OK" } else { "TODO" }
    Write-Host ("  [{0}] {1} -> {2}" -f $marker, $action.name, $action.action)
    if ($action.command) { Write-Host ("       {0}" -f $action.command) -ForegroundColor DarkGray }
  }

  Write-Host ""
  Write-Host "Environment changes:"
  foreach ($env in $plan.env) {
    Write-Host ("  {0} = {1}" -f $env.name, $env.desired)
  }

  if ($Execute) {
    Write-Host ""
    Write-Host "Execute mode is enabled. Running install commands..." -ForegroundColor Yellow
    foreach ($action in $plan.actions) {
      if ($action.action -eq "install" -and $action.command) {
        Write-Host "Running: $($action.command)" -ForegroundColor Yellow
        powershell -NoProfile -ExecutionPolicy Bypass -Command $action.command
      }
    }
    foreach ($env in $plan.env) {
      [Environment]::SetEnvironmentVariable($env.name, $env.desired, "User")
    }
  } else {
    Write-Host ""
    Write-Host "Dry run only. Add -Execute to run install commands and write user environment variables." -ForegroundColor DarkYellow
  }
}

function Invoke-EnvForgeVerifyCommand([string]$ManifestPath) {
  $results = @(Invoke-EnvForgeVerify $ManifestPath)
  if ($Json) {
    $results | ConvertTo-Json -Depth 8
    return
  }

  foreach ($result in $results) {
    $marker = if ($result.passed) { "PASS" } else { "FAIL" }
    $color = if ($result.passed) { "Green" } else { "Red" }
    Write-Host ("[{0}] {1}" -f $marker, $result.command) -ForegroundColor $color
    if ($result.output) { Write-Host ("      {0}" -f $result.output.Split("`n")[0]) -ForegroundColor DarkGray }
  }
}

switch ($Command) {
  "detect" { Write-EnvForgeObjectResult (Get-EnvForgeDetectResult) }
  "catalog" { Write-EnvForgeObjectResult (Read-EnvForgeCatalog) }
  "plan" { Invoke-EnvForgePlanCommand $Manifest }
  "verify" { Invoke-EnvForgeVerifyCommand $Manifest }
  "ui" {
    $index = Join-Path $RepoRoot "mvp/web/index.html"
    $fileUrl = "file:///" + ($index -replace "\\", "/")
    Write-Host "Open this file in your browser:" -ForegroundColor Cyan
    Write-Host $index
    Write-Host ""
    Write-Host "URL:" -ForegroundColor Cyan
    Write-Host $fileUrl
  }
  default {
    Write-Host "EnvForge"
    Write-Host ""
    Write-Host "Usage:"
    Write-Host "  powershell -ExecutionPolicy Bypass -File src/cmd/envforge.ps1 detect"
    Write-Host "  powershell -ExecutionPolicy Bypass -File src/cmd/envforge.ps1 catalog"
    Write-Host "  powershell -ExecutionPolicy Bypass -File src/cmd/envforge.ps1 plan examples/ai-backend.envforge.yaml"
    Write-Host "  powershell -ExecutionPolicy Bypass -File src/cmd/envforge.ps1 verify examples/ai-backend.envforge.yaml"
    Write-Host "  powershell -ExecutionPolicy Bypass -File src/cmd/envforge.ps1 ui"
  }
}
