package main

import "fmt"

func main() {
	fmt.Println("Defer in Go")
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
	defer fmt.Println("This will be printed second")
	defer fmt.Println("This will be printed third")
	defer fmt.Println("This will be printed fourth")
	deferFunc()
}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
