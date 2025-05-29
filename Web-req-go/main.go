package main

import (
	"fmt"
	"io"
	"net/http"
)

const getUrl = "http://localhost:8080/get"

func main() {
	fmt.Println("Web Requests in Go")
	performGetRequest(getUrl)
}

func performGetRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Status Code:", response.StatusCode)
	fmt.Println("Response Body Content:", string(bodyBytes))
}
