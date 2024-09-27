package main

import "fmt"

func main() {
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78

	fmt.Println(an)
	fmt.Printf("%#v\n",an)

	for k,v := range an {
		fmt.Println(k,v)
	}

	for _,v := range an {
		fmt.Println(v)
	}

	for k := range an { //retreives just the key
		fmt.Println(k)
	}

	xi := []int{42,43,44}
	for k,v := range xi {
		fmt.Println(k,v)
	}

	for _,v := range xi {
		fmt.Println(v)
	}

	for i := range xi { //retreives just the key
		fmt.Println(i)
	}
}