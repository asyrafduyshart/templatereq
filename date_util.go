package templatereq

import (
	"fmt"
	"time"
)

/**
- Copy from Golang Library time and add new feature time
- for example Dt, Dtz, Dtm
*/

type TimeFormat string

const defaultDt string = "2006-01-02 15:04:05"

const (
	Dt          TimeFormat = "YYYY-MM-DD hh:mm:ss"
	Dtz         TimeFormat = "YYYY-MM-DDThh:mm:ssZ"
	Dtm         TimeFormat = "YYYY-MM-DD hh:mm:ss.nnnn"
	Dtmz        TimeFormat = "YYYY-MM-DDThh:mm:ss.nnnnZ"
	Layout      TimeFormat = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       TimeFormat = "Mon Jan _2 15:04:05 2006"
	UnixDate    TimeFormat = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    TimeFormat = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      TimeFormat = "02 Jan 06 15:04 MST"
	RFC822Z     TimeFormat = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      TimeFormat = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     TimeFormat = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    TimeFormat = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     TimeFormat = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano TimeFormat = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     TimeFormat = "3:04PM"
	Stamp       TimeFormat = "Jan _2 15:04:05"
	StampMilli  TimeFormat = "Jan _2 15:04:05.000"
	StampMicro  TimeFormat = "Jan _2 15:04:05.000000"
	StampNano   TimeFormat = "Jan _2 15:04:05.000000000"
)

func (t TimeFormat) String() string {
	switch t {
	case Dt:
		return "2006-01-02 15:04:05"
	case Dtz:
		return "2006-01-02T15:04:05Z"
	case Dtm:
		return "2006-01-02 15:04:05.9999"
	case Dtmz:
		return "2006-01-02T15:04:05.9999Z"
	case Layout:
		return "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	case ANSIC:
		return "Mon Jan _2 15:04:05 2006"
	case UnixDate:
		return "Mon Jan _2 15:04:05 MST 2006"
	case RubyDate:
		return "Mon Jan 02 15:04:05 -0700 2006"
	case RFC822:
		return "02 Jan 06 15:04 MST"
	case RFC822Z:
		return "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	case RFC850:
		return "Monday, 02-Jan-06 15:04:05 MST"
	case RFC1123:
		return "Mon, 02 Jan 2006 15:04:05 MST"
	case RFC1123Z:
		return "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	case RFC3339:
		return "2006-01-02T15:04:05Z07:00"
	case RFC3339Nano:
		return "2006-01-02T15:04:05.999999999Z07:00"
	case Kitchen:
		return "3:04PM"
	case Stamp:
		return "Jan _2 15:04:05"
	case StampMilli:
		return "Jan _2 15:04:05.000"
	case StampMicro:
		return "Jan _2 15:04:05.000000"
	case StampNano:
		return "Jan _2 15:04:05.000000000"
	default:
		return defaultDt
	}
}

func parseStringToDateTime(t string) time.Time {
	date, err := time.Parse(time.RFC3339, t)
	if err != nil {
		fmt.Println(err)
	}
	return date
}

func FormatNormalDate(t string, ft string) string {
	date := parseStringToDateTime(t)
	normalDate := date.Format(TimeFormat(ft).String())
	return normalDate
}

func AddDateInSecond(t string, n int, ft TimeFormat) string {
	date := parseStringToDateTime(t)
	d1 := date.Add(time.Second * time.Duration(n))
	d2 := d1.Format(ft.String())
	return d2
}

func SubtractDateInSecond(t string, n int, ft TimeFormat) string {
	date := parseStringToDateTime(t)
	d1 := date.Add(-time.Duration(time.Second * time.Duration(n)))
	d2 := d1.Format(ft.String())
	return d2
}
