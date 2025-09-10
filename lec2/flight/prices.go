package main

import (
	"net/rpc"
	"log"
	"net"
)

type Args struct {
	
}

type FlightService int

func (t *FlightService) GetPrices(args *Args, reply *[]int) error {
	log.Println("Called GetPrices")
	*reply = []int{1,2,3}
	return nil
}


func main() {
	arith := new(FlightService)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	log.Printf("FlightService running on port 1235")
	//http.Serve(l, nil)
	rpc.Accept(l)
}
