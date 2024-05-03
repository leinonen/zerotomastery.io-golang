//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func area(rect Rectangle) int {
	return rect.Height * rect.Width
}

func perimeter(rect Rectangle) int {
	return 2 * (rect.Height + rect.Width)
}

func printInfo(rect Rectangle) {
	fmt.Println("Area:", area(rect))
	fmt.Println("Perimeter:", perimeter(rect))
}

func main() {
	// rect := Rectangle{30, 40}
	rect := Rectangle{Width: 3, Height: 7}
	fmt.Println("Rectangle", rect)

	printInfo(rect)

	rect.Width *= 2
	rect.Height *= 2

	printInfo(rect)
}
