
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


	for {
		select {
		case <-quit:
			log.Println("Quitting!")
			return
		case t := <-time.After(200 * time.Millisecond):
			log.Println("Tick at", t)
		case <- msg_ready:
			log.Println("message ready!")
		}
	}
}
