// Code generated by protoc-gen-go.
// source: bnet/protocol/friends/friends.proto
// DO NOT EDIT!

/*
Package bnet_protocol_friends is a generated protocol buffer package.

It is generated from these files:
	bnet/protocol/friends/friends.proto

It has these top-level messages:
	AssignRoleRequest
	Friend
	FriendInvitation
	FriendInvitationParams
	FriendNotification
	GenericFriendRequest
	GenericFriendResponse
	SubscribeToFriendsRequest
	UnsubscribeToFriendsRequest
	UpdateFriendStateNotification
	UpdateFriendStateRequest
	ViewFriendsRequest
	ViewFriendsResponse
*/
package bnet_protocol_friends

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

type AssignRoleRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	TargetId         *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=target_id,json=targetId" json:"target_id,omitempty"`
	Role             []int32                 `protobuf:"varint,3,rep,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *AssignRoleRequest) Reset()                    { *m = AssignRoleRequest{} }
func (m *AssignRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*AssignRoleRequest) ProtoMessage()               {}
func (*AssignRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AssignRoleRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *AssignRoleRequest) GetTargetId() *bnet_protocol.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *AssignRoleRequest) GetRole() []int32 {
	if m != nil {
		return m.Role
	}
	return nil
}

type Friend struct {
	Id               *bnet_protocol.EntityId    `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Attribute        []*bnet_protocol.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	Role             []uint32                   `protobuf:"varint,3,rep,packed,name=role" json:"role,omitempty"`
	Privileges       *uint64                    `protobuf:"varint,4,opt,name=privileges,def=0" json:"privileges,omitempty"`
	AttributesEpoch  *uint64                    `protobuf:"varint,5,opt,name=attributes_epoch,json=attributesEpoch" json:"attributes_epoch,omitempty"`
	FullName         *string                    `protobuf:"bytes,6,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	BattleTag        *string                    `protobuf:"bytes,7,opt,name=battle_tag,json=battleTag" json:"battle_tag,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Friend) Reset()                    { *m = Friend{} }
func (m *Friend) String() string            { return proto.CompactTextString(m) }
func (*Friend) ProtoMessage()               {}
func (*Friend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_Friend_Privileges uint64 = 0

func (m *Friend) GetId() *bnet_protocol.EntityId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Friend) GetAttribute() []*bnet_protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *Friend) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *Friend) GetPrivileges() uint64 {
	if m != nil && m.Privileges != nil {
		return *m.Privileges
	}
	return Default_Friend_Privileges
}

func (m *Friend) GetAttributesEpoch() uint64 {
	if m != nil && m.AttributesEpoch != nil {
		return *m.AttributesEpoch
	}
	return 0
}

func (m *Friend) GetFullName() string {
	if m != nil && m.FullName != nil {
		return *m.FullName
	}
	return ""
}

func (m *Friend) GetBattleTag() string {
	if m != nil && m.BattleTag != nil {
		return *m.BattleTag
	}
	return ""
}

