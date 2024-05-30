package usecase

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func SaveFile(req entity.Request, dir string) *entity.Response {
	fileName := strings.TrimPrefix(req.Target, "/files/")
	file, err := os.OpenFile(dir+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return InternalError()
	}
	defer file.Close()

	length, err := strconv.Atoi(req.Headers["Content-Length"])
	if err != nil {
		return BadRequestError()
	}

	buff := make([]byte, 4*1024)
	for length > 0 {
		println("length", length)
		n, err := req.Reader.Read(buff)
		length -= n
		if err != nil {
			break
		}
		file.Write(buff[:n])
	}

	if length != 0 {
		return BadRequestError()
	}

	resp := entity.NewResponse()

	resp.SetVersion("HTTP/1.1")
	resp.SetStatus(http.StatusCreated, "Created")

	return resp
}
