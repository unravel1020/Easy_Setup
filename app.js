const translations = {
  en: {
    brandSubtitle: "Environment setup",
    searchLabel: "Search",
    eyebrow: "MVP planner",
    headline: "Pick a stack. Learn the basics. Review the plan.",
    aiTemplate: "AI Python",
    webTemplate: "Fullstack",
    catalogTitle: "Catalog",
    planTitle: "Plan",
    planSubtitle: "Dry-run preview",
    clear: "Clear",
    selected: "selected",
    actions: "actions",
    checks: "checks",
    actionsTitle: "Actions",
    environmentTitle: "Environment",
    cliPreviewTitle: "CLI Preview",
    jobsTitle: "Jobs",
    noJobs: "No jobs yet.",
    jobsOffline: "Start a local Agent to view job history.",
    jobScript: "Script",
    jobLog: "Log",
    jobExitCode: "Exit code",
    jobCompletedAt: "Completed",
    viewLog: "View Log",
    hideLog: "Hide Log",
    cancelJob: "Cancel",
    cancelConfirm: "Request cancellation for this job?",
    retryJob: "Retry",
    emptyLog: "Log is empty.",
    emptyPlan: "Select a stack to generate an install plan.",
    itemCount: "items",
    themeDark: "Dark",
    themeLight: "Light",
    languageToggle: "中文",
    installSelected: "Execute Install",
    installOne: "Install",
    copyCli: "Copy",
    copied: "Copied.",
    agentOnline: "Agent online",
    agentOffline: "Agent offline",
    agentNoExecute: "Agent dry-run",
    executeConfirm: "Execute selected install commands in a local PowerShell window?",
    official: "Official",
    hide: "Hide",
    menu: "Menu",
    resources: "Start here",
    useCase: "Use",
    categories: {
      all: "All",
      templates: "Templates",
      languages: "Languages",
      frontend: "Frontend",
      backend: "Backend",
      ai: "AI/ML",
      testing: "Testing",
      operations: "Operations",
      skills: "Skills",
      vibe: "Vibe Coding",
      harness: "Harness",
      linux: "Linux",
      virtualization: "Virtualization",
      devops: "DevOps"
    }
  },
  zh: {
    brandSubtitle: "开发环境配置",
    searchLabel: "搜索",
    eyebrow: "MVP 规划器",
    headline: "选择环境栈，了解入门资料，审阅安装计划。",
    aiTemplate: "AI Python",
    webTemplate: "全栈 Web",
    catalogTitle: "环境目录",
    planTitle: "计划",
    planSubtitle: "安全预览，不会安装",
    clear: "清空",
    selected: "已选",
    actions: "操作",
    checks: "验证",
    actionsTitle: "操作",
    environmentTitle: "环境变量",
    cliPreviewTitle: "CLI 预览",
    jobsTitle: "执行历史",
    noJobs: "暂无执行记录。",
    jobsOffline: "启动本地 Agent 后可查看执行历史。",
    jobScript: "脚本",
    jobLog: "日志",
    jobExitCode: "退出码",
    jobCompletedAt: "完成时间",
    viewLog: "查看日志",
    hideLog: "收起日志",
    cancelJob: "取消",
    cancelConfirm: "是否请求取消这个任务？",
    retryJob: "重试",
    emptyLog: "日志为空。",
    emptyPlan: "选择一个环境栈后生成安装计划。",
    itemCount: "项",
    themeDark: "深色",
    themeLight: "浅色",
    languageToggle: "EN",
    installSelected: "执行安装",
    installOne: "安装",
    copyCli: "复制",
    copied: "已复制。",
    agentOnline: "Agent 在线",
    agentOffline: "Agent 离线",
    agentNoExecute: "Agent 预览",
    executeConfirm: "是否在本机 PowerShell 窗口中执行所选安装命令？",
    official: "官网",
    hide: "收起",
    menu: "菜单",
    resources: "入门资料",
    useCase: "用途",
    categories: {
      all: "全部",
      templates: "模板",
      languages: "语言",
      frontend: "前端",
      backend: "后端",
      ai: "AI/ML",
      testing: "测试",
      operations: "运维",
      skills: "技能",
      vibe: "Vibe Coding",
      harness: "Harness",
      linux: "Linux",
      virtualization: "虚拟机",
      devops: "DevOps"
    }
  }
};

