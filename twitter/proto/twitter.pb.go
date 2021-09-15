// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/twitter.proto

package twitter

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tweet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the tweet
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// text of the tweet
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// username of the person who tweeted
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// time of tweet
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// number of times retweeted
	RetweetedCount int64 `protobuf:"varint,5,opt,name=retweeted_count,json=retweetedCount,proto3" json:"retweeted_count,omitempty"`
	// number of times favourited
	FavouritedCount int64 `protobuf:"varint,6,opt,name=favourited_count,json=favouritedCount,proto3" json:"favourited_count,omitempty"`
}

func (x *Tweet) Reset() {
	*x = Tweet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twitter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tweet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tweet) ProtoMessage() {}

func (x *Tweet) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twitter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tweet.ProtoReflect.Descriptor instead.
func (*Tweet) Descriptor() ([]byte, []int) {
	return file_proto_twitter_proto_rawDescGZIP(), []int{0}
}

func (x *Tweet) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tweet) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Tweet) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Tweet) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Tweet) GetRetweetedCount() int64 {
	if x != nil {
		return x.RetweetedCount
	}
	return 0
}

func (x *Tweet) GetFavouritedCount() int64 {
	if x != nil {
		return x.FavouritedCount
	}
	return 0
}

// Get the timeline for a given user
type TimelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the username to request the timeline for
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// number of tweets to return. default: 20
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *TimelineRequest) Reset() {
	*x = TimelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twitter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimelineRequest) ProtoMessage() {}

func (x *TimelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twitter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimelineRequest.ProtoReflect.Descriptor instead.
func (*TimelineRequest) Descriptor() ([]byte, []int) {
	return file_proto_twitter_proto_rawDescGZIP(), []int{1}
}

func (x *TimelineRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TimelineRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type TimelineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The recent tweets for the user
	Tweets []*Tweet `protobuf:"bytes,1,rep,name=tweets,proto3" json:"tweets,omitempty"`
}

func (x *TimelineResponse) Reset() {
	*x = TimelineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twitter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimelineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimelineResponse) ProtoMessage() {}

func (x *TimelineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twitter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimelineResponse.ProtoReflect.Descriptor instead.
func (*TimelineResponse) Descriptor() ([]byte, []int) {
	return file_proto_twitter_proto_rawDescGZIP(), []int{2}
}

func (x *TimelineResponse) GetTweets() []*Tweet {
	if x != nil {
		return x.Tweets
	}
	return nil
}

// Search for tweets with a simple query
type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the query to search for
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// number of tweets to return. default: 20
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twitter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twitter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_twitter_proto_rawDescGZIP(), []int{3}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the related tweets for the search
	Tweets []*Tweet `protobuf:"bytes,2,rep,name=tweets,proto3" json:"tweets,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twitter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twitter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_proto_twitter_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResponse) GetTweets() []*Tweet {
	if x != nil {
		return x.Tweets
	}
	return nil
}

var File_proto_twitter_proto protoreflect.FileDescriptor

var file_proto_twitter_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x22, 0xba,
	0x01, 0x0a, 0x05, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x74, 0x77, 0x65,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x72, 0x65, 0x74, 0x77, 0x65, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x69, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x0f, 0x54,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x3a, 0x0a, 0x10, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x77, 0x65, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x52, 0x06, 0x74, 0x77, 0x65, 0x65, 0x74, 0x73, 0x22, 0x3b, 0x0a, 0x0d,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x74,
	0x77, 0x65, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x77,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x06, 0x74, 0x77, 0x65,
	0x65, 0x74, 0x73, 0x32, 0x89, 0x01, 0x0a, 0x07, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12,
	0x41, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x2e, 0x74, 0x77,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x74,
	0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x74, 0x77, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_twitter_proto_rawDescOnce sync.Once
	file_proto_twitter_proto_rawDescData = file_proto_twitter_proto_rawDesc
)

func file_proto_twitter_proto_rawDescGZIP() []byte {
	file_proto_twitter_proto_rawDescOnce.Do(func() {
		file_proto_twitter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_twitter_proto_rawDescData)
	})
	return file_proto_twitter_proto_rawDescData
}

var file_proto_twitter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_twitter_proto_goTypes = []interface{}{
	(*Tweet)(nil),            // 0: twitter.Tweet
	(*TimelineRequest)(nil),  // 1: twitter.TimelineRequest
	(*TimelineResponse)(nil), // 2: twitter.TimelineResponse
	(*SearchRequest)(nil),    // 3: twitter.SearchRequest
	(*SearchResponse)(nil),   // 4: twitter.SearchResponse
}
var file_proto_twitter_proto_depIdxs = []int32{
	0, // 0: twitter.TimelineResponse.tweets:type_name -> twitter.Tweet
	0, // 1: twitter.SearchResponse.tweets:type_name -> twitter.Tweet
	1, // 2: twitter.Twitter.Timeline:input_type -> twitter.TimelineRequest
	3, // 3: twitter.Twitter.Search:input_type -> twitter.SearchRequest
	2, // 4: twitter.Twitter.Timeline:output_type -> twitter.TimelineResponse
	4, // 5: twitter.Twitter.Search:output_type -> twitter.SearchResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_twitter_proto_init() }
func file_proto_twitter_proto_init() {
	if File_proto_twitter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_twitter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tweet); i {
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
		file_proto_twitter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimelineRequest); i {
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
		file_proto_twitter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimelineResponse); i {
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
		file_proto_twitter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_proto_twitter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_twitter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_twitter_proto_goTypes,
		DependencyIndexes: file_proto_twitter_proto_depIdxs,
		MessageInfos:      file_proto_twitter_proto_msgTypes,
	}.Build()
	File_proto_twitter_proto = out.File
	file_proto_twitter_proto_rawDesc = nil
	file_proto_twitter_proto_goTypes = nil
	file_proto_twitter_proto_depIdxs = nil
}
