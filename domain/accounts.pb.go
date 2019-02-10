// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts.proto

package domain

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AccountType int32

const (
	AccountType_USER  AccountType = 0
	AccountType_ADMIN AccountType = 1
	AccountType_DEV   AccountType = 2
)

var AccountType_name = map[int32]string{
	0: "USER",
	1: "ADMIN",
	2: "DEV",
}

var AccountType_value = map[string]int32{
	"USER":  0,
	"ADMIN": 1,
	"DEV":   2,
}

func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}

func (AccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}

type AccountStatus int32

const (
	AccountStatus_CREATED  AccountStatus = 0
	AccountStatus_ACTIVE   AccountStatus = 1
	AccountStatus_LOCKED   AccountStatus = 2
	AccountStatus_INACTIVE AccountStatus = 3
	AccountStatus_DELETED  AccountStatus = 4
)

var AccountStatus_name = map[int32]string{
	0: "CREATED",
	1: "ACTIVE",
	2: "LOCKED",
	3: "INACTIVE",
	4: "DELETED",
}

var AccountStatus_value = map[string]int32{
	"CREATED":  0,
	"ACTIVE":   1,
	"LOCKED":   2,
	"INACTIVE": 3,
	"DELETED":  4,
}

func (x AccountStatus) String() string {
	return proto.EnumName(AccountStatus_name, int32(x))
}

func (AccountStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}

//AccountEntity entity (timestamps in seconds)
type AccountEntity struct {
	// `datastore:"-"`
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty" datastore:"-"`
	Em  string `protobuf:"bytes,2,opt,name=em,proto3" json:"em,omitempty"`
	// `json:"-"`
	Pw                   string        `protobuf:"bytes,3,opt,name=pw,json=-,proto3" json:"-"`
	CAt                  int64         `protobuf:"varint,4,opt,name=c_at,json=cAt,proto3" json:"c_at,omitempty"`
	UAt                  int64         `protobuf:"varint,5,opt,name=u_at,json=uAt,proto3" json:"u_at,omitempty"`
	Type                 AccountType   `protobuf:"varint,6,opt,name=type,proto3,enum=domain.AccountType" json:"type,omitempty"`
	Status               AccountStatus `protobuf:"varint,7,opt,name=status,proto3,enum=domain.AccountStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AccountEntity) Reset()         { *m = AccountEntity{} }
func (m *AccountEntity) String() string { return proto.CompactTextString(m) }
func (*AccountEntity) ProtoMessage()    {}
func (*AccountEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}

func (m *AccountEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEntity.Unmarshal(m, b)
}
func (m *AccountEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEntity.Marshal(b, m, deterministic)
}
func (m *AccountEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEntity.Merge(m, src)
}
func (m *AccountEntity) XXX_Size() int {
	return xxx_messageInfo_AccountEntity.Size(m)
}
func (m *AccountEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountEntity.DiscardUnknown(m)
}

var xxx_messageInfo_AccountEntity proto.InternalMessageInfo

func (m *AccountEntity) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AccountEntity) GetEm() string {
	if m != nil {
		return m.Em
	}
	return ""
}

func (m *AccountEntity) GetPw() string {
	if m != nil {
		return m.Pw
	}
	return ""
}

func (m *AccountEntity) GetCAt() int64 {
	if m != nil {
		return m.CAt
	}
	return 0
}

func (m *AccountEntity) GetUAt() int64 {
	if m != nil {
		return m.UAt
	}
	return 0
}

func (m *AccountEntity) GetType() AccountType {
	if m != nil {
		return m.Type
	}
	return AccountType_USER
}

func (m *AccountEntity) GetStatus() AccountStatus {
	if m != nil {
		return m.Status
	}
	return AccountStatus_CREATED
}

type AccountID struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountID) Reset()         { *m = AccountID{} }
func (m *AccountID) String() string { return proto.CompactTextString(m) }
func (*AccountID) ProtoMessage()    {}
func (*AccountID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}

func (m *AccountID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountID.Unmarshal(m, b)
}
func (m *AccountID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountID.Marshal(b, m, deterministic)
}
func (m *AccountID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountID.Merge(m, src)
}
func (m *AccountID) XXX_Size() int {
	return xxx_messageInfo_AccountID.Size(m)
}
func (m *AccountID) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountID.DiscardUnknown(m)
}

var xxx_messageInfo_AccountID proto.InternalMessageInfo

func (m *AccountID) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type AccountEmail struct {
	Em                   string   `protobuf:"bytes,1,opt,name=em,proto3" json:"em,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountEmail) Reset()         { *m = AccountEmail{} }
func (m *AccountEmail) String() string { return proto.CompactTextString(m) }
func (*AccountEmail) ProtoMessage()    {}
func (*AccountEmail) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{2}
}

