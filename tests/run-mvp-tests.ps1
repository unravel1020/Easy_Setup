param(
  [switch]$VerboseOutput
)

$ErrorActionPreference = "Stop"
$Root = Split-Path -Parent $PSScriptRoot
$Cli = Join-Path $Root "mvp/envforge.ps1"
$SourceCli = Join-Path $Root "src/cmd/envforge.ps1"
$Failures = New-Object System.Collections.Generic.List[string]

function Add-Failure([string]$Message) {
  $script:Failures.Add($Message)
  Write-Host "[FAIL] $Message" -ForegroundColor Red
}

function Assert-True([bool]$Condition, [string]$Message) {
  if ($Condition) {
    Write-Host "[PASS] $Message" -ForegroundColor Green
  } else {
    Add-Failure $Message
  }
}

function Invoke-CliJson([string[]]$ArgsList) {
  $output = & powershell -NoProfile -ExecutionPolicy Bypass -File $Cli @ArgsList -Json 2>&1
  if ($LASTEXITCODE -ne 0) {
    throw "CLI failed: $($ArgsList -join ' ')`n$output"
  }
  if ($VerboseOutput) {
    Write-Host $output
  }
  return ($output | Out-String | ConvertFrom-Json)
}

function Invoke-CliText([string[]]$ArgsList) {
  $output = & powershell -NoProfile -ExecutionPolicy Bypass -File $Cli @ArgsList 2>&1
  if ($LASTEXITCODE -ne 0) {
    throw "CLI failed: $($ArgsList -join ' ')`n$output"
  }
  if ($VerboseOutput) {
    Write-Host $output
  }
  return ($output | Out-String)
}

Write-Host "Running EnvForge MVP tests..." -ForegroundColor Cyan

$catalog = Invoke-CliJson @("catalog")
Assert-True ($catalog.version -eq "0.2.0") "catalog exposes expected version"
$catalogRecipes = @($catalog.recipes)
$catalogStacks = @($catalog.stacks)
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "python" }).Count -eq 1) "catalog contains Python recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "ai.pytorch" }).Count -eq 1) "catalog contains PyTorch recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "ai.transformers" }).Count -eq 1) "catalog contains Transformers recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "ai.langchain" }).Count -eq 1) "catalog contains LangChain recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "testing.playwright" }).Count -eq 1) "catalog contains Playwright recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "testing.pytest" }).Count -eq 1) "catalog contains pytest recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "frontend.vue" }).Count -eq 1) "catalog contains Vue recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "frontend.electron" }).Count -eq 1) "catalog contains Electron recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "backend.spring" }).Count -eq 1) "catalog contains Spring recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "backend.gin" }).Count -eq 1) "catalog contains Go Gin recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "backend.axum" }).Count -eq 1) "catalog contains Rust Axum recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "kubernetes.kubectl" }).Count -eq 1) "catalog contains kubectl recipe"
Assert-True (@($catalogRecipes | Where-Object { $_.id -eq "terraform" }).Count -eq 1) "catalog contains Terraform recipe"
Assert-True (@($catalogStacks | Where-Object { $_.id -eq "template.ai-python-workstation" }).Count -eq 1) "catalog contains AI workstation template"
Assert-True (@($catalogStacks | Where-Object { $_.id -eq "template.fullstack-web" }).Count -eq 1) "catalog contains fullstack web template"
Assert-True (@($catalogStacks | Where-Object { $_.id -eq "testing.web" }).Count -eq 1) "catalog contains web testing stack"
Assert-True (@($catalogStacks | Where-Object { $_.id -eq "operations.kubernetes" }).Count -eq 1) "catalog contains Kubernetes ops stack"
Assert-True (@($catalogStacks | Where-Object { $_.id -eq "harness.delivery" }).Count -eq 1) "catalog contains Harness delivery stack"
Assert-True (@($catalogStacks | Where-Object { $_.id -eq "virtualization.vm" }).Count -eq 1) "catalog contains virtualization stack"

$detect = Invoke-CliJson @("detect")
Assert-True ([string]::IsNullOrWhiteSpace($detect.os) -eq $false) "detect returns OS description"
Assert-True ($detect.platformKey -in @("windows", "macos", "linux")) "detect returns supported platform key"
Assert-True (@($detect.tools | Where-Object { $_.id -eq "git" }).Count -eq 1) "detect includes git check"

$aiPlan = Invoke-CliJson @("plan", "examples/ai-backend.envforge.yaml")
$aiActionIds = @($aiPlan.actions | ForEach-Object id)
Assert-True ($aiPlan.manifest.name -eq "ai-backend-workstation") "AI manifest name parsed"
Assert-True ($aiActionIds -contains "python") "AI plan includes Python"
Assert-True ($aiActionIds -contains "ai.pytorch") "AI plan includes PyTorch"
Assert-True ($aiActionIds -contains "ai.transformers") "AI plan includes Transformers"
Assert-True ($aiActionIds -contains "ai.langchain") "AI plan includes LangChain"
Assert-True (@($aiPlan.env | Where-Object { $_.name -eq "HF_HOME" }).Count -eq 1) "AI plan includes HF_HOME environment variable"

