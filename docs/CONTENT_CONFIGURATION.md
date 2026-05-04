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

The static MVP UI loads its presentation catalog from:

```text
mvp/web/catalog.web.json
```

`mvp/web/app.js` still keeps an embedded fallback catalog so local `file://` usage continues to work in browsers that block JSON fetches. GitHub Pages and normal HTTP hosting use `catalog.web.json` first.

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

The next step is to generate `mvp/web/catalog.web.json` from `src/catalog/catalog.json` plus presentation metadata, so execution recipes and UI cards cannot drift apart.
