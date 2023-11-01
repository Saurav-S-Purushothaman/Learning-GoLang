package main 

import "fmt"


// func Add(a int, b int) int {
// 	return a+b
// }

// func AddFloat(a float64, b float64) float64{
// 	return a+b
// }

// for adding float different versions of the functions need to be kept. 
// so there is generics 

type num interface {
	int | int16 | int32 | float32 | float64
}

func Add[T num]( a T , b T) T {
	return a+b
}


func main(){
	fmt.Println(Add(1,2))
	fmt.Println(Add(1.1,2.1)) // this is now possible 
	fmt.Println(Add(1.3,2))
}