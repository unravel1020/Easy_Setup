# 环境目录设计

## 目标

环境目录是 GUI 中“可勾选环境”的数据来源。用户不需要记住包名、安装命令和验证命令，只需要选择语言、框架、AI 栈或团队模板，EnvForge 会展开为具体安装计划。

## 分层模型

```text
Catalog Item
└── Stack Template
    └── Recipe
```

- Catalog Item：界面展示层，包含名称、分类、说明、标签、推荐度、体积、预计耗时。
- Stack Template：组合层，把多个 recipe 组织成一个可安装环境栈。
- Recipe：执行层，定义检测、安装、配置、验证和回滚。

## 主流语言范围

第一阶段环境目录应覆盖：

- C/C++：GCC、Clang、MSVC Build Tools、CMake、Ninja、Make、GDB/LLDB、vcpkg、Conan。
- Java：JDK、Maven、Gradle、Spring Boot、Quarkus、Micronaut。
- JavaScript/TypeScript：Node.js、pnpm、npm、Yarn、Bun、Vite、React、Vue、Angular、Svelte、Next.js、Nuxt、NestJS。
- Python：CPython、uv、pipx、virtualenv、Poetry、Django、FastAPI、Flask、Jupyter。
- Go：Go toolchain、gofmt、golangci-lint、Gin、Fiber。
- Rust：rustup、stable toolchain、cargo、clippy、rustfmt、Actix、Axum、Tauri CLI。
- C#/.NET：.NET SDK、ASP.NET Core、Entity Framework CLI。
- PHP：PHP、Composer、Laravel、Symfony。
- Ruby：Ruby、Bundler、Rails。
- Kotlin：JDK、Gradle、Kotlin compiler、Ktor。
- Swift：Swift toolchain，macOS 上支持 Xcode Command Line Tools。
- Dart/Flutter：Dart SDK、Flutter SDK、Android SDK 依赖。
- R / Julia：数据分析与科学计算基础工具链。

## AI/ML 范围

AI 开发目录应覆盖：

- 基础环境：Python、uv、Jupyter、CUDA 检测、GPU 驱动状态检测。
- 深度学习：PyTorch、TensorFlow、JAX。
- 模型生态：Transformers、Datasets、Tokenizers、Accelerate、Diffusers。
- LLM 应用：LangChain、LlamaIndex、OpenAI SDK、Anthropic SDK、Ollama。
- 推理与部署：ONNX Runtime、TensorRT 检测、vLLM、llama.cpp。
- 数据工具：NumPy、Pandas、Polars、Scikit-learn、Matplotlib。

AI 栈需要支持变体选择：

- CPU only。
- NVIDIA CUDA。
- Apple Silicon。
- AMD ROCm。

不同平台无法支持的变体应在 GUI 中明确置灰并说明原因。

## 推荐模板

### AI Python 工作站

包含：

- Python 3.12。
- uv。
- Jupyter。
- PyTorch。
- Transformers。
- LangChain。
- LlamaIndex。
- NumPy、Pandas、Scikit-learn。

验证：

- Python 版本。
- 虚拟环境创建。
- `import torch`。
- `import transformers`。
- `import langchain`。

### C/C++ 原生开发

包含：

- 平台编译器。
- CMake。
- Ninja。
- 调试器。
- vcpkg 或 Conan。

验证：

- 编译 hello world。
- 运行可执行文件。
- 输出编译器版本。

### Java 后端

包含：

- JDK。
- Maven 或 Gradle。
- Spring Boot。
- Docker 可选。

验证：

- `java -version`。
- 构建最小 Spring Boot 项目。
- 运行测试命令。

### 全栈 Web

包含：

- Git。
- Node.js。
- pnpm。
- TypeScript。
- Vite。
- React 或 Vue。
- 后端框架可选。

验证：

- 包管理器版本。
- 创建临时项目。
- 安装依赖 dry-run 或锁文件验证。

## GUI 选择规则

- 用户可以从模板开始，也可以逐项勾选。
- 勾选项必须能展示依赖关系，例如选择 LangChain 会自动建议 Python 和 uv。
- 冲突项必须提前提示，例如多个 Node 版本管理器同时启用。
- 大型组件必须展示预计下载体积和耗时，尤其是 CUDA、Android SDK、Docker Desktop。
- 一键执行前必须展示最终计划。

## 可审计输出

用户在 GUI 中勾选的结果应能保存为 `*.envforge.yaml`，便于团队审查和复用。
