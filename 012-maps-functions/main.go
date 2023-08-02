package main

import "fmt"

// reference types (maps)
type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s.\n", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs\n", a.Name, a.NumberOfLegs)
}

func main() {
	fmt.Println("\n****************** Maps ******************")
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5
	for key, value := range intMap {
		fmt.Println(key, value)
	}
	delete(intMap, "four") //delete an element with key
	fmt.Println(intMap)

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}

	intMap["two"] = 4
	fmt.Println(intMap)

	fmt.Println("\n****************** Functions ******************")
	x := 5
	y := 7
	z := addTwoNumbers(x, y)
	fmt.Printf("Sum of %d and %d is %d\n", x, y, z)

	myTotal := sumMany(2, 3, 4, 5)
	fmt.Println(myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "Miaw",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()
}

func addTwoNumbers(x, y int) int {
	return x + y
}

// variadic function
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}
	return total
}
