package templatereq

import (
	"fmt"
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

func TestDateTimeOffset(t *testing.T) {
	date := replaceFuncWithValue(`$func("dateOffset:::subtract:60*60*4:format:Dtmz")`)
	fmt.Println(date)
}

func TestGetDateNowInUnixMilli(t *testing.T) {
	date := replaceFuncWithValue(`$func("dateNow:UnixMilli")`)
	fmt.Println(date)
}
