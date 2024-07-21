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
	u := &User{userId: GetDatabaseInstance().GetUserId(), fileId: fileId,
		username: username, Chats: make([]string, 0)}

	GetDatabaseInstance().UsernameToId[u.username] = u.userId
	GetDatabaseInstance().Users[u.userId] = u
	return u
}

func (u *User) GetUserId() int64 {
	return u.userId
}

func NewMessage(content *messenger.Chat_Message_Content, senderId int64, receiverId int64, timestamp *timestamppb.Timestamp) (*messenger.Chat_Message, error) {
	contentType, str := utils.ContentType(content)
	if contentType == "file" || contentType == "image" {
		ok, err := utils.ValidateFileId(str)
		if !ok {
			return nil, err
		}
	}
	m := &messenger.Chat_Message{
		MessageId:  GetDatabaseInstance().GetMessageId(),
		Content:    content,
		SenderId:   senderId,
		ReceiverId: receiverId,
		Timestamp:  timestamp,
	}

	GetDatabaseInstance().Messages[m.MessageId] = m

	chatId := fmt.Sprintf("%d_%d", m.ReceiverId, m.SenderId)
	if senderId < receiverId {
		chatId = fmt.Sprintf("%d_%d", m.SenderId, m.ReceiverId)
	}
	chat, ok := GetDatabaseInstance().Chats[chatId]

	if ok {
		chat.Messages = append(chat.Messages, m)
		return m, nil
	}

	newChat := &messenger.Chat{
		ChatId:   GetDatabaseInstance().GetChatId(),
		User1:    receiverId,
		User2:    senderId,
		Messages: make([]*messenger.Chat_Message, 0),
	}

	newChat.Messages = append(newChat.Messages, m)
	GetDatabaseInstance().Chats[chatId] = newChat
	GetDatabaseInstance().Users[receiverId].Chats = append(GetDatabaseInstance().Users[receiverId].Chats, chatId)
	GetDatabaseInstance().Users[senderId].Chats = append(GetDatabaseInstance().Users[senderId].Chats, chatId)

	return m, nil
}
