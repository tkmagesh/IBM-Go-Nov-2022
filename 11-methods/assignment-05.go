/*
	Take the solution for assignment-04
	Convert the "PrintProduct" & "ApplyDiscount" functions as methods of "Product" type

*/

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

func main() {
	pen := Product{Id: 100, Name: "Pen", Cost: 10, Category: "Stationary"}
	pen.PrintProduct()
	//(&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10) // a method with a pointer receiver is invoked with a value
	pen.PrintProduct()

	penPtr := &Product{Id: 100, Name: "Pen", Cost: 10, Category: "Stationary"}
	//fmt.Println(pen.Name)
	//fmt.Println((*penPtr).Name)
	//fmt.Println(penPtr.Name)
	penPtr.PrintProduct() // a method with a value receiver is invoked with a pointer
}
