package entity

type Response struct {
	HTTPVersion   string
	StatusCode    int
	StatusMessage string
	Headers       map[string]string
	Body          []byte
}

func (r *Response) Marshal() []byte {
	return []byte{}
}
