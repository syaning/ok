package ok

const (
	GET  = "GET"
	POST = "POST"
)

type Request struct {
	Method  string
	Url     string
	Headers map[string]string
}

// new HTTP request
func New() *Request {
	r := &Request{
		Method:  GET,
		Headers: make(map[string]string),
	}
	return r
}

// new GET HTTP request
func Get(url string) *Request {
	return New()
}

// new POST HTTP request
func Post(url string) *Request {
	r := New()
	r.Method = POST
	return r
}

// set request header
func (r *Request) Set(field string, value string) *Request {
	r.Headers[field] = value
	return r
}

// alias for Set(field, value)
func (r *Request) SetHeader(field string, value string) *Request {
	return r.Set(field, value)
}

// set Content-Type header
func (r *Request) Type(t string) *Request {
	switch t {
	case "form":
		r.Set("Content-Type", "application/x-www-form-urlencoded")
	case "json":
		r.Set("Content-Type", "application/json")
	}
	return r
}
