package templatereq

import (
	"fmt"
	"time"
)

func parseStringToDateTime(t string) time.Time {
	dateString := "2006-01-02T15:04:05Z"
	date, err := time.Parse(dateString, t)
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

func AddDateInMinute(t string, n int) string {
	date := parseStringToDateTime(t)
	d1 := date.Add(time.Minute * time.Duration(n))
	d2 := d1.String()[0:19]
	return d2
}

func AddDateInHour(t string, n int) string {
	date := parseStringToDateTime(t)
	d1 := date.Add(time.Hour * time.Duration(n))
	d2 := d1.String()[0:19]
	return d2
}

func AddDateInDay(t string, n int) string {
	date := parseStringToDateTime(t)
	d1 := date.AddDate(0, 0, n)
	d2 := d1.String()[0:19]
	return d2
}

func SubtractDateInMinute(t string, n int) string {
	date := parseStringToDateTime(t)
	d1 := date.Add(time.Minute * time.Duration(-n))
	d2 := d1.String()[0:19]
	return d2
}

func SubtractDateInHour(t string, n int) string {
	date := parseStringToDateTime(t)
	d1 := date.Add(time.Hour * time.Duration(-n))
	d2 := d1.String()[0:19]
	return d2
}

func SubtractDateInDay(t string, n int) string {
	date := parseStringToDateTime(t)
	d1 := date.AddDate(0, 0, -n)
	d2 := d1.String()[0:19]
	return d2
}
