package main

import (
	"net/rpc"
	"log"
	"fmt"
)

type Args struct {
	Vals []int
}
func main() {
	client := make([] *rpc.Client, 2)
	addresses := []string{"localhost:1235",  "localhost:1234"}

	for i := 0; i < 2; i++ {
		c, err := rpc.Dial("tcp",addresses[i])
		client[i] = c
		if err != nil {
			log.Fatal("dialing:", err)
		}
	}
	
	log.Println("Connected to Servers")

	var prices []int
	err := client[0].Call("FlightService.GetPrices", struct{}{}, &prices)
	if err != nil {
		log.Fatal("FlightService.GetPrices error:", err)
	}

	log.Println("Prices:", prices)

	args := Args{Vals:prices}
	var avg float64
	err = client[1].Call("MathService.Average", &args, &avg)
	if err != nil {
		log.Fatal("MathService.Average error:", err)
	}
	fmt.Printf("Average: %.2f\n", avg)
}


