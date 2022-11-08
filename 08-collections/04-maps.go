package main

import "fmt"

func main() {
	//var productRanks map[string]int
	/*
		var productRanks map[string]int = make(map[string]int)
		productRanks["pen"] = 4
		fmt.Println(productRanks)
	*/

	//var productRanks map[string]int = map[string]int{} // == make(map[string]int)
	//var productRanks  = map[string]int{}
	//productRanks := map[string]int{}

	//productRanks := map[string]int{"Pen": 4, "Pencil": 1, "Marker": 3}
	productRanks := map[string]int{
		"Pen":    4,
		"Pencil": 1,
		"Marker": 3,
	}
	fmt.Println(productRanks)

	//adding a new item
	productRanks["Notebook"] = 2
	fmt.Println(productRanks)

	fmt.Println("Count : ", len(productRanks))

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Check if a key exists")
	keyToCheck := "Notebook"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("%q does not exist\n", keyToCheck)
	}

	fmt.Println("Deleting....")
	keyToDelete := "Charger"
	delete(productRanks, keyToDelete)
	fmt.Printf("After deleting %q, productRanks = %v\n", keyToDelete, productRanks)
}
