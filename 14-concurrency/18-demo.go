package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(6 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	go func() {
		time.Sleep(4 * time.Second)
		data3 := <-ch3
		fmt.Println(data3)
	}()

	for i := 0; i < 3; i++ {
		select {
		case data1 := <-ch1:
			fmt.Println(data1)
		case data2 := <-ch2:
			fmt.Println(data2)
		case ch3 <- 300:
			fmt.Println("data sent to ch3")
			/*
				default:
						fmt.Println("No channel operation was successful")
			*/
		}
	}
}
