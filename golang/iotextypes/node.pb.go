// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// To compile the proto, run:
//      protoc --go_out=plugins=grpc:$GOPATH/src *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: proto/types/node.proto

package iotextypes

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

// Server Metadata
type ServerMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageVersion  string `protobuf:"bytes,1,opt,name=packageVersion,proto3" json:"packageVersion,omitempty"`
	PackageCommitID string `protobuf:"bytes,2,opt,name=packageCommitID,proto3" json:"packageCommitID,omitempty"`
	GitStatus       string `protobuf:"bytes,3,opt,name=gitStatus,proto3" json:"gitStatus,omitempty"`
	GoVersion       string `protobuf:"bytes,4,opt,name=goVersion,proto3" json:"goVersion,omitempty"`
	BuildTime       string `protobuf:"bytes,5,opt,name=buildTime,proto3" json:"buildTime,omitempty"`
}

func (x *ServerMeta) Reset() {
	*x = ServerMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMeta) ProtoMessage() {}

func (x *ServerMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMeta.ProtoReflect.Descriptor instead.
func (*ServerMeta) Descriptor() ([]byte, []int) {
	return file_proto_types_node_proto_rawDescGZIP(), []int{0}
}

func (x *ServerMeta) GetPackageVersion() string {
	if x != nil {
		return x.PackageVersion
	}
	return ""
}

func (x *ServerMeta) GetPackageCommitID() string {
	if x != nil {
		return x.PackageCommitID
	}
	return ""
}

func (x *ServerMeta) GetGitStatus() string {
	if x != nil {
		return x.GitStatus
	}
	return ""
}

func (x *ServerMeta) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

func (x *ServerMeta) GetBuildTime() string {
	if x != nil {
		return x.BuildTime
	}
	return ""
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Height    uint64                 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Address   string                 `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_proto_types_node_proto_rawDescGZIP(), []int{1}
}

func (x *NodeInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NodeInfo) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *NodeInfo) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *NodeInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ResponseNodeInfoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *NodeInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Signature []byte    `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ResponseNodeInfoMessage) Reset() {
	*x = ResponseNodeInfoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseNodeInfoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseNodeInfoMessage) ProtoMessage() {}

func (x *ResponseNodeInfoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseNodeInfoMessage.ProtoReflect.Descriptor instead.
func (*ResponseNodeInfoMessage) Descriptor() ([]byte, []int) {
	return file_proto_types_node_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseNodeInfoMessage) GetInfo() *NodeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ResponseNodeInfoMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type RequestNodeInfoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestNodeInfoMessage) Reset() {
	*x = RequestNodeInfoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestNodeInfoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestNodeInfoMessage) ProtoMessage() {}

func (x *RequestNodeInfoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestNodeInfoMessage.ProtoReflect.Descriptor instead.
func (*RequestNodeInfoMessage) Descriptor() ([]byte, []int) {
	return file_proto_types_node_proto_rawDescGZIP(), []int{3}
}

var File_proto_types_node_proto protoreflect.FileDescriptor

var file_proto_types_node_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x69, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x90, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x61, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69,
	0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x5d, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69,
	0x6f, 0x74, 0x65, 0x78, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_node_proto_rawDescOnce sync.Once
	file_proto_types_node_proto_rawDescData = file_proto_types_node_proto_rawDesc
)

func file_proto_types_node_proto_rawDescGZIP() []byte {
	file_proto_types_node_proto_rawDescOnce.Do(func() {
		file_proto_types_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_node_proto_rawDescData)
	})
	return file_proto_types_node_proto_rawDescData
}

var file_proto_types_node_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_types_node_proto_goTypes = []interface{}{
	(*ServerMeta)(nil),              // 0: iotextypes.ServerMeta
	(*NodeInfo)(nil),                // 1: iotextypes.NodeInfo
	(*ResponseNodeInfoMessage)(nil), // 2: iotextypes.ResponseNodeInfoMessage
	(*RequestNodeInfoMessage)(nil),  // 3: iotextypes.RequestNodeInfoMessage
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
}
var file_proto_types_node_proto_depIdxs = []int32{
	4, // 0: iotextypes.NodeInfo.timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: iotextypes.ResponseNodeInfoMessage.info:type_name -> iotextypes.NodeInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_types_node_proto_init() }
func file_proto_types_node_proto_init() {
	if File_proto_types_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMeta); i {
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
		file_proto_types_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_proto_types_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseNodeInfoMessage); i {
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
		file_proto_types_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestNodeInfoMessage); i {
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
			RawDescriptor: file_proto_types_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_node_proto_goTypes,
		DependencyIndexes: file_proto_types_node_proto_depIdxs,
		MessageInfos:      file_proto_types_node_proto_msgTypes,
	}.Build()
	File_proto_types_node_proto = out.File
	file_proto_types_node_proto_rawDesc = nil
	file_proto_types_node_proto_goTypes = nil
	file_proto_types_node_proto_depIdxs = nil
}
