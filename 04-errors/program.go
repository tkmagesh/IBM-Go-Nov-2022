package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be zero")

func main() {
	quotient, remainder, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Please donot attempt to divide by zero", err)
		return
	}
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	fmt.Printf("Quotient = %d, Remainder = %d\n", quotient, remainder)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := fmt.Errorf("invalid divisor. divisor = %d", y)
		return 0, 0, err
	}
	quotient, remainder := x/y, x%y
	fmt.Println("Returing results")
	return quotient, remainder, nil
}
*/

/*
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = fmt.Errorf("invalid divisor. divisor = %d", y)
		return
	}
	quotient, remainder = x/y, x%y
	fmt.Println("Returing results")
	return
}
*/

/*
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient, remainder = x/y, x%y
	fmt.Println("Returing results")
	return
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	fmt.Println("Returing results")
	return
}
