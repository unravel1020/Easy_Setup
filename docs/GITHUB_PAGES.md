# GitHub Pages 发布说明

当前 MVP 是纯静态页面，可以直接通过 GitHub Pages 发布 `mvp/web`。

## 当前状态

本地目录还不是 Git 仓库，并且当前机器没有安装 `gh` 命令，所以 Codex 不能在本轮直接把页面上传到 GitHub。

## 推荐发布方式

1. 在 GitHub 创建一个新仓库，例如 `envforge`。
2. 在本地项目根目录执行：

```powershell
git init
git add .
git commit -m "Initial EnvForge MVP"
git branch -M main
git remote add origin https://github.com/<your-name>/envforge.git
git push -u origin main
```

3. 打开 GitHub 仓库的 `Settings -> Pages`。
4. Source 选择 `GitHub Actions`。
5. 手动运行工作流 `Deploy EnvForge MVP to GitHub Pages`，或再次 push 到 `main`。

工作流文件位于：

```text
.github/workflows/pages.yml
```

它会把 `mvp/web` 上传为 GitHub Pages 站点。

## 直接预览

本地仍可直接打开：

```text
mvp/web/index.html
```
