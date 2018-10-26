package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)
	if err == nil {
		fmt.Println(greeting)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getGreeting(hour int) (string, error) {

	greeting := ""
	if hour < 7 {
		err := errors.New("Too early for greetings!")
		return greeting, err
	}
	if hour < 12 {
		greeting = "Good Morning"
	} else if hour < 18 {
		greeting = "Good Afternoon"
	} else {
		greeting = "Good Evening"
	}

	return greeting, nil
}
