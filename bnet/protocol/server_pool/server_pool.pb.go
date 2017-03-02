// Code generated by protoc-gen-go.
// source: bnet/protocol/server_pool/server_pool.proto
// DO NOT EDIT!

/*
Package bnet_protocol_server_pool is a generated protocol buffer package.

It is generated from these files:
	bnet/protocol/server_pool/server_pool.proto

It has these top-level messages:
	GetLoadRequest
	PoolStateRequest
	PoolStateResponse
	ServerInfo
	ServerState
*/
package bnet_protocol_server_pool

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bnet_protocol "github.com/HearthSim/hs-proto-go/bnet/protocol"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetLoadRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetLoadRequest) Reset()                    { *m = GetLoadRequest{} }
func (m *GetLoadRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLoadRequest) ProtoMessage()               {}
func (*GetLoadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PoolStateRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PoolStateRequest) Reset()                    { *m = PoolStateRequest{} }
func (m *PoolStateRequest) String() string            { return proto.CompactTextString(m) }
func (*PoolStateRequest) ProtoMessage()               {}
func (*PoolStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PoolStateResponse struct {
	Info             []*ServerInfo `protobuf:"bytes,1,rep,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *PoolStateResponse) Reset()                    { *m = PoolStateResponse{} }
func (m *PoolStateResponse) String() string            { return proto.CompactTextString(m) }
func (*PoolStateResponse) ProtoMessage()               {}
func (*PoolStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PoolStateResponse) GetInfo() []*ServerInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ServerInfo struct {
	Host             *bnet_protocol.ProcessId       `protobuf:"bytes,1,req,name=host" json:"host,omitempty"`
	Replace          *bool                          `protobuf:"varint,2,opt,name=replace,def=0" json:"replace,omitempty"`
	State            *ServerState                   `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Attribute        []*bnet_protocol.AttributeA29E `protobuf:"bytes,4,rep,name=attribute" json:"attribute,omitempty"`
	ProgramId        *uint32                        `protobuf:"fixed32,5,opt,name=program_id,json=programId" json:"program_id,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_ServerInfo_Replace bool = false

func (m *ServerInfo) GetHost() *bnet_protocol.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ServerInfo) GetReplace() bool {
	if m != nil && m.Replace != nil {
		return *m.Replace
	}
	return Default_ServerInfo_Replace
}

func (m *ServerInfo) GetState() *ServerState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ServerInfo) GetAttribute() []*bnet_protocol.AttributeA29E {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *ServerInfo) GetProgramId() uint32 {
	if m != nil && m.ProgramId != nil {
		return *m.ProgramId
	}
	return 0
}

type ServerState struct {
	CurrentLoad      *float32 `protobuf:"fixed32,1,opt,name=current_load,json=currentLoad,def=1" json:"current_load,omitempty"`
	GameCount        *uint32  `protobuf:"varint,2,opt,name=game_count,json=gameCount,def=0" json:"game_count,omitempty"`
	PlayerCount      *uint32  `protobuf:"varint,3,opt,name=player_count,json=playerCount,def=0" json:"player_count,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ServerState) Reset()                    { *m = ServerState{} }
func (m *ServerState) String() string            { return proto.CompactTextString(m) }
func (*ServerState) ProtoMessage()               {}
func (*ServerState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

const Default_ServerState_CurrentLoad float32 = 1
const Default_ServerState_GameCount uint32 = 0
const Default_ServerState_PlayerCount uint32 = 0

func (m *ServerState) GetCurrentLoad() float32 {
	if m != nil && m.CurrentLoad != nil {
		return *m.CurrentLoad
	}
	return Default_ServerState_CurrentLoad
}

func (m *ServerState) GetGameCount() uint32 {
	if m != nil && m.GameCount != nil {
		return *m.GameCount
	}
	return Default_ServerState_GameCount
}

func (m *ServerState) GetPlayerCount() uint32 {
	if m != nil && m.PlayerCount != nil {
		return *m.PlayerCount
	}
	return Default_ServerState_PlayerCount
}

func init() {
	proto.RegisterType((*GetLoadRequest)(nil), "bnet.protocol.server_pool.GetLoadRequest")
	proto.RegisterType((*PoolStateRequest)(nil), "bnet.protocol.server_pool.PoolStateRequest")
	proto.RegisterType((*PoolStateResponse)(nil), "bnet.protocol.server_pool.PoolStateResponse")
	proto.RegisterType((*ServerInfo)(nil), "bnet.protocol.server_pool.ServerInfo")
	proto.RegisterType((*ServerState)(nil), "bnet.protocol.server_pool.ServerState")
}

func init() { proto.RegisterFile("bnet/protocol/server_pool/server_pool.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x86, 0x59, 0x7f, 0xe0, 0x7a, 0xd4, 0x16, 0x77, 0x4f, 0xdb, 0x52, 0x53, 0x21, 0xda, 0x22,
	0x68, 0x91, 0x13, 0xdd, 0xec, 0xe4, 0x12, 0x72, 0x08, 0x86, 0x10, 0xcc, 0xfa, 0x07, 0x88, 0xb5,
	0x34, 0x76, 0x0c, 0xb2, 0x46, 0xd9, 0x5d, 0x05, 0x42, 0x7e, 0x77, 0xee, 0x41, 0x6b, 0x2b, 0xfe,
	0x80, 0x90, 0xdb, 0xf0, 0xe8, 0x99, 0x57, 0xf3, 0xb2, 0xf0, 0x6f, 0x51, 0xa0, 0x1d, 0x95, 0x9a,
	0x2c, 0xa5, 0x94, 0x8f, 0x0c, 0xea, 0x47, 0xd4, 0x49, 0x49, 0xc7, 0x73, 0xe4, 0x04, 0xfe, 0xbd,
	0x96, 0xa3, 0x46, 0x8e, 0x0e, 0x84, 0x1f, 0x3f, 0x8f, 0x73, 0xde, 0x1c, 0x37, 0x04, 0x03, 0xf8,
	0x7a, 0x83, 0xf6, 0x96, 0x54, 0x26, 0xf1, 0xa1, 0x42, 0x63, 0x03, 0x0e, 0x83, 0x19, 0x51, 0x3e,
	0xb7, 0xca, 0x62, 0xc3, 0xee, 0xe0, 0xdb, 0x01, 0x33, 0x25, 0x15, 0x06, 0xf9, 0x18, 0x3a, 0xeb,
	0x62, 0x49, 0x82, 0xf9, 0xed, 0xd0, 0x8b, 0xff, 0x44, 0xef, 0x9e, 0x10, 0xcd, 0xdd, 0x3c, 0x2d,
	0x96, 0x24, 0xdd, 0x4a, 0xf0, 0xc2, 0x00, 0xf6, 0x90, 0xff, 0x87, 0xce, 0x3d, 0x19, 0x2b, 0x98,
	0xdf, 0x0a, 0xbd, 0x58, 0x9c, 0x24, 0xcd, 0x34, 0xa5, 0x68, 0xcc, 0x34, 0x93, 0xce, 0xe2, 0xbf,
	0xa0, 0xa7, 0xb1, 0xcc, 0x55, 0x8a, 0xa2, 0xe5, 0xb3, 0xf0, 0xd3, 0xa4, 0xbb, 0x54, 0xb9, 0x41,
	0xd9, 0x50, 0x7e, 0x09, 0x5d, 0x53, 0x5f, 0x2a, 0xda, 0x3e, 0x0b, 0xbd, 0xf8, 0xef, 0x87, 0x97,
	0x6d, 0x7b, 0x6d, 0x97, 0xf8, 0x05, 0xf4, 0x95, 0xb5, 0x7a, 0xbd, 0xa8, 0x2c, 0x8a, 0x8e, 0xeb,
	0x36, 0x3c, 0x49, 0xb8, 0x6a, 0xbe, 0x27, 0x2a, 0x1e, 0xa3, 0xdc, 0xfb, 0x7c, 0x08, 0x50, 0x6a,
	0x5a, 0x69, 0xb5, 0x49, 0xd6, 0x99, 0xe8, 0xfa, 0x2c, 0xec, 0xc9, 0xfe, 0x8e, 0x4c, 0xb3, 0xe0,
	0x19, 0xbc, 0x83, 0x3f, 0xf2, 0xdf, 0xf0, 0x39, 0xad, 0xb4, 0xc6, 0xc2, 0x26, 0x39, 0xa9, 0x4c,
	0x30, 0x9f, 0x85, 0xad, 0x09, 0x3b, 0x97, 0xde, 0x0e, 0xd7, 0xef, 0xc2, 0x7d, 0x80, 0x95, 0xda,
	0x60, 0x92, 0x52, 0x55, 0x58, 0x57, 0xf9, 0xcb, 0x84, 0x9d, 0xc9, 0x7e, 0x0d, 0xaf, 0x6b, 0x56,
	0xe7, 0x94, 0xb9, 0x7a, 0x42, 0xbd, 0x73, 0xda, 0x8d, 0xe3, 0x6d, 0xb1, 0xb3, 0x5e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xee, 0xde, 0xde, 0x5a, 0x51, 0x02, 0x00, 0x00,
}
