package main

import "fmt"

func main() {
	
	channel := make(chan string,1)
	go func() {
		channel <- "message"
	}()
	channel <- "message1" 
	
    fmt.Println(<-channel)
}
