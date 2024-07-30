package main

import (
	"bytes"
	"image"
	"log"

	_ "embed"
	"image/png"

	eb "github.com/hajimehoshi/ebiten"
)

//go:embed assets/icon.png
var iconData []byte

const (
	WIDTH        = 1280
	HEIGHT       = 720
	FPS          = 10
	WINDOW_TITLE = "hello mother"
)

func main() {
	iconImage, err := png.Decode(bytes.NewReader(iconData))
	if err != nil {
		log.Fatal(err)
	}

	eb.SetWindowIcon([]image.Image{iconImage})

	eb.SetWindowSize(WIDTH, HEIGHT)
	eb.SetWindowTitle(WINDOW_TITLE)
	eb.SetMaxTPS(FPS)

	if err := eb.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
