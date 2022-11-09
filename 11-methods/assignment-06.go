package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
	Write the apis for the following

	IndexOf => return the index of the given product
		ex:  returns the index of the given product

	Includes => return true if the given product is present in the collection else return false
		ex:  returns true if the given product is present in the collection

	Filter => return a new collection of products that satisfy the given condition
		use cases:
			1. filter all costly products (cost > 1000)
				OR
			2. filter all stationary products (category = "Stationary")
				OR
			3. filter all costly stationary products
			etc

	Any => return true if any of the product in the collection satifies the given criteria
		use cases:
			1. are there any costly product (cost > 1000)?
			OR
			2. are there any stationary product (category = "Stationary")?
			OR
			3. are there any costly stationary product?
			etc

	All => return true if all the products in the collection satifies the given criteria
		use cases:
			1. are all the products costly products (cost > 1000)?
			2. are all the products stationary products (category = "Stationary")?
			3. are all the products costly stationary products?
			etc

*/

type Products []Product

func (products Products) Format() string {
	var result string
	for _, product := range products {
		result += product.Format() + "\n"
	}
	return result
}

func (products Products) IndexOf(product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	for _, p := range products {
		if p == product {
			return true
		}
	}
	return false
}

/*
func (products Products) FilterCostlyProducts() Products {
	var result Products
	for _, product := range products {
		if product.Cost > 1000 {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) FilterStationaryProducts() Products {
	var result Products
	for _, product := range products {
		if product.Category == "Stationary" {
			result = append(result, product)
		}
	}
	return result
}
*/

func (products Products) Filter(predicate func(Product) bool) Products {
	var result Products
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	stove := Product{102, "Stove", 5000, 5, "Utencil"}
	//fmt.Println("Index of stove : ", IndexOf(products, stove))
	fmt.Println("Index of stove : ", products.IndexOf(stove))
	fmt.Println("Does products include a stove ?:", products.Includes(stove))

	fmt.Println("Costly Products")
	//costlyProducts := products.FilterCostlyProducts()
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts.Format())

	fmt.Println("Stationary Products")
	//stationaryProducts := products.FilterStationaryProducts()
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts.Format())

	fmt.Println("Costly Stationary Products")
	costlyStationaryProductPredicate := func(product Product) bool {
		return costlyProductPredicate(product) && stationaryProductPredicate(product)
	}
	costlyStationaryProducts := products.Filter(costlyStationaryProductPredicate)
	fmt.Println(costlyStationaryProducts.Format())

	fmt.Println("Are there ANY stationary products? :", products.Any(stationaryProductPredicate))
	fmt.Println("Are ALL the products stationary products? :", products.All(stationaryProductPredicate))
}
