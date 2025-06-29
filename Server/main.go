package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func getResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a GET request response")
}

// func postResponse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is a POST request response")
// }

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", getResponse)
	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
