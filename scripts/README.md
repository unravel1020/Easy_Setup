# 脚本目录

本目录存放开发、测试、启动和发布辅助脚本。脚本只服务工程流程，不承载核心业务逻辑。

## 常用命令

运行完整测试：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/test.ps1
```

运行 Go 后端测试：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/go-test.ps1
```

启动 PowerShell Agent：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-agent.ps1 -AllowExecute
```

启动 Go Agent：

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/start-go-agent.ps1 -AllowExecute
```

## 维护原则

- 不把核心规划、执行和验证逻辑写进脚本。
- 可复用能力应进入 Go 核心、PowerShell core 或 recipe。
- 脚本可以负责设置本地环境变量、调用测试、启动 Agent 和生成发布产物。
