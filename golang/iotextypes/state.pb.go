// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/state.proto

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
	return fileDescriptor_d165ee162385f520, []int{0}
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
	return fileDescriptor_d165ee162385f520, []int{1}
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
	return fileDescriptor_d165ee162385f520, []int{2}
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
	return fileDescriptor_d165ee162385f520, []int{3}
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
	return fileDescriptor_d165ee162385f520, []int{4}
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
	return fileDescriptor_d165ee162385f520, []int{5}
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

type ReadStakingData struct {
	// Types that are valid to be assigned to Request:
	//	*ReadStakingData_GetBuckets
	//	*ReadStakingData_GetBucketsByVoter
	//	*ReadStakingData_GetBucketsByCandidate
	//	*ReadStakingData_GetCandidates
	//	*ReadStakingData_GetCandidateByName
	Request              isReadStakingData_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ReadStakingData) Reset()         { *m = ReadStakingData{} }
func (m *ReadStakingData) String() string { return proto.CompactTextString(m) }
func (*ReadStakingData) ProtoMessage()    {}
func (*ReadStakingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d165ee162385f520, []int{6}
}

func (m *ReadStakingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingData.Unmarshal(m, b)
}
func (m *ReadStakingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingData.Marshal(b, m, deterministic)
}
func (m *ReadStakingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingData.Merge(m, src)
}
func (m *ReadStakingData) XXX_Size() int {
	return xxx_messageInfo_ReadStakingData.Size(m)
}
func (m *ReadStakingData) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingData.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingData proto.InternalMessageInfo

type isReadStakingData_Request interface {
	isReadStakingData_Request()
}

type ReadStakingData_GetBuckets struct {
	GetBuckets *ReadStakingVoteBuckets `protobuf:"bytes,1,opt,name=getBuckets,proto3,oneof"`
}

type ReadStakingData_GetBucketsByVoter struct {
	GetBucketsByVoter *ReadStakingVoteBucketsByVoter `protobuf:"bytes,2,opt,name=getBucketsByVoter,proto3,oneof"`
}

type ReadStakingData_GetBucketsByCandidate struct {
	GetBucketsByCandidate *ReadStakingVoteBucketsByCandidate `protobuf:"bytes,3,opt,name=getBucketsByCandidate,proto3,oneof"`
}

type ReadStakingData_GetCandidates struct {
	GetCandidates *ReadStakingCandidates `protobuf:"bytes,4,opt,name=getCandidates,proto3,oneof"`
}

type ReadStakingData_GetCandidateByName struct {
	GetCandidateByName *ReadStakingCandidateByName `protobuf:"bytes,5,opt,name=getCandidateByName,proto3,oneof"`
}

func (*ReadStakingData_GetBuckets) isReadStakingData_Request() {}

func (*ReadStakingData_GetBucketsByVoter) isReadStakingData_Request() {}

func (*ReadStakingData_GetBucketsByCandidate) isReadStakingData_Request() {}

func (*ReadStakingData_GetCandidates) isReadStakingData_Request() {}

func (*ReadStakingData_GetCandidateByName) isReadStakingData_Request() {}

func (m *ReadStakingData) GetRequest() isReadStakingData_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ReadStakingData) GetGetBuckets() *ReadStakingVoteBuckets {
	if x, ok := m.GetRequest().(*ReadStakingData_GetBuckets); ok {
		return x.GetBuckets
	}
	return nil
}

func (m *ReadStakingData) GetGetBucketsByVoter() *ReadStakingVoteBucketsByVoter {
	if x, ok := m.GetRequest().(*ReadStakingData_GetBucketsByVoter); ok {
		return x.GetBucketsByVoter
	}
	return nil
}

func (m *ReadStakingData) GetGetBucketsByCandidate() *ReadStakingVoteBucketsByCandidate {
	if x, ok := m.GetRequest().(*ReadStakingData_GetBucketsByCandidate); ok {
		return x.GetBucketsByCandidate
	}
	return nil
}

func (m *ReadStakingData) GetGetCandidates() *ReadStakingCandidates {
	if x, ok := m.GetRequest().(*ReadStakingData_GetCandidates); ok {
		return x.GetCandidates
	}
	return nil
}

func (m *ReadStakingData) GetGetCandidateByName() *ReadStakingCandidateByName {
	if x, ok := m.GetRequest().(*ReadStakingData_GetCandidateByName); ok {
		return x.GetCandidateByName
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReadStakingData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReadStakingData_GetBuckets)(nil),
		(*ReadStakingData_GetBucketsByVoter)(nil),
		(*ReadStakingData_GetBucketsByCandidate)(nil),
		(*ReadStakingData_GetCandidates)(nil),
		(*ReadStakingData_GetCandidateByName)(nil),
	}
}

