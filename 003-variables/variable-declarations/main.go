package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready."

func main() {
	//declare then assign
	var firstNumber int
	firstNumber = 2

	//declare type and name and assign
	var secondNumber = 6

	//declare name, assign value and let go figure out the type
	subtraction := 7

	//display a welcome/instructions
	fmt.Println("\nGuess the number game")
	fmt.Println("---------------------------")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nThink of a number between 1 and 10 ", prompt)
	reader.ReadString('\n')

	//take them through the game
	fmt.Println("Multiply your number by ", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of ", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract ", subtraction, prompt)
	reader.ReadString('\n')

	//give them the answer
	answer := firstNumber*secondNumber - subtraction
	fmt.Println("The answer is ", answer)
}
