package main

import (
	"log"
	"motus/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Jill", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Sam", LastName: "Smithers", Salary: 80000, FullTime: true},
	{FirstName: "Andrew", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Grimm", LastName: "Carter", Salary: 10000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	//staff.OverPaidLimit = 10000

	//log.Println(myStaff.All())
	log.Println("Overpaid staff", myStaff.Overpaid())
	log.Println("Underpaid staff", myStaff.Underpaid())

}
