package utils

import (
	"strings"
	"time"
)

func FormatDate(val time.Time, dateTime bool) string {
	if dateTime {
		return val.Format("2006-01-02 15:04:05")
	} else {
		res := strings.Split(val.Format("2006-01-02 15:04:05"), " ")[0]
		// 如果该日期实际上是空的
		if res == "0001-01-01" {
			return ""
		}
		return res
	}
}

func FormatProjectStatus(status int) string {
	switch status {
	case 0:
		return "立项"
	case 1:
		return "进行中"
	case 2:
		return "中止"
	case 3:
		return "作废"
	case 4:
		return "结项"
	default:
		return "undefined"
	}
}
