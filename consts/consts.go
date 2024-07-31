package consts

import (
	"math"
)

const (
	WIDTH        = 1000
	HEIGHT       = 800
	FPS          = 10
	WINDOW_TITLE = "Game of Life"
	CELL_SIZE    = 20
)

var (
	CELL_COLS = int(math.Floor(WIDTH / CELL_SIZE))
	CELL_ROWS = int(math.Floor(HEIGHT / CELL_SIZE))
)
