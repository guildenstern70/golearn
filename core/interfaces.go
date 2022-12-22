/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021-22
 * Licensed under MIT Licence - See LICENSE
 */

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

// Interfaces demonstrates the Interfaces code usage
func Interfaces(a, b float64) float64 {
	t1 := newTriangle(a, a*4)
	r1 := newRect(b, b+4)
	return t1.area() + r1.area()
}
