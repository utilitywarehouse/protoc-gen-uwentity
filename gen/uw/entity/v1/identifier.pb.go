// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: uw/entity/v1/identifier.proto

package entity

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_uw_entity_v1_identifier_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         75630,
		Name:          "uw.entity.v1.ignore",
		Tag:           "varint,75630,opt,name=ignore",
		Filename:      "uw/entity/v1/identifier.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         75530,
		Name:          "uw.entity.v1.identifier",
		Tag:           "varint,75530,opt,name=identifier",
		Filename:      "uw/entity/v1/identifier.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Skip checking or generating for this message
	//
	// optional bool ignore = 75630;
	E_Ignore = &file_uw_entity_v1_identifier_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// Mark as the field you want to generate the getter for
	//
	// optional bool identifier = 75530;
	E_Identifier = &file_uw_entity_v1_identifier_proto_extTypes[1]
)

var File_uw_entity_v1_identifier_proto protoreflect.FileDescriptor

var file_uw_entity_v1_identifier_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x75, 0x77, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x75, 0x77, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a,
	0x39, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xee, 0xce, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x3a, 0x3f, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a, 0xce, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x49, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x75, 0x77, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x75, 0x77, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_uw_entity_v1_identifier_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
}
var file_uw_entity_v1_identifier_proto_depIdxs = []int32{
	0, // 0: uw.entity.v1.ignore:extendee -> google.protobuf.MessageOptions
	1, // 1: uw.entity.v1.identifier:extendee -> google.protobuf.FieldOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_uw_entity_v1_identifier_proto_init() }
func file_uw_entity_v1_identifier_proto_init() {
	if File_uw_entity_v1_identifier_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_uw_entity_v1_identifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_uw_entity_v1_identifier_proto_goTypes,
		DependencyIndexes: file_uw_entity_v1_identifier_proto_depIdxs,
		ExtensionInfos:    file_uw_entity_v1_identifier_proto_extTypes,
	}.Build()
	File_uw_entity_v1_identifier_proto = out.File
	file_uw_entity_v1_identifier_proto_rawDesc = nil
	file_uw_entity_v1_identifier_proto_goTypes = nil
	file_uw_entity_v1_identifier_proto_depIdxs = nil
}
