package main

import (
	"fmt"
)

// aggregate types(array, struct)
type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	fmt.Println("\n************* Arrays **********************")
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	for i := 0; i < 3; i++ {
		fmt.Printf("%d. element in array is %s\n", i+1, myStrings[i])
	}

	fmt.Println("\n************* Structs **********************")
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}
	fmt.Printf("My car is %d years old, its make is %s and its model is %s,Luxury: %t, BucketSeats: %t\n",
		myCar.NumberOfTires,
		myCar.Make,
		myCar.Model,
		myCar.Luxury,
		myCar.BucketSeats)

	fmt.Println("***************************************")
}
