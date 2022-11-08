/* higher order functions - return a function as a return value from another function */
package main

import "fmt"

func main() {
	fn := getFn()
	fn()
}

func getFn() func() {
	fn := func() {
		fmt.Println("fn invoked")
	}
	return fn
}
