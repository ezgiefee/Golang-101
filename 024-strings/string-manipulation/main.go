package main

import (
	"fmt"
	"strings"
)

func main() {

	newString := "Go is a great programming language. Go for it!"
	if strings.Contains(newString, "Go") {
		//newString = strings.Replace(newString, "Go", "Golang", 1)
		newString = strings.ReplaceAll(newString, "Go", "Golang")
	}
	fmt.Println(newString)
	fmt.Println()

	if "Alpha" > "Absolute" {
		fmt.Println("Alpha is greater than Absolute")
	} else {
		fmt.Println("Alpha is not greater than Absolute ")
	}

	badEmail := " me@here.com  "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Printf("=%s=", badEmail)
	fmt.Println()

	str := "alpha alpha alpha alpha"
	str = replaceNth(str, "alpha", "beta", 3)
	fmt.Println(str)

	//case
	myString := "This is a clear example of why we search in one case only."
	searchString := strings.ToLower(myString)
	if strings.Contains(searchString, "this") {
		fmt.Println("Found it!")
	} else {
		fmt.Println("Did not find it!")
	}

	//other case functions
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
}

func replaceNth(s, old, new string, n int) string {
	//index
	i := 0
	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			//we didnt find it
			break
		}
		i = i + x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)

	}
	return s
}
