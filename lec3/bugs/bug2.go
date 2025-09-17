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
	c := make(chan int)
	rc := make(chan int)

	fmt.Println("writing data!")
	for i:=0; i < 5; i+= 1 {
		c <- i
	}
	fmt.Println("COMPLETE: writing data!")


	for i := 0; i< 5; i+= 1 {
		go worker(c,rc)
	}

	for i := 0; i< 5; i+= 1 {
		fmt.Println(i)
	}

}
