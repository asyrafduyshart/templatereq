package templatereq

import (
	"fmt"
	"time"
)

func FormatNormalDate(t string) string {
	dateString := "2006-01-02T15:04:05Z"
	str := t
	date, err := time.Parse(dateString, str)

	if err != nil {
		fmt.Println(err)
	}
	normalDate := date.String()[0:19]

	return normalDate
}

func AdjustDateInMin(t string, minute int) string {
	dateString := "2006-01-02T15:04:05Z"
	str := t
	date, err := time.Parse(dateString, str)

	if err != nil {
		fmt.Println(err)
	}
	d1 := date.Add(time.Minute * time.Duration(minute))
	adjustedDateMinute := d1.String()[0:19]

	return adjustedDateMinute
}
