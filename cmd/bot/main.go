package main

import (
	"os"

	botrouter "github.com/mobasity-web-landing/internal/apps/bot/router"
	"github.com/mobasity-web-landing/internal/envs"
	"github.com/mobasity-web-landing/internal/server"
)

func main() {

	var envVars []string = []string{
		"BOT_ADDRESS",
		"BOT_PORT",
	}

	if err := envs.Load(envVars); err != nil {
		panic(err)
	}
	addr := os.Getenv("BOT_ADDRESS") + ":" + os.Getenv("BOT_PORT")

	mux := botrouter.DefineRoutes()

	srv := server.NewServer(mux, addr)

	server.Run(srv)

}
