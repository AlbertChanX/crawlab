// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: services/model_delegate.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_model_delegate_proto protoreflect.FileDescriptor

var file_services_model_delegate_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x14, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x36, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x02, 0x44, 0x6f, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_model_delegate_proto_goTypes = []any{
	(*Request)(nil),  // 0: grpc.Request
	(*Response)(nil), // 1: grpc.Response
}
var file_services_model_delegate_proto_depIdxs = []int32{
	0, // 0: grpc.ModelDelegate.Do:input_type -> grpc.Request
	1, // 1: grpc.ModelDelegate.Do:output_type -> grpc.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_model_delegate_proto_init() }
func file_services_model_delegate_proto_init() {
	if File_services_model_delegate_proto != nil {
		return
	}
	file_entity_request_proto_init()
	file_entity_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_model_delegate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_model_delegate_proto_goTypes,
		DependencyIndexes: file_services_model_delegate_proto_depIdxs,
	}.Build()
	File_services_model_delegate_proto = out.File
	file_services_model_delegate_proto_rawDesc = nil
	file_services_model_delegate_proto_goTypes = nil
	file_services_model_delegate_proto_depIdxs = nil
}
