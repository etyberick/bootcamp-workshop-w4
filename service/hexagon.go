package service

import (
	"fmt"
	"math"
)

// Hexagon struct definition
type Hexagon struct {
	shape
	side float64
}

// NewHexagon returns a new Hexagon struct
func NewHexagon(n, st string, s float64) *Hexagon {
	return &Hexagon{
		shape: shape{
			name:      n,
			shapeType: st,
		},
		side: s,
	}
}

// Area - calculates Hexagon's area
func (h *Hexagon) Area() float64 {
	return 6 * h.side
}

// Perimeter - calculates Hexagon's perimeter
func (h *Hexagon) Perimeter() float64 {
	return 3 * math.Sqrt(3) / 2 * math.Pow(h.side, 2)
}

// Info - displays Hexagon's information
func (h *Hexagon) Info() string {
	return fmt.Sprintf("Howdy, my name is '%s', I'm a '%s' with side [%.03f units], area [%.03f units] and perimeter [%.03f units²]\n",
		h.name, h.shapeType, h.side, h.Area(), h.Perimeter())
}

// NewSide sets a new side to hexagon h
func (h *Hexagon) NewSide(s float64) {
	h.side = s
}
