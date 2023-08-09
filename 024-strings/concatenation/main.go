package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello World!"
	fmt.Println("String :", name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x", name[i])
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println("---------------------------------")
	fmt.Println("Three ways to concatenate strings")
	fmt.Println()

	//using + not very efficient
	h := "Hello"
	w := "World."
	myString := h + w
	fmt.Println("With + sign:", myString)

	//using fmt more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println("With fmt:", myString)

	// using stringbuilder- very efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println("With stringbuilder:", sb.String())

	fmt.Println()
	name = "ABCDEFGHIJKLMNOPRS"
	fmt.Println("Getting a substring")
	fmt.Println(name[3:9])
}
