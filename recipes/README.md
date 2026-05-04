# Recipe 目录

本目录预留给独立 recipe 文件。Recipe 用来描述某个语言、框架、SDK、AI 库或工具在不同平台上的安装、检测、验证和卸载方式。

当前 MVP 的 recipe 仍集中在：

```text
src/catalog/catalog.json
```

未来可以把共享目录拆分为多个文件，例如：

```text
recipes/
|-- node.yaml
|-- python.yaml
|-- go.yaml
|-- java.yaml
|-- rust.yaml
|-- cpp.yaml
|-- pytorch.yaml
|-- transformers.yaml
|-- langchain.yaml
|-- docker.yaml
`-- postgres-client.yaml
```

## 设计原则

- Recipe 应保持声明式。
- 优先使用系统包管理器或语言包管理器。
- 只有在包管理器无法表达时，才允许脚本步骤。
- 环境栈通过 catalog 或 stack template 组合多个 recipe。
- 命令必须可审计，危险操作必须显式声明。
