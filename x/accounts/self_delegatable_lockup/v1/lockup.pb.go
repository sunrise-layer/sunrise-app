// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/accounts/self_delegatable_lockup/v1/lockup.proto

package v1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Period defines a length of time and amount of coins that will be lock.
type Period struct {
	// Period duration
	Length time.Duration                            `protobuf:"bytes,1,opt,name=length,proto3,stdduration" json:"length"`
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a93bf358cf97793, []int{0}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetLength() time.Duration {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Period) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

type UnbondingEntries struct {
	Entries []*UnbondingEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (m *UnbondingEntries) Reset()         { *m = UnbondingEntries{} }
func (m *UnbondingEntries) String() string { return proto.CompactTextString(m) }
func (*UnbondingEntries) ProtoMessage()    {}
func (*UnbondingEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a93bf358cf97793, []int{1}
}
func (m *UnbondingEntries) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnbondingEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnbondingEntries.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnbondingEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbondingEntries.Merge(m, src)
}
func (m *UnbondingEntries) XXX_Size() int {
	return m.Size()
}
func (m *UnbondingEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbondingEntries.DiscardUnknown(m)
}

var xxx_messageInfo_UnbondingEntries proto.InternalMessageInfo

func (m *UnbondingEntries) GetEntries() []*UnbondingEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// UnbondingEntry defines an entry tracking the lockup account unbonding operation.
type UnbondingEntry struct {
	CreationHeight int64 `protobuf:"varint,1,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height,omitempty"`
	// end time of entry
	EndTime time.Time `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
	// unbond amount
	Amount types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	// validator address
	ValidatorAddress string `protobuf:"bytes,4,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *UnbondingEntry) Reset()         { *m = UnbondingEntry{} }
func (m *UnbondingEntry) String() string { return proto.CompactTextString(m) }
func (*UnbondingEntry) ProtoMessage()    {}
func (*UnbondingEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a93bf358cf97793, []int{2}
}
func (m *UnbondingEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnbondingEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnbondingEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnbondingEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbondingEntry.Merge(m, src)
}
func (m *UnbondingEntry) XXX_Size() int {
	return m.Size()
}
func (m *UnbondingEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbondingEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UnbondingEntry proto.InternalMessageInfo

func (m *UnbondingEntry) GetCreationHeight() int64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

func (m *UnbondingEntry) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *UnbondingEntry) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *UnbondingEntry) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Period)(nil), "sunrise.accounts.self_delegatable_lockup.v1.Period")
	proto.RegisterType((*UnbondingEntries)(nil), "sunrise.accounts.self_delegatable_lockup.v1.UnbondingEntries")
	proto.RegisterType((*UnbondingEntry)(nil), "sunrise.accounts.self_delegatable_lockup.v1.UnbondingEntry")
}

func init() {
	proto.RegisterFile("sunrise/accounts/self_delegatable_lockup/v1/lockup.proto", fileDescriptor_0a93bf358cf97793)
}

