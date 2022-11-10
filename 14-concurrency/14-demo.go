package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fn(ch chan int) {
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	ch <- 50
}
