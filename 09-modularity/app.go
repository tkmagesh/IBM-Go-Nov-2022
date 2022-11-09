package main

import (
	"fmt"
	_ "modularity-demo/calculator" // ONLY to execute the init functions
	appUtils "modularity-demo/calculator/utils"
	"modularity-demo/models"

	"github.com/fatih/color"
	carbon "github.com/golang-module/carbon/v2"
)

func main() {
	color.Red("app executed")
	fmt.Println(carbon.Now())
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("OpCount =", calculator.OpCount())
	*/

	//fmt.Println("Is 21 even number ? :", utils.IsEven(21))
	fmt.Println("Is 21 even number ? :", appUtils.IsEven(21))
	var emp = models.Employee{Id: 100, Name: "Magesh"}
	models.PrintEmployee(emp)

}
