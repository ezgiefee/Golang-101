package main

import (
	"fmt"
	"sort"
)

//refernce types(pointers,slices,maps,functions,chanels)

func main() {
	fmt.Println("************ Pointers***************")
	x := 10

	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	changeValueOfPointer(&x) //change value of x
	fmt.Println("After function call x is now", x)

	fmt.Println("************ Slices ***************")

	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("First two elements are", animals[0:2])
	fmt.Println("The slice is", len(animals), "elements long")
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = deleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func changeValueOfPointer(num *int) {
	*num = 25
}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
