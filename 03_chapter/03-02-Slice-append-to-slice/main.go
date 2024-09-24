package main

import "fmt"

func main() {
	xi := []int{43,44,45}
	fmt.Println(xi)
	fmt.Println("-------------")
	// variadic parameter
	xi = append(xi, 45,46,47,99,777)
	fmt.Println(xi)
	fmt.Println("-------------")
	fmt.Println("Slicing a Slice")
	fmt.Println("-------------")
	fmt.Printf("xi - %#v\n",xi)
	// [inclusive:exclusive]
	fmt.Printf("xi - %#v\n",xi[0:4])
	//[:exclusive]
	fmt.Printf("xi - %#v\n",xi[:7])
	// [inclusive:]
	fmt.Printf("xi - %#v\n",xi[3:])

	// [:]
	fmt.Printf("xi - %#v\n",xi[:])
	fmt.Println("-------------")

	fmt.Println("Deleting from a slice")
	fmt.Println("-------------")

	xi = append(xi[:4], xi[5:]...)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")
}