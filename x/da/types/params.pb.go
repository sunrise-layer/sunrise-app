// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/da/v1/params.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the module.
type Params struct {
	VoteThreshold       string                                   `protobuf:"bytes,1,opt,name=vote_threshold,json=voteThreshold,proto3" json:"vote_threshold,omitempty"`
	SlashEpoch          uint64                                   `protobuf:"varint,2,opt,name=slash_epoch,json=slashEpoch,proto3" json:"slash_epoch,omitempty"`
	EpochMaxFault       uint64                                   `protobuf:"varint,3,opt,name=epoch_max_fault,json=epochMaxFault,proto3" json:"epoch_max_fault,omitempty"`
	SlashFraction       string                                   `protobuf:"bytes,4,opt,name=slash_fraction,json=slashFraction,proto3" json:"slash_fraction,omitempty"`
	ReplicationFactor   string                                   `protobuf:"bytes,5,opt,name=replication_factor,json=replicationFactor,proto3" json:"replication_factor,omitempty"`
	MinShardCount       uint64                                   `protobuf:"varint,6,opt,name=min_shard_count,json=minShardCount,proto3" json:"min_shard_count,omitempty"`
	MaxShardCount       uint64                                   `protobuf:"varint,7,opt,name=max_shard_count,json=maxShardCount,proto3" json:"max_shard_count,omitempty"`
	MaxShardSize        uint64                                   `protobuf:"varint,8,opt,name=max_shard_size,json=maxShardSize,proto3" json:"max_shard_size,omitempty"`
	ChallengePeriod     time.Duration                            `protobuf:"bytes,9,opt,name=challenge_period,json=challengePeriod,proto3,stdduration" json:"challenge_period"`
	ProofPeriod         time.Duration                            `protobuf:"bytes,10,opt,name=proof_period,json=proofPeriod,proto3,stdduration" json:"proof_period"`
	ChallengeCollateral github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,11,rep,name=challenge_collateral,json=challengeCollateral,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"challenge_collateral"`
	ZkpProvingKey       []byte                                   `protobuf:"bytes,12,opt,name=zkp_proving_key,json=zkpProvingKey,proto3" json:"zkp_proving_key,omitempty"`
	ZkpVerifyingKey     []byte                                   `protobuf:"bytes,13,opt,name=zkp_verifying_key,json=zkpVerifyingKey,proto3" json:"zkp_verifying_key,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_70863a5b1732e9c8, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetVoteThreshold() string {
	if m != nil {
		return m.VoteThreshold
	}
	return ""
}

func (m *Params) GetSlashEpoch() uint64 {
	if m != nil {
		return m.SlashEpoch
	}
	return 0
}

func (m *Params) GetEpochMaxFault() uint64 {
	if m != nil {
		return m.EpochMaxFault
	}
	return 0
}

func (m *Params) GetSlashFraction() string {
	if m != nil {
		return m.SlashFraction
	}
	return ""
}

func (m *Params) GetReplicationFactor() string {
	if m != nil {
		return m.ReplicationFactor
	}
	return ""
}

func (m *Params) GetMinShardCount() uint64 {
	if m != nil {
		return m.MinShardCount
	}
	return 0
}

func (m *Params) GetMaxShardCount() uint64 {
	if m != nil {
		return m.MaxShardCount
	}
	return 0
}

func (m *Params) GetMaxShardSize() uint64 {
	if m != nil {
		return m.MaxShardSize
	}
	return 0
}

func (m *Params) GetChallengePeriod() time.Duration {
	if m != nil {
		return m.ChallengePeriod
	}
	return 0
}

func (m *Params) GetProofPeriod() time.Duration {
	if m != nil {
		return m.ProofPeriod
	}
	return 0
}

func (m *Params) GetChallengeCollateral() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ChallengeCollateral
	}
	return nil
}

func (m *Params) GetZkpProvingKey() []byte {
	if m != nil {
		return m.ZkpProvingKey
	}
	return nil
}

func (m *Params) GetZkpVerifyingKey() []byte {
	if m != nil {
		return m.ZkpVerifyingKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "sunrise.da.v1.Params")
}

func init() { proto.RegisterFile("sunrise/da/v1/params.proto", fileDescriptor_70863a5b1732e9c8) }

