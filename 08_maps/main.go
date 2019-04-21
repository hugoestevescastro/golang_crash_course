package main

import "fmt"

func main() {
	// Map is a key-value structure

	// default map  map[key-type]value-type
	emails := make(map[string]string)

	// Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Bob2"] = "bob2@gmail.com"
	emails["Bob_del"] = "bobdel@gmail.com"

	// shorthand to create and assign
	// emails := map[string]string{"Bob": "bob@gmail.com", "Bob2": "bob2@gmail.com", "Bobdel": "bobdel@gmail.com"}
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])
	fmt.Println("after delete")

	// Delete values
	delete(emails, "Bob_del")
	fmt.Println(emails)
	fmt.Println(len(emails))

}
