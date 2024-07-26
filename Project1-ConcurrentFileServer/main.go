package main

import (
	"ConcurrentFileServer/server"
	"fmt"
)

func main() {
	serv := server.NewFileServer()
	addr := "0.0.0.0:8080"
	if err := serv.Serve(addr); err != nil {
		fmt.Println("serve err")
	}
}
