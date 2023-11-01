package main

import (
	"fmt"
	"math"
)

func main() {
	// variable declartion

	fmt.Println("go" + "lang")
	fmt.Println(1+1)
	fmt.Println(1, "1" , "Will this work")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(!true)
	var a = "initial"
	fmt.Println(a)
	var b, d  int = 1,2
	fmt.Println(d,b)
	f := "apple"
	fmt.Println(f)
	ba := "this"
	fmt.Println(ba)

	// constants declartion
	const g = 3e20
	const n = 500000000000000
	const h =  g/n
	fmt.Println(h)
	fmt.Println(math.Sin(n))
	// for loops in go 
}
