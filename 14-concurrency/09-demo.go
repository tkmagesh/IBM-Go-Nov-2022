package main

import (
	"fmt"
	"sync"
	"time"
)

//Communicate by sharing memory (NOT ADVISABLE)
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	result = x + y
	wg.Done()
}