var fileDescriptor_0a93bf358cf97793 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xe3, 0x04, 0xa5, 0xd4, 0x40, 0x69, 0x2d, 0x86, 0x34, 0x12, 0x4e, 0xc8, 0x42, 0x54,
	0xd4, 0x3b, 0x05, 0x16, 0x24, 0x18, 0x20, 0xb4, 0x88, 0x09, 0xa1, 0x40, 0x19, 0x58, 0xac, 0xb3,
	0xfd, 0xf6, 0x72, 0xaa, 0x7d, 0x17, 0xdd, 0x9d, 0x23, 0xf2, 0x21, 0x90, 0x3a, 0x22, 0x3e, 0x01,
	0x62, 0xea, 0xc0, 0x57, 0x40, 0xea, 0x58, 0x31, 0x31, 0x51, 0x94, 0x0c, 0xfd, 0x1a, 0xc8, 0x77,
	0xe7, 0x8a, 0x14, 0x55, 0x82, 0xc5, 0x7e, 0xff, 0xdc, 0xf3, 0x9c, 0xdf, 0xdf, 0x2b, 0xfb, 0x0f,
	0x55, 0xc1, 0x25, 0x53, 0x80, 0x49, 0x92, 0x88, 0x82, 0x6b, 0x85, 0x15, 0x64, 0xfb, 0x51, 0x0a,
	0x19, 0x50, 0xa2, 0x49, 0x9c, 0x41, 0x94, 0x89, 0xe4, 0xa0, 0x98, 0xe0, 0xe9, 0x00, 0xdb, 0x08,
	0x4d, 0xa4, 0xd0, 0x22, 0xb8, 0xe7, 0x94, 0xa8, 0x52, 0xa2, 0x4b, 0x94, 0x68, 0x3a, 0x68, 0x6f,
	0x90, 0x9c, 0x71, 0x81, 0xcd, 0xd3, 0xea, 0xdb, 0x61, 0x22, 0x54, 0x2e, 0x14, 0x8e, 0x89, 0x02,
	0x3c, 0x1d, 0xc4, 0xa0, 0xc9, 0x00, 0x27, 0x82, 0x71, 0xd7, 0xdf, 0xb4, 0xfd, 0xc8, 0x64, 0xd8,
	0x26, 0xae, 0x75, 0x8b, 0x0a, 0x2a, 0x6c, 0xbd, 0x8c, 0x2a, 0x43, 0x2a, 0x04, 0xcd, 0x00, 0x9b,
	0x2c, 0x2e, 0xf6, 0x71, 0x5a, 0x48, 0xa2, 0x99, 0xa8, 0x0c, 0x3b, 0x17, 0xfb, 0x9a, 0xe5, 0xa0,
	0x34, 0xc9, 0xdd, 0x44, 0xbd, 0x6f, 0x9e, 0xdf, 0x7c, 0x05, 0x92, 0x89, 0x34, 0x78, 0xe2, 0x37,
	0x33, 0xe0, 0x54, 0x8f, 0x5b, 0x5e, 0xd7, 0xeb, 0x5f, 0xbb, 0xbf, 0x89, 0xac, 0x18, 0x55, 0x62,
	0xb4, 0xe3, 0xcc, 0x87, 0x37, 0x8e, 0x7f, 0x76, 0x6a, 0x1f, 0x4f, 0x3b, 0xde, 0xe7, 0xb3, 0xa3,
	0x2d, 0x6f, 0xe4, 0x74, 0xc1, 0xcc, 0x6f, 0x92, 0xbc, 0xe4, 0xd2, 0xaa, 0x77, 0x1b, 0xc6, 0xc1,
	0x8d, 0x50, 0xce, 0x8b, 0xdc, 0xbc, 0xe8, 0x99, 0x60, 0x7c, 0xf8, 0xbc, 0x74, 0xf8, 0x72, 0xda,
	0xe9, 0x53, 0xa6, 0xc7, 0x45, 0x8c, 0x12, 0x91, 0xbb, 0x79, 0xdd, 0x6b, 0x5b, 0xa5, 0x07, 0x58,
	0xcf, 0x26, 0xa0, 0x8c, 0x40, 0x7d, 0x3a, 0x3b, 0xda, 0xba, 0x5e, 0x62, 0x4e, 0x66, 0x51, 0x49,
	0x4c, 0xb9, 0xab, 0xed, 0x85, 0x3d, 0xe6, 0xaf, 0xef, 0xf1, 0x58, 0xf0, 0x94, 0x71, 0xba, 0xcb,
	0xb5, 0x64, 0xa0, 0x82, 0x3d, 0x7f, 0x05, 0x6c, 0xd8, 0xf2, 0xcc, 0xf7, 0x3c, 0x42, 0xff, 0xb1,
	0x3f, 0xb4, 0xe4, 0x37, 0x1b, 0x55, 0x5e, 0xbd, 0x0f, 0x75, 0x7f, 0x6d, 0xb9, 0x17, 0xdc, 0xf5,
	0x6f, 0x26, 0x12, 0x0c, 0x9b, 0x68, 0x0c, 0x8c, 0x8e, 0xb5, 0x61, 0xd8, 0x18, 0xad, 0x55, 0xe5,
	0x17, 0xa6, 0x1a, 0xec, 0xf8, 0x57, 0x81, 0xa7, 0x51, 0xb9, 0x85, 0x56, 0xdd, 0x50, 0x6e, 0xff,
	0x45, 0xf9, 0x4d, 0xb5, 0x22, 0x8b, 0xf9, 0xf0, 0x1c, 0xf3, 0x0a, 0xf0, 0xb4, 0x6c, 0x06, 0x8f,
	0xcf, 0x39, 0x37, 0xdc, 0xa6, 0x2e, 0xe5, 0xbc, 0x5a, 0x5a, 0x2c, 0xa1, 0x0a, 0x5e, 0xfa, 0x1b,
	0x53, 0x92, 0xb1, 0x94, 0x68, 0x21, 0x23, 0x92, 0xa6, 0x12, 0x94, 0x6a, 0x5d, 0xe9, 0x7a, 0xfd,
	0xd5, 0xe1, 0x9d, 0xef, 0x5f, 0xb7, 0x6f, 0x3b, 0xaf, 0xb7, 0xd5, 0x99, 0xa7, 0xf6, 0xc8, 0x6b,
	0x2d, 0x19, 0xa7, 0xa3, 0xf5, 0xe9, 0x85, 0xfa, 0x30, 0x3a, 0x9e, 0x87, 0xde, 0xc9, 0x3c, 0xf4,
	0x7e, 0xcd, 0x43, 0xef, 0x70, 0x11, 0xd6, 0x4e, 0x16, 0x61, 0xed, 0xc7, 0x22, 0xac, 0xbd, 0xdb,
	0xfd, 0x63, 0xb9, 0x8e, 0x7c, 0x46, 0x66, 0x20, 0xab, 0x04, 0xbf, 0xff, 0x97, 0x5f, 0x30, 0x6e,
	0x1a, 0x34, 0x0f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xd1, 0xd9, 0x13, 0xb8, 0x03, 0x00,
	0x00,
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLockup(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Length, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Length):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLockup(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *UnbondingEntries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnbondingEntries) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnbondingEntries) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLockup(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UnbondingEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnbondingEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnbondingEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintLockup(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLockup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.EndTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintLockup(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if m.CreationHeight != 0 {
		i = encodeVarintLockup(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLockup(dAtA []byte, offset int, v uint64) int {
	offset -= sovLockup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Length)
	n += 1 + l + sovLockup(uint64(l))
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovLockup(uint64(l))
		}
	}
	return n
}

func (m *UnbondingEntries) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovLockup(uint64(l))
		}
	}
	return n
}

