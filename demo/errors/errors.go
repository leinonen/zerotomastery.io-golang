package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, fmt.Errorf("no element at index %v", index)
	} else {
		return s.values[index], nil
	}
}

func main() {

	stuff := Stuff{}
	val, err := stuff.Get(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Got val", val)

	stuff.values = append(stuff.values, 1, 2, 3)

	val2, err := stuff.Get(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Got val", val2)
}
