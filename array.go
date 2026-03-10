package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Andhika"
	names[1] = "Juliawan"
	names[2] = "Dwi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90, 
		80, 
		95, 
		100, 
		110,
	}
	fmt.Println(values)
	values[0] = 100
	fmt.Println(values[0])	
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
}
