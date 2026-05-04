# Content Configuration

Easy_Setup is moving toward configuration-driven content. The goal is to add or update languages, frameworks, AI stacks, mirrors and templates without changing execution code.

## Shared Catalog

The current shared catalog lives at:

```text
src/catalog/catalog.json
```

It is used by:

- `src/core/EnvForgeCore.psm1`
- `src/cmd/envforge.ps1`
- `src/agent/EasySetupAgent.ps1`

The static MVP UI still contains a richer presentation catalog in `mvp/web/app.js`. That UI data will be migrated into a web catalog file in the next step, but install execution already resolves through the shared catalog.

## Main Objects

`categories` define navigable groups.

`recipes` define executable units:

```json
{
  "id": "python",
  "name": "Python",
  "category": "languages",
  "detect": "python",
  "verify": ["python --version"],
  "install": {
    "windows": "winget install --id Python.Python.3.12 -e",
    "macos": "brew install python@3.12",
    "linux": "sudo apt install -y python3 python3-venv python3-pip"
  }
}
```

`stacks` group recipes into one selectable environment:

```json
{
  "id": "template.fullstack-web",
  "name": "Fullstack Web",
  "description": "Git, Node.js, pnpm, React, Vite and FastAPI.",
  "recipeIds": ["git", "node", "pnpm", "frontend.react", "frontend.vite", "python", "backend.fastapi"]
}
```

## Update Rules

- Add a tool as a `recipe`.
- Add a selectable card/template as a `stack`.
- Keep recipe IDs stable once published.
- Use platform-specific install commands under `install.windows`, `install.macos` and `install.linux`.
- Keep commands auditable and non-destructive.
- Prefer official package managers first: `winget`, `brew`, `apt`, then specialized installers.
- Put smoke checks in `verify` so the Agent can test installations after execution.

## Next Engineering Step

The next step is to extract the UI presentation content into `mvp/web/catalog.web.json`, then load it with `fetch`. After that, GitHub Pages content can be updated by editing JSON instead of editing JavaScript.
