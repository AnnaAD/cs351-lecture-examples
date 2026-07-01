package main

import (
	"fmt"
 	"time"
)

func worker(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 100; i++ {
		go worker(i)
	}
}

