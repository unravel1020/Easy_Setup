# Easy_Setup

Easy_Setup is a lightweight, visual environment setup planner for developers. It helps users choose mainstream programming languages, frameworks, AI stacks, testing tools, DevOps tools, Linux mirrors, virtualization tools, and team-oriented workflows, then turns those choices into an auditable install plan.

The current project is an MVP: a static web UI plus a PowerShell-based CLI prototype. It is designed to validate the product flow before moving into the full engineering implementation.

## Highlights

- Visual environment catalog with Chinese / English switching.
- Light and dark themes.
- Collapsible sidebar and independent scrolling panels.
- Mainstream languages: C/C++, Java, JavaScript/TypeScript, Python, Go/Golang, Rust and more.
- Frontend stacks: React, Vue, Vite, Electron and related tooling.
- Backend stacks: Spring Boot, FastAPI, Go web frameworks, Rust web frameworks.
- AI development stacks: PyTorch, Transformers, LangChain, LlamaIndex, Jupyter.
- Testing and operations stacks: Playwright, pytest, Postman, Docker, kubectl, Helm, Terraform, Ansible.
- Extra sections for Skills, Vibe Coding, Harness, Linux mirrors and virtual machines.
- Learning resource links and official website links for environment cards.
- One-click copy for CLI preview and generated install commands.

## Live Demo

The static MVP is prepared for GitHub Pages deployment from:

```text
mvp/web
```

After GitHub Pages is enabled, the site will be available at:

```text
https://unravel1020.github.io/Easy_Setup/
```

## Quick Start

Open the local MVP UI:

```powershell
cd mvp
.\envforge.cmd
```

If PowerShell execution policy allows scripts, this also works:

```powershell
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 ui
```

Then open the printed `file:///.../mvp/web/index.html` URL in a browser.

## Local Agent MVP

The web UI stays lightweight and static. To let it execute real setup commands on the current Windows machine, start the local Agent first:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1
```

This starts a safe dry-run Agent at `http://127.0.0.1:17771`. The UI can detect it, but install buttons still fall back to copying commands.

To allow real execution in a new PowerShell window:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1 -AllowExecute
```

When `-AllowExecute` is enabled, clicking `Execute Install` or a single card's `Install` button sends the selected stack to the local Agent. The Agent writes a job script and log under `.easy-setup/logs/`, then opens a PowerShell window to run the commands.

## CLI MVP

The MVP CLI is intentionally safe by default. Planning and verification do not install anything.

```powershell
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 detect
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 catalog
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 plan examples/ai-backend.envforge.yaml
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 verify examples/ai-backend.envforge.yaml
```

The engineering CLI entrypoint is:

```powershell
powershell -ExecutionPolicy Bypass -File src/cmd/envforge.ps1 plan examples/ai-backend.envforge.yaml
```

The Go backend prototype is now available as a lightweight shared core and HTTP Agent:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/go-test.ps1
go run ./cmd/easysetup plan template.fullstack-web
go run ./cmd/easysetup-agent -port 17772
```

The Go code reads the same `src/catalog/catalog.json` used by the PowerShell MVP.

For convenience on Windows:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1 -AllowExecute
```

## Content Configuration

Executable recipes and stack templates are now centralized in:

```text
src/catalog/catalog.json
```

The CLI, core planner and local Agent all read this shared catalog. See `docs/CONTENT_CONFIGURATION.md` for the schema and update rules.

The static website also has a configurable presentation catalog:

```text
mvp/web/catalog.web.json
```

GitHub Pages loads this file first, while `app.js` keeps a fallback copy for local `file://` usage.

## Safety

The GitHub Pages UI cannot directly execute local commands by itself. Real execution requires the user to run the trusted local Agent with `-AllowExecute` on their own machine. Without that Agent, the UI only generates and copies auditable PowerShell commands.

## Test

Run the full MVP test suite:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/test.ps1
```

## Repository Structure

```text
.
├── .github/workflows/pages.yml
├── docs/
├── examples/
├── mvp/
│   ├── envforge.cmd
│   ├── envforge.ps1
│   └── web/
├── scripts/
├── src/
│   ├── catalog/
│   ├── cmd/
│   └── core/
└── tests/
```

## Roadmap

- Replace the minimal YAML reader with a robust manifest parser.
- Move the current PowerShell MVP core into a production-grade Go core.
- Build the desktop app with Wails and the shared core engine.
- Add safe execution, state tracking, rollback and audit logs.
- Expand recipe/catalog coverage for more languages, frameworks and AI tools.
- Support team templates, offline mirrors and enterprise package sources.

## License

License is not selected yet.
