package main

import "fmt"

func main() {
	//var nos []int
	//var nos []int = []int{3, 1, 4, 2, 5}
	//var nos = []int{3, 1, 4, 2, 5}
	//nos := []int{3, 1, 4, 2, 5}
	//fmt.Println(nos)

	//var nos []int
	/*
		var nos []int = make([]int, 5)
		fmt.Println(nos)
		nos[0] = 10
		nos[1] = 20
		nos[2] = 30
		nos[3] = 40
		nos[4] = 50
		fmt.Println(nos)
	*/
	//exhausted the memory allocated & initialized
	var nos []int
	nos = append(nos, 60)
	fmt.Println(nos)
	//nos = append(nos, 70, 80)
	vals := []int{70, 80, 90}
	nos = append(nos, vals...)
	fmt.Println(nos)

	values := []int{3, 1, 4, 2, 5}
	fmt.Println("Before sort : ", values)
	sort(values)
	fmt.Println("After sort : ", values)

	//slicing
	/*
		[lo:hi] => from lo to hi-1
		[lo:] => from lo to end
		[:hi] => from 0 to hi-1
	*/
	x := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(x[2:5])
	fmt.Println(x[2:])
	fmt.Println(x[:5])
}

func sort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
