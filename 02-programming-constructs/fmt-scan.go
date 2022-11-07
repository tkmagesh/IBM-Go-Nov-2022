package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Print("Enter your name :")
		//fmt.Scanln(&name)
		fmt.Scanf("%s\n", &name)
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/
	/*
		var no int
		fmt.Print("Enter a number :")
		fmt.Scanln(&no)
		fmt.Printf("The given number is %d\n", no)
	*/
	var n1, n2 int
	fmt.Print("Enter two numbers :")
	fmt.Scanln(&n1, &n2)
	fmt.Println(n1, n2)

}
