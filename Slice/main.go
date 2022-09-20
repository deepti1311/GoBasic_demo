package main

import "fmt"

func main() {

	//syntax:slice_name := []datatype{values}
	//len() function - returns the length of the slice (the number of elements in the slice)
	//cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	//empty slice
	slice1 := []int{}
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)

	//slice having elements
	slice2 := []string{"Go", "slice", "Powerful"}
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(slice2)

	//how to create slice from array
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice := array[2:7]
	fmt.Println(slice)
	fmt.Println("lenght", len(slice))
	fmt.Println("capacity", cap(slice))

	// create slices using the make() function:
	slice3 := make([]int, 2, 8)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	//without capacity
	slice4 := make([]int, 2)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	//appent element to slice
	slice5 := append(slice4, 23, 45, 47)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	slice7 := append(slice2)
	fmt.Println(slice7)
	fmt.Println(len(slice7))
	fmt.Println(cap(slice7))

}
