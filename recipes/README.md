# Recipes

此目录用于存放工具、框架和环境栈 recipe。Recipe 描述某个语言、框架、SDK、AI 库或工具在不同平台上的安装、检测、验证和卸载方式。

示例设想：

```text
recipes/
├── node.yaml
├── python.yaml
├── go.yaml
├── java.yaml
├── rust.yaml
├── cpp.yaml
├── pytorch.yaml
├── transformers.yaml
├── langchain.yaml
├── docker.yaml
└── postgres-client.yaml
```

Recipe 应保持声明式，只有确实无法通过包管理器表达时才允许脚本步骤。

环境栈应通过 catalog 和 stack template 组合多个 recipe，例如：

- AI Python 工作站：Python、uv、Jupyter、PyTorch、Transformers、LangChain、LlamaIndex。
- C/C++ 原生开发：编译器、CMake、Ninja、调试器、vcpkg 或 Conan。
- Java 后端：JDK、Maven 或 Gradle、Spring Boot CLI、容器工具。
