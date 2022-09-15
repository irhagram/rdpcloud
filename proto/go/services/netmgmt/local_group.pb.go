// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: services/netmgmt/local_group.proto

package netmgmt

import (
	netmgmt "github.com/s77rt/rdpcloud/proto/go/models/netmgmt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddUserToLocalGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *netmgmt.User_1       `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	LocalGroup *netmgmt.LocalGroup_1 `protobuf:"bytes,2,opt,name=local_group,json=localGroup,proto3" json:"local_group,omitempty"`
}

func (x *AddUserToLocalGroupRequest) Reset() {
	*x = AddUserToLocalGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_local_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToLocalGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToLocalGroupRequest) ProtoMessage() {}

func (x *AddUserToLocalGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_local_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToLocalGroupRequest.ProtoReflect.Descriptor instead.
func (*AddUserToLocalGroupRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_local_group_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserToLocalGroupRequest) GetUser() *netmgmt.User_1 {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *AddUserToLocalGroupRequest) GetLocalGroup() *netmgmt.LocalGroup_1 {
	if x != nil {
		return x.LocalGroup
	}
	return nil
}

type AddUserToLocalGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddUserToLocalGroupResponse) Reset() {
	*x = AddUserToLocalGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_local_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToLocalGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToLocalGroupResponse) ProtoMessage() {}

func (x *AddUserToLocalGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_local_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToLocalGroupResponse.ProtoReflect.Descriptor instead.
func (*AddUserToLocalGroupResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_local_group_proto_rawDescGZIP(), []int{1}
}

type RemoveUserFromLocalGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *netmgmt.User_1       `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	LocalGroup *netmgmt.LocalGroup_1 `protobuf:"bytes,2,opt,name=local_group,json=localGroup,proto3" json:"local_group,omitempty"`
}

func (x *RemoveUserFromLocalGroupRequest) Reset() {
	*x = RemoveUserFromLocalGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_local_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromLocalGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromLocalGroupRequest) ProtoMessage() {}

func (x *RemoveUserFromLocalGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_local_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromLocalGroupRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserFromLocalGroupRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_local_group_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveUserFromLocalGroupRequest) GetUser() *netmgmt.User_1 {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RemoveUserFromLocalGroupRequest) GetLocalGroup() *netmgmt.LocalGroup_1 {
	if x != nil {
		return x.LocalGroup
	}
	return nil
}

type RemoveUserFromLocalGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveUserFromLocalGroupResponse) Reset() {
	*x = RemoveUserFromLocalGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_local_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromLocalGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromLocalGroupResponse) ProtoMessage() {}

func (x *RemoveUserFromLocalGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_local_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromLocalGroupResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserFromLocalGroupResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_local_group_proto_rawDescGZIP(), []int{3}
}

type GetLocalGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLocalGroupsRequest) Reset() {
	*x = GetLocalGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_local_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocalGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocalGroupsRequest) ProtoMessage() {}

func (x *GetLocalGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_local_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocalGroupsRequest.ProtoReflect.Descriptor instead.
func (*GetLocalGroupsRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_local_group_proto_rawDescGZIP(), []int{4}
}

type GetLocalGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalGroups []*netmgmt.LocalGroup `protobuf:"bytes,1,rep,name=local_groups,json=localGroups,proto3" json:"local_groups,omitempty"`
}

func (x *GetLocalGroupsResponse) Reset() {
	*x = GetLocalGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_local_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocalGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocalGroupsResponse) ProtoMessage() {}

func (x *GetLocalGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_local_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocalGroupsResponse.ProtoReflect.Descriptor instead.
func (*GetLocalGroupsResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_local_group_proto_rawDescGZIP(), []int{5}
}

func (x *GetLocalGroupsResponse) GetLocalGroups() []*netmgmt.LocalGroup {
	if x != nil {
		return x.LocalGroups
	}
	return nil
}

type GetUsersInLocalGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalGroup *netmgmt.LocalGroup_1 `protobuf:"bytes,1,opt,name=local_group,json=localGroup,proto3" json:"local_group,omitempty"`
}

func (x *GetUsersInLocalGroupRequest) Reset() {
	*x = GetUsersInLocalGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_local_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersInLocalGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersInLocalGroupRequest) ProtoMessage() {}

func (x *GetUsersInLocalGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_local_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersInLocalGroupRequest.ProtoReflect.Descriptor instead.
func (*GetUsersInLocalGroupRequest) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_local_group_proto_rawDescGZIP(), []int{6}
}

func (x *GetUsersInLocalGroupRequest) GetLocalGroup() *netmgmt.LocalGroup_1 {
	if x != nil {
		return x.LocalGroup
	}
	return nil
}

type GetUsersInLocalGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*netmgmt.User_1 `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersInLocalGroupResponse) Reset() {
	*x = GetUsersInLocalGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_netmgmt_local_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersInLocalGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersInLocalGroupResponse) ProtoMessage() {}

