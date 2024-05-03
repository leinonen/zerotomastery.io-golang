package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("quiz1 scored higher than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz1 scored lower than quiz2")
	} else {
		fmt.Println("quiz1 and quiz2 scored the same")
	}

	if avg := average(quiz1, quiz2, quiz3); avg > 7 {
		fmt.Println("You passed the test. Average:", avg)
	} else {
		fmt.Println("You failed the test. Average:", avg)
	}
}
