# Golang Common Tools

[![Go Report Card](https://goreportcard.com/badge/github.com/viodo/gct)](https://goreportcard.com/report/github.com/viodo/gct)
[![Build Status](https://github.com/viodo/gct/workflows/gct-test/badge.svg)](https://github.com/viodo/gct/actions)
[![codecov](https://codecov.io/gh/viodo/gct/branch/master/graph/badge.svg)](https://codecov.io/gh/viodo/gct)
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

#### 日期时间操作

支持的常用格式

- yyyy/MM/dd HH:mm:ss
- yyyy.MM.dd HH:mm:ss
- yyyy年MM月dd日 HH时mm分ss秒
- yyyy-MM-dd
- yyyy/MM/dd
- yyyy.MM.dd
- HH:mm:ss
- HH时mm分ss秒
- yyyy-MM-dd HH:mm
- yyyy-MM-dd HH:mm:ss.SSS
- yyyyMMddHHmmss
- yyyyMMddHHmmssSSS
- yyyyMMdd


根据特定格式格式化日期

```go
s, err := DateUtils.Format(time.Now(), "yyyy-MM-dd")
```

将日期字符串转换为Time

```go
t, err := DateUtils.Parse("2006-01-02", "yyyy-MM-dd")
```

日期时间偏移

```go
// 按天偏移
t, err := DateUtils.OffsetDay(time.Now(), -2)
// 按小时偏移
t, err := DateUtils.OffsetHour(time.Now(), -2)
// 按分钟偏移
t, err := DateUtils.OffsetMinute(time.Now(), -2)
// 按秒偏移
t, err := DateUtils.OffsetSecond(time.Now(), -2)
// 按毫秒偏移
t, err := DateUtils.OffsetMillisecond(time.Now(), -2)
```

日期时间差

```go
// SubDays 日期差
days := DateUtils.SubDays(time.Now(), time.Now()) // n=0
// SubHours 小时差
hours := DateUtils.SubHours(time.Now(), time.Now()) // n=0
// SubMinutes 分钟差
minutes := DateUtils.SubMinutes(time.Now(), time.Now()) // n=0
// SubSeconds 秒差
seconds := DateUtils.SubSeconds(time.Now(), time.Now()) // n=0
// SubMilliseconds 毫秒差
milliseconds := DateUtils.SubMilliseconds(time.Now(), time.Now()) // n=0
```

#### 文件操作

将字符串写入文件，追加模式

```go
f, err := FileUtils.AppendString("test content", "E:/test.txt")
```

判断文件是否存在

```go
is := FileUtils.Exist("E:/test.txt")
```

#### 字符串操作

判断字符串是否为空

```go
is := StringUtils.IsBlank(" ")
```
