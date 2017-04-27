package core

import "fmt"

var histogram map[string]int = map[string]int{
	"Jan": 100, "Feb": 445, "Mar": 514, "Apr": 233,
	"May": 321, "Jun": 644, "Jul": 113, "Aug": 734,
	"Sep": 553, "Oct": 344, "Nov": 831, "Dec": 312}

func Maps() {
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
