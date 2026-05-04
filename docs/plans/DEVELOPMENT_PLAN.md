# 开发计划

## 阶段 0：方案确认

目标：确认边界、术语、清单格式、首批平台和 GUI 产品形态。

交付物：

- 项目设计文档。
- 架构草案。
- 技术选型说明。
- GUI 设计说明。
- 示例环境清单。

## 阶段 1：核心引擎与只读 GUI 原型

目标：实现只读规划能力，并让用户可以在界面中看懂当前环境状态。

功能：

- `envforge detect`
- `envforge validate <manifest>`
- `envforge plan <manifest>`
- GUI 项目选择和清单加载。
- GUI 环境目录浏览、搜索和勾选。
- GUI 环境状态看板。
- GUI 安装计划预览。
- 支持 Windows、macOS、Linux 的基础系统探测。

验收：

- 对示例清单输出明确的安装计划。
- GUI 能识别已安装、缺失、版本不匹配和可选工具。
- GUI 能把用户勾选的语言、框架和 AI 栈展开为安装计划。
- GUI 能展示将要写入的环境变量。
- 所有 GUI 数据来自 Core Engine，不写第二套业务逻辑。

## 阶段 2：用户级安装与环境变量

目标：在用户权限下完成真实部署，并提供可视化执行过程。

功能：

- `envforge apply <manifest>`
- GUI 执行按钮、风险确认、进度步骤和实时日志。
- Windows 支持 winget 和 scoop。
- macOS 支持 Homebrew。
- Linux 支持 apt。
- 用户级环境变量和 PATH 托管片段写入。
- 状态文件记录。

验收：

- 能安装并验证第一批主流语言工具链：C/C++、Java、JavaScript/TypeScript、Python、Go、Rust、Git。
- 能安装并验证第一批常用框架：React、Vue、Next.js、Spring Boot、Django、FastAPI。
- 能安装并验证第一批 AI 开发栈：PyTorch、Transformers、LangChain、LlamaIndex。
- 能写入并验证用户级环境变量。
- 重复执行具备幂等性。
- GUI 执行过程可以清楚定位当前步骤、耗时和失败原因。

## 阶段 3：验证、诊断与回滚

目标：让变更可信、可恢复，并让错误更容易理解。

功能：

- `envforge verify <manifest>`
- `envforge rollback --last`
- GUI 一键验证。
- GUI 回滚中心。
- 版本验证、命令验证、文件验证、环境变量验证。
- 失败时输出下一步建议。

验收：

- 回滚能撤销 EnvForge 管理的 profile 片段。
- 验证失败有清晰错误码。
- GUI 能把底层错误转成用户可执行的建议。
- 日志可定位到具体 recipe 和命令。

## 阶段 4：Recipe 生态

目标：让新增工具不需要修改核心代码。

功能：

- recipe 目录结构。
- recipe schema。
- catalog schema。
- stack template schema。
- recipe 来源锁定和哈希校验。
- 本地 recipe 优先级。
- GUI recipe 详情查看。

验收：

- 新增一个数据库客户端 recipe 不改核心代码。
- 新增一个语言、框架或 AI 栈模板不改核心代码。
- recipe 可以声明不同平台的安装方式。
- recipe 更新不会静默改变已锁定环境。

## 阶段 5：团队化能力

目标：支持团队规模使用。

功能：

- lock 文件。
- 组织级 recipe 索引。
- 代理、镜像源、离线包缓存。
- CI 中验证环境清单。
- GUI 导出诊断包，便于团队支持。

验收：

- 项目仓库能固定工具版本。
- CI 能检查环境清单合法性。
- 内网环境可通过镜像部署。
- 用户可把诊断包发给维护者复现问题。

## 阶段 6：性能与体验打磨

目标：确保跨平台体验稳定、启动快、资源占用低。

功能：

- 启动性能 profiling。
- 大日志虚拟滚动。
- 安装任务取消和恢复。
- 离线模式体验。
- 跨平台安装包签名和更新策略。

验收：

- GUI 冷启动接近原生小工具体验。
- 空闲内存保持在轻量桌面应用范围内。
- 长时间安装不会造成界面卡顿。
- Windows、macOS、Linux 的核心流程一致。

## 风险与对策

- 跨平台 profile 写入差异大：先收敛到用户级托管片段，并用状态文件记录修改。
- 包管理器行为不一致：adapter 明确能力边界，plan 阶段暴露差异。
- 权限操作危险：默认 user-only，高风险操作必须显式声明。
- “任意环境”范围过大：用 catalog、stack template 和 recipe 机制扩展，不把所有知识塞进核心。
- GUI 增加体积和复杂度：选择 Wails + 原生 WebView，保持 Core Engine 独立，避免 Electron 类重量。
