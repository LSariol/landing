package main

import (
	"os"

	dashboardrouter "github.com/mobasity-web-landing/internal/apps/dashboard/router"
	"github.com/mobasity-web-landing/internal/envs"
	"github.com/mobasity-web-landing/internal/server"
)

func main() {

	var envVars []string = []string{
		"DASHBOARD_ADDRESS",
		"DASHBOARD_PORT",
	}

	if err := envs.Load(envVars); err != nil {
		panic(err)
	}
	addr := os.Getenv("DASHBOARD_ADDRESS") + ":" + os.Getenv("DASHBOARD_PORT")

	mux := dashboardrouter.DefineRoutes()

	srv := server.NewServer(mux, addr)

	server.Run(srv)

}
