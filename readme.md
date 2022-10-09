### build
```cmd
go build -ldflags="-H windowsgui" -o build/wreader.exe
```

### 说明
- wreader.exe不能脱离WeView2Loader.dll运行，请放置在统一目录

### 引用工具
- icon打包工具 rsrc https://github.com/akavel/rsrc
- gui工具 https://github.com/webview/webview