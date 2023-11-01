package main

import (
	"fmt"
)

func main(){
	m := make(map[string]int)

	newMap := make(map[int]string)
	m["k"] =1 
	m["s"] =2

	fmt.Println("Map", m)

	v := m["p"]
	fmt.Println(v)
	newMap[1] ="this"
	fmt.Println(newMap[2])

	fmt.Println("len", len(newMap))

	delete(m, "k")
	fmt.Println("Deleted m", m)

	clear(m)
	fmt.Println("cleared m", m)
	
	_, prs := m["k2"]

	fmt.Println(prs)

}

