package main

import (
	"fmt"
	"github.com/guildenstern70/golearn/core"
)

const VERSION = "0.2"

func main() {
	fmt.Printf("GoLearn v.%s", VERSION)
	core.OnlyInFunction()
	core.ForLoops()
	core.Test()
	core.Arrays()
	core.Slices()
	core.Structs()
	core.Pointers()
	core.Triangles()
	core.Strings()
	core.Interfaces()
}
