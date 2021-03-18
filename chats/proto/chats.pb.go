// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/chats.proto

package chats

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Chat struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserIds              []string             `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Chat) Reset()         { *m = Chat{} }
func (m *Chat) String() string { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()    {}
func (*Chat) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f41ca1a87b945e7, []int{0}
}

func (m *Chat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chat.Unmarshal(m, b)
}
func (m *Chat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chat.Marshal(b, m, deterministic)
}
func (m *Chat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chat.Merge(m, src)
}
func (m *Chat) XXX_Size() int {
	return xxx_messageInfo_Chat.Size(m)
}
func (m *Chat) XXX_DiscardUnknown() {
	xxx_messageInfo_Chat.DiscardUnknown(m)
}

var xxx_messageInfo_Chat proto.InternalMessageInfo

func (m *Chat) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Chat) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *Chat) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type Message struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string               `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	ChatId               string               `protobuf:"bytes,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Text                 string               `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	SentAt               *timestamp.Timestamp `protobuf:"bytes,5,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f41ca1a87b945e7, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Message) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetSentAt() *timestamp.Timestamp {
	if m != nil {
		return m.SentAt
	}
	return nil
}

type CreateChatRequest struct {
	UserIds              []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateChatRequest) Reset()         { *m = CreateChatRequest{} }
func (m *CreateChatRequest) String() string { return proto.CompactTextString(m) }
func (*CreateChatRequest) ProtoMessage()    {}
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f41ca1a87b945e7, []int{2}
}

func (m *CreateChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateChatRequest.Unmarshal(m, b)
}
func (m *CreateChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateChatRequest.Marshal(b, m, deterministic)
}
func (m *CreateChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateChatRequest.Merge(m, src)
}
func (m *CreateChatRequest) XXX_Size() int {
	return xxx_messageInfo_CreateChatRequest.Size(m)
}
func (m *CreateChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateChatRequest proto.InternalMessageInfo

func (m *CreateChatRequest) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type CreateChatResponse struct {
	Chat                 *Chat    `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateChatResponse) Reset()         { *m = CreateChatResponse{} }
func (m *CreateChatResponse) String() string { return proto.CompactTextString(m) }
func (*CreateChatResponse) ProtoMessage()    {}
func (*CreateChatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f41ca1a87b945e7, []int{3}
}

func (m *CreateChatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateChatResponse.Unmarshal(m, b)
}
func (m *CreateChatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateChatResponse.Marshal(b, m, deterministic)
}
func (m *CreateChatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateChatResponse.Merge(m, src)
}
func (m *CreateChatResponse) XXX_Size() int {
	return xxx_messageInfo_CreateChatResponse.Size(m)
}
func (m *CreateChatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateChatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateChatResponse proto.InternalMessageInfo

func (m *CreateChatResponse) GetChat() *Chat {
	if m != nil {
		return m.Chat
	}
	return nil
}

type CreateMessageRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChatId               string   `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	AuthorId             string   `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageRequest) Reset()         { *m = CreateMessageRequest{} }
func (m *CreateMessageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMessageRequest) ProtoMessage()    {}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f41ca1a87b945e7, []int{4}
}

func (m *CreateMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageRequest.Unmarshal(m, b)
}
func (m *CreateMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageRequest.Marshal(b, m, deterministic)
}
func (m *CreateMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageRequest.Merge(m, src)
}
func (m *CreateMessageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMessageRequest.Size(m)
}
func (m *CreateMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageRequest proto.InternalMessageInfo

func (m *CreateMessageRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateMessageRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *CreateMessageRequest) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *CreateMessageRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CreateMessageResponse struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageResponse) Reset()         { *m = CreateMessageResponse{} }
func (m *CreateMessageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMessageResponse) ProtoMessage()    {}
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f41ca1a87b945e7, []int{5}
}

