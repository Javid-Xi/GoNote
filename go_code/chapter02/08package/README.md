**出现问题** : .../main is not in GOROOT ...

**问题原因**:
通过 GO111MODULE 可以开启或关闭 go module 工具。
它可以设置以下三个值：off, on或者auto(默认)
1. GO111MODULE=off: 禁用 go module，编译时会在vendor目录下和GOPATH目录中查找依赖包。也把这种模式叫GOPATH模式。
2. GO111MODULE=on: 启用 go module，编译时会忽略GOPATH和vendor文件夹,只根据go.mod下载依赖,这种模式称作module-aware模式，这种模式下，GOPATH不再在build时扮演导入的角色，但是尽管如此，它还是承担着存储下载依赖包的角色。它会将依赖包放在GOPATH/pkg/mod目录下。
3. GO111MODULE=auto（默认值），默认值,也就是说在你不设置的情况下，就是auto。当项目在 GOPATH/src 目录之外,并且项目根目录有 go.mod 文件时，才开启 go module。

**解决方案**:
Windows cmd中:
```
set GO111MODULE=on 或者 set GO111MODULE=auto
```

[go module使用](https://blog.csdn.net/qq_34021712/article/details/109146367)