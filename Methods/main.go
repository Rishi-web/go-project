package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	fmt.Println("Structs in Go")

	user1 := User{"RV", 30, "rv.com", true}

	fmt.Println("User 1:", user1)
	fmt.Printf("User1 details %+v\n", user1)
	fmt.Printf("Name is %v \n", user1.Age)
	user1.getStatus()
}

func (u User) getStatus() {
	fmt.Printf("User %s is active: %t\n", u.Name, u.Status)
}
