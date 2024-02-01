package templatereq

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

// function map
var functionMap = map[string]string{
	"hash":               "hash",
	"base64":             "base64",
	"sha256":             "sha256",
	"aes-ecb":            "aes-ecb",
	"des-cbc":            "des-cbc",
	"dateOffset":         "dateOffset",
	"dateTimeFormat":     "dateTimeFormat",
	"dateTimeZoneFormat": "dateTimeZoneFormat",
	"dateFormat":         "dateFormat",
	"lowercase":          "lowercase",
	"uppercase":          "uppercase",
	"uuid":               "uuid",
	"base64ToStr":        "base64ToStr",
	"encodeBase64ToStr":  "encodeBase64ToStr",
	"arrayPos":           "arrayPos",
	"chain":              "chain",
	"dateNow":            "dateNow",
	"httpReq":            "httpReq",
}

// method map
var methodMap = map[string]string{
	"encrypt": "encrypt",
}

// command enum
const (
	Prepend string = "prepend"
	Append  string = "append"
)

// distribute enum
const (
	RoundRobin string = "roundRobin"
	SecondTick string = "secondTick"
	Random     string = "random"
)

// date type
const (
	Unix     string = "Unix"
	Standard string = "Standard"
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
	case "aes-ecb":
		return funcAESECB(v)
	case "des-cbc":
		return funcDESCBC(v)
	case "dateOffset":
		return funcGetDateTimeWithOffset(v)
	case "dateFormat":
		return funcNormalizeDateWithAdjustment(v)
	case "dateTimeFormat":
		return funcDateTimeFormat(v)
	case "dateTimeZoneFormat":
		return funcDateTimeZoneFormat(v)
	case "lowercase":
		return funcEncrToLowerCase(v)
	case "uppercase":
		return funcEncrToUpperCase(v)
	case "uuid":
		return funcGenUUID()
	case "base64ToStr":
		return funcDecodeBase64ToStr(v)
	case "encodeBase64ToStr":
		return funcEncodeBase64ToStr(v)
	case "chain":
		return funcChaining(v)
	case "arrayPos":
		return funcGetArrayPosition(v)
	case "dateNow":
		return funcDateNow(v)
	case "httpReq":
		return funcHttpReq(v)
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

func funcAESECB(text string) string {
	arr := strings.Split(text, ":param:")
	pt := []byte(arr[0])  // plaintext to be encrypted
	key := []byte(arr[1]) // encryption key

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	blockMode := ecb.NewECBEncrypter(block)

	// blockMode.BlockSize() for AES-128 would be 16.
	// Need padding if len(pt) % 16 != 0.
	padder := padding.NewPkcs7Padding(blockMode.BlockSize())
	plainText, err := padder.Pad(pt)
	if err != nil {
		return ""
	}

	crypted := make([]byte, len(plainText))
	blockMode.CryptBlocks(crypted, plainText)
	q := base64.StdEncoding.EncodeToString(crypted)

	return q
}

func funcDESCBC(text string) string {
	arr := strings.Split(text, ":param:")
	pt := []byte(arr[0])  // plaintext to be encrypted
	key := []byte(arr[1]) // encryption key
	iv := []byte(arr[2])  // encription iv

	block, err := des.NewCipher(key)
	if err != nil {
		return ""
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)

	// blockMode.BlockSize() for DES would be 8.
	// Need padding if len(pt) % 8 != 0.
	padder := padding.NewPkcs5Padding()
	plainText, err := padder.Pad(pt)
	if err != nil {
		return ""
	}

	crypted := make([]byte, len(plainText))
	blockMode.CryptBlocks(crypted, plainText)
	q := base64.StdEncoding.EncodeToString(crypted)

	return q
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func funcGetDateTimeWithOffset(v string) string {
	return GetDateTimeOffset(v)
}

func funcNormalizeDateWithAdjustment(date string) string {
	arr := make([]string, 0)
	format := ""

	arrFrm := strings.Split(date, ":format:")
	arrAdd := strings.Split(arrFrm[0], ":add:")
	arrSub := strings.Split(arrFrm[0], ":subtract:")

	if len(arrAdd) > 1 {
		arr = arrAdd
	} else if len(arrSub) > 1 {
		arr = arrSub
	}

	if len(arrFrm) > 1 {
		format = arrFrm[1]
	} else {
		format = Dt.String()
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
			date = AddDateInSecond(datetime, durationtime, TimeFormat(format))
		} else if len(arrSub) > 1 {
			date = SubtractDateInSecond(datetime, durationtime, TimeFormat(format))
		}
	} else {
		date = FormatNormalDate(arrFrm[0], format)
	}

	return date
}

func funcDateTimeFormat(date string) string {
	f := ""
	r := ""
	arrFrm := strings.Split(date, ":format:")

	if len(arrFrm) > 1 && arrFrm[1] != "" {
		f = arrFrm[1]
	} else {
		f = time.RFC3339
	}

	dt := parseStringToDateTime(arrFrm[0])
	if DateTimeFormat[f] != "" {
		r = dt.Format(DateTimeFormat[f])
	} else {
		r = dt.Format(f)
	}

	return r
}

func funcDateTimeZoneFormat(date string) string {
	f := "Asia/Jakarta"
	r := ""
	arrFrm := strings.Split(date, ":format:")
	if len(arrFrm) > 1 && arrFrm[1] != "" {
		f = arrFrm[1]
	}

	dt := parseStringToDateTime(arrFrm[0])

	loc, err := time.LoadLocation(f)
	if err != nil {
		// could be providing unknown TimeZone Location for example "Konoha"
		return r
	}

	r = dt.In(loc).Format(time.RFC3339)

	return r
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

func funcEncodeBase64ToStr(str string) string {
	data := []byte(str)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)

	return string(dst)
}

func funcChaining(v string) string {
	chains := strings.Split(v, "@")
	prev := ""
	next := ""

	for i := 0; i < len(chains); i++ {
		arr := strings.Split(chains[i], ">>")
		next = prev + chain(arr)
		prev = next
	}

	return next
}

func chain(arrChain []string) string {
	val := ""
	for i := 0; i < len(arrChain); i++ {
		arr := strings.SplitN(arrChain[i], "::", 2)
		_, isFunc := functionMap[arr[0]]
		_, isEncrypt := methodMap[arr[0]]

		// array must >= 2
		if len(arr) < 2 {
			val = "[400] ERROR FUNCTION:[" + arr[0] + "] array length minimum is 2"
			break
		}

		// if encryption method
		if isEncrypt {
			val = funcSwitch(arr[1], val)
		}

		// if not encryption method
		if !isEncrypt {
			append := arr[0] == Append
			prepend := arr[0] == Prepend
			if append {
				val += arr[1]
			} else if prepend {
				val = arr[1] + val
			} else if isFunc {
				val += funcSwitch(arr[0], arr[1])
			}
		}
	}

	return val
}

func funcGetArrayPosition(svr string) string {
	arr := strings.Split(svr, "::")
	method := arr[0]
	servers := strings.Split(arr[1], ",")

	var serverLen int = len(servers)
	var serverId string = servers[0]

	switch method {
	case RoundRobin:
		break
	case SecondTick:
		partition := int(math.Round(float64(60 / serverLen)))
		currSecond := GetNow().Second()
		for i := 0; i < serverLen; i++ {
			a := i * partition
			b := ((i + 1) * partition)
			if currSecond > a && currSecond <= b {
				serverId = servers[i]
				break
			}
		}
	case Random:
		serverId = servers[rand.Intn(serverLen)]
	default:
		break
	}

	return serverId
}

func funcDateNow(v string) string {
	switch v {
	case Unix:
		return GetDateNowUnix()
	case Standard:
		return time.Now().Format(time.RFC3339)
	default:
		return v
	}
}

func funcHttpReq(v string) string {
	arr := strings.Split(v, "::")
	method, url, headers, body, response := arr[0], arr[1], arr[2], arr[3], arr[4]

	var Method string
	var Headers map[string]string
	var Body interface{}

	switch strings.ToUpper(method) {
	case "GET":
		Method = http.MethodGet
	case "POST":
		Method = http.MethodPost
	case "PUT":
		Method = http.MethodPut
	case "PATCH":
		Method = http.MethodPatch
	case "DELETE":
		Method = http.MethodDelete
	default:
		Method = http.MethodGet
	}

	err := json.Unmarshal([]byte(headers), &Headers)
	if err != nil {
		return err.Error()
	}

	err = json.Unmarshal([]byte(body), &Body)
	if err != nil {
		Body = ""
	}

	ur := &URLReq{
		Url:     url,
		Body:    Body,
		Method:  Method,
		Headers: Headers,
	}

	resp, err := ur.RequestUrl()
	if err != nil {
		return err.Error()
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	r := gjson.GetBytes(bodyBytes, response)
	if r.Exists() {
		return r.Str
	}

	return string(bodyBytes)
}

// ALL
func Replace(t string, v map[string]string) string {
	return replaceByMap(replaceFuncWithValue(replaceByMap(t, v)), v)
}
