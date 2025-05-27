package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	var val = 10
	var ptr = &val
	// default value of ptr is nil
	fmt.Println("Value of val:", ptr)
	fmt.Println("Value of val:", *ptr)

	*ptr = *ptr + 10
	fmt.Println("Value of val after increment:", *ptr)
}
