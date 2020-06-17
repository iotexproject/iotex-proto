// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/election.proto

package iotextypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ElectionBucket struct {
	Voter                []byte               `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	Candidate            []byte               `protobuf:"bytes,2,opt,name=candidate,proto3" json:"candidate,omitempty"`
	Amount               []byte               `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	Decay                bool                 `protobuf:"varint,6,opt,name=decay,proto3" json:"decay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ElectionBucket) Reset()         { *m = ElectionBucket{} }
func (m *ElectionBucket) String() string { return proto.CompactTextString(m) }
func (*ElectionBucket) ProtoMessage()    {}
func (*ElectionBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_79eba0ff0cac3135, []int{0}
}

func (m *ElectionBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElectionBucket.Unmarshal(m, b)
}
func (m *ElectionBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElectionBucket.Marshal(b, m, deterministic)
}
func (m *ElectionBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElectionBucket.Merge(m, src)
}
func (m *ElectionBucket) XXX_Size() int {
	return xxx_messageInfo_ElectionBucket.Size(m)
}
func (m *ElectionBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_ElectionBucket.DiscardUnknown(m)
}

var xxx_messageInfo_ElectionBucket proto.InternalMessageInfo

func (m *ElectionBucket) GetVoter() []byte {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *ElectionBucket) GetCandidate() []byte {
	if m != nil {
		return m.Candidate
	}
	return nil
}

func (m *ElectionBucket) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *ElectionBucket) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ElectionBucket) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *ElectionBucket) GetDecay() bool {
	if m != nil {
		return m.Decay
	}
	return false
}

func init() {
	proto.RegisterType((*ElectionBucket)(nil), "iotextypes.ElectionBucket")
}

func init() { proto.RegisterFile("proto/types/election.proto", fileDescriptor_79eba0ff0cac3135) }

var fileDescriptor_79eba0ff0cac3135 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0xba, 0x65, 0x37, 0x8a, 0x87, 0x20, 0x12, 0x8b, 0x68, 0xf1, 0xd4, 0x8b, 0x0d,
	0x28, 0x8b, 0x9e, 0x17, 0x7d, 0x81, 0xb2, 0x27, 0x6f, 0x69, 0x3a, 0xd6, 0x68, 0xdb, 0x29, 0xe9,
	0x44, 0xdc, 0xe7, 0xf5, 0x45, 0xc4, 0xa4, 0xb5, 0xe0, 0x1e, 0xff, 0xf9, 0xbe, 0x09, 0x7f, 0x86,
	0x25, 0xbd, 0x45, 0x42, 0x49, 0xbb, 0x1e, 0x06, 0x09, 0x0d, 0x68, 0x32, 0xd8, 0xe5, 0x7e, 0xc8,
	0x99, 0x41, 0x82, 0x2f, 0x8f, 0x92, 0xab, 0x1a, 0xb1, 0x6e, 0x40, 0x7a, 0x52, 0xba, 0x57, 0x59,
	0x39, 0xab, 0x66, 0x37, 0xb9, 0xfe, 0xcf, 0xc9, 0xb4, 0x30, 0x90, 0x6a, 0xfb, 0x20, 0xdc, 0x7c,
	0x47, 0xec, 0xf4, 0x79, 0x7c, 0x7f, 0xe3, 0xf4, 0x07, 0x10, 0x3f, 0x63, 0x8b, 0x4f, 0x24, 0xb0,
	0x22, 0x4a, 0xa3, 0xec, 0xa4, 0x08, 0x81, 0x5f, 0xb2, 0x95, 0x56, 0x5d, 0x65, 0x2a, 0x45, 0x20,
	0x0e, 0x3c, 0x99, 0x07, 0xfc, 0x9c, 0xc5, 0xaa, 0x45, 0xd7, 0x91, 0x38, 0xf4, 0x68, 0x4c, 0xfc,
	0x91, 0xad, 0x06, 0x52, 0x96, 0xb6, 0xa6, 0x05, 0x71, 0x94, 0x46, 0xd9, 0xf1, 0x5d, 0x92, 0x87,
	0x4e, 0xf9, 0xd4, 0x29, 0xdf, 0x4e, 0x9d, 0x8a, 0x59, 0xe6, 0x6b, 0xb6, 0x9c, 0xfe, 0x22, 0x16,
	0x7e, 0xf1, 0x62, 0x6f, 0xf1, 0x69, 0x14, 0x8a, 0x3f, 0xf5, 0xb7, 0x7c, 0x05, 0x5a, 0xed, 0x44,
	0x9c, 0x46, 0xd9, 0xb2, 0x08, 0x61, 0xf3, 0xf0, 0xb2, 0xae, 0x0d, 0xbd, 0xb9, 0x32, 0xd7, 0xd8,
	0x4a, 0x7f, 0xbf, 0xde, 0xe2, 0x3b, 0x68, 0x0a, 0xe1, 0x36, 0x9c, 0xbb, 0xc6, 0x46, 0x75, 0xb5,
	0x9c, 0xef, 0x5b, 0xc6, 0x1e, 0xdc, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x80, 0x8c, 0x08, 0x9b,
	0x90, 0x01, 0x00, 0x00,
}
