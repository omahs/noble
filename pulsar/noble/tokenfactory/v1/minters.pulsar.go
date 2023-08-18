// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package tokenfactoryv1

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Minters           protoreflect.MessageDescriptor
	fd_Minters_address   protoreflect.FieldDescriptor
	fd_Minters_allowance protoreflect.FieldDescriptor
)

func init() {
	file_noble_tokenfactory_v1_minters_proto_init()
	md_Minters = File_noble_tokenfactory_v1_minters_proto.Messages().ByName("Minters")
	fd_Minters_address = md_Minters.Fields().ByName("address")
	fd_Minters_allowance = md_Minters.Fields().ByName("allowance")
}

var _ protoreflect.Message = (*fastReflection_Minters)(nil)

type fastReflection_Minters Minters

func (x *Minters) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Minters)(x)
}

func (x *Minters) slowProtoReflect() protoreflect.Message {
	mi := &file_noble_tokenfactory_v1_minters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Minters_messageType fastReflection_Minters_messageType
var _ protoreflect.MessageType = fastReflection_Minters_messageType{}

type fastReflection_Minters_messageType struct{}

func (x fastReflection_Minters_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Minters)(nil)
}
func (x fastReflection_Minters_messageType) New() protoreflect.Message {
	return new(fastReflection_Minters)
}
func (x fastReflection_Minters_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Minters
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Minters) Descriptor() protoreflect.MessageDescriptor {
	return md_Minters
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Minters) Type() protoreflect.MessageType {
	return _fastReflection_Minters_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Minters) New() protoreflect.Message {
	return new(fastReflection_Minters)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Minters) Interface() protoreflect.ProtoMessage {
	return (*Minters)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Minters) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_Minters_address, value) {
			return
		}
	}
	if x.Allowance != nil {
		value := protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
		if !f(fd_Minters_allowance, value) {
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
func (x *fastReflection_Minters) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "noble.tokenfactory.v1.Minters.address":
		return x.Address != ""
	case "noble.tokenfactory.v1.Minters.allowance":
		return x.Allowance != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.tokenfactory.v1.Minters"))
		}
		panic(fmt.Errorf("message noble.tokenfactory.v1.Minters does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Minters) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "noble.tokenfactory.v1.Minters.address":
		x.Address = ""
	case "noble.tokenfactory.v1.Minters.allowance":
		x.Allowance = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.tokenfactory.v1.Minters"))
		}
		panic(fmt.Errorf("message noble.tokenfactory.v1.Minters does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Minters) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "noble.tokenfactory.v1.Minters.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "noble.tokenfactory.v1.Minters.allowance":
		value := x.Allowance
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.tokenfactory.v1.Minters"))
		}
		panic(fmt.Errorf("message noble.tokenfactory.v1.Minters does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Minters) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "noble.tokenfactory.v1.Minters.address":
		x.Address = value.Interface().(string)
	case "noble.tokenfactory.v1.Minters.allowance":
		x.Allowance = value.Message().Interface().(*v1beta1.Coin)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.tokenfactory.v1.Minters"))
		}
		panic(fmt.Errorf("message noble.tokenfactory.v1.Minters does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Minters) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "noble.tokenfactory.v1.Minters.allowance":
		if x.Allowance == nil {
			x.Allowance = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
	case "noble.tokenfactory.v1.Minters.address":
		panic(fmt.Errorf("field address of message noble.tokenfactory.v1.Minters is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.tokenfactory.v1.Minters"))
		}
		panic(fmt.Errorf("message noble.tokenfactory.v1.Minters does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Minters) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "noble.tokenfactory.v1.Minters.address":
		return protoreflect.ValueOfString("")
	case "noble.tokenfactory.v1.Minters.allowance":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: noble.tokenfactory.v1.Minters"))
		}
		panic(fmt.Errorf("message noble.tokenfactory.v1.Minters does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Minters) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in noble.tokenfactory.v1.Minters", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Minters) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Minters) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Minters) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Minters) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Minters)
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
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Allowance != nil {
			l = options.Size(x.Allowance)
			n += 1 + l + runtime.Sov(uint64(l))
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
		x := input.Message.Interface().(*Minters)
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
		if x.Allowance != nil {
			encoded, err := options.Marshal(x.Allowance)
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
			dAtA[i] = 0x12
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
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
		x := input.Message.Interface().(*Minters)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Minters: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Minters: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
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
				if x.Allowance == nil {
					x.Allowance = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Allowance); err != nil {
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
// source: noble/tokenfactory/v1/minters.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Minters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Allowance *v1beta1.Coin `protobuf:"bytes,2,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (x *Minters) Reset() {
	*x = Minters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noble_tokenfactory_v1_minters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Minters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Minters) ProtoMessage() {}

// Deprecated: Use Minters.ProtoReflect.Descriptor instead.
func (*Minters) Descriptor() ([]byte, []int) {
	return file_noble_tokenfactory_v1_minters_proto_rawDescGZIP(), []int{0}
}

func (x *Minters) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Minters) GetAllowance() *v1beta1.Coin {
	if x != nil {
		return x.Allowance
	}
	return nil
}

var File_noble_tokenfactory_v1_minters_proto protoreflect.FileDescriptor

var file_noble_tokenfactory_v1_minters_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x62, 0x0a, 0x07, 0x4d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x42, 0xf2, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x6e,
	0x6f, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x76, 0x65, 0x2d, 0x76, 0x65, 0x6e,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x2f, 0x70, 0x75, 0x6c, 0x73,
	0x61, 0x72, 0x2f, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x54, 0x58, 0xaa, 0x02, 0x15,
	0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x5c, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21,
	0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x5c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x17, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x3a, 0x3a, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_noble_tokenfactory_v1_minters_proto_rawDescOnce sync.Once
	file_noble_tokenfactory_v1_minters_proto_rawDescData = file_noble_tokenfactory_v1_minters_proto_rawDesc
)

func file_noble_tokenfactory_v1_minters_proto_rawDescGZIP() []byte {
	file_noble_tokenfactory_v1_minters_proto_rawDescOnce.Do(func() {
		file_noble_tokenfactory_v1_minters_proto_rawDescData = protoimpl.X.CompressGZIP(file_noble_tokenfactory_v1_minters_proto_rawDescData)
	})
	return file_noble_tokenfactory_v1_minters_proto_rawDescData
}

var file_noble_tokenfactory_v1_minters_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_noble_tokenfactory_v1_minters_proto_goTypes = []interface{}{
	(*Minters)(nil),      // 0: noble.tokenfactory.v1.Minters
	(*v1beta1.Coin)(nil), // 1: cosmos.base.v1beta1.Coin
}
var file_noble_tokenfactory_v1_minters_proto_depIdxs = []int32{
	1, // 0: noble.tokenfactory.v1.Minters.allowance:type_name -> cosmos.base.v1beta1.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_noble_tokenfactory_v1_minters_proto_init() }
func file_noble_tokenfactory_v1_minters_proto_init() {
	if File_noble_tokenfactory_v1_minters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_noble_tokenfactory_v1_minters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Minters); i {
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
			RawDescriptor: file_noble_tokenfactory_v1_minters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_noble_tokenfactory_v1_minters_proto_goTypes,
		DependencyIndexes: file_noble_tokenfactory_v1_minters_proto_depIdxs,
		MessageInfos:      file_noble_tokenfactory_v1_minters_proto_msgTypes,
	}.Build()
	File_noble_tokenfactory_v1_minters_proto = out.File
	file_noble_tokenfactory_v1_minters_proto_rawDesc = nil
	file_noble_tokenfactory_v1_minters_proto_goTypes = nil
	file_noble_tokenfactory_v1_minters_proto_depIdxs = nil
}
