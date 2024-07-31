package gui

type Color struct {
	R, G, B, A uint8
}

var (
	White = Color{R: 255, G: 255, B: 255, A: 255}
	Black = Color{R: 0, G: 0, B: 0, A: 255}
)
