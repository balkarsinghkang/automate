// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.0
// source: external/annotations/iam/annotations.proto

package iam

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

var file_external_annotations_iam_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*PolicyInfo)(nil),
		Field:         50001,
		Name:          "chef.automate.api.iam.policy",
		Tag:           "bytes,50001,opt,name=policy",
		Filename:      "external/annotations/iam/annotations.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// ... 50000-9999.
	// This range is reserved for internal use within individual organizations, so
	// you can use numbers in this range freely for in-house applications. If you
	// intend to use custom options in public applications, however, then it is
	// important that you make sure that your field numbers are globally unique.
	// To obtain globally unique field numbers, please send a request to
	// protobuf-global-extension-registry@google.com. Simply provide your project
	// name (e.g. Object-C plugin) and your project website (if available).
	// Usually you only need one extension number.
	// See https://developers.google.com/protocol-buffers/docs/proto#customoptions
	//
	// optional chef.automate.api.iam.PolicyInfo policy = 50001;
	E_Policy = &file_external_annotations_iam_annotations_proto_extTypes[0]
)

var File_external_annotations_iam_annotations_proto protoreflect.FileDescriptor

var file_external_annotations_iam_annotations_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x61, 0x6d, 0x1a, 0x25, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x5b, 0x0a, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69,
	0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_external_annotations_iam_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.MethodOptions)(nil), // 0: google.protobuf.MethodOptions
	(*PolicyInfo)(nil),                 // 1: chef.automate.api.iam.PolicyInfo
}
var file_external_annotations_iam_annotations_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.iam.policy:extendee -> google.protobuf.MethodOptions
	1, // 1: chef.automate.api.iam.policy:type_name -> chef.automate.api.iam.PolicyInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_annotations_iam_annotations_proto_init() }
func file_external_annotations_iam_annotations_proto_init() {
	if File_external_annotations_iam_annotations_proto != nil {
		return
	}
	file_external_annotations_iam_policy_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_annotations_iam_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_external_annotations_iam_annotations_proto_goTypes,
		DependencyIndexes: file_external_annotations_iam_annotations_proto_depIdxs,
		ExtensionInfos:    file_external_annotations_iam_annotations_proto_extTypes,
	}.Build()
	File_external_annotations_iam_annotations_proto = out.File
	file_external_annotations_iam_annotations_proto_rawDesc = nil
	file_external_annotations_iam_annotations_proto_goTypes = nil
	file_external_annotations_iam_annotations_proto_depIdxs = nil
}
