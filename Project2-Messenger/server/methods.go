package server

import (
	"Messenger/data"
	"Messenger/messenger"
	"Messenger/utils"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (s MessengerServer) AddUser(c context.Context, r *messenger.AddUserRequest) (*messenger.AddUserResponse, error) {
	// check uniqueness
	if _, ok := s.data.UsernameToId[r.GetUsername()]; ok {
		return nil, status.Error(codes.AlreadyExists, "invalid username")
	}

	// validate username
	if !utils.ValidateUsername(r.GetUsername()) {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	// validate profile_id
	ok, err := utils.ValidateFileId(r.GetFileId())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid file_id")
	}

	user := data.NewUser(r.GetUsername(), r.GetFileId())
	return &messenger.AddUserResponse{UserId: user.GetUserId()}, nil
}

func (s MessengerServer) SendMessage(c context.Context, r *messenger.SendMessageRequest) (*messenger.SendMessageResponse, error) {
	var senderId int64 = 0
	switch (r.GetSender()).(type) {
	case *messenger.SendMessageRequest_SenderId:
		senderId = r.GetSenderId()
	case *messenger.SendMessageRequest_SenderUsername:
		senderId = s.data.UsernameToId[r.GetSenderUsername()]
	}

	if senderId == 0 {
		return nil, status.Error(codes.NotFound, "sender user does not exist")
	}

	var receiverId int64 = 0
	switch (r.GetReceiver()).(type) {
	case *messenger.SendMessageRequest_ReceiverId:
		receiverId = r.GetReceiverId()
	case *messenger.SendMessageRequest_ReceiverUsername:
		receiverId = s.data.UsernameToId[r.GetReceiverUsername()]
	}

	if receiverId == 0 {
		return nil, status.Error(codes.NotFound, "receiver user does not exist")
	}

	r.GetContent().GetContent()

	message, _ := data.NewMessage(r.GetContent(), senderId, receiverId, timestamppb.New(time.Now()))
	return &messenger.SendMessageResponse{MessageId: message.MessageId}, nil
}
func (s MessengerServer) FetchMessage(c context.Context, r *messenger.FetchMessageRequest) (*messenger.FetchMessageResponse, error) {
	message, err := s.data.GetMessage(r.GetMessageId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "message does not exist")
	}

	return &messenger.FetchMessageResponse{Message: message}, nil
}

func (s MessengerServer) GetUserMessages(c context.Context, r *messenger.GetUserMessagesRequest) (*messenger.GetUserMessagesResponse, error) {
	user, err := s.data.GetUser(r.GetUserId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "user does not exist")
	}

	fmt.Println(user)
	fmt.Println(user.Chats)
	fmt.Println(data.GetDatabaseInstance().Chats[user.Chats[0]])

	chats := make([]*messenger.Chat, 0)
	for _, chatId := range user.Chats {
		chats = append(chats, data.GetDatabaseInstance().Chats[chatId])
	}

	return &messenger.GetUserMessagesResponse{
		Chats: chats,
	}, nil
}
