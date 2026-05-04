# 架构设计

## 总体思路

EnvForge 采用“轻量桌面壳 + CLI + 共享核心引擎 + 声明式 recipe”的结构。GUI 和 CLI 只负责交互入口，所有探测、规划、执行、状态记录和验证逻辑都集中在 Core Engine 中。

## 架构分层

```text
Desktop App
├── Project Picker
├── Environment Catalog
├── Stack Selector
├── Environment Dashboard
├── Plan Review
├── Apply Progress
├── Logs & Diagnostics
└── Rollback Center

CLI
├── detect
├── validate
├── plan
├── apply
├── verify
└── rollback

Core Engine
├── Catalog Resolver
├── Manifest Parser
├── System Detector
├── Planner
├── Executor
├── Environment Manager
├── Verifier
├── State Store
└── Adapters
    ├── Windows: winget / scoop / PowerShell profile
    ├── macOS: brew / zsh profile
    └── Linux: apt / dnf / pacman / zsh/bash profile
```

## GUI 与核心通信

推荐使用进程内绑定或轻量 IPC，而不是本地 Web 服务：

- 桌面应用直接调用 Core Engine 暴露的 API。
- 长任务通过事件流上报进度、日志、警告和确认请求。
- CLI 调用同一套 API，并将结果格式化为文本或 JSON。
- 不启动常驻后台服务，避免端口占用和额外资源消耗。

## 关键流程

1. `detect`：识别 OS、CPU 架构、shell、权限、包管理器、已安装工具。
2. `select`：GUI 读取环境目录，让用户勾选语言、框架、AI/ML、数据库和工具链。
3. `resolve`：把用户选择的 stack template 展开为具体 recipe、版本、环境变量和验证项。
4. `parse`：读取项目清单和引用的 recipe。
5. `plan`：比较目标状态与当前状态，生成操作图。
6. `review`：GUI 展示计划、风险、环境变量差异和权限需求。
7. `apply`：按依赖顺序执行安装、配置和验证。
8. `record`：写入状态文件，记录 EnvForge 管理过的变更。
9. `rollback`：按状态文件撤销最近一次变更。

## 配置模型

清单文件描述“想要什么”，而不是写死“怎么安装”：

```yaml
schema: envforge/v1
name: fullstack-web
tools:
  - id: node
    version: "22"
  - id: python
    version: "3.12"
  - id: go
    version: "1.24"
env:
  PROJECT_ENV: development
  NODE_OPTIONS: --max-old-space-size=4096
verify:
  commands:
    - node --version
    - python --version
    - go version
```

GUI 勾选后的清单也可以保存为声明式配置：

```yaml
schema: envforge/v1
name: ai-backend-workstation
stacks:
  - id: language.python
    version: "3.12"
  - id: language.rust
    version: stable
  - id: ai.pytorch
    variant: cpu
  - id: ai.transformers
  - id: ai.langchain
  - id: backend.fastapi
tools:
  - id: git
  - id: docker
verify:
  commands:
    - python -c "import torch; print(torch.__version__)"
    - python -c "import transformers; print(transformers.__version__)"
    - rustc --version
```

## 环境目录模型

环境目录分为三层：

- Catalog：用户在 GUI 里看到的分类、搜索、说明、标签和推荐组合。
- Stack Template：一组可复用环境栈，例如 `AI Python Workstation`、`Java Backend`、`C/C++ Native`。
- Recipe：具体工具或框架的安装、配置、检测、验证和回滚规则。

核心只理解统一模型，不把某个语言或框架的安装细节写死。

## 环境变量策略

- 优先支持项目级环境文件，例如 `.envrc`、`.env`、`.envforge/session.env`。
- 需要持久化到用户 shell 时，写入带标记的托管片段。
- GUI 必须在写入前展示变量名、旧值、新值、作用域和回滚方式。
- 所有写入内容必须记录到状态文件。
- 对 PATH 修改做去重、排序保持和来源标记。
- 默认不写系统级变量，除非用户显式声明并确认。

## 状态存储

默认位置：

```text
~/.envforge/state.json
~/.envforge/logs/
```

项目级状态：

```text
.envforge/lock.json
```

状态文件记录：

- 清单哈希。
- 执行时间。
- 操作列表。
- 修改前后的环境变量值。
- 已安装工具的版本和安装方式。
- GUI 展示所需的执行摘要。
- 回滚所需的信息。

## 性能与资源约束

- 冷启动目标：常规机器上 1 秒左右进入主界面。
- 空闲内存目标：明显低于 Electron 类应用，优先控制在几十 MB 级别。
- 执行安装时只保留必要日志缓冲，完整日志落盘。
- 不轮询扫描系统；使用用户触发或操作完成后的增量刷新。
- 下载、安装和验证任务串并结合，避免对磁盘和网络造成突发压力。

## 安全边界

- `plan` 永远不修改系统。
- `apply` 对高风险操作要求确认，例如系统级变量、sudo、修改全局 profile。
- recipe 默认不可执行任意脚本；需要脚本时必须声明权限和来源。
- 支持 `--offline`、`--no-scripts`、`--user-only` 等约束模式。
- GUI 中高风险步骤必须可展开查看详细命令和变更。
