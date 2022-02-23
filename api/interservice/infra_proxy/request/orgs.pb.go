// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: interservice/infra_proxy/request/orgs.proto

package request

import (
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

type CreateOrg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef organization name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,5,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// List of projects this chef organization belongs to. May be empty.
	Projects []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
}

func (x *CreateOrg) Reset() {
	*x = CreateOrg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrg) ProtoMessage() {}

func (x *CreateOrg) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrg.ProtoReflect.Descriptor instead.
func (*CreateOrg) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_orgs_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOrg) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CreateOrg) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

type UpdateOrg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef organization name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Chef organization admin user.
	AdminUser string `protobuf:"bytes,3,opt,name=admin_user,json=adminUser,proto3" json:"admin_user,omitempty" toml:"admin_user,omitempty" mapstructure:"admin_user,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,4,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// List of projects this chef organization belongs to. May be empty.
	Projects []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
}

func (x *UpdateOrg) Reset() {
	*x = UpdateOrg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrg) ProtoMessage() {}

func (x *UpdateOrg) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrg.ProtoReflect.Descriptor instead.
func (*UpdateOrg) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_orgs_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateOrg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateOrg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateOrg) GetAdminUser() string {
	if x != nil {
		return x.AdminUser
	}
	return ""
}

func (x *UpdateOrg) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *UpdateOrg) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

type DeleteOrg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
}

func (x *DeleteOrg) Reset() {
	*x = DeleteOrg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrg) ProtoMessage() {}

func (x *DeleteOrg) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrg.ProtoReflect.Descriptor instead.
func (*DeleteOrg) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_orgs_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteOrg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteOrg) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type GetOrgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
}

func (x *GetOrgs) Reset() {
	*x = GetOrgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrgs) ProtoMessage() {}

func (x *GetOrgs) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrgs.ProtoReflect.Descriptor instead.
func (*GetOrgs) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_orgs_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrgs) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type GetOrg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
}

func (x *GetOrg) Reset() {
	*x = GetOrg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrg) ProtoMessage() {}

func (x *GetOrg) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrg.ProtoReflect.Descriptor instead.
func (*GetOrg) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_orgs_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetOrg) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type ResetOrgAdminKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Chef organization admin key.
	AdminKey string `protobuf:"bytes,3,opt,name=admin_key,json=adminKey,proto3" json:"admin_key,omitempty" toml:"admin_key,omitempty" mapstructure:"admin_key,omitempty"`
}

func (x *ResetOrgAdminKey) Reset() {
	*x = ResetOrgAdminKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetOrgAdminKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetOrgAdminKey) ProtoMessage() {}

func (x *ResetOrgAdminKey) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetOrgAdminKey.ProtoReflect.Descriptor instead.
func (*ResetOrgAdminKey) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_orgs_proto_rawDescGZIP(), []int{5}
}

func (x *ResetOrgAdminKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResetOrgAdminKey) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *ResetOrgAdminKey) GetAdminKey() string {
	if x != nil {
		return x.AdminKey
	}
	return ""
}

type GetInfraServerOrgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Automate Infra Server ID
	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
}

func (x *GetInfraServerOrgs) Reset() {
	*x = GetInfraServerOrgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfraServerOrgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfraServerOrgs) ProtoMessage() {}

func (x *GetInfraServerOrgs) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_orgs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfraServerOrgs.ProtoReflect.Descriptor instead.
func (*GetInfraServerOrgs) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_orgs_proto_rawDescGZIP(), []int{6}
}

func (x *GetInfraServerOrgs) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

var File_interservice_infra_proxy_request_orgs_proto protoreflect.FileDescriptor

var file_interservice_infra_proxy_request_orgs_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x6f, 0x72, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x68, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x35, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4f, 0x72, 0x67,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x22, 0x31, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_infra_proxy_request_orgs_proto_rawDescOnce sync.Once
	file_interservice_infra_proxy_request_orgs_proto_rawDescData = file_interservice_infra_proxy_request_orgs_proto_rawDesc
)

func file_interservice_infra_proxy_request_orgs_proto_rawDescGZIP() []byte {
	file_interservice_infra_proxy_request_orgs_proto_rawDescOnce.Do(func() {
		file_interservice_infra_proxy_request_orgs_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_infra_proxy_request_orgs_proto_rawDescData)
	})
	return file_interservice_infra_proxy_request_orgs_proto_rawDescData
}

var file_interservice_infra_proxy_request_orgs_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_interservice_infra_proxy_request_orgs_proto_goTypes = []interface{}{
	(*CreateOrg)(nil),          // 0: chef.automate.domain.infra_proxy.request.CreateOrg
	(*UpdateOrg)(nil),          // 1: chef.automate.domain.infra_proxy.request.UpdateOrg
	(*DeleteOrg)(nil),          // 2: chef.automate.domain.infra_proxy.request.DeleteOrg
	(*GetOrgs)(nil),            // 3: chef.automate.domain.infra_proxy.request.GetOrgs
	(*GetOrg)(nil),             // 4: chef.automate.domain.infra_proxy.request.GetOrg
	(*ResetOrgAdminKey)(nil),   // 5: chef.automate.domain.infra_proxy.request.ResetOrgAdminKey
	(*GetInfraServerOrgs)(nil), // 6: chef.automate.domain.infra_proxy.request.GetInfraServerOrgs
}
var file_interservice_infra_proxy_request_orgs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_infra_proxy_request_orgs_proto_init() }
func file_interservice_infra_proxy_request_orgs_proto_init() {
	if File_interservice_infra_proxy_request_orgs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_infra_proxy_request_orgs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrg); i {
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
		file_interservice_infra_proxy_request_orgs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrg); i {
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
		file_interservice_infra_proxy_request_orgs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrg); i {
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
		file_interservice_infra_proxy_request_orgs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrgs); i {
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
		file_interservice_infra_proxy_request_orgs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrg); i {
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
		file_interservice_infra_proxy_request_orgs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetOrgAdminKey); i {
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
		file_interservice_infra_proxy_request_orgs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfraServerOrgs); i {
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
			RawDescriptor: file_interservice_infra_proxy_request_orgs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_infra_proxy_request_orgs_proto_goTypes,
		DependencyIndexes: file_interservice_infra_proxy_request_orgs_proto_depIdxs,
		MessageInfos:      file_interservice_infra_proxy_request_orgs_proto_msgTypes,
	}.Build()
	File_interservice_infra_proxy_request_orgs_proto = out.File
	file_interservice_infra_proxy_request_orgs_proto_rawDesc = nil
	file_interservice_infra_proxy_request_orgs_proto_goTypes = nil
	file_interservice_infra_proxy_request_orgs_proto_depIdxs = nil
}
