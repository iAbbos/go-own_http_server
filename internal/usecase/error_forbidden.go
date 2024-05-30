package usecase

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"net/http"
)

func ForbiddenError() *entity.Response {
	resp := entity.NewResponse()
	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(http.StatusForbidden, "Forbidden")

	return resp
}
