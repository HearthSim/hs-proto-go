// Code generated by protoc-gen-go.
// source: bnet/protocol/authentication/authentication.proto
// DO NOT EDIT!

/*
Package bnet_protocol_authentication is a generated protocol buffer package.

It is generated from these files:
	bnet/protocol/authentication/authentication.proto

It has these top-level messages:
	AccountSettingsNotification
	GameAccountSelectedRequest
	GenerateSSOTokenRequest
	GenerateSSOTokenResponse
	LogonQueueUpdateRequest
	LogonRequest
	LogonResult
	LogonUpdateRequest
	MemModuleLoadRequest
	MemModuleLoadResponse
	ModuleLoadRequest
	ModuleMessageRequest
	ModuleNotification
	SelectGameAccountRequest
	ServerStateChangeRequest
	VerifyWebCredentialsRequest
	VersionInfo
	VersionInfoNotification
*/
package bnet_protocol_authentication

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bnet_protocol_account "github.com/HearthSim/hs-proto-go/bnet/protocol/account"
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

// ref: bnet.protocol.authentication.AccountSettingsNotification
type AccountSettingsNotification struct {
	Licenses         []*bnet_protocol_account.AccountLicense `protobuf:"bytes,1,rep,name=licenses" json:"licenses,omitempty"`
	IsUsingRid       *bool                                   `protobuf:"varint,2,opt,name=is_using_rid,json=isUsingRid" json:"is_using_rid,omitempty"`
	IsPlayingFromIgr *bool                                   `protobuf:"varint,3,opt,name=is_playing_from_igr,json=isPlayingFromIgr" json:"is_playing_from_igr,omitempty"`
	CanReceiveVoice  *bool                                   `protobuf:"varint,4,opt,name=can_receive_voice,json=canReceiveVoice" json:"can_receive_voice,omitempty"`
	CanSendVoice     *bool                                   `protobuf:"varint,5,opt,name=can_send_voice,json=canSendVoice" json:"can_send_voice,omitempty"`
	XXX_unrecognized []byte                                  `json:"-"`
}

func (m *AccountSettingsNotification) Reset()                    { *m = AccountSettingsNotification{} }
func (m *AccountSettingsNotification) String() string            { return proto.CompactTextString(m) }
func (*AccountSettingsNotification) ProtoMessage()               {}
func (*AccountSettingsNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AccountSettingsNotification) GetLicenses() []*bnet_protocol_account.AccountLicense {
	if m != nil {
		return m.Licenses
	}
	return nil
}

func (m *AccountSettingsNotification) GetIsUsingRid() bool {
	if m != nil && m.IsUsingRid != nil {
		return *m.IsUsingRid
	}
	return false
}

func (m *AccountSettingsNotification) GetIsPlayingFromIgr() bool {
	if m != nil && m.IsPlayingFromIgr != nil {
		return *m.IsPlayingFromIgr
	}
	return false
}

func (m *AccountSettingsNotification) GetCanReceiveVoice() bool {
	if m != nil && m.CanReceiveVoice != nil {
		return *m.CanReceiveVoice
	}
	return false
}

func (m *AccountSettingsNotification) GetCanSendVoice() bool {
	if m != nil && m.CanSendVoice != nil {
		return *m.CanSendVoice
	}
	return false
}

// ref: bnet.protocol.authentication.GameAccountSelectedRequest
type GameAccountSelectedRequest struct {
	Result           *uint32                 `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	GameAccount      *bnet_protocol.EntityId `protobuf:"bytes,2,opt,name=game_account,json=gameAccount" json:"game_account,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *GameAccountSelectedRequest) Reset()                    { *m = GameAccountSelectedRequest{} }
func (m *GameAccountSelectedRequest) String() string            { return proto.CompactTextString(m) }
func (*GameAccountSelectedRequest) ProtoMessage()               {}
func (*GameAccountSelectedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GameAccountSelectedRequest) GetResult() uint32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *GameAccountSelectedRequest) GetGameAccount() *bnet_protocol.EntityId {
	if m != nil {
		return m.GameAccount
	}
	return nil
}

