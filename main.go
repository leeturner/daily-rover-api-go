package main

import (
	"github.com/leeturner/daily-rover-api-go/pkg/server"
	"log"
)

const port = "8181"

func main() {
	httpServer := server.InitServer(port)
	err := httpServer.Start()
	if err != nil {
		log.Fatal(err)
	}
}
