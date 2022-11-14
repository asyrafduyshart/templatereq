package templatereq

import (
	"fmt"
	"testing"
	"time"
)

func TestNormalize(t *testing.T) {
	str := "2022-11-10T08:09:25Z"
	date := parseStringToDateTime(str)
	normalDate := date.String()[0:19]
	fmt.Println("normalDate: ", normalDate)
}

func TestAdjustDateInMin(t *testing.T) {
	str := "2022-11-10T08:09:25Z"
	date := parseStringToDateTime(str)
	d1 := date.Add(time.Minute * time.Duration(4))
	adjustedDateMinute := d1.String()[0:19]
	fmt.Println("originalDate: ", date)
	fmt.Println("adjustedDate: ", adjustedDateMinute)
}
