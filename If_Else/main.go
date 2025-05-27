package main

import "fmt"

func main() {
	fmt.Println("if else in Go")

	loginCOunt := 25

	var result string
	if loginCOunt < 10 {
		result = "Welcome! You are a new user."
	} else {
		result = "Welcome back! You are a returning user."
	}
	fmt.Println(result)

	if num := 15; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than or equal to 10")
	}
}
