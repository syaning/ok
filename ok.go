package ok

const (
	MethodGET  = "GET"
	MethodPOST = "POST"
)

type Request struct {
	method  string
	url     string
	headers map[string]string
}

func (r *Request) Method(method string) {
	r.method = method
}

func (r *Request) Url(url string) {
	r.url = url
}

func (r *Request) Header(field string, value string) {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	r.headers[field] = value
}
