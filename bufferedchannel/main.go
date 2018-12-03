package main

import (
	"fmt"
	"strconv"
	"time"
)

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

type Salutations []Salutation

func (salutation *Salutation) Greet(do Printer) {
	do(salutation.greeting + " " + salutation.name)
}

func (salutations *Salutations) Greet(do Printer) {
	for _, salutation := range *salutations {
		salutation.Greet(do)
	}
}

func (salutations *Salutations) CreatePrintFunction(addon string) Printer {
	return func(greeting string) {
		fmt.Println(greeting + addon)
	}
}

func main() {
	salutations := Salutations{
		{"Bob", "Hello"},
		{"Ravi", "Hey"},
		{"Vincent", "Ola"},
	}

	done := make(chan bool, 2)
	go func() {
		salutations.Greet(salutations.CreatePrintFunction("<c>"))
		done <- true
		time.Sleep(100 * time.Millisecond)
		done <- true
		close(done) // use thic to close the channel and te remove the dead lock condition
		fmt.Println("Done !")
	}()
	salutations.Greet(salutations.CreatePrintFunction("?"))
	// <-done
	// for {
	// 	time.Sleep(100 * time.Millisecond)
	// }
	count := 0
	for value := range done { // async looping over the data fromt he channel
		if value {
			count = count + 1
			fmt.Println("Value came in channel " + strconv.FormatBool(value))
		}
		// if count == 2 {
		// 	close(done)
		// 	break
		// }
	}
}
