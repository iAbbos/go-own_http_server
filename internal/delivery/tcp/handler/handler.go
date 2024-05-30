package handler

import (
	"fmt"
	configpkg "github.com/codecrafters-io/http-server-starter-go/internal/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/internal/pkg/parser"
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
	if err != nil {
		return err
	}

	fmt.Println(req)

	//writer := writer.NewWriter(option.Conn)
	//
	//
	//
	//writer.Write(res)

	return nil
}
