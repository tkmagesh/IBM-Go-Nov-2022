/*
	Modify the below program to keep generating the numbers until the user hits ENTER key
	DO NOT read the ENTER key in the genFib() function
*/
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
