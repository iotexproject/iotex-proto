// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api/read_state.proto

package iotexapi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ReadStakingDataMethod_Name int32

const (
	ReadStakingDataMethod_INVALID                  ReadStakingDataMethod_Name = 0
	ReadStakingDataMethod_GET_BUCKETS              ReadStakingDataMethod_Name = 1
	ReadStakingDataMethod_GET_BUCKETS_BY_VOTER     ReadStakingDataMethod_Name = 2
	ReadStakingDataMethod_GET_BUCKETS_BY_CANDIDATE ReadStakingDataMethod_Name = 3
	ReadStakingDataMethod_GET_CANDIDATE            ReadStakingDataMethod_Name = 4
	ReadStakingDataMethod_GET_CANDIDATE_BY_NAME    ReadStakingDataMethod_Name = 5
)

var ReadStakingDataMethod_Name_name = map[int32]string{
	0: "INVALID",
	1: "GET_BUCKETS",
	2: "GET_BUCKETS_BY_VOTER",
	3: "GET_BUCKETS_BY_CANDIDATE",
	4: "GET_CANDIDATE",
	5: "GET_CANDIDATE_BY_NAME",
}

var ReadStakingDataMethod_Name_value = map[string]int32{
	"INVALID":                  0,
	"GET_BUCKETS":              1,
	"GET_BUCKETS_BY_VOTER":     2,
	"GET_BUCKETS_BY_CANDIDATE": 3,
	"GET_CANDIDATE":            4,
	"GET_CANDIDATE_BY_NAME":    5,
}

func (x ReadStakingDataMethod_Name) String() string {
	return proto.EnumName(ReadStakingDataMethod_Name_name, int32(x))
}

func (ReadStakingDataMethod_Name) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{0, 0}
}

