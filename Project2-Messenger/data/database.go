package data

import (
	"Messenger/messenger"
	"errors"
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

func (db *Database) GetUserId() int64 {
	db.UserIdsCounter++
	return db.UserIdsCounter
}

func (db *Database) GetMessageId() int64 {
	db.MessagesCounter++
	return db.MessagesCounter
}

func (db *Database) GetChatId() int64 {
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
