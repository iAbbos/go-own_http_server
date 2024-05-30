package usecase

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"os"
	"strings"
)

func CheckFile(target string, dir string) *entity.Response {
	fileName := strings.TrimPrefix(target, "/files/")

	fileInfo, err := os.Stat(dir + "/" + fileName)
	if err != nil {
		return NotFoundError()
	} else if fileInfo.IsDir() {
		return ForbiddenError()
	}

	file, err := os.ReadFile(dir + "/" + fileName)
	if err != nil {
		return InternalError()
	}

	resp := entity.NewResponse()

	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(200, "OK")
	resp.SetHeader("Content-Type", "application/octet-stream")
	resp.SetBody(file)

	return resp
}
