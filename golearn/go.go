/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021-23
 * Licensed under MIT Licence - See LICENSE
 */

package main

import (
	"fmt"
)

// VERSION OF GO-LEARN
const VERSION = "0.4"

func main() {
	fmt.Printf("\nGoLearn v.%s\n\n", VERSION)
	fmt.Printf("Run tests with \n")
	fmt.Printf("- go test github.com/guildenstern70/golearn/core \n\n")
}
