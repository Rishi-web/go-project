package main

import (
	"fmt"
	"net/url"
)

const urlW = "https://www.chrismytton.com/plain-text-websites/coursename=golang&course=web-requests#section1"

func main() {
	fmt.Println("URL Requests in Go")
	result, err := url.Parse(urlW)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed URL:", result)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Raw Query:", result.RawQuery)
	fmt.Println("Query Parameters:", result.Query())

	// constructing a new URL
	newUrls := url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "path/to/resource",
		RawPath:  "path/to/resource",
		RawQuery: "param1=value1&param2=value2",
	}
	fmt.Println("Constructed URL:", newUrls.String())

}
