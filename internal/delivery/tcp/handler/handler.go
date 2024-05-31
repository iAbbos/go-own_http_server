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

	if req.Target == "/" {
		res = usecase.BaseResponse()
	} else if strings.HasPrefix(req.Target, "/echo/") {
		res = usecase.Echo(req)
	} else if strings.HasPrefix(req.Target, "/user-agent") {
		res = usecase.UserAgent(req.Headers)
	} else if strings.HasPrefix(req.Target, "/files/") {
		dir := option.Config.FilesDir
		switch req.Method {
		case types.METHOD_GET:
			res = usecase.CheckFile(req.Target, dir)
		case types.METHOD_POST:
			res = usecase.SaveFile(req, dir)
		}
	} else {
		res = usecase.NotFoundError()
	}

	err = newWriter.Write(res)
	if err != nil {
		return err
	}

	return nil
}
