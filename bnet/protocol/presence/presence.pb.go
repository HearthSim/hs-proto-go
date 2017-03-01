// Code generated by protoc-gen-go.
// source: bnet/protocol/presence/presence.proto
// DO NOT EDIT!

/*
Package bnet_protocol_presence is a generated protocol buffer package.

It is generated from these files:
	bnet/protocol/presence/presence.proto

It has these top-level messages:
	ChannelState
	Field
	FieldKey
	FieldOperation
	OwnershipRequest
	QueryRequest
	QueryResponse
	RichPresence
	SubscribeNotificationRequest
	SubscribeRequest
	UnsubscribeRequest
	UpdateRequest
*/
package bnet_protocol_presence

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bnet_protocol "bnet/protocol"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FieldOperation_OperationType int32

const (
	FieldOperation_SET   FieldOperation_OperationType = 0
	FieldOperation_CLEAR FieldOperation_OperationType = 1
)

var FieldOperation_OperationType_name = map[int32]string{
	0: "SET",
	1: "CLEAR",
}
var FieldOperation_OperationType_value = map[string]int32{
	"SET":   0,
	"CLEAR": 1,
}

func (x FieldOperation_OperationType) Enum() *FieldOperation_OperationType {
	p := new(FieldOperation_OperationType)
	*p = x
	return p
}
func (x FieldOperation_OperationType) String() string {
	return proto.EnumName(FieldOperation_OperationType_name, int32(x))
}
func (x *FieldOperation_OperationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FieldOperation_OperationType_value, data, "FieldOperation_OperationType")
	if err != nil {
		return err
	}
	*x = FieldOperation_OperationType(value)
	return nil
}
func (FieldOperation_OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type ChannelState struct {
	EntityId         *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	FieldOperation   []*FieldOperation       `protobuf:"bytes,2,rep,name=field_operation,json=fieldOperation" json:"field_operation,omitempty"`
	Healing          *bool                   `protobuf:"varint,3,opt,name=healing,def=0" json:"healing,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ChannelState) Reset()                    { *m = ChannelState{} }
func (m *ChannelState) String() string            { return proto.CompactTextString(m) }
func (*ChannelState) ProtoMessage()               {}
func (*ChannelState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ChannelState_Healing bool = false

func (m *ChannelState) GetEntityId() *bnet_protocol.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *ChannelState) GetFieldOperation() []*FieldOperation {
	if m != nil {
		return m.FieldOperation
	}
	return nil
}

func (m *ChannelState) GetHealing() bool {
	if m != nil && m.Healing != nil {
		return *m.Healing
	}
	return Default_ChannelState_Healing
}

type Field struct {
	Key              *FieldKey              `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *bnet_protocol.Variant `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Field) GetKey() *FieldKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Field) GetValue() *bnet_protocol.Variant {
	if m != nil {
		return m.Value
	}
	return nil
}

type FieldKey struct {
	Program          *uint32 `protobuf:"varint,1,req,name=program" json:"program,omitempty"`
	Group            *uint32 `protobuf:"varint,2,req,name=group" json:"group,omitempty"`
	Field            *uint32 `protobuf:"varint,3,req,name=field" json:"field,omitempty"`
	Index            *uint64 `protobuf:"varint,4,opt,name=index,def=0" json:"index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FieldKey) Reset()                    { *m = FieldKey{} }
func (m *FieldKey) String() string            { return proto.CompactTextString(m) }
func (*FieldKey) ProtoMessage()               {}
func (*FieldKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_FieldKey_Index uint64 = 0

func (m *FieldKey) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *FieldKey) GetGroup() uint32 {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return 0
}

func (m *FieldKey) GetField() uint32 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *FieldKey) GetIndex() uint64 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return Default_FieldKey_Index
}

type FieldOperation struct {
	Field            *Field                        `protobuf:"bytes,1,req,name=field" json:"field,omitempty"`
	Operation        *FieldOperation_OperationType `protobuf:"varint,2,opt,name=operation,enum=bnet.protocol.presence.FieldOperation_OperationType,def=0" json:"operation,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *FieldOperation) Reset()                    { *m = FieldOperation{} }
func (m *FieldOperation) String() string            { return proto.CompactTextString(m) }
func (*FieldOperation) ProtoMessage()               {}
func (*FieldOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_FieldOperation_Operation FieldOperation_OperationType = FieldOperation_SET

func (m *FieldOperation) GetField() *Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *FieldOperation) GetOperation() FieldOperation_OperationType {
	if m != nil && m.Operation != nil {
		return *m.Operation
	}
	return Default_FieldOperation_Operation
}

type OwnershipRequest struct {
	EntityId         *bnet_protocol.EntityId `protobuf:"bytes,1,req,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	ReleaseOwnership *bool                   `protobuf:"varint,2,opt,name=release_ownership,json=releaseOwnership,def=0" json:"release_ownership,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *OwnershipRequest) Reset()                    { *m = OwnershipRequest{} }
