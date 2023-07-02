// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: protobuf/relationship.proto

package api

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

type GetRelationsRequestResponse_RelationStatus int32

const (
	GetRelationsRequestResponse_Relation_PENDING  GetRelationsRequestResponse_RelationStatus = 0
	GetRelationsRequestResponse_Relation_ACCEPTED GetRelationsRequestResponse_RelationStatus = 1
	GetRelationsRequestResponse_Relation_DECLINED GetRelationsRequestResponse_RelationStatus = 2
)

// Enum value maps for GetRelationsRequestResponse_RelationStatus.
var (
	GetRelationsRequestResponse_RelationStatus_name = map[int32]string{
		0: "PENDING",
		1: "ACCEPTED",
		2: "DECLINED",
	}
	GetRelationsRequestResponse_RelationStatus_value = map[string]int32{
		"PENDING":  0,
		"ACCEPTED": 1,
		"DECLINED": 2,
	}
)

func (x GetRelationsRequestResponse_RelationStatus) Enum() *GetRelationsRequestResponse_RelationStatus {
	p := new(GetRelationsRequestResponse_RelationStatus)
	*p = x
	return p
}

func (x GetRelationsRequestResponse_RelationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetRelationsRequestResponse_RelationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_relationship_proto_enumTypes[0].Descriptor()
}

func (GetRelationsRequestResponse_RelationStatus) Type() protoreflect.EnumType {
	return &file_protobuf_relationship_proto_enumTypes[0]
}

func (x GetRelationsRequestResponse_RelationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetRelationsRequestResponse_RelationStatus.Descriptor instead.
func (GetRelationsRequestResponse_RelationStatus) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{6, 0, 0}
}

type CreateRelationRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TargetId int64 `protobuf:"varint,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
}

func (x *CreateRelationRequestRequest) Reset() {
	*x = CreateRelationRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRelationRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationRequestRequest) ProtoMessage() {}

func (x *CreateRelationRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRelationRequestRequest.ProtoReflect.Descriptor instead.
func (*CreateRelationRequestRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRelationRequestRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRelationRequestRequest) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

type CreateRelationRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TargetId int64 `protobuf:"varint,3,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
}

func (x *CreateRelationRequestResponse) Reset() {
	*x = CreateRelationRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRelationRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationRequestResponse) ProtoMessage() {}

func (x *CreateRelationRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRelationRequestResponse.ProtoReflect.Descriptor instead.
func (*CreateRelationRequestResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRelationRequestResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateRelationRequestResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRelationRequestResponse) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

type SubmitRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubmitRelationRequest) Reset() {
	*x = SubmitRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRelationRequest) ProtoMessage() {}

