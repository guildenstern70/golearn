package core

import "fmt"

var months = []string{
	"Jan", "Feb", "Mar", "Apr",
	"May", "Jun", "Jul", "Aug",
	"Sep", "Oct", "Nov", "Dec",
}

var halfyr = months[:6]
var q1 = halfyr[:3]
var q2 = halfyr[3:6]
var q3 = months[6:9]
var q4 = months[9:]

func Slices() {
	// Make and add
	var somestrings = make([]string, 10)
	somestrings = append(somestrings, "XXX")

	fmt.Println(len(months), cap(months))
	fmt.Println(len(somestrings), cap(somestrings))

	for i := range months {
		fmt.Println(months[i])
	}

}