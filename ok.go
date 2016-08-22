package ok

import (
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

type request struct {
	req *http.Request
}

func Request() *request {
	req := &http.Request{
		Header: make(http.Header),
	}
	return &request{req}
}

func NewRequest(method, uri string) *request {
	r := Request()
	r.Method(method).Url(uri)
	return r
}

func Get(uri string) *request {
	return NewRequest(GET, uri)
}

func Post(uri string) *request {
	return NewRequest(POST, uri)
}

func (r *request) GetRequest() *http.Request {
	return r.req
}

// set request method
func (r *request) Method(method string) *request {
	r.req.Method = method
	return r
}

func (r *request) Url(uri string) *request {
	// todo
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
