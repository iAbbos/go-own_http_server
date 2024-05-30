package server

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/internal/delivery/tcp/handler"
	"github.com/codecrafters-io/http-server-starter-go/internal/pkg/config"
	"net"
)

type Server struct {
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", s.Config.Server.Host+s.Config.Server.Port)
	if err != nil {
		return fmt.Errorf("tsp server error on listen: %s %w", s.Config.Server.Port, err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return fmt.Errorf("tcp server error on accept: %w", err)
		}
		go handler.HandleConnection(handler.HandleOption{
			Conn:   conn,
			Config: s.Config,
		})
	}
}
