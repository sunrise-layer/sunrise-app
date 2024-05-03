// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/liquiditypool/v1/pool.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Pool struct {
	Id         uint64                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Admin      string                      `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
	BaseDenom  string                      `protobuf:"bytes,3,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	QuoteDenom string                      `protobuf:"bytes,4,opt,name=quote_denom,json=quoteDenom,proto3" json:"quote_denom,omitempty"`
	FeeRate    cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=fee_rate,json=feeRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"fee_rate"`
	FX         string                      `protobuf:"bytes,6,opt,name=f_x,json=fX,proto3" json:"f_x,omitempty"`
	FY         string                      `protobuf:"bytes,7,opt,name=f_y,json=fY,proto3" json:"f_y,omitempty"`
	FK         string                      `protobuf:"bytes,8,opt,name=f_k,json=fK,proto3" json:"f_k,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_69f2038cccc8b5a8, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pool) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *Pool) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func (m *Pool) GetQuoteDenom() string {
	if m != nil {
		return m.QuoteDenom
	}
	return ""
}

func (m *Pool) GetFX() string {
	if m != nil {
		return m.FX
	}
	return ""
}

func (m *Pool) GetFY() string {
	if m != nil {
		return m.FY
	}
	return ""
}

func (m *Pool) GetFK() string {
	if m != nil {
		return m.FK
	}
	return ""
}

func init() {
	proto.RegisterType((*Pool)(nil), "sunrise.liquiditypool.v1.Pool")
}

func init() {
	proto.RegisterFile("sunrise/liquiditypool/v1/pool.proto", fileDescriptor_69f2038cccc8b5a8)
}

var fileDescriptor_69f2038cccc8b5a8 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x3d, 0x6e, 0xdb, 0x30,
	0x18, 0x86, 0x45, 0xf9, 0x9f, 0x05, 0x5a, 0x54, 0xf0, 0xc0, 0xba, 0xa8, 0x6c, 0xb4, 0x8b, 0x51,
	0xa0, 0x22, 0x8c, 0xa2, 0x3d, 0x80, 0xe1, 0x2d, 0x01, 0x92, 0x68, 0x4a, 0xb2, 0x08, 0xb4, 0x44,
	0xc9, 0x84, 0x25, 0x7d, 0xb6, 0x44, 0x19, 0xd6, 0x2d, 0x72, 0x8c, 0x8c, 0x19, 0x72, 0x08, 0x8f,
	0x46, 0xa6, 0x20, 0x83, 0x11, 0xd8, 0x43, 0xf6, 0x9c, 0x20, 0x90, 0x28, 0x0f, 0xc9, 0x22, 0xf1,
	0x7d, 0x9e, 0x17, 0x20, 0xf9, 0x11, 0xff, 0x4a, 0xb3, 0x38, 0x11, 0x29, 0xa7, 0xa1, 0x58, 0x66,
	0xc2, 0x13, 0x32, 0x5f, 0x00, 0x84, 0x74, 0x35, 0xa2, 0xc5, 0xdf, 0x5a, 0x24, 0x20, 0xc1, 0x20,
	0x55, 0xc9, 0x7a, 0x57, 0xb2, 0x56, 0xa3, 0xde, 0x57, 0x16, 0x89, 0x18, 0x68, 0xf9, 0x55, 0xe5,
	0xde, 0x37, 0x17, 0xd2, 0x08, 0x52, 0xa7, 0x4c, 0x54, 0x85, 0x4a, 0x75, 0x03, 0x08, 0x40, 0xf1,
	0x62, 0xa5, 0xe8, 0xcf, 0x57, 0x84, 0xeb, 0xe7, 0x00, 0xa1, 0xf1, 0x19, 0xeb, 0xc2, 0x23, 0x68,
	0x80, 0x86, 0x75, 0x5b, 0x17, 0x9e, 0xd1, 0xc5, 0x0d, 0xe6, 0x45, 0x22, 0x26, 0xfa, 0x00, 0x0d,
	0x3b, 0xb6, 0x0a, 0xc6, 0x0f, 0x8c, 0xa7, 0x2c, 0xe5, 0x8e, 0xc7, 0x63, 0x88, 0x48, 0xad, 0x54,
	0x9d, 0x82, 0x4c, 0x0a, 0x60, 0xf4, 0xf1, 0xa7, 0x65, 0x06, 0xf2, 0xe8, 0xeb, 0xa5, 0xc7, 0x25,
	0x52, 0x85, 0x0b, 0xdc, 0xf6, 0x39, 0x77, 0x12, 0x26, 0x39, 0x69, 0x14, 0x76, 0xfc, 0x7f, 0xb3,
	0xeb, 0x6b, 0x4f, 0xbb, 0xfe, 0x77, 0x75, 0xd8, 0xd4, 0x9b, 0x5b, 0x02, 0x68, 0xc4, 0xe4, 0xcc,
	0x3a, 0xe5, 0x01, 0x73, 0xf3, 0x09, 0x77, 0x1f, 0xee, 0xff, 0xe0, 0xea, 0x2e, 0x13, 0xee, 0xde,
	0xbe, 0xdc, 0xfd, 0x46, 0x76, 0xcb, 0xe7, 0xdc, 0x66, 0x92, 0x1b, 0x5f, 0x70, 0xcd, 0x77, 0xd6,
	0xa4, 0x59, 0xee, 0xa5, 0xfb, 0x97, 0x0a, 0xe4, 0xa4, 0x55, 0x81, 0x2b, 0x05, 0xe6, 0xa4, 0x5d,
	0x81, 0x93, 0xf1, 0xd9, 0x66, 0x6f, 0xa2, 0xed, 0xde, 0x44, 0xcf, 0x7b, 0x13, 0xdd, 0x1c, 0x4c,
	0x6d, 0x7b, 0x30, 0xb5, 0xc7, 0x83, 0xa9, 0x5d, 0xff, 0x0b, 0x84, 0x9c, 0x65, 0x53, 0xcb, 0x85,
	0x88, 0x56, 0x73, 0x0f, 0x59, 0xce, 0x93, 0x63, 0xa0, 0xeb, 0x0f, 0x6f, 0x25, 0xf3, 0x05, 0x4f,
	0xa7, 0xcd, 0x72, 0x98, 0x7f, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0x0f, 0x3d, 0x46, 0xd1,
	0x01, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FK) > 0 {
		i -= len(m.FK)
		copy(dAtA[i:], m.FK)
		i = encodeVarintPool(dAtA, i, uint64(len(m.FK)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.FY) > 0 {
		i -= len(m.FY)
		copy(dAtA[i:], m.FY)
		i = encodeVarintPool(dAtA, i, uint64(len(m.FY)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FX) > 0 {
		i -= len(m.FX)
		copy(dAtA[i:], m.FX)
		i = encodeVarintPool(dAtA, i, uint64(len(m.FX)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.FeeRate.Size()
		i -= size
		if _, err := m.FeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.QuoteDenom) > 0 {
		i -= len(m.QuoteDenom)
		copy(dAtA[i:], m.QuoteDenom)
		i = encodeVarintPool(dAtA, i, uint64(len(m.QuoteDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintPool(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.QuoteDenom)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.FeeRate.Size()
	n += 1 + l + sovPool(uint64(l))
	l = len(m.FX)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.FY)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.FK)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FX", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FX = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FY", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FY = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FK", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FK = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)