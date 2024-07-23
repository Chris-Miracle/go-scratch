package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string){
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string){
	fmt.Printf("Good bye %v \n", n)
}

func cycleNames(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main() {
	// sayGreeting("mario");
	// sayGreeting("jake");
	// sayGreeting("ola");

	cycleNames([]string{"nike", "addidas", "puma"}, sayGreeting)
	cycleNames([]string{"nike", "addidas", "puma"}, sayBye)

	a1 := circleArea(92.22)
	a2 := circleArea(15)

	fmt.Println(a1, a2)

	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2 )

}
