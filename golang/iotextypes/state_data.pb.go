// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/state_data.proto

package iotextypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type KickoutInfo struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Count                uint32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KickoutInfo) Reset()         { *m = KickoutInfo{} }
func (m *KickoutInfo) String() string { return proto.CompactTextString(m) }
func (*KickoutInfo) ProtoMessage()    {}
func (*KickoutInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e4d0bfa2d89671, []int{0}
}

func (m *KickoutInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickoutInfo.Unmarshal(m, b)
}
func (m *KickoutInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickoutInfo.Marshal(b, m, deterministic)
}
func (m *KickoutInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickoutInfo.Merge(m, src)
}
func (m *KickoutInfo) XXX_Size() int {
	return xxx_messageInfo_KickoutInfo.Size(m)
}
func (m *KickoutInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KickoutInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KickoutInfo proto.InternalMessageInfo

func (m *KickoutInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *KickoutInfo) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// KickoutInfo and list of KickoutCandidateList
type KickoutCandidateList struct {
	Blacklists           []*KickoutInfo `protobuf:"bytes,1,rep,name=blacklists,proto3" json:"blacklists,omitempty"`
	IntensityRate        uint32         `protobuf:"varint,2,opt,name=intensityRate,proto3" json:"intensityRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KickoutCandidateList) Reset()         { *m = KickoutCandidateList{} }
func (m *KickoutCandidateList) String() string { return proto.CompactTextString(m) }
func (*KickoutCandidateList) ProtoMessage()    {}
func (*KickoutCandidateList) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e4d0bfa2d89671, []int{1}
}

func (m *KickoutCandidateList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickoutCandidateList.Unmarshal(m, b)
}
func (m *KickoutCandidateList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickoutCandidateList.Marshal(b, m, deterministic)
}
func (m *KickoutCandidateList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickoutCandidateList.Merge(m, src)
}
func (m *KickoutCandidateList) XXX_Size() int {
	return xxx_messageInfo_KickoutCandidateList.Size(m)
}
func (m *KickoutCandidateList) XXX_DiscardUnknown() {
	xxx_messageInfo_KickoutCandidateList.DiscardUnknown(m)
}

var xxx_messageInfo_KickoutCandidateList proto.InternalMessageInfo

func (m *KickoutCandidateList) GetBlacklists() []*KickoutInfo {
	if m != nil {
		return m.Blacklists
	}
	return nil
}

func (m *KickoutCandidateList) GetIntensityRate() uint32 {
	if m != nil {
		return m.IntensityRate
	}
	return 0
}

type VoteBucket struct {
	CandidateAddress     string               `protobuf:"bytes,1,opt,name=candidateAddress,proto3" json:"candidateAddress,omitempty"`
	StakedAmount         string               `protobuf:"bytes,2,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
	StakedDuration       uint32               `protobuf:"varint,3,opt,name=stakedDuration,proto3" json:"stakedDuration,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	StakeStartTime       *timestamp.Timestamp `protobuf:"bytes,5,opt,name=stakeStartTime,proto3" json:"stakeStartTime,omitempty"`
	UnstakeStartTime     *timestamp.Timestamp `protobuf:"bytes,6,opt,name=unstakeStartTime,proto3" json:"unstakeStartTime,omitempty"`
	AutoStake            bool                 `protobuf:"varint,7,opt,name=autoStake,proto3" json:"autoStake,omitempty"`
	Owner                string               `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VoteBucket) Reset()         { *m = VoteBucket{} }
func (m *VoteBucket) String() string { return proto.CompactTextString(m) }
func (*VoteBucket) ProtoMessage()    {}
func (*VoteBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e4d0bfa2d89671, []int{2}
}

func (m *VoteBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteBucket.Unmarshal(m, b)
}
func (m *VoteBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteBucket.Marshal(b, m, deterministic)
}
func (m *VoteBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteBucket.Merge(m, src)
}
func (m *VoteBucket) XXX_Size() int {
	return xxx_messageInfo_VoteBucket.Size(m)
}
func (m *VoteBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteBucket.DiscardUnknown(m)
}

var xxx_messageInfo_VoteBucket proto.InternalMessageInfo

func (m *VoteBucket) GetCandidateAddress() string {
	if m != nil {
		return m.CandidateAddress
	}
	return ""
}

func (m *VoteBucket) GetStakedAmount() string {
	if m != nil {
		return m.StakedAmount
	}
	return ""
}

func (m *VoteBucket) GetStakedDuration() uint32 {
	if m != nil {
		return m.StakedDuration
	}
	return 0
}

func (m *VoteBucket) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *VoteBucket) GetStakeStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StakeStartTime
	}
	return nil
}

