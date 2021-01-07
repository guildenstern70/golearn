/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package core

// ForLoops returns absurd numbers out of loops
func ForLoops() uint32 {

	var sum uint32 = 0

	// Classic
	for x := 0; x < 10; x++ {
		sum += uint32(x) // type cast
	}

	// With type
	for y := uint32(0); y < 100; y++ {
		sum += y
	}

	// Iterator
	var fruits = [3]string{"Apple", "Banana", "Kiwi"}
	for f := range fruits {
		sum += uint32(f)
	}

	// As while
	var k = 1
	for k < 10 {
		k++
		sum += uint32(k)
	}

	return sum
}
