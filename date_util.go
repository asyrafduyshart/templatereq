package templatereq

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/**
- Copy from Golang Library time and add new feature time
- for example Dt, Dtz, Dtm
*/

type TimeFormat string

const defaultDt string = "2006-01-02 15:04:05"

const (
	Add      string = "add"
	Subtract string = "subtract"
)

const (
	Dt          TimeFormat = "YYYY-MM-DD hh:mm:ss"
	Dtz         TimeFormat = "YYYY-MM-DDThh:mm:ssZ"
	Dtm         TimeFormat = "YYYY-MM-DD hh:mm:ss.nnnn"
	Dtmz        TimeFormat = "YYYY-MM-DDThh:mm:ss.nnnnZ"
	YYMMDD      TimeFormat = "YYMMDD"
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

var DateTimeFormat = map[string]string{
	"Layout":      time.Layout,
	"ANSIC":       time.ANSIC,
	"UnixDate":    time.UnixDate,
	"RubyDate":    time.RubyDate,
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Kitchen":     time.Kitchen,
	"Stamp":       time.Stamp,
	"StampMilli":  time.StampMilli,
	"StampMicro":  time.StampMicro,
	"StampNano":   time.StampNano,
	"DateTime":    time.DateTime,
	"DateOnly":    time.DateOnly,
	"TimeOnly":    time.TimeOnly,
}

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

func (t TimeFormat) CustomFormat(v time.Time) string {
	switch t {
	case YYMMDD:
		return formatYYMMDD(v)
	default:
		return v.Format(defaultDt)
	}
}

func parseStringToDateTime(t string) time.Time {
	date, err := time.Parse(time.RFC3339Nano, t)
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

func GetNow() time.Time {
	return time.Now()
}

func GetDateTimeOffset(v string) string {
	var arr []string
	format := ""
	method := ""
	durationtime := 1

	date := time.Now()
	_, offset := date.Zone() // get timezone environment's offset in seconds

	arrFrm := strings.Split(v, ":format:")
	arrAdd := strings.Split(arrFrm[0], ":add:")
	arrSub := strings.Split(arrFrm[0], ":subtract:")

	if len(arrAdd) > 1 {
		arr = arrAdd
		method = Subtract
		format = arrFrm[1]
	} else if len(arrSub) > 1 {
		arr = arrSub
		method = Subtract
		format = arrFrm[1]
	} else {
		return date.Format(Dtz.String())
	}

	durationcount := strings.Split(arr[1], "*")

	for _, v := range durationcount {
		i, _ := strconv.Atoi(v)
		durationtime = durationtime * i
	}

	durationtime += offset

	switch method {
	case Add:
		subtract := date.Add(time.Duration(time.Second * time.Duration(durationtime)))
		return TimeFormat(format).CustomFormat(subtract)
	case Subtract:
		subtract := date.Add(-time.Duration(time.Second * time.Duration(durationtime)))
		return TimeFormat(format).CustomFormat(subtract)
	default:
		return date.Format(Dtz.String())
	}
}

func formatYYMMDD(v time.Time) string {
	year := strconv.Itoa(v.Year())
	day := strconv.Itoa(v.Day())
	month := ""

	if int(v.Month()) < 10 {
		month = "0" + strconv.Itoa(int(v.Month()))
	} else {
		month = strconv.Itoa(int(v.Month()))
	}

	return year[len(year)-2:] + month + day
}

func GetDateNowUnix() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
