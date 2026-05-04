# 内容配置说明

Easy_Setup 正在向配置驱动演进。目标是让语言、框架、AI 栈、镜像源和模板的新增或更新尽量通过 JSON/YAML 完成，而不是修改执行代码。

## 共享执行目录

当前共享目录位于：

```text
src/catalog/catalog.json
```

使用它的模块包括：

- `src/core/EnvForgeCore.psm1`
- `src/cmd/envforge.ps1`
- `src/agent/EasySetupAgent.ps1`
- `internal/catalog`
- `cmd/easysetup`
- `cmd/easysetup-agent`

## 静态页面展示目录

静态 MVP 页面优先加载：

```text
mvp/web/catalog.web.json
```

`mvp/web/app.js` 仍保留内置兜底目录，用于浏览器阻止本地 `file://` 读取 JSON 的场景。GitHub Pages 和普通 HTTP 托管会优先使用 `catalog.web.json`。

## 主要对象

`categories` 定义导航分组。

`recipes` 定义一个可执行单元：

```json
{
  "id": "python",
  "name": "Python",
  "category": "languages",
  "detect": "python",
  "verify": ["python --version"],
  "install": {
    "windows": "winget install --id Python.Python.3.12 -e",
    "macos": "brew install python@3.12",
    "linux": "sudo apt install -y python3 python3-venv python3-pip"
  }
}
```

`stacks` 把多个 recipe 组合成一个可选择环境：

```json
{
  "id": "template.fullstack-web",
  "name": "Fullstack Web",
  "description": "Git, Node.js, pnpm, React, Vite and FastAPI.",
  "recipeIds": ["git", "node", "pnpm", "frontend.react", "frontend.vite", "python", "backend.fastapi"]
}
```

## 更新规则

- 新增工具时添加 `recipe`。
- 新增可选择卡片或模板时添加 `stack`。
- 已发布的 recipe ID 应保持稳定。
- 平台安装命令写在 `install.windows`、`install.macos`、`install.linux` 下。
- 命令必须可审计，避免默认执行破坏性操作。
- 优先使用官方包管理器：`winget`、`brew`、`apt`，再考虑专用安装器。
- 在 `verify` 中提供冒烟验证命令，便于安装后检查。

## 下一步

下一步是从 `src/catalog/catalog.json` 和展示元数据生成 `mvp/web/catalog.web.json`，避免执行目录和 UI 卡片长期漂移。
