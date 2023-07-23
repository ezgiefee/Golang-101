package main

import (
	"bufio"
	"fmt"
	"motus/eliza/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		//The actual return appended to "quit" word. Strip off the part where the user pressed enter.
		//Overwrite the user input where it equals carriage return.
		userInput = strings.Replace(userInput, "\r\n", "", -1) //for windows users
		userInput = strings.Replace(userInput, "\n", "", -1)   //for other users

		if userInput == "quit" {
			break //get out of the loop
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}

}
