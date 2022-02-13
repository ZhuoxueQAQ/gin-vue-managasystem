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
