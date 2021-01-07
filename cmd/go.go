/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package main

import (
	"fmt"
)

// VERSION OF GO-LEARN
const VERSION = "0.3"

func main() {
	fmt.Printf("GoLearn v.%s\n\n", VERSION)
	fmt.Printf("Run tests with \n")
	fmt.Printf("- go test \n\n")
}
