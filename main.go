package main

import "fmt"

func main() {
	age := 35
	name := "Chris"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("new line \n")

	// Println 
	fmt.Println("Hello Ninjas!")
	fmt.Println("Goodbye Ninjas!")
	fmt.Println("My name is", name, "and my age is", age)

	// Printf Formatted strings %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)

	fmt.Printf("my age is %q and my name is %q \n", age, name)

	fmt.Printf("my age is of type %T \n", age)

	fmt.Printf("you scored %0.2f points! \n", 489.3039)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println(str)


}
