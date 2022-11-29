package templatereq

import (
	"fmt"
	"time"
)

func parseStringToDateTime(t string) time.Time {
	dateFormat := time.RFC3339
	date, err := time.Parse(dateFormat, t)
	if err != nil {
		fmt.Println(err)
	}
	return date
}

func FormatNormalDate(t string) string {
	date := parseStringToDateTime(t)
	normalDate := date.String()[0:19]
	return normalDate
}

func AddDateInSecond(t string, n int) string {
	date := parseStringToDateTime(t)
	d1 := date.Add(time.Second * time.Duration(n))
	d2 := d1.String()[0:19]
	return d2
}

func SubtractDateInSecond(t string, n int) string {
	date := parseStringToDateTime(t)
	d1 := date.Add(time.Second * time.Duration(-n))
	d2 := d1.String()[0:19]
	return d2
}
