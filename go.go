package main

import (
	"fmt"
	"github.com/guildenstern70/golearn/core"
)

const VERSION = "0.1"

func main() {
	fmt.Printf("GoLearn v.%s", VERSION)
	core.OnlyInFunction()
}
