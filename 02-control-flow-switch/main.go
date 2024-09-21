package main

import "fmt"

func main() {
	x := 41
	switch {
	case x == 42:
		fmt.Print("Value is 42")
	case x > 42:
		fmt.Print("Value GREATER than 42")
	case x < 42:
		fmt.Print("Value LESS than 42")
	}
}