package main

import (
	"log"

	"github.com/be99inner/rolab-bot-client/internal/networking"
)

var serverAddr = "localhost:3000"

func main() {
	log.Printf("Client is connecting to server at %s...\n", serverAddr)

	url := "ws://" + serverAddr + "/ws"
	networking.ConnectAndServe(url)
}
