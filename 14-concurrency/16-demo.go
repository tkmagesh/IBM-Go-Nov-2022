package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for data := range ch {
		fmt.Println(data)
	}
}

func fn(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * 10
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}
