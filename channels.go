package main 

import "fmt"
import "reflect"

func main() {
	messages := make(chan string) 
	go func(){
		messages <- "ping"
	} ()
	msg := <-messages
	fmt.Println(reflect.TypeOf(messages))
	fmt.Println(reflect.TypeOf(msg))
	fmt.Println(msg)
}