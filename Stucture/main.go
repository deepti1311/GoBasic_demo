package main

import "fmt"

//syntex
//type struct_name struct {
//	member1 datatype;
//	member2 datatype;
//	member3 datatype;
//	...
//}

func main() {

	var per1 person

	per1.name = "deepti"
	per1.age = 22

	fmt.Println("Name: ", per1.name)
	fmt.Println("Age", per1.age)

	//call a Structure using function

	print(per1)

}

type person struct {
	name string
	age  int
}

func print(per person) {

	fmt.Println("name: ", per.name)
	fmt.Println("age: ", per.age)
}
