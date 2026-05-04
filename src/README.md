# 源码目录

`src/` 保存 PowerShell MVP 的工程化入口和共享目录。Go 后端代码位于仓库根目录下的 `internal/` 和 `cmd/`。

## 当前结构

```text
src/
|-- agent/
|   `-- EasySetupAgent.ps1
|-- catalog/
|   `-- catalog.json
|-- cmd/
|   `-- envforge.ps1
`-- core/
    `-- EnvForgeCore.psm1
```

## 说明

- `core/` 是 PowerShell MVP 的核心规划模块。
- `cmd/envforge.ps1` 是 PowerShell CLI 工程化入口。
- `agent/` 是 PowerShell 本地 Agent 原型。
- `catalog/catalog.json` 是 CLI、Agent 和 Go 后端共享的执行目录。

后续迁移方向是让 Go 的 `internal/catalog`、`internal/agent` 和 Wails 桌面应用逐步替代 PowerShell MVP，但在迁移完成前保留 PowerShell 入口，保证原型可运行。
