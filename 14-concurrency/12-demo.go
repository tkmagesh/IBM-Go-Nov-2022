package main

import (
	"fmt"
	"time"
)

func main() {
	//Share memory by communicating
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch //RECEIVE the result from the channel
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result // SENDING the result to the channel
}