func (x *SubmitRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRelationRequest.ProtoReflect.Descriptor instead.
func (*SubmitRelationRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitRelationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SubmitRelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubmitRelationResponse) Reset() {
	*x = SubmitRelationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRelationResponse) ProtoMessage() {}

func (x *SubmitRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRelationResponse.ProtoReflect.Descriptor instead.
func (*SubmitRelationResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitRelationResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeclineRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeclineRelationRequest) Reset() {
	*x = DeclineRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeclineRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclineRelationRequest) ProtoMessage() {}

func (x *DeclineRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclineRelationRequest.ProtoReflect.Descriptor instead.
func (*DeclineRelationRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{4}
}

func (x *DeclineRelationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeclineRelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeclineRelationResponse) Reset() {
	*x = DeclineRelationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeclineRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclineRelationResponse) ProtoMessage() {}

func (x *DeclineRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclineRelationResponse.ProtoReflect.Descriptor instead.
func (*DeclineRelationResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{5}
}

func (x *DeclineRelationResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRelationsRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relations []*GetRelationsRequestResponse_Relation `protobuf:"bytes,1,rep,name=relations,proto3" json:"relations,omitempty"`
}

func (x *GetRelationsRequestResponse) Reset() {
	*x = GetRelationsRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsRequestResponse) ProtoMessage() {}

func (x *GetRelationsRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsRequestResponse.ProtoReflect.Descriptor instead.
func (*GetRelationsRequestResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{6}
}

func (x *GetRelationsRequestResponse) GetRelations() []*GetRelationsRequestResponse_Relation {
	if x != nil {
		return x.Relations
	}
	return nil
}

type GetRelationsRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetRelationsRequestRequest) Reset() {
	*x = GetRelationsRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsRequestRequest) ProtoMessage() {}

func (x *GetRelationsRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsRequestRequest.ProtoReflect.Descriptor instead.
func (*GetRelationsRequestRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{7}
}

func (x *GetRelationsRequestRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetRelationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRelationsRequest) Reset() {
	*x = GetRelationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsRequest) ProtoMessage() {}

func (x *GetRelationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsRequest.ProtoReflect.Descriptor instead.
func (*GetRelationsRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{8}
}

func (x *GetRelationsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRelationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aliens []*GetRelationsResponse_Aliens `protobuf:"bytes,1,rep,name=aliens,proto3" json:"aliens,omitempty"`
}

func (x *GetRelationsResponse) Reset() {
	*x = GetRelationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsResponse) ProtoMessage() {}

func (x *GetRelationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsResponse.ProtoReflect.Descriptor instead.
func (*GetRelationsResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{9}
}

func (x *GetRelationsResponse) GetAliens() []*GetRelationsResponse_Aliens {
	if x != nil {
		return x.Aliens
	}
	return nil
}

type GetRelationsRequestResponse_Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRelationsRequestResponse_Relation) Reset() {
	*x = GetRelationsRequestResponse_Relation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsRequestResponse_Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsRequestResponse_Relation) ProtoMessage() {}

func (x *GetRelationsRequestResponse_Relation) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsRequestResponse_Relation.ProtoReflect.Descriptor instead.
func (*GetRelationsRequestResponse_Relation) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{6, 0}
}

func (x *GetRelationsRequestResponse_Relation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRelationsRequestResponse_Relation_To struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetRelationsRequestResponse_Relation_To) Reset() {
	*x = GetRelationsRequestResponse_Relation_To{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsRequestResponse_Relation_To) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsRequestResponse_Relation_To) ProtoMessage() {}

func (x *GetRelationsRequestResponse_Relation_To) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsRequestResponse_Relation_To.ProtoReflect.Descriptor instead.
func (*GetRelationsRequestResponse_Relation_To) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{6, 0, 0}
}

func (x *GetRelationsRequestResponse_Relation_To) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRelationsRequestResponse_Relation_To) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetRelationsRequestResponse_Relation_To) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetRelationsResponse_Aliens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetRelationsResponse_Aliens) Reset() {
	*x = GetRelationsResponse_Aliens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_relationship_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationsResponse_Aliens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationsResponse_Aliens) ProtoMessage() {}

func (x *GetRelationsResponse_Aliens) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_relationship_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationsResponse_Aliens.ProtoReflect.Descriptor instead.
func (*GetRelationsResponse_Aliens) Descriptor() ([]byte, []int) {
	return file_protobuf_relationship_proto_rawDescGZIP(), []int{9, 0}
}

func (x *GetRelationsResponse_Aliens) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRelationsResponse_Aliens) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetRelationsResponse_Aliens) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_protobuf_relationship_proto protoreflect.FileDescriptor

var file_protobuf_relationship_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x22, 0x54, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22,
	0x27, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x17,
	0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8a, 0x02, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0xa1, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x52, 0x0a,
	0x02, 0x54, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45,
	0x50, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x43, 0x4c, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x02, 0x22, 0x35, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x61,
	0x6c, 0x69, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x69, 0x65, 0x6e, 0x73, 0x52, 0x06, 0x61,
	0x6c, 0x69, 0x65, 0x6e, 0x73, 0x1a, 0x56, 0x0a, 0x06, 0x41, 0x6c, 0x69, 0x65, 0x6e, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xf9, 0x02,
	0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_relationship_proto_rawDescOnce sync.Once
	file_protobuf_relationship_proto_rawDescData = file_protobuf_relationship_proto_rawDesc
)

func file_protobuf_relationship_proto_rawDescGZIP() []byte {
	file_protobuf_relationship_proto_rawDescOnce.Do(func() {
		file_protobuf_relationship_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_relationship_proto_rawDescData)
	})
	return file_protobuf_relationship_proto_rawDescData
}

