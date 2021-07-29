# Golang Common Tools

[![Go Report Card](https://goreportcard.com/badge/github.com/viodo/gct)](https://goreportcard.com/report/github.com/viodo/gct)
[![Build Status](https://github.com/viodo/gct/workflows/gct-test/badge.svg)](https://github.com/viodo/gct-test/actions)
[![Code Size](https://img.shields.io/github/languages/code-size/viodo/gct.svg?style=flat-square)](https://github.com/viodo/gct)

### 安装使用

安装

```shell
go get -u github.com/viodo/gct
```

使用

```go
import . "github.com/viodo/gct"
```
### 功能

#### 文件操作

将String写入文件，追加模式

```go
file, err := FileUtils.AppendString("test content", "E:/test.txt")
```

判断文件是否存在

```go
isExist := FileUtils.Exist("E:/test.txt")
```

