// Code generated by protoc-gen-go.
// source: bnet/protocol/attribute_468/attribute_468.proto
// DO NOT EDIT!

package bnet_protocol_attribute_468

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"
import bnet_protocol "github.com/HearthSim/hs-proto-go/bnet/protocol"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// ref: bnet.protocol.attribute.AttributeFilter/Types/Operation
type AttributeFilter_Operation int32

const (
	AttributeFilter_MATCH_NONE              AttributeFilter_Operation = 0
	AttributeFilter_MATCH_ANY               AttributeFilter_Operation = 1
	AttributeFilter_MATCH_ALL               AttributeFilter_Operation = 2
	AttributeFilter_MATCH_ALL_MOST_SPECIFIC AttributeFilter_Operation = 3
)

var AttributeFilter_Operation_name = map[int32]string{
	0: "MATCH_NONE",
	1: "MATCH_ANY",
	2: "MATCH_ALL",
	3: "MATCH_ALL_MOST_SPECIFIC",
}
var AttributeFilter_Operation_value = map[string]int32{
	"MATCH_NONE":              0,
	"MATCH_ANY":               1,
	"MATCH_ALL":               2,
	"MATCH_ALL_MOST_SPECIFIC": 3,
}

func (x AttributeFilter_Operation) Enum() *AttributeFilter_Operation {
	p := new(AttributeFilter_Operation)
	*p = x
	return p
}
func (x AttributeFilter_Operation) String() string {
	return proto.EnumName(AttributeFilter_Operation_name, int32(x))
}
func (x AttributeFilter_Operation) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *AttributeFilter_Operation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AttributeFilter_Operation_value, data, "AttributeFilter_Operation")
	if err != nil {
		return err
	}
	*x = AttributeFilter_Operation(value)
	return nil
}

// ref: bnet.protocol.attribute.AttributeFilter
type AttributeFilter struct {
	Op               *AttributeFilter_Operation `protobuf:"varint,1,req,name=op,enum=bnet.protocol.attribute_468.AttributeFilter_Operation" json:"op,omitempty"`
	Attribute        []*bnet_protocol.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *AttributeFilter) Reset()         { *m = AttributeFilter{} }
func (m *AttributeFilter) String() string { return proto.CompactTextString(m) }
func (*AttributeFilter) ProtoMessage()    {}

func (m *AttributeFilter) GetOp() AttributeFilter_Operation {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return 0
}

func (m *AttributeFilter) GetAttribute() []*bnet_protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func init() {
	proto.RegisterEnum("bnet.protocol.attribute_468.AttributeFilter_Operation", AttributeFilter_Operation_name, AttributeFilter_Operation_value)
}
