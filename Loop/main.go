package main

import "fmt"

func main() {

	//for true {
	//	//fmt.Printf("this for loop will run forever.\n")
	//}

	for i := 0; i <= 5; i++ {
		fmt.Println("for loop running", i)
	}

	fmt.Println("for loop with condition")
	//for loop with condition
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("for loop with break condition ")
	//for loop with condition
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	adj := [2]string{"Big", "Tasty"}
	fruit := [3]string{"apple", "orange", "mango"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruit); j++ {

			fmt.Println(adj[i], fruit[j])

		}
	}

	fruits := [3]string{"1a", "3x", "4y"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}
}
