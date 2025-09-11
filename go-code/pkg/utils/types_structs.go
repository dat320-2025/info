package utils

import "math"

// Shape is an interface for geometric shapes.
type Shape interface {
	Area() float64
}

// Circle implements Shape.
type Circle struct {
	Radius float64
}
type SomeStuff struct{}

func (s *SomeStuff) Area() float64 {
	return 0.0 // Placeholder
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle implements Shape.
type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Demo demonstrates usage of Shape interface.
func DemoShapes() {
	shapes := []Shape{
		&SomeStuff{},
		&Circle{Radius: 3},
		&Rectangle{Width: 4, Height: 5},
	}

	for _, s := range shapes {
		println("Area:", s.Area())
	}
}
