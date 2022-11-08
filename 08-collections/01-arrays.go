package main

import "fmt"

func main() {
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iteration using for (indexer)")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iteration using for (range)")
	for _, val := range nos {
		//fmt.Printf("nos[%d] = %d\n", idx, val)
		fmt.Println(val)
	}

	x := [3]int{10, 20, 30}
	y := x
	x[0] = 100
	fmt.Println(x, y)

	sort(&nos)
	fmt.Println(nos)
}

func sort(nos *[5]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if nos[i] > nos[j] {
				nos[i], nos[j] = nos[j], nos[i]
			}
		}
	}
}
