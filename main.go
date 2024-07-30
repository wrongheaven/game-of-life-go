package main

import (
	"bytes"
	"image"
	"log"

	_ "embed"
	"image/png"

	eb "github.com/hajimehoshi/ebiten"
	"github.com/wrongheaven/game-of-life-go/utils"
)

//go:embed assets/icon.png
var iconData []byte

const (
	WIDTH        = 400
	HEIGHT       = 400
	FPS          = 10
	WINDOW_TITLE = "hello mother"
)

func main() {
	iconImage, err := png.Decode(bytes.NewReader(iconData))
	utils.Bekind(err)

	eb.SetWindowIcon([]image.Image{iconImage})

	eb.SetWindowSize(WIDTH, HEIGHT)
	eb.SetWindowTitle(WINDOW_TITLE)
	eb.SetMaxTPS(FPS)

	if err := eb.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
