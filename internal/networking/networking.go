package networking

import (
	"log"
	"time"

	"github.com/be99inner/rolab-bot-client/internal/gameinterface"
	"github.com/be99inner/rolab-bot-client/internal/mock"
	"github.com/be99inner/rolab-bot-utility/networking"
)

// ConnectAndServe connects to the server and handle communication
func ConnectAndServe(url string) {
	conn, err := networking.ConnectWebSocket(url)
	if err != nil {
		log.Fatalf("Connect error: %v\n", err)
	}
	defer conn.Close()

	// ConnectAndServe connects to the server and handles communication
	for {
		// img := gameinterface.CaptureGameInterface()
		img := mock.GenerateRandomImage()
		encodedImage := gameinterface.EncodeImage(img)

		data := networking.GameData{
			EventType: "game_interface",
			Payload:   encodedImage,
		}

		err = networking.SendData(conn, data)
		if err != nil {
			log.Fatalf("Send error: %v\n", err)
		}

		// Receive and log the server's response
		response, err := networking.ReceiveData(conn)
		if err != nil {
			log.Printf("Receive error: %v\n", err)
			continue
		}

		log.Printf("Receive response: %+v\n", response)

		time.Sleep(10 * time.Second)
	}
}
