package gameinterface

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"log"
)

// CaptureGameInterface captures the game interface as an image
func CaptureGameInterface() image.Image {
	// Placeholder function to capture a screenshot of the game interface
	// In practice, this function should capture the actual game screen
	img := image.NewRGBA(image.Rect(0, 0, 800, 600))
	return img
}

// EncodeImage encodes the imagess to a base64 string
func EncodeImage(img image.Image) string {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Fatalf("Failed to encode image: %v\n", err)
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}
