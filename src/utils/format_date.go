package utils

import "time"

func FormatDate(date time.Time) string {
	loc, _ := time.LoadLocation("America/Santiago")
	return date.In(loc).Format("2006-01-02 15:04:05")
}
