/* closures */

package main

import "fmt"

var counter int

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	counter = 1000
	fmt.Println(increment())
	fmt.Println(increment())
}

func increment() int {
	counter++
	return counter
}