func (m *AccountEmail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEmail.Unmarshal(m, b)
}
func (m *AccountEmail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEmail.Marshal(b, m, deterministic)
}
func (m *AccountEmail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEmail.Merge(m, src)
}
func (m *AccountEmail) XXX_Size() int {
	return xxx_messageInfo_AccountEmail.Size(m)
}
func (m *AccountEmail) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountEmail.DiscardUnknown(m)
}

var xxx_messageInfo_AccountEmail proto.InternalMessageInfo

func (m *AccountEmail) GetEm() string {
	if m != nil {
		return m.Em
	}
	return ""
}

//AccountParams holds payload to create/update an Account
type AccountParams struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Em                   string   `protobuf:"bytes,2,opt,name=em,proto3" json:"em,omitempty"`
	Pw                   string   `protobuf:"bytes,3,opt,name=pw,proto3" json:"pw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountParams) Reset()         { *m = AccountParams{} }
func (m *AccountParams) String() string { return proto.CompactTextString(m) }
func (*AccountParams) ProtoMessage()    {}
func (*AccountParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{3}
}

func (m *AccountParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountParams.Unmarshal(m, b)
}
func (m *AccountParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountParams.Marshal(b, m, deterministic)
}
func (m *AccountParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountParams.Merge(m, src)
}
func (m *AccountParams) XXX_Size() int {
	return xxx_messageInfo_AccountParams.Size(m)
}
func (m *AccountParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountParams.DiscardUnknown(m)
}

var xxx_messageInfo_AccountParams proto.InternalMessageInfo

func (m *AccountParams) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AccountParams) GetEm() string {
	if m != nil {
		return m.Em
	}
	return ""
}

func (m *AccountParams) GetPw() string {
	if m != nil {
		return m.Pw
	}
	return ""
}

//AccountResp holds response for Account creation or update
type AccountResp struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountResp) Reset()         { *m = AccountResp{} }
func (m *AccountResp) String() string { return proto.CompactTextString(m) }
func (*AccountResp) ProtoMessage()    {}
func (*AccountResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{4}
}

func (m *AccountResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResp.Unmarshal(m, b)
}
func (m *AccountResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResp.Marshal(b, m, deterministic)
}
func (m *AccountResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResp.Merge(m, src)
}
func (m *AccountResp) XXX_Size() int {
	return xxx_messageInfo_AccountResp.Size(m)
}
func (m *AccountResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResp.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResp proto.InternalMessageInfo

func (m *AccountResp) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AccountResp) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

//AccountJwtTokens holds authentication tokens
type AccountJwtTokens struct {
	Access               string   `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh              string   `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountJwtTokens) Reset()         { *m = AccountJwtTokens{} }
func (m *AccountJwtTokens) String() string { return proto.CompactTextString(m) }
func (*AccountJwtTokens) ProtoMessage()    {}
func (*AccountJwtTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{5}
}

func (m *AccountJwtTokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountJwtTokens.Unmarshal(m, b)
}
func (m *AccountJwtTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountJwtTokens.Marshal(b, m, deterministic)
}
func (m *AccountJwtTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountJwtTokens.Merge(m, src)
}
func (m *AccountJwtTokens) XXX_Size() int {
	return xxx_messageInfo_AccountJwtTokens.Size(m)
}
func (m *AccountJwtTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountJwtTokens.DiscardUnknown(m)
}

var xxx_messageInfo_AccountJwtTokens proto.InternalMessageInfo

func (m *AccountJwtTokens) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *AccountJwtTokens) GetRefresh() string {
	if m != nil {
		return m.Refresh
	}
	return ""
}

//AccountCredentials holds credentials to authenticate a user
type AccountCredentials struct {
	Em                   string   `protobuf:"bytes,1,opt,name=em,proto3" json:"em,omitempty"`
	Pw                   string   `protobuf:"bytes,2,opt,name=pw,proto3" json:"pw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountCredentials) Reset()         { *m = AccountCredentials{} }