// ref: bnet.protocol.authentication.GenerateSSOTokenRequest
type GenerateSSOTokenRequest struct {
	Program          *uint32 `protobuf:"fixed32,1,opt,name=program" json:"program,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GenerateSSOTokenRequest) Reset()                    { *m = GenerateSSOTokenRequest{} }
func (m *GenerateSSOTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*GenerateSSOTokenRequest) ProtoMessage()               {}
func (*GenerateSSOTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GenerateSSOTokenRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

// ref: bnet.protocol.authentication.GenerateSSOTokenResponse
type GenerateSSOTokenResponse struct {
	SsoId            []byte `protobuf:"bytes,1,opt,name=sso_id,json=ssoId" json:"sso_id,omitempty"`
	SsoSecret        []byte `protobuf:"bytes,2,opt,name=sso_secret,json=ssoSecret" json:"sso_secret,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GenerateSSOTokenResponse) Reset()                    { *m = GenerateSSOTokenResponse{} }
func (m *GenerateSSOTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*GenerateSSOTokenResponse) ProtoMessage()               {}
func (*GenerateSSOTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GenerateSSOTokenResponse) GetSsoId() []byte {
	if m != nil {
		return m.SsoId
	}
	return nil
}

func (m *GenerateSSOTokenResponse) GetSsoSecret() []byte {
	if m != nil {
		return m.SsoSecret
	}
	return nil
}

// ref: bnet.protocol.authentication.LogonQueueUpdateRequest
type LogonQueueUpdateRequest struct {
	Position          *uint32 `protobuf:"varint,1,req,name=position" json:"position,omitempty"`
	EstimatedTime     *uint64 `protobuf:"varint,2,req,name=estimated_time,json=estimatedTime" json:"estimated_time,omitempty"`
	EtaDeviationInSec *uint64 `protobuf:"varint,3,req,name=eta_deviation_in_sec,json=etaDeviationInSec" json:"eta_deviation_in_sec,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *LogonQueueUpdateRequest) Reset()                    { *m = LogonQueueUpdateRequest{} }
func (m *LogonQueueUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*LogonQueueUpdateRequest) ProtoMessage()               {}
func (*LogonQueueUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogonQueueUpdateRequest) GetPosition() uint32 {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return 0
}

func (m *LogonQueueUpdateRequest) GetEstimatedTime() uint64 {
	if m != nil && m.EstimatedTime != nil {
		return *m.EstimatedTime
	}
	return 0
}

func (m *LogonQueueUpdateRequest) GetEtaDeviationInSec() uint64 {
	if m != nil && m.EtaDeviationInSec != nil {
		return *m.EtaDeviationInSec
	}
	return 0
}

// ref: bnet.protocol.authentication.LogonRequest
type LogonRequest struct {
	Program                      *string `protobuf:"bytes,1,opt,name=program" json:"program,omitempty"`
	Platform                     *string `protobuf:"bytes,2,opt,name=platform" json:"platform,omitempty"`
	Locale                       *string `protobuf:"bytes,3,opt,name=locale" json:"locale,omitempty"`
	Email                        *string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Version                      *string `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	ApplicationVersion           *int32  `protobuf:"varint,6,opt,name=application_version,json=applicationVersion" json:"application_version,omitempty"`
	PublicComputer               *bool   `protobuf:"varint,7,opt,name=public_computer,json=publicComputer" json:"public_computer,omitempty"`
	SsoId                        []byte  `protobuf:"bytes,8,opt,name=sso_id,json=ssoId" json:"sso_id,omitempty"`
	DisconnectOnCookieFail       *bool   `protobuf:"varint,9,opt,name=disconnect_on_cookie_fail,json=disconnectOnCookieFail,def=0" json:"disconnect_on_cookie_fail,omitempty"`
	AllowLogonQueueNotifications *bool   `protobuf:"varint,10,opt,name=allow_logon_queue_notifications,json=allowLogonQueueNotifications,def=0" json:"allow_logon_queue_notifications,omitempty"`
	WebClientVerification        *bool   `protobuf:"varint,11,opt,name=web_client_verification,json=webClientVerification,def=0" json:"web_client_verification,omitempty"`
	CachedWebCredentials         []byte  `protobuf:"bytes,12,opt,name=cached_web_credentials,json=cachedWebCredentials" json:"cached_web_credentials,omitempty"`
	UserAgent                    *string `protobuf:"bytes,14,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	XXX_unrecognized             []byte  `json:"-"`
}

func (m *LogonRequest) Reset()                    { *m = LogonRequest{} }
func (m *LogonRequest) String() string            { return proto.CompactTextString(m) }
func (*LogonRequest) ProtoMessage()               {}
func (*LogonRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

const Default_LogonRequest_DisconnectOnCookieFail bool = false
const Default_LogonRequest_AllowLogonQueueNotifications bool = false
const Default_LogonRequest_WebClientVerification bool = false

func (m *LogonRequest) GetProgram() string {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return ""
}

func (m *LogonRequest) GetPlatform() string {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return ""
}

func (m *LogonRequest) GetLocale() string {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return ""
}

func (m *LogonRequest) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *LogonRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *LogonRequest) GetApplicationVersion() int32 {
	if m != nil && m.ApplicationVersion != nil {
		return *m.ApplicationVersion
	}
	return 0
}

func (m *LogonRequest) GetPublicComputer() bool {
	if m != nil && m.PublicComputer != nil {
		return *m.PublicComputer
	}
	return false
}

func (m *LogonRequest) GetSsoId() []byte {
	if m != nil {
		return m.SsoId
	}
	return nil
}

func (m *LogonRequest) GetDisconnectOnCookieFail() bool {
	if m != nil && m.DisconnectOnCookieFail != nil {
		return *m.DisconnectOnCookieFail
	}
	return Default_LogonRequest_DisconnectOnCookieFail
}

func (m *LogonRequest) GetAllowLogonQueueNotifications() bool {
	if m != nil && m.AllowLogonQueueNotifications != nil {
		return *m.AllowLogonQueueNotifications
	}
	return Default_LogonRequest_AllowLogonQueueNotifications
}

func (m *LogonRequest) GetWebClientVerification() bool {
	if m != nil && m.WebClientVerification != nil {
		return *m.WebClientVerification
	}
	return Default_LogonRequest_WebClientVerification
}

func (m *LogonRequest) GetCachedWebCredentials() []byte {
	if m != nil {
		return m.CachedWebCredentials
	}
	return nil
}

func (m *LogonRequest) GetUserAgent() string {
	if m != nil && m.UserAgent != nil {
		return *m.UserAgent
	}
	return ""
}

// ref: bnet.protocol.authentication.LogonResult
type LogonResult struct {
	ErrorCode        *uint32                   `protobuf:"varint,1,req,name=error_code,json=errorCode" json:"error_code,omitempty"`
	Account          *bnet_protocol.EntityId   `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
	GameAccount      []*bnet_protocol.EntityId `protobuf:"bytes,3,rep,name=game_account,json=gameAccount" json:"game_account,omitempty"`
	Email            *string                   `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	AvailableRegion  []uint32                  `protobuf:"varint,5,rep,name=available_region,json=availableRegion" json:"available_region,omitempty"`
	ConnectedRegion  *uint32                   `protobuf:"varint,6,opt,name=connected_region,json=connectedRegion" json:"connected_region,omitempty"`
	BattleTag        *string                   `protobuf:"bytes,7,opt,name=battle_tag,json=battleTag" json:"battle_tag,omitempty"`
	GeoipCountry     *string                   `protobuf:"bytes,8,opt,name=geoip_country,json=geoipCountry" json:"geoip_country,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *LogonResult) Reset()                    { *m = LogonResult{} }
func (m *LogonResult) String() string            { return proto.CompactTextString(m) }
func (*LogonResult) ProtoMessage()               {}
func (*LogonResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LogonResult) GetErrorCode() uint32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *LogonResult) GetAccount() *bnet_protocol.EntityId {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *LogonResult) GetGameAccount() []*bnet_protocol.EntityId {
	if m != nil {
		return m.GameAccount
	}
	return nil
}

func (m *LogonResult) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *LogonResult) GetAvailableRegion() []uint32 {
	if m != nil {
		return m.AvailableRegion
	}
	return nil
}

