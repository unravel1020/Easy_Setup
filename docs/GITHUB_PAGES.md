# GitHub Pages 发布说明

当前 MVP 的 Web 页面是纯静态资源，可以通过 GitHub Pages 发布 `mvp/web`。

## 发布分支

线上页面使用 `gh-pages` 分支。该分支由 `mvp/web` 子目录生成。

本地重新生成并推送：

```powershell
git -c safe.directory=C:/Users/15517/Desktop/project branch -D gh-pages
git -c safe.directory=C:/Users/15517/Desktop/project subtree split --prefix=mvp/web -b gh-pages
git -c safe.directory=C:/Users/15517/Desktop/project push -f origin gh-pages:gh-pages
```

## 在线地址

```text
https://unravel1020.github.io/Easy_Setup/
```

## 静态资源

页面入口：

```text
mvp/web/index.html
```

展示目录：

```text
mvp/web/catalog.web.json
```

GitHub Pages 会加载 `catalog.web.json`。如果本地浏览器阻止 `file://` 读取 JSON，`app.js` 会使用内置兜底目录。

## 注意事项

- GitHub Pages 只能托管静态页面，不能直接执行本机命令。
- 真实安装必须依赖用户本机启动的 PowerShell Agent 或 Go Agent。
- 推送后 GitHub Pages 可能有短暂发布延迟。
