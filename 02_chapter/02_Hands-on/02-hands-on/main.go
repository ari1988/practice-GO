package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)"}
	fmt.Println(xs)
	fmt.Println(len(xs))

	for i, v := range xs {
		fmt.Printf("Index %v\t Value %v\n", i, v)
	}

	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}
	fmt.Println("--------------------")
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}
}
