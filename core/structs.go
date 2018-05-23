package core

import "fmt"

type currency struct {
	name, country string
}

type price struct {
	value int
	cur   currency
}

type car struct {
	brand string
	model string
	price price
}

func Structs() {

	var euro = currency{
		"Euro",
		"EU"}

	var price10k = price{
		10000,
		euro}

	var mycar = car{
		brand: "Fiat",
		model: "500",
		price: price10k}
	fmt.Println(mycar)
}
