# EnvForge MVP

这是 EnvForge 的最小可用原型，目标是先验证产品闭环：

- 检测本机开发工具。
- 读取示例环境清单。
- 展开用户选择的语言、框架和 AI 栈。
- 生成安装计划。
- 执行验证命令。
- 提供一个可勾选的静态 GUI 原型。

## CLI

```powershell
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 detect
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 catalog
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 plan examples/fullstack-web.envforge.yaml
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 plan examples/ai-backend.envforge.yaml
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 verify examples/ai-backend.envforge.yaml
```

`plan` 默认是 dry-run，不会修改系统。添加 `-Execute` 才会运行计划中的安装命令。

## GUI

```powershell
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 ui
```

在浏览器中打开输出的 `mvp/web/index.html` 即可查看可勾选环境目录原型。

在 `mvp/` 目录下也可以直接运行：

```powershell
.\envforge.ps1
```

无参数运行会默认进入 UI 模式，输出本地 HTML 路径和 `file:///` URL。

如果系统阻止直接运行 `.ps1`，可以使用：

```powershell
.\envforge.cmd
```

## 当前限制

- YAML 解析器只覆盖当前示例清单所需的简单结构。
- 静态 GUI 暂不直接调用 PowerShell 后端，只用于验证选择和计划体验。
- 真正桌面版仍建议后续迁移到 Go Core + Wails。
