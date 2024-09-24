package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, v := range x {
		fmt.Printf("%v\t %T\n", v, v)
	}

	fmt.Printf("%T \t %#v\n",x,x) // Gives type in the value as well
	fmt.Printf("%T \t %v",x,x)
}
