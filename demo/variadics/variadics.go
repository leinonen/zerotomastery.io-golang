package main

import "fmt"

func sum(nums ...int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	answer := sum(1, 2, 3, 4)
	fmt.Println(answer)
}
