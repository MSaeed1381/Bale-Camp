// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: messenger/messenger.proto

package messenger

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId   int64           `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	User1    int64           `protobuf:"varint,2,opt,name=user1,proto3" json:"user1,omitempty"`
	User2    int64           `protobuf:"varint,3,opt,name=user2,proto3" json:"user2,omitempty"`
	Messages []*Chat_Message `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *Chat) GetUser1() int64 {
	if x != nil {
		return x.User1
	}
	return 0
}

func (x *Chat) GetUser2() int64 {
	if x != nil {
		return x.User2
	}
	return 0
}

func (x *Chat) GetMessages() []*Chat_Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type AddUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FileId   string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *AddUserRequest) Reset() {
	*x = AddUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRequest) ProtoMessage() {}

func (x *AddUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRequest.ProtoReflect.Descriptor instead.
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddUserRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type AddUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AddUserResponse) Reset() {
	*x = AddUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserResponse) ProtoMessage() {}

func (x *AddUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserResponse.ProtoReflect.Descriptor instead.
func (*AddUserResponse) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{2}
}

func (x *AddUserResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sender:
	//
	//	*SendMessageRequest_SenderId
	//	*SendMessageRequest_SenderUsername
	Sender isSendMessageRequest_Sender `protobuf_oneof:"sender"`
	// Types that are assignable to Receiver:
	//
	//	*SendMessageRequest_ReceiverId
	//	*SendMessageRequest_ReceiverUsername
	Receiver isSendMessageRequest_Receiver `protobuf_oneof:"receiver"`
	Content  *Chat_Message_Content         `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{3}
}

func (m *SendMessageRequest) GetSender() isSendMessageRequest_Sender {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (x *SendMessageRequest) GetSenderId() int64 {
	if x, ok := x.GetSender().(*SendMessageRequest_SenderId); ok {
		return x.SenderId
	}
	return 0
}

func (x *SendMessageRequest) GetSenderUsername() string {
	if x, ok := x.GetSender().(*SendMessageRequest_SenderUsername); ok {
		return x.SenderUsername
	}
	return ""
}

func (m *SendMessageRequest) GetReceiver() isSendMessageRequest_Receiver {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (x *SendMessageRequest) GetReceiverId() int64 {
	if x, ok := x.GetReceiver().(*SendMessageRequest_ReceiverId); ok {
		return x.ReceiverId
	}
	return 0
}

func (x *SendMessageRequest) GetReceiverUsername() string {
	if x, ok := x.GetReceiver().(*SendMessageRequest_ReceiverUsername); ok {
		return x.ReceiverUsername
	}
	return ""
}

func (x *SendMessageRequest) GetContent() *Chat_Message_Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type isSendMessageRequest_Sender interface {
	isSendMessageRequest_Sender()
}

type SendMessageRequest_SenderId struct {
	SenderId int64 `protobuf:"varint,1,opt,name=sender_id,json=senderId,proto3,oneof"`
}

type SendMessageRequest_SenderUsername struct {
	SenderUsername string `protobuf:"bytes,2,opt,name=sender_username,json=senderUsername,proto3,oneof"`
}

func (*SendMessageRequest_SenderId) isSendMessageRequest_Sender() {}

func (*SendMessageRequest_SenderUsername) isSendMessageRequest_Sender() {}

type isSendMessageRequest_Receiver interface {
	isSendMessageRequest_Receiver()
}

type SendMessageRequest_ReceiverId struct {
	ReceiverId int64 `protobuf:"varint,3,opt,name=receiver_id,json=receiverId,proto3,oneof"`
}

type SendMessageRequest_ReceiverUsername struct {
	ReceiverUsername string `protobuf:"bytes,4,opt,name=receiver_username,json=receiverUsername,proto3,oneof"`
}

func (*SendMessageRequest_ReceiverId) isSendMessageRequest_Receiver() {}

func (*SendMessageRequest_ReceiverUsername) isSendMessageRequest_Receiver() {}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64 `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{4}
}

