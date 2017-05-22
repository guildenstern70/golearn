package core

import (
	"fmt"
	"strconv"
)

func multipleReturnValues( a, b int,
			   str_a, str_b string) (string,int) {
	return str_a + str_b, a + b
}

func PrintPlanet(shortname string,
	desc string,
	radius int32,
	mass float64) {

	str_x, x := multipleReturnValues(5, 7, "5", "7")

	fmt.Println(str_x + " " + strconv.Itoa(x))
	fmt.Printf("\t%s\t%s\t%d\t%f\n", shortname, desc, radius, mass)
}
