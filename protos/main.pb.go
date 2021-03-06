// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.1
// source: protos/main.proto

package protos

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

// BidInput indicates if getBids should get new bids
type BidInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refresh bool `protobuf:"varint,1,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *BidInput) Reset() {
	*x = BidInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidInput) ProtoMessage() {}

func (x *BidInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidInput.ProtoReflect.Descriptor instead.
func (*BidInput) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{0}
}

func (x *BidInput) GetRefresh() bool {
	if x != nil {
		return x.Refresh
	}
	return false
}

// BidOutput returns number of bids
type BidOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bids int32 `protobuf:"varint,1,opt,name=bids,proto3" json:"bids,omitempty"`
}

func (x *BidOutput) Reset() {
	*x = BidOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidOutput) ProtoMessage() {}

func (x *BidOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidOutput.ProtoReflect.Descriptor instead.
func (*BidOutput) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{1}
}

func (x *BidOutput) GetBids() int32 {
	if x != nil {
		return x.Bids
	}
	return 0
}

// Represent an Empty Type for Project
type ProjectEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectEmpty) Reset() {
	*x = ProjectEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectEmpty) ProtoMessage() {}

func (x *ProjectEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectEmpty.ProtoReflect.Descriptor instead.
func (*ProjectEmpty) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{2}
}

type ProjectStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ProjectStatus) Reset() {
	*x = ProjectStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectStatus) ProtoMessage() {}

func (x *ProjectStatus) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectStatus.ProtoReflect.Descriptor instead.
func (*ProjectStatus) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectStatus) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

// Represent an Empty Type for Project
type AuthEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthEmpty) Reset() {
	*x = AuthEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthEmpty) ProtoMessage() {}

func (x *AuthEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthEmpty.ProtoReflect.Descriptor instead.
func (*AuthEmpty) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{4}
}

type AuthStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCookie bool `protobuf:"varint,1,opt,name=isCookie,proto3" json:"isCookie,omitempty"`
}

func (x *AuthStatus) Reset() {
	*x = AuthStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthStatus) ProtoMessage() {}

func (x *AuthStatus) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthStatus.ProtoReflect.Descriptor instead.
func (*AuthStatus) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{5}
}

func (x *AuthStatus) GetIsCookie() bool {
	if x != nil {
		return x.IsCookie
	}
	return false
}

type AuthCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	UsernameInput string `protobuf:"bytes,3,opt,name=usernameInput,proto3" json:"usernameInput,omitempty"`
	PasswordInput string `protobuf:"bytes,4,opt,name=passwordInput,proto3" json:"passwordInput,omitempty"`
	ButtonInput   string `protobuf:"bytes,5,opt,name=buttonInput,proto3" json:"buttonInput,omitempty"`
}

func (x *AuthCredentials) Reset() {
	*x = AuthCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCredentials) ProtoMessage() {}

func (x *AuthCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthCredentials.ProtoReflect.Descriptor instead.
func (*AuthCredentials) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{6}
}

func (x *AuthCredentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthCredentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthCredentials) GetUsernameInput() string {
	if x != nil {
		return x.UsernameInput
	}
	return ""
}

func (x *AuthCredentials) GetPasswordInput() string {
	if x != nil {
		return x.PasswordInput
	}
	return ""
}

func (x *AuthCredentials) GetButtonInput() string {
	if x != nil {
		return x.ButtonInput
	}
	return ""
}

type ProjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link                   string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	BidAmount              string `protobuf:"bytes,2,opt,name=bidAmount,proto3" json:"bidAmount,omitempty"`
	ProjectDelivery        string `protobuf:"bytes,3,opt,name=projectDelivery,proto3" json:"projectDelivery,omitempty"`
	ProjectDescription     string `protobuf:"bytes,4,opt,name=projectDescription,proto3" json:"projectDescription,omitempty"`
	BidAmountPath          string `protobuf:"bytes,5,opt,name=bidAmountPath,proto3" json:"bidAmountPath,omitempty"`
	ProjectDeliveryPath    string `protobuf:"bytes,6,opt,name=projectDeliveryPath,proto3" json:"projectDeliveryPath,omitempty"`
	ProjectDescriptionPath string `protobuf:"bytes,7,opt,name=projectDescriptionPath,proto3" json:"projectDescriptionPath,omitempty"`
	BidButtonPath          string `protobuf:"bytes,8,opt,name=bidButtonPath,proto3" json:"bidButtonPath,omitempty"`
}

func (x *ProjectInfo) Reset() {
	*x = ProjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInfo) ProtoMessage() {}

func (x *ProjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectInfo.ProtoReflect.Descriptor instead.
func (*ProjectInfo) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{7}
}

func (x *ProjectInfo) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *ProjectInfo) GetBidAmount() string {
	if x != nil {
		return x.BidAmount
	}
	return ""
}

func (x *ProjectInfo) GetProjectDelivery() string {
	if x != nil {
		return x.ProjectDelivery
	}
	return ""
}

func (x *ProjectInfo) GetProjectDescription() string {
	if x != nil {
		return x.ProjectDescription
	}
	return ""
}

func (x *ProjectInfo) GetBidAmountPath() string {
	if x != nil {
		return x.BidAmountPath
	}
	return ""
}

func (x *ProjectInfo) GetProjectDeliveryPath() string {
	if x != nil {
		return x.ProjectDeliveryPath
	}
	return ""
}

func (x *ProjectInfo) GetProjectDescriptionPath() string {
	if x != nil {
		return x.ProjectDescriptionPath
	}
	return ""
}

func (x *ProjectInfo) GetBidButtonPath() string {
	if x != nil {
		return x.BidButtonPath
	}
	return ""
}

type ProjectPathInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidAmountPath          string `protobuf:"bytes,1,opt,name=bidAmountPath,proto3" json:"bidAmountPath,omitempty"`
	ProjectDeliveryPath    string `protobuf:"bytes,2,opt,name=projectDeliveryPath,proto3" json:"projectDeliveryPath,omitempty"`
	ProjectDescriptionPath string `protobuf:"bytes,3,opt,name=projectDescriptionPath,proto3" json:"projectDescriptionPath,omitempty"`
	BidButtonPath          string `protobuf:"bytes,4,opt,name=bidButtonPath,proto3" json:"bidButtonPath,omitempty"`
}

func (x *ProjectPathInfo) Reset() {
	*x = ProjectPathInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectPathInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectPathInfo) ProtoMessage() {}

func (x *ProjectPathInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectPathInfo.ProtoReflect.Descriptor instead.
func (*ProjectPathInfo) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{8}
}

func (x *ProjectPathInfo) GetBidAmountPath() string {
	if x != nil {
		return x.BidAmountPath
	}
	return ""
}

func (x *ProjectPathInfo) GetProjectDeliveryPath() string {
	if x != nil {
		return x.ProjectDeliveryPath
	}
	return ""
}

func (x *ProjectPathInfo) GetProjectDescriptionPath() string {
	if x != nil {
		return x.ProjectDescriptionPath
	}
	return ""
}

func (x *ProjectPathInfo) GetBidButtonPath() string {
	if x != nil {
		return x.BidButtonPath
	}
	return ""
}

type Projects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skills       string `protobuf:"bytes,1,opt,name=skills,proto3" json:"skills,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Link         string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	BiddingPrice string `protobuf:"bytes,5,opt,name=biddingPrice,proto3" json:"biddingPrice,omitempty"`
	Selected     bool   `protobuf:"varint,6,opt,name=selected,proto3" json:"selected,omitempty"`
	Currency     string `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *Projects) Reset() {
	*x = Projects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_main_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Projects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Projects) ProtoMessage() {}

func (x *Projects) ProtoReflect() protoreflect.Message {
	mi := &file_protos_main_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Projects.ProtoReflect.Descriptor instead.
func (*Projects) Descriptor() ([]byte, []int) {
	return file_protos_main_proto_rawDescGZIP(), []int{9}
}

func (x *Projects) GetSkills() string {
	if x != nil {
		return x.Skills
	}
	return ""
}

func (x *Projects) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Projects) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Projects) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Projects) GetBiddingPrice() string {
	if x != nil {
		return x.BiddingPrice
	}
	return ""
}

func (x *Projects) GetSelected() bool {
	if x != nil {
		return x.Selected
	}
	return false
}

func (x *Projects) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

var File_protos_main_proto protoreflect.FileDescriptor

var file_protos_main_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x5f, 0x52, 0x50, 0x43, 0x22, 0x24, 0x0a, 0x08, 0x42, 0x69, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x22, 0x1f, 0x0a, 0x09, 0x42, 0x69,
	0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x0b, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x28, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x73, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0f,
	0x41, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xcf, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x69, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x36, 0x0a, 0x16, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x69, 0x64, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x69, 0x64, 0x42, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x22, 0xc7, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x62,
	0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x62, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x36, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x62,
	0x69, 0x64, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x62, 0x69, 0x64, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x50, 0x61, 0x74,
	0x68, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x32, 0x49,
	0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x69, 0x64, 0x73,
	0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x52,
	0x50, 0x43, 0x2e, 0x42, 0x69, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x50, 0x43, 0x2e, 0x42, 0x69,
	0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x32, 0xfd, 0x01, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4e, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64,
	0x65, 0x72, 0x5f, 0x52, 0x50, 0x43, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x50, 0x43, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0c, 0x62, 0x69, 0x64, 0x4f, 0x6e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x5f, 0x52, 0x50, 0x43, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x5f, 0x52, 0x50, 0x43, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x50, 0x43, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x50, 0x43, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x00, 0x30, 0x01, 0x32, 0xa1, 0x01, 0x0a, 0x04, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x48, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f,
	0x52, 0x50, 0x43, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x50, 0x43, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x50, 0x43, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x1b,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x50, 0x43,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x18, 0x5a,
	0x16, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_main_proto_rawDescOnce sync.Once
	file_protos_main_proto_rawDescData = file_protos_main_proto_rawDesc
)

func file_protos_main_proto_rawDescGZIP() []byte {
	file_protos_main_proto_rawDescOnce.Do(func() {
		file_protos_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_main_proto_rawDescData)
	})
	return file_protos_main_proto_rawDescData
}

var file_protos_main_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_main_proto_goTypes = []interface{}{
	(*BidInput)(nil),        // 0: auto_bidder_RPC.BidInput
	(*BidOutput)(nil),       // 1: auto_bidder_RPC.BidOutput
	(*ProjectEmpty)(nil),    // 2: auto_bidder_RPC.ProjectEmpty
	(*ProjectStatus)(nil),   // 3: auto_bidder_RPC.ProjectStatus
	(*AuthEmpty)(nil),       // 4: auto_bidder_RPC.AuthEmpty
	(*AuthStatus)(nil),      // 5: auto_bidder_RPC.AuthStatus
	(*AuthCredentials)(nil), // 6: auto_bidder_RPC.AuthCredentials
	(*ProjectInfo)(nil),     // 7: auto_bidder_RPC.ProjectInfo
	(*ProjectPathInfo)(nil), // 8: auto_bidder_RPC.ProjectPathInfo
	(*Projects)(nil),        // 9: auto_bidder_RPC.Projects
}
var file_protos_main_proto_depIdxs = []int32{
	0, // 0: auto_bidder_RPC.Bid.getBids:input_type -> auto_bidder_RPC.BidInput
	8, // 1: auto_bidder_RPC.Project.getProjects:input_type -> auto_bidder_RPC.ProjectPathInfo
	7, // 2: auto_bidder_RPC.Project.bidOnProject:input_type -> auto_bidder_RPC.ProjectInfo
	2, // 3: auto_bidder_RPC.Project.subscribeToProject:input_type -> auto_bidder_RPC.ProjectEmpty
	4, // 4: auto_bidder_RPC.Auth.checkStatus:input_type -> auto_bidder_RPC.AuthEmpty
	6, // 5: auto_bidder_RPC.Auth.authenticate:input_type -> auto_bidder_RPC.AuthCredentials
	1, // 6: auto_bidder_RPC.Bid.getBids:output_type -> auto_bidder_RPC.BidOutput
	9, // 7: auto_bidder_RPC.Project.getProjects:output_type -> auto_bidder_RPC.Projects
	3, // 8: auto_bidder_RPC.Project.bidOnProject:output_type -> auto_bidder_RPC.ProjectStatus
	9, // 9: auto_bidder_RPC.Project.subscribeToProject:output_type -> auto_bidder_RPC.Projects
	5, // 10: auto_bidder_RPC.Auth.checkStatus:output_type -> auto_bidder_RPC.AuthStatus
	5, // 11: auto_bidder_RPC.Auth.authenticate:output_type -> auto_bidder_RPC.AuthStatus
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_main_proto_init() }
func file_protos_main_proto_init() {
	if File_protos_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidInput); i {
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
		file_protos_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidOutput); i {
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
		file_protos_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectEmpty); i {
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
		file_protos_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectStatus); i {
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
		file_protos_main_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthEmpty); i {
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
		file_protos_main_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthStatus); i {
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
		file_protos_main_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthCredentials); i {
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
		file_protos_main_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInfo); i {
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
		file_protos_main_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectPathInfo); i {
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
		file_protos_main_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Projects); i {
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
			RawDescriptor: file_protos_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_protos_main_proto_goTypes,
		DependencyIndexes: file_protos_main_proto_depIdxs,
		MessageInfos:      file_protos_main_proto_msgTypes,
	}.Build()
	File_protos_main_proto = out.File
	file_protos_main_proto_rawDesc = nil
	file_protos_main_proto_goTypes = nil
	file_protos_main_proto_depIdxs = nil
}
