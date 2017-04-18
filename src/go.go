package main

import (
	"fmt"
	"net/littlelite/golearn"
)

const VERSION = "0.1"

func main() {
	fmt.Printf("GoLearn v.%s", VERSION)
	golearn.OnlyInFunction()
}
