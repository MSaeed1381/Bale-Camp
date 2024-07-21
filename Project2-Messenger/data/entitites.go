package data

import (
	"Messenger/messenger"
	"Messenger/utils"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	userId   int64
	username string
	fileId   string
	Chats    []string
}

func NewUser(username string, fileId string) *User {
	db := GetDatabaseInstance()
	u := &User{userId: db.GenerateUserId(), fileId: fileId,
		username: username, Chats: make([]string, 0)}

	db.UsernameToId[u.username] = u.userId
	db.Users[u.userId] = u
	return u
}

func (u *User) GetUserId() int64 {
	return u.userId
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetFileId() string {
	return u.fileId
}

func (u *User) GetChats() []string {
	return u.Chats
}

func NewMessage(content *messenger.Chat_Message_Content, senderId int64, receiverId int64, timestamp *timestamppb.Timestamp) (*messenger.Chat_Message, error) {
	db := GetDatabaseInstance()
	contentType, str := utils.ContentType(content)
	if contentType == "file" || contentType == "image" {
		ok, err := utils.ValidateFileId(str)
		if !ok {
			return nil, err
		}
	}
	m := &messenger.Chat_Message{
		MessageId:  db.GenerateMessageId(),
		Content:    content,
		SenderId:   senderId,
		ReceiverId: receiverId,
		Timestamp:  timestamp,
	}

	db.Messages[m.MessageId] = m

	chatId := fmt.Sprintf("%d_%d", m.ReceiverId, m.SenderId)
	if senderId < receiverId {
		chatId = fmt.Sprintf("%d_%d", m.SenderId, m.ReceiverId)
	}
	chat, ok := db.Chats[chatId]

	if ok {
		chat.Messages = append(chat.Messages, m)
		return m, nil
	}

	newChat := &messenger.Chat{
		ChatId:   db.GenerateChatId(),
		User1:    receiverId,
		User2:    senderId,
		Messages: make([]*messenger.Chat_Message, 0),
	}

	newChat.Messages = append(newChat.Messages, m)
	db.Chats[chatId] = newChat
	db.Users[receiverId].Chats = append(db.Users[receiverId].Chats, chatId)
	db.Users[senderId].Chats = append(db.Users[senderId].Chats, chatId)

	return m, nil
}
