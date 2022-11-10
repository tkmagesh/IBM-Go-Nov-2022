package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Share memory by communicating
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch //RECEIVE the result from the channel
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result // SENDING the result to the channel
	wg.Done()
}
