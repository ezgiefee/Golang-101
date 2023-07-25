package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	//seed the random number generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	//declare then assign
	var firstNumber = random.Intn(8) + 2
	var secondNumber = random.Intn(8) + 2
	var subtraction = random.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	reader := bufio.NewReader(os.Stdin)

	//display a welcome/instructions
	fmt.Println("\nGuess the number game")
	fmt.Println("---------------------------")

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
	fmt.Println("The answer is ", answer)
}
