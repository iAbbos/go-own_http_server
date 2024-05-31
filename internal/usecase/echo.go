package usecase

import (
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

	encoding, ok := req.Headers["Accept-Encoding"]
	if strings.Contains(encoding, "gzip") && ok {
		resp.SetHeader("Content-Encoding", "gzip")
	}

	resp.SetBody([]byte(msg))

	return resp
}
