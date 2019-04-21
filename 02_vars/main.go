package main

import "fmt"

func main() {
	//usingVar()
	shorthand()
}

func usingVar() {
	var name = "Hugo"   // or 	var name string = "Hugo"
	var age int32 = 21  // if it is a specific type like int32, no inference is made
	const isCool = true // just like var and const on JS

	fmt.Println(name, age, isCool)

	// Get a variable type
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
}

func shorthand() {
	// Only usable within a function
	name := "Hugo"
	size := 1.3

	email, age := "hugo@gmail.com", 2.3 // assign new variable
	email, age = "hugo2@gmail.com", 3.3 // updated used variable

	fmt.Println(name)
	fmt.Println(size)
	fmt.Println(email)
	fmt.Println(age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", email)
	fmt.Printf("%T\n", age)
}
