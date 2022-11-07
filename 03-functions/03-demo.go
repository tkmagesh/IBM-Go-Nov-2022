/* anonymous functions */
/*
	1. it should have a name
	2. it should be immediately invoked
*/
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	message := func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	}("Suresh")
	fmt.Print(message)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("Result =", result)

	quotient, remainder := func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Println(quotient, remainder)
}
