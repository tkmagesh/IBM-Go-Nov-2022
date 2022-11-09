package main

import "fmt"

func main() {
	//var x interface{} //before go 1.18
	var x any //from go 1.18
	x = 10
	x = "This is a string"
	//x = true
	//x = 1000.98
	//x = struct{}{}
	fmt.Println(x)

	//x = 200
	//y := x + 300
	fmt.Println("Type Asserting (using if)")
	if val, ok := x.(int); ok {
		y := val + 300
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	fmt.Println("Type Asserting (using switch)")
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 300 =", val+300)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, x + 200.987 =", val+200.987)
	case struct{}:
		fmt.Println("x is an empty struct")
	default:
		fmt.Println("unknown type")
	}

}