func (x *SendMessageResponse) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type FetchMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64 `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *FetchMessageRequest) Reset() {
	*x = FetchMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMessageRequest) ProtoMessage() {}

func (x *FetchMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMessageRequest.ProtoReflect.Descriptor instead.
func (*FetchMessageRequest) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{5}
}

func (x *FetchMessageRequest) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type FetchMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Chat_Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FetchMessageResponse) Reset() {
	*x = FetchMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMessageResponse) ProtoMessage() {}

func (x *FetchMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMessageResponse.ProtoReflect.Descriptor instead.
func (*FetchMessageResponse) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{6}
}

func (x *FetchMessageResponse) GetMessage() *Chat_Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type GetUserMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserMessagesRequest) Reset() {
	*x = GetUserMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMessagesRequest) ProtoMessage() {}

func (x *GetUserMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetUserMessagesRequest) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserMessagesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*Chat `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *GetUserMessagesResponse) Reset() {
	*x = GetUserMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMessagesResponse) ProtoMessage() {}

func (x *GetUserMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetUserMessagesResponse) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserMessagesResponse) GetChats() []*Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

type Chat_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64                 `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Content   *Chat_Message_Content `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	SenderId  int64                 `protobuf:"varint,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// int64 receiver_id = 4;
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Chat_Message) Reset() {
	*x = Chat_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat_Message) ProtoMessage() {}

func (x *Chat_Message) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat_Message.ProtoReflect.Descriptor instead.
func (*Chat_Message) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Chat_Message) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *Chat_Message) GetContent() *Chat_Message_Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chat_Message) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *Chat_Message) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Chat_Message_Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*Chat_Message_Content_Text
	//	*Chat_Message_Content_File
	//	*Chat_Message_Content_Image
	Content isChat_Message_Content_Content `protobuf_oneof:"content"`
}

func (x *Chat_Message_Content) Reset() {
	*x = Chat_Message_Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messenger_messenger_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat_Message_Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat_Message_Content) ProtoMessage() {}

func (x *Chat_Message_Content) ProtoReflect() protoreflect.Message {
	mi := &file_messenger_messenger_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat_Message_Content.ProtoReflect.Descriptor instead.
func (*Chat_Message_Content) Descriptor() ([]byte, []int) {
	return file_messenger_messenger_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *Chat_Message_Content) GetContent() isChat_Message_Content_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Chat_Message_Content) GetText() string {
	if x, ok := x.GetContent().(*Chat_Message_Content_Text); ok {
		return x.Text
	}
	return ""
}

func (x *Chat_Message_Content) GetFile() string {
	if x, ok := x.GetContent().(*Chat_Message_Content_File); ok {
		return x.File
	}
	return ""
}

func (x *Chat_Message_Content) GetImage() string {
	if x, ok := x.GetContent().(*Chat_Message_Content_Image); ok {
		return x.Image
	}
	return ""
}

type isChat_Message_Content_Content interface {
	isChat_Message_Content_Content()
}

type Chat_Message_Content_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type Chat_Message_Content_File struct {
	File string `protobuf:"bytes,2,opt,name=file,proto3,oneof"`
}

type Chat_Message_Content_Image struct {
	Image string `protobuf:"bytes,3,opt,name=image,proto3,oneof"`
}

func (*Chat_Message_Content_Text) isChat_Message_Content_Content() {}

func (*Chat_Message_Content_File) isChat_Message_Content_Content() {}

func (*Chat_Message_Content_Image) isChat_Message_Content_Content() {}

var File_messenger_messenger_proto protoreflect.FileDescriptor

var file_messenger_messenger_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x03, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x12,
	0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x32, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x9a, 0x02, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x58, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22, 0x45, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x81, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x08,
	0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x13, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x22, 0x49, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x68, 0x61,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73,
	0x32, 0xce, 0x02, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x42,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3a, 0x5a, 0x38, 0x2e, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x61, 0x65,
	0x65, 0x64, 0x7a, 0x61, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x6c, 0x65, 0x2d, 0x63, 0x61, 0x6d, 0x70,
	0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x2d, 0x4d, 0x65, 0x73, 0x73, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messenger_messenger_proto_rawDescOnce sync.Once
	file_messenger_messenger_proto_rawDescData = file_messenger_messenger_proto_rawDesc
)

func file_messenger_messenger_proto_rawDescGZIP() []byte {
	file_messenger_messenger_proto_rawDescOnce.Do(func() {
		file_messenger_messenger_proto_rawDescData = protoimpl.X.CompressGZIP(file_messenger_messenger_proto_rawDescData)
	})
	return file_messenger_messenger_proto_rawDescData
}

var file_messenger_messenger_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_messenger_messenger_proto_goTypes = []any{
	(*Chat)(nil),                    // 0: messenger.Chat
	(*AddUserRequest)(nil),          // 1: messenger.AddUserRequest
	(*AddUserResponse)(nil),         // 2: messenger.AddUserResponse
	(*SendMessageRequest)(nil),      // 3: messenger.SendMessageRequest
	(*SendMessageResponse)(nil),     // 4: messenger.SendMessageResponse
	(*FetchMessageRequest)(nil),     // 5: messenger.FetchMessageRequest
	(*FetchMessageResponse)(nil),    // 6: messenger.FetchMessageResponse
	(*GetUserMessagesRequest)(nil),  // 7: messenger.GetUserMessagesRequest
	(*GetUserMessagesResponse)(nil), // 8: messenger.GetUserMessagesResponse
	(*Chat_Message)(nil),            // 9: messenger.Chat.Message
	(*Chat_Message_Content)(nil),    // 10: messenger.Chat.Message.Content
	(*timestamppb.Timestamp)(nil),   // 11: google.protobuf.Timestamp
}
var file_messenger_messenger_proto_depIdxs = []int32{
	9,  // 0: messenger.Chat.messages:type_name -> messenger.Chat.Message
	10, // 1: messenger.SendMessageRequest.content:type_name -> messenger.Chat.Message.Content
	9,  // 2: messenger.FetchMessageResponse.message:type_name -> messenger.Chat.Message
	0,  // 3: messenger.GetUserMessagesResponse.chats:type_name -> messenger.Chat
	10, // 4: messenger.Chat.Message.content:type_name -> messenger.Chat.Message.Content
	11, // 5: messenger.Chat.Message.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 6: messenger.messenger.AddUser:input_type -> messenger.AddUserRequest
	3,  // 7: messenger.messenger.SendMessage:input_type -> messenger.SendMessageRequest
	5,  // 8: messenger.messenger.FetchMessage:input_type -> messenger.FetchMessageRequest
	7,  // 9: messenger.messenger.GetUserMessages:input_type -> messenger.GetUserMessagesRequest
	2,  // 10: messenger.messenger.AddUser:output_type -> messenger.AddUserResponse
	4,  // 11: messenger.messenger.SendMessage:output_type -> messenger.SendMessageResponse
	6,  // 12: messenger.messenger.FetchMessage:output_type -> messenger.FetchMessageResponse
	8,  // 13: messenger.messenger.GetUserMessages:output_type -> messenger.GetUserMessagesResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_messenger_messenger_proto_init() }
func file_messenger_messenger_proto_init() {
	if File_messenger_messenger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messenger_messenger_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Chat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SendMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SendMessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*FetchMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*FetchMessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserMessagesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserMessagesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Chat_Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messenger_messenger_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Chat_Message_Content); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_messenger_messenger_proto_msgTypes[3].OneofWrappers = []any{
		(*SendMessageRequest_SenderId)(nil),
		(*SendMessageRequest_SenderUsername)(nil),
		(*SendMessageRequest_ReceiverId)(nil),
		(*SendMessageRequest_ReceiverUsername)(nil),
	}
	file_messenger_messenger_proto_msgTypes[10].OneofWrappers = []any{
		(*Chat_Message_Content_Text)(nil),
		(*Chat_Message_Content_File)(nil),
		(*Chat_Message_Content_Image)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messenger_messenger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messenger_messenger_proto_goTypes,
		DependencyIndexes: file_messenger_messenger_proto_depIdxs,
		MessageInfos:      file_messenger_messenger_proto_msgTypes,
	}.Build()
	File_messenger_messenger_proto = out.File
	file_messenger_messenger_proto_rawDesc = nil
	file_messenger_messenger_proto_goTypes = nil
	file_messenger_messenger_proto_depIdxs = nil
}