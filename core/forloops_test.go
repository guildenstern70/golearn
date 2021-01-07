/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "testing"

func TestForLoops(t *testing.T) {

	var result uint32 = ForLoops()
	expectedResult := uint32(5052)

	if result == expectedResult {
		t.Logf("ForLoops Test: expected %v - got %v", result, expectedResult)
	} else {
		t.Errorf("ForLoops Test ERROR: expected %v - got %v", result, expectedResult)
	}

}
