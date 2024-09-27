package main

import "fmt"

func main() {
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78

	fmt.Println(an)
	fmt.Printf("%#v\n",an)
	
	delete(an, "George")
	fmt.Println(an)
	fmt.Printf("%#v\n",an)

	fmt.Println("--- accessing keys that don't exist")
	delete(an, "George") //won't panic
	fmt.Println(an["Georgey"]) //won't panic
	fmt.Println("------------------------------")
}
