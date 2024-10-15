package templatereq

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type URLReq struct {
	Url     string
	Method  string
	Body    interface{}
	Headers map[string]string
	Type    string
}

func (ur *URLReq) RequestUrl() (*http.Response, error) {

	duration := 5000 * time.Millisecond

	// Check if a timeout header is specified
	if timeout, ok := ur.Headers["TimeoutDuration"]; ok && timeout != "" {
		var val, err = time.ParseDuration(timeout)
		if err == nil {
			duration = val
		}
	}

	client := &http.Client{
		Timeout: duration,
	}

	// JSON
	switch t := ur.Body.(type) {
	case string:
		r, _ := http.NewRequest(ur.Method, ur.Url, bytes.NewBuffer([]byte(t))) // JSON payload
		for k, e := range ur.Headers {
			r.Header.Add(k, e)
		}
		return client.Do(r)
	case map[string]interface{}:
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
