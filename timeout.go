package main 

import (
    "fmt"
    "time"
)

func main() {
	c1 := make(chan string,1) 
	go func(){
		time.Sleep(2* time.Second)
		c1 <- "result1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(3*time.Second):
		fmt.Println("timeout1")
	}
}
