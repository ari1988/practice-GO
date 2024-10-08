package main

import "fmt"

func main() {
	m := map[string]int {
		"James": 42,
		"MoneyPenny": 32,
	}
	fmt.Println("-----------------------------")
	for k,v := range m{
		fmt.Printf("Key %v \t Value %v \n",k,v)
	}
	fmt.Println("-----------------------------")
	age1 := m["James"]
	fmt.Println("The age of Bond : ",age1)
	if v, ok := m["James"]; ok {
		fmt.Println("There is a BOND lookup, and the value is",v)
	}
	age2 := m["Q"]
	fmt.Println("The age of Q : ",age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int",v)
	}
}