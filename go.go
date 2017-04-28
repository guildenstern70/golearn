package main

import (
	"fmt"
	"github.com/guildenstern70/GoLearn/core"
)

const VERSION = "0.1"

func main() {
	fmt.Printf("GoLearn v.%s", VERSION)
	core.OnlyInFunction()
	core.Test()
	core.Arrays()
	core.Slices()
	core.Structs()
	core.Pointers()
}
