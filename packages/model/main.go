package model

type gopher struct {
	name string
	age  int
}

func (g gopher) Jump() string {
	if g.age > 65 {
		return g.name + " can still jump"
	}
	return g.name + " can jump high."
}

type horse struct {
	name   string
	weight int
}

func (g horse) Jump() string {
	if g.weight > 2500 {
		return g.name + " can not jump. Iam too heavy."
	}
	return g.name + " can jump !Neigh"
}

type jumper interface {
	Jump() string
}

func GetJumperList() []jumper {
	phill := &gopher{name: "phill", age: 21}
	john := &gopher{name: "john", age: 81}
	rider := &horse{name: "rider", weight: 1800}
	list := []jumper{phill, john, rider}
	return list
}
