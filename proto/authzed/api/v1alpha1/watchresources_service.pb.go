// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: authzed/api/v1alpha1/watchresources_service.proto

package v1alpha1

import (
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// todo: work this into the v1 core API at some point since it's used
// across services.
type PermissionUpdate_Permissionship int32

const (
	PermissionUpdate_PERMISSIONSHIP_UNSPECIFIED    PermissionUpdate_Permissionship = 0
	PermissionUpdate_PERMISSIONSHIP_NO_PERMISSION  PermissionUpdate_Permissionship = 1
	PermissionUpdate_PERMISSIONSHIP_HAS_PERMISSION PermissionUpdate_Permissionship = 2
)

// Enum value maps for PermissionUpdate_Permissionship.
var (
	PermissionUpdate_Permissionship_name = map[int32]string{
		0: "PERMISSIONSHIP_UNSPECIFIED",
		1: "PERMISSIONSHIP_NO_PERMISSION",
		2: "PERMISSIONSHIP_HAS_PERMISSION",
	}
	PermissionUpdate_Permissionship_value = map[string]int32{
		"PERMISSIONSHIP_UNSPECIFIED":    0,
		"PERMISSIONSHIP_NO_PERMISSION":  1,
		"PERMISSIONSHIP_HAS_PERMISSION": 2,
	}
)

func (x PermissionUpdate_Permissionship) Enum() *PermissionUpdate_Permissionship {
	p := new(PermissionUpdate_Permissionship)
	*p = x
	return p
}

func (x PermissionUpdate_Permissionship) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionUpdate_Permissionship) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1alpha1_watchresources_service_proto_enumTypes[0].Descriptor()
}

func (PermissionUpdate_Permissionship) Type() protoreflect.EnumType {
	return &file_authzed_api_v1alpha1_watchresources_service_proto_enumTypes[0]
}

func (x PermissionUpdate_Permissionship) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionUpdate_Permissionship.Descriptor instead.
func (PermissionUpdate_Permissionship) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1alpha1_watchresources_service_proto_rawDescGZIP(), []int{1, 0}
}

// WatchResourcesRequest starts a watch for specific permission updates
// for the given resource and subject types.
type WatchResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource_object_type is the type of resource object for which we will
	// watch for changes.
	ResourceObjectType string `protobuf:"bytes,1,opt,name=resource_object_type,json=resourceObjectType,proto3" json:"resource_object_type,omitempty"`
	// permission is the name of the permission or relation for which we will
	// watch for changes.
	Permission string `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	// subject_object_type is the type of the subject resource for which we will
	// watch for changes.
	SubjectObjectType string `protobuf:"bytes,3,opt,name=subject_object_type,json=subjectObjectType,proto3" json:"subject_object_type,omitempty"`
	// optional_subject_relation allows you to specify a group of subjects to watch
	// for a given subject type.
	OptionalSubjectRelation string       `protobuf:"bytes,4,opt,name=optional_subject_relation,json=optionalSubjectRelation,proto3" json:"optional_subject_relation,omitempty"`
	OptionalStartCursor     *v1.ZedToken `protobuf:"bytes,5,opt,name=optional_start_cursor,json=optionalStartCursor,proto3" json:"optional_start_cursor,omitempty"`
}

func (x *WatchResourcesRequest) Reset() {
	*x = WatchResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResourcesRequest) ProtoMessage() {}

func (x *WatchResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResourcesRequest.ProtoReflect.Descriptor instead.
func (*WatchResourcesRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1alpha1_watchresources_service_proto_rawDescGZIP(), []int{0}
}

func (x *WatchResourcesRequest) GetResourceObjectType() string {
	if x != nil {
		return x.ResourceObjectType
	}
	return ""
}

func (x *WatchResourcesRequest) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *WatchResourcesRequest) GetSubjectObjectType() string {
	if x != nil {
		return x.SubjectObjectType
	}
	return ""
}

func (x *WatchResourcesRequest) GetOptionalSubjectRelation() string {
	if x != nil {
		return x.OptionalSubjectRelation
	}
	return ""
}

func (x *WatchResourcesRequest) GetOptionalStartCursor() *v1.ZedToken {
	if x != nil {
		return x.OptionalStartCursor
	}
	return nil
}

// PermissionUpdate represents a single permission update for a specific
// subject's permissions.
type PermissionUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// subject defines the subject resource whose permissions have changed.
	Subject *v1.SubjectReference `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// resource defines the specific object in the system.
	Resource          *v1.ObjectReference             `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Relation          string                          `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	UpdatedPermission PermissionUpdate_Permissionship `protobuf:"varint,4,opt,name=updated_permission,json=updatedPermission,proto3,enum=authzed.api.v1alpha1.PermissionUpdate_Permissionship" json:"updated_permission,omitempty"`
}

func (x *PermissionUpdate) Reset() {
	*x = PermissionUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUpdate) ProtoMessage() {}

func (x *PermissionUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUpdate.ProtoReflect.Descriptor instead.
func (*PermissionUpdate) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1alpha1_watchresources_service_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionUpdate) GetSubject() *v1.SubjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *PermissionUpdate) GetResource() *v1.ObjectReference {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *PermissionUpdate) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *PermissionUpdate) GetUpdatedPermission() PermissionUpdate_Permissionship {
	if x != nil {
		return x.UpdatedPermission
	}
	return PermissionUpdate_PERMISSIONSHIP_UNSPECIFIED
}

// WatchResourcesResponse enumerates the list of permission updates that have
// occurred as a result of one or more relationship updates.
type WatchResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updates        []*PermissionUpdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
	ChangesThrough *v1.ZedToken        `protobuf:"bytes,2,opt,name=changes_through,json=changesThrough,proto3" json:"changes_through,omitempty"`
}

