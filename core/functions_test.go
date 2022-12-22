/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021-22
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "testing"

func TestFunctions(t *testing.T) {

	result := Functions()
	expectedResult := 277

	if result == expectedResult {
		t.Logf("Functions Test: expected %v - got %v", result, expectedResult)
	} else {
		t.Errorf("Functions Test ERROR: expected %v - got %v", result, expectedResult)
	}

}