let catalog = {
  categories: ["all", "templates", "languages", "frontend", "backend", "ai", "testing", "operations", "skills", "vibe", "harness", "linux", "virtualization", "devops"],
  items: [
    item("template.ai-python-workstation", "templates", ["git", "python", "notebook.jupyter", "ai.pytorch", "ai.transformers", "ai.langchain", "ai.llamaindex"], {
      en: ["AI Python Workstation", "Python, notebooks and LLM app tooling.", "AI app prototypes, RAG demos, model experiments.", ["Python basics", "PyTorch 60-minute blitz", "LangChain quickstart"]],
      zh: ["AI Python 工作站", "Python、Notebook 和 LLM 应用工具。", "AI 应用原型、RAG 演示、模型实验。", ["Python 入门", "PyTorch 60 分钟教程", "LangChain 快速开始"]]
    }),
    item("template.fullstack-web", "templates", ["git", "node", "pnpm", "frontend.react", "frontend.vite", "python", "backend.fastapi"], {
      en: ["Fullstack Web", "React, Vite and FastAPI baseline.", "Small SaaS, admin tools, web prototypes.", ["MDN Web Docs", "React learn", "FastAPI tutorial"]],
      zh: ["全栈 Web", "React、Vite 和 FastAPI 基线。", "小型 SaaS、管理后台、Web 原型。", ["MDN Web 文档", "React 官方教程", "FastAPI 教程"]]
    }),
    item("language.go", "languages", ["go"], {
      en: ["Go / Golang", "Go compiler, formatter and module workflow.", "CLI tools, services, infrastructure agents.", ["A Tour of Go", "Effective Go", "Go by Example"]],
      zh: ["Go / Golang", "Go 编译器、格式化和模块工作流。", "CLI 工具、服务端、基础设施代理。", ["Go 指南", "Effective Go", "Go by Example"]]
    }),
    item("language.cpp", "languages", ["cpp"], {
      en: ["C/C++ Native", "Compiler, CMake, Ninja and debugger.", "Native libraries, performance tools, embedded code.", ["CMake tutorial", "Modern C++ tour", "vcpkg docs"]],
      zh: ["C/C++ 原生开发", "编译器、CMake、Ninja 和调试器。", "原生库、性能工具、嵌入式代码。", ["CMake 教程", "现代 C++ 导览", "vcpkg 文档"]]
    }),
    item("language.java", "languages", ["java"], {
      en: ["Java JDK", "JDK baseline for JVM services.", "Enterprise backend and Android-adjacent tooling.", ["Java tutorials", "Maven guide", "Gradle basics"]],
      zh: ["Java JDK", "JVM 服务开发的 JDK 基线。", "企业后端和 Android 相关工具。", ["Java 教程", "Maven 指南", "Gradle 基础"]]
    }),
    item("language.python", "languages", ["python"], {
      en: ["Python", "Runtime and package tooling.", "Automation, data, AI and backend scripting.", ["Python tutorial", "uv guide", "Packaging guide"]],
      zh: ["Python", "运行时和包管理工具。", "自动化、数据、AI 和后端脚本。", ["Python 教程", "uv 指南", "打包指南"]]
    }),
    item("language.rust", "languages", ["rust"], {
      en: ["Rust", "Compiler, cargo and formatters.", "Fast CLIs, systems code, web services.", ["Rust Book", "Rustlings", "Cargo guide"]],
      zh: ["Rust", "编译器、cargo 和格式化工具。", "高性能 CLI、系统代码、Web 服务。", ["Rust 程序设计语言", "Rustlings", "Cargo 指南"]]
    }),
    item("frontend.react", "frontend", ["node", "pnpm", "frontend.react"], {
      en: ["React", "React development baseline.", "Component apps and dashboards.", ["React learn", "Vite guide", "TanStack Query"]],
      zh: ["React", "React 开发基线。", "组件应用和管理看板。", ["React 官方教程", "Vite 指南", "TanStack Query"]]
    }),
    item("frontend.vue", "frontend", ["node", "pnpm", "frontend.vue"], {
      en: ["Vue", "Vue and Vite tooling.", "Progressive web apps and internal tools.", ["Vue guide", "Pinia", "Vue Router"]],
      zh: ["Vue", "Vue 与 Vite 工具链。", "渐进式 Web 应用和内部工具。", ["Vue 指南", "Pinia", "Vue Router"]]
    }),
    item("frontend.electron", "frontend", ["node", "pnpm", "frontend.electron"], {
      en: ["Electron", "Desktop app shell for web frontends.", "Cross-platform desktop tools.", ["Electron docs", "Electron Forge", "Security checklist"]],
      zh: ["Electron", "面向 Web 前端的桌面应用壳。", "跨平台桌面工具。", ["Electron 文档", "Electron Forge", "安全清单"]]
    }),
    item("backend.spring", "backend", ["java", "backend.spring"], {
      en: ["Spring Series", "Spring Boot, Maven or Gradle workflow.", "Java services, APIs and enterprise apps.", ["Spring guides", "Spring Boot docs", "Spring Security"]],
      zh: ["Spring 系列", "Spring Boot、Maven 或 Gradle 工作流。", "Java 服务、API 和企业应用。", ["Spring 指南", "Spring Boot 文档", "Spring Security"]]
    }),
    item("backend.go-web", "backend", ["go", "backend.gin", "backend.fiber", "backend.echo"], {
      en: ["Go Web Frameworks", "Gin, Fiber and Echo quick-start stack.", "Fast APIs and small cloud services.", ["Gin docs", "Fiber docs", "Echo cookbook"]],
      zh: ["Go Web 框架", "Gin、Fiber 和 Echo 快速开发栈。", "高性能 API 和小型云服务。", ["Gin 文档", "Fiber 文档", "Echo 手册"]]
    }),
    item("backend.rust-web", "backend", ["rust", "backend.axum", "backend.actix"], {
      en: ["Rust Web Frameworks", "Axum and Actix web service stack.", "Safe high-performance APIs.", ["Axum examples", "Actix guide", "Tokio tutorial"]],
      zh: ["Rust Web 框架", "Axum 和 Actix 服务端开发栈。", "安全高性能 API。", ["Axum 示例", "Actix 指南", "Tokio 教程"]]
    }),
    item("backend.fastapi", "backend", ["python", "backend.fastapi"], {
      en: ["FastAPI", "Python API stack.", "Typed APIs and AI service backends.", ["FastAPI tutorial", "Pydantic", "Uvicorn"]],
      zh: ["FastAPI", "Python API 开发栈。", "类型化 API 和 AI 服务后端。", ["FastAPI 教程", "Pydantic", "Uvicorn"]]
    }),
    item("ai.pytorch", "ai", ["python", "ai.pytorch"], {
      en: ["PyTorch CPU", "Deep learning CPU stack.", "Model experiments without GPU setup.", ["PyTorch tutorials", "TorchVision", "Accelerate"]],
      zh: ["PyTorch CPU", "CPU 深度学习开发栈。", "无需 GPU 配置的模型实验。", ["PyTorch 教程", "TorchVision", "Accelerate"]]
    }),
    item("ai.langchain", "ai", ["python", "ai.langchain", "ai.llamaindex"], {
      en: ["LLM App Stack", "LangChain and LlamaIndex.", "RAG, agents and document assistants.", ["LangChain quickstart", "LlamaIndex starter", "RAG patterns"]],
      zh: ["LLM 应用栈", "LangChain 和 LlamaIndex。", "RAG、Agent 和文档助手。", ["LangChain 快速开始", "LlamaIndex 入门", "RAG 模式"]]
    }),
    item("testing.web", "testing", ["node", "testing.playwright"], {
      en: ["Web Testing", "Playwright browser tests.", "End-to-end UI test automation.", ["Playwright intro", "Selectors", "Trace viewer"]],
      zh: ["Web 测试", "Playwright 浏览器测试。", "端到端 UI 自动化测试。", ["Playwright 入门", "选择器", "Trace Viewer"]]
    }),
    item("testing.python", "testing", ["python", "testing.pytest"], {
      en: ["Python Testing", "pytest for Python projects.", "Unit tests and API test suites.", ["pytest docs", "fixtures", "coverage.py"]],
      zh: ["Python 测试", "用于 Python 项目的 pytest。", "单元测试和 API 测试套件。", ["pytest 文档", "fixtures", "coverage.py"]]
    }),
    item("testing.postman", "testing", ["testing.postman"], {
      en: ["Postman", "API testing client.", "Manual and collection-based API checks.", ["Collections", "Environments", "Newman CLI"]],
      zh: ["Postman", "API 测试客户端。", "手动和集合式 API 检查。", ["Collections", "Environments", "Newman CLI"]]
    }),
    item("operations.kubernetes", "operations", ["kubernetes.kubectl", "helm"], {
      en: ["Kubernetes Ops", "kubectl and Helm.", "Cluster inspection and chart deployment.", ["kubectl cheatsheet", "Helm charts", "K9s"]],
      zh: ["Kubernetes 运维", "kubectl 和 Helm。", "集群查看和 Chart 部署。", ["kubectl 速查", "Helm Charts", "K9s"]]
    }),
    item("operations.infra", "operations", ["terraform", "ansible"], {
      en: ["Infrastructure Ops", "Terraform and Ansible.", "Provisioning and configuration automation.", ["Terraform tutorials", "Ansible guide", "IaC patterns"]],
      zh: ["基础设施运维", "Terraform 和 Ansible。", "资源编排和配置自动化。", ["Terraform 教程", "Ansible 指南", "IaC 模式"]]
    }),
    item("skills.codex", "skills", ["skill.codex"], {
      en: ["Codex Skills", "Reusable workflows and local instructions.", "Team playbooks for repeatable AI-assisted work.", ["Skill design", "Prompt patterns", "Review checklists"]],
      zh: ["Codex Skills", "可复用工作流和本地指令。", "团队可复用的 AI 协作手册。", ["Skill 设计", "提示词模式", "审查清单"]]
    }),
    item("vibe.webapp", "vibe", ["vibe.prototype"], {
      en: ["Vibe Coding Kit", "Idea-to-prototype web app stack.", "Fast UI experiments and product spikes.", ["PRD prompts", "UI critique loop", "Ship checklist"]],
      zh: ["Vibe Coding 套件", "从想法到原型的 Web 应用栈。", "快速 UI 实验和产品验证。", ["PRD 提示词", "UI 评审循环", "发布清单"]]
    }),
    item("harness.delivery", "harness", ["harness.cli"], {
      en: ["Harness Delivery", "Harness CLI and CI/CD setup notes.", "Pipelines, feature flags and deployment checks.", ["Harness docs", "Pipeline basics", "Feature flags"]],
      zh: ["Harness 交付", "Harness CLI 和 CI/CD 配置说明。", "流水线、功能开关和部署检查。", ["Harness 文档", "流水线基础", "功能开关"]]
    }),
    item("linux.mirrors", "linux", ["linux.mirrors"], {
      en: ["Linux Mirrors", "Apt, dnf, pacman and pip/npm mirror presets.", "Faster installs in regional or intranet networks.", ["apt sources", "pip config", "npm registry"]],
      zh: ["Linux 镜像源", "apt、dnf、pacman、pip/npm 镜像预设。", "区域网络或内网中更快安装。", ["apt sources", "pip config", "npm registry"]]
    }),
    item("virtualization.vm", "virtualization", ["vm.virtualbox", "vm.vagrant"], {
      en: ["Virtual Machines", "VirtualBox and Vagrant baseline.", "Disposable Linux labs and isolated dev boxes.", ["Vagrant boxes", "VirtualBox networking", "WSL comparison"]],
      zh: ["虚拟机", "VirtualBox 和 Vagrant 基线。", "一次性 Linux 实验环境和隔离开发机。", ["Vagrant boxes", "VirtualBox 网络", "WSL 对比"]]
    }),
    item("git", "devops", ["git"], {
      en: ["Git", "Version control baseline.", "Every project and team workflow.", ["Git book", "Branching", "Conventional commits"]],
      zh: ["Git", "版本控制基线。", "所有项目和团队协作流程。", ["Git Book", "分支模型", "约定式提交"]]
    }),
    item("docker", "devops", ["docker"], {
      en: ["Docker", "Container runtime.", "Local services and reproducible app runtime.", ["Docker getting started", "Compose", "Dockerfile best practices"]],
      zh: ["Docker", "容器运行时。", "本地服务和可复现应用运行时。", ["Docker 入门", "Compose", "Dockerfile 最佳实践"]]
    })
  ],
  recipes: {
    git: recipe("Git", "winget install --id Git.Git -e", "git --version"),
    node: recipe("Node.js", "winget install --id OpenJS.NodeJS.LTS -e", "node --version"),
    pnpm: recipe("pnpm", "corepack enable", "pnpm --version"),
    python: recipe("Python", "winget install --id Python.Python.3.12 -e", "python --version"),
    go: recipe("Go", "winget install --id GoLang.Go -e", "go version"),
    rust: recipe("Rust", "winget install --id Rustlang.Rustup -e", "rustc --version"),
    java: recipe("Java JDK", "winget install --id EclipseAdoptium.Temurin.21.JDK -e", "java -version"),
    cpp: recipe("C/C++ Toolchain", "winget install --id Kitware.CMake -e", "cmake --version"),
    docker: recipe("Docker", "winget install --id Docker.DockerDesktop -e", "docker --version"),
    "frontend.react": recipe("React", "npm create vite@latest envforge-react-smoke -- --template react-ts", "node --version"),
    "frontend.vue": recipe("Vue", "npm create vue@latest", "node --version"),
    "frontend.electron": recipe("Electron", "npm create electron-app@latest envforge-electron-smoke", "node --version"),
    "backend.spring": recipe("Spring Boot", "winget install --id SpringSource.SpringToolSuite.4 -e", "java -version"),
    "backend.gin": recipe("Gin", "go install github.com/gin-gonic/gin@latest", "go version"),
    "backend.fiber": recipe("Fiber", "go env GOPATH", "go version"),
    "backend.echo": recipe("Echo", "go env GOPATH", "go version"),
    "backend.axum": recipe("Axum", "cargo search axum", "cargo --version"),
    "backend.actix": recipe("Actix", "cargo search actix-web", "cargo --version"),
    "backend.fastapi": recipe("FastAPI", "python -m pip install fastapi uvicorn", "python -c \"import fastapi; print(fastapi.__version__)\""),
    "ai.pytorch": recipe("PyTorch", "python -m pip install torch --index-url https://download.pytorch.org/whl/cpu", "python -c \"import torch; print(torch.__version__)\""),
    "ai.transformers": recipe("Transformers", "python -m pip install transformers datasets accelerate", "python -c \"import transformers; print(transformers.__version__)\""),
    "ai.langchain": recipe("LangChain", "python -m pip install langchain", "python -c \"import langchain; print(langchain.__version__)\""),
    "ai.llamaindex": recipe("LlamaIndex", "python -m pip install llama-index", "python -c \"import llama_index; print('llama-index ok')\""),
    "notebook.jupyter": recipe("Jupyter", "python -m pip install jupyterlab", "jupyter --version"),
    "testing.playwright": recipe("Playwright", "npm init playwright@latest", "npx playwright --version"),
    "testing.pytest": recipe("pytest", "python -m pip install pytest", "python -m pytest --version"),
    "testing.postman": recipe("Postman", "winget install --id Postman.Postman -e", "postman --version"),
    "kubernetes.kubectl": recipe("kubectl", "winget install --id Kubernetes.kubectl -e", "kubectl version --client"),
    helm: recipe("Helm", "winget install --id Helm.Helm -e", "helm version"),
    terraform: recipe("Terraform", "winget install --id Hashicorp.Terraform -e", "terraform version"),
    ansible: recipe("Ansible", "python -m pip install ansible", "ansible --version"),
    "skill.codex": recipe("Codex Skills", "envforge notes skill.codex", "git --version"),
    "vibe.prototype": recipe("Vibe Prototype", "envforge notes vibe.prototype", "git --version"),
    "harness.cli": recipe("Harness CLI", "winget install --id Harness.HarnessCLI -e", "harness --version"),
    "linux.mirrors": recipe("Linux Mirrors", "envforge configure mirrors --interactive", "git --version"),
    "vm.virtualbox": recipe("VirtualBox", "winget install --id Oracle.VirtualBox -e", "VBoxManage --version"),
    "vm.vagrant": recipe("Vagrant", "winget install --id Hashicorp.Vagrant -e", "vagrant --version")
  }
};

