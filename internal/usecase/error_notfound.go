package usecase

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
)

func NotFoundError() *entity.Response {
	resp := entity.NewResponse()
	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(404, "Not Found")

	return resp
}
