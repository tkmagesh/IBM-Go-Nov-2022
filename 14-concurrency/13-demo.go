package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	/*
		1. invoke the add function
		2. receive the channel from the add function
		3. wait for the data to arrive through the channel
		4. receive the data and print it
	*/
	ch := add(100, 200)
	result := <-ch
	fmt.Println(result)
}

//producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
