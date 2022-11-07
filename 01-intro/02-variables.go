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

	//type conversion (NOTE : Use the type name like a function)
	var i int = 100
	var f float32 = float32(i)
	fmt.Println(f)

	//complex numbers
	//var c1 complex64 = 4 + 5i
	//var c1 = 4 + 5i
	c1 := 4 + 5i
	fmt.Println(c1)
	var c2 = 7 + 3i
	c1c2 := c1 + c2
	fmt.Println(c1c2)

	//accessing the real & imaginary parts of the complex number
	fmt.Println("real : ", real(c1c2))
	fmt.Println("imaginary : ", imag(c1c2))

	//constants
	// convention - ALL UPPER CASE
	//const pi float32 = 3.14
	const pi = 3.14
	//fmt.Println("pi :", pi)

	//iota

	/*
		const red int = 0
		const green int = 1
		const blue int = 2
	*/
	/*
		const (
			red   int = 0
			green int = 1
			blue  int = 2
		)
	*/
	/*
		const (
			red   int = iota
			green int = iota
			blue  int = iota
		)
	*/

	/*
		const (
			red int = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 2
			green
			blue
		)
	*/

	/*
		const (
			red = iota * 2
			green
			blue
		)
	*/

	const (
		red = iota * 2
		green
		_
		blue
	)
	fmt.Printf("red = %d, green = %d, blue = %d\n", red, green, blue)

	//IOTA Usage:
	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}