func (x *GetUsersInLocalGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_netmgmt_local_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersInLocalGroupResponse.ProtoReflect.Descriptor instead.
func (*GetUsersInLocalGroupResponse) Descriptor() ([]byte, []int) {
	return file_services_netmgmt_local_group_proto_rawDescGZIP(), []int{7}
}

func (x *GetUsersInLocalGroupResponse) GetUsers() []*netmgmt.User_1 {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_services_netmgmt_local_group_proto protoreflect.FileDescriptor

var file_services_netmgmt_local_group_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x1a, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6e,
	0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x31, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x3d,
	0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x31, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x1d, 0x0a,
	0x1b, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8c, 0x01, 0x0a,
	0x1f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x5f, 0x31, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0b,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x31, 0x52,
	0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x22, 0x0a, 0x20, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x22, 0x5c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3d, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x31, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22,
	0x4c, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x5f, 0x31, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x37, 0x37, 0x72,
	0x74, 0x2f, 0x72, 0x64, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74,
	0x6d, 0x67, 0x6d, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_netmgmt_local_group_proto_rawDescOnce sync.Once
	file_services_netmgmt_local_group_proto_rawDescData = file_services_netmgmt_local_group_proto_rawDesc
)

func file_services_netmgmt_local_group_proto_rawDescGZIP() []byte {
	file_services_netmgmt_local_group_proto_rawDescOnce.Do(func() {
		file_services_netmgmt_local_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_netmgmt_local_group_proto_rawDescData)
	})
	return file_services_netmgmt_local_group_proto_rawDescData
}

var file_services_netmgmt_local_group_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_netmgmt_local_group_proto_goTypes = []interface{}{
	(*AddUserToLocalGroupRequest)(nil),       // 0: services.netmgmt.AddUserToLocalGroupRequest
	(*AddUserToLocalGroupResponse)(nil),      // 1: services.netmgmt.AddUserToLocalGroupResponse
	(*RemoveUserFromLocalGroupRequest)(nil),  // 2: services.netmgmt.RemoveUserFromLocalGroupRequest
	(*RemoveUserFromLocalGroupResponse)(nil), // 3: services.netmgmt.RemoveUserFromLocalGroupResponse
	(*GetLocalGroupsRequest)(nil),            // 4: services.netmgmt.GetLocalGroupsRequest
	(*GetLocalGroupsResponse)(nil),           // 5: services.netmgmt.GetLocalGroupsResponse
	(*GetUsersInLocalGroupRequest)(nil),      // 6: services.netmgmt.GetUsersInLocalGroupRequest
	(*GetUsersInLocalGroupResponse)(nil),     // 7: services.netmgmt.GetUsersInLocalGroupResponse
	(*netmgmt.User_1)(nil),                   // 8: models.netmgmt.User_1
	(*netmgmt.LocalGroup_1)(nil),             // 9: models.netmgmt.LocalGroup_1
	(*netmgmt.LocalGroup)(nil),               // 10: models.netmgmt.LocalGroup
}
var file_services_netmgmt_local_group_proto_depIdxs = []int32{
	8,  // 0: services.netmgmt.AddUserToLocalGroupRequest.user:type_name -> models.netmgmt.User_1
	9,  // 1: services.netmgmt.AddUserToLocalGroupRequest.local_group:type_name -> models.netmgmt.LocalGroup_1
	8,  // 2: services.netmgmt.RemoveUserFromLocalGroupRequest.user:type_name -> models.netmgmt.User_1
	9,  // 3: services.netmgmt.RemoveUserFromLocalGroupRequest.local_group:type_name -> models.netmgmt.LocalGroup_1
	10, // 4: services.netmgmt.GetLocalGroupsResponse.local_groups:type_name -> models.netmgmt.LocalGroup
	9,  // 5: services.netmgmt.GetUsersInLocalGroupRequest.local_group:type_name -> models.netmgmt.LocalGroup_1
	8,  // 6: services.netmgmt.GetUsersInLocalGroupResponse.users:type_name -> models.netmgmt.User_1
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_services_netmgmt_local_group_proto_init() }
func file_services_netmgmt_local_group_proto_init() {
	if File_services_netmgmt_local_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_netmgmt_local_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToLocalGroupRequest); i {
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
		file_services_netmgmt_local_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToLocalGroupResponse); i {
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
		file_services_netmgmt_local_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromLocalGroupRequest); i {
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
		file_services_netmgmt_local_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromLocalGroupResponse); i {
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
		file_services_netmgmt_local_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocalGroupsRequest); i {
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
		file_services_netmgmt_local_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocalGroupsResponse); i {
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
		file_services_netmgmt_local_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersInLocalGroupRequest); i {
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
		file_services_netmgmt_local_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersInLocalGroupResponse); i {
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
			RawDescriptor: file_services_netmgmt_local_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_netmgmt_local_group_proto_goTypes,
		DependencyIndexes: file_services_netmgmt_local_group_proto_depIdxs,
		MessageInfos:      file_services_netmgmt_local_group_proto_msgTypes,
	}.Build()
	File_services_netmgmt_local_group_proto = out.File
	file_services_netmgmt_local_group_proto_rawDesc = nil
	file_services_netmgmt_local_group_proto_goTypes = nil
	file_services_netmgmt_local_group_proto_depIdxs = nil
}
