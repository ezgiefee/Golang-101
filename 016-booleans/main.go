package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	apples := 18
	oranges := 9

	//boolean expression
	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// > < >= <=
	fmt.Printf("%d > %d : %t\n", apples, oranges, apples > oranges)
	fmt.Printf("%d < %d : %t\n", apples, oranges, apples < oranges)
	fmt.Printf("%d >= %d : %t\n", apples, oranges, apples >= oranges)
	fmt.Printf("%d <= %d : %t\n", apples, oranges, apples <= oranges)

	jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   40000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill Jones",
		Age:      32,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 or older")
		} else {
			fmt.Println(x.Name, "is under 30")
		}
		//compound boolean
		if x.Age > 30 && x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50000 and is over 30")
		} else {
			fmt.Println("Either ", x.Name, "makes less than 50000 and under 30")
		}

		if x.Age > 30 || x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50000 or is over 30")
		} else {
			fmt.Println(x.Name, "makes less than 50000 or under 30")
		}

		if x.Age > 30 || x.Salary < 50000 && x.FullTime {
			fmt.Println(x.Name, "matches our unclear criteria")
		}
	}
}
