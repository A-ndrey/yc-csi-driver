// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/lockbox/v1/payload_service.proto

package lockbox

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetPayloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the secret.
	SecretId string `protobuf:"bytes,1,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	// Optional ID of the version.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
}

func (x *GetPayloadRequest) Reset() {
	*x = GetPayloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPayloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPayloadRequest) ProtoMessage() {}

func (x *GetPayloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPayloadRequest.ProtoReflect.Descriptor instead.
func (*GetPayloadRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPayloadRequest) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *GetPayloadRequest) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

type GetExRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifier:
	//
	//	*GetExRequest_SecretId
	//	*GetExRequest_FolderAndName
	Identifier isGetExRequest_Identifier `protobuf_oneof:"identifier"`
	VersionId  string                    `protobuf:"bytes,21,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
}

func (x *GetExRequest) Reset() {
	*x = GetExRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExRequest) ProtoMessage() {}

func (x *GetExRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExRequest.ProtoReflect.Descriptor instead.
func (*GetExRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescGZIP(), []int{1}
}

func (m *GetExRequest) GetIdentifier() isGetExRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *GetExRequest) GetSecretId() string {
	if x, ok := x.GetIdentifier().(*GetExRequest_SecretId); ok {
		return x.SecretId
	}
	return ""
}

func (x *GetExRequest) GetFolderAndName() *FolderAndName {
	if x, ok := x.GetIdentifier().(*GetExRequest_FolderAndName); ok {
		return x.FolderAndName
	}
	return nil
}

func (x *GetExRequest) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

type isGetExRequest_Identifier interface {
	isGetExRequest_Identifier()
}

type GetExRequest_SecretId struct {
	SecretId string `protobuf:"bytes,1,opt,name=secret_id,json=secretId,proto3,oneof"`
}

type GetExRequest_FolderAndName struct {
	FolderAndName *FolderAndName `protobuf:"bytes,2,opt,name=folder_and_name,json=folderAndName,proto3,oneof"`
}

func (*GetExRequest_SecretId) isGetExRequest_Identifier() {}

func (*GetExRequest_FolderAndName) isGetExRequest_Identifier() {}

type FolderAndName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FolderId   string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	SecretName string `protobuf:"bytes,2,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
}

func (x *FolderAndName) Reset() {
	*x = FolderAndName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderAndName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderAndName) ProtoMessage() {}

func (x *FolderAndName) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderAndName.ProtoReflect.Descriptor instead.
func (*FolderAndName) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescGZIP(), []int{2}
}

func (x *FolderAndName) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *FolderAndName) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

type GetExResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretId  string            `protobuf:"bytes,1,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	VersionId string            `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	Entries   map[string][]byte `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetExResponse) Reset() {
	*x = GetExResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExResponse) ProtoMessage() {}

func (x *GetExResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExResponse.ProtoReflect.Descriptor instead.
func (*GetExResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetExResponse) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *GetExResponse) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *GetExResponse) GetEntries() map[string][]byte {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_yandex_cloud_lockbox_v1_payload_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_lockbox_v1_payload_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x35, 0x30, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0a,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xcc, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c,
	0x3d, 0x35, 0x30, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x50, 0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x48, 0x00, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52,
	0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x12, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x4a, 0x04,
	0x08, 0x03, 0x10, 0x15, 0x22, 0x6a, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x41, 0x6e,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8,
	0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x05, 0x3c,
	0x3d, 0x31, 0x30, 0x30, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xd6, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4d,
	0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x3a, 0x0a,
	0x0c, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x92, 0x02, 0x0a, 0x0e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x6c, 0x6f, 0x63,
	0x6b, 0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f,
	0x7b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x79, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x45, 0x78, 0x12, 0x25, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x63, 0x6b,
	0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x45, 0x78, 0x42, 0x62,
	0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x63, 0x6b, 0x62,
	0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescData = file_yandex_cloud_lockbox_v1_payload_service_proto_rawDesc
)

func file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescData)
	})
	return file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescData
}

var file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_lockbox_v1_payload_service_proto_goTypes = []any{
	(*GetPayloadRequest)(nil), // 0: yandex.cloud.lockbox.v1.GetPayloadRequest
	(*GetExRequest)(nil),      // 1: yandex.cloud.lockbox.v1.GetExRequest
	(*FolderAndName)(nil),     // 2: yandex.cloud.lockbox.v1.FolderAndName
	(*GetExResponse)(nil),     // 3: yandex.cloud.lockbox.v1.GetExResponse
	nil,                       // 4: yandex.cloud.lockbox.v1.GetExResponse.EntriesEntry
	(*Payload)(nil),           // 5: yandex.cloud.lockbox.v1.Payload
}
var file_yandex_cloud_lockbox_v1_payload_service_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.lockbox.v1.GetExRequest.folder_and_name:type_name -> yandex.cloud.lockbox.v1.FolderAndName
	4, // 1: yandex.cloud.lockbox.v1.GetExResponse.entries:type_name -> yandex.cloud.lockbox.v1.GetExResponse.EntriesEntry
	0, // 2: yandex.cloud.lockbox.v1.PayloadService.Get:input_type -> yandex.cloud.lockbox.v1.GetPayloadRequest
	1, // 3: yandex.cloud.lockbox.v1.PayloadService.GetEx:input_type -> yandex.cloud.lockbox.v1.GetExRequest
	5, // 4: yandex.cloud.lockbox.v1.PayloadService.Get:output_type -> yandex.cloud.lockbox.v1.Payload
	3, // 5: yandex.cloud.lockbox.v1.PayloadService.GetEx:output_type -> yandex.cloud.lockbox.v1.GetExResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_lockbox_v1_payload_service_proto_init() }
func file_yandex_cloud_lockbox_v1_payload_service_proto_init() {
	if File_yandex_cloud_lockbox_v1_payload_service_proto != nil {
		return
	}
	file_yandex_cloud_lockbox_v1_payload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetPayloadRequest); i {
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
		file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetExRequest); i {
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
		file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FolderAndName); i {
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
		file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetExResponse); i {
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
	file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[1].OneofWrappers = []any{
		(*GetExRequest_SecretId)(nil),
		(*GetExRequest_FolderAndName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_lockbox_v1_payload_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_lockbox_v1_payload_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_lockbox_v1_payload_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_lockbox_v1_payload_service_proto = out.File
	file_yandex_cloud_lockbox_v1_payload_service_proto_rawDesc = nil
	file_yandex_cloud_lockbox_v1_payload_service_proto_goTypes = nil
	file_yandex_cloud_lockbox_v1_payload_service_proto_depIdxs = nil
}
