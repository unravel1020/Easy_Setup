# 可视化界面设计

## 设计目标

EnvForge 的界面要让用户知道三件事：

1. 我当前机器缺什么。
2. EnvForge 准备改什么。
3. 出错后我该怎么恢复或继续。

界面不追求炫技，重点是清楚、可信、低干扰。

## 主要页面

### 1. 项目选择

用户可以选择一个包含 `*.envforge.yaml` 的项目目录，也可以直接打开清单文件。

页面元素：

- 最近项目列表。
- 选择目录按钮。
- 清单文件状态。
- schema 校验结果。

### 2. 环境选择

用户可以像选择应用组件一样勾选开发环境。

分类：

- 语言工具链：C/C++、Java、JavaScript/TypeScript、Python、Go、Rust、C#/.NET、PHP、Ruby、Kotlin、Swift、Dart/Flutter、R、Julia。
- Web 前端：React、Vue、Angular、Svelte、Next.js、Nuxt、Vite。
- 后端框架：Spring Boot、Quarkus、Micronaut、Django、FastAPI、Flask、Express、NestJS、Gin、Rails、Laravel、ASP.NET Core。
- AI/ML：PyTorch、TensorFlow、Transformers、LangChain、LlamaIndex、ONNX Runtime、Jupyter、CUDA 工具链。
- 数据库与中间件：PostgreSQL、MySQL、Redis、MongoDB、SQLite、Kafka、RabbitMQ。
- DevOps：Docker、Kubernetes CLI、Helm、Terraform、Ansible、云厂商 CLI。
- 移动端：Android SDK、Flutter、React Native、iOS 相关工具。

交互：

- 支持搜索。
- 支持推荐组合，例如“AI Python 工作站”“Java 后端”“C/C++ 原生开发”“全栈 Web”。
- 每个勾选项展示体积、预计耗时、依赖、是否需要重启 shell、是否需要管理员权限。
- 勾选后立即生成可审阅的安装计划。

### 3. 环境看板

展示当前机器与目标环境的差异。

模块：

- 系统信息：OS、架构、shell、包管理器。
- 工具状态：已安装、缺失、版本不匹配、可选。
- 环境变量状态：已满足、将新增、将修改、冲突。
- 风险提示：需要管理员权限、需要网络、需要重启 shell。

### 4. 安装计划

在执行前展示完整计划，避免“黑盒安装”。

每个步骤展示：

- 动作类型：安装、升级、写入变量、修改 PATH、验证。
- 目标工具或变量。
- 当前状态和目标状态。
- 预计使用的包管理器。
- 是否可回滚。

### 5. 执行进度

安装时展示有节奏的进度，而不是只丢一堆日志。

模块：

- 当前步骤。
- 总进度。
- 实时日志。
- 可展开的底层命令。
- 失败原因和建议操作。

### 6. 验证与诊断

执行完成后，用户可以一键验证环境。

展示：

- 每个验证项的通过或失败状态。
- 版本输出。
- 缺失命令。
- 环境变量实际值。
- 诊断包导出入口。

### 7. 回滚中心

展示 EnvForge 管理过的变更记录，并允许撤销最近一次或指定一次变更。

回滚前展示：

- 将撤销的环境变量。
- 将移除的 PATH 片段。
- 将恢复的 profile 内容。
- 不可自动卸载或需要人工处理的工具。

## 交互原则

- 默认只做用户级变更。
- 用户勾选环境后，必须先生成计划，再执行安装。
- 所有高风险操作都要在执行前明确展示。
- 一键安装应包含安装、配置、验证三个阶段，而不是只执行下载。
- 错误信息要包含“发生了什么”和“下一步做什么”。
- 长日志默认折叠，但用户可以展开查看完整命令。
- 不把用户困在向导里，允许返回计划页调整清单。

## 性能原则

- 首屏只加载必要数据，重型扫描按需触发。
- 日志使用虚拟列表。
- 安装任务在后端执行，前端只订阅事件。
- 不使用常驻后台服务。
- 不引入大型 UI 组件库。
