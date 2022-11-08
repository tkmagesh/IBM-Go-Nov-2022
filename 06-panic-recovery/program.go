package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("something went wrong")
			fmt.Println(err)
			return
		}
		fmt.Println("Thank you!")
	}()
	var n1, n2 int
	for {
		fmt.Scanln(&n1, &n2)
		quotient, remainder, err := divideClient(n1, n2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(quotient, remainder)
		break
	}
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("unable to divide")
		}
		return
	}()
	quotient, remainder = divide(x, y)
	return
}

//3rd party library
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic("divisor cannot be zero")
	}
	quotient, remainder = x/y, x%y
	return
}
