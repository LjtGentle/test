// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: resdata.option.proto

package osproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResOutput int32

const (
	ResOutput_Public ResOutput = 0
	ResOutput_Server ResOutput = 1
)

// Enum value maps for ResOutput.
var (
	ResOutput_name = map[int32]string{
		0: "Public",
		1: "Server",
	}
	ResOutput_value = map[string]int32{
		"Public": 0,
		"Server": 1,
	}
)

func (x ResOutput) Enum() *ResOutput {
	p := new(ResOutput)
	*p = x
	return p
}

func (x ResOutput) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResOutput) Descriptor() protoreflect.EnumDescriptor {
	return file_resdata_option_proto_enumTypes[0].Descriptor()
}

func (ResOutput) Type() protoreflect.EnumType {
	return &file_resdata_option_proto_enumTypes[0]
}

func (x ResOutput) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResOutput.Descriptor instead.
func (ResOutput) EnumDescriptor() ([]byte, []int) {
	return file_resdata_option_proto_rawDescGZIP(), []int{0}
}

var file_resdata_option_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ResOutput)(nil),
		Field:         50001,
		Name:          "resdata.output",
		Tag:           "varint,50001,opt,name=output,enum=resdata.ResOutput",
		Filename:      "resdata.option.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50011,
		Name:          "resdata.count",
		Tag:           "bytes,50011,opt,name=count",
		Filename:      "resdata.option.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50012,
		Name:          "resdata.referCname",
		Tag:           "bytes,50012,opt,name=referCname",
		Filename:      "resdata.option.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50013,
		Name:          "resdata.cname",
		Tag:           "bytes,50013,opt,name=cname",
		Filename:      "resdata.option.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50014,
		Name:          "resdata.bind",
		Tag:           "bytes,50014,opt,name=bind",
		Filename:      "resdata.option.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50015,
		Name:          "resdata.translate",
		Tag:           "varint,50015,opt,name=translate",
		Filename:      "resdata.option.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50021,
		Name:          "resdata.keyword_alias",
		Tag:           "bytes,50021,opt,name=keyword_alias",
		Filename:      "resdata.option.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional resdata.ResOutput output = 50001;
	E_Output = &file_resdata_option_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string count = 50011;
	E_Count = &file_resdata_option_proto_extTypes[1]
	// optional string referCname = 50012;
	E_ReferCname = &file_resdata_option_proto_extTypes[2]
	// optional string cname = 50013;
	E_Cname = &file_resdata_option_proto_extTypes[3]
	// optional string bind = 50014;
	E_Bind = &file_resdata_option_proto_extTypes[4] // 指定当前域绑定了哪个枚举
	// optional bool translate = 50015;
	E_Translate = &file_resdata_option_proto_extTypes[5] // 指定是否翻译这个域
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string keyword_alias = 50021;
	E_KeywordAlias = &file_resdata_option_proto_extTypes[6]
)

var File_resdata_option_proto protoreflect.FileDescriptor

var file_resdata_option_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x73, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x73, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x23, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0a,
	0x0a, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x10, 0x01, 0x3a, 0x4d, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x3a, 0x35, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdb, 0x86,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3f, 0x0a, 0x0a,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x43, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdc, 0x86, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x43, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x35, 0x0a,
	0x05, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdd, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x33, 0x0a, 0x04, 0x62, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xde, 0x86, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x69, 0x6e, 0x64, 0x3a, 0x3d, 0x0a, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdf, 0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x3a, 0x48, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe5, 0x86, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x41, 0x6c, 0x69,
	0x61, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x6f, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resdata_option_proto_rawDescOnce sync.Once
	file_resdata_option_proto_rawDescData = file_resdata_option_proto_rawDesc
)

func file_resdata_option_proto_rawDescGZIP() []byte {
	file_resdata_option_proto_rawDescOnce.Do(func() {
		file_resdata_option_proto_rawDescData = protoimpl.X.CompressGZIP(file_resdata_option_proto_rawDescData)
	})
	return file_resdata_option_proto_rawDescData
}

var file_resdata_option_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resdata_option_proto_goTypes = []interface{}{
	(ResOutput)(0),                        // 0: resdata.ResOutput
	(*descriptorpb.MessageOptions)(nil),   // 1: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),     // 2: google.protobuf.FieldOptions
	(*descriptorpb.EnumValueOptions)(nil), // 3: google.protobuf.EnumValueOptions
}
var file_resdata_option_proto_depIdxs = []int32{
	1, // 0: resdata.output:extendee -> google.protobuf.MessageOptions
	2, // 1: resdata.count:extendee -> google.protobuf.FieldOptions
	2, // 2: resdata.referCname:extendee -> google.protobuf.FieldOptions
	2, // 3: resdata.cname:extendee -> google.protobuf.FieldOptions
	2, // 4: resdata.bind:extendee -> google.protobuf.FieldOptions
	2, // 5: resdata.translate:extendee -> google.protobuf.FieldOptions
	3, // 6: resdata.keyword_alias:extendee -> google.protobuf.EnumValueOptions
	0, // 7: resdata.output:type_name -> resdata.ResOutput
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	7, // [7:8] is the sub-list for extension type_name
	0, // [0:7] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resdata_option_proto_init() }
func file_resdata_option_proto_init() {
	if File_resdata_option_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resdata_option_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 7,
			NumServices:   0,
		},
		GoTypes:           file_resdata_option_proto_goTypes,
		DependencyIndexes: file_resdata_option_proto_depIdxs,
		EnumInfos:         file_resdata_option_proto_enumTypes,
		ExtensionInfos:    file_resdata_option_proto_extTypes,
	}.Build()
	File_resdata_option_proto = out.File
	file_resdata_option_proto_rawDesc = nil
	file_resdata_option_proto_goTypes = nil
	file_resdata_option_proto_depIdxs = nil
}