/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021-22
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "testing"

func TestArrays(t *testing.T) {
	result := Arrays()
	expectedResult := 27

	if result == expectedResult {
		t.Logf("Arrays Test: expected %v - got %v", result, expectedResult)
	} else {
		t.Errorf("Arrays Test ERROR: expected %v - got %v", result, expectedResult)
	}

}
