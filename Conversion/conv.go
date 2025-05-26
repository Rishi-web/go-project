package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conversion in Go")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number to convert from string : ")

	input, _ := reader.ReadString('\n')

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("Converted number is :", num)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Printf("Converted number is of type %T\n", num)
	}

}
