package templatereq

import (
	"testing"
	"time"
)

func TestNormalize(t *testing.T) {
	str := "2022-11-10T08:09:25Z"
	date := parseStringToDateTime(str)
	init := date.String()[0:19]
	expect := "2022-11-10 08:09:25"

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestAdjustDate(t *testing.T) {
	str := "2022-11-10T08:09:25Z"
	date := parseStringToDateTime(str)
	d1 := date.Add(time.Second * time.Duration(5))
	init := d1.String()[0:19]
	expect := "2022-11-10 08:09:30"

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}
