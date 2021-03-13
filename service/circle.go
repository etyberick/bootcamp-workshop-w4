package service

import (
	"fmt"
	"math"
)

// Circle struct definition
type Circle struct {
	shape
	radius float64
}

// NewCircle returns a new Circle struct
func NewCircle(n, st string, r float64) *Circle {
	return &Circle{
		shape: shape{
			name:      n,
			shapeType: st,
		},
		radius: r,
	}
}

// Area - calculates Circle's area
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

// Perimeter - calculates Circle's perimeter
func (c *Circle) Perimeter() float64 {
	return math.Pi * 2 * c.radius
}

// Info - displays Circle's information
func (c *Circle) Info() string {
	return fmt.Sprintf("Hola, my name is '%s', I'm a '%s' with radius [%.03f units], area [%.03f units] and perimeter [%.03f units²]\n",
		c.name, c.shapeType, c.radius, c.Area(), c.Perimeter())
}

// NewRadius sets a new radius to circle c
func (c *Circle) NewRadius(r float64) {
	c.radius = r
}
