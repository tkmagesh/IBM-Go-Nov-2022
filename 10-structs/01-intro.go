package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{}
	//var emp Employee = Employee{100, "Magesh", 10000}
	//var emp Employee = Employee{Id: 100, Name: "Magesh", Salary: 10000}
	var emp Employee = Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	fmt.Printf("%#v\n", emp)

	//accessing the fields
	fmt.Println(emp.Id, emp.Name, emp.Salary)

	//struct pointers
	//var empPtr *Employee
	//var empPtr *Employee = &Employee{}
	//var empPtr = new(Employee)
	var empPtr = &Employee{
		Id:     200,
		Name:   "Suresh",
		Salary: 10000,
	}
	fmt.Println(empPtr)
	//fmt.Println((*empPtr).Id, (*empPtr).Name, (*empPtr).Salary)
	fmt.Println(empPtr.Id, empPtr.Name, empPtr.Salary)

}
