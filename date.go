package gct

import (
	"errors"
	"time"
)

// 格式化支持的类型
var formatMap = map[string]string{
	"yyyy/MM/dd HH:mm:ss":     "2006/01/02 15:04:05",
	"yyyy.MM.dd HH:mm:ss":     "2006.01.02 15:04:05",
	"yyyy年MM月dd日 HH时mm分ss秒":   "2006年01月02 15时04分05秒",
	"yyyy-MM-dd":              "2006-01-02",
	"yyyy/MM/dd":              "2006/01/02",
	"yyyy.MM.dd":              "2006.01.02",
	"HH:mm:ss":                "15:04:05",
	"HH时mm分ss秒":               "15时04分05秒",
	"yyyy-MM-dd HH:mm":        "2006-01-02 15:04",
	"yyyy-MM-dd HH:mm:ss.SSS": "2006-01-02 15:04:05.999",
	"yyyyMMddHHmmss":          "20060102150405",
	"yyyyMMddHHmmssSSS":       "20060102150405999",
	"yyyyMMdd":                "20060102",
}

// Format 根据特定格式格式化日期
func (td *TDate) Format(date time.Time, format string) (string, error) {
	if _, ok := formatMap[format]; !ok {
		return "", errors.New("unsupported format")
	}
	return date.Format(formatMap[format]), nil
}

// Parse 将日期字符串转换为Time
func (td *TDate) Parse(dateStr, format string) (time.Time, error) {
	if _, ok := formatMap[format]; !ok {
		return time.Time{}, errors.New("unsupported format")
	}
	return time.Parse(formatMap[format], dateStr)
}
