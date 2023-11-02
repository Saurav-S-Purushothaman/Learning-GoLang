package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string) 

	go func(){
		time.Sleep(2 * time.Second) 
		one <- "Received msg from channel one"
	}()
	go func(){
		time.Sleep(1 * time.Second) 
		one <- "Received msg from channel two"
	}()
    // this select statement would wait for the first completed channel and end the session
	for i:=0; i< 2; i++ {
	select {
	case rec1 := <-one:
		fmt.Println("await over", rec1)
	case rec2 := <-two:
		fmt.Println("await over", rec2)
	}
}

}
