# 测试目录

本目录存放集成测试、fixture 和测试说明。

## 当前测试入口

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/test.ps1
```

该入口会运行：

- PowerShell MVP 测试。
- 静态 Web UI 文件完整性检查。
- 共享 catalog 校验。
- 本地 Agent 基础能力校验。
- Go 后端单元测试。

## 后续测试分层

- 单元测试：解析、规划、环境变量合并、状态文件。
- 适配器测试：用 mock 命令验证不同平台安装器行为。
- 集成测试：在容器或虚拟机中验证真实安装流程。
- 回滚测试：确保 Easy_Setup 只撤销自己管理的变更。