var fileDescriptor_70863a5b1732e9c8 = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x63, 0xfa, 0x83, 0xf6, 0xf2, 0xa3, 0xd4, 0x74, 0x70, 0x33, 0x38, 0x11, 0x42, 0x28,
	0xaa, 0x54, 0x9b, 0x80, 0x58, 0x90, 0x58, 0x92, 0x90, 0xa5, 0x02, 0x45, 0x29, 0x62, 0x60, 0xb1,
	0x2e, 0xf6, 0xd9, 0x3e, 0xc5, 0xf6, 0x59, 0x77, 0x67, 0x2b, 0xce, 0xc0, 0xdf, 0xc0, 0xc8, 0xc8,
	0xcc, 0xcc, 0x1f, 0xd1, 0xb1, 0x62, 0x62, 0xa2, 0x28, 0x91, 0x10, 0x7f, 0x06, 0xba, 0xf3, 0x39,
	0x0d, 0x52, 0x07, 0xa6, 0xe4, 0x7d, 0xdf, 0xe7, 0x7d, 0xdf, 0xf3, 0xdd, 0x3b, 0xd0, 0x66, 0x59,
	0x42, 0x31, 0x43, 0xb6, 0x07, 0xed, 0xbc, 0x6f, 0xa7, 0x90, 0xc2, 0x98, 0x59, 0x29, 0x25, 0x9c,
	0xe8, 0x4d, 0x95, 0xb3, 0x3c, 0x68, 0xe5, 0xfd, 0xb6, 0xe9, 0x12, 0x16, 0x13, 0x66, 0xcf, 0x20,
	0x43, 0x76, 0xde, 0x9f, 0x21, 0x0e, 0xfb, 0xb6, 0x4b, 0x70, 0x52, 0xe2, 0xed, 0xd3, 0x32, 0xef,
	0xc8, 0xc8, 0x2e, 0x03, 0x95, 0x3a, 0x09, 0x48, 0x40, 0x4a, 0x5d, 0xfc, 0x53, 0xaa, 0x19, 0x10,
	0x12, 0x44, 0xc8, 0x96, 0xd1, 0x2c, 0xf3, 0x6d, 0x2f, 0xa3, 0x90, 0x63, 0xa2, 0x0c, 0x1f, 0xfd,
	0xde, 0x03, 0xfb, 0x13, 0x39, 0x90, 0xfe, 0x02, 0xb4, 0x72, 0xc2, 0x91, 0xc3, 0x43, 0x8a, 0x58,
	0x48, 0x22, 0xcf, 0xd0, 0xba, 0x5a, 0xef, 0x70, 0xd0, 0xfa, 0xfe, 0xed, 0x1c, 0xa8, 0x56, 0x23,
	0xe4, 0x4e, 0x9b, 0x82, 0x7a, 0x57, 0x41, 0x7a, 0x07, 0xd4, 0x59, 0x04, 0x59, 0xe8, 0xa0, 0x94,
	0xb8, 0xa1, 0x71, 0xaf, 0xab, 0xf5, 0x76, 0xa7, 0x40, 0x4a, 0xaf, 0x85, 0xa2, 0x3f, 0x01, 0x47,
	0x32, 0xe5, 0xc4, 0x70, 0xe1, 0xf8, 0x30, 0x8b, 0xb8, 0xb1, 0x23, 0xa1, 0xa6, 0x94, 0xdf, 0xc0,
	0xc5, 0x58, 0x88, 0xa2, 0x7f, 0x69, 0xe4, 0x53, 0xe8, 0x8a, 0x11, 0x8d, 0xdd, 0xbb, 0xfb, 0x4b,
	0x6a, 0xac, 0x20, 0xfd, 0x15, 0xd0, 0x29, 0x4a, 0x23, 0xec, 0xca, 0xcf, 0x72, 0x7c, 0xe8, 0x72,
	0x42, 0x8d, 0xbd, 0x3b, 0x4b, 0x8f, 0xb7, 0xc8, 0xb1, 0x04, 0xc5, 0x74, 0x31, 0x4e, 0x1c, 0x16,
	0x42, 0xea, 0x39, 0x2e, 0xc9, 0x12, 0x6e, 0xec, 0x97, 0xd3, 0xc5, 0x38, 0xb9, 0x14, 0xea, 0x50,
	0x88, 0x92, 0x83, 0x8b, 0x7f, 0xb8, 0xfb, 0x8a, 0x83, 0x8b, 0x2d, 0xee, 0x31, 0x68, 0xdd, 0x72,
	0x0c, 0x2f, 0x91, 0x71, 0x20, 0xb1, 0x46, 0x85, 0x5d, 0xe2, 0x25, 0xd2, 0xdf, 0x82, 0x07, 0x6e,
	0x08, 0xa3, 0x08, 0x25, 0x01, 0x72, 0x52, 0x44, 0x31, 0xf1, 0x8c, 0xc3, 0xae, 0xd6, 0xab, 0x3f,
	0x3b, 0xb5, 0xca, 0x1b, 0xb3, 0xaa, 0x1b, 0xb3, 0x46, 0xea, 0xc6, 0x06, 0x07, 0x57, 0x3f, 0x3b,
	0xb5, 0xcf, 0x37, 0x1d, 0x6d, 0x7a, 0xb4, 0x29, 0x9e, 0xc8, 0x5a, 0x7d, 0x0c, 0x1a, 0x29, 0x25,
	0xc4, 0xaf, 0xbc, 0xc0, 0xff, 0x7b, 0xd5, 0x65, 0xa1, 0xf2, 0xf9, 0x08, 0x4e, 0x6e, 0xe7, 0x72,
	0x49, 0x14, 0x41, 0x8e, 0x28, 0x8c, 0x8c, 0x7a, 0x77, 0x47, 0xfa, 0xa9, 0xb3, 0x14, 0xeb, 0x69,
	0xa9, 0xf5, 0xb4, 0x86, 0x04, 0x27, 0x83, 0xa7, 0xc2, 0xef, 0xeb, 0x4d, 0xa7, 0x17, 0x60, 0x1e,
	0x66, 0x33, 0xcb, 0x25, 0xb1, 0x5a, 0x4f, 0xf5, 0x73, 0xce, 0xbc, 0xb9, 0xcd, 0x8b, 0x14, 0x31,
	0x59, 0xc0, 0xa6, 0x0f, 0x37, 0x8d, 0x86, 0x9b, 0x3e, 0xe2, 0x94, 0x97, 0xf3, 0x54, 0xac, 0x77,
	0x8e, 0x93, 0xc0, 0x99, 0xa3, 0xc2, 0x68, 0x74, 0xb5, 0x5e, 0x63, 0xda, 0x5c, 0xce, 0xd3, 0x49,
	0xa9, 0x5e, 0xa0, 0x42, 0x3f, 0x03, 0xc7, 0x82, 0xcb, 0x11, 0xc5, 0x7e, 0x51, 0x91, 0x4d, 0x49,
	0x0a, 0x83, 0xf7, 0x95, 0x7e, 0x81, 0x8a, 0x97, 0xbb, 0x7f, 0xbe, 0x74, 0xb4, 0xc1, 0xe8, 0x6a,
	0x65, 0x6a, 0xd7, 0x2b, 0x53, 0xfb, 0xb5, 0x32, 0xb5, 0x4f, 0x6b, 0xb3, 0x76, 0xbd, 0x36, 0x6b,
	0x3f, 0xd6, 0x66, 0xed, 0xc3, 0xd9, 0xd6, 0xc8, 0xea, 0x35, 0x46, 0xb0, 0x40, 0xb4, 0x0a, 0xec,
	0x85, 0x78, 0xb8, 0x72, 0xf4, 0xd9, 0xbe, 0x3c, 0xc9, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x2d, 0x29, 0xab, 0x36, 0xd3, 0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VoteThreshold != that1.VoteThreshold {
		return false
	}
	if this.SlashEpoch != that1.SlashEpoch {
		return false
	}
	if this.EpochMaxFault != that1.EpochMaxFault {
		return false
	}
	if this.SlashFraction != that1.SlashFraction {
		return false
	}
	if this.ReplicationFactor != that1.ReplicationFactor {
		return false
	}
	if this.MinShardCount != that1.MinShardCount {
		return false
	}
	if this.MaxShardCount != that1.MaxShardCount {
		return false
	}
	if this.MaxShardSize != that1.MaxShardSize {
		return false
	}
	if this.ChallengePeriod != that1.ChallengePeriod {
		return false
	}
	if this.ProofPeriod != that1.ProofPeriod {
		return false
	}
	if len(this.ChallengeCollateral) != len(that1.ChallengeCollateral) {
		return false
	}
	for i := range this.ChallengeCollateral {
		if !this.ChallengeCollateral[i].Equal(&that1.ChallengeCollateral[i]) {
			return false
		}
	}
	if !bytes.Equal(this.ZkpProvingKey, that1.ZkpProvingKey) {
		return false
	}
	if !bytes.Equal(this.ZkpVerifyingKey, that1.ZkpVerifyingKey) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ZkpVerifyingKey) > 0 {
		i -= len(m.ZkpVerifyingKey)
		copy(dAtA[i:], m.ZkpVerifyingKey)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ZkpVerifyingKey)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.ZkpProvingKey) > 0 {
		i -= len(m.ZkpProvingKey)
		copy(dAtA[i:], m.ZkpProvingKey)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ZkpProvingKey)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.ChallengeCollateral) > 0 {
		for iNdEx := len(m.ChallengeCollateral) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChallengeCollateral[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.ProofPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ProofPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.ChallengePeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ChallengePeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x4a
	if m.MaxShardSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxShardSize))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxShardCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxShardCount))
		i--
		dAtA[i] = 0x38
	}
	if m.MinShardCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinShardCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ReplicationFactor) > 0 {
		i -= len(m.ReplicationFactor)
		copy(dAtA[i:], m.ReplicationFactor)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ReplicationFactor)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SlashFraction) > 0 {
		i -= len(m.SlashFraction)
		copy(dAtA[i:], m.SlashFraction)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SlashFraction)))
		i--
		dAtA[i] = 0x22
	}
	if m.EpochMaxFault != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EpochMaxFault))
		i--
		dAtA[i] = 0x18
	}
	if m.SlashEpoch != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SlashEpoch))
		i--
		dAtA[i] = 0x10
	}
	if len(m.VoteThreshold) > 0 {
		i -= len(m.VoteThreshold)
		copy(dAtA[i:], m.VoteThreshold)
		i = encodeVarintParams(dAtA, i, uint64(len(m.VoteThreshold)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VoteThreshold)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.SlashEpoch != 0 {
		n += 1 + sovParams(uint64(m.SlashEpoch))
	}
	if m.EpochMaxFault != 0 {
		n += 1 + sovParams(uint64(m.EpochMaxFault))
	}
	l = len(m.SlashFraction)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ReplicationFactor)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MinShardCount != 0 {
		n += 1 + sovParams(uint64(m.MinShardCount))
	}
	if m.MaxShardCount != 0 {
		n += 1 + sovParams(uint64(m.MaxShardCount))
	}
	if m.MaxShardSize != 0 {
		n += 1 + sovParams(uint64(m.MaxShardSize))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ChallengePeriod)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ProofPeriod)
	n += 1 + l + sovParams(uint64(l))
	if len(m.ChallengeCollateral) > 0 {
		for _, e := range m.ChallengeCollateral {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = len(m.ZkpProvingKey)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ZkpVerifyingKey)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoteThreshold = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashEpoch", wireType)
			}
			m.SlashEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochMaxFault", wireType)
			}
			m.EpochMaxFault = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochMaxFault |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFraction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashFraction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicationFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReplicationFactor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinShardCount", wireType)
			}
			m.MinShardCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinShardCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxShardCount", wireType)
			}
			m.MaxShardCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxShardCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxShardSize", wireType)
			}
			m.MaxShardSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxShardSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengePeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.ChallengePeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.ProofPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeCollateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengeCollateral = append(m.ChallengeCollateral, types.Coin{})
			if err := m.ChallengeCollateral[len(m.ChallengeCollateral)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZkpProvingKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZkpProvingKey = append(m.ZkpProvingKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ZkpProvingKey == nil {
				m.ZkpProvingKey = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZkpVerifyingKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZkpVerifyingKey = append(m.ZkpVerifyingKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ZkpVerifyingKey == nil {
				m.ZkpVerifyingKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
