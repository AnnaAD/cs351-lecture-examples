package main

import (
	"fmt"
 	"time"
	"sync"
)

var sumLock sync.Mutex
var sum int

func worker() {
	for i := 0; i < 100; i++ {
		sumLock.Lock()
		new_val := sum + 1
		sum = new_val
		sumLock.Unlock()
	}
}

func main() {
	sum = 0
	for i:=0; i < 100; i++ {
		go worker()
	}
	time.Sleep(1 * time.Second)
	sumLock.Lock()
	fmt.Printf("sum: %v\n", sum)
	sumLock.Unlock()
}
