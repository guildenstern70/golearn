/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021-22
 * Licensed under MIT Licence - See LICENSE
 */

package core

var histogram = map[string]int{
	"Jan": 100, "Feb": 445, "Mar": 514, "Apr": 233,
	"May": 321, "Jun": 644, "Jul": 113, "Aug": 734,
	"Sep": 553, "Oct": 344, "Nov": 831, "Dec": 312}

// Maps prints out maps examples
func Maps() int {

	var febValue = histogram["Feb"]

	// There is no functional programming in Go
	// so you must do filtering the old way
	var sum = 0
	for _, val := range histogram {
		if val > 100 {
			sum += val
		}
	}

	sum += febValue

	// Initialize a new map
	hist := make(map[string]int)
	hist["Jan"] = 100
	hist["Feb"] = 445
	hist["Mar"] = 514

	for _, val := range hist {
		adjVal := int(float64(val) * 0.100)
		sum += adjVal
	}

	return sum
}
