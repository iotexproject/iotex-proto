// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// To compile the proto, run:
//      protoc --go_out=plugins=grpc:$GOPATH/src *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/types/receiptstatus.proto

package iotextypes

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// BELOW ARE DEFINITIONS FOR EVM ERROR CLASSIFICATION IN RECEIPT STATUS
type ReceiptStatus int32

const (
	ReceiptStatus_Failure ReceiptStatus = 0
	ReceiptStatus_Success ReceiptStatus = 1
	//1xx for EVM ErrorCode
	ReceiptStatus_ErrUnknown                  ReceiptStatus = 100
	ReceiptStatus_ErrOutOfGas                 ReceiptStatus = 101
	ReceiptStatus_ErrCodeStoreOutOfGas        ReceiptStatus = 102
	ReceiptStatus_ErrDepth                    ReceiptStatus = 103
	ReceiptStatus_ErrContractAddressCollision ReceiptStatus = 104
	ReceiptStatus_ErrNoCompatibleInterpreter  ReceiptStatus = 105
	ReceiptStatus_ErrExecutionReverted        ReceiptStatus = 106
	ReceiptStatus_ErrMaxCodeSizeExceeded      ReceiptStatus = 107
	ReceiptStatus_ErrWriteProtection          ReceiptStatus = 108
	//2xx for Staking ErrorCode
	ReceiptStatus_ErrLoadAccount                  ReceiptStatus = 200
	ReceiptStatus_ErrNotEnoughBalance             ReceiptStatus = 201
	ReceiptStatus_ErrInvalidBucketIndex           ReceiptStatus = 202
	ReceiptStatus_ErrUnauthorizedOperator         ReceiptStatus = 203
	ReceiptStatus_ErrInvalidBucketType            ReceiptStatus = 204
	ReceiptStatus_ErrCandidateNotExist            ReceiptStatus = 205
	ReceiptStatus_ErrReduceDurationBeforeMaturity ReceiptStatus = 206
	ReceiptStatus_ErrUnstakeBeforeMaturity        ReceiptStatus = 207
	ReceiptStatus_ErrWithdrawBeforeUnstake        ReceiptStatus = 208
	ReceiptStatus_ErrWithdrawBeforeMaturity       ReceiptStatus = 209
	ReceiptStatus_ErrCandidateAlreadyExist        ReceiptStatus = 210
	ReceiptStatus_ErrCandidateConflict            ReceiptStatus = 211
	ReceiptStatus_ErrInvalidBucketAmount          ReceiptStatus = 212
	ReceiptStatus_ErrWriteAccount                 ReceiptStatus = 213
	ReceiptStatus_ErrWriteBucket                  ReceiptStatus = 214
	ReceiptStatus_ErrWriteCandidate               ReceiptStatus = 215
)

// Enum value maps for ReceiptStatus.
var (
	ReceiptStatus_name = map[int32]string{
		0:   "Failure",
		1:   "Success",
		100: "ErrUnknown",
		101: "ErrOutOfGas",
		102: "ErrCodeStoreOutOfGas",
		103: "ErrDepth",
		104: "ErrContractAddressCollision",
		105: "ErrNoCompatibleInterpreter",
		106: "ErrExecutionReverted",
		107: "ErrMaxCodeSizeExceeded",
		108: "ErrWriteProtection",
		200: "ErrLoadAccount",
		201: "ErrNotEnoughBalance",
		202: "ErrInvalidBucketIndex",
		203: "ErrUnauthorizedOperator",
		204: "ErrInvalidBucketType",
		205: "ErrCandidateNotExist",
		206: "ErrReduceDurationBeforeMaturity",
		207: "ErrUnstakeBeforeMaturity",
		208: "ErrWithdrawBeforeUnstake",
		209: "ErrWithdrawBeforeMaturity",
		210: "ErrCandidateAlreadyExist",
		211: "ErrCandidateConflict",
		212: "ErrInvalidBucketAmount",
		213: "ErrWriteAccount",
		214: "ErrWriteBucket",
		215: "ErrWriteCandidate",
	}
	ReceiptStatus_value = map[string]int32{
		"Failure":                         0,
		"Success":                         1,
		"ErrUnknown":                      100,
		"ErrOutOfGas":                     101,
		"ErrCodeStoreOutOfGas":            102,
		"ErrDepth":                        103,
		"ErrContractAddressCollision":     104,
		"ErrNoCompatibleInterpreter":      105,
		"ErrExecutionReverted":            106,
		"ErrMaxCodeSizeExceeded":          107,
		"ErrWriteProtection":              108,
		"ErrLoadAccount":                  200,
		"ErrNotEnoughBalance":             201,
		"ErrInvalidBucketIndex":           202,
		"ErrUnauthorizedOperator":         203,
		"ErrInvalidBucketType":            204,
		"ErrCandidateNotExist":            205,
		"ErrReduceDurationBeforeMaturity": 206,
		"ErrUnstakeBeforeMaturity":        207,
		"ErrWithdrawBeforeUnstake":        208,
		"ErrWithdrawBeforeMaturity":       209,
		"ErrCandidateAlreadyExist":        210,
		"ErrCandidateConflict":            211,
		"ErrInvalidBucketAmount":          212,
		"ErrWriteAccount":                 213,
		"ErrWriteBucket":                  214,
		"ErrWriteCandidate":               215,
	}
)

