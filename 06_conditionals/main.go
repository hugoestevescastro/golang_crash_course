package main

import "fmt"

func main() {
	x := 15
	y := 10

	// if else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d \n", x, y)
	} else {
		fmt.Printf("%d is less than %d \n", y, x)
	}

	// else if
	color := "blue"

	if color == "red" {
		fmt.Println("It is red")
	} else if color == "blue" {
		fmt.Println("It is blue")
	} else {
		fmt.Println("Color not blue or red")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("is red")
	case "blue":
		fmt.Println("is blue")
	default:
		fmt.Println("not blue or red")
	}
}
