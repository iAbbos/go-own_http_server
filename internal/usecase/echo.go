package usecase

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"strings"
)

func Echo(target string) *entity.Response {
	fmt.Println(target)
	msg := strings.TrimPrefix(target, "/echo/")

	resp := entity.NewResponse()

	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(200, "OK")
	resp.SetHeader("Content-Type", "text/plain")
	resp.SetBody([]byte(msg))

	return resp
}
