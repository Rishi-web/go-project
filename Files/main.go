package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Operations in Go")

	content := " File operations in Go are simple and efficient."
	file, err := os.Create("./file_op.txt")
	checkErrorNil(err)

	length, err := io.WriteString(file, content)
	checkErrorNil(err)

	fmt.Println("Length of file content written:", length)
	defer file.Close()
	readFile("./file_op.txt")
	fmt.Println("File operations completed successfully.")

}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)
	checkErrorNil(err)
	fmt.Println("File content:", string(dataByte))
}

func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
