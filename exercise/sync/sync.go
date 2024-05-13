//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"unicode"
)

type LetterCount struct {
	count int
	sync.Mutex
}

func analyze(str string, wg *sync.WaitGroup, counter *LetterCount) {
	numLetters := 0
	for _, char := range str {
		if unicode.IsLetter(char) {
			numLetters += 1
		}
	}
	counter.Lock()
	defer counter.Unlock()
	defer wg.Done()
	counter.count += numLetters
}

func main() {
	var wg sync.WaitGroup
	letterCount := LetterCount{}
	rd := bufio.NewReader(os.Stdin)

	for {
		str, err := rd.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		} else {
			wg.Add(1)
			go analyze(strings.TrimSpace(str), &wg, &letterCount)
		}
	}
	wg.Wait()

	letterCount.Lock()
	count := letterCount.count
	letterCount.Unlock()

	fmt.Println("letter count =", count)
}
