/* higher order functions - composability */

package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	//Log
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	//Profile
	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)

		profileSubtract := getProfileOperation(subtract)
		profileSubtract(100, 200)
	*/

	//Log & Profile
	/*
		logAdd := getLogOperation(add)
		profileLogAdd := getProfileOperation(logAdd)
		profileLogAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		profileLogSubtract := getProfileOperation(logSubtract)
		profileLogSubtract(100, 200)
	*/

	profileLogAdd := getProfileOperation(getLogOperation(add))
	profileLogAdd(100, 200)

	profileLogSubtract := getProfileOperation(getLogOperation(subtract))
	profileLogSubtract(100, 200)

}

func getProfileOperation(oper func(int, int)) func(int, int) {
	return func(n1, n2 int) {
		start := time.Now()
		oper(n1, n2)
		elapsed := time.Now().Sub(start)
		log.Println("Time taken :", elapsed)
	}
}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(n1, n2 int) {
		log.Println("Operation started")
		oper(n1, n2)
		log.Println("Operation completed")
	}
}

func add(n1, n2 int) {
	result := n1 + n2
	fmt.Println("Result :", result)
}

func subtract(n1, n2 int) {
	result := n1 - n2
	fmt.Println("Result :", result)
}

func multiply(n1, n2 int) {
	result := n1 * n2
	fmt.Println("Result :", result)
}
