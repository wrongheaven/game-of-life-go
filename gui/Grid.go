package gui

import (
	"errors"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

type Grid struct {
	Cells []Cell
}

func (g *Grid) AddCell(x int, y int) {
	var state CellState = CellState(rand.Intn(2))
	g.Cells = append(g.Cells, Cell{}.New(x, y, state))
}

// Executes a callback function on all the cells in the grid.
func (g *Grid) Each(callback func(c *Cell)) {
	for i := range g.Cells {
		callback(&g.Cells[i])
	}
}

func (g *Grid) AdvanceTick() {
	var newCells []Cell = []Cell{}
	for _, cell := range g.Cells {
		neighbors := g.GetNeighbors(cell)
		aliveCount := 0
		for _, neighbor := range neighbors {
			if neighbor.State == Alive {
				aliveCount++
			}
		}

		if aliveCount < 2 || aliveCount > 3 {
			cell.State = Dead
		} else if aliveCount == 3 {
			cell.State = Alive
		} else {
			// do nothing
		}

		newCells = append(newCells, cell)
	}
	g.Cells = newCells
}

func (g *Grid) GetNeighbors(cell Cell) []Cell {
	x, y := cell.Coords.X, cell.Coords.Y
	var neighbors []Cell = []Cell{}

	for row := -1; row <= 1; row++ {
		for col := -1; col <= 1; col++ {
			if !(row == 0 && col == 0) {
				neighborCell, err := g.GetCell(x+col, y+row)
				if err == nil {
					neighbors = append(neighbors, neighborCell)
				}
			}
		}
	}

	return neighbors
}

func (g *Grid) GetCell(col int, row int) (Cell, error) {
	for _, cell := range g.Cells {
		if cell.Coords.X == col && cell.Coords.Y == row {
			return cell, nil
		}
	}
	return Cell{}, errors.New("cell not found")
}

func (g Grid) Draw(screen *ebiten.Image) {
	for _, cell := range g.Cells {
		cell.Draw(screen)
	}
}
