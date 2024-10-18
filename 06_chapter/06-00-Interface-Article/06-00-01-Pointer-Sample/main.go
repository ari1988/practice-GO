package main

import "fmt"

func main() {
	answer := 42
	answerPtr := &answer

	fmt.Println(answerPtr)
	fmt.Println(*answerPtr)
}