/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2019.
 * MIT Licence - See LICENSE
 */

package core

import "fmt"

// Slices are dynamic length arrays
// aka collections
var months = []string{
	"Jan", "Feb", "Mar", "Apr",
	"May", "Jun", "Jul", "Aug",
	"Sep", "Oct", "Nov", "Dec",
}

var halfyr = months[:6]
var q1 = halfyr[:3]
var q2 = halfyr[3:6]
var q3 = months[6:9]
var q4 = months[9:]

// Slices shows how to use slices
// This function returns a slice of strings
func Slices() []string {
	// A slice can be created with the built-in function called make
	// Make 10 empty strings
	var somestrings = make([]string, 10)
	// Append 11th string
	somestrings = append(somestrings, "XXX")

	fmt.Println(len(months), cap(months))
	fmt.Println(len(somestrings), cap(somestrings))

	for i := range months {
		fmt.Println(months[i])
	}

	return createEmptySliceAndAppend()
}

func createEmptySliceAndAppend() []string {

	emptySlice := make([]string, 0)

	emptySlice = append(emptySlice, "Alessio")
	emptySlice = append(emptySlice, "Elena")
	emptySlice = append(emptySlice, "Isabella")

	return emptySlice

}
