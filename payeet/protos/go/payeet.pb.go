// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: payeet.proto

package payeet

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type StatusResponseCode int32

const (
	StatusResponse_SUCCESS         StatusResponseCode = 0
	StatusResponse_WORNG_CREDS     StatusResponseCode = 1
	StatusResponse_INVALID_MAIL    StatusResponseCode = 2
	StatusResponse_INVALID_SESSION StatusResponseCode = 3
)

var StatusResponseCode_name = map[int32]string{
	0: "SUCCESS",
	1: "WORNG_CREDS",
	2: "INVALID_MAIL",
	3: "INVALID_SESSION",
}

var StatusResponseCode_value = map[string]int32{
	"SUCCESS":         0,
	"WORNG_CREDS":     1,
	"INVALID_MAIL":    2,
	"INVALID_SESSION": 3,
}

func (x StatusResponseCode) String() string {
	return proto.EnumName(StatusResponseCode_name, int32(x))
}

func (StatusResponseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{4, 0}
}

type LoginRequest_S struct {
	Session              string   `protobuf:"bytes,1,opt,name=Session,proto3" json:"Session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest_S) Reset()         { *m = LoginRequest_S{} }
func (m *LoginRequest_S) String() string { return proto.CompactTextString(m) }
func (*LoginRequest_S) ProtoMessage()    {}
func (*LoginRequest_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{0}
}
func (m *LoginRequest_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest_S.Unmarshal(m, b)
}
func (m *LoginRequest_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest_S.Marshal(b, m, deterministic)
}
func (m *LoginRequest_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest_S.Merge(m, src)
}
func (m *LoginRequest_S) XXX_Size() int {
	return xxx_messageInfo_LoginRequest_S.Size(m)
}
func (m *LoginRequest_S) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest_S.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest_S proto.InternalMessageInfo

func (m *LoginRequest_S) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type LoginRequest struct {
	Mail                 string   `protobuf:"bytes,1,opt,name=Mail,proto3" json:"Mail,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{1}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Session              string   `protobuf:"bytes,3,opt,name=Session,proto3" json:"Session,omitempty"`
	User_ID              string   `protobuf:"bytes,4,opt,name=User_ID,json=UserID,proto3" json:"User_ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{2}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *LoginResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *LoginResponse) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *LoginResponse) GetUser_ID() string {
	if m != nil {
		return m.User_ID
	}
	return ""
}

type RegisterRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Mail                 string   `protobuf:"bytes,3,opt,name=Mail,proto3" json:"Mail,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{3}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterRequest) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type StatusResponse struct {
	StatusCode           StatusResponseCode `protobuf:"varint,1,opt,name=StatusCode,proto3,enum=payeet.StatusResponseCode" json:"StatusCode,omitempty"`
	Message              string             `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{4}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetStatusCode() StatusResponseCode {
	if m != nil {
		return m.StatusCode
	}
	return StatusResponse_SUCCESS
}

func (m *StatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type BalanceRequest struct {
	User_ID              string   `protobuf:"bytes,1,opt,name=User_ID,json=UserID,proto3" json:"User_ID,omitempty"`
	Session              string   `protobuf:"bytes,2,opt,name=Session,proto3" json:"Session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceRequest) Reset()         { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()    {}
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{5}
}
func (m *BalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceRequest.Unmarshal(m, b)
}
func (m *BalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceRequest.Marshal(b, m, deterministic)
}
func (m *BalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceRequest.Merge(m, src)
}
func (m *BalanceRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceRequest.Size(m)
}
func (m *BalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceRequest proto.InternalMessageInfo

func (m *BalanceRequest) GetUser_ID() string {
	if m != nil {
		return m.User_ID
	}
	return ""
}

func (m *BalanceRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type BalanceResponse struct {
	Balance              string   `protobuf:"bytes,1,opt,name=Balance,proto3" json:"Balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceResponse) Reset()         { *m = BalanceResponse{} }
func (m *BalanceResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse) ProtoMessage()    {}
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{6}
}
func (m *BalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceResponse.Unmarshal(m, b)
}
func (m *BalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceResponse.Marshal(b, m, deterministic)
}
func (m *BalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse.Merge(m, src)
}
func (m *BalanceResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceResponse.Size(m)
}
func (m *BalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse proto.InternalMessageInfo

func (m *BalanceResponse) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

type TransferRequest struct {
	ReceiverMail         string   `protobuf:"bytes,1,opt,name=ReceiverMail,proto3" json:"ReceiverMail,omitempty"`
	SenderID             string   `protobuf:"bytes,2,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	Session              string   `protobuf:"bytes,3,opt,name=Session,proto3" json:"Session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{7}
}
func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (m *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(m, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetReceiverMail() string {
	if m != nil {
		return m.ReceiverMail
	}
	return ""
}

func (m *TransferRequest) GetSenderID() string {
	if m != nil {
		return m.SenderID
	}
	return ""
}

func (m *TransferRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type HistoryRequest struct {
	Index                int32    `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	SenderID             string   `protobuf:"bytes,2,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	Session              string   `protobuf:"bytes,3,opt,name=Session,proto3" json:"Session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistoryRequest) Reset()         { *m = HistoryRequest{} }
func (m *HistoryRequest) String() string { return proto.CompactTextString(m) }
func (*HistoryRequest) ProtoMessage()    {}
func (*HistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49889347cc5f18bf, []int{8}
}
func (m *HistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryRequest.Unmarshal(m, b)
}
func (m *HistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryRequest.Marshal(b, m, deterministic)
}
func (m *HistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryRequest.Merge(m, src)
}
func (m *HistoryRequest) XXX_Size() int {
	return xxx_messageInfo_HistoryRequest.Size(m)
}
func (m *HistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryRequest proto.InternalMessageInfo

func (m *HistoryRequest) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *HistoryRequest) GetSenderID() string {
	if m != nil {
		return m.SenderID
	}
	return ""
}

func (m *HistoryRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func init() {
	proto.RegisterEnum("payeet.StatusResponseCode", StatusResponseCode_name, StatusResponseCode_value)
	proto.RegisterType((*LoginRequest_S)(nil), "payeet.LoginRequest_S")
	proto.RegisterType((*LoginRequest)(nil), "payeet.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "payeet.LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "payeet.RegisterRequest")
	proto.RegisterType((*StatusResponse)(nil), "payeet.StatusResponse")
	proto.RegisterType((*BalanceRequest)(nil), "payeet.BalanceRequest")
	proto.RegisterType((*BalanceResponse)(nil), "payeet.BalanceResponse")
	proto.RegisterType((*TransferRequest)(nil), "payeet.TransferRequest")
	proto.RegisterType((*HistoryRequest)(nil), "payeet.historyRequest")
}

func init() { proto.RegisterFile("payeet.proto", fileDescriptor_49889347cc5f18bf) }

var fileDescriptor_49889347cc5f18bf = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0xc4, 0xf9, 0x6c, 0x5f, 0x83, 0x6d, 0x2d, 0xa5, 0xb1, 0x02, 0x07, 0xb4, 0x27, 0x04, 0x52,
	0x0e, 0x05, 0x89, 0x43, 0x05, 0x52, 0xeb, 0x84, 0xca, 0x22, 0x49, 0x91, 0x97, 0xc2, 0x05, 0xc9,
	0x5a, 0xe2, 0x47, 0xb0, 0x28, 0xde, 0xe0, 0xdd, 0x02, 0x15, 0xfc, 0x2c, 0x24, 0xfe, 0x1e, 0xb2,
	0xb3, 0x9b, 0xd8, 0x69, 0x7d, 0xe9, 0x2d, 0x33, 0x7e, 0xdf, 0x33, 0x1b, 0xe8, 0x2d, 0xf9, 0x15,
	0xa2, 0x1a, 0x2e, 0x33, 0xa1, 0x04, 0xe9, 0xac, 0x10, 0x7d, 0x02, 0xf6, 0x44, 0x2c, 0x92, 0x34,
	0xc4, 0xef, 0x97, 0x28, 0x55, 0xc4, 0x88, 0x07, 0x5d, 0x86, 0x52, 0x26, 0x22, 0xf5, 0xac, 0x47,
	0xd6, 0xe3, 0xdd, 0xd0, 0x40, 0xfa, 0x0a, 0x7a, 0xe5, 0x58, 0x42, 0xa0, 0x35, 0xe5, 0xc9, 0x85,
	0x0e, 0x2b, 0x7e, 0x93, 0x01, 0xec, 0xbc, 0xe5, 0x52, 0xfe, 0x14, 0x59, 0xec, 0x35, 0x0a, 0x7e,
	0x8d, 0xe9, 0x1f, 0xb8, 0xab, 0xf3, 0xe5, 0x52, 0xa4, 0x12, 0xc9, 0x43, 0xd8, 0x7d, 0x9d, 0x64,
	0x52, 0xcd, 0xf8, 0x37, 0xd4, 0x55, 0x36, 0x44, 0x5e, 0x6a, 0xc2, 0xf5, 0x47, 0x5d, 0xca, 0xe0,
	0xf2, 0x90, 0xcd, 0xca, 0x90, 0xa4, 0x0f, 0xdd, 0x73, 0x89, 0x59, 0x14, 0x8c, 0xbc, 0x56, 0xf1,
	0xa5, 0x93, 0xc3, 0x60, 0x44, 0x7f, 0x83, 0x13, 0xe2, 0x22, 0x91, 0x0a, 0x33, 0xb3, 0xc0, 0xed,
	0xfb, 0x9b, 0xd5, 0x9b, 0x35, 0xab, 0xb7, 0xb6, 0x56, 0xff, 0x6b, 0x81, 0xcd, 0x14, 0x57, 0x97,
	0x72, 0xbd, 0xfc, 0x11, 0xc0, 0x8a, 0xf1, 0x45, 0xbc, 0xea, 0x6e, 0x1f, 0x3e, 0x18, 0x6a, 0x91,
	0xaa, 0xb1, 0xc3, 0xb9, 0x88, 0x31, 0x2c, 0x85, 0xe7, 0xfb, 0x4f, 0x51, 0x4a, 0xbe, 0x30, 0xa3,
	0x19, 0x48, 0xdf, 0x40, 0x2b, 0x8f, 0x26, 0x7b, 0xd0, 0x65, 0xe7, 0xbe, 0x3f, 0x66, 0xcc, 0xbd,
	0x43, 0x1c, 0xd8, 0xfb, 0x70, 0x16, 0xce, 0x4e, 0x23, 0x3f, 0x1c, 0x8f, 0x98, 0x6b, 0x11, 0x17,
	0x7a, 0xc1, 0xec, 0xfd, 0xf1, 0x24, 0x18, 0x45, 0xd3, 0xe3, 0x60, 0xe2, 0x36, 0xc8, 0x3d, 0x70,
	0x0c, 0xc3, 0xc6, 0x8c, 0x05, 0x67, 0x33, 0xb7, 0x49, 0x7d, 0xb0, 0x4f, 0xf8, 0x05, 0x4f, 0xe7,
	0x68, 0x4e, 0x56, 0x3a, 0xaf, 0x55, 0x3e, 0x6f, 0x59, 0x91, 0x46, 0xd5, 0x36, 0x4f, 0xc1, 0x59,
	0x17, 0xd1, 0xbb, 0x7b, 0xd0, 0xd5, 0x94, 0xf1, 0x98, 0x86, 0xf4, 0x2b, 0x38, 0xef, 0x32, 0x9e,
	0xca, 0xcf, 0x1b, 0x95, 0x28, 0xf4, 0x42, 0x9c, 0x63, 0xf2, 0x03, 0xb3, 0x92, 0xdd, 0x2a, 0x5c,
	0x7e, 0x7b, 0x86, 0x69, 0x9c, 0x4f, 0x62, 0xb4, 0x32, 0xb8, 0xde, 0x2b, 0xf4, 0x23, 0xd8, 0x5f,
	0x12, 0xa9, 0x44, 0x76, 0x65, 0x7a, 0xed, 0x43, 0x3b, 0x48, 0x63, 0xfc, 0x55, 0x34, 0x69, 0x87,
	0x2b, 0x70, 0xbb, 0xea, 0x87, 0xff, 0x1a, 0xa0, 0x5f, 0x19, 0x79, 0x0e, 0xed, 0xc2, 0xf9, 0x64,
	0xdf, 0x08, 0x5c, 0x7e, 0x48, 0x83, 0xfb, 0x5b, 0xac, 0xbe, 0xd2, 0x0b, 0xe8, 0x14, 0x04, 0x23,
	0x07, 0x37, 0xa5, 0x45, 0xac, 0x2e, 0xf1, 0x08, 0x76, 0x8c, 0xd5, 0x49, 0xdf, 0x84, 0x6c, 0x99,
	0x7f, 0x70, 0x70, 0xb3, 0xd7, 0xc8, 0x4b, 0x80, 0x53, 0x54, 0x5a, 0x8f, 0x4d, 0xe7, 0xaa, 0x0f,
	0x06, 0xfd, 0x6b, 0xbc, 0x4e, 0x3f, 0xd9, 0x08, 0x68, 0x6a, 0xac, 0x63, 0xb7, 0x94, 0xad, 0x1b,
	0xe1, 0x53, 0xa7, 0xf8, 0x8f, 0x7a, 0xf6, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x95, 0x82, 0x35,
	0xb3, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PayeetClient is the client API for Payeet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PayeetClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	LoginS(ctx context.Context, in *LoginRequest_S, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	TransferBalance(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type payeetClient struct {
	cc *grpc.ClientConn
}

func NewPayeetClient(cc *grpc.ClientConn) PayeetClient {
	return &payeetClient{cc}
}

func (c *payeetClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/payeet.payeet/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payeetClient) LoginS(ctx context.Context, in *LoginRequest_S, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/payeet.payeet/LoginS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payeetClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/payeet.payeet/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payeetClient) GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/payeet.payeet/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payeetClient) TransferBalance(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/payeet.payeet/TransferBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayeetServer is the server API for Payeet service.
type PayeetServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	LoginS(context.Context, *LoginRequest_S) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*StatusResponse, error)
	GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	TransferBalance(context.Context, *TransferRequest) (*StatusResponse, error)
}

// UnimplementedPayeetServer can be embedded to have forward compatible implementations.
type UnimplementedPayeetServer struct {
}

func (*UnimplementedPayeetServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedPayeetServer) LoginS(ctx context.Context, req *LoginRequest_S) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginS not implemented")
}
func (*UnimplementedPayeetServer) Register(ctx context.Context, req *RegisterRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedPayeetServer) GetBalance(ctx context.Context, req *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (*UnimplementedPayeetServer) TransferBalance(ctx context.Context, req *TransferRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferBalance not implemented")
}

func RegisterPayeetServer(s *grpc.Server, srv PayeetServer) {
	s.RegisterService(&_Payeet_serviceDesc, srv)
}

func _Payeet_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayeetServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payeet.payeet/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayeetServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payeet_LoginS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest_S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayeetServer).LoginS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payeet.payeet/LoginS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayeetServer).LoginS(ctx, req.(*LoginRequest_S))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payeet_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayeetServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payeet.payeet/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayeetServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payeet_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayeetServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payeet.payeet/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayeetServer).GetBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payeet_TransferBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayeetServer).TransferBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payeet.payeet/TransferBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayeetServer).TransferBalance(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Payeet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "payeet.payeet",
	HandlerType: (*PayeetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Payeet_Login_Handler,
		},
		{
			MethodName: "LoginS",
			Handler:    _Payeet_LoginS_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Payeet_Register_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Payeet_GetBalance_Handler,
		},
		{
			MethodName: "TransferBalance",
			Handler:    _Payeet_TransferBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payeet.proto",
}
