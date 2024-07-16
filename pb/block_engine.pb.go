// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: block_engine.proto

package jito_pb

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

type SubscribePacketsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribePacketsRequest) Reset() {
	*x = SubscribePacketsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribePacketsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribePacketsRequest) ProtoMessage() {}

func (x *SubscribePacketsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribePacketsRequest.ProtoReflect.Descriptor instead.
func (*SubscribePacketsRequest) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{0}
}

type SubscribePacketsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Batch  *PacketBatch `protobuf:"bytes,2,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (x *SubscribePacketsResponse) Reset() {
	*x = SubscribePacketsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribePacketsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribePacketsResponse) ProtoMessage() {}

func (x *SubscribePacketsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribePacketsResponse.ProtoReflect.Descriptor instead.
func (*SubscribePacketsResponse) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribePacketsResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SubscribePacketsResponse) GetBatch() *PacketBatch {
	if x != nil {
		return x.Batch
	}
	return nil
}

type SubscribeBundlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeBundlesRequest) Reset() {
	*x = SubscribeBundlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeBundlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeBundlesRequest) ProtoMessage() {}

func (x *SubscribeBundlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeBundlesRequest.ProtoReflect.Descriptor instead.
func (*SubscribeBundlesRequest) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{2}
}

type SubscribeBundlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bundles []*BundleUuid `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles,omitempty"`
}

func (x *SubscribeBundlesResponse) Reset() {
	*x = SubscribeBundlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeBundlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeBundlesResponse) ProtoMessage() {}

func (x *SubscribeBundlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeBundlesResponse.ProtoReflect.Descriptor instead.
func (*SubscribeBundlesResponse) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeBundlesResponse) GetBundles() []*BundleUuid {
	if x != nil {
		return x.Bundles
	}
	return nil
}

type BlockBuilderFeeInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlockBuilderFeeInfoRequest) Reset() {
	*x = BlockBuilderFeeInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockBuilderFeeInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockBuilderFeeInfoRequest) ProtoMessage() {}

func (x *BlockBuilderFeeInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockBuilderFeeInfoRequest.ProtoReflect.Descriptor instead.
func (*BlockBuilderFeeInfoRequest) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{4}
}

type BlockBuilderFeeInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pubkey string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	// commission (0-100)
	Commission uint64 `protobuf:"varint,2,opt,name=commission,proto3" json:"commission,omitempty"`
}

func (x *BlockBuilderFeeInfoResponse) Reset() {
	*x = BlockBuilderFeeInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockBuilderFeeInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockBuilderFeeInfoResponse) ProtoMessage() {}

func (x *BlockBuilderFeeInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockBuilderFeeInfoResponse.ProtoReflect.Descriptor instead.
func (*BlockBuilderFeeInfoResponse) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{5}
}

func (x *BlockBuilderFeeInfoResponse) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

func (x *BlockBuilderFeeInfoResponse) GetCommission() uint64 {
	if x != nil {
		return x.Commission
	}
	return 0
}

type AccountsOfInterest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// use * for all accounts
	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *AccountsOfInterest) Reset() {
	*x = AccountsOfInterest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsOfInterest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsOfInterest) ProtoMessage() {}

func (x *AccountsOfInterest) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsOfInterest.ProtoReflect.Descriptor instead.
func (*AccountsOfInterest) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{6}
}

func (x *AccountsOfInterest) GetAccounts() []string {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type AccountsOfInterestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AccountsOfInterestRequest) Reset() {
	*x = AccountsOfInterestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsOfInterestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsOfInterestRequest) ProtoMessage() {}

func (x *AccountsOfInterestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsOfInterestRequest.ProtoReflect.Descriptor instead.
func (*AccountsOfInterestRequest) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{7}
}

type AccountsOfInterestUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *AccountsOfInterestUpdate) Reset() {
	*x = AccountsOfInterestUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsOfInterestUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsOfInterestUpdate) ProtoMessage() {}

func (x *AccountsOfInterestUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsOfInterestUpdate.ProtoReflect.Descriptor instead.
func (*AccountsOfInterestUpdate) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{8}
}

