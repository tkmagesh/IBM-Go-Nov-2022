package calculator

import "fmt"

func init() {
	fmt.Println("Calculator[subtract.go] initialized")
}

func Subtract(x, y int) int {
	opCount["subtract"]++
	return x - y
}
