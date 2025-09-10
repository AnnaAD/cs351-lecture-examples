package main

import (
	"fmt"
)


func worker(inChan chan int, resChan chan int) {
	data := <-inChan
	data *= 2
	resChan <- data
}

func main() {
	c := make(chan int, 5) 
	rc := make(chan int)

	for i :=0; i < 5; i+= 1 {
		c <- i
	}

	for i := 0; i< 5; i+= 1 {
		go worker(c,rc)
	}

	for i :=0; i < 5; i+= 1 {
		fmt.Println(i)
	}

}
