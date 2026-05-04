$ErrorActionPreference = "Stop"
$Root = Split-Path -Parent $PSScriptRoot
$TestScript = Join-Path $Root "tests/run-mvp-tests.ps1"

& powershell -NoProfile -ExecutionPolicy Bypass -File $TestScript
$mvpExit = $LASTEXITCODE
if ($mvpExit -ne 0) {
  exit $mvpExit
}

$GoTestScript = Join-Path $Root "scripts/go-test.ps1"
if (Test-Path -LiteralPath $GoTestScript) {
  & powershell -NoProfile -ExecutionPolicy Bypass -File $GoTestScript
  exit $LASTEXITCODE
}
