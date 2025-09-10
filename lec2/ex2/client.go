package main

import (
	"net/rpc"
	"log"
	"fmt"
	"sync"
	"strconv"
)

type Args struct {
	Message string
}

func main() {
	client, err := rpc.Dial("tcp","localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	log.Println("Connected to Servers")

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { // anonymous function
			defer wg.Done() // at end of function, decrement 1 from wg
			args := Args{Message:"Hi from thread: " + strconv.Itoa(i)}
			var res int
			err := client.Call("MessageService.AddMessage", &args, &res)
			if err != nil {
				log.Fatal("MessageService.AddMessage error:", err)
			}
		}()
	}

	
	wg.Wait() // blocks until all goroutines finish

	args := Args{}
	var res []string
	err = client.Call("MessageService.GetMessages", &args, &res)
	if err != nil {
		log.Fatal("MessageService.AddMessage error:", err)
	}

	fmt.Println("Number of Messages:", len(res))
	
	
}


