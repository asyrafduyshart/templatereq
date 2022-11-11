package templatereq

import (
	"fmt"
	"testing"
	"time"
)

func TestNormalize(t *testing.T) {
	dateString := "2006-01-02T15:04:05Z"
	str := "2022-11-10T08:09:25Z"
	date, err := time.Parse(dateString, str)

	if err != nil {
		t.Errorf("got %v", err)
	}

	normalDate := date.String()[0:19]

	fmt.Println("normalDate: ", normalDate)

}

func TestAdjustDateInMin(t *testing.T) {
	dateString := "2006-01-02T15:04:05Z"
	str := "2022-11-10T08:09:25Z"
	date, err := time.Parse(dateString, str)

	if err != nil {
		fmt.Println(err)
	}
	d1 := date.Add(time.Minute * time.Duration(4))
	adjustedDateMinute := d1.String()[0:19]

	fmt.Println("adjustedDateMinute: ", adjustedDateMinute)
}
