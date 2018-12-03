package main

// Like switch but on communications
// Rules
// 	Execute case that is ready
// 	If more than one is ready execute one at randon
// 	If none are ready block unles edfault is defined.

import (
	"fmt"
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
	for {
		select {
		case value, ok := <-done:
			if value {
				fmt.Println("Value in channel done")
			}
			if !ok {
				fmt.Println("done channel closed")
				return
			}
		default:
			//fmt.Println("Printing fromdeafult block")
		}
	}
}
