/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package core

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// Top level
var (
	topname         = "Earth"
	topdesc         = "Planet"
	topradius int32 = 6378
	topmass         = 5.972e+24
)

// Typed variables
var name, desc string
var radius int32
var mass float64
var active bool
var satellites = []string{
	"Moon",
}

// Automatic typed
var autostring = "Prova"
var autonumber = 0.0

// PrintPlanet prints out useless details about a planet
func PrintPlanet(shortname string,
	desc string,
	radius int32,
	mass float64) {

	strX, x := multipleReturnValues(5, 7, "5", "7")

	fmt.Println(strX + " " + strconv.Itoa(x))
	fmt.Printf("\t%s\t%s\t%d\t%f\n", shortname, desc, radius, mass)
}

// OnlyInFunction demonstrates the shortcut of declaring a variable (with value)
// inside a function
func OnlyInFunction() {

	shortname := "Neptune"
	desc := "Planet"
	radius := int32(24764)
	mass := 1.024e26

	PrintPlanet(topname, topdesc, topradius, topmass)
	PrintPlanet(shortname, desc, radius, mass)

}

// Typed constants
const d int32 = 111009

// Untyped constants
const m1 = math.Pi / 3.141592

// Constants declaration block
const (
	a1, a2 string        = "Mastering", "Go"
	b      rune          = 'G'
	c      bool          = false
	e      float32       = 2.71828
	f      float64       = math.Pi * 2.0e+3
	g      complex64     = 5.0i
	h      time.Duration = 4 * time.Second
)

// Enumeration const
const (
	StarHyperGiant = iota
	StarSuperGiant
	StarBrightGiant
	StarGiant
	StarSubGiant
	StarDwarf
	StarSubDwarf
	StarWhiteDwarf
	StarRedDwarf
	StarBrownDwarf
)
