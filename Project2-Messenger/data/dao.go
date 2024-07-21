package data

import "Messenger/messenger"

type DatabaseDao interface {
	GenerateUserId() int64
	GenerateMessageId() int64
	GenerateChatId() int64
	GetMessage(messageId int64) (*messenger.Chat_Message, error)
	GetUser(userId int64) (*User, error)
	GetUserIdByUsername(username string) (int64, error)
	GetChatIdByCode(chatCode string) (int64, error)
	GetChat(chatCode string) (*messenger.Chat, error)
}
