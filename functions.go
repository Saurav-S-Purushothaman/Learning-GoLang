package main

import "fmt"

func plus(a , b int) int {
	return a+b
}

// function with multiple return value 

func multReturnValFunc() (int, int) {
	return 1,2
}


// variadic functions 
// func variadicFunction(nums... int){
// 	fmt.Println(nums, " ")
// }
func main(){
	fmt.Println(plus(12,12))
	a,b := multReturnValFunc()
	fmt.Println("a", a, "b", b)
	// fmt.Println(variadicFunction(1,2,3,4,5))
}
