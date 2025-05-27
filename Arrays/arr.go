package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")

	var fl = [4]int{1, 2, 3}

	fmt.Println("Array:", fl)
	fmt.Println("Length of array:", len(fl))
	fmt.Println("Capacity of array:", cap(fl))
}