type ReadStakingVoteBuckets struct {
	Offset               uint32   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingVoteBuckets) Reset()         { *m = ReadStakingVoteBuckets{} }
func (m *ReadStakingVoteBuckets) String() string { return proto.CompactTextString(m) }
func (*ReadStakingVoteBuckets) ProtoMessage()    {}
func (*ReadStakingVoteBuckets) Descriptor() ([]byte, []int) {
	return fileDescriptor_d165ee162385f520, []int{7}
}

func (m *ReadStakingVoteBuckets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingVoteBuckets.Unmarshal(m, b)
}
func (m *ReadStakingVoteBuckets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingVoteBuckets.Marshal(b, m, deterministic)
}
func (m *ReadStakingVoteBuckets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingVoteBuckets.Merge(m, src)
}
func (m *ReadStakingVoteBuckets) XXX_Size() int {
	return xxx_messageInfo_ReadStakingVoteBuckets.Size(m)
}
func (m *ReadStakingVoteBuckets) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingVoteBuckets.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingVoteBuckets proto.InternalMessageInfo

func (m *ReadStakingVoteBuckets) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadStakingVoteBuckets) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadStakingVoteBucketsByVoter struct {
	VoterAddress         string   `protobuf:"bytes,1,opt,name=voterAddress,proto3" json:"voterAddress,omitempty"`
	Offset               uint32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingVoteBucketsByVoter) Reset()         { *m = ReadStakingVoteBucketsByVoter{} }
func (m *ReadStakingVoteBucketsByVoter) String() string { return proto.CompactTextString(m) }
func (*ReadStakingVoteBucketsByVoter) ProtoMessage()    {}
func (*ReadStakingVoteBucketsByVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d165ee162385f520, []int{8}
}

func (m *ReadStakingVoteBucketsByVoter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingVoteBucketsByVoter.Unmarshal(m, b)
}
func (m *ReadStakingVoteBucketsByVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingVoteBucketsByVoter.Marshal(b, m, deterministic)
}
func (m *ReadStakingVoteBucketsByVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingVoteBucketsByVoter.Merge(m, src)
}
func (m *ReadStakingVoteBucketsByVoter) XXX_Size() int {
	return xxx_messageInfo_ReadStakingVoteBucketsByVoter.Size(m)
}
func (m *ReadStakingVoteBucketsByVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingVoteBucketsByVoter.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingVoteBucketsByVoter proto.InternalMessageInfo

func (m *ReadStakingVoteBucketsByVoter) GetVoterAddress() string {
	if m != nil {
		return m.VoterAddress
	}
	return ""
}

func (m *ReadStakingVoteBucketsByVoter) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadStakingVoteBucketsByVoter) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadStakingVoteBucketsByCandidate struct {
	CandName             string   `protobuf:"bytes,1,opt,name=candName,proto3" json:"candName,omitempty"`
	Offset               uint32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingVoteBucketsByCandidate) Reset()         { *m = ReadStakingVoteBucketsByCandidate{} }
func (m *ReadStakingVoteBucketsByCandidate) String() string { return proto.CompactTextString(m) }
func (*ReadStakingVoteBucketsByCandidate) ProtoMessage()    {}
func (*ReadStakingVoteBucketsByCandidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d165ee162385f520, []int{9}
}

func (m *ReadStakingVoteBucketsByCandidate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingVoteBucketsByCandidate.Unmarshal(m, b)
}
func (m *ReadStakingVoteBucketsByCandidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingVoteBucketsByCandidate.Marshal(b, m, deterministic)
}
func (m *ReadStakingVoteBucketsByCandidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingVoteBucketsByCandidate.Merge(m, src)
}
func (m *ReadStakingVoteBucketsByCandidate) XXX_Size() int {
	return xxx_messageInfo_ReadStakingVoteBucketsByCandidate.Size(m)
}
func (m *ReadStakingVoteBucketsByCandidate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingVoteBucketsByCandidate.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingVoteBucketsByCandidate proto.InternalMessageInfo

func (m *ReadStakingVoteBucketsByCandidate) GetCandName() string {
	if m != nil {
		return m.CandName
	}
	return ""
}

func (m *ReadStakingVoteBucketsByCandidate) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadStakingVoteBucketsByCandidate) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadStakingCandidates struct {
	Offset               uint32   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingCandidates) Reset()         { *m = ReadStakingCandidates{} }
