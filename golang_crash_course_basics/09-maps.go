package main

import "fmt"

func main() {
	// maps are unordered like in most other languages
	// declare with values:
	emails := map[string]string{"Bob": "bob@gmail.com", "Joe": "joe@gmail.com"}
	fmt.Println(emails["Bob"])

	// declare multiline
	emails2 := map[string]string{
		"Bob": "bob@gmail.com",
		"Joe": "joe@gmail.com", // must have trailing comma
	}
	fmt.Println(emails2)

	// create empty mape: make(map[key-type]val-type)
	cart := make(map[string]int)
	cart["milk"] = 3
	cart["cheese"] = 2
	fmt.Println(cart) // map[cheese:2 milk:3]

	// create new entry and add 3 to its default value (0)
	cart["someNewItem"] += 3
	fmt.Println(cart) // map[cheese:2 milk:3 someNewItem:3]

	// accessing values
	fmt.Println(cart["chees"])

	// check if a value exists and assign it
	// the second value is a boolean that is true if the item was found in the map
	milkInStore, foundItem := cart["milk"]
	if foundItem {
		fmt.Println(milkInStore)
	}

	// delete items
	delete(cart, "cheese")

	fmt.Println(len(cart)) // length like usual with len

	// iterate over map
	for key := range emails {
		fmt.Println(key) // Bob Joe
	}

	for key, value := range emails {
		fmt.Println(key + " " + value) // Bob bob@gmail.com Joe joe@gmail.com
	}
}
