// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// data common to all messages - Top level msg format
type CommonMessageData struct {
	SessionId            []byte   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonMessageData) Reset()         { *m = CommonMessageData{} }
func (m *CommonMessageData) String() string { return proto.CompactTextString(m) }
func (*CommonMessageData) ProtoMessage()    {}
func (*CommonMessageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_6cf0f80e208218d6, []int{0}
}
func (m *CommonMessageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonMessageData.Unmarshal(m, b)
}
func (m *CommonMessageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonMessageData.Marshal(b, m, deterministic)
}
func (dst *CommonMessageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonMessageData.Merge(dst, src)
}
func (m *CommonMessageData) XXX_Size() int {
	return xxx_messageInfo_CommonMessageData.Size(m)
}
func (m *CommonMessageData) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonMessageData.DiscardUnknown(m)
}

var xxx_messageInfo_CommonMessageData proto.InternalMessageInfo

func (m *CommonMessageData) GetSessionId() []byte {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *CommonMessageData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CommonMessageData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// Handshake protocol data used for both request and response - sent unencrypted over the wire
type HandshakeData struct {
	SessionId            []byte   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ClientVersion        string   `protobuf:"bytes,4,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	NetworkID            int32    `protobuf:"varint,5,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Protocol             string   `protobuf:"bytes,6,opt,name=protocol,proto3" json:"protocol,omitempty"`
	NodePubKey           []byte   `protobuf:"bytes,7,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`
	Iv                   []byte   `protobuf:"bytes,8,opt,name=iv,proto3" json:"iv,omitempty"`
	PubKey               []byte   `protobuf:"bytes,9,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Hmac                 []byte   `protobuf:"bytes,10,opt,name=hmac,proto3" json:"hmac,omitempty"`
	Sign                 string   `protobuf:"bytes,11,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeData) Reset()         { *m = HandshakeData{} }
func (m *HandshakeData) String() string { return proto.CompactTextString(m) }
func (*HandshakeData) ProtoMessage()    {}
func (*HandshakeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_6cf0f80e208218d6, []int{1}
}
func (m *HandshakeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeData.Unmarshal(m, b)
}
func (m *HandshakeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeData.Marshal(b, m, deterministic)
}
func (dst *HandshakeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeData.Merge(dst, src)
}
func (m *HandshakeData) XXX_Size() int {
	return xxx_messageInfo_HandshakeData.Size(m)
}
func (m *HandshakeData) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeData.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeData proto.InternalMessageInfo

func (m *HandshakeData) GetSessionId() []byte {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *HandshakeData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HandshakeData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *HandshakeData) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *HandshakeData) GetNetworkID() int32 {
	if m != nil {
		return m.NetworkID
	}
	return 0
}

func (m *HandshakeData) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *HandshakeData) GetNodePubKey() []byte {
	if m != nil {
		return m.NodePubKey
	}
	return nil
}

func (m *HandshakeData) GetIv() []byte {
	if m != nil {
		return m.Iv
	}
	return nil
}

func (m *HandshakeData) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *HandshakeData) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

func (m *HandshakeData) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

