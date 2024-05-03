//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Inactive = false
)

type SecurityTag bool
type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	/*for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}*/
	for i := range items {
		deactivate(&items[i].tag)
	}
}

func main() {
	//  - Create at least 4 items, all with active security items
	//  - Store them in a slice or array
	items := []Item{
		{"first", Active},
		{"second", Active},
		{"third", Active},
		{"fourth", Active},
	}
	fmt.Println("initial", items)

	deactivate(&items[1].tag)
	fmt.Println("deactivate second", items)

	checkout(items)
	fmt.Println("after checkout", items)
}
