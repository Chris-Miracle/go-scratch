package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hello Charles, Hoow are you?!"

	// fmt.Println(strings.Contains(greeting, "Charles"))
	// fmt.Println(strings.ReplaceAll(greeting, "are", "is"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "?"))
	// fmt.Println(strings.Split(greeting, " "))

	// fmt.Println("Original string =", greeting)

	ages := []int{48, 94, 96, 39, 30, 18, 33, 30}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 96)

	fmt.Println(index)

	names := []string{"hey", "life", "isnt", "funny"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "life"))

}
