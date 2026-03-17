package main

import (
	"fmt"
	"math"
)

func main() {
	var age int // variable declaration
	fmt.Println("Enter your age: ", age)
	age = 25 // variable assignment
	fmt.Println("My age after assignment is:", age)
	age = 54 // variable reassignment
	fmt.Println("My age after reassignment is:", age)

	var newAge int = 67
	fmt.Println("my new age : ", newAge)

	// multiple varaibles declaration
	var price, quantity int = 5000, 100 // declaring multiple vairables
	fmt.Println("price is ", price, "quantity is ", quantity)

	var (
		anotherName = "Zeinab"
		herAge      = 38
		height      int
	)

	fmt.Println("The Queens name is : ", anotherName)
	fmt.Println("My age is: ", herAge)
	fmt.Println("my height is ", height)

	count := 10
	fmt.Println("Count =", count)

	//lets do more variable declaration
	a, b := 20, 30 //declare variables a and b
	fmt.Println("a is ", a, "b is ", b)
	b, c := 50, 69 // b is already declared but c is new
	fmt.Println("b is ", b, "c is ", c)
	b, c = 90, 100 // assign new values to already assigned variables
	fmt.Println("changed b is ", b, "c is ", c)

	fmt.Println()

	x, y := 145.4, 67.9
	k := math.Min(x, y)
	fmt.Println("Minimum value is ", k)
}
