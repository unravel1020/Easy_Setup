# EnvForge MVP

这是 Easy_Setup 的最小可用原型。目标是先验证产品闭环：

- 检测本机开发工具。
- 读取示例环境清单。
- 展开用户选择的语言、框架和 AI 栈。
- 生成安装计划。
- 执行验证命令。
- 提供可勾选的静态 GUI 原型。

## CLI

```powershell
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 detect
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 catalog
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 plan examples/fullstack-web.envforge.yaml
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 plan examples/ai-backend.envforge.yaml
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 verify examples/ai-backend.envforge.yaml
```

`plan` 默认是 dry-run，不会修改系统。只有显式执行模式才会运行安装命令。

## GUI

```powershell
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 ui
```

在浏览器中打开输出的 `mvp/web/index.html` 即可查看可勾选环境目录原型。

也可以在 `mvp/` 目录中运行：

```powershell
.\envforge.ps1
```

如果系统阻止直接运行 `.ps1`，可以使用：

```powershell
.\envforge.cmd
```

## 当前限制

- YAML 解析器只覆盖当前示例清单需要的简单结构。
- 静态 GUI 主要验证选择、计划和本地 Agent 调用体验。
- 真正桌面版仍建议迁移到 Go Core + Wails。
