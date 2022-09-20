package main

import "fmt"

func main() {

	//single and multiple switch cases
	day := 8

	switch day {
	case 1:
		fmt.Println("day 1")
	case 2:
		fmt.Println("day 2")

	case 5:
		fmt.Println("day 5")

	}

	switch day {
	case 1, 2, 3:
		fmt.Println("day 1 2 3")
	case 4, 5:
		fmt.Println("day 4 5")

	default:
		fmt.Println("Invalid day")

	}

}
