package main

import "fmt"
import "reflect"

func main(){
	var arr[5] int 
	fmt.Println("emp:", arr)

	arr[4] = 100
	fmt.Println("printing the entire array :", arr)
	fmt.Println("arr[4] =", arr[4])
	fmt.Println("len(arr)=", len(arr))

	// best way to initialize an array
	
	arr2 := [5] int {1,2,3,4,5}
	fmt.Println("arr2 = " , arr2)

	// 2d arrays 
	var twoD [2][3] int
	for i:=0; i< 2; i++{
		for j:=0; j<3; j++ {
			twoD[i][j] = i+j
		}
	}
    fmt.Println("twoD",  twoD)
	fmt.Println("2d ", reflect.TypeOf(twoD))
}
