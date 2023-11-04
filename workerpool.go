package main

import (
	"fmt"
	"net/http"
	"log"
)

type Site struct{
	URL string
}

type Result struct{
	Status int
}

func crawl(wId int, jobs <-chan Site, results chan <-Result){
	for site:= range jobs{
		log.Printf("Worker ID: %d\n", wId)
		resp,err := http.Get(site.URL)
		if err != nil {
			log.Printf(err.Error())
		}
		results <- Result{Status : resp.StatusCode}
	}
}


func main(){
	fmt.Println("Worker pool in go")

	jobs := make(chan Site, 3)
    results := make(chan Result, 3)

	for w:=1; w<=3; w++ {
		go crawl(w,jobs,results)
	}
    
	urls :=[]string{
		"https://www.google.com",
		"https://google.com",
		"https://yahoo.com",
	}
	for _, url := range urls{
		jobs <-Site{URL:url}
	}

    
	for a:=1; a<=3; a++{
		result := <-results
		log.Println(result)
	}

	// 

}