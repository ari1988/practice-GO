package main

import "fmt"

func main() {
	answer := 42
	answerPtr := &answer

	*answerPtr = 99
	fmt.Println(answer)
}