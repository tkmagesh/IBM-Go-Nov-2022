/* higher order functions - pass a function as an argument to another function */
package main

import "fmt"

func main() {
	exec(f1)
	exec(f2)
	exec(func() {
		fmt.Println("anonymous fn invoked")
	})
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func exec(fn func()) {
	fn()
}
