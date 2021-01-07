/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "testing"

func TestIfSwitch(t *testing.T) {

	if IfSwitch(10) {
		t.Logf("Functions Test: expected %v - got %v", "true", "false")
	} else {
		t.Errorf("Functions Test ERROR: expected %v - got %v", "true", "false")
	}

}
