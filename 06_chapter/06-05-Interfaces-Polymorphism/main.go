package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am",p.first)
}

func (sa secretAgent) speak() { // If you have the methods I have, you are my type
	fmt.Println("I am secret agent",sa.first)
}

type human interface {
	speak()
}

func saySomething (h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent {
		person: person{
			first: "James",
		},
		ltk: true,
	}

	p2 := person {
		first: "Jenny",
	}
	// sa1.speak()
	// p2.speak()

	saySomething(sa1)
	saySomething(p2)
}

// An Interface in Go defines a set of method signatures
// Polymorphism is the ability of a VALUE of a certain TYPE to also be of another TYPE
// In Go, values can be of more than one type.
