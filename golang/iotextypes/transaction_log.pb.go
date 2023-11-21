// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// To compile the proto, run:
//      protoc --go_out=plugins=grpc:$GOPATH/src *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.17.3
// source: proto/types/transaction_log.proto

package iotextypes

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

type TransactionLogType int32

const (
	TransactionLogType_IN_CONTRACT_TRANSFER       TransactionLogType = 0
	TransactionLogType_WITHDRAW_BUCKET            TransactionLogType = 1
	TransactionLogType_CREATE_BUCKET              TransactionLogType = 2
	TransactionLogType_DEPOSIT_TO_BUCKET          TransactionLogType = 3
	TransactionLogType_CANDIDATE_SELF_STAKE       TransactionLogType = 4
	TransactionLogType_CANDIDATE_REGISTRATION_FEE TransactionLogType = 5
	TransactionLogType_GAS_FEE                    TransactionLogType = 6
	TransactionLogType_NATIVE_TRANSFER            TransactionLogType = 7
	TransactionLogType_DEPOSIT_TO_REWARDING_FUND  TransactionLogType = 8
	TransactionLogType_CLAIM_FROM_REWARDING_FUND  TransactionLogType = 9
)

// Enum value maps for TransactionLogType.
var (
	TransactionLogType_name = map[int32]string{
		0: "IN_CONTRACT_TRANSFER",
		1: "WITHDRAW_BUCKET",
		2: "CREATE_BUCKET",
		3: "DEPOSIT_TO_BUCKET",
		4: "CANDIDATE_SELF_STAKE",
		5: "CANDIDATE_REGISTRATION_FEE",
		6: "GAS_FEE",
		7: "NATIVE_TRANSFER",
		8: "DEPOSIT_TO_REWARDING_FUND",
		9: "CLAIM_FROM_REWARDING_FUND",
	}
	TransactionLogType_value = map[string]int32{
		"IN_CONTRACT_TRANSFER":       0,
		"WITHDRAW_BUCKET":            1,
		"CREATE_BUCKET":              2,
		"DEPOSIT_TO_BUCKET":          3,
		"CANDIDATE_SELF_STAKE":       4,
		"CANDIDATE_REGISTRATION_FEE": 5,
		"GAS_FEE":                    6,
		"NATIVE_TRANSFER":            7,
		"DEPOSIT_TO_REWARDING_FUND":  8,
		"CLAIM_FROM_REWARDING_FUND":  9,
	}
)

func (x TransactionLogType) Enum() *TransactionLogType {
	p := new(TransactionLogType)
	*p = x
	return p
}

func (x TransactionLogType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionLogType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_transaction_log_proto_enumTypes[0].Descriptor()
}

func (TransactionLogType) Type() protoreflect.EnumType {
	return &file_proto_types_transaction_log_proto_enumTypes[0]
}

func (x TransactionLogType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionLogType.Descriptor instead.
func (TransactionLogType) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_transaction_log_proto_rawDescGZIP(), []int{0}
}

type TransactionLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionHash      []byte                        `protobuf:"bytes,1,opt,name=actionHash,proto3" json:"actionHash,omitempty"`
	NumTransactions uint64                        `protobuf:"varint,2,opt,name=numTransactions,proto3" json:"numTransactions,omitempty"`
	Transactions    []*TransactionLog_Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *TransactionLog) Reset() {
	*x = TransactionLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_transaction_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionLog) ProtoMessage() {}

