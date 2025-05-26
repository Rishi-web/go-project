package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User input in Go")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter whats your name?")
	// Comma-ok syntax is used to read input from the user
	input, _ := reader.ReadString('\n')

	fmt.Println("Hello", input)
}
