package main

import (
	"log"
	"net/rpc"

	"github.com/integralist/rpcremote/functions"
)

func main() {
	// Make connection to rpc server
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}

	// Make arguments object
	args := &functions.Args{
		A: 2,
		B: 3,
	}

	// This will store returned result
	var result functions.Result

	// Call remote procedure with args
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		log.Fatalf("Error in Arith/ %s ", err)
	}

	// We got our result in result
	log.Printf("%d*%d=%d\n", args.A, args.B, result)
}
