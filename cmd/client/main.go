package main

import (
	"flag"
	"log"

	"github.com/be99inner/rolab-bot-client/internal/networking"
)

var serverAddr = flag.String("addr", "localhost:3000", "HTTP Server address")

func main() {
	log.Printf("Client is connecting to server at %s...\n", *serverAddr)

	url := "ws://" + *serverAddr + "/ws"
	networking.ConnectAndServe(url)
}
