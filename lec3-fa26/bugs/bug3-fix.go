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
	for i := 0; i< 5; i+= 1 {
		batch += 1
		go func(b int) {
			do_processing(b)
		}(batch)
	}

	time.Sleep(10 * time.Second)
}