type FriendInvitation struct {
	FirstReceived    *bool    `protobuf:"varint,1,opt,name=first_received,json=firstReceived,def=0" json:"first_received,omitempty"`
	Role             []uint32 `protobuf:"varint,2,rep,packed,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FriendInvitation) Reset()                    { *m = FriendInvitation{} }
func (m *FriendInvitation) String() string            { return proto.CompactTextString(m) }
func (*FriendInvitation) ProtoMessage()               {}
func (*FriendInvitation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_FriendInvitation_FirstReceived bool = false

func (m *FriendInvitation) GetFirstReceived() bool {
	if m != nil && m.FirstReceived != nil {
		return *m.FirstReceived
	}
	return Default_FriendInvitation_FirstReceived
}

func (m *FriendInvitation) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

type FriendInvitationParams struct {
	TargetEmail        *string  `protobuf:"bytes,1,opt,name=target_email,json=targetEmail" json:"target_email,omitempty"`
	TargetBattleTag    *string  `protobuf:"bytes,2,opt,name=target_battle_tag,json=targetBattleTag" json:"target_battle_tag,omitempty"`
	InviterBattleTag   *string  `protobuf:"bytes,3,opt,name=inviter_battle_tag,json=inviterBattleTag" json:"inviter_battle_tag,omitempty"`
	InviterFullName    *string  `protobuf:"bytes,4,opt,name=inviter_full_name,json=inviterFullName" json:"inviter_full_name,omitempty"`
	InviteeDisplayName *string  `protobuf:"bytes,5,opt,name=invitee_display_name,json=inviteeDisplayName" json:"invitee_display_name,omitempty"`
	Role               []uint32 `protobuf:"varint,6,rep,packed,name=role" json:"role,omitempty"`
	XXX_unrecognized   []byte   `json:"-"`
}

func (m *FriendInvitationParams) Reset()                    { *m = FriendInvitationParams{} }
func (m *FriendInvitationParams) String() string            { return proto.CompactTextString(m) }
func (*FriendInvitationParams) ProtoMessage()               {}
func (*FriendInvitationParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FriendInvitationParams) GetTargetEmail() string {
	if m != nil && m.TargetEmail != nil {
		return *m.TargetEmail
	}
	return ""
}

func (m *FriendInvitationParams) GetTargetBattleTag() string {
	if m != nil && m.TargetBattleTag != nil {
		return *m.TargetBattleTag
	}
	return ""
}

func (m *FriendInvitationParams) GetInviterBattleTag() string {
	if m != nil && m.InviterBattleTag != nil {
		return *m.InviterBattleTag
	}
	return ""
}

func (m *FriendInvitationParams) GetInviterFullName() string {
	if m != nil && m.InviterFullName != nil {
		return *m.InviterFullName
	}
	return ""
}

func (m *FriendInvitationParams) GetInviteeDisplayName() string {
	if m != nil && m.InviteeDisplayName != nil {
		return *m.InviteeDisplayName
	}
	return ""
}

func (m *FriendInvitationParams) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

type FriendNotification struct {
	Target           *Friend                 `protobuf:"bytes,1,req,name=target" json:"target,omitempty"`
	GameAccountId    *bnet_protocol.EntityId `protobuf:"bytes,2,opt,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *FriendNotification) Reset()                    { *m = FriendNotification{} }
func (m *FriendNotification) String() string            { return proto.CompactTextString(m) }
func (*FriendNotification) ProtoMessage()               {}
func (*FriendNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FriendNotification) GetTarget() *Friend {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *FriendNotification) GetGameAccountId() *bnet_protocol.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

type GenericFriendRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	TargetId         *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=target_id,json=targetId" json:"target_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *GenericFriendRequest) Reset()                    { *m = GenericFriendRequest{} }
func (m *GenericFriendRequest) String() string            { return proto.CompactTextString(m) }
func (*GenericFriendRequest) ProtoMessage()               {}
func (*GenericFriendRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GenericFriendRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *GenericFriendRequest) GetTargetId() *bnet_protocol.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