type ReadStakingDataMethod struct {
	Method               ReadStakingDataMethod_Name `protobuf:"varint,1,opt,name=method,proto3,enum=iotexapi.ReadStakingDataMethod_Name" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ReadStakingDataMethod) Reset()         { *m = ReadStakingDataMethod{} }
func (m *ReadStakingDataMethod) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataMethod) ProtoMessage()    {}
func (*ReadStakingDataMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{0}
}

func (m *ReadStakingDataMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataMethod.Unmarshal(m, b)
}
func (m *ReadStakingDataMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataMethod.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataMethod.Merge(m, src)
}
func (m *ReadStakingDataMethod) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataMethod.Size(m)
}
func (m *ReadStakingDataMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataMethod.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataMethod proto.InternalMessageInfo

func (m *ReadStakingDataMethod) GetMethod() ReadStakingDataMethod_Name {
	if m != nil {
		return m.Method
	}
	return ReadStakingDataMethod_INVALID
}

type ReadStakingDataRequest struct {
	// Types that are valid to be assigned to Request:
	//	*ReadStakingDataRequest_VoteBuckets_
	//	*ReadStakingDataRequest_BucketsByVoter
	//	*ReadStakingDataRequest_BucketsByCandidate
	//	*ReadStakingDataRequest_Candidates_
	//	*ReadStakingDataRequest_CandidateByName_
	Request              isReadStakingDataRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ReadStakingDataRequest) Reset()         { *m = ReadStakingDataRequest{} }
func (m *ReadStakingDataRequest) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest) ProtoMessage()    {}
func (*ReadStakingDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{1}
}

func (m *ReadStakingDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest.Merge(m, src)
}
func (m *ReadStakingDataRequest) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest.Size(m)
}
func (m *ReadStakingDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest proto.InternalMessageInfo

type isReadStakingDataRequest_Request interface {
	isReadStakingDataRequest_Request()
}

type ReadStakingDataRequest_VoteBuckets_ struct {
	VoteBuckets *ReadStakingDataRequest_VoteBuckets `protobuf:"bytes,1,opt,name=voteBuckets,proto3,oneof"`
}

type ReadStakingDataRequest_BucketsByVoter struct {
	BucketsByVoter *ReadStakingDataRequest_VoteBucketsByVoter `protobuf:"bytes,2,opt,name=bucketsByVoter,proto3,oneof"`
}

type ReadStakingDataRequest_BucketsByCandidate struct {
	BucketsByCandidate *ReadStakingDataRequest_VoteBucketsByCandidate `protobuf:"bytes,3,opt,name=bucketsByCandidate,proto3,oneof"`
}

type ReadStakingDataRequest_Candidates_ struct {
	Candidates *ReadStakingDataRequest_Candidates `protobuf:"bytes,4,opt,name=candidates,proto3,oneof"`
}

type ReadStakingDataRequest_CandidateByName_ struct {
	CandidateByName *ReadStakingDataRequest_CandidateByName `protobuf:"bytes,5,opt,name=candidateByName,proto3,oneof"`
}

func (*ReadStakingDataRequest_VoteBuckets_) isReadStakingDataRequest_Request() {}

func (*ReadStakingDataRequest_BucketsByVoter) isReadStakingDataRequest_Request() {}

func (*ReadStakingDataRequest_BucketsByCandidate) isReadStakingDataRequest_Request() {}

func (*ReadStakingDataRequest_Candidates_) isReadStakingDataRequest_Request() {}

func (*ReadStakingDataRequest_CandidateByName_) isReadStakingDataRequest_Request() {}

func (m *ReadStakingDataRequest) GetRequest() isReadStakingDataRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ReadStakingDataRequest) GetVoteBuckets() *ReadStakingDataRequest_VoteBuckets {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_VoteBuckets_); ok {
		return x.VoteBuckets
	}
	return nil
}

func (m *ReadStakingDataRequest) GetBucketsByVoter() *ReadStakingDataRequest_VoteBucketsByVoter {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_BucketsByVoter); ok {
		return x.BucketsByVoter
	}
	return nil
}

func (m *ReadStakingDataRequest) GetBucketsByCandidate() *ReadStakingDataRequest_VoteBucketsByCandidate {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_BucketsByCandidate); ok {
		return x.BucketsByCandidate
	}
	return nil
}

func (m *ReadStakingDataRequest) GetCandidates() *ReadStakingDataRequest_Candidates {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_Candidates_); ok {
		return x.Candidates
	}
	return nil
}

func (m *ReadStakingDataRequest) GetCandidateByName() *ReadStakingDataRequest_CandidateByName {
	if x, ok := m.GetRequest().(*ReadStakingDataRequest_CandidateByName_); ok {
		return x.CandidateByName
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReadStakingDataRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReadStakingDataRequest_VoteBuckets_)(nil),
		(*ReadStakingDataRequest_BucketsByVoter)(nil),
		(*ReadStakingDataRequest_BucketsByCandidate)(nil),
		(*ReadStakingDataRequest_Candidates_)(nil),
		(*ReadStakingDataRequest_CandidateByName_)(nil),
	}
}

type ReadStakingDataRequest_VoteBuckets struct {
	Offset               uint32   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingDataRequest_VoteBuckets) Reset()         { *m = ReadStakingDataRequest_VoteBuckets{} }
func (m *ReadStakingDataRequest_VoteBuckets) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest_VoteBuckets) ProtoMessage()    {}
func (*ReadStakingDataRequest_VoteBuckets) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{1, 0}
}

func (m *ReadStakingDataRequest_VoteBuckets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_VoteBuckets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_VoteBuckets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.Merge(m, src)
}
func (m *ReadStakingDataRequest_VoteBuckets) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.Size(m)
}
func (m *ReadStakingDataRequest_VoteBuckets) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_VoteBuckets.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_VoteBuckets proto.InternalMessageInfo

func (m *ReadStakingDataRequest_VoteBuckets) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadStakingDataRequest_VoteBuckets) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadStakingDataRequest_VoteBucketsByVoter struct {
	VoterAddress         string   `protobuf:"bytes,1,opt,name=voterAddress,proto3" json:"voterAddress,omitempty"`
	Offset               uint32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingDataRequest_VoteBucketsByVoter) Reset() {
	*m = ReadStakingDataRequest_VoteBucketsByVoter{}
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest_VoteBucketsByVoter) ProtoMessage()    {}
func (*ReadStakingDataRequest_VoteBucketsByVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{1, 1}
}

func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.Merge(m, src)
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.Size(m)
}
func (m *ReadStakingDataRequest_VoteBucketsByVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByVoter proto.InternalMessageInfo

func (m *ReadStakingDataRequest_VoteBucketsByVoter) GetVoterAddress() string {
	if m != nil {
		return m.VoterAddress
	}
	return ""
}

func (m *ReadStakingDataRequest_VoteBucketsByVoter) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadStakingDataRequest_VoteBucketsByVoter) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadStakingDataRequest_VoteBucketsByCandidate struct {
	CandName             string   `protobuf:"bytes,1,opt,name=candName,proto3" json:"candName,omitempty"`
	Offset               uint32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingDataRequest_VoteBucketsByCandidate) Reset() {
	*m = ReadStakingDataRequest_VoteBucketsByCandidate{}
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) String() string {
	return proto.CompactTextString(m)
}
func (*ReadStakingDataRequest_VoteBucketsByCandidate) ProtoMessage() {}
func (*ReadStakingDataRequest_VoteBucketsByCandidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{1, 2}
}

func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.Merge(m, src)
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.Size(m)
}
func (m *ReadStakingDataRequest_VoteBucketsByCandidate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_VoteBucketsByCandidate proto.InternalMessageInfo

func (m *ReadStakingDataRequest_VoteBucketsByCandidate) GetCandName() string {
	if m != nil {
		return m.CandName
	}
	return ""
}

func (m *ReadStakingDataRequest_VoteBucketsByCandidate) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadStakingDataRequest_VoteBucketsByCandidate) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadStakingDataRequest_Candidates struct {
	Offset               uint32   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingDataRequest_Candidates) Reset()         { *m = ReadStakingDataRequest_Candidates{} }
func (m *ReadStakingDataRequest_Candidates) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest_Candidates) ProtoMessage()    {}
func (*ReadStakingDataRequest_Candidates) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{1, 3}
}

func (m *ReadStakingDataRequest_Candidates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_Candidates.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_Candidates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_Candidates.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_Candidates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_Candidates.Merge(m, src)
}
func (m *ReadStakingDataRequest_Candidates) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_Candidates.Size(m)
}
func (m *ReadStakingDataRequest_Candidates) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_Candidates.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_Candidates proto.InternalMessageInfo

func (m *ReadStakingDataRequest_Candidates) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadStakingDataRequest_Candidates) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadStakingDataRequest_CandidateByName struct {
	CandName             string   `protobuf:"bytes,1,opt,name=candName,proto3" json:"candName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingDataRequest_CandidateByName) Reset() {
	*m = ReadStakingDataRequest_CandidateByName{}
}
func (m *ReadStakingDataRequest_CandidateByName) String() string { return proto.CompactTextString(m) }
func (*ReadStakingDataRequest_CandidateByName) ProtoMessage()    {}
func (*ReadStakingDataRequest_CandidateByName) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08103f271b0c8be, []int{1, 4}
}

