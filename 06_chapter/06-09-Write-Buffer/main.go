package main

import (
	"bytes"
	"fmt"
)

/*
// String to a byte slice

str := "Hello"
bytes := []byte(str)

// Byte slice to string

bytes := []byte{72,101,108,108,111}
str := string(bytes)
*/

func main() {
	b := bytes.NewBufferString("Hello ")
  fmt.Println(b.String())
  b.WriteString("Gophers")
  fmt.Println(b.String())
  b.Reset() // Resets the buffer
  b.WriteString("It's Thursday")
  fmt.Println(b.String())
  b.Write([]byte("--Happy Happy Day"))
  fmt.Println(b.String())
}