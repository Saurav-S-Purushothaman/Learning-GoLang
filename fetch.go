package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()

	// create a response channel that buffers
	responseChannel := make(chan any,2) 
	wg := &sync.WaitGroup{}
	
	wg.Add(2)

	// changeed to goroutine. Therefore cannot assing to loal variable. Lets create a channel
	go fetchUserLikes(userName, responseChannel,wg)
	go fetchUserMatch(userName, responseChannel,wg)
 

	wg.Wait() // block until we have two wg done(). 
	// close the responseChannel 

	close(responseChannel)
	// fmt.Println("username:", userName)
	// fmt.Println("likes:", likes)

	for resp := range responseChannel{
        fmt.Println("response:",resp)
	}


	fmt.Println()
	fmt.Println("took us:", time.Since(start))

}


func fetchUser() string {
    time.Sleep(time.Millisecond *200)
	return "Bob"
} 

func fetchUserLikes(userName string, responseChannel chan any, wg *sync.WaitGroup)  {
    time.Sleep(time.Millisecond *500)
	responseChannel <- 11 
	wg.Done()
}

func fetchUserMatch(userName string, responseChannel chan any, wg *sync.WaitGroup) {
    time.Sleep(time.Millisecond *499)
	responseChannel <- "Anna"
	wg.Done()
}

