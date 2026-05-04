param(
  [int]$Port = 17771,
  [switch]$AllowExecute
)

$ErrorActionPreference = "Stop"
$Root = Split-Path -Parent (Split-Path -Parent $PSScriptRoot)
$CatalogPath = Join-Path $Root "src/catalog/catalog.json"
$LogDir = Join-Path $Root ".easy-setup/logs"
New-Item -ItemType Directory -Force -Path $LogDir | Out-Null

function Read-AgentCatalog {
  Get-Content -Raw -Encoding UTF8 -LiteralPath $CatalogPath | ConvertFrom-Json
}

function ConvertTo-AgentJson($Value) {
  $Value | ConvertTo-Json -Depth 16 -Compress
}

function Write-AgentResponse($Stream, [int]$StatusCode, $Value) {
  $json = ConvertTo-AgentJson $Value
  $bodyBytes = [Text.Encoding]::UTF8.GetBytes($json)
  $statusText = switch ($StatusCode) {
    200 { "OK" }
    403 { "Forbidden" }
    404 { "Not Found" }
    default { "Internal Server Error" }
  }
  $headers = @(
    "HTTP/1.1 $StatusCode $statusText"
    "Content-Type: application/json; charset=utf-8"
    "Content-Length: $($bodyBytes.Length)"
    "Access-Control-Allow-Origin: *"
    "Access-Control-Allow-Methods: GET,POST,OPTIONS"
    "Access-Control-Allow-Headers: Content-Type"
    "Connection: close"
    ""
    ""
  ) -join "`r`n"
  $headerBytes = [Text.Encoding]::ASCII.GetBytes($headers)
  $Stream.Write($headerBytes, 0, $headerBytes.Length)
  $Stream.Write($bodyBytes, 0, $bodyBytes.Length)
}

function Read-AgentRequest($Client) {
  $stream = $Client.GetStream()
  $buffer = New-Object byte[] 8192
  $memory = New-Object IO.MemoryStream
  $headerEnd = -1
  $contentLength = 0

  while ($true) {
    $read = $stream.Read($buffer, 0, $buffer.Length)
    if ($read -le 0) { break }
    $memory.Write($buffer, 0, $read)
    $bytes = $memory.ToArray()
    $text = [Text.Encoding]::ASCII.GetString($bytes)
    $headerEnd = $text.IndexOf("`r`n`r`n")
    if ($headerEnd -ge 0) {
      $headerText = $text.Substring(0, $headerEnd)
      if ($headerText -match "(?im)^Content-Length:\s*(\d+)") {
        $contentLength = [int]$matches[1]
      }
      if ($bytes.Length -ge ($headerEnd + 4 + $contentLength)) { break }
    }
  }

  $bytes = $memory.ToArray()
  $requestText = [Text.Encoding]::ASCII.GetString($bytes)
  $headerEnd = $requestText.IndexOf("`r`n`r`n")
  if ($headerEnd -lt 0) {
    return [pscustomobject]@{ stream = $stream; method = ""; path = ""; body = [pscustomobject]@{} }
  }

  $headerText = $requestText.Substring(0, $headerEnd)
  $firstLine = ($headerText -split "`r`n")[0]
  $parts = $firstLine -split " "
  $method = $parts[0]
  $target = $parts[1]
  $path = ($target -split "\?")[0].TrimEnd("/")
  if ($path -eq "") { $path = "/" }

  $bodyObject = [pscustomobject]@{}
  if ($contentLength -gt 0) {
    $bodyStart = $headerEnd + 4
    $bodyText = [Text.Encoding]::UTF8.GetString($bytes, $bodyStart, $contentLength)
    if (-not [string]::IsNullOrWhiteSpace($bodyText)) {
      $bodyObject = $bodyText | ConvertFrom-Json
    }
  }

  [pscustomobject]@{
    stream = $stream
    method = $method
    path = $path
    body = $bodyObject
  }
}

