package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hit ENTER to start....")
	fmt.Scanln()
	wg := &sync.WaitGroup{}
	var goroutineCount = flag.Int("count", 1, "Number of goroutines to execute")
	flag.Parse()
	rand.Seed(7)
	for i := 1; i <= *goroutineCount; i++ {
		wg.Add(1)
		go f1(wg, i, time.Duration(rand.Intn(10))*time.Second)
	}
	f2()
	wg.Wait() //wait for the counter to become 0 (blocking)
	fmt.Println("Done")
	fmt.Scanln()
}

func f1(wg *sync.WaitGroup, id int, delay time.Duration) {
	fmt.Printf("f1[%d] started\n", id)
	time.Sleep(delay) //simulating a time consuming operation
	fmt.Printf("f1[%d] completed\n", id)
	wg.Done() //decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
