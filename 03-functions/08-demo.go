/* higher order functions - return a function as a return value from another function - use case*/

package main

import (
	"fmt"
	"log"
)

func main() {

	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	logAdd := getLogOperation(add)
	logAdd(100, 200)

	logSubtract := getLogOperation(subtract)
	logSubtract(100, 200)

	getLogOperation(multiply)(100, 200)

}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(n1, n2 int) {
		log.Println("Operation started")
		oper(n1, n2)
		log.Println("Operation completed")
	}
}

/*
func logOperation(oper func(int, int), n1, n2 int) {
	log.Println("Operation started")
	oper(n1, n2)
	log.Println("Operation completed")
}
*/

/*
func logAdd(n1, n2 int) {
	log.Println("Operation started")
	add(n1, n2)
	log.Println("Operation completed")
}

func logSubtract(n1, n2 int) {
	log.Println("Operation started")
	subtract(n1, n2)
	log.Println("Operation completed")
}
*/

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
