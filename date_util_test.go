package templatereq

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	str := "2022-11-10T08:09:25Z"
	date := parseStringToDateTime(str)
	init := date.Format(defaultDt)
	expect := "2022-11-10 08:09:25"

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestAdjustDate(t *testing.T) {
	str := "2022-11-10T08:09:25Z"
	date := AddDateInSecond(str, 5, Dt)
	init := date
	expect := "2022-11-10 08:09:30"

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}
