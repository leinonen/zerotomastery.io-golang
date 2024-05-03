//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type RuneFunc func(r rune) bool

func countBy(line string, check RuneFunc) int {
	var count int
	for _, char := range line {
		if check(char) {
			count += 1
		}
	}
	return count
}

func lineIterator(lines []string, callback func(line string)) {
	for _, line := range lines {
		callback(line)
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	letters := 0
	digits := 0
	spaces := 0
	puncts := 0

	lineFunc := func(line string) {
		letters += countBy(line, unicode.IsLetter)
		digits += countBy(line, unicode.IsDigit)
		spaces += countBy(line, unicode.IsSpace)
		puncts += countBy(line, unicode.IsPunct)
	}
	lineIterator(lines, lineFunc)

	// fmt.Printf("'%v' contains the following:\n", line)
	fmt.Printf("%v letters\n", letters)
	fmt.Printf("%v digits\n", digits)
	fmt.Printf("%v spaces\n", spaces)
	fmt.Printf("%v punctuations\n\n", puncts)
}
