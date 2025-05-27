package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	lang := make(map[string]string)

	lang["go"] = "Go is a statically typed, compiled programming language designed at Google."
	lang["py"] = "Python is an interpreted, high-level, general-purpose programming language."
	lang["ja"] = "Java is a high-level, class-based, object-oriented programming language that is designed to have as few implementation dependencies as possible."
	lang["js"] = "JavaScript is a high-level, dynamic, untyped, and interpreted programming language."

	//fmt.Println("Languages Map:", lang)
	fmt.Println("Go Language:", lang["go"])
	fmt.Println("Python Language:", lang["py"])

	delete(lang, "ja")
	fmt.Println("Language after deletion:", lang)

	// loop through the map
	for key, value := range lang {
		fmt.Println("Key:", key, "Value:", value)
	}

}
