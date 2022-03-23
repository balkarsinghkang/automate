// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: interservice/infra_proxy/request/automate_infra_server_users.proto

package request

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AutomateInfraServerUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Automate Infra Server ID.
	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
}

func (x *AutomateInfraServerUsers) Reset() {
	*x = AutomateInfraServerUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutomateInfraServerUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutomateInfraServerUsers) ProtoMessage() {}

func (x *AutomateInfraServerUsers) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutomateInfraServerUsers.ProtoReflect.Descriptor instead.
func (*AutomateInfraServerUsers) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescGZIP(), []int{0}
}

func (x *AutomateInfraServerUsers) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type AutomateInfraOrgUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Automate Infra Server ID.
	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Automate Infra Org ID.
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
}

func (x *AutomateInfraOrgUsers) Reset() {
	*x = AutomateInfraOrgUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutomateInfraOrgUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutomateInfraOrgUsers) ProtoMessage() {}

func (x *AutomateInfraOrgUsers) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutomateInfraOrgUsers.ProtoReflect.Descriptor instead.
func (*AutomateInfraOrgUsers) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescGZIP(), []int{1}
}

func (x *AutomateInfraOrgUsers) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *AutomateInfraOrgUsers) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

type ResetInfraServerUserKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef Server User's Id
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" toml:"user_id,omitempty" mapstructure:"user_id,omitempty"`
	// Chef Server Username
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty" toml:"user_name,omitempty" mapstructure:"user_name,omitempty"`
	// Chef Server Id
	ServerId string `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
}

func (x *ResetInfraServerUserKeyReq) Reset() {
	*x = ResetInfraServerUserKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetInfraServerUserKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetInfraServerUserKeyReq) ProtoMessage() {}

func (x *ResetInfraServerUserKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetInfraServerUserKeyReq.ProtoReflect.Descriptor instead.
func (*ResetInfraServerUserKeyReq) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescGZIP(), []int{2}
}

func (x *ResetInfraServerUserKeyReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ResetInfraServerUserKeyReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ResetInfraServerUserKeyReq) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

var File_interservice_infra_proxy_request_automate_infra_server_users_proto protoreflect.FileDescriptor

var file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDesc = []byte{
	0x0a, 0x42, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37,
	0x0a, 0x18, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x4f, 0x72, 0x67, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x1a, 0x52, 0x65, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescOnce sync.Once
	file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescData = file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDesc
)

func file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescGZIP() []byte {
	file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescOnce.Do(func() {
		file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescData)
	})
	return file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDescData
}

var file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interservice_infra_proxy_request_automate_infra_server_users_proto_goTypes = []interface{}{
	(*AutomateInfraServerUsers)(nil),   // 0: chef.automate.domain.infra_proxy.request.AutomateInfraServerUsers
	(*AutomateInfraOrgUsers)(nil),      // 1: chef.automate.domain.infra_proxy.request.AutomateInfraOrgUsers
	(*ResetInfraServerUserKeyReq)(nil), // 2: chef.automate.domain.infra_proxy.request.ResetInfraServerUserKeyReq
}
var file_interservice_infra_proxy_request_automate_infra_server_users_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_infra_proxy_request_automate_infra_server_users_proto_init() }
func file_interservice_infra_proxy_request_automate_infra_server_users_proto_init() {
	if File_interservice_infra_proxy_request_automate_infra_server_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutomateInfraServerUsers); i {
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
		file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutomateInfraOrgUsers); i {
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
		file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetInfraServerUserKeyReq); i {
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
			RawDescriptor: file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_infra_proxy_request_automate_infra_server_users_proto_goTypes,
		DependencyIndexes: file_interservice_infra_proxy_request_automate_infra_server_users_proto_depIdxs,
		MessageInfos:      file_interservice_infra_proxy_request_automate_infra_server_users_proto_msgTypes,
	}.Build()
	File_interservice_infra_proxy_request_automate_infra_server_users_proto = out.File
	file_interservice_infra_proxy_request_automate_infra_server_users_proto_rawDesc = nil
	file_interservice_infra_proxy_request_automate_infra_server_users_proto_goTypes = nil
	file_interservice_infra_proxy_request_automate_infra_server_users_proto_depIdxs = nil
}
