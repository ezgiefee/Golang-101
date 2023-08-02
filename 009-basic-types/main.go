package main

import (
	"fmt"
	"log"
)

// basic types(numbers,strings,booleans)
var myInt int

//var myInt16 int16
//var myInt32 int32
//var myInt64 int64

var myUint uint //positive values
var myFloat float32
var myFloat64 float64 //much larger float number

func main() {
	fmt.Println("************* Basic Types *****************")
	myInt = 10
	myUint = 20
	myFloat = 10.1
	myFloat64 = 100.1

	fmt.Printf("Integer : %d\nUint : %d\nFloat32: %.2f\nFloat64: %.2f\n", myInt, myUint, myFloat, myFloat64)

	myString := "Ezgi"
	log.Println(myString)
	myString = "John" //it is wrong, strings are immutable in go

	var myBool = true
	fmt.Printf("Boolean : %t\n", myBool)

}
