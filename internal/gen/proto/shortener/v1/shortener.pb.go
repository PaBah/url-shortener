// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: proto/shortener/v1/shortener.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	reflect "reflect"
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortRequest) Reset() {
	*x = ShortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortRequest) ProtoMessage() {}

func (x *ShortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortRequest.ProtoReflect.Descriptor instead.
func (*ShortRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *ShortRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ShortRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ShortResponse) Reset() {
	*x = ShortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortResponse) ProtoMessage() {}

func (x *ShortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortResponse.ProtoReflect.Descriptor instead.
func (*ShortResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *ShortResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ExpandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId string `protobuf:"bytes,1,opt,name=short_id,json=shortId,proto3" json:"short_id,omitempty"`
}

func (x *ExpandRequest) Reset() {
	*x = ExpandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandRequest) ProtoMessage() {}

func (x *ExpandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandRequest.ProtoReflect.Descriptor instead.
func (*ExpandRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *ExpandRequest) GetShortId() string {
	if x != nil {
		return x.ShortId
	}
	return ""
}

type ExpandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ExpandResponse) Reset() {
	*x = ExpandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandResponse) ProtoMessage() {}

func (x *ExpandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandResponse.ProtoReflect.Descriptor instead.
func (*ExpandResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *ExpandResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     []string `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{5}
}

type GetUserBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserBucketRequest) Reset() {
	*x = GetUserBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBucketRequest) ProtoMessage() {}

func (x *GetUserBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBucketRequest.ProtoReflect.Descriptor instead.
func (*GetUserBucketRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserBucketRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type OriginalAndShort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *OriginalAndShort) Reset() {
	*x = OriginalAndShort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginalAndShort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalAndShort) ProtoMessage() {}

func (x *OriginalAndShort) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalAndShort.ProtoReflect.Descriptor instead.
func (*OriginalAndShort) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *OriginalAndShort) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *OriginalAndShort) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type GetUserBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*OriginalAndShort `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserBucketResponse) Reset() {
	*x = GetUserBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBucketResponse) ProtoMessage() {}

func (x *GetUserBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBucketResponse.ProtoReflect.Descriptor instead.
func (*GetUserBucketResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserBucketResponse) GetData() []*OriginalAndShort {
	if x != nil {
		return x.Data
	}
	return nil
}

type CorrelatedOriginalURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *CorrelatedOriginalURL) Reset() {
	*x = CorrelatedOriginalURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorrelatedOriginalURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorrelatedOriginalURL) ProtoMessage() {}

func (x *CorrelatedOriginalURL) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorrelatedOriginalURL.ProtoReflect.Descriptor instead.
func (*CorrelatedOriginalURL) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *CorrelatedOriginalURL) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CorrelatedOriginalURL) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type ShortBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string                   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Original []*CorrelatedOriginalURL `protobuf:"bytes,2,rep,name=original,proto3" json:"original,omitempty"`
}

func (x *ShortBatchRequest) Reset() {
	*x = ShortBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortBatchRequest) ProtoMessage() {}

func (x *ShortBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortBatchRequest.ProtoReflect.Descriptor instead.
func (*ShortBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *ShortBatchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ShortBatchRequest) GetOriginal() []*CorrelatedOriginalURL {
	if x != nil {
		return x.Original
	}
	return nil
}

type CorrelatedShortURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl      string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *CorrelatedShortURL) Reset() {
	*x = CorrelatedShortURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorrelatedShortURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorrelatedShortURL) ProtoMessage() {}

func (x *CorrelatedShortURL) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorrelatedShortURL.ProtoReflect.Descriptor instead.
func (*CorrelatedShortURL) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{11}
}

func (x *CorrelatedShortURL) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CorrelatedShortURL) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ShortBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Short []*CorrelatedShortURL `protobuf:"bytes,1,rep,name=short,proto3" json:"short,omitempty"`
}

func (x *ShortBatchResponse) Reset() {
	*x = ShortBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortBatchResponse) ProtoMessage() {}

func (x *ShortBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortBatchResponse.ProtoReflect.Descriptor instead.
func (*ShortBatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{12}
}

func (x *ShortBatchResponse) GetShort() []*CorrelatedShortURL {
	if x != nil {
		return x.Short
	}
	return nil
}

type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{13}
}

