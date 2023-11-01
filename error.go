package main

import "fmt"
import "errors"

func main(){
	result,err := userJustFarted()
	if result,err:=userJustFarted(); err !=nil{
		fmt.Println("User has to change their pants")
		fmt.Println(result)
	}
	_ = result
	_ = err

}

func userJustFarted() (int, error){
	smell := 1000;
	if smell >  100 {
		return smell, errors.New("this is my error")
	}
	return smell, nil
}





