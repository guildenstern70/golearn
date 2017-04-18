package src

import (
	"fmt"
	"math"
	"time"
)

// Top level
var (
	topname       string  = "Earth"
	topdesc       string  = "Planet"
	topradius     int32   = 6378
	topmass       float64 = 5.972E+24
	topactive     bool    = true
	topsatellites []string
)

// Typed variables
var name, desc string = "Earth", "Planet"
var radius int32 = 6378
var mass float64 = 5.972E+24
var active bool = true
var satellites = []string{
	"Moon",
}

// Automatic typed
var autostring = "Prova"
var autonumber = 0.0

// Short
func onlyInFunction() {

	shortname := "Neptune"
	desc := "Planet"
	radius := 24764
	mass := 1.024e26

	fmt.Print(shortname, desc, radius, mass)

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
	StarHyperGiant  = iota
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
