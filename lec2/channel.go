package main

import (
	"fmt"
	"time"
)


func worker(inChan chan int) {
	fmt.Println("WORKER: 1. starting...");
	time.Sleep(5 * time.Second);
	
	data := <-inChan
	fmt.Println("WORKER: 2. Got data!", data);
	
}

func main() {
	c := make(chan int)

	go worker(c)

	fmt.Println("MAIN: 1. sending data!")
	c<-5

	fmt.Println("MAIN: 2. sent data!")

	time.Sleep(10*time.Second)

}
