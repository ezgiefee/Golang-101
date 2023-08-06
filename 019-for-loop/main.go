package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"motus/mylogger"
	"os"
	"time"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("i is", i)
	}

	//execute a loop while i>100
	rand.NewSource(time.Now().UnixNano())
	i := 1000

	for i > 100 {
		//get a random between 1 and 1001
		i = rand.Intn(1000) + 1
		fmt.Println("i is", i)
		if i > 100 {
			fmt.Println("i is", i)
		} else {
			fmt.Println("Got", i, "and broke out of the loop")
		}
	}

	//infinite loop
	//read input from the user 5 times and write it to a log
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)
	go mylogger.ListenForLog(ch)

	fmt.Println("Enter something")

	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
