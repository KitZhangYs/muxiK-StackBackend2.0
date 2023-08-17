// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: usercenter.proto

package pb

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentID string `protobuf:"bytes,1,opt,name=StudentID,proto3" json:"StudentID,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetStudentID() string {
	if x != nil {
		return x.StudentID
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	IfNew bool   `protobuf:"varint,2,opt,name=IfNew,proto3" json:"IfNew,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetIfNew() bool {
	if x != nil {
		return x.IfNew
	}
	return false
}

type GenerateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentID string `protobuf:"bytes,1,opt,name=StudentID,proto3" json:"StudentID,omitempty"`
	RoleGrade int64  `protobuf:"varint,2,opt,name=RoleGrade,proto3" json:"RoleGrade,omitempty"`
	Licence   int64  `protobuf:"varint,3,opt,name=Licence,proto3" json:"Licence,omitempty"`
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateTokenRequest) GetStudentID() string {
	if x != nil {
		return x.StudentID
	}
	return ""
}

func (x *GenerateTokenRequest) GetRoleGrade() int64 {
	if x != nil {
		return x.RoleGrade
	}
	return 0
}

func (x *GenerateTokenRequest) GetLicence() int64 {
	if x != nil {
		return x.Licence
	}
	return 0
}

type GenerateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *GenerateTokenResponse) Reset() {
	*x = GenerateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResponse) ProtoMessage() {}

func (x *GenerateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SID      string `protobuf:"bytes,1,opt,name=SID,proto3" json:"SID,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
}

func (x *UpdateInfoRequest) Reset() {
	*x = UpdateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfoRequest) ProtoMessage() {}

func (x *UpdateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateInfoRequest) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateInfoRequest) GetSID() string {
	if x != nil {
		return x.SID
	}
	return ""
}

func (x *UpdateInfoRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UpdateInfoRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type UpdateInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateInfoResponse) Reset() {
	*x = UpdateInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfoResponse) ProtoMessage() {}

func (x *UpdateInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateInfoResponse) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{5}
}

type SharingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SID string `protobuf:"bytes,1,opt,name=SID,proto3" json:"SID,omitempty"`
}

func (x *SharingRequest) Reset() {
	*x = SharingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharingRequest) ProtoMessage() {}

func (x *SharingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharingRequest.ProtoReflect.Descriptor instead.
func (*SharingRequest) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{6}
}

func (x *SharingRequest) GetSID() string {
	if x != nil {
		return x.SID
	}
	return ""
}

type SharingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SharingResponse) Reset() {
	*x = SharingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharingResponse) ProtoMessage() {}

func (x *SharingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharingResponse.ProtoReflect.Descriptor instead.
func (*SharingResponse) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{7}
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SID string `protobuf:"bytes,1,opt,name=SID,proto3" json:"SID,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{8}
}

func (x *GetInfoRequest) GetSID() string {
	if x != nil {
		return x.SID
	}
	return ""
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentID string `protobuf:"bytes,1,opt,name=StudentID,proto3" json:"StudentID,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Avatar    string `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	IsBlocked int64  `protobuf:"varint,4,opt,name=IsBlocked,proto3" json:"IsBlocked,omitempty"`
	RoleGrade int64  `protobuf:"varint,5,opt,name=RoleGrade,proto3" json:"RoleGrade,omitempty"`
	Integral  int64  `protobuf:"varint,6,opt,name=Integral,proto3" json:"Integral,omitempty"`
	Licence   int64  `protobuf:"varint,7,opt,name=Licence,proto3" json:"Licence,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{9}
}

func (x *GetInfoResponse) GetStudentID() string {
	if x != nil {
		return x.StudentID
	}
	return ""
}

func (x *GetInfoResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetInfoResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GetInfoResponse) GetIsBlocked() int64 {
	if x != nil {
		return x.IsBlocked
	}
	return 0
}

func (x *GetInfoResponse) GetRoleGrade() int64 {
	if x != nil {
		return x.RoleGrade
	}
	return 0
}

func (x *GetInfoResponse) GetIntegral() int64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *GetInfoResponse) GetLicence() int64 {
	if x != nil {
		return x.Licence
	}
	return 0
}

type IncreaseIntegralRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=StudentId,proto3" json:"StudentId,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Integral  int64  `protobuf:"varint,3,opt,name=Integral,proto3" json:"Integral,omitempty"`
}

func (x *IncreaseIntegralRequest) Reset() {
	*x = IncreaseIntegralRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseIntegralRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseIntegralRequest) ProtoMessage() {}

func (x *IncreaseIntegralRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseIntegralRequest.ProtoReflect.Descriptor instead.
func (*IncreaseIntegralRequest) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{10}
}

func (x *IncreaseIntegralRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *IncreaseIntegralRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IncreaseIntegralRequest) GetIntegral() int64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

type IncreaseIntegralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=StudentId,proto3" json:"StudentId,omitempty"`
	Integral  int64  `protobuf:"varint,2,opt,name=Integral,proto3" json:"Integral,omitempty"`
	RoleGrade int64  `protobuf:"varint,3,opt,name=RoleGrade,proto3" json:"RoleGrade,omitempty"`
}

