package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 30, 79}

	var ages = [3]int{20, 30, 79};

	names := [4]string{"chris", "james", "jake", "hillary"}
	names[1] = "chuks"

	fmt.Println(ages, len(ages))

	fmt.Println(names, len(names))

	// Arrsy slices
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 69)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo:= names[1:]
	rangeThree:= names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Jill")

	fmt.Println(rangeOne)

}
