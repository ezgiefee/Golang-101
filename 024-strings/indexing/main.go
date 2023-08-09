package main

import "fmt"

func main() {
	//Indexes start at 0
	courseName := "Learn Go for Beginners Crash Course"
	fmt.Println(courseName[6])
	fmt.Println(string(courseName[0])) //casting the rune to string

	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()

	for i := 13; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()

	fmt.Println("Length of courseName is", len(courseName))
	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")

	fmt.Println("mySlice has ", len(mySlice), " elements")
	fmt.Println("The last element of mySlice is", mySlice[len(mySlice)-1])

}
