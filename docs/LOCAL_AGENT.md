# 本地 Agent 设计

本地 Agent 是静态 Easy_Setup 页面与真实本机安装能力之间的桥接层。

## 目标

- 保持 GitHub Pages 应用静态、快速、低成本。
- 任何命令执行前都要求用户本机显式授权。
- 真实安装时打开可见 PowerShell 窗口。
- 写入 job 脚本和日志，方便审阅与排错。
- Agent 不在线时，网页仍可复制安装命令。

## 当前 MVP

启动 PowerShell Agent 预览模式：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1
```

启动 PowerShell Agent 执行模式：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1 -AllowExecute
```

启动 Go Agent 原型：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1 -AllowExecute
```

如需指定 job 目录：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1 -JobDir .easy-setup/logs
```

监听地址：

```text
http://127.0.0.1:17771
http://127.0.0.1:17772
```

`17771` 是 PowerShell Agent；`17772` 是 Go Agent 原型。

## API

```text
GET  /health
GET  /catalog
POST /plan
POST /execute
GET  /jobs
GET  /jobs/{id}
GET  /jobs/{id}/log
```

`POST /execute` 必须启用执行参数。PowerShell Agent 使用 `-AllowExecute`，Go Agent 使用 `-allow-execute`。否则接口会返回安全拒绝。

Go Agent 会在执行时创建 job 记录：

```text
.easy-setup/logs/<job-id>.ps1
.easy-setup/logs/<job-id>.log
.easy-setup/logs/<job-id>.json
```

`GET /jobs` 返回当前 job 目录中的历史记录，按创建时间倒序排列。`GET /jobs/{id}` 返回单个 job 的脚本路径、日志路径、状态、环境选择和展开后的 actions。

`GET /jobs/{id}/log` 返回该 job 日志文件的尾部内容。当前上限为 32 KB，后续会演进为实时日志流。

## 目录来源

Agent 读取共享目录：

```text
src/catalog/catalog.json
```

它会用 UI 传入的 stack ID 匹配 `stacks[].id`，展开 `recipeIds`，并读取 `recipes[].install.windows` 作为当前 Windows MVP 的安装命令。

## 安全说明

- 监听地址只绑定 `127.0.0.1`。
- CORS 仅用于本地原型与静态页面联调。
- 未显式启用执行模式时，`/execute` 不会运行命令。
- job 脚本和日志写入 `.easy-setup/logs/`。
- MVP 使用可见 PowerShell 窗口，让用户可以检查和中断命令。

## 工程路线

1. 保留 PowerShell Agent 作为 MVP 执行桥。
2. 使用 `cmd/easysetup-agent` 承载 Go HTTP 后端。
3. 在 Go Agent 中完善进度流、取消和重试。
4. 让 UI 展示 `/jobs` 和 `/jobs/{id}` 返回的执行历史。
5. 增加 recipe 签名和允许列表校验。
6. 增加 `winget`、`choco`、`scoop`、`brew`、`apt`、`dnf`、`pacman` 适配器。
7. 增加内容配置后端，让目录更新不依赖 UI 代码修改。
