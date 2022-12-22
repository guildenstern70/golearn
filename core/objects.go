/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021-22
 * Licensed under MIT Licence - See LICENSE
 */

package core

import (
	"math"
	"strconv"
)

// Triangle object. Note that "triangle" is "private" because it's written in lowercase.
// To use it as a library type, the name must be exported, that is written in Title case,
// ie: Triangle.
// Even the "base" and "height" attributes are "private", and can be exported using
// Title case
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

// Objects Triangles shows how to use a class of objects
func Objects(h, b float64) string {

	t1 := newTriangle(b, h)
	area := strconv.FormatFloat(t1.area(), 'f', 3, 32)
	return area

}
