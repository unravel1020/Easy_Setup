param(
  [switch]$AllowExecute,
  [int]$Port = 17771
)

$ErrorActionPreference = "Stop"
$Root = Split-Path -Parent $PSScriptRoot
$Agent = Join-Path $Root "src/agent/EasySetupAgent.ps1"

if ($AllowExecute) {
  & powershell -NoProfile -ExecutionPolicy Bypass -File $Agent -Port $Port -AllowExecute
} else {
  & powershell -NoProfile -ExecutionPolicy Bypass -File $Agent -Port $Port
}
