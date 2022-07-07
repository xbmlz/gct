<p align = "center">
<h1>Golang Common Tools</h1>
<a title="Build Status" target="_blank" href="https://github.com/xbmlz/gct/actions/workflows/test.yml"><img src="https://img.shields.io/github/workflow/status/xbmlz/gct/Go%20Test?style=flat-square"></a>
<a title="GoDoc" target="_blank" href="https://godoc.org/github.com/xbmlz/gct"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/xbmlz/gct"><img src="https://goreportcard.com/badge/github.com/xbmlz/gct?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/github/xbmlz/gct"><img src="https://img.shields.io/coveralls/github/xbmlz/gct.svg?style=flat-square&color=CC9933"></a>
<a title="Code Size" target="_blank" href="https://github.com/xbmlz/gct"><img src="https://img.shields.io/github/languages/code-size/xbmlz/gct.svg?style=flat-square"></a>
</p>

## 安装使用

安装

```shell
go get -u github.com/xbmlz/gct
```

使用

```go
import . "github.com/xbmlz/gct"
```
## 功能

### 日期时间操作

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

年龄

```go
// 根据身份证号计算年龄
age := DateUtils.AgeOfIDCard("123456199501170016") // 26
```

星座和属相

```go
// 星座
zodiac := DateUtils.GetZodiac(1, 17) // 摩羯座
// 属相
chineseZodiac := DateUtils.GetChineseZodiac(1995) // 猪
```

### 文件操作

将字符串写入文件，追加模式

```go
f, err := FileUtils.AppendString("test content", "E:/test.txt")
```

判断文件是否存在

```go
is := FileUtils.Exist("E:/test.txt")
```

移除后缀名

```go
s := FileUtils.RemoveSuffix("main.go") // main.go -> main
```

移除前缀

```go
s := FileUtils.RemovePrefix("main.go") // main.go -> go
```

压缩zip

```go
err := FileUtils.Zip("./testdata/archive.zip", "./testdata/csv", "./testdata/file.txt")
```

解压zip

```go
err := FileUtils.Unzip("./testdata/archive.zip", "./testdata/unzip")
```

### 字符串操作

判断字符串是否为空

```go
is := StringUtils.IsBlank(" ")
```