func (m *ReadStakingCandidates) String() string { return proto.CompactTextString(m) }
func (*ReadStakingCandidates) ProtoMessage()    {}
func (*ReadStakingCandidates) Descriptor() ([]byte, []int) {
	return fileDescriptor_d165ee162385f520, []int{10}
}

func (m *ReadStakingCandidates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingCandidates.Unmarshal(m, b)
}
func (m *ReadStakingCandidates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingCandidates.Marshal(b, m, deterministic)
}
func (m *ReadStakingCandidates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingCandidates.Merge(m, src)
}
func (m *ReadStakingCandidates) XXX_Size() int {
	return xxx_messageInfo_ReadStakingCandidates.Size(m)
}
func (m *ReadStakingCandidates) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingCandidates.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingCandidates proto.InternalMessageInfo

func (m *ReadStakingCandidates) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadStakingCandidates) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadStakingCandidateByName struct {
	CandName             string   `protobuf:"bytes,1,opt,name=candName,proto3" json:"candName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadStakingCandidateByName) Reset()         { *m = ReadStakingCandidateByName{} }
func (m *ReadStakingCandidateByName) String() string { return proto.CompactTextString(m) }
func (*ReadStakingCandidateByName) ProtoMessage()    {}
func (*ReadStakingCandidateByName) Descriptor() ([]byte, []int) {
	return fileDescriptor_d165ee162385f520, []int{11}
}

func (m *ReadStakingCandidateByName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadStakingCandidateByName.Unmarshal(m, b)
}
func (m *ReadStakingCandidateByName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadStakingCandidateByName.Marshal(b, m, deterministic)
}
func (m *ReadStakingCandidateByName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStakingCandidateByName.Merge(m, src)
}
func (m *ReadStakingCandidateByName) XXX_Size() int {
	return xxx_messageInfo_ReadStakingCandidateByName.Size(m)
}
func (m *ReadStakingCandidateByName) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStakingCandidateByName.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStakingCandidateByName proto.InternalMessageInfo

func (m *ReadStakingCandidateByName) GetCandName() string {
	if m != nil {
		return m.CandName
	}
	return ""
}

func init() {
	proto.RegisterType((*KickoutInfo)(nil), "iotextypes.KickoutInfo")
	proto.RegisterType((*KickoutCandidateList)(nil), "iotextypes.KickoutCandidateList")
	proto.RegisterType((*VoteBucket)(nil), "iotextypes.VoteBucket")
	proto.RegisterType((*CandidateV2)(nil), "iotextypes.CandidateV2")
	proto.RegisterType((*CandidateListV2)(nil), "iotextypes.CandidateListV2")
	proto.RegisterType((*VoteBucketList)(nil), "iotextypes.VoteBucketList")
	proto.RegisterType((*ReadStakingData)(nil), "iotextypes.ReadStakingData")
	proto.RegisterType((*ReadStakingVoteBuckets)(nil), "iotextypes.ReadStakingVoteBuckets")
	proto.RegisterType((*ReadStakingVoteBucketsByVoter)(nil), "iotextypes.ReadStakingVoteBucketsByVoter")
	proto.RegisterType((*ReadStakingVoteBucketsByCandidate)(nil), "iotextypes.ReadStakingVoteBucketsByCandidate")
	proto.RegisterType((*ReadStakingCandidates)(nil), "iotextypes.ReadStakingCandidates")
	proto.RegisterType((*ReadStakingCandidateByName)(nil), "iotextypes.ReadStakingCandidateByName")
}

func init() { proto.RegisterFile("proto/types/state.proto", fileDescriptor_d165ee162385f520) }

var fileDescriptor_d165ee162385f520 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdb, 0x6e, 0x1a, 0x49,
	0x10, 0x05, 0x63, 0x1b, 0x53, 0x2c, 0xbe, 0xb4, 0x7c, 0x41, 0x68, 0x57, 0x8b, 0x47, 0x2b, 0x8b,
	0x5d, 0xad, 0x87, 0x15, 0xab, 0xc8, 0x51, 0xa4, 0x3c, 0x78, 0xe2, 0x58, 0x38, 0x89, 0xf2, 0xd0,
	0xb6, 0x9c, 0xcb, 0x5b, 0x33, 0xd3, 0x8c, 0x27, 0xc0, 0x34, 0x9e, 0xae, 0x89, 0xcd, 0x5b, 0xfe,
	0x25, 0x5f, 0x93, 0xbf, 0x8a, 0xa6, 0xe7, 0x0e, 0x63, 0x13, 0xbf, 0x51, 0xa7, 0x4f, 0x9d, 0xea,
	0x3e, 0x55, 0x35, 0xc0, 0xc1, 0xd4, 0x13, 0x28, 0xba, 0x38, 0x9b, 0x72, 0xd9, 0x95, 0xc8, 0x90,
	0xeb, 0x0a, 0x21, 0xe0, 0x08, 0xe4, 0xf7, 0x0a, 0x6f, 0xfd, 0x69, 0x0b, 0x61, 0x8f, 0x79, 0x57,
	0x9d, 0x0c, 0xfc, 0x61, 0x17, 0x9d, 0x09, 0x97, 0xc8, 0x26, 0xd3, 0x90, 0xac, 0xbd, 0x84, 0xfa,
	0x5b, 0xc7, 0x1c, 0x09, 0x1f, 0x2f, 0xdc, 0xa1, 0x20, 0x4d, 0xa8, 0x32, 0xcb, 0xf2, 0xb8, 0x94,
	0xcd, 0x72, 0xbb, 0xdc, 0xa9, 0xd1, 0x38, 0x24, 0xbb, 0xb0, 0x66, 0x0a, 0xdf, 0xc5, 0xe6, 0x4a,
	0xbb, 0xdc, 0x69, 0xd0, 0x30, 0xd0, 0x7c, 0xd8, 0x8d, 0xd2, 0x5f, 0x31, 0xd7, 0x72, 0x2c, 0x86,
	0xfc, 0x9d, 0x23, 0x91, 0x9c, 0x00, 0x0c, 0xc6, 0xcc, 0x1c, 0x8d, 0x1d, 0x89, 0x81, 0x54, 0xa5,
	0x53, 0xef, 0x1d, 0xe8, 0xe9, 0xc5, 0xf4, 0x4c, 0x51, 0x9a, 0xa1, 0x92, 0xbf, 0xa0, 0xe1, 0xb8,
	0xc8, 0x5d, 0xe9, 0xe0, 0x8c, 0x32, 0xe4, 0x51, 0xb9, 0x3c, 0xa8, 0x7d, 0xab, 0x00, 0x5c, 0x0b,
	0xe4, 0x86, 0x6f, 0x8e, 0x38, 0x92, 0x7f, 0x60, 0xdb, 0x8c, 0xcb, 0x9f, 0xe6, 0xae, 0xbf, 0x80,
	0x13, 0x0d, 0x7e, 0x93, 0xc8, 0x46, 0xdc, 0x3a, 0x9d, 0x24, 0xcf, 0xa9, 0xd1, 0x1c, 0x46, 0x8e,
	0x60, 0x33, 0x8c, 0xcf, 0x7c, 0x8f, 0xa1, 0x23, 0xdc, 0x66, 0x45, 0xdd, 0x62, 0x0e, 0x25, 0x2f,
	0x00, 0x4c, 0x8f, 0x33, 0xe4, 0x57, 0xce, 0x84, 0x37, 0x57, 0xdb, 0xe5, 0x4e, 0xbd, 0xd7, 0xd2,
	0x43, 0xcb, 0xf5, 0xd8, 0x72, 0xfd, 0x2a, 0xb6, 0x9c, 0x66, 0xd8, 0xc4, 0x88, 0x6a, 0x5c, 0x22,
	0xf3, 0x50, 0xe5, 0xaf, 0x2d, 0xcd, 0x9f, 0xcb, 0x20, 0xe7, 0xb0, 0xed, 0xbb, 0x73, 0x2a, 0xeb,
	0x4b, 0x55, 0x16, 0x72, 0xc8, 0xef, 0x50, 0x63, 0x3e, 0x8a, 0xcb, 0x00, 0x6d, 0x56, 0xdb, 0xe5,
	0xce, 0x06, 0x4d, 0x81, 0xa0, 0xf3, 0xe2, 0xce, 0xe5, 0x5e, 0x73, 0x43, 0x59, 0x15, 0x06, 0xda,
	0xf7, 0x15, 0xa8, 0x27, 0x3d, 0xbf, 0xee, 0x05, 0xbe, 0xaa, 0x83, 0xbc, 0xff, 0x39, 0x8c, 0x74,
	0x60, 0x4b, 0x4c, 0xb9, 0xc7, 0x50, 0x24, 0xb4, 0xd0, 0xfe, 0x79, 0x38, 0x18, 0x03, 0x8f, 0xdf,
	0x31, 0xcf, 0x8a, 0x79, 0x15, 0xc5, 0xcb, 0x83, 0x84, 0xc0, 0xaa, 0xcb, 0x22, 0xe7, 0x6b, 0x54,
	0xfd, 0x26, 0x3a, 0x10, 0x14, 0xc8, 0xc6, 0x1f, 0xb8, 0x63, 0xdf, 0x20, 0xb7, 0x82, 0x31, 0x91,
	0xca, 0xdb, 0x1a, 0x2d, 0x38, 0x09, 0xf8, 0x92, 0x8f, 0x87, 0xea, 0xa9, 0xe1, 0x38, 0x5d, 0x58,
	0xf7, 0xca, 0xc5, 0x55, 0x5a, 0x70, 0x42, 0xfe, 0x85, 0x9d, 0x18, 0x75, 0x5c, 0xfb, 0x4a, 0x8c,
	0xb8, 0x2b, 0x95, 0x67, 0x35, 0xba, 0x78, 0xa0, 0xbd, 0x81, 0xad, 0xdc, 0x62, 0x5c, 0xf7, 0x82,
	0xd5, 0x48, 0x86, 0xb2, 0x70, 0x35, 0x32, 0xae, 0xd2, 0x0c, 0x55, 0x33, 0x60, 0x33, 0x9d, 0x79,
	0xb5, 0x65, 0xff, 0x41, 0x75, 0xa0, 0xa2, 0x58, 0x67, 0x3f, 0xab, 0x93, 0x92, 0x69, 0x4c, 0xd3,
	0x7e, 0x54, 0x60, 0x8b, 0x72, 0x66, 0x45, 0xb7, 0x3c, 0x63, 0xc8, 0xc8, 0x19, 0x80, 0xcd, 0xd1,
	0x48, 0x84, 0x82, 0xf9, 0xd1, 0xb2, 0x42, 0x99, 0x84, 0x54, 0x53, 0xf6, 0x4b, 0x34, 0x93, 0x47,
	0x3e, 0xc1, 0x4e, 0x1a, 0x19, 0xb3, 0x80, 0xe8, 0xa9, 0xee, 0xd6, 0x7b, 0x7f, 0x2f, 0x17, 0x8b,
	0x12, 0xfa, 0x25, 0xba, 0xa8, 0x42, 0x38, 0xec, 0x65, 0xc1, 0xc4, 0x1f, 0x35, 0x14, 0xf5, 0xde,
	0xf1, 0xaf, 0xc8, 0x27, 0x49, 0xfd, 0x12, 0x2d, 0x56, 0x23, 0x17, 0xd0, 0xb0, 0x79, 0xfa, 0x1d,
	0x93, 0xd1, 0x42, 0x1f, 0x3e, 0x20, 0x9f, 0x12, 0xfb, 0x25, 0x9a, 0xcf, 0x24, 0x1f, 0x81, 0x64,
	0x01, 0x63, 0xf6, 0x9e, 0x25, 0x0b, 0x7e, 0xb4, 0x4c, 0x2f, 0x64, 0xf7, 0x4b, 0xb4, 0x40, 0xc3,
	0xa8, 0x41, 0xd5, 0xe3, 0xb7, 0x3e, 0x97, 0xa8, 0x9d, 0xc3, 0x7e, 0xf1, 0x6b, 0xc9, 0x3e, 0xac,
	0x8b, 0xe1, 0x50, 0x72, 0x54, 0xdd, 0x6c, 0xd0, 0x28, 0x0a, 0x36, 0x79, 0xec, 0x4c, 0x9c, 0xe4,
	0x1b, 0xae, 0x02, 0xed, 0x16, 0xfe, 0x78, 0xb4, 0x29, 0xc1, 0x6a, 0x7f, 0x0d, 0x7e, 0xcc, 0xad,
	0x76, 0x16, 0xcb, 0x94, 0x5c, 0x29, 0x2e, 0x59, 0xc9, 0x96, 0x9c, 0xc0, 0xe1, 0xd2, 0x46, 0x91,
	0x16, 0x6c, 0x04, 0xd3, 0xaf, 0xac, 0x0b, 0x4b, 0x26, 0xf1, 0x13, 0xcb, 0xbd, 0x86, 0xbd, 0xc2,
	0xc6, 0x3d, 0xd1, 0xa8, 0xe7, 0xd0, 0x7a, 0xb8, 0x5f, 0x8f, 0x5d, 0xd7, 0x38, 0xf9, 0xfc, 0xcc,
	0x76, 0xf0, 0xc6, 0x1f, 0xe8, 0xa6, 0x98, 0x74, 0x55, 0xff, 0xa7, 0x9e, 0xf8, 0xc2, 0x4d, 0x0c,
	0x83, 0xe3, 0xf0, 0xbf, 0xdc, 0x16, 0x63, 0xe6, 0xda, 0xdd, 0x74, 0x3e, 0x06, 0xeb, 0xea, 0xe0,
	0xff, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x8f, 0x37, 0xba, 0xed, 0x07, 0x00, 0x00,
}
