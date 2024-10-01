package main

import "fmt"

type person struct {
	first string
	last string
	favIC []string
}

func main() {
	p1 := person {
		first: "James",
		last:"Bond",
		favIC : []string{"chocolate", "banana", "passion fruit with mango and guvava"},
	}

	p2 := person {
		first: "Jenny",
		last: "Jennypenny",
		favIC: []string{"Strawberry", "Chocolate"},
	}

	// fmt.Println(p1)
	// fmt.Println(p2)

	// fmt.Println(p1.favIC)
	// fmt.Println(p2.favIC)

	m := map[string]person {
		p1.last : p1,
		p2.last : p2,
	}

	for _,v := range m {
		fmt.Println(v)
		for _,v2:= range v.favIC {
			fmt.Println(v.first, v.last, v2)
		}
	}
}