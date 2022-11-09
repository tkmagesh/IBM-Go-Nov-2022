package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization
}

func main() {
	ibm := &Organization{Name: "IBM", City: "Bangalore"}
	emp1 := Employee{100, "Magesh", 10000, ibm}
	emp2 := Employee{200, "Suresh", 20000, ibm}
	fmt.Println(emp1)
	fmt.Println(emp2)

	ibm.City = "Pune"
	fmt.Println(emp1.Org.City)
	fmt.Println(emp2.Org.City)

	/*
		e1 := Employee{100, "Magesh", 10000, ibm}
		e2 := Employee{100, "Magesh", 10000, ibm}
		fmt.Println(e1 == e2)
	*/
}
