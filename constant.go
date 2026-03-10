package main

import "fmt"

func main() {
	const (
		firstName string = "Andhika"
		lastName         = "Juliawan"
	)

	fmt.Println(firstName, lastName)
	fmt.Println(lastName)

	// Error
	// firstName = "Andhika1"
	// lastName = "Juliawan1"
}
