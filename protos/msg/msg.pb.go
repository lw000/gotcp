// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package Tmsg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//请求注册接口
type ReqRegister struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRegister) Reset()         { *m = ReqRegister{} }
func (m *ReqRegister) String() string { return proto.CompactTextString(m) }
func (*ReqRegister) ProtoMessage()    {}
func (*ReqRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *ReqRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRegister.Unmarshal(m, b)
}
func (m *ReqRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRegister.Marshal(b, m, deterministic)
}
func (m *ReqRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRegister.Merge(m, src)
}
func (m *ReqRegister) XXX_Size() int {
	return xxx_messageInfo_ReqRegister.Size(m)
}
func (m *ReqRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRegister.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRegister proto.InternalMessageInfo

func (m *ReqRegister) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *ReqRegister) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

//回应注册接口
type AckRegister struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckRegister) Reset()         { *m = AckRegister{} }
func (m *AckRegister) String() string { return proto.CompactTextString(m) }
func (*AckRegister) ProtoMessage()    {}
func (*AckRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *AckRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckRegister.Unmarshal(m, b)
}
func (m *AckRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckRegister.Marshal(b, m, deterministic)
}
func (m *AckRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckRegister.Merge(m, src)
}
func (m *AckRegister) XXX_Size() int {
	return xxx_messageInfo_AckRegister.Size(m)
}
func (m *AckRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_AckRegister.DiscardUnknown(m)
}

var xxx_messageInfo_AckRegister proto.InternalMessageInfo

func (m *AckRegister) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AckRegister) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AckRegister) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//请求登录接口
type ReqLogin struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqLogin) Reset()         { *m = ReqLogin{} }
func (m *ReqLogin) String() string { return proto.CompactTextString(m) }
func (*ReqLogin) ProtoMessage()    {}
func (*ReqLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *ReqLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqLogin.Unmarshal(m, b)
}
func (m *ReqLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqLogin.Marshal(b, m, deterministic)
}
func (m *ReqLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqLogin.Merge(m, src)
}
func (m *ReqLogin) XXX_Size() int {
	return xxx_messageInfo_ReqLogin.Size(m)
}
func (m *ReqLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqLogin.DiscardUnknown(m)
}

var xxx_messageInfo_ReqLogin proto.InternalMessageInfo

func (m *ReqLogin) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *ReqLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

