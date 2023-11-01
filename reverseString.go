package main
import (
	"fmt"
)

func reverseString(str string) string{
	var reversed string

	for i:=len(str)-1; i>=0; i--{
		reversed += string(str[i])
	}

	return reversed
}

func main(){
	var str string = "Hello world"
	str  = reverseString(str)


	fmt.Println(str)
}
