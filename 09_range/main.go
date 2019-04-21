package main

import "fmt"

func main() {
	ids := []int{23, 345, 65, 343, 233}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Bob2": "bob2@gmail.com", "Bobdel": "bobdel@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
