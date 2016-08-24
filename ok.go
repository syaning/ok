package ok

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type request struct {
	client *http.Client
	req    *http.Request
	res    *http.Response
	err    error
}

func NewRequest(method, urlStr string) *request {
	req, err := http.NewRequest(method, urlStr, nil)

	if err != nil {
		return nil
	}

	return &request{req: req}
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

func Put(urlStr string) *request {
	return NewRequest(PUT, urlStr)
}

func Delete(urlStr string) *request {
	return NewRequest(DELETE, urlStr)
}

func (r *request) Client() *http.Client {
	return r.client
}

func (r *request) Request() *http.Request {
	return r.req
}

func (r *request) Response() (*http.Response, error) {
	return r.res, r.err
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
		return r
	}
	r.req.URL = u
	return r
}

// set query string
func (r *request) Query(query string) *request {
	r.req.URL.RawQuery = query
	return r
}

// set param
func (r *request) Param(key, value string) *request {
	query := r.req.URL.Query()
	query.Add(key, value)
	return r.Query(query.Encode())
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

// send form
func (r *request) Form(data string) *request {
	reader := strings.NewReader(data)
	r.req.Body = ioutil.NopCloser(reader)
	r.req.ContentLength = int64(reader.Len())
	r.Set("Content-Type", "application/x-www-form-urlencoded")
	return r
}

// send json
func (r *request) Json(data string) *request {
	reader := strings.NewReader(data)
	r.req.Body = ioutil.NopCloser(reader)
	r.req.ContentLength = int64(reader.Len())
	r.Set("Content-Type", "application/json")
	return r
}

// alias for Json(data string)
func (r *request) JSON(data string) *request {
	return r.Json(data)
}

// using proxy
func (r *request) Proxy(proxy string) *request {
	r.lazyClient()
	r.client.Transport = &http.Transport{
		Proxy: func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		},
	}
	return r
}

// using proxy function
func (r *request) ProxyFn(proxyFn func(*http.Request) (*url.URL, error)) *request {
	r.lazyClient()
	r.client.Transport = &http.Transport{Proxy: proxyFn}
	return r
}

// lazy client allocation
func (r *request) lazyClient() {
	if r.client == nil {
		r.client = &http.Client{}
	}
}

// use client
func (r *request) Use(client *http.Client) *request {
	if client != nil {
		r.client = client
	}
	return r
}

// send request
func (r *request) OK() *request {
	client := http.DefaultClient
	if r.client != nil {
		client = r.client
	}

	r.res, r.err = client.Do(r.req)
	return r
}