func (x *AccountsOfInterestUpdate) GetAccounts() []string {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type ProgramsOfInterestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProgramsOfInterestRequest) Reset() {
	*x = ProgramsOfInterestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgramsOfInterestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramsOfInterestRequest) ProtoMessage() {}

func (x *ProgramsOfInterestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgramsOfInterestRequest.ProtoReflect.Descriptor instead.
func (*ProgramsOfInterestRequest) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{9}
}

type ProgramsOfInterestUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Programs []string `protobuf:"bytes,1,rep,name=programs,proto3" json:"programs,omitempty"`
}

func (x *ProgramsOfInterestUpdate) Reset() {
	*x = ProgramsOfInterestUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgramsOfInterestUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramsOfInterestUpdate) ProtoMessage() {}

func (x *ProgramsOfInterestUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgramsOfInterestUpdate.ProtoReflect.Descriptor instead.
func (*ProgramsOfInterestUpdate) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{10}
}

func (x *ProgramsOfInterestUpdate) GetPrograms() []string {
	if x != nil {
		return x.Programs
	}
	return nil
}

// A series of packets with an expiration attached to them.
// The header contains a timestamp for when this packet was generated.
// The expiry is how long the packet batches have before they expire and are forwarded to the validator.
// This provides a more censorship resistant method to MEV than block engines receiving packets directly.
type ExpiringPacketBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header   *Header      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Batch    *PacketBatch `protobuf:"bytes,2,opt,name=batch,proto3" json:"batch,omitempty"`
	ExpiryMs uint32       `protobuf:"varint,3,opt,name=expiry_ms,json=expiryMs,proto3" json:"expiry_ms,omitempty"`
}

func (x *ExpiringPacketBatch) Reset() {
	*x = ExpiringPacketBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpiringPacketBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpiringPacketBatch) ProtoMessage() {}

func (x *ExpiringPacketBatch) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpiringPacketBatch.ProtoReflect.Descriptor instead.
func (*ExpiringPacketBatch) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{11}
}

func (x *ExpiringPacketBatch) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ExpiringPacketBatch) GetBatch() *PacketBatch {
	if x != nil {
		return x.Batch
	}
	return nil
}

func (x *ExpiringPacketBatch) GetExpiryMs() uint32 {
	if x != nil {
		return x.ExpiryMs
	}
	return 0
}

// Packets and heartbeats are sent over the same stream.
// ExpiringPacketBatches have an expiration attached to them so the block engine can track
// how long it has until the relayer forwards the packets to the validator.
// Heartbeats contain a timestamp from the system and is used as a simple and naive time-sync mechanism
// so the block engine has some idea on how far their clocks are apart.
type PacketBatchUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*PacketBatchUpdate_Batches
	//	*PacketBatchUpdate_Heartbeat
	Msg isPacketBatchUpdate_Msg `protobuf_oneof:"msg"`
}

func (x *PacketBatchUpdate) Reset() {
	*x = PacketBatchUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketBatchUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketBatchUpdate) ProtoMessage() {}

func (x *PacketBatchUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketBatchUpdate.ProtoReflect.Descriptor instead.
func (*PacketBatchUpdate) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{12}
}

func (m *PacketBatchUpdate) GetMsg() isPacketBatchUpdate_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *PacketBatchUpdate) GetBatches() *ExpiringPacketBatch {
	if x, ok := x.GetMsg().(*PacketBatchUpdate_Batches); ok {
		return x.Batches
	}
	return nil
}

