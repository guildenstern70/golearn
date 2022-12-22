/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021-22
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "testing"

func TestPointers(t *testing.T) {

	result := Pointers()
	expectedResult := "1500.140000 {US Dollar US}"

	if result == expectedResult {
		t.Logf("Pointers Test: expected %v - got %v", result, expectedResult)
	} else {
		t.Errorf("Pointers Test ERROR: expected %v - got %v", result, expectedResult)
	}

}
