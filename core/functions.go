package core

import (
	"fmt"
	"strconv"
)

func multipleReturnValues(a, b int,
	strA, strB string) (string, int) {
	return strA + strB, a + b
}

// PrintPlanet prints out details about a planet
func PrintPlanet(shortname string,
	desc string,
	radius int32,
	mass float64) {

	strX, x := multipleReturnValues(5, 7, "5", "7")

	fmt.Println(strX + " " + strconv.Itoa(x))
	fmt.Printf("\t%s\t%s\t%d\t%f\n", shortname, desc, radius, mass)
}