let officialLinks = {
  "template.ai-python-workstation": "https://www.python.org/",
  "template.fullstack-web": "https://developer.mozilla.org/",
  "language.go": "https://go.dev/",
  "language.cpp": "https://isocpp.org/",
  "language.java": "https://dev.java/",
  "language.python": "https://www.python.org/",
  "language.rust": "https://www.rust-lang.org/",
  "frontend.react": "https://react.dev/",
  "frontend.vue": "https://vuejs.org/",
  "frontend.electron": "https://www.electronjs.org/",
  "backend.spring": "https://spring.io/projects/spring-boot",
  "backend.go-web": "https://gin-gonic.com/",
  "backend.rust-web": "https://github.com/tokio-rs/axum",
  "backend.fastapi": "https://fastapi.tiangolo.com/",
  "ai.pytorch": "https://pytorch.org/",
  "ai.langchain": "https://www.langchain.com/",
  "testing.web": "https://playwright.dev/",
  "testing.python": "https://docs.pytest.org/",
  "testing.postman": "https://www.postman.com/",
  "operations.kubernetes": "https://kubernetes.io/",
  "operations.infra": "https://www.terraform.io/",
  "skills.codex": "https://developers.openai.com/",
  "vibe.webapp": "https://developer.mozilla.org/",
  "harness.delivery": "https://www.harness.io/",
  "linux.mirrors": "https://repogen.simplylinux.ch/",
  "virtualization.vm": "https://www.virtualbox.org/",
  git: "https://git-scm.com/",
  docker: "https://www.docker.com/"
};

