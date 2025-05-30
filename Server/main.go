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

func postResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a POST request response")
}

func main() {
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", getResponse)
}
