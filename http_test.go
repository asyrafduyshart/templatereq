package templatereq

import (
	"io"
	"net/http"
	"testing"
)

func TestUrlRequest(t *testing.T) {
	ur := &URLReq{
		Url:    "https://httpbin.org/post",
		Body:   `{"hello": "there"}`,
		Method: http.MethodPost,
		Headers: map[string]string{
			"Accept": "application/json",
		},
	}

	resp, err := ur.RequestUrl()
	if err != nil {
		t.Errorf("got %v, want %v", "failed", "success")
		return
	}
	if resp.StatusCode >= 200 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("got %v, want %v", "failed", "success")
		}
		bodyString := string(bodyBytes)
		t.Log(bodyString)
		return
	}

	t.Errorf("got %v, want %v", "failed", "success")
}

func TestUrlRequestWithTimeOutError(t *testing.T) {
	ur := &URLReq{
		Url:    "https://httpbin.org/post",
		Body:   `{"hello": "there"}`,
		Method: http.MethodPost,
		Headers: map[string]string{
			"Accept":          "application/json",
			"TimeoutDuration": "50ms",
		},
	}

	resp, err := ur.RequestUrl()
	if err != nil {
		t.Log("Success return timeout error")
		return
	}
	if resp.StatusCode >= 200 {
		t.Errorf("got %v, want %v", "success", "error timeout")
	}

}