func (x *WatchResourcesResponse) Reset() {
	*x = WatchResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResourcesResponse) ProtoMessage() {}

func (x *WatchResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResourcesResponse.ProtoReflect.Descriptor instead.
func (*WatchResourcesResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1alpha1_watchresources_service_proto_rawDescGZIP(), []int{2}
}

func (x *WatchResourcesResponse) GetUpdates() []*PermissionUpdate {
	if x != nil {
		return x.Updates
	}
	return nil
}

func (x *WatchResourcesResponse) GetChangesThrough() *v1.ZedToken {
	if x != nil {
		return x.ChangesThrough
	}
	return nil
}

var File_authzed_api_v1alpha1_watchresources_service_proto protoreflect.FileDescriptor

var file_authzed_api_v1alpha1_watchresources_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x15,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x7a, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x48, 0xfa, 0x42, 0x45, 0x72, 0x43, 0x28, 0x80, 0x01, 0x32, 0x3e, 0x5e,
	0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b,
	0x31, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2f, 0x29, 0x2a,
	0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31,
	0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x12, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xfa, 0x42, 0x24, 0x72, 0x22, 0x28, 0x40, 0x32, 0x1e,
	0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b,
	0x31, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x5a, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x13, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x22, 0x84, 0x03, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x64,
	0x0a, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x52, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x4e, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x9d, 0x01, 0x0a, 0x16,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x5a, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0e, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x32, 0xa9, 0x01, 0x0a, 0x15,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x0e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x30, 0x01, 0x42, 0x54, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2d,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authzed_api_v1alpha1_watchresources_service_proto_rawDescOnce sync.Once
	file_authzed_api_v1alpha1_watchresources_service_proto_rawDescData = file_authzed_api_v1alpha1_watchresources_service_proto_rawDesc
)

func file_authzed_api_v1alpha1_watchresources_service_proto_rawDescGZIP() []byte {
	file_authzed_api_v1alpha1_watchresources_service_proto_rawDescOnce.Do(func() {
		file_authzed_api_v1alpha1_watchresources_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v1alpha1_watchresources_service_proto_rawDescData)
	})
	return file_authzed_api_v1alpha1_watchresources_service_proto_rawDescData
}

var file_authzed_api_v1alpha1_watchresources_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_authzed_api_v1alpha1_watchresources_service_proto_goTypes = []interface{}{
	(PermissionUpdate_Permissionship)(0), // 0: authzed.api.v1alpha1.PermissionUpdate.Permissionship
	(*WatchResourcesRequest)(nil),        // 1: authzed.api.v1alpha1.WatchResourcesRequest
	(*PermissionUpdate)(nil),             // 2: authzed.api.v1alpha1.PermissionUpdate
	(*WatchResourcesResponse)(nil),       // 3: authzed.api.v1alpha1.WatchResourcesResponse
	(*v1.ZedToken)(nil),                  // 4: authzed.api.v1.ZedToken
	(*v1.SubjectReference)(nil),          // 5: authzed.api.v1.SubjectReference
	(*v1.ObjectReference)(nil),           // 6: authzed.api.v1.ObjectReference
}
var file_authzed_api_v1alpha1_watchresources_service_proto_depIdxs = []int32{
	4, // 0: authzed.api.v1alpha1.WatchResourcesRequest.optional_start_cursor:type_name -> authzed.api.v1.ZedToken
	5, // 1: authzed.api.v1alpha1.PermissionUpdate.subject:type_name -> authzed.api.v1.SubjectReference
	6, // 2: authzed.api.v1alpha1.PermissionUpdate.resource:type_name -> authzed.api.v1.ObjectReference
	0, // 3: authzed.api.v1alpha1.PermissionUpdate.updated_permission:type_name -> authzed.api.v1alpha1.PermissionUpdate.Permissionship
	2, // 4: authzed.api.v1alpha1.WatchResourcesResponse.updates:type_name -> authzed.api.v1alpha1.PermissionUpdate
	4, // 5: authzed.api.v1alpha1.WatchResourcesResponse.changes_through:type_name -> authzed.api.v1.ZedToken
	1, // 6: authzed.api.v1alpha1.WatchResourcesService.WatchResources:input_type -> authzed.api.v1alpha1.WatchResourcesRequest
	3, // 7: authzed.api.v1alpha1.WatchResourcesService.WatchResources:output_type -> authzed.api.v1alpha1.WatchResourcesResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_authzed_api_v1alpha1_watchresources_service_proto_init() }
func file_authzed_api_v1alpha1_watchresources_service_proto_init() {
	if File_authzed_api_v1alpha1_watchresources_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResourcesRequest); i {
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
		file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUpdate); i {
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
		file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResourcesResponse); i {
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
			RawDescriptor: file_authzed_api_v1alpha1_watchresources_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authzed_api_v1alpha1_watchresources_service_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1alpha1_watchresources_service_proto_depIdxs,
		EnumInfos:         file_authzed_api_v1alpha1_watchresources_service_proto_enumTypes,
		MessageInfos:      file_authzed_api_v1alpha1_watchresources_service_proto_msgTypes,
	}.Build()
	File_authzed_api_v1alpha1_watchresources_service_proto = out.File
	file_authzed_api_v1alpha1_watchresources_service_proto_rawDesc = nil
	file_authzed_api_v1alpha1_watchresources_service_proto_goTypes = nil
	file_authzed_api_v1alpha1_watchresources_service_proto_depIdxs = nil
}
