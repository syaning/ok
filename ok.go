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
	return &request{&http.Request{}}
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

func (r *request) Method(method string) *request {
	r.req.Method = method
	return r
}

func (r *request) Url(uri string) *request {
	// todo
	return r
}

func (r *request) Set(key, value string) *request {
	// todo
	return r
}

func (r *request) Header(key, value string) *request {
	return r.Set(key, value)
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
