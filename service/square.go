package service

import (
	"fmt"
	"math"
)

// Square struct definition
type Square struct {
	shape
	side float64
}

// NewSquare returns a new Square struct
func NewSquare(n, st string, s float64) *Square {
	return &Square{
		shape: shape{
			name:      n,
			shapeType: st,
		},
		side: s,
	}
}

// Area - calculates Square's area
func (s *Square) Area() float64 {
	return 4 * s.side
}

// Perimeter - calculates Square's perimeter
func (s *Square) Perimeter() float64 {
	return math.Pow(s.side, 2)
}

// Info - displays Square's information
func (s *Square) Info() string {
	return fmt.Sprintf("Hello, my name is '%s', I'm a '%s' with side [%.03f units], area [%.03f units] and perimeter [%.03f units²]\n",
		s.name, s.shapeType, s.side, s.Area(), s.Perimeter())
}

// NewSide sets a new side to square sq
func (s *Square) NewSide(side float64) {
	s.side = side
}
