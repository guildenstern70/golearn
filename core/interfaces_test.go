/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "testing"

func TestInterfaces(t *testing.T) {

	result := Interfaces(10.0, 14.5)
	expectedResult := 468.25

	if result == expectedResult {
		t.Logf("Interfaces Test: expected %v - got %v", result, expectedResult)
	} else {
		t.Errorf("Interfaces Test ERROR: expected %v - got %v", result, expectedResult)
	}

}
