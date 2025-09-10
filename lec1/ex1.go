package main

import (
	"fmt"
 	"time"
)

var sum int

func worker() {
	for i := 0; i < 1000; i++ {
		sum = sum + 1
	}
}

func main() {
	sum = 0
	for i:=0; i < 100; i++ {
		go worker()
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("sum: %v\n", sum)
}
