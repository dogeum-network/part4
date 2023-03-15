package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var method string
	var num1, num2 int
	var reply int

	for {
		fmt.Print("Enter method name (Add/Sub) and numbers: ")
		fmt.Scanln(&method, &num1, &num2)
		args := &Args{num1, num2}
		err = client.Call(method, args, &reply)
		if err != nil {
			log.Fatal("Error:", err)
		}
		fmt.Printf("Result: %d\n", reply)
	}
}
