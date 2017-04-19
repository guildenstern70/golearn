package core

import "fmt"

func PrintPlanet(shortname string,
	desc string,
	radius int32,
	mass float64) {
	fmt.Printf("\t%s\t%s\t%d\t%f\n", shortname, desc, radius, mass)
}
