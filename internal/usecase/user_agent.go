package usecase

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
)

func UserAgent(headers map[string]string) *entity.Response {
	userAgent := headers["User-Agent"]

	resp := entity.NewResponse()

	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(200, "OK")
	resp.SetHeader("Content-Type", "text/plain")
	resp.SetBody([]byte(userAgent))

	return resp
}
