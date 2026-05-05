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
GET  /jobs/{id}/events
POST /jobs/{id}/cancel
POST /jobs/{id}/retry
```

`POST /execute` 必须启用执行参数。PowerShell Agent 使用 `-AllowExecute`，Go Agent 使用 `-allow-execute`。否则接口会返回安全拒绝。

`POST /plan` 返回的每个 action 都包含 `risk` 字段。当前风险分级根据命令文本启发式生成，用于提示包管理器安装、语言生态包安装、配置修改、远程脚本执行、提权和破坏性文件操作等信号。

如果计划中存在 `risk.level = high` 的 action，`POST /execute` 和 `POST /jobs/{id}/retry` 必须在请求体中传入 `confirmHighRisk: true`。缺少确认时，Go Agent 会返回 `409 Conflict`，并在响应中带回被拦截的高风险 actions。

每个 action 还包含 `trust` 字段。内置目录默认可信；外部目录可以通过 recipe 的 `trust` 元数据标记来源。如果计划中存在 `trust.trusted = false` 的 action，`POST /execute` 和 `POST /jobs/{id}/retry` 必须传入 `confirmUntrusted: true`，否则 Go Agent 会返回 `409 Conflict`。

Go Agent 会在执行时创建 job 记录：

```text
.easy-setup/logs/<job-id>.ps1
.easy-setup/logs/<job-id>.log
.easy-setup/logs/<job-id>.json
.easy-setup/logs/<job-id>.cancel
```

`GET /jobs` 返回当前 job 目录中的历史记录，按创建时间倒序排列。`GET /jobs/{id}` 返回单个 job 的 JSON 路径、脚本路径、日志路径、状态、退出码、完成时间、环境选择和展开后的 actions。

当前状态包括：

- `created`：记录已创建，尚未拉起脚本。
- `launched`：脚本窗口已拉起，等待脚本回写运行状态。
- `running`：脚本已经开始执行。
- `cancel-requested`：用户已请求取消，脚本会在下一个步骤边界停止。
- `completed`：脚本执行完成，退出码为 0。
- `failed`：脚本执行完成，但安装或验证命令返回非 0 退出码。
- `canceled`：脚本响应取消请求后停止，退出码为 130。
- `launch-failed`：Agent 未能拉起本地脚本窗口。

`GET /jobs/{id}/log` 返回该 job 日志文件的尾部内容。当前上限为 32 KB，后续会演进为实时日志流。

`GET /jobs/{id}/events` 是 Server-Sent Events 进度流，会持续推送当前 job 快照和日志尾部内容。当前事件名为 `job`，数据结构包含 `job` 和 `log`。任务进入 `completed`、`failed`、`canceled` 或 `launch-failed` 后，流会自动结束。UI 展开日志面板时会优先使用该接口，轮询仍作为兜底。

`POST /jobs/{id}/cancel` 会写入取消标记并把 job 状态改为 `cancel-requested`。当前取消是协作式取消：正在执行的命令不会被强制杀死，但脚本会在每个安装或验证步骤之间检查取消标记，并尽快以 `canceled` 状态结束。

`POST /jobs/{id}/retry` 会读取原 job 的 `itemIds`，重新生成计划并创建一个新的 job。重试不会覆盖原 job 的脚本和日志，因此失败记录可以继续保留用于排错。该接口和 `/execute` 一样，要求 Go Agent 使用 `-allow-execute` 启动。

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
3. 增加更细的步骤级进度事件。
4. 让 UI 展示 `/jobs`、`/jobs/{id}` 和 `/jobs/{id}/events` 返回的执行历史。
5. 增加 recipe 签名和允许列表校验。
6. 增加 `winget`、`choco`、`scoop`、`brew`、`apt`、`dnf`、`pacman` 适配器。
7. 增加内容配置后端，让目录更新不依赖 UI 代码修改。
