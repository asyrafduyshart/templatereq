package templatereq

import (
	"fmt"
	"hash/fnv"
	"regexp"
	"strings"
)

func replaceByMap(t string, v map[string]string) string {
	re := regexp.MustCompile(`\$[A-Z_]+`)

	sub := func(match string) string {
		return v[match[1:]]
	}
	return re.ReplaceAllStringFunc(t, sub)
}

func funcName(t string) string {
	left := `func("`
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
	default:
		return v
	}
}

// code hash
func funcHash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	res := fmt.Sprint(h.Sum32())
	fmt.Println("hash", res)
	return res
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

//ALL

func Replace(t string, v map[string]string) string {
	return replaceByMap(replaceFuncWithValue(t), v)
}
