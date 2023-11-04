package main

import (
	"fmt"
	"time"
)

func main(){
	ticker := time.NewTicker(500*time.Millisecond)
	check := make(chan bool)
	
	go func(){
		for {
			select {
			case <- check:
				return
			case t:=<-ticker.C:
				fmt.Println("Ticket at ", t)
			}

		}
	}()
	time.Sleep(1600* time.Millisecond)
	check<-true
	fmt.Println("Ticker stopped")

}
