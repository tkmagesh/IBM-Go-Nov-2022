package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
}

func (product Product) PrintProduct() {
	fmt.Printf("Id = %d, Name = %q, Cost = %v, Category = %q\n", product.Id, product.Name, product.Cost, product.Category)
}

func (product *Product) ApplyDiscount(discountPercent float32) {
	product.Cost = product.Cost * ((100 - discountPercent) / 100)
}

type PerishableProduct struct {
	Expiry string
	Product
}

func (pp PerishableProduct) PrintProduct() {
	fmt.Printf("Id = %d, Name = %q, Cost = %v, Category = %q, Expiry = %q\n", pp.Id, pp.Name, pp.Cost, pp.Category, pp.Expiry)
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

	/*
		PrintProduct(apple.Product)
		fmt.Println("Applying 10% discount on apples")
		ApplyDiscount(&apple.Product, 10)
		PrintProduct(apple.Product)
	*/
	apple.Product.PrintProduct()
	apple.PrintProduct()
	apple.ApplyDiscount(10)
	apple.PrintProduct()
}
