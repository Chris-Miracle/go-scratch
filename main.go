package main

import (
	"fmt"
)


func main() {
	
	menu := map[string]float64{
		"soup" : 4.5,
		"pie" : 7.99,
		"salad" : 6.99,
		"pudding" : 3.55,
	}

	fmt.Println(menu)

	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		1234567890 : "mario",
		9876543210 : "luigi",
		1111111111 : "yoshi",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1234567890])

	phonebook[1234567890] = "mario2"

	fmt.Println(phonebook)

	phonebook[1111111111] = "jake"

	fmt.Println(phonebook)
}
