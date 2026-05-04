# 技术选型

## 核心语言：Go

选择 Go 的原因：

- 易于编译为单文件二进制，符合轻量化目标。
- 跨平台系统调用、文件处理和命令执行能力成熟。
- 启动速度快，适合 CLI 和桌面应用后端核心。
- 标准库覆盖面广，可以减少依赖。
- 交叉编译和发布流程简单。

## GUI 推荐：Wails + 原生 WebView

推荐使用 Wails 作为桌面应用框架：

- Go 后端可以直接复用 Core Engine。
- 前端使用 Web 技术构建界面，但运行时依赖系统原生 WebView。
- 不捆绑 Chromium，安装包和内存占用通常显著低于 Electron。
- 支持 Windows、macOS、Linux。
- 适合“轻应用 + 本地系统能力 + 可视化流程”的产品形态。

## GUI 备选方案

### Tauri

Tauri 同样轻量且跨平台表现好，但后端主要使用 Rust。若未来核心转向 Rust，Tauri 会成为强候选；在 Go 核心优先的前提下，Wails 的整合成本更低。

### Fyne

Fyne 是纯 Go GUI，打包简单，资源占用较低。但复杂表格、日志、差异视图和现代交互体验的开发效率不如 Web UI。

### Electron

Electron 生态成熟，但捆绑 Chromium，安装包和运行内存偏大，不符合“运行负担小”的核心诉求，因此不作为首选。

## 前端技术

建议使用：

- TypeScript。
- Svelte 或 Solid，优先考虑低运行时开销。
- CSS Modules 或轻量 CSS 方案。
- 虚拟列表用于长日志。
- 本地化文案独立维护，便于后续中英文切换。

不建议第一阶段引入重型前端框架、复杂状态管理库或大型组件库。

## 配置格式：YAML + JSON Schema

选择 YAML 作为用户编写格式，原因是可读性好，适合声明式环境配置。使用 JSON Schema 约束结构，便于 IDE 提示、校验和文档生成。

## 本地状态：JSON

状态文件使用 JSON，原因是跨语言可读、便于调试、无需引入数据库。后续如果状态复杂，可追加 SQLite 作为可选后端，但不作为第一阶段依赖。

## 包管理适配

第一阶段建议支持：

- Windows：winget、scoop。
- macOS：Homebrew。
- Debian/Ubuntu：apt。
- Fedora/RHEL：dnf。
- Arch：pacman。
- 通用版本管理：mise 或 asdf 作为可选适配器。

## 初始依赖建议

- CLI 框架：先用 Go 标准 `flag`，需要子命令体验时再评估 `spf13/cobra`。
- 桌面框架：Wails。
- 前端构建：Vite。
- UI：Svelte 或 Solid。
- YAML 解析：`gopkg.in/yaml.v3`。
- Schema 校验：先内置轻量校验，后续再引入 JSON Schema 库。
- 日志：标准库 `log/slog`。
- 测试：Go 标准测试框架 + 前端单元测试。

## 发布方式

- GitHub Releases 提供 Windows、macOS、Linux 安装包和 CLI 二进制。
- GUI 安装包作为普通用户推荐入口。
- Homebrew tap、Scoop bucket、winget manifest 作为后续安装入口。
- 不要求用户先安装项目本体依赖。
