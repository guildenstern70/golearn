package core

import "fmt"

var val = [100]int{44, 72, 12, 55, 64, 1, 4, 90, 13, 54}
var days = [7]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

// Arrays prints out the example
func Arrays() {
	fmt.Println(len(val), cap(val))
	fmt.Println(days)
}
