/* function basics */
package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	msg := getGreetMessage("Suresh")
	fmt.Print(msg)
	fmt.Println(add(100, 200))

	//fmt.Println(divide(100, 7))

	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/

	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
	*/

	/*
		quotient := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
	*/

	quotient, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
}

func sayHi() {
	fmt.Println("Hi!")
}

func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

func getGreetMessage(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

//named return values
/* func getGreetMessage(userName string) (message string) {
	message = fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	return
} */

//multiple parameters
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/* multiple return results */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/*
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
