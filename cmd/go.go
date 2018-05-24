package main

import (
	"fmt"
	"github.com/guildenstern70/golearn/core"
)

// VERSION OF GO-LEARN
const VERSION = "0.2"

func main() {
	fmt.Printf("GoLearn v.%s", VERSION)


	fmt.Println("\n== VARIABLES ==")
	core.OnlyInFunction()
	fmt.Println("")

	fmt.Println("\n== FOR LOOPS ==")
	core.ForLoops()

	fmt.Println("\n== IF/SWITCH ==")
	core.IfSwitch(20)

	fmt.Println("\n== ARRAYS ==")
	core.Arrays()

	fmt.Println("\n== SLICES ==")
	core.Slices()

	fmt.Println("\n== STRUCTS ==")
	core.Structs()

	fmt.Println("\n== POINTERS ==")
	core.Pointers()

	fmt.Println("\n== CLASSES ==")
	core.Triangles()

	fmt.Println("\n== STRINGS ==")
	core.Strings()

	fmt.Println("\n== INTERFACES ==")
	core.Interfaces()
}