type GenericFriendResponse struct {
	TargetFriend     *Friend `protobuf:"bytes,1,opt,name=target_friend,json=targetFriend" json:"target_friend,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GenericFriendResponse) Reset()                    { *m = GenericFriendResponse{} }
func (m *GenericFriendResponse) String() string            { return proto.CompactTextString(m) }
func (*GenericFriendResponse) ProtoMessage()               {}
func (*GenericFriendResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GenericFriendResponse) GetTargetFriend() *Friend {
	if m != nil {
		return m.TargetFriend
	}
	return nil
}

type SubscribeToFriendsRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	ObjectId         *uint64                 `protobuf:"varint,2,req,name=object_id,json=objectId" json:"object_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *SubscribeToFriendsRequest) Reset()                    { *m = SubscribeToFriendsRequest{} }
func (m *SubscribeToFriendsRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeToFriendsRequest) ProtoMessage()               {}
func (*SubscribeToFriendsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SubscribeToFriendsRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SubscribeToFriendsRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

type UnsubscribeToFriendsRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	ObjectId         *uint64                 `protobuf:"varint,2,opt,name=object_id,json=objectId" json:"object_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UnsubscribeToFriendsRequest) Reset()                    { *m = UnsubscribeToFriendsRequest{} }
func (m *UnsubscribeToFriendsRequest) String() string            { return proto.CompactTextString(m) }
func (*UnsubscribeToFriendsRequest) ProtoMessage()               {}
func (*UnsubscribeToFriendsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UnsubscribeToFriendsRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UnsubscribeToFriendsRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

type UpdateFriendStateNotification struct {
	ChangedFriend    *Friend                 `protobuf:"bytes,1,req,name=changed_friend,json=changedFriend" json:"changed_friend,omitempty"`
	GameAccountId    *bnet_protocol.EntityId `protobuf:"bytes,2,opt,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdateFriendStateNotification) Reset()                    { *m = UpdateFriendStateNotification{} }
func (m *UpdateFriendStateNotification) String() string            { return proto.CompactTextString(m) }
func (*UpdateFriendStateNotification) ProtoMessage()               {}
func (*UpdateFriendStateNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateFriendStateNotification) GetChangedFriend() *Friend {
	if m != nil {
		return m.ChangedFriend
	}
	return nil
}

func (m *UpdateFriendStateNotification) GetGameAccountId() *bnet_protocol.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

type UpdateFriendStateRequest struct {
	AgentId          *bnet_protocol.EntityId    `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	TargetId         *bnet_protocol.EntityId    `protobuf:"bytes,2,req,name=target_id,json=targetId" json:"target_id,omitempty"`
	Attribute        []*bnet_protocol.Attribute `protobuf:"bytes,3,rep,name=attribute" json:"attribute,omitempty"`
	AttributesEpoch  *uint64                    `protobuf:"varint,4,opt,name=attributes_epoch,json=attributesEpoch" json:"attributes_epoch,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *UpdateFriendStateRequest) Reset()                    { *m = UpdateFriendStateRequest{} }
func (m *UpdateFriendStateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateFriendStateRequest) ProtoMessage()               {}
func (*UpdateFriendStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateFriendStateRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateFriendStateRequest) GetTargetId() *bnet_protocol.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *UpdateFriendStateRequest) GetAttribute() []*bnet_protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *UpdateFriendStateRequest) GetAttributesEpoch() uint64 {
	if m != nil && m.AttributesEpoch != nil {
		return *m.AttributesEpoch
	}
	return 0
}

type ViewFriendsRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	TargetId         *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=target_id,json=targetId" json:"target_id,omitempty"`
	Role             []uint32                `protobuf:"varint,3,rep,packed,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ViewFriendsRequest) Reset()                    { *m = ViewFriendsRequest{} }
func (m *ViewFriendsRequest) String() string            { return proto.CompactTextString(m) }
func (*ViewFriendsRequest) ProtoMessage()               {}
func (*ViewFriendsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ViewFriendsRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *ViewFriendsRequest) GetTargetId() *bnet_protocol.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *ViewFriendsRequest) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

type ViewFriendsResponse struct {
	Friends          []*Friend `protobuf:"bytes,1,rep,name=friends" json:"friends,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ViewFriendsResponse) Reset()                    { *m = ViewFriendsResponse{} }
func (m *ViewFriendsResponse) String() string            { return proto.CompactTextString(m) }
func (*ViewFriendsResponse) ProtoMessage()               {}
func (*ViewFriendsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ViewFriendsResponse) GetFriends() []*Friend {
	if m != nil {
		return m.Friends
	}
	return nil
}

func init() {
	proto.RegisterType((*AssignRoleRequest)(nil), "bnet.protocol.friends.AssignRoleRequest")
	proto.RegisterType((*Friend)(nil), "bnet.protocol.friends.Friend")
	proto.RegisterType((*FriendInvitation)(nil), "bnet.protocol.friends.FriendInvitation")
	proto.RegisterType((*FriendInvitationParams)(nil), "bnet.protocol.friends.FriendInvitationParams")
	proto.RegisterType((*FriendNotification)(nil), "bnet.protocol.friends.FriendNotification")
	proto.RegisterType((*GenericFriendRequest)(nil), "bnet.protocol.friends.GenericFriendRequest")
	proto.RegisterType((*GenericFriendResponse)(nil), "bnet.protocol.friends.GenericFriendResponse")
	proto.RegisterType((*SubscribeToFriendsRequest)(nil), "bnet.protocol.friends.SubscribeToFriendsRequest")
	proto.RegisterType((*UnsubscribeToFriendsRequest)(nil), "bnet.protocol.friends.UnsubscribeToFriendsRequest")
	proto.RegisterType((*UpdateFriendStateNotification)(nil), "bnet.protocol.friends.UpdateFriendStateNotification")
	proto.RegisterType((*UpdateFriendStateRequest)(nil), "bnet.protocol.friends.UpdateFriendStateRequest")
	proto.RegisterType((*ViewFriendsRequest)(nil), "bnet.protocol.friends.ViewFriendsRequest")
	proto.RegisterType((*ViewFriendsResponse)(nil), "bnet.protocol.friends.ViewFriendsResponse")
}

func init() { proto.RegisterFile("bnet/protocol/friends/friends.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0x5d, 0x4f, 0x13, 0x4d,
	0x14, 0xce, 0x6e, 0x3f, 0x68, 0x0f, 0x6f, 0xa1, 0xcc, 0x0b, 0xb8, 0x8a, 0x24, 0x65, 0xbd, 0xb0,
	0x12, 0x52, 0x08, 0xf1, 0x23, 0xe1, 0xc6, 0x40, 0x00, 0xd3, 0x1b, 0x62, 0x16, 0x30, 0x26, 0x5e,
	0x34, 0xd3, 0xdd, 0xd3, 0x65, 0xcc, 0x76, 0xb6, 0xee, 0x4c, 0x31, 0xdc, 0x79, 0xaf, 0x37, 0xde,
	0xf8, 0x13, 0xfc, 0x6b, 0xfe, 0x0d, 0xb3, 0x33, 0xb3, 0x6d, 0xb7, 0x22, 0x10, 0x51, 0xae, 0xba,
	0x3d, 0xcf, 0x73, 0xbe, 0x9e, 0x39, 0x73, 0x06, 0x1e, 0x75, 0x39, 0xca, 0xcd, 0x41, 0x12, 0xcb,
	0xd8, 0x8f, 0xa3, 0xcd, 0x5e, 0xc2, 0x90, 0x07, 0x22, 0xfb, 0x6d, 0x29, 0x80, 0x2c, 0xa5, 0xa4,
	0x56, 0x46, 0x6a, 0x19, 0xf0, 0xc1, 0xc3, 0xbc, 0xef, 0x08, 0x57, 0x1f, 0xee, 0x57, 0x0b, 0x16,
	0x76, 0x85, 0x60, 0x21, 0xf7, 0xe2, 0x08, 0x3d, 0xfc, 0x30, 0x44, 0x21, 0xc9, 0x36, 0x54, 0x68,
	0x88, 0x5c, 0x76, 0x58, 0xe0, 0x58, 0x0d, 0xab, 0x39, 0xbb, 0x7d, 0xaf, 0x95, 0x8f, 0x7e, 0xc0,
	0x25, 0x93, 0x17, 0xed, 0xc0, 0x9b, 0x51, 0xc4, 0x76, 0x40, 0x9e, 0x42, 0x55, 0xd2, 0x24, 0x44,
	0xe5, 0x64, 0x37, 0xec, 0xab, 0x9c, 0x2a, 0x9a, 0xd9, 0x0e, 0x08, 0x81, 0x62, 0x12, 0x47, 0xe8,
	0x14, 0x1a, 0x85, 0x66, 0xc9, 0x53, 0xdf, 0xee, 0x67, 0x1b, 0xca, 0x87, 0xaa, 0x7a, 0xf2, 0x18,
	0x6c, 0x55, 0xc2, 0x95, 0xd1, 0x6c, 0x16, 0x90, 0xe7, 0x50, 0xa5, 0x52, 0x26, 0xac, 0x3b, 0x94,
	0xe8, 0xd8, 0x8d, 0x42, 0x73, 0x76, 0xdb, 0x99, 0xe2, 0xef, 0x66, 0xb8, 0x37, 0xa6, 0x92, 0xe5,
	0x89, 0xfc, 0xb5, 0x3d, 0xbb, 0x6e, 0xe9, 0x1a, 0xc8, 0x1a, 0xc0, 0x20, 0x61, 0xe7, 0x2c, 0xc2,
	0x10, 0x85, 0x53, 0x6c, 0x58, 0xcd, 0xe2, 0x8e, 0xb5, 0xe5, 0x4d, 0x18, 0xc9, 0x13, 0xa8, 0x8f,
	0xe2, 0x88, 0x0e, 0x0e, 0x62, 0xff, 0xcc, 0x29, 0xa5, 0x44, 0x6f, 0x7e, 0x6c, 0x3f, 0x48, 0xcd,
	0x64, 0x05, 0xaa, 0xbd, 0x61, 0x14, 0x75, 0x38, 0xed, 0xa3, 0x53, 0x6e, 0x58, 0xcd, 0xaa, 0x57,
	0x49, 0x0d, 0x47, 0xb4, 0x8f, 0x64, 0x15, 0xa0, 0x4b, 0xa5, 0x8c, 0xb0, 0x23, 0x69, 0xe8, 0xcc,
	0x28, 0xb4, 0xaa, 0x2d, 0x27, 0x34, 0x74, 0xdf, 0x42, 0x5d, 0x8b, 0xd1, 0xe6, 0xe7, 0x4c, 0x52,
	0xc9, 0x62, 0x4e, 0x36, 0x60, 0xae, 0xc7, 0x12, 0x21, 0x3b, 0x09, 0xfa, 0xc8, 0xce, 0x51, 0x9f,
	0x52, 0x65, 0xa7, 0xd4, 0xa3, 0x91, 0x40, 0xaf, 0xa6, 0x40, 0xcf, 0x60, 0xa3, 0x1e, 0xed, 0x7c,
	0x8f, 0xa9, 0xce, 0xcb, 0xd3, 0xa1, 0x5f, 0xd3, 0x84, 0xf6, 0x05, 0x59, 0x83, 0xff, 0xcc, 0x61,
	0x62, 0x9f, 0xb2, 0x48, 0x85, 0xaf, 0x7a, 0xb3, 0xda, 0x76, 0x90, 0x9a, 0xc8, 0x3a, 0x2c, 0x18,
	0xca, 0x44, 0xf5, 0xb6, 0xe2, 0xcd, 0x6b, 0x60, 0x2f, 0xeb, 0x81, 0x6c, 0x00, 0x61, 0x69, 0x0a,
	0x4c, 0x26, 0xc9, 0x05, 0x45, 0xae, 0x1b, 0x64, 0xcc, 0x5e, 0x87, 0x85, 0x8c, 0x3d, 0x56, 0xad,
	0xa8, 0x23, 0x1b, 0xe0, 0x30, 0x13, 0x6f, 0x0b, 0x16, 0xb5, 0x09, 0x3b, 0x01, 0x13, 0x83, 0x88,
	0x5e, 0x68, 0x7a, 0x49, 0xd1, 0x4d, 0x56, 0xdc, 0xd7, 0x90, 0xf2, 0xc8, 0xd4, 0x28, 0x4f, 0xa9,
	0xf1, 0xc5, 0x02, 0xa2, 0xd5, 0x38, 0x8a, 0x25, 0xeb, 0x31, 0x5f, 0x4b, 0xfd, 0x0c, 0xca, 0xba,
	0x1b, 0x33, 0x85, 0xab, 0xad, 0x4b, 0xaf, 0x59, 0x4b, 0xbb, 0x7a, 0x86, 0x4c, 0x5e, 0xc2, 0x7c,
	0x48, 0xfb, 0xd8, 0xa1, 0xbe, 0x1f, 0x0f, 0xb9, 0xb9, 0x13, 0x57, 0x5e, 0xa4, 0x5a, 0xca, 0xdf,
	0xd5, 0xf4, 0x76, 0xe0, 0x7e, 0xb2, 0x60, 0xf1, 0x15, 0x72, 0x4c, 0x98, 0x6f, 0x42, 0xdf, 0xf5,
	0xdd, 0x74, 0xdf, 0xc1, 0xd2, 0x54, 0x05, 0x62, 0x10, 0x73, 0x81, 0x64, 0x0f, 0x6a, 0x26, 0x9c,
	0xee, 0xde, 0xd4, 0x71, 0x8d, 0x34, 0x66, 0xa2, 0xf4, 0x3f, 0x37, 0x82, 0xfb, 0xc7, 0xc3, 0xae,
	0xf0, 0x13, 0xd6, 0xc5, 0x93, 0x58, 0x1b, 0xc5, 0x6d, 0x7a, 0x5c, 0x81, 0x6a, 0xdc, 0x7d, 0x8f,
	0xfe, 0xa8, 0xc7, 0xa2, 0x57, 0xd1, 0x86, 0x76, 0xe0, 0x72, 0x58, 0x39, 0xe5, 0xe2, 0x5f, 0xe6,
	0xb3, 0x72, 0xf9, 0xbe, 0x5b, 0xb0, 0x7a, 0x3a, 0x08, 0xa8, 0x44, 0x9d, 0xe9, 0x58, 0x52, 0x89,
	0xb9, 0xb9, 0xda, 0x87, 0x39, 0xff, 0x8c, 0xf2, 0x10, 0x83, 0xb1, 0x88, 0x37, 0x98, 0xaf, 0x9a,
	0x71, 0x32, 0xfb, 0xf1, 0xd6, 0x63, 0xf6, 0xc3, 0x02, 0xe7, 0x97, 0x42, 0xef, 0xfe, 0x19, 0xc8,
	0xad, 0xef, 0xc2, 0xcd, 0xd7, 0xf7, 0x65, 0x3b, 0xb8, 0x78, 0xe9, 0x0e, 0x76, 0xbf, 0x59, 0x40,
	0xde, 0x30, 0xfc, 0xf8, 0x17, 0x8e, 0xfe, 0xcf, 0x7a, 0xfc, 0xcd, 0x53, 0xe3, 0x1e, 0xc1, 0xff,
	0xb9, 0xba, 0xcc, 0x25, 0x7b, 0x01, 0x33, 0xe6, 0xec, 0x1d, 0x4b, 0x09, 0x72, 0xcd, 0x64, 0x64,
	0xec, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x97, 0x6d, 0x78, 0x2d, 0x08, 0x00, 0x00,
}
