package main

import (
	"fmt"
	"sync"
	"time"
)

var ALock sync.Mutex
var A []int

var BLock sync.Mutex
var B []int

func workerA(c chan int) {
	ALock.Lock()
	defer ALock.Unlock()
	time.Sleep(100*time.Millisecond)

	for i := 0; i < 1000; i++ {
		if(A[i] == 0) {
			BLock.Lock()
			A[i] = B[i]
			BLock.Unlock()
		}
	}

	c<-1
	fmt.Println("Worker A done!\n");
}

func workerB() {
	BLock.Lock()
	defer BLock.Unlock()
	time.Sleep(100*time.Millisecond)

	for i := 0; i < 1000; i++ {
		if(B[i] == 0) {
			ALock.Lock()
			B[i] = A[i]
			ALock.Unlock()
		}
	}
	fmt.Println("Worker B done!\n");
}

func main() {
	A = make([]int, 1000)
	B = make([]int, 1000)
	
	for i := 0; i < 1000; i++ { 
		if(i % 3 == 0) {
			B[i] = 312
		}
		if(i % 2 == 0) {
			A[i] = 555
		}
	}

	fmt.Println(A)
	fmt.Println(B)

	fmt.Println("STARTING workers...")
	
	c := make(chan int)
	go workerA(c)

	workerB()

	<-c // indicates that workerA is done
}
