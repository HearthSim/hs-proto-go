// Code generated by protoc-gen-go.
// source: Pegasus/PegasusClient/PegasusClient.proto
// DO NOT EDIT!

/*
Package pegasus_pegasusclient is a generated protocol buffer package.

It is generated from these files:
	Pegasus/PegasusClient/PegasusClient.proto

It has these top-level messages:
	SessionRecord
*/
package pegasus_pegasusclient

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

type SessionRecordType int32

const (
	SessionRecordType_ARENA        SessionRecordType = 0
	SessionRecordType_TAVERN_BRAWL SessionRecordType = 1
)

var SessionRecordType_name = map[int32]string{
	0: "ARENA",
	1: "TAVERN_BRAWL",
}
var SessionRecordType_value = map[string]int32{
	"ARENA":        0,
	"TAVERN_BRAWL": 1,
}

func (x SessionRecordType) Enum() *SessionRecordType {
	p := new(SessionRecordType)
	*p = x
	return p
}
func (x SessionRecordType) String() string {
	return proto.EnumName(SessionRecordType_name, int32(x))
}
func (x *SessionRecordType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SessionRecordType_value, data, "SessionRecordType")
	if err != nil {
		return err
	}
	*x = SessionRecordType(value)
	return nil
}
func (SessionRecordType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SessionRecord struct {
	Wins              *uint32            `protobuf:"varint,1,req,name=wins" json:"wins,omitempty"`
	Losses            *uint32            `protobuf:"varint,2,req,name=losses" json:"losses,omitempty"`
	RunFinished       *bool              `protobuf:"varint,3,req,name=run_finished,json=runFinished" json:"run_finished,omitempty"`
	SessionRecordType *SessionRecordType `protobuf:"varint,4,req,name=session_record_type,json=sessionRecordType,enum=pegasus.pegasusclient.SessionRecordType" json:"session_record_type,omitempty"`
	XXX_unrecognized  []byte             `json:"-"`
}

func (m *SessionRecord) Reset()                    { *m = SessionRecord{} }
func (m *SessionRecord) String() string            { return proto.CompactTextString(m) }
func (*SessionRecord) ProtoMessage()               {}
func (*SessionRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SessionRecord) GetWins() uint32 {
	if m != nil && m.Wins != nil {
		return *m.Wins
	}
	return 0
}

func (m *SessionRecord) GetLosses() uint32 {
	if m != nil && m.Losses != nil {
		return *m.Losses
	}
	return 0
}

func (m *SessionRecord) GetRunFinished() bool {
	if m != nil && m.RunFinished != nil {
		return *m.RunFinished
	}
	return false
}

func (m *SessionRecord) GetSessionRecordType() SessionRecordType {
	if m != nil && m.SessionRecordType != nil {
		return *m.SessionRecordType
	}
	return SessionRecordType_ARENA
}

func init() {
	proto.RegisterType((*SessionRecord)(nil), "pegasus.pegasusclient.SessionRecord")
	proto.RegisterEnum("pegasus.pegasusclient.SessionRecordType", SessionRecordType_name, SessionRecordType_value)
}

func init() { proto.RegisterFile("Pegasus/PegasusClient/PegasusClient.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x0c, 0x48, 0x4d, 0x4f,
	0x2c, 0x2e, 0x2d, 0xd6, 0x87, 0xd2, 0xce, 0x39, 0x99, 0xa9, 0x79, 0x25, 0xa8, 0x3c, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xd1, 0x02, 0x88, 0xa0, 0x1e, 0x94, 0x4e, 0x06, 0x4b, 0x2a, 0xed,
	0x60, 0xe4, 0xe2, 0x0d, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x0b, 0x4a, 0x4d, 0xce, 0x2f, 0x4a,
	0x11, 0x12, 0xe2, 0x62, 0x29, 0xcf, 0xcc, 0x2b, 0x96, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0d, 0x02,
	0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x72, 0xf2, 0x8b, 0x8b, 0x53, 0x8b, 0x25, 0x98, 0xc0, 0xa2, 0x50,
	0x9e, 0x90, 0x22, 0x17, 0x4f, 0x51, 0x69, 0x5e, 0x7c, 0x5a, 0x66, 0x5e, 0x66, 0x71, 0x46, 0x6a,
	0x8a, 0x04, 0xb3, 0x02, 0x93, 0x06, 0x47, 0x10, 0x77, 0x51, 0x69, 0x9e, 0x1b, 0x54, 0x48, 0x28,
	0x82, 0x4b, 0xb8, 0x18, 0x62, 0x7e, 0x7c, 0x11, 0xd8, 0x82, 0xf8, 0x92, 0xca, 0x82, 0x54, 0x09,
	0x16, 0x05, 0x26, 0x0d, 0x3e, 0x23, 0x0d, 0x3d, 0xac, 0xae, 0xd2, 0x43, 0x71, 0x51, 0x48, 0x65,
	0x41, 0x6a, 0x90, 0x60, 0x31, 0xba, 0x90, 0x96, 0x01, 0x97, 0x20, 0x86, 0x3a, 0x21, 0x4e, 0x2e,
	0x56, 0xc7, 0x20, 0x57, 0x3f, 0x47, 0x01, 0x06, 0x21, 0x01, 0x2e, 0x9e, 0x10, 0xc7, 0x30, 0xd7,
	0x20, 0xbf, 0x78, 0xa7, 0x20, 0xc7, 0x70, 0x1f, 0x01, 0x46, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3f, 0x84, 0x35, 0x4f, 0x2f, 0x01, 0x00, 0x00,
}
