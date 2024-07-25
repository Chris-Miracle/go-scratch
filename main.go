package main

import (
	"fmt"
)

func updateName(name *string) {
	*name = "Nairobi"
}

func main() {
	
	// group A types -> string, int, float, bool, arrays, structs
	name := "Ikeja"


	m := &name

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}
