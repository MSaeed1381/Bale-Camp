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
	"sort"
	"time"
)

func (s MessengerServer) AddUser(c context.Context, r *messenger.AddUserRequest) (*messenger.AddUserResponse, error) {
	// check uniqueness
	if _, err := s.data.GetUserIdByUsername(r.GetUsername()); err == nil {
		return nil, status.Error(codes.AlreadyExists, "invalid username")
	}

	// validate username
	if !utils.ValidateUsername(r.GetUsername()) {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	// validate profile_id
	err := utils.ValidateFileId(r.GetFileId())
	if err != nil {
		return nil, err
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

	_, err := data.GetDatabaseInstance().GetUser(senderId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "sender user does not exist")
	}

	var receiverId int64 = 0
	switch (r.GetReceiver()).(type) {
	case *messenger.SendMessageRequest_ReceiverId:
		receiverId = r.GetReceiverId()
	case *messenger.SendMessageRequest_ReceiverUsername:
		receiverId = s.data.UsernameToId[r.GetReceiverUsername()]
	}

	_, err = data.GetDatabaseInstance().GetUser(receiverId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "receiver user does not exist")
	}

	r.GetContent().GetContent()

	message, err := data.NewMessage(r.GetContent(), senderId, receiverId, timestamppb.New(time.Now()))
	if err != nil {
		return nil, err
	}

	fmt.Println(message)
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

	chats := make([]*messenger.Chat, 0)
	for _, chatCode := range user.Chats {
		chat, err := s.data.GetChat(chatCode)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	for _, chat := range chats {
		sort.Slice(chat.Messages, func(i, j int) bool {
			if chat.Messages[i].Timestamp.GetSeconds() < chat.Messages[j].Timestamp.GetSeconds() {
				return true
			} else if chat.Messages[i].Timestamp.Seconds == chat.Messages[j].Timestamp.Seconds &&
				chat.Messages[i].Timestamp.GetNanos() == chat.Messages[j].Timestamp.GetNanos() {
				return true
			}
			return false
		})
	}

	return &messenger.GetUserMessagesResponse{
		Chats: chats,
	}, nil
}
