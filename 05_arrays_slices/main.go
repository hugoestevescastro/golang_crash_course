package main

import (
	"fmt"
	"strconv"
)

func arrays() {
	//Array
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	fruitArr2 := [2]string{"apple", "orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(fruitArr2)
}

func slices() {
	fruitSlice := []string{"apple", "orange", "grape"}

	fmt.Print("fruitSlice: ")
	fmt.Println(fruitSlice)
	fmt.Println("length: " + strconv.Itoa(len(fruitSlice)))
	fmt.Print("fruitSlice[0:2]: ")
	fmt.Println(fruitSlice[0:2])
}

func main() {
	slices()
}
