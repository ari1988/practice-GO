package main

import "fmt"

func main(){
	a := []int{0,1,2,3,4,5}
	// b := a // any change in it references to original array

	b := make([]int, 6)
	copy(b, a) // If a value is updated now, b is going to be the same

	fmt.Println("a ",a)
	fmt.Println("b ",b)
	fmt.Println("----------------")

	a[0] = 7
	fmt.Println("a ",a) // both a & b refers to the same underlying array
	fmt.Println("b ",b)
	fmt.Println("----------------")
}