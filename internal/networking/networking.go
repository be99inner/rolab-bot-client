package networking

import (
	"log"

	"github.com/be99inner/rolab-bot-client/internal/gameinterface"
	"github.com/be99inner/rolab-bot-utility/networking"
)

// ConnectAndServe connects to the server and handle communication
func ConnectAndServe(url string) {
	conn, err := networking.ConnectWebSocket(url)
	if err != nil {
		log.Fatalf("Connect error: %v\n", err)
	}
	defer conn.Close()

	// capture and send the game interface
	img := gameinterface.CaptureGameInterface()
	encodedImage := gameinterface.EncodeImage(img)

	data := networking.GameData{
		EventType: "game_interface",
		Payload:   encodedImage,
	}

	err = networking.SendData(conn, data)
	if err != nil {
		log.Fatalf("Send error: %v\n", err)
	}

	response, err := networking.ReceiveData(conn)
	if err != nil {
		log.Fatalf("Receive error: %v\n", err)
	}

	log.Printf("Receive response: %+v\n", response)
}