let resourceLinks = {
  "Python basics": "https://docs.python.org/3/tutorial/",
  "Python 入门": "https://docs.python.org/zh-cn/3/tutorial/",
  "PyTorch 60-minute blitz": "https://pytorch.org/tutorials/beginner/deep_learning_60min_blitz.html",
  "PyTorch 60 分钟教程": "https://pytorch.org/tutorials/beginner/deep_learning_60min_blitz.html",
  "LangChain quickstart": "https://python.langchain.com/docs/get_started/quickstart/",
  "LangChain 快速开始": "https://python.langchain.com/docs/get_started/quickstart/",
  "MDN Web Docs": "https://developer.mozilla.org/",
  "MDN Web 文档": "https://developer.mozilla.org/zh-CN/",
  "React learn": "https://react.dev/learn",
  "React 官方教程": "https://react.dev/learn",
  "FastAPI tutorial": "https://fastapi.tiangolo.com/tutorial/",
  "FastAPI 教程": "https://fastapi.tiangolo.com/tutorial/",
  "A Tour of Go": "https://go.dev/tour/",
  "Go 指南": "https://go.dev/tour/",
  "Effective Go": "https://go.dev/doc/effective_go",
  "Go by Example": "https://gobyexample.com/",
  "CMake tutorial": "https://cmake.org/cmake/help/latest/guide/tutorial/",
  "CMake 教程": "https://cmake.org/cmake/help/latest/guide/tutorial/",
  "vcpkg docs": "https://learn.microsoft.com/vcpkg/",
  "vcpkg 文档": "https://learn.microsoft.com/vcpkg/",
  "Java tutorials": "https://dev.java/learn/",
  "Java 教程": "https://dev.java/learn/",
  "Vue guide": "https://vuejs.org/guide/",
  "Vue 指南": "https://cn.vuejs.org/guide/",
  "Electron docs": "https://www.electronjs.org/docs/latest/",
  "Electron 文档": "https://www.electronjs.org/docs/latest/",
  "Spring guides": "https://spring.io/guides",
  "Spring 指南": "https://spring.io/guides",
  "Gin docs": "https://gin-gonic.com/docs/",
  "Gin 文档": "https://gin-gonic.com/docs/",
  "Axum examples": "https://github.com/tokio-rs/axum/tree/main/examples",
  "Axum 示例": "https://github.com/tokio-rs/axum/tree/main/examples",
  "Playwright intro": "https://playwright.dev/docs/intro",
  "Playwright 入门": "https://playwright.dev/docs/intro",
  "pytest docs": "https://docs.pytest.org/",
  "pytest 文档": "https://docs.pytest.org/",
  "Harness docs": "https://developer.harness.io/docs/",
  "Harness 文档": "https://developer.harness.io/docs/",
  "Terraform tutorials": "https://developer.hashicorp.com/terraform/tutorials",
  "Terraform 教程": "https://developer.hashicorp.com/terraform/tutorials",
  "Docker getting started": "https://docs.docker.com/get-started/",
  "Docker 入门": "https://docs.docker.com/get-started/"
};

