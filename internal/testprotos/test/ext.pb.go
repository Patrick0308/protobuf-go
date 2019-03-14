// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/ext.proto

package test

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var xxx_File_test_ext_proto_extDescs = []proto.ExtensionDesc{
	{
		ExtendedType:  (*TestAllExtensions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         2000,
		Name:          "goproto.proto.test.foreign_int32_extension",
		Tag:           "varint,2000,opt,name=foreign_int32_extension",
		Filename:      "test/ext.proto",
	},
}
var (
	// extend goproto.proto.test.TestAllExtensions { optional int32 foreign_int32_extension = 2000; }
	E_ForeignInt32Extension = &xxx_File_test_ext_proto_extDescs[0]
)

func init() {
	proto.RegisterFile("test/ext.proto", xxx_File_test_ext_proto_rawdesc_gzipped)
	proto.RegisterExtension(E_ForeignInt32Extension)
}

var xxx_File_test_ext_proto_rawdesc = []byte{
	// 207 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x5e, 0x0a, 0x17, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e,
	0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15,
	0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74,
}

var xxx_File_test_ext_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_test_ext_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_test_ext_proto protoreflect.FileDescriptor

var xxx_File_test_ext_proto_goTypes = []interface{}{
	(*TestAllExtensions)(nil), // 0: goproto.proto.test.TestAllExtensions
}
var xxx_File_test_ext_proto_depIdxs = []int32{
	0, // goproto.proto.test.foreign_int32_extension:extendee -> goproto.proto.test.TestAllExtensions
}

func init() { xxx_File_test_ext_proto_init() }
func xxx_File_test_ext_proto_init() {
	if File_test_ext_proto != nil {
		return
	}
	xxx_File_test_test_proto_init()
	extensionTypes := make([]protoreflect.ExtensionType, 1)
	File_test_ext_proto = protoimpl.FileBuilder{
		RawDescriptor:        xxx_File_test_ext_proto_rawdesc,
		GoTypes:              xxx_File_test_ext_proto_goTypes,
		DependencyIndexes:    xxx_File_test_ext_proto_depIdxs,
		LegacyExtensions:     xxx_File_test_ext_proto_extDescs,
		ExtensionOutputTypes: extensionTypes,
	}.Init()
	xxx_File_test_ext_proto_goTypes = nil
	xxx_File_test_ext_proto_depIdxs = nil
}
