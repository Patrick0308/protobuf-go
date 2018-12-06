// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/enum.proto

package proto2

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	prototype "github.com/golang/protobuf/v2/reflect/prototype"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// EnumType1 comment.
type EnumType1 int32

const (
	// EnumType1_ONE comment.
	EnumType1_ONE EnumType1 = 1
	// EnumType1_TWO comment.
	EnumType1_TWO EnumType1 = 2
)

type xxx_EnumType1 EnumType1

func (e EnumType1) ProtoReflect() protoreflect.Enum {
	return (xxx_EnumType1)(e)
}
func (e xxx_EnumType1) Type() protoreflect.EnumType {
	return xxx_Enum_ProtoFile_EnumTypes[0]
}
func (e xxx_EnumType1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var EnumType1_name = map[int32]string{
	1: "ONE",
	2: "TWO",
}

var EnumType1_value = map[string]int32{
	"ONE": 1,
	"TWO": 2,
}

func (x EnumType1) Enum() *EnumType1 {
	p := new(EnumType1)
	*p = x
	return p
}

func (x EnumType1) String() string {
	return proto.EnumName(EnumType1_name, int32(x))
}

func (x *EnumType1) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumType1_value, data, "EnumType1")
	if err != nil {
		return err
	}
	*x = EnumType1(value)
	return nil
}

func (EnumType1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0}
}

type EnumType2 int32

const (
	EnumType2_duplicate1 EnumType2 = 1
	EnumType2_duplicate2 EnumType2 = 1
)

type xxx_EnumType2 EnumType2

func (e EnumType2) ProtoReflect() protoreflect.Enum {
	return (xxx_EnumType2)(e)
}
func (e xxx_EnumType2) Type() protoreflect.EnumType {
	return xxx_Enum_ProtoFile_EnumTypes[1]
}
func (e xxx_EnumType2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var EnumType2_name = map[int32]string{
	1: "duplicate1",
	// Duplicate value: 1: "duplicate2",
}

var EnumType2_value = map[string]int32{
	"duplicate1": 1,
	"duplicate2": 1,
}

func (x EnumType2) Enum() *EnumType2 {
	p := new(EnumType2)
	*p = x
	return p
}

func (x EnumType2) String() string {
	return proto.EnumName(EnumType2_name, int32(x))
}

func (x *EnumType2) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumType2_value, data, "EnumType2")
	if err != nil {
		return err
	}
	*x = EnumType2(value)
	return nil
}

func (EnumType2) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{1}
}

// NestedEnumType1A comment.
type EnumContainerMessage1_NestedEnumType1A int32

const (
	// NestedEnumType1A_VALUE comment.
	EnumContainerMessage1_NESTED_1A_VALUE EnumContainerMessage1_NestedEnumType1A = 0
)

type xxx_EnumContainerMessage1_NestedEnumType1A EnumContainerMessage1_NestedEnumType1A

func (e EnumContainerMessage1_NestedEnumType1A) ProtoReflect() protoreflect.Enum {
	return (xxx_EnumContainerMessage1_NestedEnumType1A)(e)
}
func (e xxx_EnumContainerMessage1_NestedEnumType1A) Type() protoreflect.EnumType {
	return xxx_Enum_ProtoFile_EnumTypes[2]
}
func (e xxx_EnumContainerMessage1_NestedEnumType1A) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var EnumContainerMessage1_NestedEnumType1A_name = map[int32]string{
	0: "NESTED_1A_VALUE",
}

var EnumContainerMessage1_NestedEnumType1A_value = map[string]int32{
	"NESTED_1A_VALUE": 0,
}

func (x EnumContainerMessage1_NestedEnumType1A) Enum() *EnumContainerMessage1_NestedEnumType1A {
	p := new(EnumContainerMessage1_NestedEnumType1A)
	*p = x
	return p
}

func (x EnumContainerMessage1_NestedEnumType1A) String() string {
	return proto.EnumName(EnumContainerMessage1_NestedEnumType1A_name, int32(x))
}

func (x *EnumContainerMessage1_NestedEnumType1A) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumContainerMessage1_NestedEnumType1A_value, data, "EnumContainerMessage1_NestedEnumType1A")
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_NestedEnumType1A(value)
	return nil
}

func (EnumContainerMessage1_NestedEnumType1A) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 0}
}

type EnumContainerMessage1_NestedEnumType1B int32

const (
	EnumContainerMessage1_NESTED_1B_VALUE EnumContainerMessage1_NestedEnumType1B = 0
)

