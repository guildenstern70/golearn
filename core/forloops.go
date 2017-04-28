package core

import "fmt"

func ForLoops() {

	// Classic
	for x:=0; x < 10; x++ {
		fmt.Println(x)
	}

	// Iterator
	var fruits [3]string = [3]string{ "Apple", "Banana", "Kiwi" }
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

