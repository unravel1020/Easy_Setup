# Local Agent Design

The Local Agent is the bridge between the static Easy_Setup UI and real environment installation on a user's machine.

## Goals

- Keep the public GitHub Pages app static, fast and cheap to host.
- Require explicit local opt-in before any command can execute.
- Open a visible PowerShell window for real installation work.
- Store generated job scripts and logs for review.
- Let the website keep working without the Agent by copying commands.

## Current MVP

Start dry-run mode:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1
```

Start execution mode:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1 -AllowExecute
```

Start the Go Agent prototype:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1 -AllowExecute
```

The Agent listens on:

```text
http://127.0.0.1:17771
http://127.0.0.1:17772
```

`17771` is the PowerShell Agent. `17772` is the Go Agent prototype.

## API

```text
GET  /health
GET  /catalog
POST /plan
POST /execute
```

`POST /execute` requires `-AllowExecute`. Otherwise it returns a safe refusal plus the generated plan.

## Catalog Source

The Agent reads the shared catalog from:

```text
src/catalog/catalog.json
```

It resolves selected UI stack IDs against `stacks[].id`, expands their `recipeIds`, and reads platform install commands from `recipes[].install.windows` for the current MVP.

## Security Notes

- The listener binds only to `127.0.0.1`.
- CORS is intentionally narrow enough for local prototype use.
- Execution is disabled unless the user passes `-AllowExecute`.
- Jobs are written to `.easy-setup/logs/` before execution.
- The MVP opens a visible PowerShell window so the user can inspect and interrupt commands.

## Engineering Path

1. Keep the PowerShell Agent as the MVP execution bridge while the Go Agent matures.
2. Use `cmd/easysetup-agent` as the new Go HTTP backend for `/health`, `/catalog` and `/plan`.
3. Port job logging, visible execution windows and explicit confirmation into the Go Agent.
4. Add recipe signatures and allowlist validation.
5. Add job state, progress streaming, cancellation and retry.
6. Add package-manager adapters for winget, choco, scoop, brew, apt, dnf and pacman.
7. Add a content configuration backend so catalog updates can be pulled without changing UI code.
