package main

import "Messenger/server"

func main() {
	serv := server.NewMessengerServer()
	const ADDRESS = "127.0.0.1:8000"
	if err := serv.Serve(ADDRESS); err != nil {
		return
	}
}