func (m *ReadStakingDataRequest_CandidateByName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingDataRequest_CandidateByName.Unmarshal(m, b)
}
func (m *ReadStakingDataRequest_CandidateByName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingDataRequest_CandidateByName.Marshal(b, m, deterministic)
}
func (m *ReadStakingDataRequest_CandidateByName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingDataRequest_CandidateByName.Merge(m, src)
}
func (m *ReadStakingDataRequest_CandidateByName) XXX_Size() int {
	return xxx_messageInfo_ReadStakingDataRequest_CandidateByName.Size(m)
}
func (m *ReadStakingDataRequest_CandidateByName) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingDataRequest_CandidateByName.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingDataRequest_CandidateByName proto.InternalMessageInfo

func (m *ReadStakingDataRequest_CandidateByName) GetCandName() string {
	if m != nil {
		return m.CandName
	}
	return ""
}

func init() {
	proto.RegisterEnum("iotexapi.ReadStakingDataMethod_Name", ReadStakingDataMethod_Name_name, ReadStakingDataMethod_Name_value)
	proto.RegisterType((*ReadStakingDataMethod)(nil), "iotexapi.ReadStakingDataMethod")
	proto.RegisterType((*ReadStakingDataRequest)(nil), "iotexapi.ReadStakingDataRequest")
	proto.RegisterType((*ReadStakingDataRequest_VoteBuckets)(nil), "iotexapi.ReadStakingDataRequest.VoteBuckets")
	proto.RegisterType((*ReadStakingDataRequest_VoteBucketsByVoter)(nil), "iotexapi.ReadStakingDataRequest.VoteBucketsByVoter")
	proto.RegisterType((*ReadStakingDataRequest_VoteBucketsByCandidate)(nil), "iotexapi.ReadStakingDataRequest.VoteBucketsByCandidate")
	proto.RegisterType((*ReadStakingDataRequest_Candidates)(nil), "iotexapi.ReadStakingDataRequest.Candidates")
	proto.RegisterType((*ReadStakingDataRequest_CandidateByName)(nil), "iotexapi.ReadStakingDataRequest.CandidateByName")
}

