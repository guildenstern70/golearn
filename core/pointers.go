/*
 * Project GoLearn
 * Copyright (c) Alessio Saltarin 2021-22
 * Licensed under MIT Licence - See LICENSE
 */

package core

import "fmt"

// Pointers shows why pointers are useful
func Pointers() string {

	var xvalPtr *float64
	var xval = 1500.14

	var usdPtr *currency
	var usd = currency{
		"US Dollar",
		"US"}

	xvalPtr = &xval
	usdPtr = &usd

	return fmt.Sprintf("%f %s", *xvalPtr, *usdPtr)
}
