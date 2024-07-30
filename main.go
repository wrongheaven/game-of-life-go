package main

import (
	"bytes"
	"image"
	"log"

	_ "embed"
	"image/png"

	"github.com/hajimehoshi/ebiten"
)

//go:embed assets/icon.png
var iconData []byte

const (
	WIDTH        = 600
	HEIGHT       = 600
	FPS          = 30
	WINDOW_TITLE = "hello mother"
)

func main() {
	iconImage, err := png.Decode(bytes.NewReader(iconData))
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowIcon([]image.Image{iconImage})

	ebiten.SetWindowPosition(2500, 200)
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle(WINDOW_TITLE)
	ebiten.SetMaxTPS(FPS)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
