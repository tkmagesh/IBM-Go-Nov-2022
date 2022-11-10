package main

import (
	"fmt"
	"time"
)

func main() {
	result := go add(100, 200)
	fmt.Println(result)
}

func add(x, y int) int {
	time.Sleep(3 * time.Second)
	return x + y
}
