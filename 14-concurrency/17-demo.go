package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		data1 := <-ch1
		fmt.Println(data1)
		wg.Done()
	}()

	go func() {
		data2 := <-ch2
		fmt.Println(data2)
		wg.Done()
	}()
	wg.Wait()
}
