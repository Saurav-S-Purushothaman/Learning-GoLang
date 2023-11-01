package main

import "fmt"

func main(){
    this := []int {1,2,3,4,5}
	for _,num := range this {
		fmt.Println("index =      number = " , num)
	}
	 kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
}
