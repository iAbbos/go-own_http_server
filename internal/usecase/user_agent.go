package usecase

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
)

func UserAgent(reqHeaders map[string]string) *entity.Response {
	userAgent := reqHeaders["User-Agent"]

	resp := entity.NewResponse()

	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(200, "OK")
	resp.SetHeader("Content-Type", "text/plain")
	resp.SetBody([]byte(userAgent))

	return resp
}
