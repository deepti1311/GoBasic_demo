package main

import "fmt"

func main() {
	msg()
	name("deepti")

	fullName("deepti", "jakhotra")
	fmt.Println(sum(1, 2))
	fmt.Println(somefun(2, "Hi"))

	//ways to call a function
	a, b := somefun(2, "hi")
	fmt.Println(a, b)

	_, b = somefun(5, "Hello")
	fmt.Println(b)

	a, _ = somefun(2, "hi")
	fmt.Println(a)

	fmt.Println("Recursion funtion call")
	fmt.Println(count(5))

	fmt.Println(factorial(9))

}

func msg() {

	fmt.Println("just a msg!!!")

}

//parameters
func name(name string) {
	fmt.Println("my name is ", name)

}

//multi-para
func fullName(fname string, lname string) {

	fmt.Println("my full name is ", fname, " ", lname)

}

//return func
func sum(x int, y int) int {
	return x + y

}

func somefun(x int, y string) (result int, text string) {
	result = x + x
	text = y + " World!"
	return
}

//recursion function
func count(x int) int {

	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return count(x + 1)
}

func factorial(x float64) float64 {

	if x > 0 {
		return x * factorial(x-1)
	} else {
		return 1
	}
}
