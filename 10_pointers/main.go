package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b) // value
	fmt.Println(b)  // address

	// change val with pointer
	*b = 10
}
