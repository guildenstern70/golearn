package core

import "fmt"

func Pointers() {

	var xvalPtr *float64
	var xval = 1500.14

	var usdPtr *currency
	var usd = currency{
		"US Dollar",
		"US"}

	xvalPtr = &xval
	usdPtr = &usd

	fmt.Println(xvalPtr, xval, *xvalPtr, *usdPtr)
}
