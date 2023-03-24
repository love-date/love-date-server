package main

import (
	"github.com/joho/godotenv"
	"log"
	"love-date/config"
	"love-date/delivery/httpsserver"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()

	server := httpsserver.NewHttpServer(conf.Server.Host, conf.Server.Port)
	server.Start()
}
