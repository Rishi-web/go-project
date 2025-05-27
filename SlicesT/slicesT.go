package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")

	var fruitList = []string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Println("Fruit List:", fruitList)

	fruitList = append(fruitList, "Mango", "Orange")
	fmt.Println("Updated Fruit List:", fruitList)

	fruitList = append(fruitList[1:3], "Grapes")
	fmt.Println("Sliced Fruit List:", fruitList)

	valS := make([]int, 4)
	valS[0] = 10
	valS[1] = 20
	valS[2] = 30
	valS[3] = 100

	fmt.Println("Slice of integers:", valS)
	valS = append(valS, 50, 60)
	fmt.Println("Updated Slice of integers:", valS)

	sort.Ints(valS)
	fmt.Println("Sorted Slice of integers:", valS)

	// how to remove an element from a slice based on index
	index := 2
	var courses = []string{"Go", "Python", "Java", "C++", "JavaScript"}
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after removing index 2:", courses)
}
