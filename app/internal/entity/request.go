package entity

import "io"

type Request struct {
	Version string
	Method  string
	Target  string
	Headers map[string]string
	Reader  io.Reader
}
