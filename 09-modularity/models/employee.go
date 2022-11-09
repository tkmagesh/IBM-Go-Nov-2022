package models

import "fmt"

type Employee struct {
	Id     int
	Name   string
	salary float32
}

func PrintEmployee(emp Employee) {
	fmt.Printf("Id = %d, Name = %q, Salary = %v\n", emp.Id, emp.Name, emp.salary)
}