func (x *PacketBatchUpdate) GetHeartbeat() *Heartbeat {
	if x, ok := x.GetMsg().(*PacketBatchUpdate_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

type isPacketBatchUpdate_Msg interface {
	isPacketBatchUpdate_Msg()
}

type PacketBatchUpdate_Batches struct {
	Batches *ExpiringPacketBatch `protobuf:"bytes,1,opt,name=batches,proto3,oneof"`
}

type PacketBatchUpdate_Heartbeat struct {
	Heartbeat *Heartbeat `protobuf:"bytes,2,opt,name=heartbeat,proto3,oneof"`
}

func (*PacketBatchUpdate_Batches) isPacketBatchUpdate_Msg() {}

func (*PacketBatchUpdate_Heartbeat) isPacketBatchUpdate_Msg() {}

type StartExpiringPacketStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Heartbeat *Heartbeat `protobuf:"bytes,1,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *StartExpiringPacketStreamResponse) Reset() {
	*x = StartExpiringPacketStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_engine_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartExpiringPacketStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartExpiringPacketStreamResponse) ProtoMessage() {}

func (x *StartExpiringPacketStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_block_engine_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartExpiringPacketStreamResponse.ProtoReflect.Descriptor instead.
func (*StartExpiringPacketStreamResponse) Descriptor() ([]byte, []int) {
	return file_block_engine_proto_rawDescGZIP(), []int{13}
}

func (x *StartExpiringPacketStreamResponse) GetHeartbeat() *Heartbeat {
	if x != nil {
		return x.Heartbeat
	}
	return nil
}

var File_block_engine_proto protoreflect.FileDescriptor

var file_block_engine_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x1a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x48, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x07, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x55, 0x75,
	0x69, 0x64, 0x52, 0x07, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x1b, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x30, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x4f, 0x66, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x4f, 0x66,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x36, 0x0a, 0x18, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x4f, 0x66, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x73, 0x4f, 0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73,
	0x4f, 0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x85, 0x01, 0x0a,
	0x13, 0x45, 0x78, 0x70, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x79, 0x4d, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x48, 0x00,
	0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x48,
	0x00, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x42, 0x05, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x54, 0x0a, 0x21, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x09,
	0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x32, 0xd5, 0x02, 0x0a, 0x14, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x65, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x65, 0x0a, 0x10, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x25, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x6f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x46, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0xf1, 0x02, 0x0a, 0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x72, 0x0a, 0x1b, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x4f, 0x66, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x4f,
	0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x4f, 0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x72, 0x0a, 0x1b,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x73, 0x4f, 0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x73, 0x4f, 0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x4f, 0x66, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x73, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1f, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x2f,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x6a, 0x69, 0x74, 0x6f, 0x2f, 0x6a, 0x69,
	0x74, 0x6f, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_engine_proto_rawDescOnce sync.Once
	file_block_engine_proto_rawDescData = file_block_engine_proto_rawDesc
)

func file_block_engine_proto_rawDescGZIP() []byte {
	file_block_engine_proto_rawDescOnce.Do(func() {
		file_block_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_engine_proto_rawDescData)
	})
	return file_block_engine_proto_rawDescData
}

var file_block_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_block_engine_proto_goTypes = []interface{}{
	(*SubscribePacketsRequest)(nil),           // 0: block_engine.SubscribePacketsRequest
	(*SubscribePacketsResponse)(nil),          // 1: block_engine.SubscribePacketsResponse
	(*SubscribeBundlesRequest)(nil),           // 2: block_engine.SubscribeBundlesRequest
	(*SubscribeBundlesResponse)(nil),          // 3: block_engine.SubscribeBundlesResponse
	(*BlockBuilderFeeInfoRequest)(nil),        // 4: block_engine.BlockBuilderFeeInfoRequest
	(*BlockBuilderFeeInfoResponse)(nil),       // 5: block_engine.BlockBuilderFeeInfoResponse
	(*AccountsOfInterest)(nil),                // 6: block_engine.AccountsOfInterest
	(*AccountsOfInterestRequest)(nil),         // 7: block_engine.AccountsOfInterestRequest
	(*AccountsOfInterestUpdate)(nil),          // 8: block_engine.AccountsOfInterestUpdate
	(*ProgramsOfInterestRequest)(nil),         // 9: block_engine.ProgramsOfInterestRequest
	(*ProgramsOfInterestUpdate)(nil),          // 10: block_engine.ProgramsOfInterestUpdate
	(*ExpiringPacketBatch)(nil),               // 11: block_engine.ExpiringPacketBatch
	(*PacketBatchUpdate)(nil),                 // 12: block_engine.PacketBatchUpdate
	(*StartExpiringPacketStreamResponse)(nil), // 13: block_engine.StartExpiringPacketStreamResponse
	(*Header)(nil),                            // 14: shared.Header
	(*PacketBatch)(nil),                       // 15: packet.PacketBatch
	(*BundleUuid)(nil),                        // 16: bundle.BundleUuid
	(*Heartbeat)(nil),                         // 17: shared.Heartbeat
}
var file_block_engine_proto_depIdxs = []int32{
	14, // 0: block_engine.SubscribePacketsResponse.header:type_name -> shared.Header
	15, // 1: block_engine.SubscribePacketsResponse.batch:type_name -> packet.PacketBatch
	16, // 2: block_engine.SubscribeBundlesResponse.bundles:type_name -> bundle.BundleUuid
	14, // 3: block_engine.ExpiringPacketBatch.header:type_name -> shared.Header
	15, // 4: block_engine.ExpiringPacketBatch.batch:type_name -> packet.PacketBatch
	11, // 5: block_engine.PacketBatchUpdate.batches:type_name -> block_engine.ExpiringPacketBatch
	17, // 6: block_engine.PacketBatchUpdate.heartbeat:type_name -> shared.Heartbeat
	17, // 7: block_engine.StartExpiringPacketStreamResponse.heartbeat:type_name -> shared.Heartbeat
	0,  // 8: block_engine.BlockEngineValidator.SubscribePackets:input_type -> block_engine.SubscribePacketsRequest
	2,  // 9: block_engine.BlockEngineValidator.SubscribeBundles:input_type -> block_engine.SubscribeBundlesRequest
	4,  // 10: block_engine.BlockEngineValidator.GetBlockBuilderFeeInfo:input_type -> block_engine.BlockBuilderFeeInfoRequest
	7,  // 11: block_engine.BlockEngineRelayer.SubscribeAccountsOfInterest:input_type -> block_engine.AccountsOfInterestRequest
	9,  // 12: block_engine.BlockEngineRelayer.SubscribeProgramsOfInterest:input_type -> block_engine.ProgramsOfInterestRequest
	12, // 13: block_engine.BlockEngineRelayer.StartExpiringPacketStream:input_type -> block_engine.PacketBatchUpdate
	1,  // 14: block_engine.BlockEngineValidator.SubscribePackets:output_type -> block_engine.SubscribePacketsResponse
	3,  // 15: block_engine.BlockEngineValidator.SubscribeBundles:output_type -> block_engine.SubscribeBundlesResponse
	5,  // 16: block_engine.BlockEngineValidator.GetBlockBuilderFeeInfo:output_type -> block_engine.BlockBuilderFeeInfoResponse
	8,  // 17: block_engine.BlockEngineRelayer.SubscribeAccountsOfInterest:output_type -> block_engine.AccountsOfInterestUpdate
	10, // 18: block_engine.BlockEngineRelayer.SubscribeProgramsOfInterest:output_type -> block_engine.ProgramsOfInterestUpdate
	13, // 19: block_engine.BlockEngineRelayer.StartExpiringPacketStream:output_type -> block_engine.StartExpiringPacketStreamResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_block_engine_proto_init() }
func file_block_engine_proto_init() {
	if File_block_engine_proto != nil {
		return
	}
	file_packet_proto_init()
	file_shared_proto_init()
	file_bundle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_block_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribePacketsRequest); i {
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
		file_block_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribePacketsResponse); i {
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
		file_block_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeBundlesRequest); i {
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
		file_block_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeBundlesResponse); i {
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
		file_block_engine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockBuilderFeeInfoRequest); i {
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
		file_block_engine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockBuilderFeeInfoResponse); i {
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
		file_block_engine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountsOfInterest); i {
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
		file_block_engine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountsOfInterestRequest); i {
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
		file_block_engine_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountsOfInterestUpdate); i {
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
		file_block_engine_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgramsOfInterestRequest); i {
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
		file_block_engine_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgramsOfInterestUpdate); i {
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
		file_block_engine_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpiringPacketBatch); i {
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
		file_block_engine_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketBatchUpdate); i {
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
		file_block_engine_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartExpiringPacketStreamResponse); i {
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
	file_block_engine_proto_msgTypes[12].OneofWrappers = []interface{}{
		(*PacketBatchUpdate_Batches)(nil),
		(*PacketBatchUpdate_Heartbeat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_block_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_block_engine_proto_goTypes,
		DependencyIndexes: file_block_engine_proto_depIdxs,
		MessageInfos:      file_block_engine_proto_msgTypes,
	}.Build()
	File_block_engine_proto = out.File
	file_block_engine_proto_rawDesc = nil
	file_block_engine_proto_goTypes = nil
	file_block_engine_proto_depIdxs = nil
}
