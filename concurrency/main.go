package main

import (
	"fmt"
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

	done := make(chan bool)
	go func() {
		salutations.Greet(salutations.CreatePrintFunction("<c>"))
		done <- true
	}()
	salutations.Greet(salutations.CreatePrintFunction("?"))
	// time.Sleep(100 * time.Millisecond)
	<-done
}
