package main

import (
	"net/rpc"
	"log"
	"net"
	"sync"
)

type Args struct {
	Message string
}

type MessageService int

var messages []string
var messageLock sync.Mutex

func (t MessageService) AddMessage(args *Args, reply *int) error {
	messageLock.Lock()
	defer messageLock.Unlock()
	log.Println("Added a Message")
	messages = append(messages, args.Message)
	return nil
}

func (t MessageService) GetMessages(args *Args, reply *[]string) error {
	messageLock.Lock()
	defer messageLock.Unlock()
	log.Println("Returning Messages", messages)
	*reply = messages
	messages = []string{}
	return nil
}


func main() {
	arith := new(MessageService)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	log.Printf("MessageService running on port 1234")
	rpc.Accept(l)
}
