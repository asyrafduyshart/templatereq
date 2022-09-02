package templatereq

import (
	"fmt"
	"net/http"
)

func (c *URLReq) ReplaceRequest(r map[string]string) (*http.Response, error) {

	headers := make(map[string]string)

	for i, j := range c.Headers {
		headers[Replace(i, r)] = Replace(j, r)
	}

	var body interface{}

	switch v := c.Body.(type) {
	case string:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		body = Replace(v, r)
	case map[string]interface{}:
		var bdy = map[string]interface{}{}
		for i, j := range v {
			bdy[Replace(i, r)] = Replace(fmt.Sprintf("%v", j), r)
		}
		body = bdy
	default:
		// And here I'm feeling dumb. ;)
		fmt.Printf("I don't know, ask stackoverflow.")
		body = c.Body
	}

	rr := &URLReq{
		Url:     Replace(c.Url, r),
		Method:  c.Method,
		Body:    body,
		Headers: headers,
	}

	return rr.RequestUrl()
}
