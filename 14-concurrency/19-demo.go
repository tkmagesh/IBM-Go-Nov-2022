package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFib()
	for fibNo := range ch {
		fmt.Println(fibNo)
	}
	fmt.Println("Done")
}

func genFib() <-chan int {
	ch := make(chan int)
	//timeoutCh := timeout(5 * time.Second)
	timeoutCh := time.After(5 * time.Second)
	go func() {
		x, y := 1, 1
	LOOP:
		for {
			select {
			case <-timeoutCh:
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

func timeout(t time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(t)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
