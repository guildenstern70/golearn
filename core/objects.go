/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package core

import (
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

// Triangles shows how to use a class of objects
func Objects(h, b float64) string {

	t1 := newTriangle(b, h)
	area := strconv.FormatFloat(t1.area(), 'f', 3, 32)
	return area

}
