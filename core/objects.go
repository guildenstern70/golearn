package core

import (
	"fmt"
	"strconv"
)

// Triangle object
type triangle struct {
	base float32
	height float32
}

func (t *triangle) area() float32 {
	return (t.base * t.height) / 2
}

func newTriangle(b, h float32) *triangle {
	t := &triangle{}
	t.height = h
	t.base = b
	return t
}

func Triangles() {

	t1 := newTriangle(10,20)
	area := strconv.FormatFloat(float64(t1.area()), 'f', 3, 32)
	fmt.Println("Area of triangle with b=10 and h=20 => " + area)

}




