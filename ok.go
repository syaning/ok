package ok

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type RequestWrapper struct {
	client *http.Client
	req    *http.Request
	res    *http.Response
	err    error
}

func NewRequest(method, urlStr string) *RequestWrapper {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil
	}
	return &RequestWrapper{req: req}
}

func Request() *RequestWrapper {
	return NewRequest("", "")
}

func Get(urlStr string) *RequestWrapper {
	return NewRequest("GET", urlStr)
}

func Post(urlStr string) *RequestWrapper {
	return NewRequest("POST", urlStr)
}

func Put(urlStr string) *RequestWrapper {
	return NewRequest("PUT", urlStr)
}

func Delete(urlStr string) *RequestWrapper {
	return NewRequest("DELETE", urlStr)
}

func (r *RequestWrapper) Client() *http.Client {
	return r.client
}

func (r *RequestWrapper) Request() *http.Request {
	return r.req
}

func (r *RequestWrapper) Response() (*http.Response, error) {
	return r.res, r.err
}

func (r *RequestWrapper) Method(method string) *RequestWrapper {
	r.req.Method = method
	return r
}

func (r *RequestWrapper) Url(urlStr string) *RequestWrapper {
	u, err := url.Parse(urlStr)
	if err != nil {
		return r
	}
	r.req.URL = u
	return r
}

func (r *RequestWrapper) Query(query string) *RequestWrapper {
	r.req.URL.RawQuery = query
	return r
}

func (r *RequestWrapper) Param(key, value string) *RequestWrapper {
	query := r.req.URL.Query()
	query.Add(key, value)
	return r.Query(query.Encode())
}

func (r *RequestWrapper) Set(key, value string) *RequestWrapper {
	r.req.Header.Set(key, value)
	return r
}

func (r *RequestWrapper) Header(key, value string) *RequestWrapper {
	return r.Set(key, value)
}

func (r *RequestWrapper) BasicAuth(username, password string) *RequestWrapper {
	r.req.SetBasicAuth(username, password)
	return r
}

func (r *RequestWrapper) Form(data string) *RequestWrapper {
	reader := strings.NewReader(data)
	r.req.Body = ioutil.NopCloser(reader)
	r.req.ContentLength = int64(reader.Len())
	r.Set("Content-Type", "application/x-www-form-urlencoded")
	return r
}

func (r *RequestWrapper) Json(data string) *RequestWrapper {
	reader := strings.NewReader(data)
	r.req.Body = ioutil.NopCloser(reader)
	r.req.ContentLength = int64(reader.Len())
	r.Set("Content-Type", "application/json")
	return r
}

func (r *RequestWrapper) JSON(data string) *RequestWrapper {
	return r.Json(data)
}

func (r *RequestWrapper) Proxy(proxy string) *RequestWrapper {
	r.lazyClient()
	r.client.Transport = &http.Transport{
		Proxy: func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		},
	}
	return r
}

func (r *RequestWrapper) ProxyFn(proxyFn func(*http.Request) (*url.URL, error)) *RequestWrapper {
	r.lazyClient()
	r.client.Transport = &http.Transport{Proxy: proxyFn}
	return r
}

func (r *RequestWrapper) lazyClient() {
	if r.client == nil {
		r.client = &http.Client{}
	}
}

func (r *RequestWrapper) Use(client *http.Client) *RequestWrapper {
	if client != nil {
		r.client = client
	}
	return r
}

func (r *RequestWrapper) OK() *RequestWrapper {
	client := http.DefaultClient
	if r.client != nil {
		client = r.client
	}
	r.res, r.err = client.Do(r.req)
	return r
}

func (r *RequestWrapper) ToBytes() ([]byte, error) {
	res, err := r.Response()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (r *RequestWrapper) ToString() (string, error) {
	data, err := r.ToBytes()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *RequestWrapper) Pipe(w io.Writer) (written int64, err error) {
	res, err := r.Response()
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	written, err = io.Copy(w, res.Body)
	return
}

func (r *RequestWrapper) ToFile(filename string) (size int64, err error) {
	file, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	size, err = r.Pipe(file)
	return
}

func Download(urlStr, filename string) (size int64, err error) {
	size, err = Get(urlStr).OK().ToFile(filename)
	return
}
