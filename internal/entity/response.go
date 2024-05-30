package entity

import "fmt"

type Response struct {
	Version       string
	StatusCode    int
	StatusMessage string
	Headers       map[string]string
	Body          []byte
}

func NewResponse() *Response {
	return &Response{
		Headers: make(map[string]string),
	}
}

func (r *Response) Marshal() []byte {
	res := make([]byte, 0)
	startLine := []byte(fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, r.StatusMessage))
	res = append(res, startLine...)

	if r.Body != nil {
		r.Headers["Content-Length"] = fmt.Sprintf("%d", len(r.Body))
	}

	for k, v := range r.Headers {
		res = append(res, []byte(fmt.Sprintf("%s: %s\r\n", k, v))...)
	}
	res = append(res, []byte("\r\n")...)

	if r.Body != nil {
		res = append(res, r.Body...)
	} else {
		res = append(res, []byte("\r\n")...)
	}

	return res
}

func (r *Response) SetVersion(version string) {
	r.Version = version
}

func (r *Response) SetStatus(code int, message string) {
	r.StatusCode = code
	r.StatusMessage = message
}

func (r *Response) SetHeader(key string, value string) {
	r.Headers[key] = value
}

func (r *Response) SetBody(body []byte) {
	r.Body = body
}
