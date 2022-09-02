package templatereq

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type URLReq struct {
	Url     string
	Method  string
	Body    interface{}
	Headers map[string]string
	Type    string
}

func (ur *URLReq) RequestUrl() (*http.Response, error) {

	// JSON
	switch t := ur.Body.(type) {
	case string:
		client := &http.Client{}
		r, _ := http.NewRequest(ur.Method, ur.Url, bytes.NewBuffer([]byte(t))) // JSON payload
		for k, e := range ur.Headers {
			r.Header.Add(k, e)
		}
		return client.Do(r)
	case map[string]interface{}:
		client := &http.Client{}
		data := url.Values{}
		for i, n := range t {
			data.Set(i, fmt.Sprintf("%v", n))
		}
		r, _ := http.NewRequest(ur.Method, ur.Url, strings.NewReader(data.Encode())) // URL-encoded payload
		for k, e := range ur.Headers {
			r.Header.Add(k, e)
		}
		return client.Do(r)
	default:
		var r = reflect.TypeOf(t)
		fmt.Printf("Other:%v\n", r)
	}
	return nil, errors.New("not declared")
}
