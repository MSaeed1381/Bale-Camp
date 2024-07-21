package data

import (
	"Messenger/messenger"
	"Messenger/utils"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
)

type User struct {
	userId   int64
	username string
	fileId   string
	Chats    []string
}

func NewUser(username string, fileId string) (*User, error) {
	db := GetDatabaseInstance()

	// check uniqueness
	if _, err := db.GetUserIdByUsername(username); err == nil {
		return nil, status.Error(codes.AlreadyExists, "invalid username")
	}

	// validate username
	if !utils.ValidateUsername(username) {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	// validate profile_id
	if err := utils.ValidateFileId(fileId); err != nil {
		return nil, err
	}

	u := &User{userId: db.GenerateUserId(), fileId: fileId,
		username: username, Chats: make([]string, 0)}

	db.UsernameToId[u.GetUsername()] = u.GetUserId()
	db.Users[u.GetUserId()] = u
	return u, nil
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
	contentType, str := utils.GetContentType(content)
	if contentType == 1 || contentType == 2 { // file or image
		err := utils.ValidateFileId(str)
		if err != nil {
			return nil, err
		}
	}

	m := &messenger.Chat_Message{
		MessageId: db.GenerateMessageId(),
		Content:   content,
		SenderId:  senderId,
		Timestamp: timestamp,
	}

	db.Messages[m.GetMessageId()] = m

	minId := int(math.Min(float64(receiverId), float64(senderId)))
	maxId := int(math.Max(float64(receiverId), float64(senderId)))
	chatId := fmt.Sprintf("%d_%d", minId, maxId)
	chat, ok := db.Chats[chatId]

	if ok {
		chat.Messages = append(chat.Messages, m)
		return m, nil
	}

	newChat := &messenger.Chat{
		ChatId:   db.GenerateChatId(),
		User1:    senderId,
		User2:    receiverId,
		Messages: make([]*messenger.Chat_Message, 0),
	}

	newChat.Messages = append(newChat.Messages, m)
	db.Chats[chatId] = newChat
	db.Users[receiverId].Chats = append(db.Users[receiverId].Chats, chatId)
	db.Users[senderId].Chats = append(db.Users[senderId].Chats, chatId)

	return m, nil
}
