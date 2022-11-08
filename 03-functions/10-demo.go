/* closures */

package main

import "fmt"

func main() {
	increment := getIncrement()
	fmt.Println(increment())
	fmt.Println(increment())
	//counter = 1000
	fmt.Println(increment())
	fmt.Println(increment())
}

func getIncrement() func() int {
	var counter int
	return func() int {
		counter++
		return counter
	}
}
