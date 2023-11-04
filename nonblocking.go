package main

import "fmt"

func main() {
	message := make(chan string) 
	//signal := make(chan bool)

	msg := "HI"
    
	select {
	case message <- msg:
		fmt.Println("received message", msg) 
	default:
		fmt.Println("No message received")
	}
}