func (m *CreateMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageResponse.Unmarshal(m, b)
}
func (m *CreateMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageResponse.Marshal(b, m, deterministic)
}
func (m *CreateMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageResponse.Merge(m, src)
}
func (m *CreateMessageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMessageResponse.Size(m)
}
func (m *CreateMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageResponse proto.InternalMessageInfo

func (m *CreateMessageResponse) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type ListMessagesRequest struct {
	ChatId               string               `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	SentBefore           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=sent_before,json=sentBefore,proto3" json:"sent_before,omitempty"`
	Limit                *wrappers.Int32Value `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListMessagesRequest) Reset()         { *m = ListMessagesRequest{} }
func (m *ListMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMessagesRequest) ProtoMessage()    {}
func (*ListMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f41ca1a87b945e7, []int{6}
}

func (m *ListMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessagesRequest.Unmarshal(m, b)
}
func (m *ListMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessagesRequest.Marshal(b, m, deterministic)
}
func (m *ListMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessagesRequest.Merge(m, src)
}
func (m *ListMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMessagesRequest.Size(m)
}
func (m *ListMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessagesRequest proto.InternalMessageInfo

func (m *ListMessagesRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *ListMessagesRequest) GetSentBefore() *timestamp.Timestamp {
	if m != nil {
		return m.SentBefore
	}
	return nil
}

func (m *ListMessagesRequest) GetLimit() *wrappers.Int32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

type ListMessagesResponse struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListMessagesResponse) Reset()         { *m = ListMessagesResponse{} }
func (m *ListMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMessagesResponse) ProtoMessage()    {}
func (*ListMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f41ca1a87b945e7, []int{7}
}

func (m *ListMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessagesResponse.Unmarshal(m, b)
}
func (m *ListMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessagesResponse.Marshal(b, m, deterministic)
}
func (m *ListMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessagesResponse.Merge(m, src)
}
func (m *ListMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMessagesResponse.Size(m)
}
func (m *ListMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessagesResponse proto.InternalMessageInfo

func (m *ListMessagesResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*Chat)(nil), "chats.Chat")
	proto.RegisterType((*Message)(nil), "chats.Message")
	proto.RegisterType((*CreateChatRequest)(nil), "chats.CreateChatRequest")
	proto.RegisterType((*CreateChatResponse)(nil), "chats.CreateChatResponse")
	proto.RegisterType((*CreateMessageRequest)(nil), "chats.CreateMessageRequest")
	proto.RegisterType((*CreateMessageResponse)(nil), "chats.CreateMessageResponse")
	proto.RegisterType((*ListMessagesRequest)(nil), "chats.ListMessagesRequest")
	proto.RegisterType((*ListMessagesResponse)(nil), "chats.ListMessagesResponse")
}

func init() { proto.RegisterFile("proto/chats.proto", fileDescriptor_5f41ca1a87b945e7) }

var fileDescriptor_5f41ca1a87b945e7 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0xd2, 0x9f, 0xb4, 0x13, 0x76, 0xa5, 0x35, 0x8b, 0xc8, 0xa6, 0x88, 0xad, 0x72, 0xaa,
	0x38, 0xa4, 0xa2, 0x15, 0x07, 0xb4, 0xa7, 0x76, 0x0f, 0xa8, 0x08, 0x2e, 0x11, 0xe2, 0xc0, 0x65,
	0xe5, 0x6e, 0xdc, 0x36, 0x52, 0xdb, 0x84, 0x78, 0x22, 0x78, 0x14, 0x4e, 0xbc, 0x1a, 0xaf, 0x82,
	0x3c, 0x76, 0xa0, 0x4e, 0x23, 0xb8, 0xc5, 0xf3, 0x7d, 0xf1, 0x7c, 0x3f, 0x86, 0xab, 0xa2, 0xcc,
	0x31, 0x9f, 0x3e, 0xee, 0x38, 0xca, 0x98, 0xbe, 0x59, 0x8f, 0x0e, 0xe1, 0xed, 0x36, 0xcf, 0xb7,
	0x7b, 0x31, 0xa5, 0xe1, 0xba, 0xda, 0x4c, 0x31, 0x3b, 0x08, 0x89, 0xfc, 0x50, 0x68, 0x5e, 0xf8,
	0xb2, 0x49, 0xf8, 0x56, 0xf2, 0xa2, 0x10, 0xa5, 0xb9, 0x27, 0xda, 0x43, 0xf7, 0x7e, 0xc7, 0x91,
	0x5d, 0x82, 0x9b, 0xa5, 0x81, 0x33, 0x76, 0x26, 0xc3, 0xc4, 0xcd, 0x52, 0x76, 0x03, 0x83, 0x4a,
	0x8a, 0xf2, 0x21, 0x4b, 0x65, 0xe0, 0x8e, 0x3b, 0x93, 0x61, 0xe2, 0xa9, 0xf3, 0x2a, 0x95, 0xec,
	0x2d, 0xc0, 0x63, 0x29, 0x38, 0x8a, 0xf4, 0x81, 0x63, 0xd0, 0x19, 0x3b, 0x13, 0x7f, 0x16, 0xc6,
	0x7a, 0x4f, 0x5c, 0xef, 0x89, 0x3f, 0xd5, 0x42, 0x92, 0xa1, 0x61, 0x2f, 0x30, 0xfa, 0xe1, 0x80,
	0xf7, 0x51, 0x48, 0xc9, 0xb7, 0xe2, 0x6c, 0xe3, 0x08, 0x86, 0xbc, 0xc2, 0x5d, 0xae, 0x76, 0x06,
	0x2e, 0x8d, 0x07, 0x7a, 0xb0, 0x4a, 0xd9, 0x73, 0xf0, 0x94, 0x61, 0x05, 0x75, 0x08, 0xea, 0xab,
	0xe3, 0x2a, 0x65, 0x0c, 0xba, 0x28, 0xbe, 0x63, 0xd0, 0xa5, 0x29, 0x7d, 0xb3, 0x39, 0x78, 0x52,
	0x1c, 0x51, 0xa9, 0xeb, 0xfd, 0x57, 0x5d, 0x5f, 0x51, 0x17, 0x18, 0xc5, 0x70, 0x75, 0x4f, 0x3a,
	0x55, 0x1c, 0x89, 0xf8, 0x5a, 0x09, 0x89, 0x56, 0x0a, 0x8e, 0x95, 0x42, 0xf4, 0x06, 0xd8, 0x29,
	0x5f, 0x16, 0xf9, 0x51, 0x0a, 0x76, 0x0b, 0x5d, 0x25, 0x8c, 0x6c, 0xf9, 0x33, 0x3f, 0xd6, 0x95,
	0x11, 0x85, 0x80, 0xa8, 0x80, 0x6b, 0xfd, 0x9b, 0x89, 0xa1, 0xde, 0xd4, 0x4c, 0xe3, 0xc4, 0xb0,
	0x6b, 0x19, 0xb6, 0x62, 0xea, 0x34, 0x62, 0x6a, 0x49, 0x23, 0x5a, 0xc0, 0xb3, 0xc6, 0x46, 0xa3,
	0x75, 0x02, 0xde, 0x41, 0x8f, 0x8c, 0xdc, 0x4b, 0x23, 0xb7, 0x26, 0xd6, 0x70, 0xf4, 0xd3, 0x81,
	0xa7, 0x1f, 0x32, 0x89, 0x06, 0x90, 0xb5, 0xe8, 0x13, 0x91, 0x8e, 0x25, 0xf2, 0x0e, 0x7c, 0x6a,
	0x60, 0x2d, 0x36, 0x79, 0x29, 0xc8, 0xc1, 0xbf, 0x5b, 0x00, 0x45, 0x5f, 0x12, 0x9b, 0xbd, 0x86,
	0xde, 0x3e, 0x3b, 0x64, 0xf5, 0xd3, 0x1a, 0x9d, 0xfd, 0xb6, 0x3a, 0xe2, 0x7c, 0xf6, 0x99, 0xef,
	0x2b, 0x91, 0x68, 0x66, 0xb4, 0x84, 0x6b, 0x5b, 0x9f, 0xb1, 0xf8, 0x0a, 0x06, 0xc6, 0x83, 0xee,
	0xef, 0xdc, 0xe3, 0x1f, 0x7c, 0xf6, 0xcb, 0x81, 0x9e, 0x2a, 0x4a, 0xb2, 0x05, 0xc0, 0xdf, 0x6a,
	0x59, 0x50, 0x97, 0xd8, 0x7c, 0x1d, 0xe1, 0x4d, 0x0b, 0x62, 0x16, 0xbf, 0x87, 0x0b, 0x2b, 0x74,
	0x36, 0xb2, 0xb8, 0x76, 0xf9, 0xe1, 0x8b, 0x76, 0xd0, 0xdc, 0xf5, 0x0e, 0x9e, 0x9c, 0x9a, 0x63,
	0xa1, 0x61, 0xb7, 0x34, 0x12, 0x8e, 0x5a, 0x31, 0x7d, 0xd1, 0xf2, 0xe2, 0x8b, 0x4f, 0x19, 0xde,
	0x11, 0x67, 0xdd, 0xa7, 0xc3, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x2d, 0x4f, 0xe4,
	0x5e, 0x04, 0x00, 0x00,
}
