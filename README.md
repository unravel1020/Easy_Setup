# Easy_Setup

Easy_Setup 是一个轻量化、可视化的开发环境配置工具。它帮助用户选择主流编程语言、框架、AI 开发栈、测试工具、DevOps 工具、Linux 镜像源、虚拟化工具和团队工作流，并把选择结果转换为可审计的安装计划。

当前项目处于 MVP 到工程化迁移阶段：已经包含静态 Web UI、PowerShell CLI/Agent 原型，以及正在推进的 Go 核心与 Go Agent。

## 功能亮点

- 可视化环境目录，支持中文/英文切换。
- 浅色/深色主题。
- 可收起侧边栏和独立滚动面板。
- 支持 C/C++、Java、JavaScript/TypeScript、Python、Go、Rust 等主流语言。
- 支持 React、Vue、Vite、Electron、Spring Boot、FastAPI、Go Web、Rust Web 等框架。
- 支持 PyTorch、Transformers、LangChain、LlamaIndex、Jupyter 等 AI 开发栈。
- 支持 Playwright、pytest、Postman、Docker、kubectl、Helm、Terraform、Ansible 等测试和运维工具。
- 支持 Skills、Vibe Coding、Harness、Linux 镜像源、虚拟机等扩展板块。
- 提供入门资料链接和官网链接。
- 支持 CLI 预览复制、安装命令复制和本地 Agent 执行。

## 在线演示

GitHub Pages 静态页面目录：

```text
mvp/web
```

访问地址：

```text
https://unravel1020.github.io/Easy_Setup/
```

## 快速开始

打开本地 MVP UI：

```powershell
cd mvp
.\envforge.cmd
```

也可以直接调用 PowerShell 入口：

```powershell
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 ui
```

然后在浏览器中打开输出的 `file:///.../mvp/web/index.html` 地址。

## 本地 Agent

静态 Web 页面本身不能直接执行本机命令。需要先启动受信任的本地 Agent。

PowerShell Agent 安全预览模式：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1
```

PowerShell Agent 执行模式：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1 -AllowExecute
```

Go Agent 原型：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1 -AllowExecute
```

启用执行模式后，页面点击 `执行安装` 或单个卡片的 `安装` 按钮，会把所选环境栈发送给本地 Agent。Agent 会在 `.easy-setup/logs/` 下写入 job 脚本和日志，并打开可见 PowerShell 窗口执行。

## CLI

PowerShell MVP CLI 默认只规划和验证，不会修改系统：

```powershell
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 detect
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 catalog
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 plan examples/ai-backend.envforge.yaml
powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 verify examples/ai-backend.envforge.yaml
```

工程化 PowerShell 入口：

```powershell
powershell -ExecutionPolicy Bypass -File src/cmd/envforge.ps1 plan examples/ai-backend.envforge.yaml
```

Go 后端原型：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/go-test.ps1
go run ./cmd/easysetup plan template.fullstack-web
go run ./cmd/easysetup-agent -port 17772
```

## 内容配置

可执行 recipe 和环境栈模板集中在：

```text
src/catalog/catalog.json
```

PowerShell CLI、Go 核心和本地 Agent 都读取这份共享目录。

静态网站展示目录位于：

```text
mvp/web/catalog.web.json
```

GitHub Pages 会优先加载这个文件；`app.js` 保留内置兜底数据，保证本地 `file://` 场景也能运行。

更多说明见 [docs/CONTENT_CONFIGURATION.md](docs/CONTENT_CONFIGURATION.md)。

## 安全边界

GitHub Pages 页面不能单独执行本机命令。真实执行必须由用户在本机显式启动 Agent，并传入 `-AllowExecute` 或 `-allow-execute`。没有 Agent 时，页面只会生成和复制可审计命令。

## 测试

运行完整测试：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/test.ps1
```

## 目录结构

```text
.
|-- .github/workflows/pages.yml
|-- cmd/
|   |-- easysetup/
|   `-- easysetup-agent/
|-- docs/
|-- examples/
|-- internal/
|   |-- agent/
|   `-- catalog/
|-- mvp/
|   |-- envforge.cmd
|   |-- envforge.ps1
|   `-- web/
|-- recipes/
|-- scripts/
|-- src/
|   |-- agent/
|   |-- catalog/
|   |-- cmd/
|   `-- core/
`-- tests/
```

## 路线图

- 用 Go 完整替代 PowerShell MVP 核心。
- 为 Go Agent 增加 job 状态查询、进度流、取消和重试。
- 增加 recipe 签名、允许列表和更细的风险确认。
- 用 Wails 构建轻量跨平台桌面应用。
- 扩展更多语言、框架、AI 工具和企业镜像源。
- 支持团队模板、离线包缓存和诊断包导出。

## 许可证

暂未选择许可证。
