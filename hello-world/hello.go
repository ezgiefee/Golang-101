package main

import "fmt"

func main() {

	whatToSay := "Hello world, again!"

	sayHelloWorld(whatToSay)

}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
