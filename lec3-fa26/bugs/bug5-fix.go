
package main

import (
	"log"
	"time"
)

func worker(msg_ready chan string) {
	for {
		time.Sleep(100 * time.Millisecond)
		msg_ready <- "hello"
	}

}


func main() {

	msg_ready := make(chan string)
	quit := make(chan bool)

	go worker(msg_ready)

	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-quit:
			log.Println("Quitting!")
			return
		case t := <-ticker.C:
			log.Println("Tick at", t)
		case <- msg_ready:
			log.Println("message ready!")
		}
	}
}
