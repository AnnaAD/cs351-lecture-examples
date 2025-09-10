package main

import (
	"net/rpc"
	"log"
	"net"
)

type Args struct {
	Vals []int
}

type MathService int

func (t *MathService) Average(args *Args, reply *float64) error {
	log.Println("Average Called")
	sum := 0
	for i := 0; i < len(args.Vals); i++ {
		sum += args.Vals[i]
	}
	*reply = float64(sum) / float64(len(args.Vals))
	return nil
}

func main() {
	arith := new(MathService)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	log.Printf("MathService running on port 1234")
	//http.Serve(l, nil)
	rpc.Accept(l)
}
