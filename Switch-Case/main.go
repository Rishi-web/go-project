package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Go")

	//rand.Seed(time.Now().UnixNano()) // deprected in latest Go versions
	// Use rand.New instead to create a new random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	diceNum := r.Intn(6) + 1

	switch diceNum {
	case 1:
		fmt.Println("You rolled a 1")
	case 2:
		fmt.Println("You rolled a 2")
	case 3:
		fmt.Println("You rolled a 3")
	case 4:
		fmt.Println("You rolled a 4")
	case 5:
		fmt.Println("You rolled a 5")
		fallthrough // fallthrough allows the next case to execute
	case 6:
		fmt.Println("You rolled a 6, roll again!")
	default:
		fmt.Println("Invalid roll")

	}
}
