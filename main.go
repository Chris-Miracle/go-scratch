package main

import (
	"fmt"
)

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 45")
	}

	names := []string{"jalee", "jill", "biden", "trump"}

	for index, value := range names {
		if(index == 1){
			fmt.Println("Coontinuing at pos", index)
			continue
		}

		if(index > 1){
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("thee value at pos %v is %v \n", index, value)
	}

}
