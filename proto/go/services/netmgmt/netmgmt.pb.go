// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: services/netmgmt/netmgmt.proto

package netmgmt

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

var File_services_netmgmt_netmgmt_proto protoreflect.FileDescriptor

var file_services_netmgmt_netmgmt_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67,
	0x6d, 0x74, 0x1a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74,
	0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xb0, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x12, 0x53, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x37, 0x37, 0x72, 0x74, 0x2f, 0x72, 0x64, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_services_netmgmt_netmgmt_proto_goTypes = []interface{}{
	(*GetUsersRequest)(nil),  // 0: services.netmgmt.GetUsersRequest
	(*AddUserRequest)(nil),   // 1: services.netmgmt.AddUserRequest
	(*GetUsersResponse)(nil), // 2: services.netmgmt.GetUsersResponse
	(*AddUserResponse)(nil),  // 3: services.netmgmt.AddUserResponse
}
var file_services_netmgmt_netmgmt_proto_depIdxs = []int32{
	0, // 0: services.netmgmt.Netmgmt.GetUsers:input_type -> services.netmgmt.GetUsersRequest
	1, // 1: services.netmgmt.Netmgmt.AddUser:input_type -> services.netmgmt.AddUserRequest
	2, // 2: services.netmgmt.Netmgmt.GetUsers:output_type -> services.netmgmt.GetUsersResponse
	3, // 3: services.netmgmt.Netmgmt.AddUser:output_type -> services.netmgmt.AddUserResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_netmgmt_netmgmt_proto_init() }
func file_services_netmgmt_netmgmt_proto_init() {
	if File_services_netmgmt_netmgmt_proto != nil {
		return
	}
	file_services_netmgmt_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_netmgmt_netmgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_netmgmt_netmgmt_proto_goTypes,
		DependencyIndexes: file_services_netmgmt_netmgmt_proto_depIdxs,
	}.Build()
	File_services_netmgmt_netmgmt_proto = out.File
	file_services_netmgmt_netmgmt_proto_rawDesc = nil
	file_services_netmgmt_netmgmt_proto_goTypes = nil
	file_services_netmgmt_netmgmt_proto_depIdxs = nil
}