func (x *TransactionLog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_transaction_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionLog.ProtoReflect.Descriptor instead.
func (*TransactionLog) Descriptor() ([]byte, []int) {
	return file_proto_types_transaction_log_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionLog) GetActionHash() []byte {
	if x != nil {
		return x.ActionHash
	}
	return nil
}

func (x *TransactionLog) GetNumTransactions() uint64 {
	if x != nil {
		return x.NumTransactions
	}
	return 0
}

func (x *TransactionLog) GetTransactions() []*TransactionLog_Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type TransactionLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*TransactionLog `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *TransactionLogs) Reset() {
	*x = TransactionLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_transaction_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionLogs) ProtoMessage() {}

func (x *TransactionLogs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_transaction_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionLogs.ProtoReflect.Descriptor instead.
func (*TransactionLogs) Descriptor() ([]byte, []int) {
	return file_proto_types_transaction_log_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionLogs) GetLogs() []*TransactionLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

type TransactionStructLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pc         uint64   `protobuf:"varint,1,opt,name=pc,proto3" json:"pc,omitempty"`
	Op         uint64   `protobuf:"varint,2,opt,name=op,proto3" json:"op,omitempty"`
	Gas        uint64   `protobuf:"varint,3,opt,name=gas,proto3" json:"gas,omitempty"`
	GasCost    uint64   `protobuf:"varint,4,opt,name=gasCost,proto3" json:"gasCost,omitempty"`
	Memory     string   `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	MemSize    int32    `protobuf:"varint,6,opt,name=memSize,proto3" json:"memSize,omitempty"`
	Stack      []string `protobuf:"bytes,7,rep,name=stack,proto3" json:"stack,omitempty"`
	ReturnData string   `protobuf:"bytes,8,opt,name=returnData,proto3" json:"returnData,omitempty"`
	Depth      int32    `protobuf:"varint,9,opt,name=depth,proto3" json:"depth,omitempty"`
	Refund     uint64   `protobuf:"varint,10,opt,name=refund,proto3" json:"refund,omitempty"`
	OpName     string   `protobuf:"bytes,11,opt,name=opName,proto3" json:"opName,omitempty"`
	Error      string   `protobuf:"bytes,12,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TransactionStructLog) Reset() {
	*x = TransactionStructLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_transaction_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionStructLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionStructLog) ProtoMessage() {}

func (x *TransactionStructLog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_transaction_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionStructLog.ProtoReflect.Descriptor instead.
func (*TransactionStructLog) Descriptor() ([]byte, []int) {
	return file_proto_types_transaction_log_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionStructLog) GetPc() uint64 {
	if x != nil {
		return x.Pc
	}
	return 0
}

func (x *TransactionStructLog) GetOp() uint64 {
	if x != nil {
		return x.Op
	}
	return 0
}

func (x *TransactionStructLog) GetGas() uint64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *TransactionStructLog) GetGasCost() uint64 {
	if x != nil {
		return x.GasCost
	}
	return 0
}

func (x *TransactionStructLog) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *TransactionStructLog) GetMemSize() int32 {
	if x != nil {
		return x.MemSize
	}
	return 0
}

func (x *TransactionStructLog) GetStack() []string {
	if x != nil {
		return x.Stack
	}
	return nil
}

func (x *TransactionStructLog) GetReturnData() string {
	if x != nil {
		return x.ReturnData
	}
	return ""
}

func (x *TransactionStructLog) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *TransactionStructLog) GetRefund() uint64 {
	if x != nil {
		return x.Refund
	}
	return 0
}

func (x *TransactionStructLog) GetOpName() string {
	if x != nil {
		return x.OpName
	}
	return ""
}

func (x *TransactionStructLog) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type TransactionLog_Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     []byte             `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Amount    string             `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Sender    string             `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient string             `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Type      TransactionLogType `protobuf:"varint,5,opt,name=type,proto3,enum=iotextypes.TransactionLogType" json:"type,omitempty"`
}

func (x *TransactionLog_Transaction) Reset() {
	*x = TransactionLog_Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_transaction_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionLog_Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionLog_Transaction) ProtoMessage() {}

func (x *TransactionLog_Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_transaction_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionLog_Transaction.ProtoReflect.Descriptor instead.
func (*TransactionLog_Transaction) Descriptor() ([]byte, []int) {
	return file_proto_types_transaction_log_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TransactionLog_Transaction) GetTopic() []byte {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *TransactionLog_Transaction) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TransactionLog_Transaction) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *TransactionLog_Transaction) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *TransactionLog_Transaction) GetType() TransactionLogType {
	if x != nil {
		return x.Type
	}
	return TransactionLogType_IN_CONTRACT_TRANSFER
}

var File_proto_types_transaction_log_proto protoreflect.FileDescriptor

var file_proto_types_transaction_log_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22,
	0xce, 0x02, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6e, 0x75, 0x6d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4a, 0x0a, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xa5, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6f, 0x74,
	0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x41, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x22, 0xa6, 0x02, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02,
	0x70, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x70, 0x63, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x67, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x61, 0x73, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x67, 0x61, 0x73, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x87, 0x02, 0x0a,
	0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41,
	0x43, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x55, 0x43,
	0x4b, 0x45, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x5f, 0x54, 0x4f, 0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14,
	0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x46, 0x5f, 0x53,
	0x54, 0x41, 0x4b, 0x45, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x46, 0x45, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x41, 0x53, 0x5f, 0x46, 0x45,
	0x45, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4f,
	0x53, 0x49, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x46, 0x55, 0x4e, 0x44, 0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4c, 0x41, 0x49, 0x4d,
	0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x5f,
	0x46, 0x55, 0x4e, 0x44, 0x10, 0x09, 0x42, 0x5d, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_transaction_log_proto_rawDescOnce sync.Once
	file_proto_types_transaction_log_proto_rawDescData = file_proto_types_transaction_log_proto_rawDesc
)

func file_proto_types_transaction_log_proto_rawDescGZIP() []byte {
	file_proto_types_transaction_log_proto_rawDescOnce.Do(func() {
		file_proto_types_transaction_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_transaction_log_proto_rawDescData)
	})
	return file_proto_types_transaction_log_proto_rawDescData
}

var file_proto_types_transaction_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_types_transaction_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_types_transaction_log_proto_goTypes = []interface{}{
	(TransactionLogType)(0),            // 0: iotextypes.TransactionLogType
	(*TransactionLog)(nil),             // 1: iotextypes.TransactionLog
	(*TransactionLogs)(nil),            // 2: iotextypes.TransactionLogs
	(*TransactionStructLog)(nil),       // 3: iotextypes.TransactionStructLog
	(*TransactionLog_Transaction)(nil), // 4: iotextypes.TransactionLog.Transaction
}
var file_proto_types_transaction_log_proto_depIdxs = []int32{
	4, // 0: iotextypes.TransactionLog.transactions:type_name -> iotextypes.TransactionLog.Transaction
	1, // 1: iotextypes.TransactionLogs.logs:type_name -> iotextypes.TransactionLog
	0, // 2: iotextypes.TransactionLog.Transaction.type:type_name -> iotextypes.TransactionLogType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_types_transaction_log_proto_init() }
func file_proto_types_transaction_log_proto_init() {
	if File_proto_types_transaction_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_transaction_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionLog); i {
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
		file_proto_types_transaction_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionLogs); i {
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
		file_proto_types_transaction_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionStructLog); i {
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
		file_proto_types_transaction_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionLog_Transaction); i {
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
			RawDescriptor: file_proto_types_transaction_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_transaction_log_proto_goTypes,
		DependencyIndexes: file_proto_types_transaction_log_proto_depIdxs,
		EnumInfos:         file_proto_types_transaction_log_proto_enumTypes,
		MessageInfos:      file_proto_types_transaction_log_proto_msgTypes,
	}.Build()
	File_proto_types_transaction_log_proto = out.File
	file_proto_types_transaction_log_proto_rawDesc = nil
	file_proto_types_transaction_log_proto_goTypes = nil
	file_proto_types_transaction_log_proto_depIdxs = nil
}
