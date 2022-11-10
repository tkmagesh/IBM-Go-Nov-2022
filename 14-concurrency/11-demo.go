/* channel behavior */
package main

import "fmt"

func main() {
	/*
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
		fmt.Println("Done")
	*/

	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)

}
