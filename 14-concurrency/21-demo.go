/* channel behavior */
package main

import "fmt"

func main() {
	ch := make(chan int, 1) /* buffered channel */
	fmt.Println("Before : len(ch) =", len(ch))
	ch <- 100
	fmt.Println("After : len(ch) =", len(ch))
	data := <-ch
	fmt.Println(data)
}
