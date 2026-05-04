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

## Safety

The static web UI cannot directly execute local install commands because browsers intentionally block local command execution from `file://` and GitHub Pages pages. For the MVP, install buttons generate and copy real PowerShell commands. A later desktop version can execute those commands through a trusted local bridge such as Wails.

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
