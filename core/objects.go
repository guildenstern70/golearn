package core

import (
	"fmt"
	"math"
	"strconv"
)

// Triangle object
type triangle struct {
	base   float64
	height float64
}

func (t *triangle) area() float64 {
	return (t.base * t.height) / 2
}

func (t *triangle) perim() float64 {
	ipotenusa := math.Sqrt((.5*t.base)*(.5*t.base) + t.height*t.height)
	return 2*ipotenusa + t.base
}

func newTriangle(b, h float64) *triangle {
	t := &triangle{}
	t.height = h
	t.base = b
	return t
}

func Triangles() {

	t1 := newTriangle(10, 20)
	area := strconv.FormatFloat(t1.area(), 'f', 3, 32)
	fmt.Println("Area of triangle with b=10 and h=20 => " + area)

}
