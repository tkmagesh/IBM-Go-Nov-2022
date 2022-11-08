package main

import (
	"fmt"
	_ "modularity-demo/calculator" // ONLY to execute the init functions
	appUtils "modularity-demo/calculator/utils"
)

func main() {
	fmt.Println("app executed")

	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("OpCount =", calculator.OpCount())
	*/

	//fmt.Println("Is 21 even number ? :", utils.IsEven(21))
	fmt.Println("Is 21 even number ? :", appUtils.IsEven(21))
}
