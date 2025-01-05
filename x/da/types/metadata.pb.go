// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/da/v1/metadata.proto

package types

import (
	fmt "fmt"
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

// Metadata
type Metadata struct {
	RecoveredDataHash []byte   `protobuf:"bytes,1,opt,name=recovered_data_hash,json=recoveredDataHash,proto3" json:"recovered_data_hash,omitempty"`
	RecoveredDataSize uint64   `protobuf:"varint,2,opt,name=recovered_data_size,json=recoveredDataSize,proto3" json:"recovered_data_size,omitempty"`
	ShardSize         uint64   `protobuf:"varint,3,opt,name=shard_size,json=shardSize,proto3" json:"shard_size,omitempty"`
	ParityShardCount  uint64   `protobuf:"varint,4,opt,name=parity_shard_count,json=parityShardCount,proto3" json:"parity_shard_count,omitempty"`
	ShardUris         []string `protobuf:"bytes,5,rep,name=shard_uris,json=shardUris,proto3" json:"shard_uris,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e04403f9958bcbf, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetRecoveredDataHash() []byte {
	if m != nil {
		return m.RecoveredDataHash
	}
	return nil
}

func (m *Metadata) GetRecoveredDataSize() uint64 {
	if m != nil {
		return m.RecoveredDataSize
	}
	return 0
}

func (m *Metadata) GetShardSize() uint64 {
	if m != nil {
		return m.ShardSize
	}
	return 0
}

func (m *Metadata) GetParityShardCount() uint64 {
	if m != nil {
		return m.ParityShardCount
	}
	return 0
}

func (m *Metadata) GetShardUris() []string {
	if m != nil {
		return m.ShardUris
	}
	return nil
}

// MetadataUriWrapper
type MetadataUriWrapper struct {
	MetadataUri string `protobuf:"bytes,1,opt,name=metadata_uri,json=metadataUri,proto3" json:"metadata_uri,omitempty"`
}

func (m *MetadataUriWrapper) Reset()         { *m = MetadataUriWrapper{} }
func (m *MetadataUriWrapper) String() string { return proto.CompactTextString(m) }
func (*MetadataUriWrapper) ProtoMessage()    {}
func (*MetadataUriWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e04403f9958bcbf, []int{1}
}
func (m *MetadataUriWrapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataUriWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataUriWrapper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataUriWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataUriWrapper.Merge(m, src)
}
func (m *MetadataUriWrapper) XXX_Size() int {
	return m.Size()
}
func (m *MetadataUriWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataUriWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataUriWrapper proto.InternalMessageInfo

func (m *MetadataUriWrapper) GetMetadataUri() string {
	if m != nil {
		return m.MetadataUri
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "sunrise.da.v1.Metadata")
	proto.RegisterType((*MetadataUriWrapper)(nil), "sunrise.da.v1.MetadataUriWrapper")
}

func init() { proto.RegisterFile("sunrise/da/v1/metadata.proto", fileDescriptor_3e04403f9958bcbf) }

var fileDescriptor_3e04403f9958bcbf = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x7b, 0xb6, 0x8a, 0x39, 0x2b, 0xe8, 0xb9, 0x64, 0xd0, 0x23, 0x76, 0x0a, 0x22, 0x09,
	0xc5, 0xc1, 0x5d, 0x3b, 0xb8, 0xb8, 0xb4, 0x14, 0xc1, 0x25, 0x7c, 0x4d, 0x0e, 0x73, 0x60, 0x9a,
	0xf0, 0xdd, 0x25, 0x98, 0xfe, 0x0a, 0x7f, 0x96, 0x63, 0x27, 0x71, 0x94, 0xe4, 0x8f, 0x48, 0x2e,
	0x89, 0x05, 0x71, 0xbc, 0xf7, 0x79, 0xee, 0x83, 0xf7, 0xa5, 0xe7, 0x2a, 0x5f, 0xa3, 0x54, 0xc2,
	0x8f, 0xc0, 0x2f, 0xa6, 0x7e, 0x22, 0x34, 0x44, 0xa0, 0xc1, 0xcb, 0x30, 0xd5, 0x29, 0x3b, 0xee,
	0xa8, 0x17, 0x81, 0x57, 0x4c, 0x27, 0x9f, 0x84, 0x1e, 0x3e, 0x76, 0x06, 0xf3, 0xe8, 0x19, 0x8a,
	0x30, 0x2d, 0x04, 0x8a, 0x28, 0x68, 0x92, 0x20, 0x06, 0x15, 0xdb, 0xc4, 0x21, 0xee, 0x78, 0x7e,
	0xfa, 0x8b, 0x66, 0xa0, 0xe1, 0x01, 0x54, 0xfc, 0x8f, 0xaf, 0xe4, 0x46, 0xd8, 0x7b, 0x0e, 0x71,
	0x47, 0x7f, 0xfc, 0x85, 0xdc, 0x08, 0x76, 0x41, 0xa9, 0x8a, 0x01, 0xa3, 0x56, 0x1b, 0x1a, 0xcd,
	0x32, 0x89, 0xc1, 0xd7, 0x94, 0x65, 0x80, 0x52, 0x97, 0x41, 0x6b, 0x85, 0x69, 0xbe, 0xd6, 0xf6,
	0xc8, 0x68, 0x27, 0x2d, 0x59, 0x34, 0xe0, 0xbe, 0xc9, 0x77, 0xc7, 0x72, 0x94, 0xca, 0xde, 0x77,
	0x86, 0xae, 0xd5, 0x1d, 0x5b, 0xa2, 0x54, 0x93, 0x5b, 0xca, 0xfa, 0x5e, 0x4b, 0x94, 0x4f, 0x08,
	0x59, 0x26, 0x90, 0x5d, 0xd2, 0x71, 0xbf, 0x47, 0xf3, 0xcf, 0x54, 0xb3, 0xe6, 0x47, 0xc9, 0xce,
	0xbc, 0x9b, 0x7d, 0x54, 0x9c, 0x6c, 0x2b, 0x4e, 0xbe, 0x2b, 0x4e, 0xde, 0x6b, 0x3e, 0xd8, 0xd6,
	0x7c, 0xf0, 0x55, 0xf3, 0xc1, 0xf3, 0xd5, 0x8b, 0xd4, 0x71, 0xbe, 0xf2, 0xc2, 0x34, 0xf1, 0xbb,
	0x15, 0x5f, 0xa1, 0x14, 0xd8, 0x3f, 0xfc, 0xb7, 0x66, 0x72, 0x5d, 0x66, 0x42, 0xad, 0x0e, 0xcc,
	0xda, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0xae, 0x27, 0x3c, 0x8d, 0x01, 0x00, 0x00,
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShardUris) > 0 {
		for iNdEx := len(m.ShardUris) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ShardUris[iNdEx])
			copy(dAtA[i:], m.ShardUris[iNdEx])
			i = encodeVarintMetadata(dAtA, i, uint64(len(m.ShardUris[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.ParityShardCount != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.ParityShardCount))
		i--
		dAtA[i] = 0x20
	}
	if m.ShardSize != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.ShardSize))
		i--
		dAtA[i] = 0x18
	}
	if m.RecoveredDataSize != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.RecoveredDataSize))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RecoveredDataHash) > 0 {
		i -= len(m.RecoveredDataHash)
		copy(dAtA[i:], m.RecoveredDataHash)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.RecoveredDataHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MetadataUriWrapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataUriWrapper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataUriWrapper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MetadataUri) > 0 {
		i -= len(m.MetadataUri)
		copy(dAtA[i:], m.MetadataUri)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.MetadataUri)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RecoveredDataHash)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.RecoveredDataSize != 0 {
		n += 1 + sovMetadata(uint64(m.RecoveredDataSize))
	}
	if m.ShardSize != 0 {
		n += 1 + sovMetadata(uint64(m.ShardSize))
	}
	if m.ParityShardCount != 0 {
		n += 1 + sovMetadata(uint64(m.ParityShardCount))
	}
	if len(m.ShardUris) > 0 {
		for _, s := range m.ShardUris {
			l = len(s)
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	return n
}

func (m *MetadataUriWrapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MetadataUri)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecoveredDataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecoveredDataHash = append(m.RecoveredDataHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RecoveredDataHash == nil {
				m.RecoveredDataHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecoveredDataSize", wireType)
			}
			m.RecoveredDataSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecoveredDataSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardSize", wireType)
			}
			m.ShardSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShardSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParityShardCount", wireType)
			}
			m.ParityShardCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParityShardCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardUris", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShardUris = append(m.ShardUris, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func (m *MetadataUriWrapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: MetadataUriWrapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataUriWrapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetadataUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadata = fmt.Errorf("proto: unexpected end of group")
)