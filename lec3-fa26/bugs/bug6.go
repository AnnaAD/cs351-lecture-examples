package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	batch := 0
	for i := 0; i < 100; i++ {
		go func(id int, batch int) {
			select {
			case <- ticker.C:
				log.Println("Worker Processes Data!")
			default:
				log.Println("Worker Collects Data!", id, batch)
			}
		}(i, batch)
		batch += 1
	}
	time.Sleep(10 * time.Second)
}