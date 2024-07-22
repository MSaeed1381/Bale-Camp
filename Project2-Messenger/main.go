package main

import "Messenger/server"

func main() {
	serv := server.NewMessengerServer()
	const ADDRESS = "0.0.0.0:8000"
	if err := serv.Serve(ADDRESS); err != nil {
		return
	}
}
