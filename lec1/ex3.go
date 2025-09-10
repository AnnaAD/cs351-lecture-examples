package main

import (
	"fmt"
)


func worker(sumChan chan int) {
	sum := 0 
	for i := 0; i < 100; i++ {
		new_val := sum + 1
		sum = new_val
	}
	sumChan <- sum
}

func main() {
	sumChan := make(chan int)
	for i:=0; i < 100; i++ {
		go worker(sumChan)
	}
	sum := 0
	for i:=0; i < 100; i++ {
		sum += <-sumChan 
	}
	fmt.Printf("sum: %v\n", sum)
}
