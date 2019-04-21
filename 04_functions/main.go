package main

import "fmt"

func main() {
	fmt.Println(getSum(1, 2))
	//fmt.Println(greeting("hugo"))
}

// func name(param param_type) return_type
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

// if num1 and num2 have the same type
func getSumShorthand(num1, num2 int) int {
	return num1 + num2
}
