package main

import (
	"fmt"
)

func main() {
	// x := 0

	// for x < 5 {
	// 	fmt.Println("value of x is:", x)

	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of x is:", i)
	// }

	names := []string{"Hey", "life", "is", "simple"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("value of names right now is ", names[i])
	// }

	// for index, value := range names {
	// 	fmt.Println(index, value)
	// }

	for _, value := range names {
		fmt.Println(value)
		value = "jake"
	}

	fmt.Println(names)

}