func init() { proto.RegisterFile("proto/api/read_state.proto", fileDescriptor_a08103f271b0c8be) }

var fileDescriptor_a08103f271b0c8be = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0x6d, 0x97, 0x8f, 0xdd, 0xbd, 0x95, 0xa5, 0xde, 0xec, 0x92, 0x4a, 0x7c, 0x30, 0xc4, 0x07,
	0x13, 0xdd, 0x62, 0x96, 0x18, 0x13, 0xf5, 0xa5, 0x85, 0x46, 0x88, 0x82, 0x66, 0x16, 0x49, 0x34,
	0x1a, 0x32, 0xd0, 0x81, 0x1d, 0x77, 0x61, 0xb0, 0x1d, 0x8c, 0xfb, 0x17, 0xfc, 0x95, 0xc6, 0x5f,
	0x62, 0x3a, 0x85, 0xd2, 0xad, 0xf8, 0xc1, 0x5b, 0xcf, 0x19, 0xce, 0x39, 0x73, 0xef, 0x99, 0x00,
	0xd5, 0x45, 0x20, 0xa4, 0xa8, 0xd3, 0x05, 0xaf, 0x07, 0x8c, 0xfa, 0xc3, 0x50, 0x52, 0xc9, 0x6c,
	0x45, 0xe2, 0x01, 0x17, 0x92, 0x7d, 0xa3, 0x0b, 0x5e, 0xfb, 0xa9, 0xc3, 0x09, 0x61, 0xd4, 0x3f,
	0x97, 0xf4, 0x92, 0xcf, 0xa7, 0x2d, 0x2a, 0x69, 0x97, 0xc9, 0x0b, 0xe1, 0xe3, 0x0b, 0x28, 0xce,
	0xd4, 0x97, 0xa5, 0xdf, 0xd3, 0x1f, 0x1c, 0x9d, 0xdd, 0xb7, 0xd7, 0x22, 0x7b, 0xab, 0xc0, 0xee,
	0xd1, 0x19, 0x23, 0x2b, 0x4d, 0xed, 0xbb, 0x0e, 0xf9, 0x88, 0x40, 0x03, 0xf6, 0x3b, 0xbd, 0x81,
	0xf3, 0xba, 0xd3, 0x32, 0x35, 0x2c, 0x83, 0xf1, 0xd2, 0xeb, 0x0f, 0xdd, 0x77, 0xcd, 0x57, 0x5e,
	0xff, 0xdc, 0xd4, 0xd1, 0x82, 0xe3, 0x14, 0x31, 0x74, 0xdf, 0x0f, 0x07, 0x6f, 0xfa, 0x1e, 0x31,
	0xf7, 0xf0, 0x2e, 0x58, 0x99, 0x93, 0xa6, 0xd3, 0x6b, 0x75, 0x5a, 0x4e, 0xdf, 0x33, 0x73, 0x78,
	0x1b, 0x4a, 0xd1, 0xe9, 0x86, 0xca, 0xe3, 0x1d, 0x38, 0xb9, 0x41, 0x45, 0x92, 0x9e, 0xd3, 0xf5,
	0xcc, 0x42, 0xed, 0x47, 0x11, 0x2a, 0x99, 0x3b, 0x13, 0xf6, 0x65, 0xc9, 0x42, 0x89, 0x6f, 0xc1,
	0xf8, 0x2a, 0x24, 0x73, 0x97, 0xe3, 0x4b, 0x26, 0x43, 0x35, 0xaa, 0x71, 0xf6, 0xe8, 0x8f, 0xa3,
	0xae, 0x64, 0xf6, 0x60, 0xa3, 0x69, 0x6b, 0x24, 0x6d, 0x81, 0x9f, 0xe0, 0x68, 0x14, 0x7f, 0xba,
	0xd7, 0xd1, 0xcf, 0x02, 0x6b, 0x4f, 0x99, 0x36, 0x76, 0x31, 0x5d, 0x49, 0xdb, 0x1a, 0xc9, 0x98,
	0x21, 0x07, 0x4c, 0x98, 0x26, 0x9d, 0xfb, 0xdc, 0xa7, 0x92, 0x59, 0x39, 0x15, 0xf1, 0x74, 0xb7,
	0x88, 0x44, 0xde, 0xd6, 0xc8, 0x16, 0x53, 0xec, 0x02, 0x8c, 0xd7, 0x20, 0xb4, 0xf2, 0x2a, 0xe2,
	0xe1, 0x3f, 0x23, 0x12, 0x7d, 0xb4, 0x99, 0x94, 0x01, 0x7e, 0x84, 0x72, 0x82, 0xdc, 0xeb, 0xe8,
	0x71, 0x58, 0x05, 0xe5, 0xf9, 0xf8, 0xff, 0x3d, 0x63, 0x5d, 0x5b, 0x23, 0x59, 0xab, 0xea, 0x73,
	0x30, 0x52, 0xc3, 0x61, 0x05, 0x8a, 0x62, 0x32, 0x09, 0x99, 0x54, 0x95, 0x96, 0xc8, 0x0a, 0xe1,
	0x31, 0x14, 0xae, 0xf8, 0x8c, 0x4b, 0x55, 0x4a, 0x89, 0xc4, 0xa0, 0x3a, 0x01, 0xfc, 0x7d, 0xf9,
	0x58, 0x83, 0x5b, 0x51, 0xb1, 0x81, 0xe3, 0xfb, 0x01, 0x0b, 0xe3, 0xc7, 0x71, 0x48, 0x6e, 0x70,
	0xa9, 0x9c, 0xbd, 0xed, 0x39, 0xb9, 0x74, 0xce, 0x08, 0x2a, 0xdb, 0x1b, 0xc0, 0x2a, 0x1c, 0x44,
	0x13, 0xa9, 0xad, 0xc4, 0x39, 0x09, 0xde, 0x31, 0xe3, 0x19, 0xc0, 0xa6, 0x82, 0x1d, 0xf7, 0x70,
	0x0a, 0xe5, 0xcc, 0xaa, 0xff, 0x76, 0x31, 0xf7, 0x10, 0xf6, 0x83, 0xb8, 0x21, 0xf7, 0xc9, 0x87,
	0xc6, 0x94, 0xcb, 0x8b, 0xe5, 0xc8, 0x1e, 0x8b, 0x59, 0x5d, 0xf5, 0xb9, 0x08, 0xc4, 0x67, 0x36,
	0x96, 0x31, 0x38, 0x8d, 0xff, 0x8d, 0xa6, 0xe2, 0x8a, 0xce, 0xa7, 0xf5, 0x75, 0xdf, 0xa3, 0xa2,
	0xa2, 0x1b, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x71, 0xc7, 0x60, 0xad, 0x04, 0x00, 0x00,
}