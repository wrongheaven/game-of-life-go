package gui

import (
	eb "github.com/hajimehoshi/ebiten"
	ebu "github.com/hajimehoshi/ebiten/ebitenutil"
)

func PrintText(screen *eb.Image, text string, pos Point) {
	ebu.DebugPrintAt(screen, text, pos.X, pos.Y)
}
