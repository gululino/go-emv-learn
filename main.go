package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"
)

func main() {

	// variables explicit type
	var name string = "Manzi"
	var age int = 90

	fmt.Println(" Name is " + name)
	fmt.Println(" Age  is", age)

	fmt.Println()

	// multiple variable at once
	var c, d int = 10, 20
	a, b := "hello", true

	fmt.Println(" C Value " , c)
	fmt.Println(" D Value ", d)
	fmt.Println(" A Value " , a)
	fmt.Println(" B Value ", b)

	fmt.Println()

	// block declaration
	var ( 
		host string = "localhost"
		port int 	= 8080
		tls  bool	= true
	)

	fmt.Println(" Host is " + host)
	fmt.Println(" Port is ", port)
	fmt.Println(" tls is ",  tls)

	fmt.Println()

	// constants
	const MaxRetries = 3
	const AppName = "PayConex"

	// Typed constant
	const Timeout time.Duration = 30 * time.Second

	// Go implicit 
	var x int = 42
	var z string = strconv.Itoa(x)  //int -> string

	fmt.Println(" Z Value " , z)
	//fmt.Println(" Y Value ", y)


	// okay lets put all these together

	//----- 1. Basic declaration & Assignment

	var firstName string = "Manzi"
	lastname := "MJ"
	var salary float64 = 1000.00

	// Concatenation 

	fullName := firstName + " " + lastname
	fmt.Println("Full Name is :", fullName)


	// -----2. Updating / reassigning 
	age = age + 1		// birthday 
	salary *= 1.10		// 10% raise
	fmt.Printf("Updated age is : age =%d, salary is =%.2f\n", age, salary)


	//-----4 Type Conversation
	var score int = 95
	scoreFloat := float64(score)
	scoreStr := strconv.Itoa(score)

	fmt.Printf("int: %d | float: %.1f | string: \"%s\" and ", score, scoreFloat, scoreStr)

	//----5. String Manipulation ----
	role := " Certification Analyst II "
	role = strings.TrimSpace(role)
	role = strings.ToUpper(role)
	fmt.Println("Role", role)

	fmt.Println()


	words := strings.Split(fullName, " ")
	fmt.Println("Name parts: ", words)			// [Manzi Joseph]
	fmt.Println("First word: ", words[0])

	fmt.Println()

	// ----6. Zero Values at work 
	var counter int
	var label string
	var active bool

	counter++			//0 --> without initializing
	label = "EMV Cert"
	active = !active	// false --> true

	fmt.Printf("counter=%d, label=%s, active=%v\n", counter,label, active)

	// -----7. Constants vs Variable 
	attempts := 0

	for attempts < MaxRetries {
		attempts++
		fmt.Printf("Attempts %d of %d\n", attempts, MaxRetries)
	}

}
