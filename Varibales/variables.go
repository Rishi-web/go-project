package main

import "fmt"

const LoginToken = "abc123"

// LoginToken is a constant variable that cannot be changed
// It is a good practice to use uppercase letters for constant variables

func main() {
	fmt.Println("Variables in Go")
	var username = "Rishi"

	fmt.Println("Username :", username)
	fmt.Printf("Varibale if of type %T\n", username)

	var isBool bool = true
	fmt.Println("Is Bool :", isBool)
	fmt.Printf("Varibale is of type %T\n", isBool)

	var intVal int = 10
	fmt.Println("Int Value :", intVal)
	fmt.Printf("Varibale is of type %T\n", intVal)

	var smalVal uint8 = 255
	fmt.Println("Small Value :", smalVal)
	fmt.Printf("Varibale is of type %T\n", smalVal)

	var floatVal float64 = 10.5
	fmt.Println("Float Value :", floatVal)
	fmt.Printf("Varibale is of type %T\n", floatVal)

	// default values and aliases
	var defaultVal int
	fmt.Println("Default Value :", defaultVal)

	// implicit type declaration
	var implicitVal = "Implicit Type"
	fmt.Println("Implicit Value :", implicitVal)
	fmt.Printf("Implicit Varibale is of type %T\n", implicitVal)

	// no var keyword
	ranNum := 48 // works only inside functions cannot be used at package level
	// := is a short variable declaration operator
	// it infers the type of the variable based on the value assigned to it
	// it can only be used inside functions, not at the package level
	// it can be used to declare multiple variables at once
	// e.g. ranNum, anotherVar := 48, "Hello"
	// it cannot be used to re-declare a variable that already exists in the same scope
	// it can be used to declare a variable with the same name in a different scope
	// e.g. func() { ranNum := 50 } // this is a different scope
	// it is a shorthand for var ranNum = 48
	fmt.Println("Random Number :", ranNum)
	fmt.Printf("Random Number is of type %T\n", ranNum)

}
