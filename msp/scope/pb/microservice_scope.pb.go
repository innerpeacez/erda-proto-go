// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: microservice_scope.proto

package pb

import (
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EncryptMicroserviceScopeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID int64  `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Env       string `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *EncryptMicroserviceScopeRequest) Reset() {
	*x = EncryptMicroserviceScopeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_scope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptMicroserviceScopeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptMicroserviceScopeRequest) ProtoMessage() {}

func (x *EncryptMicroserviceScopeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_scope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptMicroserviceScopeRequest.ProtoReflect.Descriptor instead.
func (*EncryptMicroserviceScopeRequest) Descriptor() ([]byte, []int) {
	return file_microservice_scope_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptMicroserviceScopeRequest) GetProjectID() int64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *EncryptMicroserviceScopeRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

type EncryptMicroserviceScopeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope *Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *EncryptMicroserviceScopeResponse) Reset() {
	*x = EncryptMicroserviceScopeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_scope_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptMicroserviceScopeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptMicroserviceScopeResponse) ProtoMessage() {}

func (x *EncryptMicroserviceScopeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_scope_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptMicroserviceScopeResponse.ProtoReflect.Descriptor instead.
func (*EncryptMicroserviceScopeResponse) Descriptor() ([]byte, []int) {
	return file_microservice_scope_proto_rawDescGZIP(), []int{1}
}

func (x *EncryptMicroserviceScopeResponse) GetScope() *Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

type DecryptMicroserviceScopeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScopeID string `protobuf:"bytes,1,opt,name=scopeID,proto3" json:"scopeID,omitempty"`
}

func (x *DecryptMicroserviceScopeRequest) Reset() {
	*x = DecryptMicroserviceScopeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_scope_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptMicroserviceScopeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptMicroserviceScopeRequest) ProtoMessage() {}

func (x *DecryptMicroserviceScopeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_scope_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptMicroserviceScopeRequest.ProtoReflect.Descriptor instead.
func (*DecryptMicroserviceScopeRequest) Descriptor() ([]byte, []int) {
	return file_microservice_scope_proto_rawDescGZIP(), []int{2}
}

func (x *DecryptMicroserviceScopeRequest) GetScopeID() string {
	if x != nil {
		return x.ScopeID
	}
	return ""
}

type DecryptMicroserviceScopeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope *Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *DecryptMicroserviceScopeResponse) Reset() {
	*x = DecryptMicroserviceScopeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_scope_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptMicroserviceScopeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptMicroserviceScopeResponse) ProtoMessage() {}

func (x *DecryptMicroserviceScopeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_scope_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptMicroserviceScopeResponse.ProtoReflect.Descriptor instead.
func (*DecryptMicroserviceScopeResponse) Descriptor() ([]byte, []int) {
	return file_microservice_scope_proto_rawDescGZIP(), []int{3}
}

func (x *DecryptMicroserviceScopeResponse) GetScope() *Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

// Differentiate the actual deployment environment of the service and addon
type Scope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectID int64  `protobuf:"varint,2,opt,name=projectID,json=projectId,proto3" json:"projectID,omitempty"`
	Env       string `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *Scope) Reset() {
	*x = Scope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_scope_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope) ProtoMessage() {}

func (x *Scope) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_scope_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scope.ProtoReflect.Descriptor instead.
func (*Scope) Descriptor() ([]byte, []int) {
	return file_microservice_scope_proto_rawDescGZIP(), []int{4}
}

func (x *Scope) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scope) GetProjectID() int64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *Scope) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

var File_microservice_scope_proto protoreflect.FileDescriptor

var file_microservice_scope_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x61, 0x0a, 0x1f, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x03, 0x65,
	0x6e, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0x4f, 0x0a, 0x20, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x43, 0x0a, 0x1f, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x44, 0x22, 0x4f, 0x0a, 0x20, 0x44,
	0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x05,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x32, 0x82, 0x03, 0x0a, 0x18, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xb3, 0x01, 0x0a, 0x18, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x2f, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x3d, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x7d, 0x12, 0xaf, 0x01, 0x0a, 0x18, 0x44, 0x65, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x2f, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a,
	0x12, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2f, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x3f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64,
	0x3d, 0x7b, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x44, 0x7d, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microservice_scope_proto_rawDescOnce sync.Once
	file_microservice_scope_proto_rawDescData = file_microservice_scope_proto_rawDesc
)

func file_microservice_scope_proto_rawDescGZIP() []byte {
	file_microservice_scope_proto_rawDescOnce.Do(func() {
		file_microservice_scope_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservice_scope_proto_rawDescData)
	})
	return file_microservice_scope_proto_rawDescData
}

var file_microservice_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_microservice_scope_proto_goTypes = []interface{}{
	(*EncryptMicroserviceScopeRequest)(nil),  // 0: erda.msp.scope.EncryptMicroserviceScopeRequest
	(*EncryptMicroserviceScopeResponse)(nil), // 1: erda.msp.scope.EncryptMicroserviceScopeResponse
	(*DecryptMicroserviceScopeRequest)(nil),  // 2: erda.msp.scope.DecryptMicroserviceScopeRequest
	(*DecryptMicroserviceScopeResponse)(nil), // 3: erda.msp.scope.DecryptMicroserviceScopeResponse
	(*Scope)(nil),                            // 4: erda.msp.scope.Scope
}
var file_microservice_scope_proto_depIdxs = []int32{
	4, // 0: erda.msp.scope.EncryptMicroserviceScopeResponse.scope:type_name -> erda.msp.scope.Scope
	4, // 1: erda.msp.scope.DecryptMicroserviceScopeResponse.scope:type_name -> erda.msp.scope.Scope
	0, // 2: erda.msp.scope.MicroserviceScopeService.EncryptMicroserviceScope:input_type -> erda.msp.scope.EncryptMicroserviceScopeRequest
	2, // 3: erda.msp.scope.MicroserviceScopeService.DecryptMicroserviceScope:input_type -> erda.msp.scope.DecryptMicroserviceScopeRequest
	1, // 4: erda.msp.scope.MicroserviceScopeService.EncryptMicroserviceScope:output_type -> erda.msp.scope.EncryptMicroserviceScopeResponse
	3, // 5: erda.msp.scope.MicroserviceScopeService.DecryptMicroserviceScope:output_type -> erda.msp.scope.DecryptMicroserviceScopeResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_microservice_scope_proto_init() }
func file_microservice_scope_proto_init() {
	if File_microservice_scope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservice_scope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptMicroserviceScopeRequest); i {
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
		file_microservice_scope_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptMicroserviceScopeResponse); i {
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
		file_microservice_scope_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptMicroserviceScopeRequest); i {
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
		file_microservice_scope_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptMicroserviceScopeResponse); i {
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
		file_microservice_scope_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scope); i {
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
			RawDescriptor: file_microservice_scope_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_microservice_scope_proto_goTypes,
		DependencyIndexes: file_microservice_scope_proto_depIdxs,
		MessageInfos:      file_microservice_scope_proto_msgTypes,
	}.Build()
	File_microservice_scope_proto = out.File
	file_microservice_scope_proto_rawDesc = nil
	file_microservice_scope_proto_goTypes = nil
	file_microservice_scope_proto_depIdxs = nil
}
