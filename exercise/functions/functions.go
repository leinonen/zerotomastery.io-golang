//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hello", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func getMessage() string {
	return "This is a message"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addNumbers(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func getNumber() int {
	return 42
}

// * Write a function that returns any two numbers
func getTwoNumbers() (int, int) {
	return 1, 2
}

// * Add three numbers together using any combination of the existing functions.
//   - Print the result
func main() {
	//* Call every function at least once
	greet("Peter")
	fmt.Println(getMessage())

	a := getNumber()
	b, c := getTwoNumbers()

	result := addNumbers(a, b, c)
	fmt.Println("The result is", result)
}
