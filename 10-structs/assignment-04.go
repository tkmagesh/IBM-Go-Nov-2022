/*
	Create a struct Product with Id, Name, Cost, Category fields
	Create the following functions
		PrintProduct -> prints all the field values of the given product
		ApplyDiscount -> Apply the given discount(%) on the given product (Cost)
*/

package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
}

func PrintProduct(product Product) {
	fmt.Printf("Id = %d, Name = %q, Cost = %v, Category = %q\n", product.Id, product.Name, product.Cost, product.Category)
}

func ApplyDiscount(product *Product, discountPercent float32) {
	product.Cost = product.Cost * ((100 - discountPercent) / 100)
}

func main() {
	pen := Product{Id: 100, Name: "Pen", Cost: 10, Category: "Stationary"}
	PrintProduct(pen)
	ApplyDiscount(&pen, 10)
	PrintProduct(pen)
}