// used for protocol messages (non-handshake) - this is the decrypted CommonMessageData.payload
// it allows multi310.445plexing back to higher level protocols
// data is here and not in CommonMessageData to avoid leaked data on unencrypted connections
type ProtocolMessage struct {
	Metadata             *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Payload              []byte    `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ProtocolMessage) Reset()         { *m = ProtocolMessage{} }
func (m *ProtocolMessage) String() string { return proto.CompactTextString(m) }
func (*ProtocolMessage) ProtoMessage()    {}
func (*ProtocolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_6cf0f80e208218d6, []int{2}
}
func (m *ProtocolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolMessage.Unmarshal(m, b)
}
func (m *ProtocolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolMessage.Marshal(b, m, deterministic)
}
func (dst *ProtocolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolMessage.Merge(dst, src)
}
func (m *ProtocolMessage) XXX_Size() int {
	return xxx_messageInfo_ProtocolMessage.Size(m)
}
func (m *ProtocolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolMessage proto.InternalMessageInfo

func (m *ProtocolMessage) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ProtocolMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Metadata struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ClientVersion        string   `protobuf:"bytes,2,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Gossip               bool     `protobuf:"varint,4,opt,name=gossip,proto3" json:"gossip,omitempty"`
	AuthPubKey           []byte   `protobuf:"bytes,5,opt,name=authPubKey,proto3" json:"authPubKey,omitempty"`
	AuthorSign           string   `protobuf:"bytes,6,opt,name=authorSign,proto3" json:"authorSign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_6cf0f80e208218d6, []int{3}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Metadata) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *Metadata) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Metadata) GetGossip() bool {
	if m != nil {
		return m.Gossip
	}
	return false
}

func (m *Metadata) GetAuthPubKey() []byte {
	if m != nil {
		return m.AuthPubKey
	}
	return nil
}

func (m *Metadata) GetAuthorSign() string {
	if m != nil {
		return m.AuthorSign
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonMessageData)(nil), "pb.CommonMessageData")
	proto.RegisterType((*HandshakeData)(nil), "pb.HandshakeData")
	proto.RegisterType((*ProtocolMessage)(nil), "pb.ProtocolMessage")
	proto.RegisterType((*Metadata)(nil), "pb.Metadata")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_6cf0f80e208218d6) }

var fileDescriptor_message_6cf0f80e208218d6 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xcd, 0x6e, 0xe2, 0x30,
	0x10, 0x80, 0x65, 0x03, 0x21, 0x0c, 0xb0, 0xab, 0xf5, 0x01, 0x59, 0xab, 0xd5, 0x2a, 0x8a, 0x7a,
	0xc8, 0x89, 0x43, 0xfb, 0x06, 0x2d, 0x87, 0xa2, 0x0a, 0x09, 0xa5, 0x6a, 0x0f, 0xbd, 0x39, 0xc4,
	0x02, 0x0b, 0xfc, 0xa3, 0xd8, 0x50, 0xf1, 0x60, 0x7d, 0x81, 0x3e, 0x59, 0x15, 0x27, 0x10, 0x68,
	0x45, 0x6f, 0xbd, 0x79, 0xbe, 0xf1, 0x78, 0x3c, 0xdf, 0xc0, 0x50, 0x72, 0x6b, 0xd9, 0x92, 0x8f,
	0x4d, 0xa1, 0x9d, 0x26, 0xd8, 0x64, 0xb1, 0x80, 0x3f, 0x77, 0x5a, 0x4a, 0xad, 0x66, 0x55, 0x6a,
	0xc2, 0x1c, 0x23, 0xff, 0xa0, 0x67, 0xb9, 0xb5, 0x42, 0xab, 0x69, 0x4e, 0x51, 0x84, 0x92, 0x41,
	0xda, 0x00, 0x42, 0xa1, 0x6b, 0xd8, 0x7e, 0xa3, 0x59, 0x4e, 0xb1, 0xcf, 0x1d, 0xc2, 0xb2, 0xce,
	0x09, 0xc9, 0xad, 0x63, 0xd2, 0xd0, 0x56, 0x84, 0x92, 0x56, 0xda, 0x80, 0xf8, 0x0d, 0xc3, 0xf0,
	0x9e, 0xa9, 0xdc, 0xae, 0xd8, 0xfa, 0x07, 0xfb, 0x90, 0x2b, 0x18, 0x2e, 0x36, 0x82, 0x2b, 0xf7,
	0xcc, 0x8b, 0xf2, 0x29, 0xda, 0x8e, 0x50, 0xd2, 0x4b, 0xcf, 0x61, 0xf9, 0x86, 0xe2, 0xee, 0x55,
	0x17, 0xeb, 0xe9, 0x84, 0x76, 0x22, 0x94, 0x74, 0xd2, 0x06, 0x90, 0xbf, 0x10, 0x7a, 0x47, 0x0b,
	0xbd, 0xa1, 0x81, 0x2f, 0x3f, 0xc6, 0xe4, 0x3f, 0x80, 0xd2, 0x39, 0x9f, 0x6f, 0xb3, 0x07, 0xbe,
	0xa7, 0x5d, 0xff, 0xb5, 0x13, 0x42, 0x7e, 0x01, 0x16, 0x3b, 0x1a, 0x7a, 0x8e, 0xc5, 0x8e, 0x8c,
	0x20, 0x30, 0xd5, 0xdd, 0x9e, 0x67, 0x75, 0x44, 0x08, 0xb4, 0x57, 0x92, 0x2d, 0x28, 0x78, 0xea,
	0xcf, 0x25, 0xb3, 0x62, 0xa9, 0x68, 0xdf, 0xf7, 0xf4, 0xe7, 0xf8, 0x09, 0x7e, 0xcf, 0xeb, 0xde,
	0xf5, 0x92, 0x48, 0x02, 0xa1, 0xe4, 0x8e, 0xe5, 0xcc, 0x31, 0xef, 0xad, 0x7f, 0x3d, 0x18, 0x9b,
	0x6c, 0x3c, 0xab, 0x59, 0x7a, 0xcc, 0x5e, 0x96, 0x18, 0xbf, 0x23, 0x08, 0x0f, 0x05, 0x67, 0xf3,
	0xa2, 0x4f, 0xf3, 0x7e, 0xf1, 0x89, 0x2f, 0xf8, 0xfc, 0x66, 0x27, 0x23, 0x08, 0x96, 0xda, 0x5a,
	0x61, 0xfc, 0x32, 0xc2, 0xb4, 0x8e, 0x4a, 0x97, 0x6c, 0xeb, 0x56, 0xb5, 0xcb, 0x4e, 0xe5, 0xb2,
	0x21, 0x87, 0xbc, 0x2e, 0x1e, 0x4b, 0x2b, 0xd5, 0x26, 0x4e, 0xc8, 0x6d, 0xfb, 0x05, 0x9b, 0x2c,
	0x0b, 0xfc, 0x5f, 0x6f, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xef, 0xf4, 0xb9, 0x2a, 0xe0, 0x02,
	0x00, 0x00,
}
