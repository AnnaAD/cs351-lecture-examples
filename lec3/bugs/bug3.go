package main

import (
	"log"
	"time"
)

func do_processing(batch int) {
	log.Println("Processing: ", batch)
}

func main() {
	batch := 0
	for i := 0; i< 100; i+= 1 {
		batch += 1
		go func() {
			do_processing(batch)
		}()
	}

	time.Sleep(10 * time.Second)
}