let iconPaths = {
  templates: "M4 5h16v4H4V5Zm0 6h7v8H4v-8Zm9 0h7v8h-7v-8Z",
  languages: "M8 4 3 9l5 5 1.4-1.4L5.8 9 9.4 5.4 8 4Zm8 0-1.4 1.4L18.2 9l-3.6 3.6L16 14l5-5-5-5ZM11.2 20l4-16h-2.1l-4 16h2.1Z",
  frontend: "M4 5h16v14H4V5Zm2 3v9h12V8H6Zm2 2h5v2H8v-2Z",
  backend: "M6 4h12v4H6V4Zm0 6h12v4H6v-4Zm0 6h12v4H6v-4Zm2-10v1h2V6H8Zm0 6v1h2v-1H8Zm0 6v1h2v-1H8Z",
  ai: "M12 3a5 5 0 0 1 5 5v1h1a3 3 0 0 1 0 6h-1v1a5 5 0 0 1-10 0v-1H6a3 3 0 0 1 0-6h1V8a5 5 0 0 1 5-5Zm-2 6h4v2h-4V9Zm0 4h4v2h-4v-2Z",
  testing: "M9 3h6v2h-1v5.6l4.4 6.6A2 2 0 0 1 16.7 20H7.3a2 2 0 0 1-1.7-2.8l4.4-6.6V5H9V3Zm3 8.2L7.3 18h9.4L12 11.2Z",
  operations: "M12 2 3 7v10l9 5 9-5V7l-9-5Zm0 2.3L18.6 8 12 11.7 5.4 8 12 4.3ZM5 9.7l6 3.4v6.2l-6-3.4V9.7Zm14 0v6.2l-6 3.4v-6.2l6-3.4Z",
  skills: "M4 4h7v7H4V4Zm9 0h7v7h-7V4ZM4 13h7v7H4v-7Zm9 0h7v7h-7v-7Z",
  vibe: "M12 2l2.2 6.8H21l-5.5 4 2.1 6.8-5.6-4.1-5.6 4.1 2.1-6.8L3 8.8h6.8L12 2Z",
  harness: "M5 4h14v4H5V4Zm2 6h10v4H7v-4Zm-2 6h14v4H5v-4Z",
  linux: "M12 3c3 0 5 2.5 5 6v2l2 4v4H5v-4l2-4V9c0-3.5 2-6 5-6Zm-2 7h4V8h-4v2Z",
  virtualization: "M4 5h16v10H4V5Zm2 2v6h12V7H6Zm2 10h8v2H8v-2Z",
  devops: "M8 7a4 4 0 0 1 7.5-2H17a3 3 0 0 1 0 6h-1.5A4 4 0 0 1 8 9H7a3 3 0 0 0 0 6h1.5a4 4 0 0 0 7.5 2h1a3 3 0 0 0 0-6h-1V9h1a5 5 0 0 1 0 10h-1.5A6 6 0 0 1 6.5 17H7a5 5 0 0 1 0-10h1Z"
};

function item(id, category, recipes, text) {
  return { id, category, recipes, text };
}

function recipe(name, command, verify) {
  return { name, command, verify };
}

const state = {
  category: "templates",
  query: "",
  selected: new Set(),
  jobs: [],
  jobLogs: {},
  expandedLogs: new Set(),
  theme: localStorage.getItem("envforge-theme") || "light",
  language: localStorage.getItem("envforge-language") || "en",
  sidebarCollapsed: localStorage.getItem("envforge-sidebar") === "collapsed",
  agent: {
    online: false,
    allowExecute: false,
    url: "http://127.0.0.1:17771",
    urls: ["http://127.0.0.1:17771", "http://127.0.0.1:17772"]
  }
};

const byId = (id) => document.getElementById(id);
const t = () => translations[state.language];

function normalizeWebCatalog(data) {
  if (!data || !data.catalog || !Array.isArray(data.catalog.items)) {
    throw new Error("Invalid web catalog");
  }
  catalog = data.catalog;
  officialLinks = data.officialLinks || officialLinks;
  resourceLinks = data.resourceLinks || resourceLinks;
  iconPaths = data.iconPaths || iconPaths;
}

async function loadWebCatalog() {
  try {
    const response = await fetch("./catalog.web.json", { cache: "no-store" });
    if (!response.ok) throw new Error("Catalog request failed");
    normalizeWebCatalog(await response.json());
  } catch {
    // 某些浏览器会阻止 file:// 读取 JSON；内置目录用于保证本地 MVP 仍可运行。
  }
}

