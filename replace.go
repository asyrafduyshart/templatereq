package templatereq

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func replaceByMap(t string, v map[string]string) string {
	re := regexp.MustCompile(`\$[A-Z_]+`)

	sub := func(match string) string {
		return v[match[1:]]
	}
	return re.ReplaceAllStringFunc(t, sub)
}

func funcName(t string) string {
	left := `$func("`
	right := `")`
	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
	matches := rx.FindAllStringSubmatch(t, -1)
	if len(matches) <= 0 {
		return t
	}
	res := matches[0][1]
	funcname := strings.Split(res, ":")
	return funcname[0]
}

func replaceFuncWithValue(t string) string {
	d := trimQuotes(t)
	fname := funcName(d)
	left := `$func("`
	right := `")`
	re := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
	sub := func(match string) string {
		val := strings.Replace(match[len(left):], fname+":", "", 1)
		return funcSwitch(fname, val[:len(val)-len(right)])
	}
	val := re.ReplaceAllStringFunc(d, sub)
	return val
}

func funcSwitch(f, v string) string {
	switch f {
	case "hash":
		return funcHash(v)
	case "md5":
		return funcMd5(v)
	case "base64":
		return funcBase64(v)
	case "sha256":
		return funcSha256(v)
	case "dateFormat":
		return funcNormalizeDateWithAdjustment(v)
	case "lowercase":
		return funcEncrToLowerCase(v)
	case "uppercase":
		return funcEncrToUpperCase(v)
	case "uuid":
		return funcGenUUID()
	case "base64ToStr":
		return funcDecodeBase64ToStr(v)
	default:
		return v
	}
}

// code hash
func funcHash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	res := fmt.Sprint(h.Sum32())
	return res
}

func funcBase64(text string) string {
	hash := md5.Sum([]byte(text))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func funcNormalizeDateWithAdjustment(date string) string {
	arr := make([]string, 0)
	arrAdd := strings.Split(date, ":add:")
	arrSub := strings.Split(date, ":subtract:")

	if len(arrAdd) > 1 {
		arr = arrAdd
	} else if len(arrSub) > 1 {
		arr = arrSub
	}

	if len(arr) > 1 {
		datetime := arr[0]
		durationcount := strings.Split(arr[1], "*")
		durationtime := 1

		for _, v := range durationcount {
			i, _ := strconv.Atoi(v)
			durationtime = durationtime * i
		}

		if len(arrAdd) > 1 {
			date = AddDateInSecond(datetime, durationtime)
		} else if len(arrSub) > 1 {
			date = SubtractDateInSecond(datetime, durationtime)
		}
	} else {
		date = FormatNormalDate(date)
	}

	return date
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func funcSortKey(s map[string]string) map[string]string {
	keys := make([]string, 0, len(s))
	new := map[string]string{}
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		new[k] = s[k]
	}
	return new
}

func funcUrlEncode(m map[string]string) string {
	params := url.Values{}
	for k, v := range m {
		params.Add(k, v)
	}
	r := params.Encode()
	fmt.Println(r)
	return r
}

func funcMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func funcSha256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func funcEncrToLowerCase(text string) string {
	encrs := strings.Split(text, ":encr:")
	if len(encrs) > 1 {
		encrType := encrs[1]
		oldTxt := strings.ToLower(encrs[0])
		newTxt := funcSwitch(encrType, oldTxt)
		return newTxt
	} else {
		return strings.ToLower(text)
	}
}

func funcEncrToUpperCase(text string) string {
	encrs := strings.Split(text, ":encr:")
	if len(encrs) > 1 {
		encrType := encrs[1]
		oldTxt := strings.ToUpper(encrs[0])
		newTxt := funcSwitch(encrType, oldTxt)
		return newTxt
	} else {
		return strings.ToUpper(text)
	}
}

func funcGenUUID() string {
	return uuid.New().String()
}

func funcDecodeBase64ToStr(str string) string {
	dcd, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(dcd[:])
}

//ALL

func Replace(t string, v map[string]string) string {
	return replaceByMap(replaceFuncWithValue(replaceByMap(t, v)), v)
}
