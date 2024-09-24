package main

import "fmt"

func main() {
	// Create an array
	x := [5]int{}
	// assign values to each index position
	for i := 0; i < 5; i++ {
		x[i] = i
	}
	for _, v := range x {
		fmt.Printf("%v\t %T\n", v, v)
	}
}
