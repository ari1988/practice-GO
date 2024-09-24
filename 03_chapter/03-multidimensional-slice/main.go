package main

import "fmt"

func main() {
	jb := []string{"James","Bond","Martini","Chocolate"}
	jm := []string{"Jenny","MoneyPenny","Guiness","Wolverine Tracks"}
	hm := []string{"Danny","Badrah","Sundar","Kandi"}
	dm := []string{"Danny1","Badrah1","Sundar1","Kandi1"}
	fmt.Println(jb)
	fmt.Println(jm)
	fmt.Println(hm)
	fmt.Println(dm)

	xxs := [][]string{jb,jm,hm,dm}
	fmt.Println("index\t value")
	for i,v := range xxs {
		fmt.Printf("%v \t %v\t\n",i,v)
	}
	// fmt.Println(xxs)
}