func (m *OwnershipRequest) String() string            { return proto.CompactTextString(m) }
func (*OwnershipRequest) ProtoMessage()               {}
func (*OwnershipRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

const Default_OwnershipRequest_ReleaseOwnership bool = false

func (m *OwnershipRequest) GetEntityId() *bnet_protocol.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *OwnershipRequest) GetReleaseOwnership() bool {
	if m != nil && m.ReleaseOwnership != nil {
		return *m.ReleaseOwnership
	}
	return Default_OwnershipRequest_ReleaseOwnership
}

type QueryRequest struct {
	EntityId         *bnet_protocol.EntityId `protobuf:"bytes,1,req,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	Key              []*FieldKey             `protobuf:"bytes,2,rep,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueryRequest) GetEntityId() *bnet_protocol.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *QueryRequest) GetKey() []*FieldKey {
	if m != nil {
		return m.Key
	}
	return nil
}

type QueryResponse struct {
	Field            []*Field `protobuf:"bytes,2,rep,name=field" json:"field,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueryResponse) GetField() []*Field {
	if m != nil {
		return m.Field
	}
	return nil
}

type RichPresence struct {
	ProgramId        *uint32 `protobuf:"fixed32,1,req,name=program_id,json=programId" json:"program_id,omitempty"`
	StreamId         *uint32 `protobuf:"fixed32,2,req,name=stream_id,json=streamId" json:"stream_id,omitempty"`
	Index            *uint32 `protobuf:"varint,3,req,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RichPresence) Reset()                    { *m = RichPresence{} }
func (m *RichPresence) String() string            { return proto.CompactTextString(m) }
func (*RichPresence) ProtoMessage()               {}
func (*RichPresence) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RichPresence) GetProgramId() uint32 {
	if m != nil && m.ProgramId != nil {
		return *m.ProgramId
	}
	return 0
}

func (m *RichPresence) GetStreamId() uint32 {
	if m != nil && m.StreamId != nil {
		return *m.StreamId
	}
	return 0
}

func (m *RichPresence) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

type SubscribeNotificationRequest struct {
	EntityId         *bnet_protocol.EntityId `protobuf:"bytes,1,req,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *SubscribeNotificationRequest) Reset()                    { *m = SubscribeNotificationRequest{} }
func (m *SubscribeNotificationRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeNotificationRequest) ProtoMessage()               {}
func (*SubscribeNotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SubscribeNotificationRequest) GetEntityId() *bnet_protocol.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

type SubscribeRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	EntityId         *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	ObjectId         *uint64                 `protobuf:"varint,3,req,name=object_id,json=objectId" json:"object_id,omitempty"`
	ProgramId        []uint32                `protobuf:"fixed32,4,rep,name=program_id,json=programId" json:"program_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SubscribeRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SubscribeRequest) GetEntityId() *bnet_protocol.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *SubscribeRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *SubscribeRequest) GetProgramId() []uint32 {
	if m != nil {
		return m.ProgramId
	}
	return nil
}

type UnsubscribeRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	EntityId         *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UnsubscribeRequest) Reset()                    { *m = UnsubscribeRequest{} }
func (m *UnsubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*UnsubscribeRequest) ProtoMessage()               {}
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UnsubscribeRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UnsubscribeRequest) GetEntityId() *bnet_protocol.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

type UpdateRequest struct {
	EntityId         *bnet_protocol.EntityId `protobuf:"bytes,1,req,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	FieldOperation   []*FieldOperation       `protobuf:"bytes,2,rep,name=field_operation,json=fieldOperation" json:"field_operation,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateRequest) GetEntityId() *bnet_protocol.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *UpdateRequest) GetFieldOperation() []*FieldOperation {
	if m != nil {
		return m.FieldOperation
	}
	return nil
}

