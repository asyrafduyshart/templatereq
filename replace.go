package templatereq

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"net/url"
	"regexp"
	"sort"
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

//ALL

func Replace(t string, v map[string]string) string {
	return replaceByMap(replaceFuncWithValue(t), v)
}
