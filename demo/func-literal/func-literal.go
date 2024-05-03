package main

import "fmt"

type ComputeFunc = func(lhs, rhs int) int

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op ComputeFunc) int {
	fmt.Printf("Runnig a computation with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 =", compute(2, 2, add))

	fmt.Println("2 - 3 =", compute(2, 3, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		fmt.Printf("Multiplying %v * %v\n", lhs, rhs)
		return lhs * rhs
	}

	fmt.Println(compute(3, 3, mul))
}
