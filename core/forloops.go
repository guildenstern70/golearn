package core

import "fmt"

// ForLoops just print absurd lines out of loops
func ForLoops() {

	// Classic
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	for y := uint32(0); y < 100; y++ {
		fmt.Println(y)
	}

	// Iterator
	var fruits = [3]string{"Apple", "Banana", "Kiwi"}
	for f := range fruits {
		fmt.Println(fruits[f])
	}

	// As while
	var k = 1
	for k < 10 {
		k++
		fmt.Println(k)
	}

}