func (x ReceiptStatus) Enum() *ReceiptStatus {
	p := new(ReceiptStatus)
	*p = x
	return p
}

func (x ReceiptStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReceiptStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_receiptstatus_proto_enumTypes[0].Descriptor()
}

func (ReceiptStatus) Type() protoreflect.EnumType {
	return &file_proto_types_receiptstatus_proto_enumTypes[0]
}

func (x ReceiptStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReceiptStatus.Descriptor instead.
func (ReceiptStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_receiptstatus_proto_rawDescGZIP(), []int{0}
}

var File_proto_types_receiptstatus_proto protoreflect.FileDescriptor

var file_proto_types_receiptstatus_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2a, 0xbe, 0x05,
	0x0a, 0x0d, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x72, 0x72,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x72, 0x72,
	0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x10, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47,
	0x61, 0x73, 0x10, 0x66, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x44, 0x65, 0x70, 0x74, 0x68,
	0x10, 0x67, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x10, 0x68, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x72, 0x72, 0x4e, 0x6f, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65,
	0x72, 0x10, 0x69, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x72, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x10, 0x6a, 0x12, 0x1a, 0x0a,
	0x16, 0x45, 0x72, 0x72, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x45,
	0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x6b, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x72, 0x72,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x6c, 0x12, 0x13, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x10, 0xc8, 0x01, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x72, 0x72, 0x4e, 0x6f, 0x74,
	0x45, 0x6e, 0x6f, 0x75, 0x67, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10, 0xc9, 0x01,
	0x12, 0x1a, 0x0a, 0x15, 0x45, 0x72, 0x72, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x10, 0xca, 0x01, 0x12, 0x1c, 0x0a, 0x17,
	0x45, 0x72, 0x72, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x10, 0xcb, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x72,
	0x72, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x10, 0xcc, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x72, 0x72, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xcd, 0x01,
	0x12, 0x24, 0x0a, 0x1f, 0x45, 0x72, 0x72, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x74, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x10, 0xce, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x72, 0x72, 0x55, 0x6e, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x74, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x10, 0xcf, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x72, 0x72, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b,
	0x65, 0x10, 0xd0, 0x01, 0x12, 0x1e, 0x0a, 0x19, 0x45, 0x72, 0x72, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x10, 0xd1, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x72, 0x72, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x10, 0xd2, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x45, 0x72, 0x72, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x10, 0xd3, 0x01, 0x12, 0x1b,
	0x0a, 0x16, 0x45, 0x72, 0x72, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0xd4, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x45,
	0x72, 0x72, 0x57, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0xd5,
	0x01, 0x12, 0x13, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x10, 0xd6, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x10, 0xd7, 0x01, 0x42, 0x5d,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x74,
	0x65, 0x78, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x69, 0x6f, 0x74, 0x65, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_receiptstatus_proto_rawDescOnce sync.Once
	file_proto_types_receiptstatus_proto_rawDescData = file_proto_types_receiptstatus_proto_rawDesc
)

func file_proto_types_receiptstatus_proto_rawDescGZIP() []byte {
	file_proto_types_receiptstatus_proto_rawDescOnce.Do(func() {
		file_proto_types_receiptstatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_receiptstatus_proto_rawDescData)
	})
	return file_proto_types_receiptstatus_proto_rawDescData
}

var file_proto_types_receiptstatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_types_receiptstatus_proto_goTypes = []interface{}{
	(ReceiptStatus)(0), // 0: iotextypes.ReceiptStatus
}
var file_proto_types_receiptstatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_types_receiptstatus_proto_init() }
func file_proto_types_receiptstatus_proto_init() {
	if File_proto_types_receiptstatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_types_receiptstatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_receiptstatus_proto_goTypes,
		DependencyIndexes: file_proto_types_receiptstatus_proto_depIdxs,
		EnumInfos:         file_proto_types_receiptstatus_proto_enumTypes,
	}.Build()
	File_proto_types_receiptstatus_proto = out.File
	file_proto_types_receiptstatus_proto_rawDesc = nil
	file_proto_types_receiptstatus_proto_goTypes = nil
	file_proto_types_receiptstatus_proto_depIdxs = nil
}
