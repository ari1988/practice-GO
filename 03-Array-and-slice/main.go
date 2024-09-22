package main

import "fmt"

func main() {
	// array literal
	a := [3]int{42,43,44}
	fmt.Println(a)

	b := [...]string{"Hello", "Gophers"}
	fmt.Println(b)

	// Another way
	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c[0],"\t",c[1])

	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)
}