func init() {
	proto.RegisterType((*ChannelState)(nil), "bnet.protocol.presence.ChannelState")
	proto.RegisterType((*Field)(nil), "bnet.protocol.presence.Field")
	proto.RegisterType((*FieldKey)(nil), "bnet.protocol.presence.FieldKey")
	proto.RegisterType((*FieldOperation)(nil), "bnet.protocol.presence.FieldOperation")
	proto.RegisterType((*OwnershipRequest)(nil), "bnet.protocol.presence.OwnershipRequest")
	proto.RegisterType((*QueryRequest)(nil), "bnet.protocol.presence.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "bnet.protocol.presence.QueryResponse")
	proto.RegisterType((*RichPresence)(nil), "bnet.protocol.presence.RichPresence")
	proto.RegisterType((*SubscribeNotificationRequest)(nil), "bnet.protocol.presence.SubscribeNotificationRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "bnet.protocol.presence.SubscribeRequest")
	proto.RegisterType((*UnsubscribeRequest)(nil), "bnet.protocol.presence.UnsubscribeRequest")
	proto.RegisterType((*UpdateRequest)(nil), "bnet.protocol.presence.UpdateRequest")
	proto.RegisterEnum("bnet.protocol.presence.FieldOperation_OperationType", FieldOperation_OperationType_name, FieldOperation_OperationType_value)
}

func init() { proto.RegisterFile("bnet/protocol/presence/presence.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x54, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0x26, 0xfd, 0xa1, 0xb6, 0xb7, 0x75, 0x14, 0x0b, 0x6d, 0x11, 0xdb, 0x44, 0x64, 0x04, 0xea,
	0x03, 0xca, 0x50, 0xd8, 0xd3, 0xde, 0xd0, 0x18, 0x52, 0x05, 0x62, 0xe0, 0x6d, 0xf0, 0x38, 0xdc,
	0xe6, 0xda, 0x1a, 0x82, 0x1d, 0x1c, 0x17, 0x56, 0x09, 0xf1, 0x67, 0xf0, 0xa7, 0x20, 0xde, 0xf9,
	0xc7, 0x50, 0xec, 0xa4, 0xa3, 0x15, 0x9a, 0x22, 0x0d, 0x89, 0x37, 0xdf, 0xdd, 0xe7, 0xef, 0xcb,
	0xdd, 0x7d, 0x31, 0xdc, 0x1f, 0x4a, 0x34, 0x7b, 0xa9, 0x56, 0x46, 0x8d, 0x54, 0xb2, 0x97, 0x6a,
	0xcc, 0x50, 0x8e, 0x70, 0x71, 0x08, 0x6d, 0x89, 0x6c, 0xe6, 0xb0, 0xb0, 0x84, 0x85, 0x65, 0xf5,
	0xce, 0xce, 0xea, 0xf5, 0x05, 0x40, 0x19, 0x45, 0x7f, 0x78, 0xb0, 0x7e, 0x38, 0xe5, 0x52, 0x62,
	0x72, 0x62, 0xb8, 0x41, 0xb2, 0x0f, 0x1d, 0x94, 0x46, 0x98, 0xf9, 0xb9, 0x88, 0x7d, 0x2f, 0xf0,
	0xfa, 0x6b, 0xd1, 0x56, 0xb8, 0x4c, 0x7d, 0x64, 0xeb, 0x83, 0x98, 0xb5, 0xb1, 0x38, 0x91, 0x63,
	0xb8, 0x39, 0x16, 0x98, 0xc4, 0xe7, 0x2a, 0x45, 0xcd, 0x8d, 0x50, 0xd2, 0xaf, 0x05, 0xf5, 0xfe,
	0x5a, 0xf4, 0x20, 0xfc, 0xfb, 0x67, 0x85, 0xcf, 0x72, 0xf8, 0x71, 0x89, 0x66, 0x1b, 0xe3, 0xa5,
	0x98, 0xdc, 0x85, 0xd6, 0x14, 0x79, 0x22, 0xe4, 0xc4, 0xaf, 0x07, 0x5e, 0xbf, 0x7d, 0xd0, 0x1c,
	0xf3, 0x24, 0x43, 0x56, 0x66, 0xa9, 0x80, 0xa6, 0xa5, 0x20, 0x11, 0xd4, 0x3f, 0xe0, 0xdc, 0xf7,
	0x82, 0x5a, 0x7f, 0x2d, 0x0a, 0xae, 0x94, 0x7b, 0x8e, 0x73, 0x96, 0x83, 0xc9, 0x43, 0x68, 0x7e,
	0xe6, 0xc9, 0x0c, 0xfd, 0x9a, 0xbd, 0xb5, 0xb9, 0x72, 0xeb, 0x0d, 0xd7, 0x82, 0x4b, 0xc3, 0x1c,
	0x88, 0x0a, 0x68, 0x97, 0xd7, 0x89, 0x0f, 0xad, 0x54, 0xab, 0x89, 0xe6, 0x1f, 0xad, 0x62, 0x97,
	0x95, 0x21, 0xb9, 0x0d, 0xcd, 0x89, 0x56, 0xb3, 0xd4, 0x72, 0x76, 0x99, 0x0b, 0xf2, 0xac, 0xed,
	0xcc, 0xaf, 0xbb, 0xac, 0x0d, 0xc8, 0x16, 0x34, 0x85, 0x8c, 0xf1, 0xc2, 0x6f, 0x04, 0x5e, 0xbf,
	0x71, 0xe0, 0x3d, 0x62, 0x2e, 0xa6, 0xbf, 0x3c, 0xd8, 0x58, 0x9e, 0x0c, 0x79, 0x5c, 0x32, 0xb8,
	0x0e, 0x77, 0xaf, 0xec, 0xb0, 0x14, 0x78, 0x0b, 0x9d, 0x3f, 0x37, 0xe1, 0xf5, 0x37, 0xa2, 0xfd,
	0x6a, 0x9b, 0x08, 0x17, 0xa7, 0xd3, 0x79, 0x8a, 0x07, 0xf5, 0x93, 0xa3, 0x53, 0x76, 0xc9, 0x45,
	0xef, 0x41, 0x77, 0x09, 0x40, 0x5a, 0x90, 0x43, 0x7a, 0x37, 0x48, 0x07, 0x9a, 0x87, 0x2f, 0x8e,
	0x9e, 0xb0, 0x9e, 0x47, 0xbf, 0x42, 0xef, 0xf8, 0x8b, 0x44, 0x9d, 0x4d, 0x45, 0xca, 0xf0, 0xd3,
	0x0c, 0x33, 0xb3, 0xea, 0xab, 0x5a, 0x35, 0x5f, 0x45, 0x70, 0x4b, 0x63, 0x82, 0x3c, 0xc3, 0x73,
	0x55, 0x32, 0xda, 0x7e, 0x16, 0x86, 0xe8, 0x15, 0xf5, 0x85, 0x20, 0xbd, 0x80, 0xf5, 0xd7, 0x33,
	0xd4, 0xf3, 0xeb, 0x2a, 0x5b, 0x5b, 0x39, 0x17, 0x57, 0xb3, 0x15, 0x7d, 0x0a, 0xdd, 0x42, 0x39,
	0x4b, 0x95, 0xcc, 0xf0, 0x72, 0x77, 0x8e, 0xa6, 0xd2, 0xee, 0xe8, 0x3b, 0x58, 0x67, 0x62, 0x34,
	0x7d, 0x55, 0x14, 0xc9, 0x2e, 0x40, 0xe1, 0xb1, 0xb2, 0x81, 0x16, 0xeb, 0x14, 0x99, 0x41, 0x4c,
	0xb6, 0xa1, 0x93, 0x19, 0x8d, 0xae, 0x5a, 0xb3, 0xd5, 0xb6, 0x4b, 0x0c, 0xe2, 0xdc, 0x7e, 0xce,
	0x68, 0x85, 0xfd, 0x9c, 0xcb, 0x4e, 0x61, 0xe7, 0x64, 0x36, 0xcc, 0x46, 0x5a, 0x0c, 0xf1, 0xa5,
	0x32, 0x62, 0x2c, 0x46, 0xee, 0x2f, 0xbc, 0xce, 0xc4, 0xe8, 0x4f, 0x0f, 0x7a, 0x0b, 0xda, 0x92,
	0x2a, 0x82, 0x36, 0x9f, 0xa0, 0x34, 0x15, 0x5e, 0x93, 0x96, 0x05, 0x0e, 0xe2, 0x65, 0xf9, 0x5a,
	0xd5, 0x85, 0x6d, 0x43, 0x47, 0x0d, 0xdf, 0xe3, 0xc8, 0x4a, 0xe5, 0xed, 0x36, 0x58, 0xdb, 0x25,
	0x06, 0xf1, 0xca, 0x0c, 0x1b, 0x41, 0x7d, 0x69, 0x86, 0xf4, 0x1b, 0x90, 0x33, 0x99, 0xfd, 0xb7,
	0x6f, 0xa7, 0xdf, 0x3d, 0xe8, 0x9e, 0xa5, 0x31, 0x37, 0x78, 0x3d, 0xd3, 0xfe, 0xeb, 0x67, 0xf8,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x03, 0xa5, 0xfb, 0x7c, 0x06, 0x00, 0x00,
}
