package server

import (
	"Messenger/data"
	"Messenger/messenger"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

type MessengerServer struct {
	messenger.UnimplementedMessengerServer
	data *data.Database
}

func NewMessengerServer() *MessengerServer {
	return &MessengerServer{data: data.GetDatabaseInstance()}
}

func (s MessengerServer) Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.New("failed to listen: " + err.Error())
	}

	grpcServer := grpc.NewServer()
	messenger.RegisterMessengerServer(grpcServer, NewMessengerServer())

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		return errors.New("failed to serve: " + err.Error())
	}

	return nil
}
