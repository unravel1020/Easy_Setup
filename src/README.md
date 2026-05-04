# Source

此目录预留给 EnvForge 源码。项目采用“共享核心 + CLI + 桌面端”的结构。

建议初始模块：

```text
src/
├── cmd/
│   └── envforge.ps1
├── core/
│   └── EnvForgeCore.psm1
├── catalog/
│   └── catalog.json
├── desktop/
│   ├── app/
│   └── bindings/
└── frontend/
    ├── src/
    └── package.json
```

规划说明：

- `core/` 是唯一业务核心，CLI 和 GUI 都必须调用它。
- `cmd/envforge.ps1` 放当前工程化 CLI 入口。
- `catalog/` 放语言、框架、AI 栈和模板目录。
- `desktop/` 放 Wails 桌面应用壳和 Go 绑定。
- `frontend/` 放轻量前端界面，建议使用 TypeScript + Svelte 或 Solid。

当前阶段先用 PowerShell 标准能力实现可运行核心，保证没有 Go/Node/Python 的新机器也能测试方案。后续再把 `core/` 的接口迁移到 Go module，并让 Wails 桌面端调用同一套能力。
