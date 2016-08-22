package ok

import (
	"net/http"
	"net/url"
)

const (
	GET  = "GET"
	POST = "POST"
)

type request struct {
	req *http.Request
}

func NewRequest(method, urlStr string) *request {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil
	}

	r := &request{req}
	return r
}

func Request() *request {
	return NewRequest("", "")
}

func Get(urlStr string) *request {
	return NewRequest(GET, urlStr)
}

func Post(urlStr string) *request {
	return NewRequest(POST, urlStr)
}

func (r *request) GetRequest() *http.Request {
	return r.req
}

// set request method
func (r *request) Method(method string) *request {
	r.req.Method = method
	return r
}

// set request url
func (r *request) Url(urlStr string) *request {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil
	}
	r.req.URL = u
	return r
}

// set request header
func (r *request) Set(key, value string) *request {
	r.req.Header.Set(key, value)
	return r
}

// set request header, alias for Set(key, value string)
func (r *request) Header(key, value string) *request {
	return r.Set(key, value)
}

// set basic authorization
func (r *request) BasicAuth(username, password string) *request {
	r.req.SetBasicAuth(username, password)
	return r
}

func (r *request) Type(typ string) *request {
	switch typ {
	case "form":
		r.Set("Content-Type", "application/x-www-form-urlencoded")
	case "json":
		r.Set("Content-Type", "application/json")
	}
	return r
}
