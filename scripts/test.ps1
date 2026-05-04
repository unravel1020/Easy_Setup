$ErrorActionPreference = "Stop"
$Root = Split-Path -Parent $PSScriptRoot
$TestScript = Join-Path $Root "tests/run-mvp-tests.ps1"

& powershell -NoProfile -ExecutionPolicy Bypass -File $TestScript
exit $LASTEXITCODE
