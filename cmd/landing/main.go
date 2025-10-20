package main

import (
	"github.com/joho/godotenv"
	"github.com/mobasity-web-landing/internal/server"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	server.Run()

}
