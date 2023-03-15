package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calculator struct{}

type Args struct {
	A, B int
}

func (c *Calculator) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func (c *Calculator) Sub(args *Args, reply *int) error {
	*reply = args.A - args.B
	return nil
}

func main() {
	calculator := new(Calculator)
	rpc.Register(calculator)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
		fmt.Println("new client connected")
	}
}
