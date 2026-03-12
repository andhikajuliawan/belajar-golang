package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	fmt.Println(eko)

	eko.Name = "Eko"
	eko.Address = "Jakarta"
	eko.Age = 25
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"budi", "Jakarta", 20}
	fmt.Println(budi)
}