type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls  int64 `protobuf:"varint,1,opt,name=urls,proto3" json:"urls,omitempty"`
	Users int64 `protobuf:"varint,2,opt,name=users,proto3" json:"users,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_v1_shortener_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_v1_shortener_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_v1_shortener_proto_rawDescGZIP(), []int{14}
}

func (x *StatsResponse) GetUrls() int64 {
	if x != nil {
		return x.Urls
	}
	return 0
}

func (x *StatsResponse) GetUsers() int64 {
	if x != nil {
		return x.Users
	}
	return 0
}

var File_proto_shortener_v1_shortener_proto protoreflect.FileDescriptor

var file_proto_shortener_v1_shortener_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x34, 0x0a,
	0x0d, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x08, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x51, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0x92, 0x01, 0x07, 0x22,
	0x05, 0x72, 0x03, 0x98, 0x01, 0x08, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x10, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x41, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x51, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x41, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x74,
	0x0a, 0x15, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x2e, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x72, 0x6c, 0x22, 0x73, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52,
	0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x58, 0x0a, 0x12, 0x43, 0x6f, 0x72,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x22, 0x52, 0x0a, 0x12, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x32, 0x93, 0x04, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12,
	0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a,
	0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x25, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x42, 0x61, 0x68, 0x2f, 0x75, 0x72, 0x6c,
	0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_shortener_v1_shortener_proto_rawDescOnce sync.Once
	file_proto_shortener_v1_shortener_proto_rawDescData = file_proto_shortener_v1_shortener_proto_rawDesc
)

func file_proto_shortener_v1_shortener_proto_rawDescGZIP() []byte {
	file_proto_shortener_v1_shortener_proto_rawDescOnce.Do(func() {
		file_proto_shortener_v1_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shortener_v1_shortener_proto_rawDescData)
	})
	return file_proto_shortener_v1_shortener_proto_rawDescData
}

var file_proto_shortener_v1_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_shortener_v1_shortener_proto_goTypes = []any{
	(*ShortRequest)(nil),          // 0: proto.shortener.v1.ShortRequest
	(*ShortResponse)(nil),         // 1: proto.shortener.v1.ShortResponse
	(*ExpandRequest)(nil),         // 2: proto.shortener.v1.ExpandRequest
	(*ExpandResponse)(nil),        // 3: proto.shortener.v1.ExpandResponse
	(*DeleteRequest)(nil),         // 4: proto.shortener.v1.DeleteRequest
	(*DeleteResponse)(nil),        // 5: proto.shortener.v1.DeleteResponse
	(*GetUserBucketRequest)(nil),  // 6: proto.shortener.v1.GetUserBucketRequest
	(*OriginalAndShort)(nil),      // 7: proto.shortener.v1.OriginalAndShort
	(*GetUserBucketResponse)(nil), // 8: proto.shortener.v1.GetUserBucketResponse
	(*CorrelatedOriginalURL)(nil), // 9: proto.shortener.v1.CorrelatedOriginalURL
	(*ShortBatchRequest)(nil),     // 10: proto.shortener.v1.ShortBatchRequest
	(*CorrelatedShortURL)(nil),    // 11: proto.shortener.v1.CorrelatedShortURL
	(*ShortBatchResponse)(nil),    // 12: proto.shortener.v1.ShortBatchResponse
	(*StatsRequest)(nil),          // 13: proto.shortener.v1.StatsRequest
	(*StatsResponse)(nil),         // 14: proto.shortener.v1.StatsResponse
}
var file_proto_shortener_v1_shortener_proto_depIdxs = []int32{
	7,  // 0: proto.shortener.v1.GetUserBucketResponse.data:type_name -> proto.shortener.v1.OriginalAndShort
	9,  // 1: proto.shortener.v1.ShortBatchRequest.original:type_name -> proto.shortener.v1.CorrelatedOriginalURL
	11, // 2: proto.shortener.v1.ShortBatchResponse.short:type_name -> proto.shortener.v1.CorrelatedShortURL
	0,  // 3: proto.shortener.v1.ShortenerService.Short:input_type -> proto.shortener.v1.ShortRequest
	2,  // 4: proto.shortener.v1.ShortenerService.Expand:input_type -> proto.shortener.v1.ExpandRequest
	4,  // 5: proto.shortener.v1.ShortenerService.Delete:input_type -> proto.shortener.v1.DeleteRequest
	6,  // 6: proto.shortener.v1.ShortenerService.GetUserBucket:input_type -> proto.shortener.v1.GetUserBucketRequest
	10, // 7: proto.shortener.v1.ShortenerService.ShortBatch:input_type -> proto.shortener.v1.ShortBatchRequest
	13, // 8: proto.shortener.v1.ShortenerService.Stats:input_type -> proto.shortener.v1.StatsRequest
	1,  // 9: proto.shortener.v1.ShortenerService.Short:output_type -> proto.shortener.v1.ShortResponse
	3,  // 10: proto.shortener.v1.ShortenerService.Expand:output_type -> proto.shortener.v1.ExpandResponse
	5,  // 11: proto.shortener.v1.ShortenerService.Delete:output_type -> proto.shortener.v1.DeleteResponse
	8,  // 12: proto.shortener.v1.ShortenerService.GetUserBucket:output_type -> proto.shortener.v1.GetUserBucketResponse
	12, // 13: proto.shortener.v1.ShortenerService.ShortBatch:output_type -> proto.shortener.v1.ShortBatchResponse
	14, // 14: proto.shortener.v1.ShortenerService.Stats:output_type -> proto.shortener.v1.StatsResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_shortener_v1_shortener_proto_init() }
func file_proto_shortener_v1_shortener_proto_init() {
	if File_proto_shortener_v1_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shortener_v1_shortener_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ShortRequest); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ShortResponse); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ExpandRequest); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ExpandResponse); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteResponse); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserBucketRequest); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*OriginalAndShort); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserBucketResponse); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CorrelatedOriginalURL); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ShortBatchRequest); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*CorrelatedShortURL); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*ShortBatchResponse); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*StatsRequest); i {
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
		file_proto_shortener_v1_shortener_proto_msgTypes[14].Exporter = func(v any, i int) any {
			switch v := v.(*StatsResponse); i {
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
			RawDescriptor: file_proto_shortener_v1_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shortener_v1_shortener_proto_goTypes,
		DependencyIndexes: file_proto_shortener_v1_shortener_proto_depIdxs,
		MessageInfos:      file_proto_shortener_v1_shortener_proto_msgTypes,
	}.Build()
	File_proto_shortener_v1_shortener_proto = out.File
	file_proto_shortener_v1_shortener_proto_rawDesc = nil
	file_proto_shortener_v1_shortener_proto_goTypes = nil
	file_proto_shortener_v1_shortener_proto_depIdxs = nil
}
