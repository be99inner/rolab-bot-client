package mock

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	// "time"
)

// GenerateRandomImage genreates a random image
func GenerateRandomImage() image.Image {
	//  // random.Seed() is deprecated on new Go
	// rand.Seed(time.Now().UnixNano())
	width := 800
	height := 800

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Create a random background color
	bgColor := color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	// Add random shapes or patterns
	for i := 0; i < 10; i++ {
		rectColor := color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		}
		x0 := rand.Intn(width)
		y0 := rand.Intn(height)
		x1 := x0 + rand.Intn(100)
		y1 := y0 + rand.Intn(100)
		draw.Draw(img, image.Rect(x0, y0, x1, y1), &image.Uniform{C: rectColor}, image.Point{}, draw.Src)
	}

	return img
}
