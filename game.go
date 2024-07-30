package main

import (
	"fmt"
	"image/color"

	eb "github.com/hajimehoshi/ebiten"
	"github.com/wrongheaven/game-of-life-go/gui"
)

type Game struct{}

var frameCount = 0

// var fpsTotal = 0

func (g *Game) Update(*eb.Image) error {
	frameCount++
	return nil
}

func (g *Game) Draw(screen *eb.Image) {

	var bgColor = color.RGBA{50, 127, 127, 255}
	screen.Fill(bgColor)
	gui.PrintText(screen, fmt.Sprintf("frameCount: %d", frameCount), gui.Point{X: 0, Y: 0})
	gui.PrintText(screen, fmt.Sprintf("%.1f", eb.CurrentTPS()), gui.Point{X: 0, Y: 20})
	// ebu.DebugPrintAt(screen, fmt.Sprintf("%.1f", avgFps), 0, 40)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return WIDTH, HEIGHT
}
