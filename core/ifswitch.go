/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021-22
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "fmt"

// IfSwitch returns the result of an improbable test
func IfSwitch(a int) bool {

	if a > 20 {
		fmt.Println("It's greater than 20")
	} else {
		fmt.Println("It's less than or 20")
	}

	result := false

	switch a {
	default:
		result = false
	case 0:
		result = true
	case 10:
		result = true
	case 20:
		result = true
	}

	return result

}
