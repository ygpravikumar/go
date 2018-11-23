package main

import "fmt"

type gopher struct {
	name string
	age  int
}

func (g gopher) jump() string {
	if g.age > 65 {
		return g.name + " can still jump"
	}
	return g.name + " can jump high."
}

type horse struct {
	name   string
	weight int
}

func (g horse) jump() string {
	if g.weight > 2500 {
		return g.name + " can not jump. Iam too heavy."
	}
	return g.name + " can jump !Neigh"
}

type jumper interface {
	jump() string
}

func main() {
	jumperList := getJumperList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.jump())
	}
}

func getJumperList() []jumper {
	phill := &gopher{name: "phill", age: 21}
	john := &gopher{name: "john", age: 81}
	rider := &horse{name: "rider", weight: 1800}
	list := []jumper{phill, john, rider}
	return list
}
