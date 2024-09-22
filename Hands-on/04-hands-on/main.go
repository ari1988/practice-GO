package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for i:= 0; i <100;i++ {
	// 	fmt.Println(i)
	// }

	for i :=0; i <100;i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("x and y are %v and %v\t",x,y)
	}
}