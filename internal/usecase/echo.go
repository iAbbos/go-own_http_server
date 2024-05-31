package usecase

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"strings"
)

func Echo(req entity.Request) *entity.Response {
	target := req.Target

	msg := strings.TrimPrefix(target, "/echo/")

	resp := entity.NewResponse()

	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(200, "OK")
	resp.SetHeader("Content-Type", "text/plain")

	fmt.Println("req.Headers: ", req.Headers)
	encoding, ok := req.Headers["Accept-Encoding"]

	fmt.Println("encoding: ", encoding)
	fmt.Println("ok: ", ok)

	if encoding == "gzip" && ok {
		resp.SetHeader("Accept-Encoding", encoding)
	}

	resp.SetBody([]byte(msg))

	return resp
}
