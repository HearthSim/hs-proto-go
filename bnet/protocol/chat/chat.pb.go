// Code generated by protoc-gen-go.
// source: bnet/protocol/chat/chat.proto
// DO NOT EDIT!

/*
Package bnet_protocol_chat is a generated protocol buffer package.

It is generated from these files:
	bnet/protocol/chat/chat.proto

It has these top-level messages:
	ChannelState
*/
package bnet_protocol_chat

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

type ChannelState struct {
	Identity         *string `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Program          *uint32 `protobuf:"fixed32,2,opt,name=program" json:"program,omitempty"`
	Locale           *uint32 `protobuf:"fixed32,3,opt,name=locale" json:"locale,omitempty"`
	Public           *bool   `protobuf:"varint,4,opt,name=public,def=0" json:"public,omitempty"`
	BucketIndex      *uint32 `protobuf:"varint,5,opt,name=bucket_index,json=bucketIndex" json:"bucket_index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChannelState) Reset()                    { *m = ChannelState{} }
func (m *ChannelState) String() string            { return proto.CompactTextString(m) }
func (*ChannelState) ProtoMessage()               {}
func (*ChannelState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ChannelState_Public bool = false

func (m *ChannelState) GetIdentity() string {
	if m != nil && m.Identity != nil {
		return *m.Identity
	}
	return ""
}

func (m *ChannelState) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *ChannelState) GetLocale() uint32 {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return 0
}

func (m *ChannelState) GetPublic() bool {
	if m != nil && m.Public != nil {
		return *m.Public
	}
	return Default_ChannelState_Public
}

func (m *ChannelState) GetBucketIndex() uint32 {
	if m != nil && m.BucketIndex != nil {
		return *m.BucketIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*ChannelState)(nil), "bnet.protocol.chat.ChannelState")
}

func init() { proto.RegisterFile("bnet/protocol/chat/chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0xc9, 0xbd, 0xae, 0x82, 0x30,
	0x18, 0xc6, 0xf1, 0xf4, 0x9c, 0xc3, 0xc7, 0xa9, 0xb8, 0x74, 0x30, 0x8d, 0x09, 0x49, 0x75, 0xea,
	0x04, 0xbb, 0xab, 0x93, 0x6b, 0xbd, 0x00, 0x53, 0xca, 0xab, 0x34, 0xd6, 0x96, 0xe0, 0x4b, 0xa2,
	0x37, 0xe3, 0xb5, 0x1a, 0x10, 0x5c, 0x9e, 0xe4, 0xf9, 0xff, 0x68, 0x5e, 0x79, 0xc0, 0xb2, 0xed,
	0x02, 0x06, 0x13, 0x5c, 0x69, 0x1a, 0x8d, 0xe3, 0x14, 0x63, 0x62, 0x6c, 0xe0, 0x62, 0xe6, 0x62,
	0x90, 0xed, 0x8b, 0xd0, 0x6c, 0xdf, 0x68, 0xef, 0xc1, 0x1d, 0x51, 0x23, 0xb0, 0x35, 0x4d, 0x6d,
	0x0d, 0x1e, 0x2d, 0x3e, 0x39, 0x11, 0x44, 0xfe, 0xab, 0xef, 0x67, 0x9c, 0x26, 0x6d, 0x17, 0x2e,
	0x9d, 0xbe, 0xf1, 0x1f, 0x41, 0x64, 0xa2, 0xe6, 0xcb, 0x56, 0x34, 0x76, 0xc1, 0x68, 0x07, 0xfc,
	0x77, 0x84, 0xe9, 0xb1, 0x9c, 0xc6, 0x6d, 0x5f, 0x39, 0x6b, 0xf8, 0x9f, 0x20, 0x32, 0xdd, 0x45,
	0x67, 0xed, 0xee, 0xa0, 0xa6, 0xc8, 0x36, 0x34, 0xab, 0x7a, 0x73, 0x05, 0x3c, 0x59, 0x5f, 0xc3,
	0x83, 0x47, 0x82, 0xc8, 0xa5, 0x5a, 0x7c, 0xda, 0x61, 0x48, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x84, 0x94, 0x0f, 0x01, 0xd4, 0x00, 0x00, 0x00,
}
