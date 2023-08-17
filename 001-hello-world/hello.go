package main

import "fmt"

func main() {
	whatToSay := "Hello world!"
	sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
