package main

import "fmt"

//var s string = "Hello World!"
//var s = "Hello World!"

//The following syntax (:=) is application ONLY in function scope variables and NOT int package scope variables
//s := "Hello World!"

//NOTE: unused variables at package scope are valid
var no int

func main() {
	/*
		var s string
		s = "Hello World!"
	*/

	/*
		var s string = "Hello World!"
	*/

	//type inference
	/*
		var s = "Hello World!"
	*/

	s := "Hello World!"
	fmt.Println(s)

	//NOTE : Cannot have a "function scoped" variables unused
	/*
		var no int
		no = 100
		fmt.Println(no)
	*/

	//multiple variable declarations & initialization
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of 100 & 200 is :"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 & 200 is :"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of 100 & 200 is :"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of 100 & 200 is :"
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of 100 & 200 is :"
		var result int = x + y
		fmt.Println(str, result)
	*/

	//type inference
	/*
		var x, y = 100, 200
		var str = "Sum of 100 & 200 is :"
		var result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Sum of 100 & 200 is :"
			result = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of 100 & 200 is :"
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Sum of 100 & 200 is :"
	result := x + y
	fmt.Println(str, result)
}
