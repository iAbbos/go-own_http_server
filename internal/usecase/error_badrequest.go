package usecase

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"github.com/codecrafters-io/http-server-starter-go/internal/entity/types"
	"net/http"
)

func BadRequestError() *entity.Response {

	resp := entity.NewResponse()
	resp.SetVersion(types.VERSION_1_1)
	resp.SetStatus(http.StatusBadRequest, "Bad Request")

	return resp
}