//回应登录接口
type AckLogin struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Uid                  uint32   `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckLogin) Reset()         { *m = AckLogin{} }
func (m *AckLogin) String() string { return proto.CompactTextString(m) }
func (*AckLogin) ProtoMessage()    {}
func (*AckLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *AckLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckLogin.Unmarshal(m, b)
}
func (m *AckLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckLogin.Marshal(b, m, deterministic)
}
func (m *AckLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckLogin.Merge(m, src)
}
func (m *AckLogin) XXX_Size() int {
	return xxx_messageInfo_AckLogin.Size(m)
}
func (m *AckLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_AckLogin.DiscardUnknown(m)
}

var xxx_messageInfo_AckLogin proto.InternalMessageInfo

func (m *AckLogin) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AckLogin) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AckLogin) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AckLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//错误消息
type KickMsg struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KickMsg) Reset()         { *m = KickMsg{} }
func (m *KickMsg) String() string { return proto.CompactTextString(m) }
func (*KickMsg) ProtoMessage()    {}
func (*KickMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *KickMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickMsg.Unmarshal(m, b)
}
func (m *KickMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickMsg.Marshal(b, m, deterministic)
}
func (m *KickMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickMsg.Merge(m, src)
}
func (m *KickMsg) XXX_Size() int {
	return xxx_messageInfo_KickMsg.Size(m)
}
func (m *KickMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_KickMsg.DiscardUnknown(m)
}

var xxx_messageInfo_KickMsg proto.InternalMessageInfo

func (m *KickMsg) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *KickMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//错误消息
type ErrorMsg struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorMsg) Reset()         { *m = ErrorMsg{} }
func (m *ErrorMsg) String() string { return proto.CompactTextString(m) }
func (*ErrorMsg) ProtoMessage()    {}
func (*ErrorMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}

func (m *ErrorMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorMsg.Unmarshal(m, b)
}
func (m *ErrorMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorMsg.Marshal(b, m, deterministic)
}
func (m *ErrorMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMsg.Merge(m, src)
}
func (m *ErrorMsg) XXX_Size() int {
	return xxx_messageInfo_ErrorMsg.Size(m)
}
func (m *ErrorMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMsg proto.InternalMessageInfo

func (m *ErrorMsg) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ErrorMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//请求进入游戏接口
type ReqEnterGame struct {
	RoomId               uint32   `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	GameId               uint32   `protobuf:"varint,2,opt,name=gameId,proto3" json:"gameId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqEnterGame) Reset()         { *m = ReqEnterGame{} }
func (m *ReqEnterGame) String() string { return proto.CompactTextString(m) }
func (*ReqEnterGame) ProtoMessage()    {}
func (*ReqEnterGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{6}
}

func (m *ReqEnterGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEnterGame.Unmarshal(m, b)
}
func (m *ReqEnterGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEnterGame.Marshal(b, m, deterministic)
}
func (m *ReqEnterGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEnterGame.Merge(m, src)
}
func (m *ReqEnterGame) XXX_Size() int {
	return xxx_messageInfo_ReqEnterGame.Size(m)
}
func (m *ReqEnterGame) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEnterGame.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEnterGame proto.InternalMessageInfo

func (m *ReqEnterGame) GetRoomId() uint32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *ReqEnterGame) GetGameId() uint32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

//回应进入游戏接口
type AckEnterGame struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckEnterGame) Reset()         { *m = AckEnterGame{} }
func (m *AckEnterGame) String() string { return proto.CompactTextString(m) }
func (*AckEnterGame) ProtoMessage()    {}
func (*AckEnterGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{7}
}

func (m *AckEnterGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckEnterGame.Unmarshal(m, b)
}
func (m *AckEnterGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckEnterGame.Marshal(b, m, deterministic)
}
func (m *AckEnterGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckEnterGame.Merge(m, src)
}
func (m *AckEnterGame) XXX_Size() int {
	return xxx_messageInfo_AckEnterGame.Size(m)
}
func (m *AckEnterGame) XXX_DiscardUnknown() {
	xxx_messageInfo_AckEnterGame.DiscardUnknown(m)
}

var xxx_messageInfo_AckEnterGame proto.InternalMessageInfo

func (m *AckEnterGame) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AckEnterGame) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqRegister)(nil), "Tmsg.ReqRegister")
	proto.RegisterType((*AckRegister)(nil), "Tmsg.AckRegister")
	proto.RegisterType((*ReqLogin)(nil), "Tmsg.ReqLogin")
	proto.RegisterType((*AckLogin)(nil), "Tmsg.AckLogin")
	proto.RegisterType((*KickMsg)(nil), "Tmsg.KickMsg")
	proto.RegisterType((*ErrorMsg)(nil), "Tmsg.ErrorMsg")
	proto.RegisterType((*ReqEnterGame)(nil), "Tmsg.ReqEnterGame")
	proto.RegisterType((*AckEnterGame)(nil), "Tmsg.AckEnterGame")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x99, 0xad, 0x5b, 0xf7, 0xb6, 0x81, 0x04, 0x91, 0xe2, 0x49, 0x72, 0xf2, 0xa4, 0x82,
	0x9e, 0xc5, 0x22, 0x43, 0x8a, 0x7a, 0x09, 0x1e, 0xbc, 0xd6, 0xf4, 0x11, 0x4a, 0x48, 0xdf, 0x9a,
	0x64, 0xf8, 0xef, 0x4b, 0xd3, 0x74, 0xf3, 0xa0, 0x32, 0xbc, 0x7d, 0xdf, 0x0b, 0xdf, 0x97, 0xdf,
	0x83, 0x07, 0x73, 0xe3, 0xd4, 0xd5, 0xc6, 0x92, 0x27, 0x96, 0xbe, 0x19, 0xa7, 0xf8, 0x23, 0x2c,
	0x04, 0x76, 0x02, 0x55, 0xe3, 0x3c, 0x5a, 0x96, 0xc3, 0xac, 0x92, 0x92, 0xb6, 0xad, 0xcf, 0x27,
	0x17, 0x93, 0xcb, 0xb9, 0x18, 0x2d, 0x3b, 0x87, 0x6c, 0x53, 0x39, 0xf7, 0x49, 0xb6, 0xce, 0x8f,
	0xc2, 0xd3, 0xce, 0xf3, 0x12, 0x16, 0x85, 0xd4, 0xbb, 0x12, 0x06, 0xa9, 0xa4, 0x1a, 0x43, 0xc3,
	0x4a, 0x04, 0xcd, 0x4e, 0x20, 0x31, 0x4e, 0xc5, 0x64, 0x2f, 0xd9, 0x29, 0x1c, 0x7b, 0xd2, 0xd8,
	0xe6, 0x49, 0x98, 0x0d, 0x86, 0x3f, 0x40, 0x26, 0xb0, 0x7b, 0x21, 0xd5, 0xb4, 0xff, 0x84, 0x79,
	0x87, 0xac, 0x90, 0x7a, 0x68, 0xf8, 0x85, 0x64, 0xdb, 0x0c, 0xb1, 0x95, 0xe8, 0xe5, 0xc8, 0x96,
	0xfc, 0xc0, 0x96, 0x7e, 0x67, 0xbb, 0x86, 0xd9, 0x73, 0x23, 0xf5, 0xab, 0x53, 0x7f, 0xad, 0xb8,
	0xaf, 0xe1, 0x37, 0x90, 0xad, 0xad, 0x25, 0x7b, 0x78, 0xe2, 0x1e, 0x96, 0x02, 0xbb, 0x75, 0xeb,
	0xd1, 0x3e, 0x55, 0x06, 0xd9, 0x19, 0x4c, 0x2d, 0x91, 0x29, 0xeb, 0x98, 0x8b, 0xae, 0x9f, 0xab,
	0xca, 0x60, 0x39, 0xee, 0x11, 0x1d, 0xbf, 0x83, 0x65, 0x21, 0xf5, 0x3e, 0x7f, 0xd0, 0xaf, 0x1f,
	0xd3, 0x70, 0x11, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x8b, 0x6b, 0x2c, 0x1e, 0x02,
	0x00, 0x00,
}