function getItemText(item) {
  const fallback = item.text.en;
  const value = item.text[state.language] || fallback;
  return { name: value[0], description: value[1], use: value[2], resources: value[3] || [] };
}

function iconFor(category) {
  const path = iconPaths[category] || iconPaths.templates;
  return `<svg class="icon" viewBox="0 0 24 24" aria-hidden="true"><path fill="currentColor" d="${path}"></path></svg>`;
}

function linkForResource(resource) {
  const href = resourceLinks[resource] || officialLinks[resource] || "https://www.google.com/search?q=" + encodeURIComponent(resource);
  return `<a href="${href}" target="_blank" rel="noopener noreferrer">${resource}</a>`;
}

function applyTranslations() {
  document.documentElement.lang = state.language === "zh" ? "zh-CN" : "en";
  document.querySelectorAll("[data-i18n]").forEach((node) => {
    node.textContent = t()[node.dataset.i18n] || node.textContent;
  });
  byId("languageToggle").querySelector("span").textContent = t().languageToggle;
  byId("search").placeholder = state.language === "zh" ? "Python、Playwright、Terraform、Spring..." : "Python, Playwright, Terraform, Spring...";
}

function applyShellState() {
  document.documentElement.dataset.theme = state.theme;
  byId("themeToggle").querySelector("span").textContent = state.theme === "dark" ? t().themeLight : t().themeDark;
  byId("appShell").classList.toggle("sidebar-collapsed", state.sidebarCollapsed);
  applyTranslations();
  renderAgentStatus();
}

function renderAgentStatus() {
  const node = byId("agentStatus");
  if (!node) return;
  node.classList.toggle("online", state.agent.online);
  node.classList.toggle("offline", !state.agent.online);
  if (!state.agent.online) {
    node.textContent = t().agentOffline || "Agent offline";
    return;
  }
  node.textContent = state.agent.allowExecute
    ? (t().agentOnline || "Agent online")
    : (t().agentNoExecute || "Agent dry-run");
}

async function refreshJobs() {
  if (!state.agent.online) {
    state.jobs = [];
    renderJobs();
    return;
  }
  try {
    const response = await fetch(`${state.agent.url}/jobs`, { cache: "no-store" });
    if (!response.ok) throw new Error("Jobs unavailable");
    const data = await response.json();
    state.jobs = Array.isArray(data.jobs) ? data.jobs : [];
    await refreshExpandedJobLogs();
  } catch {
    state.jobs = [];
  }
  renderJobs();
}

async function refreshAgentStatus() {
  for (const url of state.agent.urls) {
    try {
      const response = await fetch(`${url}/health`, { cache: "no-store" });
      if (!response.ok) throw new Error("Agent unavailable");
      const data = await response.json();
      state.agent.url = url;
      state.agent.online = Boolean(data.ok);
      state.agent.allowExecute = Boolean(data.allowExecute);
      renderAgentStatus();
      refreshJobs();
      return;
    } catch {
      state.agent.online = false;
      state.agent.allowExecute = false;
    }
  }
  renderAgentStatus();
  refreshJobs();
}

function renderCategories() {
  byId("categories").innerHTML = catalog.categories.map((category) => `
    <button class="category ${state.category === category ? "active" : ""}" data-category="${category}" type="button">
      ${t().categories[category]}
    </button>
  `).join("");

  document.querySelectorAll("[data-category]").forEach((button) => {
    button.addEventListener("click", () => {
      state.category = button.dataset.category;
      render();
    });
  });
}

function getVisibleItems() {
  const query = state.query.trim().toLowerCase();
  return catalog.items.filter((item) => {
    const text = getItemText(item);
    const resourceText = text.resources.join(" ");
    const matchesCategory = state.category === "all" || item.category === state.category;
    const matchesQuery = !query || `${text.name} ${text.description} ${text.use} ${resourceText} ${item.id}`.toLowerCase().includes(query);
    return matchesCategory && matchesQuery;
  });
}

function renderCatalog() {
  const items = getVisibleItems();
  byId("catalogCount").textContent = `${items.length} ${t().itemCount}`;
  byId("catalog").innerHTML = items.map((item) => {
    const text = getItemText(item);
    const resources = text.resources.slice(0, 3).map(linkForResource).join("");
    const site = officialLinks[item.id] ? `<a class="site-link" href="${officialLinks[item.id]}" target="_blank" rel="noopener noreferrer">${t().official}</a>` : "";
    return `
      <div class="item ${state.selected.has(item.id) ? "selected" : ""}" data-row="${item.id}">
        <input type="checkbox" data-item="${item.id}" ${state.selected.has(item.id) ? "checked" : ""} />
        ${iconFor(item.category)}
        <div class="item-body">
          <div class="item-title">
            <div class="title-meta">
              <h4>${text.name}</h4>
              <span class="category-label">${t().categories[item.category]}</span>
            </div>
            <div class="item-actions">
              ${site}
              <button class="install-one" type="button" data-install="${item.id}">${t().installOne}</button>
            </div>
          </div>
          <p>${text.description}</p>
          <div class="item-meta"><strong>${t().useCase}</strong><span>${text.use}</span></div>
          <div class="resources"><strong>${t().resources}</strong>${resources}</div>
        </div>
      </div>
    `;
  }).join("");

  document.querySelectorAll("[data-row]").forEach((row) => {
    row.addEventListener("click", (event) => {
      if (event.target.closest("a, button, input")) return;
      const input = row.querySelector("[data-item]");
      input.checked = !input.checked;
      input.dispatchEvent(new Event("change", { bubbles: true }));
    });
  });

  document.querySelectorAll("[data-item]").forEach((input) => {
    input.addEventListener("change", () => {
      if (input.checked) state.selected.add(input.dataset.item);
      else state.selected.delete(input.dataset.item);
      renderPlan();
      renderCatalog();
    });
  });

  document.querySelectorAll("[data-install]").forEach((button) => {
    button.addEventListener("click", (event) => {
      event.preventDefault();
      event.stopPropagation();
      const item = catalog.items.find((candidate) => candidate.id === button.dataset.install);
      if (!item) return;
      executeViaAgent([item.id], button);
    });
  });
}

