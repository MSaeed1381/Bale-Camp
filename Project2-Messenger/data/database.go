package data

import (
	"Messenger/messenger"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Database struct {
	Users        map[int64]*User
	Messages     map[int64]*messenger.Chat_Message
	Chats        map[string]*messenger.Chat
	UsernameToId map[string]int64
	ChatCodeToId map[string]int64

	UserIdsCounter  int64
	MessagesCounter int64
	ChatsCounter    int64
}

// TODO wrap function
var dbInstance *Database

func GetDatabaseInstance() *Database {
	if dbInstance != nil {
		return dbInstance
	}

	dbInstance = &Database{
		Users:        make(map[int64]*User),
		Messages:     make(map[int64]*messenger.Chat_Message),
		Chats:        make(map[string]*messenger.Chat),
		UsernameToId: make(map[string]int64),
		ChatCodeToId: make(map[string]int64),

		UserIdsCounter:  0,
		MessagesCounter: 0,
		ChatsCounter:    0,
	}
	return dbInstance
}

func (db *Database) GenerateUserId() int64 {
	db.UserIdsCounter++
	return db.UserIdsCounter
}

func (db *Database) GenerateMessageId() int64 {
	db.MessagesCounter++
	return db.MessagesCounter
}

func (db *Database) GenerateChatId() int64 {
	db.ChatsCounter++
	return db.ChatsCounter
}

func (db *Database) GetMessage(messageId int64) (*messenger.Chat_Message, error) {
	msg, ok := db.Messages[messageId]
	if !ok {
		return nil, errors.New("message not found")
	}
	return msg, nil
}

func (db *Database) GetUser(userId int64) (*User, error) {
	usr, ok := db.Users[userId]
	if !ok {
		return nil, errors.New("user not found")
	}
	return usr, nil
}

func (db *Database) GetUserIdByUsername(username string) (int64, error) {
	userId, ok := db.UsernameToId[username]
	if !ok {
		return 0, errors.New("invalid username")
	}

	return userId, nil
}

func (db *Database) GetChatIdByCode(chatCode string) (int64, error) {
	chatId, ok := db.ChatCodeToId[chatCode]
	if !ok {
		return 0, errors.New("chat not found")
	}
	return chatId, nil
}

func (db *Database) GetChat(chatCode string) (*messenger.Chat, error) {
	chat, ok := db.Chats[chatCode]
	if !ok {
		return nil, errors.New("chat not found")
	}

	return chat, nil
}

func GetUserID(sender interface{}) (int64, error) {
	db := GetDatabaseInstance()

	switch sender := sender.(type) {
	case *messenger.SendMessageRequest_SenderId:
		return sender.SenderId, nil
	case *messenger.SendMessageRequest_SenderUsername:
		if id, err := db.GetUserIdByUsername(sender.SenderUsername); err == nil {
			return id, nil
		}
		return 0, status.Error(codes.NotFound, "sender user does not exist")
	case *messenger.SendMessageRequest_ReceiverId:
		return sender.ReceiverId, nil
	case *messenger.SendMessageRequest_ReceiverUsername:
		if id, err := db.GetUserIdByUsername(sender.ReceiverUsername); err == nil {
			return id, nil
		}
		return 0, status.Error(codes.NotFound, "receiver user does not exist")
	default:
		return 0, status.Error(codes.InvalidArgument, "unknown sender/receiver type")
	}
}
