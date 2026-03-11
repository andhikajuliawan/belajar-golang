package main

import "fmt"

func getFullName() (string, string) {
	return "Andhika", "Juliawan"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
