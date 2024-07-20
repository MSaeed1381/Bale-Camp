package main

import (
	"ConcurrentFileServer/server"
	"fmt"
)

func main() {
	serv := server.NewServer()
	if err := serv.Serve("127.0.0.1:8080"); err != nil {
		fmt.Println("serve err")
	}
}
