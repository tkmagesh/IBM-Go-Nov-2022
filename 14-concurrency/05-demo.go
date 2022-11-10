package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1) //incrementing the counter by 1
		go f1(wg) //schedule the f1 in the scheduler (to be executed in future)
	}
	f2()
	wg.Wait() //wait for the counter to become 0 (blocking)
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second) //simulating a time consuming operation
	fmt.Println("f1 completed")
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
