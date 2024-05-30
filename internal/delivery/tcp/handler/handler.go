package handler

import (
	"fmt"
	configpkg "github.com/codecrafters-io/http-server-starter-go/internal/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/internal/pkg/parser"
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
		fmt.Println("stage2")
		return err
	}

	fmt.Println("Request: ", req)

	//writer := writer.NewWriter(option.Conn)
	//
	//
	//
	//writer.Write(res)

	return nil
}