func (m *VoteBucket) GetUnstakeStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.UnstakeStartTime
	}
	return nil
}

func (m *VoteBucket) GetAutoStake() bool {
	if m != nil {
		return m.AutoStake
	}
	return false
}

func (m *VoteBucket) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type CandidateV2 struct {
	OwnerAddress         string   `protobuf:"bytes,1,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	OperatorAddress      string   `protobuf:"bytes,2,opt,name=operatorAddress,proto3" json:"operatorAddress,omitempty"`
	RewardAddress        string   `protobuf:"bytes,3,opt,name=rewardAddress,proto3" json:"rewardAddress,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	TotalWeightedVotes   string   `protobuf:"bytes,5,opt,name=totalWeightedVotes,proto3" json:"totalWeightedVotes,omitempty"`
	SelfStakeBucketIdx   uint64   `protobuf:"varint,6,opt,name=selfStakeBucketIdx,proto3" json:"selfStakeBucketIdx,omitempty"`
	SelfStakingTokens    string   `protobuf:"bytes,7,opt,name=selfStakingTokens,proto3" json:"selfStakingTokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CandidateV2) Reset()         { *m = CandidateV2{} }
func (m *CandidateV2) String() string { return proto.CompactTextString(m) }
func (*CandidateV2) ProtoMessage()    {}
func (*CandidateV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e4d0bfa2d89671, []int{3}
}

func (m *CandidateV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CandidateV2.Unmarshal(m, b)
}
func (m *CandidateV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CandidateV2.Marshal(b, m, deterministic)
}
func (m *CandidateV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CandidateV2.Merge(m, src)
}
func (m *CandidateV2) XXX_Size() int {
	return xxx_messageInfo_CandidateV2.Size(m)
}
func (m *CandidateV2) XXX_DiscardUnknown() {
	xxx_messageInfo_CandidateV2.DiscardUnknown(m)
}

var xxx_messageInfo_CandidateV2 proto.InternalMessageInfo

func (m *CandidateV2) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *CandidateV2) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *CandidateV2) GetRewardAddress() string {
	if m != nil {
		return m.RewardAddress
	}
	return ""
}

func (m *CandidateV2) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CandidateV2) GetTotalWeightedVotes() string {
	if m != nil {
		return m.TotalWeightedVotes
	}
	return ""
}

func (m *CandidateV2) GetSelfStakeBucketIdx() uint64 {
	if m != nil {
		return m.SelfStakeBucketIdx
	}
	return 0
}

func (m *CandidateV2) GetSelfStakingTokens() string {
	if m != nil {
		return m.SelfStakingTokens
	}
	return ""
}

type CandidateListV2 struct {
	Candidates           []*CandidateV2 `protobuf:"bytes,1,rep,name=candidates,proto3" json:"candidates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CandidateListV2) Reset()         { *m = CandidateListV2{} }
func (m *CandidateListV2) String() string { return proto.CompactTextString(m) }
func (*CandidateListV2) ProtoMessage()    {}
func (*CandidateListV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e4d0bfa2d89671, []int{4}
}

func (m *CandidateListV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CandidateListV2.Unmarshal(m, b)
}
func (m *CandidateListV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CandidateListV2.Marshal(b, m, deterministic)
}
func (m *CandidateListV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CandidateListV2.Merge(m, src)
}
func (m *CandidateListV2) XXX_Size() int {
	return xxx_messageInfo_CandidateListV2.Size(m)
}
func (m *CandidateListV2) XXX_DiscardUnknown() {
	xxx_messageInfo_CandidateListV2.DiscardUnknown(m)
}

var xxx_messageInfo_CandidateListV2 proto.InternalMessageInfo

func (m *CandidateListV2) GetCandidates() []*CandidateV2 {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type VoteBucketList struct {
	Buckets              []*VoteBucket `protobuf:"bytes,1,rep,name=buckets,proto3" json:"buckets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VoteBucketList) Reset()         { *m = VoteBucketList{} }
func (m *VoteBucketList) String() string { return proto.CompactTextString(m) }
func (*VoteBucketList) ProtoMessage()    {}
func (*VoteBucketList) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e4d0bfa2d89671, []int{5}
}

func (m *VoteBucketList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteBucketList.Unmarshal(m, b)
}
func (m *VoteBucketList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteBucketList.Marshal(b, m, deterministic)
}
func (m *VoteBucketList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteBucketList.Merge(m, src)
}
func (m *VoteBucketList) XXX_Size() int {
	return xxx_messageInfo_VoteBucketList.Size(m)
}
func (m *VoteBucketList) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteBucketList.DiscardUnknown(m)
}

var xxx_messageInfo_VoteBucketList proto.InternalMessageInfo

func (m *VoteBucketList) GetBuckets() []*VoteBucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func init() {
	proto.RegisterType((*KickoutInfo)(nil), "iotextypes.KickoutInfo")
	proto.RegisterType((*KickoutCandidateList)(nil), "iotextypes.KickoutCandidateList")
	proto.RegisterType((*VoteBucket)(nil), "iotextypes.VoteBucket")
	proto.RegisterType((*CandidateV2)(nil), "iotextypes.CandidateV2")
	proto.RegisterType((*CandidateListV2)(nil), "iotextypes.CandidateListV2")
	proto.RegisterType((*VoteBucketList)(nil), "iotextypes.VoteBucketList")
}

func init() { proto.RegisterFile("proto/types/state_data.proto", fileDescriptor_96e4d0bfa2d89671) }

var fileDescriptor_96e4d0bfa2d89671 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0xb1, 0x13, 0x47, 0xe3, 0xe6, 0xa3, 0x4b, 0x68, 0x45, 0x08, 0xd4, 0x88, 0x52, 0x44,
	0x69, 0xa5, 0xe2, 0x52, 0x02, 0x85, 0x1e, 0xe2, 0x96, 0x42, 0xda, 0x9e, 0x36, 0xc6, 0x85, 0x5e,
	0xca, 0x5a, 0x1a, 0x2b, 0x5b, 0xcb, 0xbb, 0x46, 0x3b, 0x22, 0xc9, 0xad, 0xff, 0xa5, 0x7f, 0xb4,
	0x68, 0x65, 0xd9, 0x92, 0x6d, 0xc8, 0x4d, 0xf3, 0xde, 0x9b, 0x99, 0xdd, 0xb7, 0x33, 0x82, 0x8b,
	0x45, 0xa6, 0x49, 0x87, 0xf4, 0xb0, 0x40, 0x13, 0x1a, 0x12, 0x84, 0xbf, 0x63, 0x41, 0x22, 0xb0,
	0x30, 0x03, 0xa9, 0x09, 0xef, 0x2d, 0x79, 0xfe, 0x22, 0xd1, 0x3a, 0x49, 0x31, 0xb4, 0xcc, 0x24,
	0x9f, 0x86, 0x24, 0xe7, 0x68, 0x48, 0xcc, 0x17, 0xa5, 0xd8, 0xfb, 0x04, 0xbd, 0xef, 0x32, 0x9a,
	0xe9, 0x9c, 0xae, 0xd5, 0x54, 0x33, 0x17, 0xba, 0x22, 0x8e, 0x33, 0x34, 0xc6, 0x6d, 0xf5, 0x5b,
	0xbe, 0xc3, 0xab, 0x90, 0x9d, 0xc1, 0x7e, 0xa4, 0x73, 0x45, 0xee, 0x5e, 0xbf, 0xe5, 0x1f, 0xf1,
	0x32, 0xf0, 0x72, 0x38, 0x5b, 0xa6, 0x7f, 0x16, 0x2a, 0x96, 0xb1, 0x20, 0xfc, 0x21, 0x0d, 0xb1,
	0x4b, 0x80, 0x49, 0x2a, 0xa2, 0x59, 0x2a, 0x0d, 0x15, 0xa5, 0xda, 0x7e, 0x6f, 0xf0, 0x3c, 0x58,
	0x1f, 0x2c, 0xa8, 0x35, 0xe5, 0x35, 0x29, 0x7b, 0x09, 0x47, 0x52, 0x11, 0x2a, 0x23, 0xe9, 0x81,
	0x0b, 0xc2, 0x65, 0xbb, 0x26, 0xe8, 0xfd, 0x6d, 0x03, 0x8c, 0x35, 0xe1, 0x30, 0x8f, 0x66, 0x48,
	0xec, 0x35, 0x9c, 0x46, 0x55, 0xfb, 0xab, 0xc6, 0xf1, 0xb7, 0x70, 0xe6, 0xc1, 0x13, 0x43, 0x62,
	0x86, 0xf1, 0xd5, 0x7c, 0x75, 0x1d, 0x87, 0x37, 0x30, 0xf6, 0x0a, 0x8e, 0xcb, 0xf8, 0x4b, 0x9e,
	0x09, 0x92, 0x5a, 0xb9, 0x6d, 0x7b, 0x8a, 0x0d, 0x94, 0x7d, 0x04, 0x88, 0x32, 0x14, 0x84, 0x23,
	0x39, 0x47, 0xb7, 0xd3, 0x6f, 0xf9, 0xbd, 0xc1, 0x79, 0x50, 0x5a, 0x1e, 0x54, 0x96, 0x07, 0xa3,
	0xca, 0x72, 0x5e, 0x53, 0xb3, 0xe1, 0xb2, 0xc7, 0x0d, 0x89, 0x8c, 0x6c, 0xfe, 0xfe, 0xa3, 0xf9,
	0x1b, 0x19, 0xec, 0x2b, 0x9c, 0xe6, 0x6a, 0xa3, 0xca, 0xc1, 0xa3, 0x55, 0xb6, 0x72, 0xd8, 0x05,
	0x38, 0x22, 0x27, 0x7d, 0x53, 0xa0, 0x6e, 0xb7, 0xdf, 0xf2, 0x0f, 0xf9, 0x1a, 0x28, 0x5e, 0x5e,
	0xdf, 0x29, 0xcc, 0xdc, 0x43, 0x6b, 0x55, 0x19, 0x78, 0xff, 0xf6, 0xa0, 0xb7, 0x7a, 0xf3, 0xf1,
	0xa0, 0xf0, 0xd5, 0x12, 0x4d, 0xff, 0x1b, 0x18, 0xf3, 0xe1, 0x44, 0x2f, 0x30, 0x13, 0xa4, 0x57,
	0xb2, 0xd2, 0xfe, 0x4d, 0xb8, 0x18, 0x83, 0x0c, 0xef, 0x44, 0x16, 0x57, 0xba, 0xb6, 0xd5, 0x35,
	0x41, 0xc6, 0xa0, 0xa3, 0xc4, 0xd2, 0x79, 0x87, 0xdb, 0x6f, 0x16, 0x00, 0x23, 0x4d, 0x22, 0xfd,
	0x89, 0x32, 0xb9, 0x25, 0x8c, 0x8b, 0x31, 0x31, 0xd6, 0x5b, 0x87, 0xef, 0x60, 0x0a, 0xbd, 0xc1,
	0x74, 0x6a, 0xaf, 0x5a, 0x8e, 0xd3, 0x75, 0x7c, 0x6f, 0x5d, 0xec, 0xf0, 0x1d, 0x0c, 0x7b, 0x03,
	0x4f, 0x2b, 0x54, 0xaa, 0x64, 0xa4, 0x67, 0xa8, 0x8c, 0xf5, 0xcc, 0xe1, 0xdb, 0x84, 0xf7, 0x0d,
	0x4e, 0x1a, 0x8b, 0x31, 0x1e, 0x14, 0xab, 0xb1, 0x1a, 0xca, 0x9d, 0xab, 0x51, 0x73, 0x95, 0xd7,
	0xa4, 0xde, 0x10, 0x8e, 0xd7, 0x33, 0x6f, 0xb7, 0xec, 0x1d, 0x74, 0x27, 0x36, 0xaa, 0xea, 0x3c,
	0xab, 0xd7, 0x59, 0x8b, 0x79, 0x25, 0x1b, 0x5e, 0xfe, 0xfa, 0x90, 0x48, 0xba, 0xcd, 0x27, 0x41,
	0xa4, 0xe7, 0xa1, 0x15, 0x2f, 0x32, 0xfd, 0x07, 0x23, 0x2a, 0x83, 0xb7, 0xe5, 0x9f, 0x25, 0xd1,
	0xa9, 0x50, 0x49, 0xb8, 0x2e, 0x36, 0x39, 0xb0, 0xc4, 0xfb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x20, 0xc9, 0x47, 0xc1, 0x7b, 0x04, 0x00, 0x00,
}