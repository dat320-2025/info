package utils

import "fmt"

func Add(a int, b int) int {
	return a + b
}

// Value type example: struct
type Point struct {
	X, Y int
}

// Method with value receiver (operates on a copy)
func (p Point) MoveByValue(dx, dy int) {
	p.X += dx
	p.Y += dy
	fmt.Printf("Inside MoveByValue: %+v\n", p)
}

// Method with pointer receiver (operates on the original)
func (p *Point) MoveByReference(dx, dy int) {
	p.X += dx
	p.Y += dy
	fmt.Printf("Inside MoveByReference: %+v\n", p)
}

// Demo demonstrates the difference between value and reference receivers.
func DemoPoint() {
	p := Point{X: 1, Y: 2}
	fmt.Printf("Original Point: %+v\n", p)

	p.MoveByValue(10, 20)
	fmt.Printf("After MoveByValue: %+v\n", p)

	p.MoveByReference(10, 20)
	fmt.Printf("After MoveByReference: %+v\n", p)
}
