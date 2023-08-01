package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)

	userName := readString("What is your name?")

	age := readInt("How old are you?")

	//fmt.Println("Your name is", userName, "and you are", age, "years old.")
	//fmt.Println(fmt.Sprintf("Your name is %s and you are %d years old", userName, age))
	fmt.Printf("Your name is %s and you are %d years old.\n", userName, age) //string interpolation
}

func prompt() {
	fmt.Print("-> ")
}

func readString(str string) string {
	for {
		fmt.Println(str)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter your name")
		} else {
			return userInput
		}
	}
}

func readInt(str string) int {
	for {
		fmt.Println(str)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}

	}
}
