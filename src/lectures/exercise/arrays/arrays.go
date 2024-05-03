//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func printCart(cart [4]Product) {
	var totalCost, numItems int
	for i := 0; i < len(cart); i++ {
		current := cart[i]
		if current.name == "" {
			continue
		}
		totalCost += int(current.price)
		numItems += 1
	}
	fmt.Println("last item:", cart[numItems-1])
	fmt.Println("total items:", numItems)
	fmt.Println("total cost:", totalCost)
}

func main() {
	shoppingList := [4]Product{
		{name: "Banana", price: 100},
		{"Meat", 200},
		{name: "Salad", price: 300},
	}
	//shoppingList[0] = Product{name: "1", price: 100}
	//shoppingList[1] = Product{name: "2", price: 200}
	//shoppingList[2] = Product{name: "3", price: 300}

	printCart(shoppingList)

	fmt.Println("Adding a new item to the cart...")
	shoppingList[3] = Product{"Ketchup", 400}
	printCart(shoppingList)
}
