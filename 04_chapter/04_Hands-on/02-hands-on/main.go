package main

import "fmt"

func main() {
	m := make(map[string][]string) // key of string and value of type slice string
	m[`bond_james`] = []string{`shaken, not stirred`,`martinis`,`fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`,`literature`,`computer science`}
	m[`no_dr`] = []string{`cats`,`ice cream`,`sunsets`}
	m[`fleming_ian`] = []string{`steaks`,`cigars`,`espionage`}

	// fmt.Printf("%#v\n",m)
	
	for k,v := range m {
		fmt.Println(k)
		for i,v2 := range v {
			fmt.Println(i,v2)
		}
	}

	fmt.Println("----------------Record Deleted------------")
	delete(m, "fleming_ian")
	fmt.Println("----------------Record Deleted------------")
	for k,v := range m {
		fmt.Println(k)
		for i,v2 := range v {
			fmt.Println(i,v2)
		}
	}


}