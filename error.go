package main

import "fmt"
import "errors"

type argError struct{
	arg int 
	probStatement string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s\n", e.arg, e.probStatement)
}

fund testError(input int) (int, error) {
	if input == 32{
		return -1, &argError{input, "INFO: 32 not supported"}
	}


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





