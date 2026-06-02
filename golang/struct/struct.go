package main

import "fmt"

func main() {
	person := Person{
		"Nguyễn",
		"Huy",
		23,
		"Gò Vấp",
		"0123456789",
		"test@gmail.com",
	}

	fmt.Println(person.fullName())
}
