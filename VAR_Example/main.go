package main

import "fmt"

func main() {
	var x float64 = 20.0

	y := 42
	z := "hello"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)
}
