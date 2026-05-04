# Scripts

此目录用于维护开发和发布辅助脚本，例如格式化、打包、生成 schema、构建多平台二进制。

原则：

- 不把核心逻辑写在脚本里。
- 脚本只服务开发流程。
- 生产能力应进入 CLI 核心或 recipe。

当前 MVP 入口位于 `mvp/envforge.ps1`，用于在没有 Go/Node/Python 的机器上先跑通检测、目录、计划、验证和静态 GUI 原型。

## 测试

统一测试入口：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/test.ps1
```

测试会覆盖：

- 环境目录结构。
- 本机检测输出。
- AI 和全栈示例计划生成。
- 验证命令输出结构。
- 静态 GUI 文件完整性。
- `mvp/envforge.ps1` 兼容入口和 `src/cmd/envforge.ps1` 工程化入口。
