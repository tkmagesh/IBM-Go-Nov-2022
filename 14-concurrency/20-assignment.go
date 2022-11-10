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
	stopCh := make(chan struct{})
	ch := genFib(stopCh)
	go func() {
		fmt.Println("Hit ENTER to stop.....")
		fmt.Scanln()
		//stopCh <- struct{}{}
		close(stopCh)
	}()

	for fibNo := range ch {
		fmt.Println(fibNo)
	}
	fmt.Println("Done")
}

func genFib(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 1, 1
	LOOP:
		for {
			select {
			case <-stopCh:
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
