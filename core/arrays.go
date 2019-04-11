/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2019.
 * MIT Licence - See LICENSE
 */

package core

import "fmt"

// Arrays are fixed-size in Go
// Dynamic arrays are called 'slices' - see slices.go
var val = [10]int{44, 72, 12, 55, 64, 1, 4, 90, 13, 54}
var days = [7]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

// Arrays prints out the example
func Arrays() {
	fmt.Println(len(val), cap(val))
	fmt.Println(days)
}
