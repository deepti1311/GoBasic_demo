package main

import "fmt"

func main() {

	var arr1 = [5]int{}                           // to initialized
	var arr2 = [5]int{1, 2}                       //partially initialized
	var arr3 = [5]string{"1", "2", "3", "4", "5"} // fully initialize

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//assign value to particular index of array
	a := [3]int{1: 1, 2: 2}
	fmt.Println(a)

	//lenght of an array
	fmt.Println("lenght of array", len(a))
	b := [...]string{"1", "3", "5", "xcdax", "sdcawe"}
	fmt.Println("lenght of array", len(b))

}
