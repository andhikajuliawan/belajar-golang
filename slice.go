package main

import "fmt"

func main() {
	names := [6]string{"Andhika", "Juliawan", "Dwi", "Putra", "Pratama", "Ganteng"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[3:]
	fmt.Println(slice2)

	var slice3 []string = names[:3]
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(days)

	daysSlice1 := days[5:] // sabtu, minggu
	fmt.Println(daysSlice1)

	daysSlice1[0] = "sabtu baru"
	daysSlice1[1] = "minggu baru"
	fmt.Println(daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "libur baru")
	daysSlice2[0] = "sabtu lama"
	// 	daysBaru := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu", "libur baru"}
	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Eko"
	// newSlice[2] = "Eko" // error, harusya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eko")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5} // [5]int{1,2,3,4,5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