var file_protobuf_relationship_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_relationship_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_protobuf_relationship_proto_goTypes = []interface{}{
	(GetRelationsRequestResponse_RelationStatus)(0), // 0: api.GetRelationsRequestResponse.Relation.status
	(*CreateRelationRequestRequest)(nil),            // 1: api.CreateRelationRequestRequest
	(*CreateRelationRequestResponse)(nil),           // 2: api.CreateRelationRequestResponse
	(*SubmitRelationRequest)(nil),                   // 3: api.SubmitRelationRequest
	(*SubmitRelationResponse)(nil),                  // 4: api.SubmitRelationResponse
	(*DeclineRelationRequest)(nil),                  // 5: api.DeclineRelationRequest
	(*DeclineRelationResponse)(nil),                 // 6: api.DeclineRelationResponse
	(*GetRelationsRequestResponse)(nil),             // 7: api.GetRelationsRequestResponse
	(*GetRelationsRequestRequest)(nil),              // 8: api.GetRelationsRequestRequest
	(*GetRelationsRequest)(nil),                     // 9: api.GetRelationsRequest
	(*GetRelationsResponse)(nil),                    // 10: api.GetRelationsResponse
	(*GetRelationsRequestResponse_Relation)(nil),    // 11: api.GetRelationsRequestResponse.Relation
	(*GetRelationsRequestResponse_Relation_To)(nil), // 12: api.GetRelationsRequestResponse.Relation.To
	(*GetRelationsResponse_Aliens)(nil),             // 13: api.GetRelationsResponse.Aliens
}
var file_protobuf_relationship_proto_depIdxs = []int32{
	11, // 0: api.GetRelationsRequestResponse.relations:type_name -> api.GetRelationsRequestResponse.Relation
	13, // 1: api.GetRelationsResponse.aliens:type_name -> api.GetRelationsResponse.Aliens
	1,  // 2: api.RelationService.Create:input_type -> api.CreateRelationRequestRequest
	3,  // 3: api.RelationService.Submit:input_type -> api.SubmitRelationRequest
	5,  // 4: api.RelationService.Decline:input_type -> api.DeclineRelationRequest
	8,  // 5: api.RelationService.GetRequests:input_type -> api.GetRelationsRequestRequest
	9,  // 6: api.RelationService.Get:input_type -> api.GetRelationsRequest
	2,  // 7: api.RelationService.Create:output_type -> api.CreateRelationRequestResponse
	4,  // 8: api.RelationService.Submit:output_type -> api.SubmitRelationResponse
	6,  // 9: api.RelationService.Decline:output_type -> api.DeclineRelationResponse
	7,  // 10: api.RelationService.GetRequests:output_type -> api.GetRelationsRequestResponse
	10, // 11: api.RelationService.Get:output_type -> api.GetRelationsResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_relationship_proto_init() }
func file_protobuf_relationship_proto_init() {
	if File_protobuf_relationship_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_relationship_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRelationRequestRequest); i {
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
		file_protobuf_relationship_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRelationRequestResponse); i {
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
		file_protobuf_relationship_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRelationRequest); i {
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
		file_protobuf_relationship_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRelationResponse); i {
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
		file_protobuf_relationship_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineRelationRequest); i {
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
		file_protobuf_relationship_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineRelationResponse); i {
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
		file_protobuf_relationship_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsRequestResponse); i {
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
		file_protobuf_relationship_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsRequestRequest); i {
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
		file_protobuf_relationship_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsRequest); i {
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
		file_protobuf_relationship_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsResponse); i {
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
		file_protobuf_relationship_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsRequestResponse_Relation); i {
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
		file_protobuf_relationship_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsRequestResponse_Relation_To); i {
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
		file_protobuf_relationship_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationsResponse_Aliens); i {
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
			RawDescriptor: file_protobuf_relationship_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_relationship_proto_goTypes,
		DependencyIndexes: file_protobuf_relationship_proto_depIdxs,
		EnumInfos:         file_protobuf_relationship_proto_enumTypes,
		MessageInfos:      file_protobuf_relationship_proto_msgTypes,
	}.Build()
	File_protobuf_relationship_proto = out.File
	file_protobuf_relationship_proto_rawDesc = nil
	file_protobuf_relationship_proto_goTypes = nil
	file_protobuf_relationship_proto_depIdxs = nil
}