function getRecipesForItems(itemIds) {
  const recipeIds = new Set();
  itemIds.forEach((itemId) => {
    const item = catalog.items.find((candidate) => candidate.id === itemId);
    if (!item) return;
    item.recipes.forEach((recipeId) => recipeIds.add(recipeId));
  });
  return [...recipeIds].map((id) => ({ id, ...catalog.recipes[id] })).filter((recipeItem) => recipeItem.name);
}

function getPlanRecipes() {
  return getRecipesForItems([...state.selected]);
}

function escapePowerShellCommand(command) {
  return command.replaceAll('"', '`"');
}

function buildInstallCommand(recipes) {
  const commands = recipes.map((recipeItem) => escapePowerShellCommand(recipeItem.command)).join("; ");
  if (!commands) return "";
  return `powershell -NoProfile -ExecutionPolicy Bypass -Command "${commands}"`;
}

function copyText(text, button) {
  if (!text) return;
  if (navigator.clipboard && navigator.clipboard.writeText) {
    navigator.clipboard.writeText(text).then(() => markCopied(button)).catch(() => fallbackCopy(text, button));
  } else {
    fallbackCopy(text, button);
  }
}

function fallbackCopy(text, button) {
  const textarea = document.createElement("textarea");
  textarea.value = text;
  textarea.style.position = "fixed";
  textarea.style.opacity = "0";
  document.body.appendChild(textarea);
  textarea.select();
  document.execCommand("copy");
  textarea.remove();
  markCopied(button);
}

function setButtonText(button, text) {
  if (!button) return;
  const span = button.querySelector("span");
  if (span) span.textContent = text;
  else button.textContent = text;
}

function markCopied(button) {
  if (!button) {
    byId("toast").textContent = t().copied;
    return;
  }
  const previous = button.querySelector("span") ? button.querySelector("span").textContent : button.textContent;
  button.classList.add("copied");
  setButtonText(button, t().copied);
  window.clearTimeout(button._copiedTimer);
  button._copiedTimer = window.setTimeout(() => {
    button.classList.remove("copied");
    setButtonText(button, previous);
  }, 1500);
}

function shortJobId(id) {
  return id ? id.slice(0, 8) : "-";
}

function escapeHtml(value) {
  return String(value || "")
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#039;");
}

function statusClass(status) {
  const value = String(status || "").toLowerCase();
  if (["completed"].includes(value)) return "success";
  if (["failed", "launch-failed"].includes(value)) return "danger";
  if (["canceled", "cancel-requested"].includes(value)) return "warning";
  if (["running", "launched", "created"].includes(value)) return "active";
  return "neutral";
}

function canCancelJob(status) {
  return ["created", "launched", "running"].includes(String(status || "").toLowerCase());
}

function canRetryJob(status) {
  return ["failed", "canceled", "launch-failed"].includes(String(status || "").toLowerCase());
}

function formatDateTime(value) {
  if (!value) return "";
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return "";
  return date.toLocaleString(state.language === "zh" ? "zh-CN" : "en-US", {
    hour12: false,
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit"
  });
}

function renderJobs() {
  const node = byId("jobs");
  if (!node) return;
  if (!state.agent.online) {
    node.innerHTML = `<p class="muted-note">${t().jobsOffline}</p>`;
    return;
  }
  if (!state.jobs.length) {
    node.innerHTML = `<p class="muted-note">${t().noJobs}</p>`;
    return;
  }
  node.innerHTML = state.jobs.slice(0, 5).map((job) => {
    const expanded = state.expandedLogs.has(job.id);
    const log = state.jobLogs[job.id] || "";
    const completedAt = formatDateTime(job.completedAt);
    return `
    <div class="job">
      <div class="job-head">
        <strong>${shortJobId(job.id)}</strong>
        <span class="job-status ${statusClass(job.status)}">${job.status || "-"}</span>
      </div>
      <div class="job-meta">
        ${Number.isInteger(job.exitCode) ? `<span>${t().jobExitCode}: ${job.exitCode}</span>` : ""}
        ${completedAt ? `<span>${t().jobCompletedAt}: ${completedAt}</span>` : ""}
      </div>
      <code>${t().jobScript}: ${job.scriptPath || "-"}</code>
      <code>${t().jobLog}: ${job.logPath || "-"}</code>
      <button class="inline-action job-log-toggle" type="button" data-job-log="${job.id}">
        ${expanded ? t().hideLog : t().viewLog}
      </button>
      ${canCancelJob(job.status) ? `
        <button class="inline-action cancel-job" type="button" data-job-cancel="${job.id}">
          ${t().cancelJob}
        </button>
      ` : ""}
      ${canRetryJob(job.status) ? `
        <button class="inline-action retry-job" type="button" data-job-retry="${job.id}">
          ${t().retryJob}
        </button>
      ` : ""}
      ${expanded ? `<pre class="job-log">${escapeHtml(log || t().emptyLog)}</pre>` : ""}
    </div>
  `;
  }).join("");

  document.querySelectorAll("[data-job-log]").forEach((button) => {
    button.addEventListener("click", () => toggleJobLog(button.dataset.jobLog));
  });
  document.querySelectorAll("[data-job-cancel]").forEach((button) => {
    button.addEventListener("click", () => cancelJob(button.dataset.jobCancel, button));
  });
  document.querySelectorAll("[data-job-retry]").forEach((button) => {
    button.addEventListener("click", () => retryJob(button.dataset.jobRetry, button));
  });
}

async function refreshExpandedJobLogs() {
  if (!state.agent.online || !state.expandedLogs.size) return;
  await Promise.all([...state.expandedLogs].map(async (jobId) => {
    try {
      const response = await fetch(`${state.agent.url}/jobs/${jobId}/log`, { cache: "no-store" });
      if (response.ok) {
        const data = await response.json();
        state.jobLogs[jobId] = data.log || "";
      }
    } catch {
      state.jobLogs[jobId] = state.jobLogs[jobId] || "";
    }
  }));
}

