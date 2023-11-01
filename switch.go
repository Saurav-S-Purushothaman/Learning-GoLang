package main

import "fmt"

func whatAmI  ( i interface{}){
	switch t := i.(type){
	case bool:
		fmt.Println("I am a bool")
	case int:
		fmt.Println("I am an int")
	default:
		fmt.Println("Don't know the types %T\n",t)
	}
}
func main(){
	whatAmI(true)
	whatAmI(1)
	
}
