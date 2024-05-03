package main

import "fmt"

// returns the double of the input
func double(x int) int {
	return x * 2
}

// returns the sum of the two inputs
func add(lhs, rhs int) int {
	return lhs + rhs
}

// prints a greeting to the terminal
func greet() {
	fmt.Println("Hello!")
}

func main() {
	myAge := 40
	fmt.Println("My age is", myAge)

	ageDoubled := double(myAge)
	fmt.Println("My age doubled is", ageDoubled)

	doubledAgeWith10 := add(ageDoubled, 10)
	fmt.Println("My age doubled with 10 is", doubledAgeWith10)

	greet()
}
