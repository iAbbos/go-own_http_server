package handler

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"github.com/codecrafters-io/http-server-starter-go/internal/entity/types"
	configpkg "github.com/codecrafters-io/http-server-starter-go/internal/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/internal/pkg/parser"
	"github.com/codecrafters-io/http-server-starter-go/internal/pkg/writer"
	"github.com/codecrafters-io/http-server-starter-go/internal/usecase"
	"io"
	"net"
	"strings"
)

type HandleOption struct {
	Conn   net.Conn
	Config *configpkg.Config
}

func HandleConnection(option HandleOption) error {
	defer option.Conn.Close()
	prs := parser.NewParser(option.Conn)
	req, err := prs.Parse()
	if err != nil && err != io.EOF {
		return err
	}

	newWriter := writer.NewWriter(option.Conn)

	var res *entity.Response

	switch req.Method {
	case types.METHOD_GET:
		if req.Target == "/" {
			res = usecase.BaseResponse()
		} else if strings.HasPrefix(req.Target, "/echo/") {
			res = usecase.Echo(req.Target)
		} else {
			res = usecase.NotFoundError()
		}
	default:
		res = usecase.BadRequestError()
	}

	err = newWriter.Write(res)
	if err != nil {
		return err
	}

	return nil
}