type xxx_EnumContainerMessage1_NestedEnumType1B EnumContainerMessage1_NestedEnumType1B

func (e EnumContainerMessage1_NestedEnumType1B) ProtoReflect() protoreflect.Enum {
	return (xxx_EnumContainerMessage1_NestedEnumType1B)(e)
}
func (e xxx_EnumContainerMessage1_NestedEnumType1B) Type() protoreflect.EnumType {
	return xxx_Enum_ProtoFile_EnumTypes[3]
}
func (e xxx_EnumContainerMessage1_NestedEnumType1B) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var EnumContainerMessage1_NestedEnumType1B_name = map[int32]string{
	0: "NESTED_1B_VALUE",
}

var EnumContainerMessage1_NestedEnumType1B_value = map[string]int32{
	"NESTED_1B_VALUE": 0,
}

func (x EnumContainerMessage1_NestedEnumType1B) Enum() *EnumContainerMessage1_NestedEnumType1B {
	p := new(EnumContainerMessage1_NestedEnumType1B)
	*p = x
	return p
}

func (x EnumContainerMessage1_NestedEnumType1B) String() string {
	return proto.EnumName(EnumContainerMessage1_NestedEnumType1B_name, int32(x))
}

func (x *EnumContainerMessage1_NestedEnumType1B) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumContainerMessage1_NestedEnumType1B_value, data, "EnumContainerMessage1_NestedEnumType1B")
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_NestedEnumType1B(value)
	return nil
}

func (EnumContainerMessage1_NestedEnumType1B) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 1}
}

// NestedEnumType2A comment.
type EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A int32

const (
	// NestedEnumType2A_VALUE comment.
	EnumContainerMessage1_EnumContainerMessage2_NESTED_2A_VALUE EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A = 0
)

type xxx_EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A

