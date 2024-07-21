package server

import (
	"Messenger/data"
	"Messenger/messenger"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sort"
	"time"
)

func (s MessengerServer) AddUser(c context.Context, r *messenger.AddUserRequest) (*messenger.AddUserResponse, error) {
	user, err := data.NewUser(r.GetUsername(), r.GetFileId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &messenger.AddUserResponse{UserId: user.GetUserId()}, nil
}

func (s MessengerServer) SendMessage(c context.Context, r *messenger.SendMessageRequest) (*messenger.SendMessageResponse, error) {
	db := data.GetDatabaseInstance()

	senderId, err := data.GetUserID(r.GetSender())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	receiverId, err := data.GetUserID(r.GetReceiver())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err = db.GetUser(senderId); err != nil {
		return nil, status.Error(codes.NotFound, "sender user does not exist")
	}

	if _, err = db.GetUser(receiverId); err != nil {
		return nil, status.Error(codes.NotFound, "receiver user does not exist")
	}

	message, err := data.NewMessage(r.GetContent(), senderId, receiverId, timestamppb.New(time.Now()))
	if err != nil {
		return nil, err
	}

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

	// TODO sorting with id
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
