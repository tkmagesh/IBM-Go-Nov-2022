package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		break
	}
}

func fn(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * 10
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}
