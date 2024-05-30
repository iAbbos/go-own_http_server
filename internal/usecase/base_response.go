package usecase

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"github.com/codecrafters-io/http-server-starter-go/internal/entity/types"
	"net/http"
)

func BaseResponse() *entity.Response {
	resp := entity.NewResponse()
	resp.SetVersion(types.VERSION_1_1)
	resp.SetStatus(http.StatusOK, "OK")

	return resp
}
