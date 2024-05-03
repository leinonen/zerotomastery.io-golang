package main

import "fmt"

func main() {
	var myName = "Peter"
	fmt.Println("my name is,", myName, myName)

	var name string = "Charles"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("userName =", userName)

	// uninitialized variables are zero-valued
	var sum int
	fmt.Println("the sum is", sum)

	// comma, ok idioms
	part1, other := 1, 5
	fmt.Println("part1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName,", lessonName)
	fmt.Println("lessonType,", lessonType)

	// ignoring using underscore
	word1, word2, _ := "hello", "world", "!"
	fmt.Println("word1 is", word1, "word2 is", word2)
}
