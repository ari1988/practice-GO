package main

import "fmt"

func main() {
	defer foo() // we can use defer to close resources
	bar()
}

// func (r receiver) identifier (p parameter(s)) (return(s)) { code }

func foo() {
	fmt.Println("foo")
}

func bar(){
	fmt.Println("bar")
}