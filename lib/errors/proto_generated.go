// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errors.proto

package errors

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ErrorDetailBool
type ErrorDetailBool struct {
	//
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	//
	Value                bool     `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorDetailBool) Reset()         { *m = ErrorDetailBool{} }
func (m *ErrorDetailBool) String() string { return proto.CompactTextString(m) }
func (*ErrorDetailBool) ProtoMessage()    {}
func (*ErrorDetailBool) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0}
}
func (m *ErrorDetailBool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorDetailBool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorDetailBool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorDetailBool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDetailBool.Merge(m, src)
}
func (m *ErrorDetailBool) XXX_Size() int {
	return m.Size()
}
func (m *ErrorDetailBool) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDetailBool.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDetailBool proto.InternalMessageInfo

func (m *ErrorDetailBool) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ErrorDetailBool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// ErrorDetailString
type ErrorDetailString struct {
	//
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	//
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorDetailString) Reset()         { *m = ErrorDetailString{} }
func (m *ErrorDetailString) String() string { return proto.CompactTextString(m) }
func (*ErrorDetailString) ProtoMessage()    {}
func (*ErrorDetailString) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{1}
}
func (m *ErrorDetailString) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorDetailString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorDetailString.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorDetailString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDetailString.Merge(m, src)
}
func (m *ErrorDetailString) XXX_Size() int {
	return m.Size()
}
func (m *ErrorDetailString) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDetailString.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDetailString proto.InternalMessageInfo

func (m *ErrorDetailString) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ErrorDetailString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// ErrorDetailInt64
type ErrorDetailInt64 struct {
	//
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	//
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorDetailInt64) Reset()         { *m = ErrorDetailInt64{} }
func (m *ErrorDetailInt64) String() string { return proto.CompactTextString(m) }
func (*ErrorDetailInt64) ProtoMessage()    {}
func (*ErrorDetailInt64) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{2}
}
func (m *ErrorDetailInt64) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorDetailInt64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorDetailInt64.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorDetailInt64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDetailInt64.Merge(m, src)
}
func (m *ErrorDetailInt64) XXX_Size() int {
	return m.Size()
}
func (m *ErrorDetailInt64) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDetailInt64.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDetailInt64 proto.InternalMessageInfo

func (m *ErrorDetailInt64) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ErrorDetailInt64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// ErrorDetailDouble
type ErrorDetailDouble struct {
	//
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	//
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorDetailDouble) Reset()         { *m = ErrorDetailDouble{} }
func (m *ErrorDetailDouble) String() string { return proto.CompactTextString(m) }
func (*ErrorDetailDouble) ProtoMessage()    {}
func (*ErrorDetailDouble) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{3}
}
func (m *ErrorDetailDouble) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorDetailDouble) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorDetailDouble.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorDetailDouble) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDetailDouble.Merge(m, src)
}
func (m *ErrorDetailDouble) XXX_Size() int {
	return m.Size()
}
func (m *ErrorDetailDouble) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDetailDouble.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDetailDouble proto.InternalMessageInfo

func (m *ErrorDetailDouble) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ErrorDetailDouble) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// ErrorDetailTimestamp
type ErrorDetailTimestamp struct {
	//
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	//
	Value                *types.Timestamp `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ErrorDetailTimestamp) Reset()         { *m = ErrorDetailTimestamp{} }
func (m *ErrorDetailTimestamp) String() string { return proto.CompactTextString(m) }
func (*ErrorDetailTimestamp) ProtoMessage()    {}
func (*ErrorDetailTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{4}
}
func (m *ErrorDetailTimestamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorDetailTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorDetailTimestamp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorDetailTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDetailTimestamp.Merge(m, src)
}
func (m *ErrorDetailTimestamp) XXX_Size() int {
	return m.Size()
}
func (m *ErrorDetailTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDetailTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDetailTimestamp proto.InternalMessageInfo

func (m *ErrorDetailTimestamp) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ErrorDetailTimestamp) GetValue() *types.Timestamp {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorDetailBool)(nil), "kitchen.errors.ErrorDetailBool")
	proto.RegisterType((*ErrorDetailString)(nil), "kitchen.errors.ErrorDetailString")
	proto.RegisterType((*ErrorDetailInt64)(nil), "kitchen.errors.ErrorDetailInt64")
	proto.RegisterType((*ErrorDetailDouble)(nil), "kitchen.errors.ErrorDetailDouble")
	proto.RegisterType((*ErrorDetailTimestamp)(nil), "kitchen.errors.ErrorDetailTimestamp")
}

