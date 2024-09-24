package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 42

	if z:= 2 * rand.Intn(x); z >= x {
		fmt.Printf("Z is GREATER. Values of x is %d and z %d",x,z)
	} else {
		fmt.Printf("Z is SMALLER. Values of x is %d and z %d",x,z)
	}
}