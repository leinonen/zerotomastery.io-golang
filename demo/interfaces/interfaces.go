package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("* Cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("* Chop salad")
	fmt.Println("* Add dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for _, dish := range dishes {
		fmt.Printf("Dish: %v\n", dish)
		dish.PrepareDish()
	}
}

func main() {
	dishes := []Preparer{
		Chicken("Chicken Wings"),
		Salad("Iceberg salad"),
	}
	prepareDishes(dishes)
}
