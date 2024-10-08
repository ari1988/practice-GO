package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is from init function.")
}
func main() {
	x := rand.Intn(350)
	fmt.Printf("The value of x is %v\t", x)
	/*
		if x <= 100 {
			fmt.Println("less than 100.")
		} else if x >= 100 && x <= 200 {
			fmt.Println("between 101 and 200.")
		} else if x >= 201 && x <= 250 {
			fmt.Println("between 201 and 250.")
		} else {
			fmt.Println("this was more than 250.")
		}
	*/
	switch {
	case x <= 100:
		fmt.Println("less than 100.")
	case x >= 100 && x <= 200:
		fmt.Println("between 101 and 200.")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250.")
	default:
		fmt.Println("this was more than 250.")
	}
}