func (e EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) ProtoReflect() protoreflect.Enum {
	return (xxx_EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A)(e)
}
func (e xxx_EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Type() protoreflect.EnumType {
	return xxx_Enum_ProtoFile_EnumTypes[4]
}
func (e xxx_EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_name = map[int32]string{
	0: "NESTED_2A_VALUE",
}

var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_value = map[string]int32{
	"NESTED_2A_VALUE": 0,
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Enum() *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A {
	p := new(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A)
	*p = x
	return p
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) String() string {
	return proto.EnumName(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_name, int32(x))
}

func (x *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_value, data, "EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A")
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A(value)
	return nil
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 0, 0}
}

type EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B int32

const (
	EnumContainerMessage1_EnumContainerMessage2_NESTED_2B_VALUE EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B = 0
)

type xxx_EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B

func (e EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) ProtoReflect() protoreflect.Enum {
	return (xxx_EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B)(e)
}
func (e xxx_EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Type() protoreflect.EnumType {
	return xxx_Enum_ProtoFile_EnumTypes[5]
}
func (e xxx_EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_name = map[int32]string{
	0: "NESTED_2B_VALUE",
}

var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_value = map[string]int32{
	"NESTED_2B_VALUE": 0,
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Enum() *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B {
	p := new(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B)
	*p = x
	return p
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) String() string {
	return proto.EnumName(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_name, int32(x))
}

func (x *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_value, data, "EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B")
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B(value)
	return nil
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 0, 1}
}

type EnumContainerMessage1 struct {
	DefaultDuplicate1    *EnumType2 `protobuf:"varint,1,opt,name=default_duplicate1,json=defaultDuplicate1,enum=goproto.protoc.proto2.EnumType2,def=1" json:"default_duplicate1,omitempty"`
	DefaultDuplicate2    *EnumType2 `protobuf:"varint,2,opt,name=default_duplicate2,json=defaultDuplicate2,enum=goproto.protoc.proto2.EnumType2,def=1" json:"default_duplicate2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

type xxx_EnumContainerMessage1 struct{ m *EnumContainerMessage1 }

func (m *EnumContainerMessage1) ProtoReflect() protoreflect.Message {
	return xxx_EnumContainerMessage1{m}
}
func (m xxx_EnumContainerMessage1) Type() protoreflect.MessageType {
	return xxx_Enum_ProtoFile_MessageTypes[0].Type
}
func (m xxx_EnumContainerMessage1) KnownFields() protoreflect.KnownFields {
	return xxx_Enum_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_EnumContainerMessage1) UnknownFields() protoreflect.UnknownFields {
	return xxx_Enum_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_EnumContainerMessage1) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_EnumContainerMessage1) ProtoMutable() {}

func (m *EnumContainerMessage1) Reset()         { *m = EnumContainerMessage1{} }
func (m *EnumContainerMessage1) String() string { return proto.CompactTextString(m) }
func (*EnumContainerMessage1) ProtoMessage()    {}
func (*EnumContainerMessage1) Descriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0}
}

func (m *EnumContainerMessage1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumContainerMessage1.Unmarshal(m, b)
}
func (m *EnumContainerMessage1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumContainerMessage1.Marshal(b, m, deterministic)
}
func (m *EnumContainerMessage1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumContainerMessage1.Merge(m, src)
}
func (m *EnumContainerMessage1) XXX_Size() int {
	return xxx_messageInfo_EnumContainerMessage1.Size(m)
}
func (m *EnumContainerMessage1) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumContainerMessage1.DiscardUnknown(m)
}

var xxx_messageInfo_EnumContainerMessage1 proto.InternalMessageInfo

const Default_EnumContainerMessage1_DefaultDuplicate1 EnumType2 = EnumType2_duplicate1
const Default_EnumContainerMessage1_DefaultDuplicate2 EnumType2 = EnumType2_duplicate2

func (m *EnumContainerMessage1) GetDefaultDuplicate1() EnumType2 {
	if m != nil && m.DefaultDuplicate1 != nil {
		return *m.DefaultDuplicate1
	}
	return Default_EnumContainerMessage1_DefaultDuplicate1
}

func (m *EnumContainerMessage1) GetDefaultDuplicate2() EnumType2 {
	if m != nil && m.DefaultDuplicate2 != nil {
		return *m.DefaultDuplicate2
	}
	return Default_EnumContainerMessage1_DefaultDuplicate2
}

type EnumContainerMessage1_EnumContainerMessage2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_EnumContainerMessage1_EnumContainerMessage2 struct {
	m *EnumContainerMessage1_EnumContainerMessage2
}

func (m *EnumContainerMessage1_EnumContainerMessage2) ProtoReflect() protoreflect.Message {
	return xxx_EnumContainerMessage1_EnumContainerMessage2{m}
}
func (m xxx_EnumContainerMessage1_EnumContainerMessage2) Type() protoreflect.MessageType {
	return xxx_Enum_ProtoFile_MessageTypes[1].Type
}
func (m xxx_EnumContainerMessage1_EnumContainerMessage2) KnownFields() protoreflect.KnownFields {
	return xxx_Enum_ProtoFile_MessageTypes[1].KnownFieldsOf(m.m)
}
func (m xxx_EnumContainerMessage1_EnumContainerMessage2) UnknownFields() protoreflect.UnknownFields {
	return xxx_Enum_ProtoFile_MessageTypes[1].UnknownFieldsOf(m.m)
}
func (m xxx_EnumContainerMessage1_EnumContainerMessage2) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_EnumContainerMessage1_EnumContainerMessage2) ProtoMutable() {}

func (m *EnumContainerMessage1_EnumContainerMessage2) Reset() {
	*m = EnumContainerMessage1_EnumContainerMessage2{}
}
func (m *EnumContainerMessage1_EnumContainerMessage2) String() string {
	return proto.CompactTextString(m)
}
func (*EnumContainerMessage1_EnumContainerMessage2) ProtoMessage() {}
func (*EnumContainerMessage1_EnumContainerMessage2) Descriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 0}
}

func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.Unmarshal(m, b)
}
func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.Marshal(b, m, deterministic)
}
func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.Merge(m, src)
}
func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_Size() int {
	return xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.Size(m)
}
func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.DiscardUnknown(m)
}

var xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2 proto.InternalMessageInfo

func init() {
	proto.RegisterFile("proto2/enum.proto", fileDescriptor_de9f68860d540858)
	proto.RegisterEnum("goproto.protoc.proto2.EnumType1", EnumType1_name, EnumType1_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumType2", EnumType2_name, EnumType2_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumContainerMessage1_NestedEnumType1A", EnumContainerMessage1_NestedEnumType1A_name, EnumContainerMessage1_NestedEnumType1A_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumContainerMessage1_NestedEnumType1B", EnumContainerMessage1_NestedEnumType1B_name, EnumContainerMessage1_NestedEnumType1B_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A", EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_name, EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B", EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_name, EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_value)
	proto.RegisterType((*EnumContainerMessage1)(nil), "goproto.protoc.proto2.EnumContainerMessage1")
	proto.RegisterType((*EnumContainerMessage1_EnumContainerMessage2)(nil), "goproto.protoc.proto2.EnumContainerMessage1.EnumContainerMessage2")
}

var fileDescriptor_de9f68860d540858 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0xd1, 0x4b, 0xc2, 0x40,
	0x1c, 0x07, 0xf0, 0xb6, 0x09, 0xe9, 0x3d, 0xd8, 0xb9, 0x10, 0x44, 0x08, 0x64, 0x2f, 0x85, 0xe0,
	0xc6, 0xfd, 0x1e, 0x7b, 0x89, 0x99, 0xf7, 0x56, 0x4a, 0x6a, 0x06, 0xf5, 0x20, 0xe7, 0x76, 0x5e,
	0x03, 0x77, 0x27, 0xee, 0x2e, 0xe8, 0x9f, 0xe8, 0x6f, 0x0e, 0x5d, 0xcc, 0x45, 0x2b, 0xe8, 0x69,
	0xdf, 0xdf, 0x6f, 0xdf, 0x7d, 0x60, 0x3f, 0xd4, 0xda, 0xee, 0x94, 0x56, 0x10, 0x70, 0x69, 0x52,
	0xff, 0x90, 0xdd, 0xb6, 0x50, 0x87, 0x90, 0x8f, 0x51, 0xfe, 0x00, 0xef, 0xc3, 0x41, 0x6d, 0x2a,
	0x4d, 0x7a, 0xab, 0xa4, 0x66, 0x89, 0xe4, 0xbb, 0x7b, 0x9e, 0x65, 0x4c, 0x70, 0xe2, 0xbe, 0x20,
	0x37, 0xe6, 0x6b, 0x66, 0x36, 0x7a, 0x19, 0x9b, 0xed, 0x26, 0x89, 0x98, 0xe6, 0xa4, 0x63, 0xf5,
	0xac, 0xab, 0x26, 0xf4, 0xfc, 0x4a, 0xcd, 0xdf, 0x4b, 0xf3, 0xf7, 0x2d, 0x87, 0x6b, 0x74, 0xfc,
	0x66, 0xda, 0xfa, 0x72, 0x46, 0xc5, 0xaa, 0x12, 0x87, 0x8e, 0xfd, 0x6f, 0x1c, 0x7e, 0xe2, 0xd0,
	0x4d, 0xaa, 0x7f, 0x09, 0xbc, 0x4b, 0x84, 0xc7, 0x3c, 0xd3, 0x3c, 0x2e, 0xa8, 0xd0, 0x3d, 0x47,
	0x67, 0x63, 0x3a, 0x9b, 0xd3, 0xd1, 0x12, 0xc2, 0xe5, 0x22, 0xbc, 0x7b, 0xa4, 0xf8, 0xa4, 0xa2,
	0x38, 0x2c, 0x17, 0x87, 0xbf, 0x17, 0x49, 0x59, 0x24, 0x7f, 0x88, 0xa4, 0x2c, 0x92, 0x42, 0xec,
	0x5f, 0xa0, 0x46, 0x51, 0x71, 0x4f, 0x91, 0x33, 0x19, 0x53, 0x6c, 0xed, 0xc3, 0xfc, 0x69, 0x82,
	0xed, 0xfe, 0xc3, 0xf1, 0x35, 0xb8, 0x4d, 0x54, 0x3a, 0x33, 0xb6, 0xbe, 0xcd, 0x80, 0xad, 0xae,
	0x8d, 0x2d, 0xaf, 0x56, 0xb7, 0xb1, 0xed, 0xd5, 0xea, 0x0e, 0x76, 0xfa, 0x8d, 0x29, 0x9d, 0xd1,
	0xe9, 0x82, 0x8e, 0xc8, 0x31, 0xc2, 0x30, 0x7c, 0xbe, 0x11, 0x89, 0x7e, 0x35, 0x2b, 0x3f, 0x52,
	0x69, 0x20, 0xd4, 0x86, 0x49, 0x11, 0x1c, 0x6e, 0xbe, 0x32, 0xeb, 0xe0, 0x0d, 0x82, 0x28, 0x8d,
	0xf3, 0x39, 0x1a, 0x08, 0x2e, 0x07, 0x42, 0x05, 0x9a, 0x67, 0x3a, 0x66, 0x9a, 0xe5, 0x6b, 0xf8,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xcb, 0xe0, 0x01, 0x70, 0x02, 0x00, 0x00,
}

func init() {
	xxx_Enum_ProtoFile_FileDesc.Enums = xxx_Enum_ProtoFile_EnumDescs[0:2]
	xxx_Enum_ProtoFile_FileDesc.Messages = xxx_Enum_ProtoFile_MessageDescs[0:1]
	xxx_Enum_ProtoFile_MessageDescs[0].Enums = xxx_Enum_ProtoFile_EnumDescs[2:4]
	xxx_Enum_ProtoFile_MessageDescs[0].Messages = xxx_Enum_ProtoFile_MessageDescs[1:2]
	xxx_Enum_ProtoFile_MessageDescs[1].Enums = xxx_Enum_ProtoFile_EnumDescs[4:6]
	xxx_Enum_ProtoFile_MessageDescs[0].Fields[0].EnumType = xxx_Enum_ProtoFile_EnumTypes[1]
	xxx_Enum_ProtoFile_MessageDescs[0].Fields[1].EnumType = xxx_Enum_ProtoFile_EnumTypes[1]
	var err error
	Enum_ProtoFile, err = prototype.NewFile(&xxx_Enum_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var Enum_ProtoFile protoreflect.FileDescriptor

var xxx_Enum_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto2,
	Path:    "proto2/enum.proto",
	Package: "goproto.protoc.proto2",
}
var xxx_Enum_ProtoFile_EnumTypes = [6]protoreflect.EnumType{
	prototype.GoEnum(
		xxx_Enum_ProtoFile_EnumDescs[0].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return EnumType1(n)
		},
	),
	prototype.GoEnum(
		xxx_Enum_ProtoFile_EnumDescs[1].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return EnumType2(n)
		},
	),
	prototype.GoEnum(
		xxx_Enum_ProtoFile_EnumDescs[2].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return EnumContainerMessage1_NestedEnumType1A(n)
		},
	),
	prototype.GoEnum(
		xxx_Enum_ProtoFile_EnumDescs[3].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return EnumContainerMessage1_NestedEnumType1B(n)
		},
	),
	prototype.GoEnum(
		xxx_Enum_ProtoFile_EnumDescs[4].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A(n)
		},
	),
	prototype.GoEnum(
		xxx_Enum_ProtoFile_EnumDescs[5].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B(n)
		},
	),
}
var xxx_Enum_ProtoFile_EnumDescs = [6]prototype.Enum{
	{
		Name: "EnumType1",
		Values: []prototype.EnumValue{
			{Name: "ONE", Number: 1},
			{Name: "TWO", Number: 2},
		},
	},
	{
		Name: "EnumType2",
		Values: []prototype.EnumValue{
			{Name: "duplicate1", Number: 1},
			{Name: "duplicate2", Number: 1},
		},
		ReservedNames:  []protoreflect.Name{"RESERVED1", "RESERVED2"},
		ReservedRanges: [][2]protoreflect.EnumNumber{{2, 2}, {3, 3}},
	},
	{
		Name: "NestedEnumType1A",
		Values: []prototype.EnumValue{
			{Name: "NESTED_1A_VALUE", Number: 0},
		},
	},
	{
		Name: "NestedEnumType1B",
		Values: []prototype.EnumValue{
			{Name: "NESTED_1B_VALUE", Number: 0},
		},
	},
	{
		Name: "NestedEnumType2A",
		Values: []prototype.EnumValue{
			{Name: "NESTED_2A_VALUE", Number: 0},
		},
	},
	{
		Name: "NestedEnumType2B",
		Values: []prototype.EnumValue{
			{Name: "NESTED_2B_VALUE", Number: 0},
		},
	},
}
var xxx_Enum_ProtoFile_MessageTypes = [2]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_Enum_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(EnumContainerMessage1)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_Enum_ProtoFile_MessageDescs[1].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(EnumContainerMessage1_EnumContainerMessage2)
		},
	)},
}
var xxx_Enum_ProtoFile_MessageDescs = [2]prototype.Message{
	{
		Name: "EnumContainerMessage1",
		Fields: []prototype.Field{
			{
				Name:        "default_duplicate1",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.EnumKind,
				JSONName:    "defaultDuplicate1",
				Default:     protoreflect.ValueOf(string("duplicate1")),
			},
			{
				Name:        "default_duplicate2",
				Number:      2,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.EnumKind,
				JSONName:    "defaultDuplicate2",
				Default:     protoreflect.ValueOf(string("duplicate2")),
			},
		},
	},
	{
		Name: "EnumContainerMessage2",
	},
}
