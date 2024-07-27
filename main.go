package main

import "fmt"

func main() {
	myBill := newBill("Chris's Bill")

	myBill.addItem("Donut", 50)
	myBill.addItem("Cake", 10)
	myBill.addItem("Chocolatee", 13)

	myBill.updateTip(10)

	fmt.Println(myBill.format())
}
