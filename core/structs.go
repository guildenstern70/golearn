package core

import "fmt"

type car struct {
	make, brand string;
	model       string
}
type currency struct {
	name, country string;
	code          int
}

func Structs() {

	var mycar = car{
		brand: "Fiat",
		model: "500" }
	fmt.Println(mycar)
}
