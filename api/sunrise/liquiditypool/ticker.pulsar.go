// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package liquiditypool

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var _ protoreflect.List = (*_TickInfo_5_list)(nil)

type _TickInfo_5_list struct {
	list *[]*v1beta1.DecCoin
}

func (x *_TickInfo_5_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TickInfo_5_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TickInfo_5_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_TickInfo_5_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TickInfo_5_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TickInfo_5_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TickInfo_5_list) NewElement() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TickInfo_5_list) IsValid() bool {
	return x.list != nil
}

var (
	md_TickInfo                 protoreflect.MessageDescriptor
	fd_TickInfo_pool_id         protoreflect.FieldDescriptor
	fd_TickInfo_tick_index      protoreflect.FieldDescriptor
	fd_TickInfo_liquidity_gross protoreflect.FieldDescriptor
	fd_TickInfo_liquidity_net   protoreflect.FieldDescriptor
	fd_TickInfo_fee_growth      protoreflect.FieldDescriptor
)

func init() {
	file_sunrise_liquiditypool_ticker_proto_init()
	md_TickInfo = File_sunrise_liquiditypool_ticker_proto.Messages().ByName("TickInfo")
	fd_TickInfo_pool_id = md_TickInfo.Fields().ByName("pool_id")
	fd_TickInfo_tick_index = md_TickInfo.Fields().ByName("tick_index")
	fd_TickInfo_liquidity_gross = md_TickInfo.Fields().ByName("liquidity_gross")
	fd_TickInfo_liquidity_net = md_TickInfo.Fields().ByName("liquidity_net")
	fd_TickInfo_fee_growth = md_TickInfo.Fields().ByName("fee_growth")
}

var _ protoreflect.Message = (*fastReflection_TickInfo)(nil)

type fastReflection_TickInfo TickInfo

func (x *TickInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TickInfo)(x)
}

