package main

import (
	"fmt"
)

func main() {
	go f1() // => scheduled this function for execution (go scheduler)
	f2()
	//time.Sleep(1 * time.Second) // blocking the execution of the main function so that the go scheduler can executed other goroutines scheduled
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
