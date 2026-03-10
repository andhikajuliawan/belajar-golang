package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Andhika",
		"address": "surabaya",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Go-lang"
	book["author"] = "andhika juliawan"
	book["wrong"] = "Ups"

	fmt.Println(len(book))
	fmt.Println(book)
	delete(book, "wrong")
	fmt.Println(book)
}
