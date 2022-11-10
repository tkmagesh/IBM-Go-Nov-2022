package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // => scheduled this function for execution (go scheduler)
	f2()
	//time.Sleep(1 * time.Second) // blocking the execution of the main function so that the go scheduler can executed other goroutines scheduled
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
