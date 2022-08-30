package templatereq

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestReplaceRequestJSON(t *testing.T) {
	ur := &URLReq{
		Url:    "https://httpbin.org/post",
		Body:   `{"hello": "$HELLO"}`,
		Method: http.MethodPost,
		Headers: map[string]string{
			"Accept": "application/json",
		},
	}

	replace := map[string]string{
		"HELLO": "there",
	}

	resp, err := ur.ReplaceRequest(replace)

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
		if !strings.Contains(bodyString, `"data": "{\"hello\": \"there\"}"`) {
			t.Errorf("result does not contain expected data")
		}
		// t.Log(bodyString)
		return
	}

	t.Errorf("got %v, want %v", "failed", "success")
}

func TestReplaceRequestEncoded(t *testing.T) {
	ur := &URLReq{
		Url: "https://httpbin.org/post",
		Body: map[string]interface{}{
			"test": "hello",
		},
		Method: http.MethodPost,
		Headers: map[string]string{
			"Accept": "application/x-www-form-urlencoded",
			"Hello":  "$HELLO",
		},
	}

	replace := map[string]string{
		"HELLO": "there",
	}

	resp, err := ur.ReplaceRequest(replace)

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
		if !strings.Contains(bodyString, `"data": "test=hello"`) {
			t.Errorf("result does not contain expected data")
		}
		return
	}

	t.Errorf("got %v, want %v", "failed", "success")
}