async function toggleJobLog(jobId) {
  if (!jobId) return;
  if (state.expandedLogs.has(jobId)) {
    state.expandedLogs.delete(jobId);
    renderJobs();
    return;
  }
  if (!state.jobLogs[jobId] && state.agent.online) {
    try {
      const response = await fetch(`${state.agent.url}/jobs/${jobId}/log`, { cache: "no-store" });
      if (response.ok) {
        const data = await response.json();
        state.jobLogs[jobId] = data.log || "";
      }
    } catch {
      state.jobLogs[jobId] = "";
    }
  }
  state.expandedLogs.add(jobId);
  renderJobs();
}

async function cancelJob(jobId, button) {
  if (!jobId || !state.agent.online) return;
  if (!window.confirm(t().cancelConfirm || "Request cancellation for this job?")) return;
  const previousText = button ? button.textContent : "";
  if (button) {
    button.disabled = true;
    button.textContent = "...";
  }
  try {
    const response = await fetch(`${state.agent.url}/jobs/${jobId}/cancel`, { method: "POST" });
    if (!response.ok) throw new Error("Cancel failed");
    await refreshJobs();
  } catch {
    showToast("Cancel failed");
  } finally {
    if (button) {
      button.disabled = false;
      button.textContent = previousText || t().cancelJob;
    }
  }
}

async function retryJob(jobId, button) {
  if (!jobId || !state.agent.online || !state.agent.allowExecute) return;
  const previousText = button ? button.textContent : "";
  if (button) {
    button.disabled = true;
    button.textContent = "...";
  }
  try {
    const response = await fetch(`${state.agent.url}/jobs/${jobId}/retry`, { method: "POST" });
    if (!response.ok) throw new Error("Retry failed");
    await refreshJobs();
  } catch {
    showToast("Retry failed");
  } finally {
    if (button) {
      button.disabled = false;
      button.textContent = previousText || t().retryJob;
    }
  }
}

async function executeViaAgent(itemIds, button) {
  if (!itemIds.length) return;
  if (!state.agent.online || !state.agent.allowExecute) {
    copyText(buildInstallCommand(getRecipesForItems(itemIds)), button);
    return;
  }
  if (!window.confirm(t().executeConfirm || "Execute selected install commands in a local PowerShell window?")) {
    return;
  }
  try {
    const response = await fetch(`${state.agent.url}/execute`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ itemIds })
    });
    if (!response.ok) throw new Error("Agent execution failed");
    await response.json().catch(() => null);
    await refreshJobs();
    markCopied(button);
  } catch {
    copyText(buildInstallCommand(getRecipesForItems(itemIds)), button);
  }
}

function renderPlan() {
  const recipes = getPlanRecipes();
  const verify = recipes.map((recipeItem) => recipeItem.verify).filter(Boolean);
  byId("selectedCount").textContent = state.selected.size;
  byId("actionCount").textContent = recipes.length;
  byId("verifyCount").textContent = verify.length;

  byId("actions").innerHTML = recipes.length
    ? recipes.map((recipeItem) => `<li><strong>${recipeItem.name}</strong><code>${recipeItem.command}</code></li>`).join("")
    : `<li>${t().emptyPlan}</li>`;

  byId("env").innerHTML = `
    <code>PROJECT_ENV=development</code>
    <code>PIP_DISABLE_PIP_VERSION_CHECK=1</code>
    <code>HF_HOME=./.cache/huggingface</code>
  `;

  const stackLines = [...state.selected].map((id) => `  - id: ${id}`).join("\n") || "  - id: template.ai-python-workstation";
  const installCommand = buildInstallCommand(recipes);
  byId("cli").textContent = [
    installCommand ? "# 安装所选环境" : "# 请先选择一个环境栈",
    installCommand || "No install command yet.",
    "",
    "powershell -ExecutionPolicy Bypass -File mvp/envforge.ps1 plan examples/ai-backend.envforge.yaml",
    "",
    "# 当前选择可保存为：",
    "schema: envforge/v1",
    "name: custom-workstation",
    "stacks:",
    stackLines,
    "verify:",
    "  commands:",
    ...verify.map((command) => `    - ${command}`)
  ].join("\n");
  renderJobs();
}

function render() {
  applyShellState();
  renderCategories();
  renderCatalog();
  renderPlan();
}

byId("search").addEventListener("input", (event) => {
  state.query = event.target.value;
  renderCatalog();
});

byId("clear").addEventListener("click", () => {
  state.selected.clear();
  render();
});

byId("templateAi").addEventListener("click", () => {
  state.selected.clear();
  state.selected.add("template.ai-python-workstation");
  render();
});

byId("templateWeb").addEventListener("click", () => {
  state.selected.clear();
  state.selected.add("template.fullstack-web");
  render();
});

byId("installSelected").addEventListener("click", () => {
  executeViaAgent([...state.selected], byId("installSelected"));
});

byId("copyCli").addEventListener("click", (event) => {
  event.preventDefault();
  event.stopPropagation();
  copyText(byId("cli").textContent, byId("copyCli"));
});

byId("themeToggle").addEventListener("click", () => {
  state.theme = state.theme === "dark" ? "light" : "dark";
  localStorage.setItem("envforge-theme", state.theme);
  applyShellState();
});

byId("languageToggle").addEventListener("click", () => {
  state.language = state.language === "zh" ? "en" : "zh";
  localStorage.setItem("envforge-language", state.language);
  render();
});

byId("collapseSidebar").addEventListener("click", () => {
  state.sidebarCollapsed = true;
  localStorage.setItem("envforge-sidebar", "collapsed");
  applyShellState();
});

byId("expandSidebar").addEventListener("click", () => {
  state.sidebarCollapsed = false;
  localStorage.setItem("envforge-sidebar", "expanded");
  applyShellState();
});

async function initializeApp() {
  await loadWebCatalog();
  render();
  refreshAgentStatus();
  window.setInterval(refreshAgentStatus, 5000);
  window.setInterval(refreshJobs, 5000);
}

initializeApp();