$webPlan = Invoke-CliJson @("plan", "examples/fullstack-web.envforge.yaml")
$webActionIds = @($webPlan.actions | ForEach-Object id)
$webEnvNames = @($webPlan.env | ForEach-Object name)
Assert-True ($webPlan.manifest.name -eq "fullstack-web") "fullstack manifest name parsed"
Assert-True ($webActionIds -contains "node") "fullstack plan includes Node.js"
Assert-True ($webActionIds -contains "frontend.react") "fullstack plan includes React"
Assert-True ($webActionIds -contains "backend.fastapi") "fullstack plan includes FastAPI"
Assert-True (-not ($webEnvNames -contains "path")) "fullstack plan does not misparse path as environment variable"
Assert-True (-not ($webEnvNames -contains "append")) "fullstack plan does not misparse path.append as environment variable"

$verify = Invoke-CliJson @("verify", "examples/ai-backend.envforge.yaml")
$verifyResults = @($verify)
Assert-True (@($verifyResults | Where-Object { $_.command -eq "git --version" }).Count -ge 1) "verify emits git command result"
Assert-True (@($verifyResults | Where-Object { $_.command -eq "python --version" }).Count -ge 1) "verify emits Python command result"
Assert-True (@($verifyResults | Where-Object { $_.PSObject.Properties.Name -contains "passed" }).Count -eq $verifyResults.Count) "verify results include pass/fail field"

$uiOutput = Invoke-CliText @("ui")
$defaultMvpOutput = Invoke-CliText @()
$cmdOutput = & (Join-Path $Root "mvp/envforge.cmd") 2>&1 | Out-String
$indexPath = Join-Path $Root "mvp/web/index.html"
$appPath = Join-Path $Root "mvp/web/app.js"
$stylePath = Join-Path $Root "mvp/web/styles.css"
$agentPath = Join-Path $Root "src/agent/EasySetupAgent.ps1"
$sharedCatalogPath = Join-Path $Root "src/catalog/catalog.json"
$agentStartPath = Join-Path $Root "scripts/start-agent.ps1"
Assert-True ($uiOutput.Contains($indexPath)) "ui command prints static GUI path"
Assert-True ($uiOutput.Contains("file:///")) "ui command prints file URL"
Assert-True ($defaultMvpOutput.Contains($indexPath)) "mvp wrapper defaults to UI when run without args"
Assert-True ($cmdOutput.Contains($indexPath)) "cmd wrapper defaults to UI without PowerShell policy friction"
Assert-True (Test-Path -LiteralPath $indexPath) "static GUI index exists"
Assert-True (Test-Path -LiteralPath $appPath) "static GUI script exists"
Assert-True (Test-Path -LiteralPath $stylePath) "static GUI stylesheet exists"
Assert-True (Test-Path -LiteralPath $agentPath) "local Agent script exists"
Assert-True (Test-Path -LiteralPath $sharedCatalogPath) "shared catalog exists for CLI and Agent"
Assert-True (Test-Path -LiteralPath $agentStartPath) "local Agent start script exists"

