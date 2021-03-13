package service

// ShapeClient - Struct with all figures
type ShapeClient struct {
	Square  *Square
	Circle  *Circle
	Hexagon *Hexagon
}

// New - Creates a client with all the figures
func New() ShapeClient {
	// create a new Square, Circle and Hexagon
	square := NewSquare("S", "square", 4)
	circle := NewCircle("C", "circle", 10.2)
	hexagon := NewHexagon("H", "hexagon", 5)

	return ShapeClient{square, circle, hexagon}
}

// Shape interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
	Info() string
}

// shape struct definition
type shape struct {
	name      string
	shapeType string
}

// NewName sets a new name to shape s
func (s *shape) NewName(n string) {
	s.name = n
}

// NewShapeType sets a new FigureType to shape s
func (s *shape) NewShapeType(st string) {
	s.shapeType = st
}