function Resolve-AgentPlan($ItemIds) {
  $catalog = Read-AgentCatalog
  $stackMap = @{}
  foreach ($stack in $catalog.stacks) { $stackMap[$stack.id] = $stack }
  $recipeMap = @{}
  foreach ($recipe in $catalog.recipes) { $recipeMap[$recipe.id] = $recipe }

  $recipeIds = New-Object Collections.Generic.List[string]
  foreach ($itemId in @($ItemIds)) {
    if ($stackMap.ContainsKey($itemId)) {
      foreach ($recipeId in $stackMap[$itemId].recipeIds) { $recipeIds.Add($recipeId) }
      continue
    }
    if ($recipeMap.ContainsKey($itemId)) {
      $recipeIds.Add($itemId)
    }
  }

  $unique = [ordered]@{}
  foreach ($recipeId in $recipeIds) {
    if ($recipeMap.ContainsKey($recipeId)) { $unique[$recipeId] = $recipeMap[$recipeId] }
  }

  $recipes = @($unique.Values)
  $platform = "windows"
  $commands = @($recipes | ForEach-Object { $_.install.$platform })
  [pscustomobject]@{
    itemIds = @($ItemIds)
    recipes = $recipes
    commands = $commands
    command = ($commands -join "; ")
  }
}

function Start-AgentExecution($Plan) {
  $jobId = [guid]::NewGuid().ToString("N")
  $scriptPath = Join-Path $LogDir "$jobId.ps1"
  $logPath = Join-Path $LogDir "$jobId.log"
  $script = @(
    '$ErrorActionPreference = "Continue"'
    "Start-Transcript -Path `"$logPath`" -Force"
  )
  foreach ($command in $Plan.commands) {
    $script += "Write-Host `"Running: $command`""
    $script += $command
  }
  $script += "Stop-Transcript"
  Set-Content -Encoding UTF8 -LiteralPath $scriptPath -Value $script
  Start-Process powershell -ArgumentList @("-NoExit", "-ExecutionPolicy", "Bypass", "-File", $scriptPath) | Out-Null
  [pscustomobject]@{
    jobId = $jobId
    scriptPath = $scriptPath
    logPath = $logPath
  }
}

$listener = [Net.Sockets.TcpListener]::new([Net.IPAddress]::Parse("127.0.0.1"), $Port)
$listener.Start()
Write-Host "Easy_Setup Agent listening on http://127.0.0.1:$Port/"
Write-Host "AllowExecute: $AllowExecute"
Write-Host "Press Ctrl+C to stop."

$running = $true
while ($running) {
  $client = $listener.AcceptTcpClient()
  try {
    $request = Read-AgentRequest $client
    if ($request.method -eq "OPTIONS") {
      Write-AgentResponse $request.stream 200 ([pscustomobject]@{ ok = $true })
      continue
    }

    if ($request.method -eq "GET" -and $request.path -eq "/health") {
      Write-AgentResponse $request.stream 200 ([pscustomobject]@{
        ok = $true
        name = "Easy_Setup Agent"
        version = "0.1.0"
        allowExecute = [bool]$AllowExecute
      })
      continue
    }

    if ($request.method -eq "GET" -and $request.path -eq "/catalog") {
      Write-AgentResponse $request.stream 200 (Read-AgentCatalog)
      continue
    }

    if ($request.method -eq "POST" -and $request.path -eq "/plan") {
      $plan = Resolve-AgentPlan $request.body.itemIds
      Write-AgentResponse $request.stream 200 $plan
      continue
    }

    if ($request.method -eq "POST" -and $request.path -eq "/execute") {
      $plan = Resolve-AgentPlan $request.body.itemIds
      if (-not $AllowExecute) {
        Write-AgentResponse $request.stream 403 ([pscustomobject]@{
          ok = $false
          reason = "Agent was not started with -AllowExecute."
          plan = $plan
        })
        continue
      }
      $job = Start-AgentExecution $plan
      Write-AgentResponse $request.stream 200 ([pscustomobject]@{
        ok = $true
        job = $job
        plan = $plan
      })
      continue
    }

    Write-AgentResponse $request.stream 404 ([pscustomobject]@{ ok = $false; reason = "Not found" })
  } catch {
    try {
      Write-AgentResponse $client.GetStream() 500 ([pscustomobject]@{ ok = $false; reason = $_.Exception.Message })
    } catch {
      Write-Host $_.Exception.Message
    }
  } finally {
    $client.Close()
  }
}
