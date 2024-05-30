package entity

import "io"

type Request struct {
	Method  string
	Target  string
	Version string
	Headers map[string]string
	Body    io.Reader
}

func (r *Request) Marshal() []byte {
	return []byte{}
}
