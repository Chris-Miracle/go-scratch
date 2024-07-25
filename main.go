package main

import (
	"fmt"
)

func updateName(name string) string {
	name = "Nairobi"
	return name
}

func updateMenu(menu map[string]float64) {
	menu["Burger"] = 1000.09
}


func main() {
	
	// group A types -> string, int, float, bool, arrays, structs
	name := "Ikeja"

	name = updateName(name)

	fmt.Println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"Burger": 10.0,
		"Pizza": 20.0,
		"Fries": 1.50,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
