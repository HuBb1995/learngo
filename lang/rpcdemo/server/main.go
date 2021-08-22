package main

import (
	rpcdemo2 "learngo/lang/rpcdemo"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo2.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
