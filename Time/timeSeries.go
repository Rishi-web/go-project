package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in Go")

	pTime := time.Now()

	fmt.Println("Current Time:", pTime)

	fmt.Println(pTime.Format("01-02-2006 15:04:05 Monday"))

	time.Sleep(5)
}