$indexHtml = Get-Content -Raw -LiteralPath $indexPath
$appJs = Get-Content -Raw -LiteralPath $appPath
$stylesCss = Get-Content -Raw -LiteralPath $stylePath
$agentPs1 = Get-Content -Raw -LiteralPath $agentPath
Assert-True ($indexHtml.Contains("./app.js")) "static GUI references app.js"
Assert-True ($indexHtml.Contains("./styles.css")) "static GUI references styles.css"
Assert-True ($appJs.Contains("template.ai-python-workstation")) "static GUI exposes AI template"
Assert-True ($appJs.Contains("template.fullstack-web")) "static GUI exposes fullstack template"
Assert-True ($appJs.Contains("themeToggle")) "static GUI exposes theme toggle"
Assert-True ($appJs.Contains("collapseSidebar")) "static GUI exposes sidebar collapse"
Assert-True ($appJs.Contains("testing.web")) "static GUI exposes testing tools"
Assert-True ($appJs.Contains("operations.kubernetes")) "static GUI exposes operations tools"
Assert-True ($appJs.Contains("skills.codex")) "static GUI exposes skills section"
Assert-True ($appJs.Contains("vibe.webapp")) "static GUI exposes vibe coding section"
Assert-True ($appJs.Contains("harness.delivery")) "static GUI exposes Harness section"
Assert-True ($appJs.Contains("language.go")) "static GUI exposes Go language stack"
Assert-True ($appJs.Contains("frontend.vue")) "static GUI exposes Vue stack"
Assert-True ($appJs.Contains("frontend.electron")) "static GUI exposes Electron stack"
Assert-True ($appJs.Contains("backend.spring")) "static GUI exposes Spring stack"
Assert-True ($appJs.Contains("backend.go-web")) "static GUI exposes Go web frameworks"
Assert-True ($appJs.Contains("backend.rust-web")) "static GUI exposes Rust web frameworks"
Assert-True ($appJs.Contains("linux.mirrors")) "static GUI exposes Linux mirrors"
Assert-True ($appJs.Contains("virtualization.vm")) "static GUI exposes VM stack"
Assert-True ($appJs.Contains("resources:")) "static GUI includes learning resources"
Assert-True ($appJs.Contains("officialLinks")) "static GUI includes official website links"
Assert-True ($appJs.Contains("resourceLinks")) "static GUI includes learning resource links"
Assert-True ($appJs.Contains("iconPaths")) "static GUI includes inline vector icon paths"
Assert-True ($appJs.Contains("buildInstallCommand")) "static GUI builds install commands"
Assert-True ($appJs.Contains("copyText")) "static GUI supports copy to clipboard"
Assert-True ($appJs.Contains("markCopied")) "static GUI marks clicked button as copied"
Assert-True ($appJs.Contains("setButtonText")) "static GUI updates copied button text"
Assert-True ($appJs.Contains("data-row")) "static GUI supports row click selection"
Assert-True ($appJs.Contains("category-label")) "static GUI renders category as non-button label"
Assert-True ($appJs.Contains("title-meta")) "static GUI places category metadata beside item title"
Assert-True ($indexHtml.Contains("installSelected")) "static GUI exposes install selected action"
Assert-True ($indexHtml.Contains("copyCli")) "static GUI exposes CLI copy action"
Assert-True ($indexHtml.Contains("top-action")) "static GUI uses coordinated top action buttons"
Assert-True ($indexHtml.Contains("sr-only")) "static GUI uses accessible icon buttons"
Assert-True ($stylesCss.Contains(".icon")) "static GUI styles vector icons"
Assert-True ($stylesCss.Contains(".install-one")) "static GUI styles per-item install actions"
Assert-True ($stylesCss.Contains("button.copied")) "static GUI styles copied button state"
Assert-True ($indexHtml.Contains("languageToggle")) "static GUI exposes language toggle"
Assert-True ($appJs.Contains("const translations")) "static GUI includes translation table"
Assert-True ($appJs.Contains("zh:")) "static GUI includes Chinese locale"
Assert-True ($appJs.Contains("language: localStorage")) "static GUI persists language choice"
Assert-True ($appJs.Contains("Environment setup")) "static GUI includes English translations"
Assert-True ($indexHtml.Contains("scroll-area")) "static GUI marks independent scroll areas"
Assert-True ($stylesCss.Contains("overflow: hidden")) "static GUI prevents whole-page scrolling"
Assert-True ($stylesCss.Contains("overscroll-behavior: contain")) "static GUI contains scroll within active panel"
Assert-True ($indexHtml.Contains("agentStatus")) "static GUI exposes local Agent status"
Assert-True ($appJs.Contains("127.0.0.1:17771")) "static GUI targets local Agent loopback port"
Assert-True ($appJs.Contains("refreshAgentStatus")) "static GUI checks local Agent health"
Assert-True ($appJs.Contains("executeViaAgent")) "static GUI can execute through local Agent"
Assert-True ($stylesCss.Contains(".agent-status")) "static GUI styles local Agent status"
Assert-True ($agentPs1.Contains("/health")) "local Agent exposes health endpoint"
Assert-True ($agentPs1.Contains("/execute")) "local Agent exposes execute endpoint"
Assert-True ($agentPs1.Contains("AllowExecute")) "local Agent requires explicit execution opt-in"
Assert-True ($agentPs1.Contains("Start-Process")) "local Agent opens a visible execution window"
Assert-True ($agentPs1.Contains("src/catalog/catalog.json")) "local Agent reads the shared catalog"

$pagesWorkflow = Join-Path $Root ".github/workflows/pages.yml"
$pagesDoc = Join-Path $Root "docs/GITHUB_PAGES.md"
$contentDoc = Join-Path $Root "docs/CONTENT_CONFIGURATION.md"
Assert-True (Test-Path -LiteralPath $pagesWorkflow) "GitHub Pages workflow exists"
Assert-True (Test-Path -LiteralPath $pagesDoc) "GitHub Pages documentation exists"
Assert-True (Test-Path -LiteralPath $contentDoc) "content configuration documentation exists"

$sourceOutput = & powershell -NoProfile -ExecutionPolicy Bypass -File $SourceCli plan "examples/ai-backend.envforge.yaml" -Json 2>&1
if ($LASTEXITCODE -eq 0) {
  $sourcePlan = $sourceOutput | Out-String | ConvertFrom-Json
  Assert-True ($sourcePlan.manifest.name -eq "ai-backend-workstation") "source CLI returns AI plan"
  Assert-True (@($sourcePlan.actions | Where-Object { $_.id -eq "ai.pytorch" }).Count -eq 1) "source CLI uses extracted core planner"
} else {
  Add-Failure "source CLI plan command exits successfully"
}

if ($Failures.Count -gt 0) {
  Write-Host ""
  Write-Host "$($Failures.Count) test(s) failed." -ForegroundColor Red
  exit 1
}

Write-Host ""
Write-Host "All MVP tests passed." -ForegroundColor Green
