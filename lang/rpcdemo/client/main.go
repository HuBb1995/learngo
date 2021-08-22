package main

import (
	"fmt"
	rpcdemo2 "learngo/lang/rpcdemo"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcdemo2.Args{A: 10, B: 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	err = client.Call("DemoService.Div", rpcdemo2.Args{A: 10, B: 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
