package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
	first string
}

func main() {
	sa1 := secretAgent {
		person: person {
			first: "James",
			last: "Bond",
			age: 20,
		},
		first: "Ethan Hawk",
		ltk : true,
	}
	p2 := person{"Jenny","Moneypenny",27}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Printf("%T \t %#v\n", sa1,sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.person)
	fmt.Println("---------------------------------")
	fmt.Println(sa1.first," | ", sa1.person.first)

}