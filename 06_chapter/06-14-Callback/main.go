package main

import "fmt"

func main() {
	x := doMath(1,2,add)
	fmt.Println(x)

	y := doMath(4,2,subtract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int,int) int) int {
	return f(a,b) // function takes a and b as input, therefore
	// when calling it in doMath we just pass the integers and the func name
}

func add(a,b int) int {
	return a + b
}

func subtract(a,b int) int {
	return a - b
}