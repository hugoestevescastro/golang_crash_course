package main

import "fmt"

// Define person struct
type Person struct {
	// alternative
	//firstName string
	//lastName  string
	//city      string
	//gender    string
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (person Person) greet() string {
	return person.firstName + " greets you!"
}

// hasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried method (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "M" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Hugo", lastName: "Castro", city: "Melgaço", gender: "M", age: 21}
	// Alternative, by order
	person2 := Person{"Sara", "Rodrigues", "Melgaço", "F", 21}

	fmt.Println(person1)
	fmt.Println(person2)

	// get field
	fmt.Println(person1.firstName)
	person1.firstName = "changed"
	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.age)
	person1.getMarried("Williams")
	person2.getMarried("Williams")
	fmt.Println(person1)
	fmt.Println(person2)
}
