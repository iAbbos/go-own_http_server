package app

import (
	tspsrv "github.com/codecrafters-io/http-server-starter-go/internal/delivery/tcp/server"
	configpkg "github.com/codecrafters-io/http-server-starter-go/internal/pkg/config"
)

type App struct {
	Config *configpkg.Config
}

func NewApp(config *configpkg.Config) (*App, error) {
	return &App{
		Config: config,
	}, nil
}

func (a *App) Run() error {
	server := tspsrv.NewServer(a.Config)
	err := server.Run()
	if err != nil {
		return err
	}
	return nil
}
