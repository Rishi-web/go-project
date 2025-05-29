package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.chrismytton.com/plain-text-websites/"

func main() {
	fmt.Println("Web Requests in Go")

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Response Status Code:", response.StatusCode)
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body Content:", string(dataBytes))
}
