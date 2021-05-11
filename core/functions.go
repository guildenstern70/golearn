/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package core

func singleReturnValue(a, b int) int {
	return a + b
}

func multipleReturnValues(a, b int,
	strA, strB string) (string, int) {
	return strA + strB, a + b
}

// Functions returns crazy useless numbers
func Functions() int {

	strX, x := multipleReturnValues(5, 7, "53223ff", "7")
	y := singleReturnValue(23, 234)

	return len(strX) + x + y
}
