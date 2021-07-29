package gct

import (
	"errors"
	"fmt"
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

// Offset 时间偏移
func (td *TDate) Offset(date time.Time, format string) (time.Time, error) {
	pd, err := time.ParseDuration(format)
	if err != nil {
		return date, err
	}
	return date.Add(pd), nil
}

// OffsetDay 时间日期偏移
func (td *TDate) OffsetDay(date time.Time, day int) (time.Time, error) {
	return td.Offset(date, fmt.Sprintf("%dh", day*24))
}

// OffsetHour 按小时偏移
func (td *TDate) OffsetHour(date time.Time, hour int) (time.Time, error) {
	return td.Offset(date, fmt.Sprintf("%dh", hour))
}

// OffsetMinute 按分钟偏移
func (td *TDate) OffsetMinute(date time.Time, minute int) (time.Time, error) {
	return td.Offset(date, fmt.Sprintf("%dm", minute))
}

// OffsetSecond 按秒偏移
func (td *TDate) OffsetSecond(date time.Time, second int) (time.Time, error) {
	return td.Offset(date, fmt.Sprintf("%ds", second))
}

// OffsetMillisecond 按毫秒偏移
func (td *TDate) OffsetMillisecond(date time.Time, ms int) (time.Time, error) {
	return td.Offset(date, fmt.Sprintf("%dms", ms))
}

// SubDays 日期差
func (td *TDate) SubDays(date1, date2 time.Time) int {
	return int(date1.Sub(date2).Hours() / 24)
}

// SubHours 小时差
func (td *TDate) SubHours(date1, date2 time.Time) int {
	return int(date1.Sub(date2).Hours())
}

// SubMinutes 分钟差
func (td *TDate) SubMinutes(date1, date2 time.Time) int {
	return int(date1.Sub(date2).Minutes())
}

// SubSeconds 秒差
func (td *TDate) SubSeconds(date1, date2 time.Time) int {
	return int(date1.Sub(date2).Seconds())
}

// SubMilliseconds 毫秒差
func (td *TDate) SubMilliseconds(date1, date2 time.Time) int {
	return int(date1.Sub(date2).Milliseconds())
}

// GetZodiac 根据生日计算星座
func (td *TDate) GetZodiac(month, day int) string {
	month = month - 1
	var (
		DAY_ARR = [12]int{20, 19, 21, 20, 21, 22, 23, 23, 23, 24, 23, 22}
		ZODIACS = [13]string{"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座"}
	)

	if day < DAY_ARR[month] {
		return ZODIACS[month]
	} else {
		return ZODIACS[month+1]
	}
}

// GetChineseZodiac 根据生日计算生肖
func (td *TDate) GetChineseZodiac(year int) string {
	var CHINESE_ZODIACS = [12]string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	if year > 1900 {
		return CHINESE_ZODIACS[(year-1900)%len(CHINESE_ZODIACS)]
	} else {
		return ""
	}
}

// AgeOfIDNumber 根据身份证号计算年龄
func (td *TDate) AgeOfIDCard() int {
	return 0
}
