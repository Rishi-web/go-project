package main

import "fmt"

func main() {
	fmt.Println("functions in Go")
	greeter()
	result := adder(10, 5)
	fmt.Println("Result of adder function:", result)

	newVal := proAdder(1, 2, 8, 9, 6, 5)
	fmt.Println("Result of proAdder function:", newVal)
}

func greeter() {
	fmt.Println("Hello, World!")
}

func adder(val1 int, val2 int) int { //function with parameters and return type and its called function signature
	return val1 + val2
}

func proAdder(values ...int) int { //variadic function
	total := 0
	for _, val := range values {
		total += val
	}
	return total

}
