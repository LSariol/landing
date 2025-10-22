package main

import (
	"os"

	wwwrouter "github.com/mobasity-web-landing/internal/apps/www/router"
	"github.com/mobasity-web-landing/internal/envs"
	"github.com/mobasity-web-landing/internal/server"
)

func main() {

	var envVars []string = []string{
		"WWW_ADDRESS",
		"WWW_PORT",
	}

	if err := envs.Load(envVars); err != nil {
		panic(err)
	}

	addr := os.Getenv("WWW_ADDRESS") + ":" + os.Getenv("WWW_PORT")

	mux := wwwrouter.DefineRoutes()

	srv := server.NewServer(mux, addr)

	server.Run(srv)

}