func (m *AccountCredentials) String() string { return proto.CompactTextString(m) }
func (*AccountCredentials) ProtoMessage()    {}
func (*AccountCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{6}
}

func (m *AccountCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCredentials.Unmarshal(m, b)
}
func (m *AccountCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCredentials.Marshal(b, m, deterministic)
}
func (m *AccountCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCredentials.Merge(m, src)
}
func (m *AccountCredentials) XXX_Size() int {
	return xxx_messageInfo_AccountCredentials.Size(m)
}
func (m *AccountCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCredentials proto.InternalMessageInfo

func (m *AccountCredentials) GetEm() string {
	if m != nil {
		return m.Em
	}
	return ""
}

func (m *AccountCredentials) GetPw() string {
	if m != nil {
		return m.Pw
	}
	return ""
}

func init() {
	proto.RegisterEnum("domain.AccountType", AccountType_name, AccountType_value)
	proto.RegisterEnum("domain.AccountStatus", AccountStatus_name, AccountStatus_value)
	proto.RegisterType((*AccountEntity)(nil), "domain.AccountEntity")
	proto.RegisterType((*AccountID)(nil), "domain.AccountID")
	proto.RegisterType((*AccountEmail)(nil), "domain.AccountEmail")
	proto.RegisterType((*AccountParams)(nil), "domain.AccountParams")
	proto.RegisterType((*AccountResp)(nil), "domain.AccountResp")
	proto.RegisterType((*AccountJwtTokens)(nil), "domain.AccountJwtTokens")
	proto.RegisterType((*AccountCredentials)(nil), "domain.AccountCredentials")
}

func init() { proto.RegisterFile("accounts.proto", fileDescriptor_e1e7723af4c007b7) }

var fileDescriptor_e1e7723af4c007b7 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xda, 0x4c,
	0x14, 0x8d, 0x6d, 0x62, 0xc8, 0x85, 0x20, 0xe7, 0x7e, 0x5f, 0x2b, 0x0b, 0xa9, 0x15, 0x75, 0x17,
	0x45, 0xa9, 0x02, 0x12, 0xc9, 0x26, 0x52, 0x37, 0x8e, 0x6d, 0x55, 0x6e, 0x53, 0x82, 0x1c, 0x9a,
	0x45, 0x37, 0xd5, 0x60, 0xa6, 0xc1, 0x02, 0xff, 0xc8, 0x33, 0x96, 0xc5, 0xcb, 0xf5, 0x65, 0xfa,
	0x20, 0xad, 0x6c, 0x0f, 0x28, 0x32, 0xa9, 0x54, 0x76, 0x73, 0x7f, 0xce, 0x9d, 0x73, 0xce, 0x1d,
	0x1b, 0xba, 0xc4, 0xf7, 0xe3, 0x2c, 0xe2, 0x6c, 0x98, 0xa4, 0x31, 0x8f, 0x51, 0x5d, 0xc4, 0x21,
	0x09, 0x22, 0xe3, 0xa7, 0x04, 0xa7, 0x66, 0x55, 0x72, 0x22, 0x1e, 0xf0, 0x0d, 0x6a, 0xa0, 0x64,
	0xc1, 0x42, 0x97, 0xfa, 0xd2, 0xe0, 0xc4, 0x2b, 0x8e, 0xd8, 0x05, 0x99, 0x86, 0xba, 0x5c, 0x26,
	0x64, 0x1a, 0xe2, 0x29, 0xc8, 0x49, 0xae, 0x2b, 0x65, 0x2c, 0x5d, 0xe0, 0x19, 0x34, 0xfc, 0xef,
	0x84, 0xeb, 0x8d, 0xbe, 0x34, 0x50, 0x3c, 0xc5, 0x37, 0x79, 0x91, 0xca, 0x8a, 0xd4, 0x71, 0x95,
	0xca, 0x4c, 0x8e, 0xef, 0xa0, 0xc1, 0x37, 0x09, 0xd5, 0xd5, 0xbe, 0x34, 0xe8, 0x8e, 0xff, 0x1b,
	0x56, 0xf7, 0x0f, 0xc5, 0xdd, 0xb3, 0x4d, 0x42, 0xbd, 0xb2, 0x01, 0x2f, 0x40, 0x65, 0x9c, 0xf0,
	0x8c, 0xe9, 0xcd, 0xb2, 0xf5, 0x45, 0xad, 0xf5, 0xbe, 0x2c, 0x7a, 0xa2, 0xc9, 0x78, 0x05, 0x27,
	0xa2, 0xe0, 0xda, 0xfb, 0xdc, 0x8d, 0xd7, 0xd0, 0xd9, 0xca, 0x0b, 0x49, 0xb0, 0x16, 0x5a, 0xa4,
	0xad, 0x16, 0xc3, 0xdc, 0xc9, 0x9f, 0x92, 0x94, 0x84, 0xec, 0x1f, 0xe4, 0x77, 0x9f, 0xc8, 0x97,
	0x93, 0xdc, 0x18, 0x41, 0x5b, 0x8c, 0xf0, 0x28, 0x4b, 0x9e, 0x1f, 0x10, 0xaf, 0xca, 0x01, 0x2d,
	0x4f, 0x8e, 0x57, 0x86, 0x0d, 0x9a, 0x00, 0x7c, 0xca, 0xf9, 0x2c, 0x5e, 0xd1, 0x88, 0xe1, 0x4b,
	0x50, 0x89, 0xef, 0x53, 0xc6, 0x04, 0x50, 0x44, 0xa8, 0x43, 0x33, 0xa5, 0x3f, 0x52, 0xca, 0x96,
	0x82, 0xc1, 0x36, 0x34, 0xae, 0x00, 0xc5, 0x14, 0x2b, 0xa5, 0x0b, 0x1a, 0xf1, 0x80, 0xac, 0x59,
	0x5d, 0x9f, 0x20, 0x2b, 0x6f, 0xc9, 0x9e, 0xbf, 0xdf, 0x91, 0x2d, 0x2c, 0xc7, 0x16, 0x34, 0xbe,
	0xde, 0x3b, 0x9e, 0x76, 0x84, 0x27, 0x70, 0x6c, 0xda, 0x5f, 0xdc, 0x89, 0x26, 0x61, 0x13, 0x14,
	0xdb, 0x79, 0xd0, 0xe4, 0xf3, 0xbb, 0x9d, 0x39, 0x95, 0xe9, 0xd8, 0x86, 0xa6, 0xe5, 0x39, 0xe6,
	0xcc, 0xb1, 0xb5, 0x23, 0x04, 0x50, 0x4d, 0x6b, 0xe6, 0x3e, 0x38, 0x9a, 0x54, 0x9c, 0x6f, 0xef,
	0xac, 0xcf, 0x8e, 0xad, 0xc9, 0xd8, 0x81, 0x96, 0x3b, 0x11, 0x15, 0xa5, 0x80, 0xd8, 0xce, 0xad,
	0x53, 0x40, 0x1a, 0xe3, 0xdf, 0x12, 0xb4, 0xc4, 0x44, 0x86, 0x57, 0xa0, 0x5a, 0x29, 0x25, 0x9c,
	0x62, 0x7d, 0xc5, 0xd5, 0x2a, 0x7a, 0xf5, 0x47, 0x52, 0xda, 0x7b, 0x0d, 0x6d, 0x6b, 0x49, 0xa2,
	0x47, 0x5a, 0xed, 0xf3, 0x10, 0xe8, 0x07, 0xe8, 0x56, 0xd0, 0x29, 0x61, 0x2c, 0x8f, 0xd3, 0xc5,
	0x41, 0x68, 0x1b, 0x3a, 0x66, 0xc6, 0x97, 0x85, 0xd1, 0x7e, 0x41, 0xba, 0x57, 0x6b, 0x7a, 0xb2,
	0x85, 0x9e, 0x5e, 0xab, 0xed, 0xf6, 0x3c, 0xfe, 0x25, 0xed, 0x1e, 0x24, 0xf3, 0x68, 0x12, 0xe3,
	0x18, 0xd4, 0x69, 0xc6, 0x27, 0x34, 0xdf, 0x23, 0x53, 0x7d, 0x8f, 0xbd, 0xb3, 0x5a, 0xda, 0xb5,
	0x71, 0x04, 0xca, 0x34, 0xe3, 0x07, 0x00, 0x2e, 0xa1, 0xf9, 0x91, 0xf2, 0x9b, 0x8d, 0x6b, 0xe3,
	0x7e, 0xb5, 0xf7, 0xfc, 0x1c, 0xbc, 0x06, 0x28, 0x41, 0x95, 0xd1, 0xff, 0xd7, 0x9b, 0x8a, 0xec,
	0x5f, 0xa0, 0x37, 0x6f, 0xbf, 0xbd, 0x79, 0x0c, 0xf8, 0x32, 0x9b, 0x0f, 0xfd, 0x38, 0x1c, 0xad,
	0xd6, 0x64, 0xc9, 0x18, 0x8d, 0x46, 0x39, 0x9d, 0x93, 0x24, 0x19, 0x55, 0x90, 0xb9, 0x5a, 0xfe,
	0x89, 0x2e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xab, 0x5d, 0x63, 0x9b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsClient interface {
	Create(ctx context.Context, in *AccountParams, opts ...grpc.CallOption) (*AccountResp, error)
	ChangeEmail(ctx context.Context, in *AccountParams, opts ...grpc.CallOption) (*AccountResp, error)
	ChangePassword(ctx context.Context, in *AccountParams, opts ...grpc.CallOption) (*AccountResp, error)
	Authenticate(ctx context.Context, in *AccountCredentials, opts ...grpc.CallOption) (*AccountJwtTokens, error)
}

type accountsClient struct {
	cc *grpc.ClientConn
}

func NewAccountsClient(cc *grpc.ClientConn) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) Create(ctx context.Context, in *AccountParams, opts ...grpc.CallOption) (*AccountResp, error) {
	out := new(AccountResp)
	err := c.cc.Invoke(ctx, "/domain.Accounts/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) ChangeEmail(ctx context.Context, in *AccountParams, opts ...grpc.CallOption) (*AccountResp, error) {
	out := new(AccountResp)
	err := c.cc.Invoke(ctx, "/domain.Accounts/ChangeEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) ChangePassword(ctx context.Context, in *AccountParams, opts ...grpc.CallOption) (*AccountResp, error) {
	out := new(AccountResp)
	err := c.cc.Invoke(ctx, "/domain.Accounts/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Authenticate(ctx context.Context, in *AccountCredentials, opts ...grpc.CallOption) (*AccountJwtTokens, error) {
	out := new(AccountJwtTokens)
	err := c.cc.Invoke(ctx, "/domain.Accounts/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
type AccountsServer interface {
	Create(context.Context, *AccountParams) (*AccountResp, error)
	ChangeEmail(context.Context, *AccountParams) (*AccountResp, error)
	ChangePassword(context.Context, *AccountParams) (*AccountResp, error)
	Authenticate(context.Context, *AccountCredentials) (*AccountJwtTokens, error)
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.Accounts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Create(ctx, req.(*AccountParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_ChangeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).ChangeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.Accounts/ChangeEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).ChangeEmail(ctx, req.(*AccountParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.Accounts/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).ChangePassword(ctx, req.(*AccountParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.Accounts/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Authenticate(ctx, req.(*AccountCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "domain.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Accounts_Create_Handler,
		},
		{
			MethodName: "ChangeEmail",
			Handler:    _Accounts_ChangeEmail_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Accounts_ChangePassword_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _Accounts_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}

// AccountsRepoClient is the client API for AccountsRepo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsRepoClient interface {
	PutNew(ctx context.Context, in *AccountEntity, opts ...grpc.CallOption) (*AccountID, error)
	Put(ctx context.Context, in *AccountEntity, opts ...grpc.CallOption) (*AccountID, error)
	GetByID(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (*AccountEntity, error)
	GetByEmail(ctx context.Context, in *AccountEmail, opts ...grpc.CallOption) (*AccountEntity, error)
}

type accountsRepoClient struct {
	cc *grpc.ClientConn
}

func NewAccountsRepoClient(cc *grpc.ClientConn) AccountsRepoClient {
	return &accountsRepoClient{cc}
}

func (c *accountsRepoClient) PutNew(ctx context.Context, in *AccountEntity, opts ...grpc.CallOption) (*AccountID, error) {
	out := new(AccountID)
	err := c.cc.Invoke(ctx, "/domain.AccountsRepo/PutNew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsRepoClient) Put(ctx context.Context, in *AccountEntity, opts ...grpc.CallOption) (*AccountID, error) {
	out := new(AccountID)
	err := c.cc.Invoke(ctx, "/domain.AccountsRepo/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsRepoClient) GetByID(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (*AccountEntity, error) {
	out := new(AccountEntity)
	err := c.cc.Invoke(ctx, "/domain.AccountsRepo/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsRepoClient) GetByEmail(ctx context.Context, in *AccountEmail, opts ...grpc.CallOption) (*AccountEntity, error) {
	out := new(AccountEntity)
	err := c.cc.Invoke(ctx, "/domain.AccountsRepo/GetByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsRepoServer is the server API for AccountsRepo service.
type AccountsRepoServer interface {
	PutNew(context.Context, *AccountEntity) (*AccountID, error)
	Put(context.Context, *AccountEntity) (*AccountID, error)
	GetByID(context.Context, *AccountID) (*AccountEntity, error)
	GetByEmail(context.Context, *AccountEmail) (*AccountEntity, error)
}

func RegisterAccountsRepoServer(s *grpc.Server, srv AccountsRepoServer) {
	s.RegisterService(&_AccountsRepo_serviceDesc, srv)
}

func _AccountsRepo_PutNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsRepoServer).PutNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.AccountsRepo/PutNew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsRepoServer).PutNew(ctx, req.(*AccountEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsRepo_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsRepoServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.AccountsRepo/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsRepoServer).Put(ctx, req.(*AccountEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsRepo_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsRepoServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.AccountsRepo/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsRepoServer).GetByID(ctx, req.(*AccountID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsRepo_GetByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsRepoServer).GetByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.AccountsRepo/GetByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsRepoServer).GetByEmail(ctx, req.(*AccountEmail))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountsRepo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "domain.AccountsRepo",
	HandlerType: (*AccountsRepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutNew",
			Handler:    _AccountsRepo_PutNew_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _AccountsRepo_Put_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _AccountsRepo_GetByID_Handler,
		},
		{
			MethodName: "GetByEmail",
			Handler:    _AccountsRepo_GetByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}