package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/wrongheaven/game-of-life-go/consts"
	"github.com/wrongheaven/game-of-life-go/gui"
)

type Game struct{}

var prevTickMouseDown bool = false

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return consts.WIDTH, consts.HEIGHT
}

func (g *Game) Update(*ebiten.Image) error {
	initMouseClick()

	grid.Each(func(cell *gui.Cell) {
		// cell.MarkIfHovered()
	})

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{127, 127, 127, 255})

	grid.Draw(screen)
}

func initMouseClick() {
	mouseDown := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if true || mouseDown && !prevTickMouseDown {
		onMouseClick()
	}
	prevTickMouseDown = mouseDown
}
func onMouseClick() {
	grid.AdvanceTick()

	//

	// var clickedCell *gui.Cell
	// found := false
	// for i := range grid.Cells {
	// 	if grid.Cells[i].IsHovered() {
	// 		clickedCell = &grid.Cells[i]
	// 		found = true
	// 		break
	// 	}
	// }

	// if found {
	// 	clickedCell.FlipState()
	// }
}
