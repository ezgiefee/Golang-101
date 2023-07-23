package main

import (
	"fmt"
	"motus/eliza/doctor"
)

func main() {
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
}
