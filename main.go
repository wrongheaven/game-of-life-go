package main

import (
	"bytes"
	"image"
	"log"

	_ "embed"
	"image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/wrongheaven/game-of-life-go/consts"
	"github.com/wrongheaven/game-of-life-go/gui"
)

//go:embed assets/icon.png
var iconData []byte

var grid gui.Grid

func main() {
	iconImage, err := png.Decode(bytes.NewReader(iconData))
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowIcon([]image.Image{iconImage})
	ebiten.SetWindowPosition(2500, 200)
	ebiten.SetWindowSize(consts.WIDTH, consts.HEIGHT)
	ebiten.SetWindowTitle(consts.WINDOW_TITLE)
	ebiten.SetMaxTPS(consts.FPS)

	for row := range consts.CELL_ROWS {
		for col := range consts.CELL_COLS {
			grid.AddCell(col, row)
		}
	}

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