func (x *IncreaseIntegralResponse) Reset() {
	*x = IncreaseIntegralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usercenter_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseIntegralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseIntegralResponse) ProtoMessage() {}

func (x *IncreaseIntegralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usercenter_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseIntegralResponse.ProtoReflect.Descriptor instead.
func (*IncreaseIntegralResponse) Descriptor() ([]byte, []int) {
	return file_usercenter_proto_rawDescGZIP(), []int{11}
}

func (x *IncreaseIntegralResponse) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *IncreaseIntegralResponse) GetIntegral() int64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *IncreaseIntegralResponse) GetRoleGrade() int64 {
	if x != nil {
		return x.RoleGrade
	}
	return 0
}

var File_usercenter_proto protoreflect.FileDescriptor

var file_usercenter_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x48, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x3b, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x66, 0x4e, 0x65, 0x77,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x49, 0x66, 0x4e, 0x65, 0x77, 0x22, 0x6c, 0x0a,
	0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2d, 0x0a, 0x15, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x59, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x53, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x53,
	0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x53, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x49, 0x44, 0x22,
	0x11, 0x0a, 0x0f, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x53, 0x49, 0x44, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x6f,
	0x6c, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x52,
	0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x67,
	0x0a, 0x17, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x22, 0x72, 0x0a, 0x18, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x32, 0xf8, 0x02, 0x0a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x73,
	0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x69, 0x6e, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x1b, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usercenter_proto_rawDescOnce sync.Once
	file_usercenter_proto_rawDescData = file_usercenter_proto_rawDesc
)

func file_usercenter_proto_rawDescGZIP() []byte {
	file_usercenter_proto_rawDescOnce.Do(func() {
		file_usercenter_proto_rawDescData = protoimpl.X.CompressGZIP(file_usercenter_proto_rawDescData)
	})
	return file_usercenter_proto_rawDescData
}

var file_usercenter_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_usercenter_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),             // 0: pb.LoginRequest
	(*LoginResponse)(nil),            // 1: pb.LoginResponse
	(*GenerateTokenRequest)(nil),     // 2: pb.GenerateTokenRequest
	(*GenerateTokenResponse)(nil),    // 3: pb.GenerateTokenResponse
	(*UpdateInfoRequest)(nil),        // 4: pb.UpdateInfoRequest
	(*UpdateInfoResponse)(nil),       // 5: pb.UpdateInfoResponse
	(*SharingRequest)(nil),           // 6: pb.SharingRequest
	(*SharingResponse)(nil),          // 7: pb.SharingResponse
	(*GetInfoRequest)(nil),           // 8: pb.GetInfoRequest
	(*GetInfoResponse)(nil),          // 9: pb.GetInfoResponse
	(*IncreaseIntegralRequest)(nil),  // 10: pb.IncreaseIntegralRequest
	(*IncreaseIntegralResponse)(nil), // 11: pb.IncreaseIntegralResponse
}
var file_usercenter_proto_depIdxs = []int32{
	0,  // 0: pb.usercenter.login:input_type -> pb.LoginRequest
	2,  // 1: pb.usercenter.generateToken:input_type -> pb.GenerateTokenRequest
	4,  // 2: pb.usercenter.updateInfo:input_type -> pb.UpdateInfoRequest
	6,  // 3: pb.usercenter.sharingPlan:input_type -> pb.SharingRequest
	8,  // 4: pb.usercenter.getInfo:input_type -> pb.GetInfoRequest
	10, // 5: pb.usercenter.increaseIntegral:input_type -> pb.IncreaseIntegralRequest
	1,  // 6: pb.usercenter.login:output_type -> pb.LoginResponse
	3,  // 7: pb.usercenter.generateToken:output_type -> pb.GenerateTokenResponse
	5,  // 8: pb.usercenter.updateInfo:output_type -> pb.UpdateInfoResponse
	7,  // 9: pb.usercenter.sharingPlan:output_type -> pb.SharingResponse
	9,  // 10: pb.usercenter.getInfo:output_type -> pb.GetInfoResponse
	11, // 11: pb.usercenter.increaseIntegral:output_type -> pb.IncreaseIntegralResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_usercenter_proto_init() }
func file_usercenter_proto_init() {
	if File_usercenter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usercenter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_usercenter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_usercenter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenRequest); i {
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
		file_usercenter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenResponse); i {
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
		file_usercenter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfoRequest); i {
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
		file_usercenter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfoResponse); i {
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
		file_usercenter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharingRequest); i {
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
		file_usercenter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharingResponse); i {
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
		file_usercenter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
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
		file_usercenter_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
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
		file_usercenter_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseIntegralRequest); i {
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
		file_usercenter_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseIntegralResponse); i {
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
			RawDescriptor: file_usercenter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usercenter_proto_goTypes,
		DependencyIndexes: file_usercenter_proto_depIdxs,
		MessageInfos:      file_usercenter_proto_msgTypes,
	}.Build()
	File_usercenter_proto = out.File
	file_usercenter_proto_rawDesc = nil
	file_usercenter_proto_goTypes = nil
	file_usercenter_proto_depIdxs = nil
}
