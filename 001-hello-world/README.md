## # Hello World! 

* **package main** : Every go file must start with the *package name* statement. Packages are used to provide code compartmentalisation and reusability. Here the package name used is main

* **import "fmt"** : The *fmt* package is imported and it will be used inside the sayHelloWorld function to print text to the standard output.

* **func main()** : The *main* is a special function. The program execution starts from the main function. The main function should always reside in the main package. The { and } indicate the start and end of the main function.

* **func sayHelloWorld(whatToSay string)** : The sayHelloWorld is a function. It takes an argument with the string type and uses fmt package's Println function to print text to standard output.