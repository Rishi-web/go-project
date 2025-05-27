package main

import "fmt"

func main() {
	fmt.Println("loops in Go")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("Days of the week:", days)

	for i := 0; i < len(days); i++ {
		fmt.Println("Day", i+1, ":", days[i])
	}

	// another way to loop through the slice
	for i := range days {
		fmt.Println("Day", i+1, ":", days[i])
	}

	// another way to loop through the slice
	for i, day := range days {
		fmt.Println("Day", i+1, ":", day)
		if i == 3 {
			goto lab
		}
	}
	for _, day := range days {
		fmt.Println("Day:", day)
	}

lab:
	fmt.Println("Jummping here")
}
