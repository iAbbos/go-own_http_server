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

	if req.Method == types.METHOD_GET {
		switch req.Target {
		case "/":
			res = usecase.BaseResponse()
		case "/echo/":
			res = usecase.Echo(req.Target)
		default:
			res = usecase.NotFoundError()
		}
	} else {
		res = usecase.BadRequestError()
	}

	err = newWriter.Write(res)
	if err != nil {
		return err
	}

	return nil
}
