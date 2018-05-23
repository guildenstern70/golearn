package core

import (
	"fmt"
)

// Interface
type shape interface {
	area() float64
	perim() float64
}

// Any object that implements an interface "is-a"
// object of the given interface (without need to specify it)
// Duck typing
type rect struct {
	length, height float64
}

func (r *rect) area() float64 {
	return r.length * r.height
}

func (r *rect) perim() float64 {
	return 2*r.length + 2*r.height
}

func newRect(l, h float64) *rect {
	r := &rect{}
	r.height = h
	r.length = l
	return r
}

// Polymorphic type
func shapeInfo(s shape) string {
	return fmt.Sprintf(
		"Area = %.2f, Perim = %.2f",
		s.area(), s.perim(),
	)
}

func Interfaces() {

	t1 := newTriangle(10, 20)
	r1 := newRect(12, 15)

	fmt.Println(shapeInfo(t1))
	fmt.Println(shapeInfo(r1))

}
