package main

import (
	"fmt"
	"sync"
)

//concurrent safe type
type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("Counter :", counter.count)

}

func fn(wg *sync.WaitGroup) {
	counter.Increment()
	wg.Done()
}
