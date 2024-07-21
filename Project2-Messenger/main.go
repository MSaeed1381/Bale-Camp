package main

import "Messenger/server"

func main() {
	serv := server.NewMessengerServer()
	if err := serv.Serve("127.0.0.1:8000"); err != nil {
		return
	}
}
