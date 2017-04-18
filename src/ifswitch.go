package src

import "fmt"

func test() bool {

	a := 30

	if a > 20 {
		fmt.Println("It's greater than 20")
	} else {
		fmt.Println("It's less than or 20")
	}

	result := false

	switch a {
	default:
		result = false
	case 0:
		result = true
	case 10:
		result = true
	case 20:
		result = true
	}

	return result

}