func (x *TickInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_sunrise_liquiditypool_ticker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TickInfo_messageType fastReflection_TickInfo_messageType
var _ protoreflect.MessageType = fastReflection_TickInfo_messageType{}

type fastReflection_TickInfo_messageType struct{}

func (x fastReflection_TickInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TickInfo)(nil)
}
func (x fastReflection_TickInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_TickInfo)
}
func (x fastReflection_TickInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TickInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TickInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_TickInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TickInfo) Type() protoreflect.MessageType {
	return _fastReflection_TickInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TickInfo) New() protoreflect.Message {
	return new(fastReflection_TickInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TickInfo) Interface() protoreflect.ProtoMessage {
	return (*TickInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TickInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.PoolId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.PoolId)
		if !f(fd_TickInfo_pool_id, value) {
			return
		}
	}
	if x.TickIndex != int64(0) {
		value := protoreflect.ValueOfInt64(x.TickIndex)
		if !f(fd_TickInfo_tick_index, value) {
			return
		}
	}
	if x.LiquidityGross != "" {
		value := protoreflect.ValueOfString(x.LiquidityGross)
		if !f(fd_TickInfo_liquidity_gross, value) {
			return
		}
	}
	if x.LiquidityNet != "" {
		value := protoreflect.ValueOfString(x.LiquidityNet)
		if !f(fd_TickInfo_liquidity_net, value) {
			return
		}
	}
	if len(x.FeeGrowth) != 0 {
		value := protoreflect.ValueOfList(&_TickInfo_5_list{list: &x.FeeGrowth})
		if !f(fd_TickInfo_fee_growth, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_TickInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "sunrise.liquiditypool.TickInfo.pool_id":
		return x.PoolId != uint64(0)
	case "sunrise.liquiditypool.TickInfo.tick_index":
		return x.TickIndex != int64(0)
	case "sunrise.liquiditypool.TickInfo.liquidity_gross":
		return x.LiquidityGross != ""
	case "sunrise.liquiditypool.TickInfo.liquidity_net":
		return x.LiquidityNet != ""
	case "sunrise.liquiditypool.TickInfo.fee_growth":
		return len(x.FeeGrowth) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.TickInfo"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.TickInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TickInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "sunrise.liquiditypool.TickInfo.pool_id":
		x.PoolId = uint64(0)
	case "sunrise.liquiditypool.TickInfo.tick_index":
		x.TickIndex = int64(0)
	case "sunrise.liquiditypool.TickInfo.liquidity_gross":
		x.LiquidityGross = ""
	case "sunrise.liquiditypool.TickInfo.liquidity_net":
		x.LiquidityNet = ""
	case "sunrise.liquiditypool.TickInfo.fee_growth":
		x.FeeGrowth = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.TickInfo"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.TickInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TickInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "sunrise.liquiditypool.TickInfo.pool_id":
		value := x.PoolId
		return protoreflect.ValueOfUint64(value)
	case "sunrise.liquiditypool.TickInfo.tick_index":
		value := x.TickIndex
		return protoreflect.ValueOfInt64(value)
	case "sunrise.liquiditypool.TickInfo.liquidity_gross":
		value := x.LiquidityGross
		return protoreflect.ValueOfString(value)
	case "sunrise.liquiditypool.TickInfo.liquidity_net":
		value := x.LiquidityNet
		return protoreflect.ValueOfString(value)
	case "sunrise.liquiditypool.TickInfo.fee_growth":
		if len(x.FeeGrowth) == 0 {
			return protoreflect.ValueOfList(&_TickInfo_5_list{})
		}
		listValue := &_TickInfo_5_list{list: &x.FeeGrowth}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.TickInfo"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.TickInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TickInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "sunrise.liquiditypool.TickInfo.pool_id":
		x.PoolId = value.Uint()
	case "sunrise.liquiditypool.TickInfo.tick_index":
		x.TickIndex = value.Int()
	case "sunrise.liquiditypool.TickInfo.liquidity_gross":
		x.LiquidityGross = value.Interface().(string)
	case "sunrise.liquiditypool.TickInfo.liquidity_net":
		x.LiquidityNet = value.Interface().(string)
	case "sunrise.liquiditypool.TickInfo.fee_growth":
		lv := value.List()
		clv := lv.(*_TickInfo_5_list)
		x.FeeGrowth = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.TickInfo"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.TickInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TickInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.liquiditypool.TickInfo.fee_growth":
		if x.FeeGrowth == nil {
			x.FeeGrowth = []*v1beta1.DecCoin{}
		}
		value := &_TickInfo_5_list{list: &x.FeeGrowth}
		return protoreflect.ValueOfList(value)
	case "sunrise.liquiditypool.TickInfo.pool_id":
		panic(fmt.Errorf("field pool_id of message sunrise.liquiditypool.TickInfo is not mutable"))
	case "sunrise.liquiditypool.TickInfo.tick_index":
		panic(fmt.Errorf("field tick_index of message sunrise.liquiditypool.TickInfo is not mutable"))
	case "sunrise.liquiditypool.TickInfo.liquidity_gross":
		panic(fmt.Errorf("field liquidity_gross of message sunrise.liquiditypool.TickInfo is not mutable"))
	case "sunrise.liquiditypool.TickInfo.liquidity_net":
		panic(fmt.Errorf("field liquidity_net of message sunrise.liquiditypool.TickInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.TickInfo"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.TickInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TickInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.liquiditypool.TickInfo.pool_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "sunrise.liquiditypool.TickInfo.tick_index":
		return protoreflect.ValueOfInt64(int64(0))
	case "sunrise.liquiditypool.TickInfo.liquidity_gross":
		return protoreflect.ValueOfString("")
	case "sunrise.liquiditypool.TickInfo.liquidity_net":
		return protoreflect.ValueOfString("")
	case "sunrise.liquiditypool.TickInfo.fee_growth":
		list := []*v1beta1.DecCoin{}
		return protoreflect.ValueOfList(&_TickInfo_5_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.TickInfo"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.TickInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TickInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in sunrise.liquiditypool.TickInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TickInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TickInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_TickInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TickInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TickInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.PoolId != 0 {
			n += 1 + runtime.Sov(uint64(x.PoolId))
		}
		if x.TickIndex != 0 {
			n += 1 + runtime.Sov(uint64(x.TickIndex))
		}
		l = len(x.LiquidityGross)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.LiquidityNet)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.FeeGrowth) > 0 {
			for _, e := range x.FeeGrowth {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*TickInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.FeeGrowth) > 0 {
			for iNdEx := len(x.FeeGrowth) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.FeeGrowth[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x2a
			}
		}
		if len(x.LiquidityNet) > 0 {
			i -= len(x.LiquidityNet)
			copy(dAtA[i:], x.LiquidityNet)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.LiquidityNet)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.LiquidityGross) > 0 {
			i -= len(x.LiquidityGross)
			copy(dAtA[i:], x.LiquidityGross)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.LiquidityGross)))
			i--
			dAtA[i] = 0x1a
		}
		if x.TickIndex != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TickIndex))
			i--
			dAtA[i] = 0x10
		}
		if x.PoolId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.PoolId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*TickInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TickInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TickInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
				}
				x.PoolId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.PoolId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TickIndex", wireType)
				}
				x.TickIndex = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TickIndex |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LiquidityGross", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LiquidityGross = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LiquidityNet", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LiquidityNet = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FeeGrowth", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.FeeGrowth = append(x.FeeGrowth, &v1beta1.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.FeeGrowth[len(x.FeeGrowth)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: sunrise/liquiditypool/ticker.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TickInfo
type TickInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolId         uint64             `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TickIndex      int64              `protobuf:"varint,2,opt,name=tick_index,json=tickIndex,proto3" json:"tick_index,omitempty"`
	LiquidityGross string             `protobuf:"bytes,3,opt,name=liquidity_gross,json=liquidityGross,proto3" json:"liquidity_gross,omitempty"`
	LiquidityNet   string             `protobuf:"bytes,4,opt,name=liquidity_net,json=liquidityNet,proto3" json:"liquidity_net,omitempty"`
	FeeGrowth      []*v1beta1.DecCoin `protobuf:"bytes,5,rep,name=fee_growth,json=feeGrowth,proto3" json:"fee_growth,omitempty"`
}

func (x *TickInfo) Reset() {
	*x = TickInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sunrise_liquiditypool_ticker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TickInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickInfo) ProtoMessage() {}

// Deprecated: Use TickInfo.ProtoReflect.Descriptor instead.
func (*TickInfo) Descriptor() ([]byte, []int) {
	return file_sunrise_liquiditypool_ticker_proto_rawDescGZIP(), []int{0}
}

func (x *TickInfo) GetPoolId() uint64 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *TickInfo) GetTickIndex() int64 {
	if x != nil {
		return x.TickIndex
	}
	return 0
}

func (x *TickInfo) GetLiquidityGross() string {
	if x != nil {
		return x.LiquidityGross
	}
	return ""
}

func (x *TickInfo) GetLiquidityNet() string {
	if x != nil {
		return x.LiquidityNet
	}
	return ""
}

func (x *TickInfo) GetFeeGrowth() []*v1beta1.DecCoin {
	if x != nil {
		return x.FeeGrowth
	}
	return nil
}

var File_sunrise_liquiditypool_ticker_proto protoreflect.FileDescriptor

var file_sunrise_liquiditypool_ticker_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6c, 0x69,
	0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x1a, 0x1e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x63, 0x6b, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x63,
	0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x66, 0x0a, 0x0f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x3d, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x1b, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73,
	0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x44, 0x65, 0x63, 0xf2, 0xde, 0x1f, 0x16, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x6c, 0x69,
	0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x22, 0x52, 0x0e,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x73, 0x73, 0x12, 0x60,
	0x0a, 0x0d, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3b, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x1b, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68,
	0x2e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x44, 0x65, 0x63, 0xf2, 0xde, 0x1f, 0x14, 0x79, 0x61,
	0x6d, 0x6c, 0x3a, 0x22, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x65,
	0x74, 0x22, 0x52, 0x0c, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x4e, 0x65, 0x74,
	0x12, 0x70, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f,
	0x69, 0x6e, 0x42, 0x33, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44,
	0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x09, 0x66, 0x65, 0x65, 0x47, 0x72, 0x6f, 0x77,
	0x74, 0x68, 0x42, 0xc5, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x6e, 0x72, 0x69,
	0x73, 0x65, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c,
	0x42, 0x0b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x26, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0xa2, 0x02, 0x03, 0x53, 0x4c, 0x58, 0xaa, 0x02, 0x15,
	0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x70, 0x6f, 0x6f, 0x6c, 0xca, 0x02, 0x15, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x5c,
	0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0xe2, 0x02, 0x21,
	0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x5c, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x16, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x3a, 0x3a, 0x4c, 0x69, 0x71,
	0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sunrise_liquiditypool_ticker_proto_rawDescOnce sync.Once
	file_sunrise_liquiditypool_ticker_proto_rawDescData = file_sunrise_liquiditypool_ticker_proto_rawDesc
)

func file_sunrise_liquiditypool_ticker_proto_rawDescGZIP() []byte {
	file_sunrise_liquiditypool_ticker_proto_rawDescOnce.Do(func() {
		file_sunrise_liquiditypool_ticker_proto_rawDescData = protoimpl.X.CompressGZIP(file_sunrise_liquiditypool_ticker_proto_rawDescData)
	})
	return file_sunrise_liquiditypool_ticker_proto_rawDescData
}

var file_sunrise_liquiditypool_ticker_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sunrise_liquiditypool_ticker_proto_goTypes = []interface{}{
	(*TickInfo)(nil),        // 0: sunrise.liquiditypool.TickInfo
	(*v1beta1.DecCoin)(nil), // 1: cosmos.base.v1beta1.DecCoin
}
var file_sunrise_liquiditypool_ticker_proto_depIdxs = []int32{
	1, // 0: sunrise.liquiditypool.TickInfo.fee_growth:type_name -> cosmos.base.v1beta1.DecCoin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sunrise_liquiditypool_ticker_proto_init() }
func file_sunrise_liquiditypool_ticker_proto_init() {
	if File_sunrise_liquiditypool_ticker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sunrise_liquiditypool_ticker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TickInfo); i {
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
			RawDescriptor: file_sunrise_liquiditypool_ticker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sunrise_liquiditypool_ticker_proto_goTypes,
		DependencyIndexes: file_sunrise_liquiditypool_ticker_proto_depIdxs,
		MessageInfos:      file_sunrise_liquiditypool_ticker_proto_msgTypes,
	}.Build()
	File_sunrise_liquiditypool_ticker_proto = out.File
	file_sunrise_liquiditypool_ticker_proto_rawDesc = nil
	file_sunrise_liquiditypool_ticker_proto_goTypes = nil
	file_sunrise_liquiditypool_ticker_proto_depIdxs = nil
}
