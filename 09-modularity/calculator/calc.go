package calculator

import "fmt"

//private (to the package) variable
var opCount map[string]int

func init() {
	fmt.Println("Calculator initialized - [1]")
	opCount = make(map[string]int)
}

func init() {
	fmt.Println("Calculator initialized - [2]")
}

func OpCount() map[string]int {
	return opCount
}
