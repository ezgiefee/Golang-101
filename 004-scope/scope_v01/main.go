package main

import (
	"fmt"
	"motus/scope/packageone"
)

var second = "second" //package level variable

func main() {
	var one = "One" //block level variable
	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar //public variable
	fmt.Println(newString)

	packageone.Exported()
}

func myFunc() {
	var one = "the number one"
	fmt.Println(one)
	fmt.Println(second)
}
