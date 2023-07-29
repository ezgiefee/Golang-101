package main

import (
	"fmt"
	"motus/scope/packagetwo"
)

var myVar = "This is a package level variable"

func main() {
	//variables
	//declare a package level variable for the main
	//package named myVar

	//declare a block variable for the main function
	//called blockVar
	blockVar := "This is a block level variable"

	//declare a package level variable in the packageone
	//package named PackageVar

	//create an exported function in packageona called PrintMe

	//in the main function, print out the values of myVar,
	//blockVar and PackageVar on one line using the PrintMe
	//function in packageone
	packagetwo.PrintMe(myVar, blockVar)

	fmt.Println(blockVar)
}
