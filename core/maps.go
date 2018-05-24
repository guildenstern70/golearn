package core

import "fmt"

var histogram = map[string]int{
	"Jan": 100, "Feb": 445, "Mar": 514, "Apr": 233,
	"May": 321, "Jun": 644, "Jul": 113, "Aug": 734,
	"Sep": 553, "Oct": 344, "Nov": 831, "Dec": 312}

// Maps prints out maps examples
func Maps() {

	fmt.Println(histogram["Jan"])

	// There is no functional programming in Go
	// so you must do that the old way
	var sum = 0
	for _, val := range histogram {
		if val > 100 {
			sum += val
		}
	}
	fmt.Printf("Sum of elements > 100 = %d", sum)

	hist := make(map[string]int)
	hist["Jan"] = 100
	hist["Feb"] = 445
	hist["Mar"] = 514

	for key, val := range hist {
		adjVal := int(float64(val) * 0.100)
		fmt.Printf("%s (%d):", key, val)
		for i := 0; i < adjVal; i++ {
			fmt.Print(".")
		}
		fmt.Println()
	}
}
