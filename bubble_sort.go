package main

import "fmt"

func bubbleSort(arr []int){
	n := len(arr)
	for sorted:=true;sorted; sorted=false {
		sorted = false
		for  i :=0; i<n-1; i++ {
			if arr[i] > arr[i+1] {
				sorted = true
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func main(){
	arr := []int{ 4,1,5}
	bubbleSort(arr)
	fmt.Println(arr)
}
