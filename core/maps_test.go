/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "testing"

func TestMaps(t *testing.T) {

	result := Maps()
	expectedResult := 5594

	if result == expectedResult {
		t.Logf("Maps Test: expected %v - got %v", result, expectedResult)
	} else {
		t.Errorf("Maps Test ERROR: expected %v - got %v", result, expectedResult)
	}

}