func (m *LogonResult) GetConnectedRegion() uint32 {
	if m != nil && m.ConnectedRegion != nil {
		return *m.ConnectedRegion
	}
	return 0
}

func (m *LogonResult) GetBattleTag() string {
	if m != nil && m.BattleTag != nil {
		return *m.BattleTag
	}
	return ""
}

func (m *LogonResult) GetGeoipCountry() string {
	if m != nil && m.GeoipCountry != nil {
		return *m.GeoipCountry
	}
	return ""
}

// ref: bnet.protocol.authentication.LogonUpdateRequest
type LogonUpdateRequest struct {
	ErrorCode        *uint32 `protobuf:"varint,1,req,name=error_code,json=errorCode" json:"error_code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogonUpdateRequest) Reset()                    { *m = LogonUpdateRequest{} }
func (m *LogonUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*LogonUpdateRequest) ProtoMessage()               {}
func (*LogonUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LogonUpdateRequest) GetErrorCode() uint32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

// ref: bnet.protocol.authentication.MemModuleLoadRequest
type MemModuleLoadRequest struct {
	Handle           *bnet_protocol.ContentHandle `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	Key              []byte                       `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Input            []byte                       `protobuf:"bytes,3,req,name=input" json:"input,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *MemModuleLoadRequest) Reset()                    { *m = MemModuleLoadRequest{} }
func (m *MemModuleLoadRequest) String() string            { return proto.CompactTextString(m) }
func (*MemModuleLoadRequest) ProtoMessage()               {}
func (*MemModuleLoadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MemModuleLoadRequest) GetHandle() *bnet_protocol.ContentHandle {
	if m != nil {
		return m.Handle
	}
	return nil
}

func (m *MemModuleLoadRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MemModuleLoadRequest) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

// ref: bnet.protocol.authentication.MemModuleLoadResponse
type MemModuleLoadResponse struct {
	Data             []byte `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MemModuleLoadResponse) Reset()                    { *m = MemModuleLoadResponse{} }
func (m *MemModuleLoadResponse) String() string            { return proto.CompactTextString(m) }
func (*MemModuleLoadResponse) ProtoMessage()               {}
func (*MemModuleLoadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MemModuleLoadResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// ref: bnet.protocol.authentication.ModuleLoadRequest
type ModuleLoadRequest struct {
	ModuleHandle     *bnet_protocol.ContentHandle `protobuf:"bytes,1,req,name=module_handle,json=moduleHandle" json:"module_handle,omitempty"`
	Message          []byte                       `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *ModuleLoadRequest) Reset()                    { *m = ModuleLoadRequest{} }
func (m *ModuleLoadRequest) String() string            { return proto.CompactTextString(m) }
func (*ModuleLoadRequest) ProtoMessage()               {}
func (*ModuleLoadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ModuleLoadRequest) GetModuleHandle() *bnet_protocol.ContentHandle {
	if m != nil {
		return m.ModuleHandle
	}
	return nil
}

func (m *ModuleLoadRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

// ref: bnet.protocol.authentication.ModuleMessageRequest
type ModuleMessageRequest struct {
	ModuleId         *int32 `protobuf:"varint,1,req,name=module_id,json=moduleId" json:"module_id,omitempty"`
	Message          []byte `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ModuleMessageRequest) Reset()                    { *m = ModuleMessageRequest{} }
func (m *ModuleMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*ModuleMessageRequest) ProtoMessage()               {}
func (*ModuleMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ModuleMessageRequest) GetModuleId() int32 {
	if m != nil && m.ModuleId != nil {
		return *m.ModuleId
	}
	return 0
}

func (m *ModuleMessageRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

// ref: bnet.protocol.authentication.ModuleNotification
type ModuleNotification struct {
	ModuleId         *int32  `protobuf:"varint,2,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	Result           *uint32 `protobuf:"varint,3,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ModuleNotification) Reset()                    { *m = ModuleNotification{} }
func (m *ModuleNotification) String() string            { return proto.CompactTextString(m) }
func (*ModuleNotification) ProtoMessage()               {}
func (*ModuleNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ModuleNotification) GetModuleId() int32 {
	if m != nil && m.ModuleId != nil {
		return *m.ModuleId
	}
	return 0
}

func (m *ModuleNotification) GetResult() uint32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

// ref: bnet.protocol.authentication.SelectGameAccountRequest
type SelectGameAccountRequest struct {
	GameAccount      *bnet_protocol.EntityId `protobuf:"bytes,1,req,name=game_account,json=gameAccount" json:"game_account,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *SelectGameAccountRequest) Reset()                    { *m = SelectGameAccountRequest{} }
func (m *SelectGameAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*SelectGameAccountRequest) ProtoMessage()               {}
func (*SelectGameAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SelectGameAccountRequest) GetGameAccount() *bnet_protocol.EntityId {
	if m != nil {
		return m.GameAccount
	}
	return nil
}

// ref: bnet.protocol.authentication.ServerStateChangeRequest
type ServerStateChangeRequest struct {
	State            *uint32 `protobuf:"varint,1,req,name=state" json:"state,omitempty"`
	EventTime        *uint64 `protobuf:"varint,2,req,name=event_time,json=eventTime" json:"event_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServerStateChangeRequest) Reset()                    { *m = ServerStateChangeRequest{} }
func (m *ServerStateChangeRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerStateChangeRequest) ProtoMessage()               {}
func (*ServerStateChangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ServerStateChangeRequest) GetState() uint32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *ServerStateChangeRequest) GetEventTime() uint64 {
	if m != nil && m.EventTime != nil {
		return *m.EventTime
	}
	return 0
}

// ref: bnet.protocol.authentication.VerifyWebCredentialsRequest
type VerifyWebCredentialsRequest struct {
	WebCredentials   []byte `protobuf:"bytes,1,opt,name=web_credentials,json=webCredentials" json:"web_credentials,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *VerifyWebCredentialsRequest) Reset()                    { *m = VerifyWebCredentialsRequest{} }
func (m *VerifyWebCredentialsRequest) String() string            { return proto.CompactTextString(m) }
func (*VerifyWebCredentialsRequest) ProtoMessage()               {}
func (*VerifyWebCredentialsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *VerifyWebCredentialsRequest) GetWebCredentials() []byte {
	if m != nil {
		return m.WebCredentials
	}
	return nil
}

// ref: bnet.protocol.authentication.VersionInfo
type VersionInfo struct {
	Number           *uint32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Patch            *string `protobuf:"bytes,2,opt,name=patch" json:"patch,omitempty"`
	IsOptional       *bool   `protobuf:"varint,3,opt,name=is_optional,json=isOptional" json:"is_optional,omitempty"`
	KickTime         *uint64 `protobuf:"varint,4,opt,name=kick_time,json=kickTime" json:"kick_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *VersionInfo) GetNumber() uint32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *VersionInfo) GetPatch() string {
	if m != nil && m.Patch != nil {
		return *m.Patch
	}
	return ""
}

func (m *VersionInfo) GetIsOptional() bool {
	if m != nil && m.IsOptional != nil {
		return *m.IsOptional
	}
	return false
}

func (m *VersionInfo) GetKickTime() uint64 {
	if m != nil && m.KickTime != nil {
		return *m.KickTime
	}
	return 0
}

// ref: bnet.protocol.authentication.VersionInfoNotification
type VersionInfoNotification struct {
	VersionInfo      *VersionInfo `protobuf:"bytes,1,opt,name=version_info,json=versionInfo" json:"version_info,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *VersionInfoNotification) Reset()                    { *m = VersionInfoNotification{} }
func (m *VersionInfoNotification) String() string            { return proto.CompactTextString(m) }
func (*VersionInfoNotification) ProtoMessage()               {}
func (*VersionInfoNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *VersionInfoNotification) GetVersionInfo() *VersionInfo {
	if m != nil {
		return m.VersionInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountSettingsNotification)(nil), "bnet.protocol.authentication.AccountSettingsNotification")
	proto.RegisterType((*GameAccountSelectedRequest)(nil), "bnet.protocol.authentication.GameAccountSelectedRequest")
	proto.RegisterType((*GenerateSSOTokenRequest)(nil), "bnet.protocol.authentication.GenerateSSOTokenRequest")
	proto.RegisterType((*GenerateSSOTokenResponse)(nil), "bnet.protocol.authentication.GenerateSSOTokenResponse")
	proto.RegisterType((*LogonQueueUpdateRequest)(nil), "bnet.protocol.authentication.LogonQueueUpdateRequest")
	proto.RegisterType((*LogonRequest)(nil), "bnet.protocol.authentication.LogonRequest")
	proto.RegisterType((*LogonResult)(nil), "bnet.protocol.authentication.LogonResult")
	proto.RegisterType((*LogonUpdateRequest)(nil), "bnet.protocol.authentication.LogonUpdateRequest")
	proto.RegisterType((*MemModuleLoadRequest)(nil), "bnet.protocol.authentication.MemModuleLoadRequest")
	proto.RegisterType((*MemModuleLoadResponse)(nil), "bnet.protocol.authentication.MemModuleLoadResponse")
	proto.RegisterType((*ModuleLoadRequest)(nil), "bnet.protocol.authentication.ModuleLoadRequest")
	proto.RegisterType((*ModuleMessageRequest)(nil), "bnet.protocol.authentication.ModuleMessageRequest")
	proto.RegisterType((*ModuleNotification)(nil), "bnet.protocol.authentication.ModuleNotification")
	proto.RegisterType((*SelectGameAccountRequest)(nil), "bnet.protocol.authentication.SelectGameAccountRequest")
	proto.RegisterType((*ServerStateChangeRequest)(nil), "bnet.protocol.authentication.ServerStateChangeRequest")
	proto.RegisterType((*VerifyWebCredentialsRequest)(nil), "bnet.protocol.authentication.VerifyWebCredentialsRequest")
	proto.RegisterType((*VersionInfo)(nil), "bnet.protocol.authentication.VersionInfo")
	proto.RegisterType((*VersionInfoNotification)(nil), "bnet.protocol.authentication.VersionInfoNotification")
}

func init() { proto.RegisterFile("bnet/protocol/authentication/authentication.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0x5b, 0x6f, 0x13, 0xc7,
	0x17, 0x97, 0xed, 0x38, 0x89, 0x8f, 0xed, 0x04, 0x96, 0x40, 0xf6, 0x1f, 0x82, 0xb0, 0x96, 0x3f,
	0xc2, 0xb4, 0x6a, 0x22, 0x02, 0x4f, 0x48, 0x95, 0x1a, 0xb9, 0x85, 0x5a, 0x4a, 0x1a, 0x3a, 0x86,
	0xf4, 0x71, 0x34, 0xde, 0x3d, 0xde, 0x8c, 0xb2, 0x3b, 0xb3, 0xec, 0x8c, 0x1d, 0xe5, 0xb9, 0xcf,
	0xfd, 0x10, 0xfd, 0x3c, 0xfd, 0x52, 0xd5, 0x5c, 0xd6, 0x37, 0x04, 0x4a, 0x9f, 0xbc, 0xe7, 0x37,
	0xbf, 0x39, 0xb7, 0x39, 0x17, 0xc3, 0xab, 0xb1, 0x40, 0x7d, 0x5c, 0x94, 0x52, 0xcb, 0x58, 0x66,
	0xc7, 0x6c, 0xaa, 0xaf, 0x50, 0x68, 0x1e, 0x33, 0xcd, 0xa5, 0x58, 0x13, 0x8f, 0x2c, 0x2d, 0x38,
	0x34, 0x57, 0x8e, 0xaa, 0x2b, 0x47, 0xab, 0x9c, 0x83, 0x67, 0x6b, 0x0a, 0xe3, 0x58, 0x4e, 0x85,
	0xae, 0x7e, 0xdd, 0xb5, 0x83, 0xc3, 0x55, 0xd2, 0x5c, 0x97, 0xfd, 0x88, 0xfe, 0xac, 0xc3, 0xe3,
	0x53, 0xc7, 0x1f, 0xa1, 0xd6, 0x5c, 0xa4, 0xea, 0x37, 0xa9, 0xf9, 0xc4, 0x9b, 0x08, 0x4e, 0x61,
	0x3b, 0xe3, 0x31, 0x0a, 0x85, 0x2a, 0xac, 0xf5, 0x1a, 0xfd, 0xf6, 0xc9, 0xf3, 0xa3, 0x35, 0x9f,
	0xbc, 0x35, 0xaf, 0xe5, 0xcc, 0xb1, 0xc9, 0xfc, 0x5a, 0xd0, 0x83, 0x0e, 0x57, 0x74, 0xaa, 0xb8,
	0x48, 0x69, 0xc9, 0x93, 0xb0, 0xde, 0xab, 0xf5, 0xb7, 0x09, 0x70, 0xf5, 0xc9, 0x40, 0x84, 0x27,
	0xc1, 0x0f, 0xf0, 0x80, 0x2b, 0x5a, 0x64, 0xec, 0xd6, 0x70, 0x26, 0xa5, 0xcc, 0x29, 0x4f, 0xcb,
	0xb0, 0x61, 0x89, 0xf7, 0xb8, 0xfa, 0xe0, 0x4e, 0xde, 0x95, 0x32, 0x1f, 0xa6, 0x65, 0xf0, 0x1d,
	0xdc, 0x8f, 0x99, 0xa0, 0x25, 0xc6, 0xc8, 0x67, 0x48, 0x67, 0x92, 0xc7, 0x18, 0x6e, 0x58, 0xf2,
	0x6e, 0xcc, 0x04, 0x71, 0xf8, 0xa5, 0x81, 0x83, 0xff, 0xc3, 0x8e, 0xe1, 0x2a, 0x14, 0x89, 0x27,
	0x36, 0x2d, 0xb1, 0x13, 0x33, 0x31, 0x42, 0x91, 0x58, 0x56, 0x54, 0xc0, 0xc1, 0x7b, 0x96, 0xe3,
	0x3c, 0x11, 0x19, 0xc6, 0x1a, 0x13, 0x82, 0x9f, 0xa7, 0xa8, 0x74, 0xf0, 0x08, 0x36, 0x4b, 0x54,
	0xd3, 0x4c, 0x87, 0xb5, 0x5e, 0xbd, 0xdf, 0x25, 0x5e, 0x0a, 0xde, 0x42, 0x27, 0x65, 0x39, 0x52,
	0x9f, 0x01, 0x1b, 0x58, 0xfb, 0x64, 0x7f, 0x2d, 0x3f, 0xbf, 0x08, 0xcd, 0xf5, 0xed, 0x30, 0x21,
	0xed, 0x74, 0x61, 0x22, 0x7a, 0x0d, 0xfb, 0xef, 0x51, 0x60, 0xc9, 0x34, 0x8e, 0x46, 0x17, 0x1f,
	0xe5, 0x35, 0x8a, 0xca, 0x5c, 0x08, 0x5b, 0x45, 0x29, 0xd3, 0x92, 0xe5, 0x61, 0xad, 0x57, 0xeb,
	0x6f, 0x91, 0x4a, 0x8c, 0x3e, 0x40, 0xf8, 0xe5, 0x25, 0x55, 0x48, 0xa1, 0x30, 0x78, 0x08, 0x9b,
	0x4a, 0x49, 0xca, 0x13, 0x7b, 0xa9, 0x43, 0x9a, 0x4a, 0xc9, 0x61, 0x12, 0x3c, 0x01, 0x30, 0xb0,
	0xc2, 0xb8, 0x44, 0xe7, 0x61, 0x87, 0xb4, 0x94, 0x92, 0x23, 0x0b, 0x44, 0x7f, 0xd5, 0x60, 0xff,
	0x4c, 0xa6, 0x52, 0xfc, 0x3e, 0xc5, 0x29, 0x7e, 0x2a, 0x12, 0xa6, 0xb1, 0xf2, 0xe3, 0x00, 0xb6,
	0x0b, 0xa9, 0xb8, 0x29, 0x03, 0x1f, 0xf8, 0x5c, 0x0e, 0x9e, 0xc3, 0x0e, 0x2a, 0xcd, 0x73, 0xa6,
	0x31, 0xa1, 0x9a, 0xe7, 0x18, 0xd6, 0x7b, 0xf5, 0xfe, 0x06, 0xe9, 0xce, 0xd1, 0x8f, 0x3c, 0xc7,
	0xe0, 0x18, 0xf6, 0x50, 0x33, 0x9a, 0xe0, 0x8c, 0xdb, 0x72, 0xa2, 0xdc, 0x3c, 0x45, 0x1c, 0x36,
	0x2c, 0xf9, 0x3e, 0x6a, 0xf6, 0x73, 0x75, 0x34, 0x14, 0x23, 0x8c, 0xa3, 0xbf, 0x37, 0xa0, 0x63,
	0xfd, 0xf9, 0x4a, 0x32, 0x5a, 0xf3, 0x64, 0x58, 0xf7, 0x32, 0xa6, 0x27, 0xb2, 0xcc, 0x6d, 0x5c,
	0x2d, 0x32, 0x97, 0xcd, 0x8b, 0x65, 0x32, 0x66, 0x19, 0xda, 0x1a, 0x6a, 0x11, 0x2f, 0x05, 0x7b,
	0xd0, 0xc4, 0x9c, 0xf1, 0xcc, 0x56, 0x4b, 0x8b, 0x38, 0xc1, 0xd8, 0x98, 0x61, 0xa9, 0x4c, 0x9c,
	0x4d, 0x67, 0xc3, 0x8b, 0xc1, 0x31, 0x3c, 0x60, 0x45, 0x91, 0xf9, 0x66, 0xa0, 0x15, 0x6b, 0xb3,
	0x57, 0xeb, 0x37, 0x49, 0xb0, 0x74, 0x74, 0xe9, 0x2f, 0xbc, 0x80, 0xdd, 0x62, 0x3a, 0xce, 0x78,
	0x4c, 0x63, 0x99, 0x17, 0x53, 0x8d, 0x65, 0xb8, 0x65, 0xeb, 0x6d, 0xc7, 0xc1, 0x03, 0x8f, 0x2e,
	0x3d, 0xd7, 0xf6, 0xf2, 0x73, 0xfd, 0x04, 0xff, 0x4b, 0xb8, 0x8a, 0xa5, 0x10, 0x18, 0x6b, 0x2a,
	0x05, 0x8d, 0xa5, 0xbc, 0xe6, 0x48, 0x27, 0xc6, 0xe9, 0x96, 0xd1, 0xf4, 0xb6, 0x39, 0x61, 0x99,
	0x42, 0xf2, 0x68, 0xc1, 0xbb, 0x10, 0x03, 0xcb, 0x7a, 0x67, 0x82, 0x39, 0x83, 0xa7, 0x2c, 0xcb,
	0xe4, 0x0d, 0xcd, 0x4c, 0x1a, 0xe9, 0x67, 0xf3, 0xae, 0x54, 0x2c, 0xb5, 0xb4, 0x0a, 0x61, 0x59,
	0xcf, 0xa1, 0x65, 0x2f, 0x6a, 0x60, 0xb9, 0xfb, 0x55, 0xf0, 0x23, 0xec, 0xdf, 0xe0, 0x98, 0xc6,
	0x19, 0x47, 0xa1, 0x4d, 0xfc, 0xf3, 0xb3, 0xb0, 0xbd, 0xac, 0xe5, 0xe1, 0x0d, 0x8e, 0x07, 0x96,
	0x74, 0xb9, 0xc4, 0x09, 0xde, 0xc0, 0xa3, 0x98, 0xc5, 0x57, 0x98, 0x50, 0xab, 0xa5, 0xc4, 0xc4,
	0x4c, 0x2f, 0x96, 0xa9, 0xb0, 0x63, 0xa3, 0xde, 0x73, 0xa7, 0x7f, 0xe0, 0x78, 0xb0, 0x38, 0x33,
	0x35, 0x3b, 0x55, 0x58, 0x52, 0x96, 0xa2, 0xd0, 0xe1, 0x8e, 0x7d, 0x92, 0x96, 0x41, 0x4e, 0x0d,
	0x10, 0xfd, 0x53, 0x87, 0xb6, 0xaf, 0x11, 0xdb, 0x86, 0x4f, 0x00, 0xb0, 0x2c, 0x65, 0x49, 0x63,
	0x99, 0xa0, 0xaf, 0xd4, 0x96, 0x45, 0x06, 0x32, 0xc1, 0xe0, 0x15, 0x6c, 0xdd, 0xb1, 0x41, 0x2b,
	0xde, 0x17, 0x8d, 0xdd, 0xb0, 0x83, 0xef, 0x4e, 0x8d, 0xfd, 0x95, 0x12, 0x7b, 0x09, 0xf7, 0xd8,
	0x8c, 0xf1, 0x8c, 0x8d, 0x33, 0xa4, 0x25, 0xa6, 0xae, 0xd6, 0x1a, 0xfd, 0x2e, 0xd9, 0x9d, 0xe3,
	0xc4, 0xc2, 0x86, 0xea, 0xdf, 0x15, 0x93, 0x8a, 0x6a, 0x0a, 0xae, 0x4b, 0x76, 0xe7, 0xb8, 0xa7,
	0x3e, 0x01, 0x18, 0x33, 0xad, 0x33, 0xa4, 0x9a, 0xa5, 0xb6, 0xd0, 0x5a, 0xa4, 0xe5, 0x90, 0x8f,
	0x2c, 0x0d, 0x9e, 0x41, 0x37, 0x45, 0xc9, 0x0b, 0x6a, 0x3d, 0x2b, 0x6f, 0x6d, 0xa9, 0xb5, 0x48,
	0xc7, 0x82, 0x03, 0x87, 0x45, 0xaf, 0x21, 0xb0, 0xc9, 0x5c, 0xed, 0xfd, 0x6f, 0xe7, 0x34, 0xd2,
	0xb0, 0x77, 0x8e, 0xf9, 0xb9, 0x4c, 0xa6, 0x19, 0x9e, 0x49, 0x36, 0x9f, 0x94, 0x6f, 0x60, 0xf3,
	0x8a, 0x89, 0x24, 0x73, 0x57, 0xda, 0x27, 0x87, 0x6b, 0x29, 0x1b, 0x48, 0xa1, 0x51, 0xe8, 0x5f,
	0x2d, 0x87, 0x78, 0x6e, 0x70, 0x0f, 0x1a, 0xd7, 0x78, 0x6b, 0x27, 0x48, 0x87, 0x98, 0x4f, 0x93,
	0x44, 0x2e, 0x8a, 0xa9, 0xb6, 0x83, 0xa2, 0x43, 0x9c, 0x10, 0x7d, 0x0f, 0x0f, 0xd7, 0xac, 0xfa,
	0xd9, 0x17, 0xc0, 0x46, 0xc2, 0x34, 0xb3, 0x46, 0x3b, 0xc4, 0x7e, 0x47, 0x05, 0xdc, 0xff, 0xd2,
	0xbf, 0x53, 0xe8, 0xe6, 0x16, 0xa4, 0xff, 0xc1, 0xcd, 0x8e, 0xbb, 0xe2, 0x24, 0x33, 0x2c, 0x72,
	0x54, 0x8a, 0xa5, 0xe8, 0xa7, 0x69, 0x25, 0x46, 0xe7, 0xb0, 0xe7, 0x2c, 0x9e, 0x3b, 0xa0, 0x32,
	0xfa, 0x18, 0x5a, 0xde, 0xa8, 0x1d, 0xce, 0xf5, 0x7e, 0x93, 0x6c, 0x3b, 0x60, 0x98, 0x7c, 0x43,
	0xdd, 0x10, 0x02, 0xa7, 0x6e, 0x65, 0x1f, 0xaf, 0x28, 0xab, 0xdb, 0x39, 0xb4, 0x50, 0xb6, 0x58,
	0x54, 0x0d, 0x5b, 0x30, 0x5e, 0x8a, 0x2e, 0x21, 0x74, 0x3b, 0x6d, 0x69, 0xc9, 0x55, 0xde, 0xad,
	0xd7, 0xba, 0xcb, 0xc8, 0xdd, 0x96, 0xd8, 0x85, 0xd1, 0x5b, 0xce, 0xb0, 0x1c, 0x69, 0xa6, 0x71,
	0x70, 0xc5, 0xc4, 0x22, 0xea, 0x3d, 0x68, 0x2a, 0x83, 0xfa, 0xe2, 0x71, 0x82, 0xad, 0xab, 0x99,
	0x19, 0x25, 0x4b, 0x3b, 0xa3, 0x65, 0x11, 0xb3, 0x2f, 0xa2, 0x77, 0xf0, 0xd8, 0xce, 0x8f, 0xdb,
	0xd5, 0x89, 0x50, 0xe9, 0x7c, 0x01, 0xbb, 0xeb, 0x73, 0xc4, 0x2d, 0xbb, 0x9d, 0x9b, 0x15, 0x7e,
	0x74, 0x0b, 0x6d, 0x3f, 0x91, 0x87, 0x62, 0x22, 0x4d, 0x5e, 0xc4, 0x34, 0x1f, 0x63, 0x69, 0xe9,
	0x5d, 0xe2, 0x25, 0xe3, 0x63, 0xc1, 0x74, 0x7c, 0xe5, 0xf7, 0x87, 0x13, 0x82, 0xa7, 0xd0, 0xe6,
	0x8a, 0xca, 0xc2, 0xe4, 0x9b, 0x65, 0xfe, 0x5f, 0x08, 0x70, 0x75, 0xe1, 0x11, 0xf3, 0x06, 0xd7,
	0x3c, 0xbe, 0x76, 0x31, 0x98, 0x36, 0xdf, 0x20, 0xdb, 0x06, 0xb0, 0x21, 0xa4, 0xb0, 0xbf, 0x64,
	0x7a, 0xe5, 0xed, 0xce, 0xa0, 0xe3, 0x37, 0x08, 0xe5, 0x62, 0x22, 0xad, 0x33, 0xed, 0x93, 0x97,
	0x47, 0xdf, 0xfa, 0x8f, 0x77, 0xb4, 0xa4, 0x8c, 0xb4, 0x67, 0x0b, 0xe1, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x72, 0x1f, 0x22, 0x47, 0x4e, 0x0a, 0x00, 0x00,
}
