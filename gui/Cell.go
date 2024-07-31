package gui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/wrongheaven/game-of-life-go/consts"
)

type Cell struct {
	Coords Point
	Pos    Point
	State  CellState
}

type CellState int

const (
	Dead  CellState = 0
	Alive CellState = 1
)

func (c Cell) New(x int, y int, state CellState) Cell {
	return Cell{
		Coords: Point{X: x, Y: y},
		Pos:    Point{X: x * consts.CELL_SIZE, Y: y * consts.CELL_SIZE},
		State:  state,
	}
}

func (c *Cell) IsHovered() bool {
	mouseX, mouseY := ebiten.CursorPosition()
	return (mouseX >= c.Pos.X &&
		mouseX < c.Pos.X+consts.CELL_SIZE &&
		mouseY >= c.Pos.Y &&
		mouseY < c.Pos.Y+consts.CELL_SIZE)
}

func (c *Cell) FlipState() {
	switch c.State {
	case Alive:
		c.State = Dead
	case Dead:
		c.State = Alive
	default:
		panic("invalid cell state: " + fmt.Sprintf("%d", c.State))
	}
}

func (c *Cell) Draw(screen *ebiten.Image) {
	var col color.RGBA

	switch c.State {
	case Alive:
		col = color.RGBA(White)
	case Dead:
		col = color.RGBA(Black)
	default:
		panic("invalid cell state: " + fmt.Sprintf("%d", c.State))
	}

	ebitenutil.DrawRect(
		screen,
		float64(c.Pos.X)+1,
		float64(c.Pos.Y)+1,
		float64(consts.CELL_SIZE)-1,
		float64(consts.CELL_SIZE)-1,
		col,
	)
}
