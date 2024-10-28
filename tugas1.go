package main

import "fmt"

func main() {

	var input int
	var message string = "Hello Universe"
	fmt.Println("Welcome to", message)
	fmt.Println("Input Number: ")
	fmt.Scanln(&input)

	if(input == 42){
		fmt.Println(message)
	} else {
		fmt.Println(input)
	}
}