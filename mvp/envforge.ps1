$ErrorActionPreference = "Stop"
$RepoRoot = Split-Path -Parent $PSScriptRoot
$CommandPath = Join-Path $RepoRoot "src/cmd/envforge.ps1"

if ($args.Count -eq 0) {
  & powershell -NoProfile -ExecutionPolicy Bypass -File $CommandPath ui
} else {
  & powershell -NoProfile -ExecutionPolicy Bypass -File $CommandPath @args
}
exit $LASTEXITCODE
