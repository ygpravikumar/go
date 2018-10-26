package main

import (
	"fmt"
	"time"
)

func main() {
	hourOfDay := time.Now().Hour()
	greeting := getGreeting(hourOfDay)
	fmt.Println(greeting)
}

func getGreeting(hour int) string {
	greeting := ""
	if hour < 12 {
		greeting = "Good Morning"
	} else if hour < 18 {
		greeting = "Good Afternoon"
	} else {
		greeting = "Good Evening"
	}

	return greeting
}
