/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "testing"

func TestObjects(t *testing.T) {

	result := Objects(10.0, 5.4)
	expectedResult := "27.000"

	if result == expectedResult {
		t.Logf("Objects Test: expected %v - got %v", result, expectedResult)
	} else {
		t.Errorf("Objects Test ERROR: expected %v - got %v", result, expectedResult)
	}

}
