package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Expiry string
	//Name   string
	Product
	//Dummy
}

func PrintProduct(product Product) {
	fmt.Printf("Id = %d, Name = %q, Cost = %v, Category = %q\n", product.Id, product.Name, product.Cost, product.Category)
}

func ApplyDiscount(product *Product, discountPercent float32) {
	product.Cost = product.Cost * ((100 - discountPercent) / 100)
}

func main() {
	//apple := PerishableProduct{Product{100, "Apple", 10, "Food"}, "1 Week"}
	apple := PerishableProduct{
		Product: Product{100, "Apple", 10, "Food"},
		Expiry:  "1 Week",
	}
	fmt.Printf("%#v\n", apple)
	fmt.Println(apple.Product.Name)
	fmt.Println(apple.Name)

	PrintProduct(apple.Product)
	fmt.Println("Applying 10% discount on apples")
	ApplyDiscount(&apple.Product, 10)
	PrintProduct(apple.Product)
}
