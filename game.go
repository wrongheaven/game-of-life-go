package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/wrongheaven/game-of-life-go/gui"
)

type Cell struct {
	Pos   gui.Point
	Size  float64
	Color color.RGBA
}

func (c Cell) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen,
		float64(c.Pos.X),
		float64(c.Pos.Y),
		c.Size,
		c.Size,
		c.Color,
	)
}

type Grid struct {
	Cells []Cell
}

func (g Grid) Draw(screen *ebiten.Image) {
	for _, cell := range g.Cells {
		cell.Draw(screen)
	}
}

var grid Grid = Grid{
	Cells: []Cell{
		{
			Pos:   gui.Point{X: 0, Y: 0},
			Size:  60,
			Color: color.RGBA{0, 0, 0, 255},
		},
		{
			Pos:   gui.Point{X: 60, Y: 0},
			Size:  60,
			Color: color.RGBA{255, 255, 255, 255},
		},
		{
			Pos:   gui.Point{X: 0, Y: 60},
			Size:  60,
			Color: color.RGBA{255, 255, 255, 255},
		},
		{
			Pos:   gui.Point{X: 60, Y: 60},
			Size:  60,
			Color: color.RGBA{0, 0, 0, 255},
		},
	},
}

type Game struct{}

func (g *Game) Update(*ebiten.Image) error {
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {

	var bgColor = color.RGBA{50, 127, 127, 255}
	screen.Fill(bgColor)
	// gui.PrintText(screen, fmt.Sprintf("frameCount: %d", frameCount), gui.Point{X: 0, Y: 0})
	// gui.PrintText(screen, fmt.Sprintf("%.1f", eb.CurrentTPS()), gui.Point{X: 0, Y: 20})

	// ebitenutil.DrawRect(screen, 0, 0, 10, 10, color.RGBA{0, 0, 0, 255})
	grid.Draw(screen)
}
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return WIDTH, HEIGHT
}
