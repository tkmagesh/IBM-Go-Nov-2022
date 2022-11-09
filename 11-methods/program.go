package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

/*
func (Employee) WhoAmI() {
	fmt.Println("I am an Employee ")
}
*/

func (emp Employee) WhoAmI() {
	fmt.Printf("I am an Employee [%q]\n", emp.Name)
}

func (emp Employee) AwardBonus(bonus float32) {
	totalSalary := emp.Salary + bonus
	fmt.Println("Emp salary after bonus :", totalSalary)
}

func main() {
	/*
		var emp Employee
		emp.WhoAmI()
	*/

	emp := Employee{100, "Magesh", 10000}
	emp.WhoAmI()

	//Award bonus
	emp.AwardBonus(1000)
}
