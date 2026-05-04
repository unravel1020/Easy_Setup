param(
  [switch]$AllowExecute,
  [int]$Port = 17772
)

$ErrorActionPreference = "Stop"
$Root = Split-Path -Parent $PSScriptRoot
$LocalGo = Join-Path $Root ".easy-setup/tools/go"
$LocalGoPath = Join-Path $Root ".easy-setup/tools/gopath"
$LocalGoCache = Join-Path $Root ".easy-setup/cache/go-build"

if (Test-Path -LiteralPath $LocalGo) {
  $env:GOROOT = $LocalGo
  $env:GOPATH = $LocalGoPath
  $env:GOCACHE = $LocalGoCache
  New-Item -ItemType Directory -Force -Path $env:GOPATH, $env:GOCACHE | Out-Null
  $env:Path = "$LocalGo\bin;$LocalGoPath\bin;$env:Path"
}

$argsList = @("run", "./cmd/easysetup-agent", "-port", "$Port")
if ($AllowExecute) {
  $argsList += "-allow-execute"
}

Push-Location $Root
try {
  & go @argsList
} finally {
  Pop-Location
}