func init() { proto.RegisterFile("errors.proto", fileDescriptor_24fe73c7f0ddb19c) }

var fileDescriptor_24fe73c7f0ddb19c = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2d, 0x2a, 0xca,
	0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xce, 0x2c, 0x49, 0xce, 0x48,
	0xcd, 0xd3, 0x83, 0x88, 0x4a, 0xc9, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x65, 0x93,
	0x4a, 0xd3, 0xf4, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x20, 0x1a, 0x94, 0x6c,
	0xb9, 0xf8, 0x5d, 0x41, 0x4a, 0x5d, 0x52, 0x4b, 0x12, 0x33, 0x73, 0x9c, 0xf2, 0xf3, 0x73, 0x84,
	0x44, 0xb8, 0x58, 0xd3, 0x32, 0x53, 0x73, 0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20,
	0x1c, 0x90, 0x68, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x84,
	0xa3, 0x64, 0xcf, 0x25, 0x88, 0xa4, 0x3d, 0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x9d, 0x18, 0x03, 0x38,
	0x61, 0x06, 0xd8, 0x71, 0x09, 0x20, 0x19, 0xe0, 0x99, 0x57, 0x62, 0x66, 0x42, 0x8c, 0x7e, 0x66,
	0xec, 0x0e, 0x70, 0xc9, 0x2f, 0x4d, 0xca, 0x49, 0x25, 0xc6, 0x00, 0x46, 0x98, 0x01, 0x71, 0x5c,
	0x22, 0x48, 0x06, 0x84, 0xc0, 0x82, 0x07, 0x87, 0x19, 0x06, 0xc8, 0x66, 0x70, 0x1b, 0x49, 0xe9,
	0x41, 0xc2, 0x57, 0x0f, 0x16, 0xbe, 0x7a, 0x70, 0x03, 0xa0, 0xe6, 0x3b, 0x49, 0x9d, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0x71,
	0xe8, 0x59, 0x43, 0x62, 0x27, 0x89, 0x0d, 0xac, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0xa5, 0xbf, 0x7f, 0xc4, 0x01, 0x00, 0x00,
}

func (m *ErrorDetailBool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorDetailBool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorDetailBool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value {
		i--
		if m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ErrorDetailString) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorDetailString) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorDetailString) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ErrorDetailInt64) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorDetailInt64) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorDetailInt64) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ErrorDetailDouble) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorDetailDouble) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorDetailDouble) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ErrorDetailTimestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorDetailTimestamp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorDetailTimestamp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintErrors(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ErrorDetailBool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.Value {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ErrorDetailString) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ErrorDetailInt64) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovErrors(uint64(m.Value))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ErrorDetailDouble) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.Value != 0 {
		n += 9
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ErrorDetailTimestamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovErrors(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ErrorDetailBool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: ErrorDetailBool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorDetailBool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ErrorDetailString) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: ErrorDetailString: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorDetailString: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ErrorDetailInt64) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: ErrorDetailInt64: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorDetailInt64: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ErrorDetailDouble) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: ErrorDetailDouble: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorDetailDouble: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ErrorDetailTimestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: ErrorDetailTimestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorDetailTimestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &types.Timestamp{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
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
				return 0, ErrInvalidLengthErrors
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrors
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrors
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrors        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrors = fmt.Errorf("proto: unexpected end of group")
)