func (m *UnbondingEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreationHeight != 0 {
		n += 1 + sovLockup(uint64(m.CreationHeight))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLockup(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovLockup(uint64(l))
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovLockup(uint64(l))
	}
	return n
}

func sovLockup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLockup(x uint64) (n int) {
	return sovLockup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLockup
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockup
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
				return ErrInvalidLengthLockup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Length, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockup
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
				return ErrInvalidLengthLockup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLockup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLockup
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
func (m *UnbondingEntries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLockup
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
			return fmt.Errorf("proto: UnbondingEntries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnbondingEntries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockup
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
				return ErrInvalidLengthLockup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &UnbondingEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLockup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLockup
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
func (m *UnbondingEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLockup
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
			return fmt.Errorf("proto: UnbondingEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnbondingEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockup
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
				return ErrInvalidLengthLockup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockup
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
				return ErrInvalidLengthLockup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLockup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLockup
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
				return ErrInvalidLengthLockup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLockup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLockup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLockup
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
func skipLockup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLockup
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
					return 0, ErrIntOverflowLockup
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
					return 0, ErrIntOverflowLockup
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
				return 0, ErrInvalidLengthLockup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLockup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLockup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLockup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLockup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLockup = fmt.Errorf("proto: unexpected end of group")
)
