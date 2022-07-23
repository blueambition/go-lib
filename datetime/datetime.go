package datetime

import (
	"time"
)

// 格式化时间
func Format(timeStr, format string) string {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(timeLayout, timeStr, loc)
	ts := tmp.Unix() //转化为时间戳 类型是int64
	needTime := time.Unix(ts, 0).Format(format)
	return needTime
}

//转换成东八区时间
func ParseE8TS(timeStr string) time.Time {
	var shZone, _ = time.LoadLocation("Asia/Shanghai") //上海
	e8Time, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, shZone)
	return e8Time
}

//转换成东八区时间
func FormatE8(ts int64) string {
	var shZone, _ = time.LoadLocation("Asia/Shanghai") //上海
	e8TimeStr := time.Unix(ts, 0).In(shZone).Format("2006-01-02 15:04:05")
	return e8TimeStr
}
