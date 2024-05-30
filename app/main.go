package main

import (
	apppkg "github.com/codecrafters-io/http-server-starter-go/internal/app"
	configpkg "github.com/codecrafters-io/http-server-starter-go/internal/pkg/config"
	"log"
)

func main() {
	// init config
	config, err := configpkg.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// init app
	app, err := apppkg.NewApp(config)
	if err != nil {
		log.Fatal(err)
	}

	// run app
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
