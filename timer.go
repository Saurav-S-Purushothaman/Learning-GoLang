package main

import (
    "fmt"
    "time"
)

func main() {
   // initialize a timer
    timer1 := time.NewTimer(time.Second *2)     
    fmt.Println("Before timer starts") // this statement will be executed without waiting for the timer. 

    <-timer1.C // this will wait for the timer
    fmt.Println("After the timer is fired") // this statement will be executed only after the timer is fired.
    timer2 := time.NewTimer(time.Second)
    go func(){
        <-timer2.C
        fmt.Println("Inside goroutine timer2 triggered")
    }()
    time.Sleep(3*time.Second)
    stop:= timer2.Stop()
    if stop{
        fmt.Println("Timer stopped before triggering")
    }
}
