// Code generated by protoc-gen-go.
// source: pegasusfsg/pegasusfsg.proto
// DO NOT EDIT!

package pegasusfsg

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"
import pegasusshared "github.com/HearthSim/hs-proto-go/pegasusshared"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// ref: PegasusFSG.CheckInToFSG/PacketID
type CheckInToFSG_PacketID int32

const (
	CheckInToFSG_system CheckInToFSG_PacketID = 3
	CheckInToFSG_ID     CheckInToFSG_PacketID = 502
)

var CheckInToFSG_PacketID_name = map[int32]string{
	3:   "system",
	502: "ID",
}
var CheckInToFSG_PacketID_value = map[string]int32{
	"system": 3,
	"ID":     502,
}

func (x CheckInToFSG_PacketID) Enum() *CheckInToFSG_PacketID {
	p := new(CheckInToFSG_PacketID)
	*p = x
	return p
}
func (x CheckInToFSG_PacketID) String() string {
	return proto.EnumName(CheckInToFSG_PacketID_name, int32(x))
}
func (x CheckInToFSG_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CheckInToFSG_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CheckInToFSG_PacketID_value, data, "CheckInToFSG_PacketID")
	if err != nil {
		return err
	}
	*x = CheckInToFSG_PacketID(value)
	return nil
}

// ref: PegasusFSG.CheckInToFSGResponse/PacketID
type CheckInToFSGResponse_PacketID int32

const (
	CheckInToFSGResponse_ID CheckInToFSGResponse_PacketID = 505
)

var CheckInToFSGResponse_PacketID_name = map[int32]string{
	505: "ID",
}
var CheckInToFSGResponse_PacketID_value = map[string]int32{
	"ID": 505,
}

func (x CheckInToFSGResponse_PacketID) Enum() *CheckInToFSGResponse_PacketID {
	p := new(CheckInToFSGResponse_PacketID)
	*p = x
	return p
}
func (x CheckInToFSGResponse_PacketID) String() string {
	return proto.EnumName(CheckInToFSGResponse_PacketID_name, int32(x))
}
func (x CheckInToFSGResponse_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CheckInToFSGResponse_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CheckInToFSGResponse_PacketID_value, data, "CheckInToFSGResponse_PacketID")
	if err != nil {
		return err
	}
	*x = CheckInToFSGResponse_PacketID(value)
	return nil
}

// ref: PegasusFSG.CheckOutOfFSG/PacketID
type CheckOutOfFSG_PacketID int32

const (
	CheckOutOfFSG_system CheckOutOfFSG_PacketID = 3
	CheckOutOfFSG_ID     CheckOutOfFSG_PacketID = 503
)

var CheckOutOfFSG_PacketID_name = map[int32]string{
	3:   "system",
	503: "ID",
}
var CheckOutOfFSG_PacketID_value = map[string]int32{
	"system": 3,
	"ID":     503,
}

func (x CheckOutOfFSG_PacketID) Enum() *CheckOutOfFSG_PacketID {
	p := new(CheckOutOfFSG_PacketID)
	*p = x
	return p
}
func (x CheckOutOfFSG_PacketID) String() string {
	return proto.EnumName(CheckOutOfFSG_PacketID_name, int32(x))
}
func (x CheckOutOfFSG_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CheckOutOfFSG_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CheckOutOfFSG_PacketID_value, data, "CheckOutOfFSG_PacketID")
	if err != nil {
		return err
	}
	*x = CheckOutOfFSG_PacketID(value)
	return nil
}

// ref: PegasusFSG.CheckOutOfFSGResponse/PacketID
type CheckOutOfFSGResponse_PacketID int32

const (
	CheckOutOfFSGResponse_ID CheckOutOfFSGResponse_PacketID = 506
)

var CheckOutOfFSGResponse_PacketID_name = map[int32]string{
	506: "ID",
}
var CheckOutOfFSGResponse_PacketID_value = map[string]int32{
	"ID": 506,
}

func (x CheckOutOfFSGResponse_PacketID) Enum() *CheckOutOfFSGResponse_PacketID {
	p := new(CheckOutOfFSGResponse_PacketID)
	*p = x
	return p
}
func (x CheckOutOfFSGResponse_PacketID) String() string {
	return proto.EnumName(CheckOutOfFSGResponse_PacketID_name, int32(x))
}
func (x CheckOutOfFSGResponse_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CheckOutOfFSGResponse_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CheckOutOfFSGResponse_PacketID_value, data, "CheckOutOfFSGResponse_PacketID")
	if err != nil {
		return err
	}
	*x = CheckOutOfFSGResponse_PacketID(value)
	return nil
}

// ref: PegasusFSG.DeckValidity/PacketID
type DeckValidity_PacketID int32

const (
	DeckValidity_ID DeckValidity_PacketID = 513
)

var DeckValidity_PacketID_name = map[int32]string{
	513: "ID",
}
var DeckValidity_PacketID_value = map[string]int32{
	"ID": 513,
}

func (x DeckValidity_PacketID) Enum() *DeckValidity_PacketID {
	p := new(DeckValidity_PacketID)
	*p = x
	return p
}
func (x DeckValidity_PacketID) String() string {
	return proto.EnumName(DeckValidity_PacketID_name, int32(x))
}
func (x DeckValidity_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DeckValidity_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeckValidity_PacketID_value, data, "DeckValidity_PacketID")
	if err != nil {
		return err
	}
	*x = DeckValidity_PacketID(value)
	return nil
}

// ref: PegasusFSG.FSGFeatureConfig/PacketID
type FSGFeatureConfig_PacketID int32

const (
	FSGFeatureConfig_ID FSGFeatureConfig_PacketID = 511
)

var FSGFeatureConfig_PacketID_name = map[int32]string{
	511: "ID",
}
var FSGFeatureConfig_PacketID_value = map[string]int32{
	"ID": 511,
}

func (x FSGFeatureConfig_PacketID) Enum() *FSGFeatureConfig_PacketID {
	p := new(FSGFeatureConfig_PacketID)
	*p = x
	return p
}
func (x FSGFeatureConfig_PacketID) String() string {
	return proto.EnumName(FSGFeatureConfig_PacketID_name, int32(x))
}
func (x FSGFeatureConfig_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *FSGFeatureConfig_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FSGFeatureConfig_PacketID_value, data, "FSGFeatureConfig_PacketID")
	if err != nil {
		return err
	}
	*x = FSGFeatureConfig_PacketID(value)
	return nil
}

// ref: PegasusFSG.FSGPatronListUpdate/PacketID
type FSGPatronListUpdate_PacketID int32

const (
	FSGPatronListUpdate_ID FSGPatronListUpdate_PacketID = 512
)

var FSGPatronListUpdate_PacketID_name = map[int32]string{
	512: "ID",
}
var FSGPatronListUpdate_PacketID_value = map[string]int32{
	"ID": 512,
}

func (x FSGPatronListUpdate_PacketID) Enum() *FSGPatronListUpdate_PacketID {
	p := new(FSGPatronListUpdate_PacketID)
	*p = x
	return p
}
func (x FSGPatronListUpdate_PacketID) String() string {
	return proto.EnumName(FSGPatronListUpdate_PacketID_name, int32(x))
}
func (x FSGPatronListUpdate_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *FSGPatronListUpdate_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FSGPatronListUpdate_PacketID_value, data, "FSGPatronListUpdate_PacketID")
	if err != nil {
		return err
	}
	*x = FSGPatronListUpdate_PacketID(value)
	return nil
}

// ref: PegasusFSG.InnkeeperSetupGathering/PacketID
type InnkeeperSetupGathering_PacketID int32

const (
	InnkeeperSetupGathering_system InnkeeperSetupGathering_PacketID = 3
	InnkeeperSetupGathering_ID     InnkeeperSetupGathering_PacketID = 507
)

var InnkeeperSetupGathering_PacketID_name = map[int32]string{
	3:   "system",
	507: "ID",
}
var InnkeeperSetupGathering_PacketID_value = map[string]int32{
	"system": 3,
	"ID":     507,
}

func (x InnkeeperSetupGathering_PacketID) Enum() *InnkeeperSetupGathering_PacketID {
	p := new(InnkeeperSetupGathering_PacketID)
	*p = x
	return p
}
func (x InnkeeperSetupGathering_PacketID) String() string {
	return proto.EnumName(InnkeeperSetupGathering_PacketID_name, int32(x))
}
func (x InnkeeperSetupGathering_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *InnkeeperSetupGathering_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InnkeeperSetupGathering_PacketID_value, data, "InnkeeperSetupGathering_PacketID")
	if err != nil {
		return err
	}
	*x = InnkeeperSetupGathering_PacketID(value)
	return nil
}

// ref: PegasusFSG.InnkeeperSetupGatheringResponse/PacketID
type InnkeeperSetupGatheringResponse_PacketID int32

const (
	InnkeeperSetupGatheringResponse_ID InnkeeperSetupGatheringResponse_PacketID = 508
)

var InnkeeperSetupGatheringResponse_PacketID_name = map[int32]string{
	508: "ID",
}
var InnkeeperSetupGatheringResponse_PacketID_value = map[string]int32{
	"ID": 508,
}

func (x InnkeeperSetupGatheringResponse_PacketID) Enum() *InnkeeperSetupGatheringResponse_PacketID {
	p := new(InnkeeperSetupGatheringResponse_PacketID)
	*p = x
	return p
}
func (x InnkeeperSetupGatheringResponse_PacketID) String() string {
	return proto.EnumName(InnkeeperSetupGatheringResponse_PacketID_name, int32(x))
}
func (x InnkeeperSetupGatheringResponse_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *InnkeeperSetupGatheringResponse_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InnkeeperSetupGatheringResponse_PacketID_value, data, "InnkeeperSetupGatheringResponse_PacketID")
	if err != nil {
		return err
	}
	*x = InnkeeperSetupGatheringResponse_PacketID(value)
	return nil
}

// ref: PegasusFSG.PatronCheckedInToFSG/PacketID
type PatronCheckedInToFSG_PacketID int32

const (
	PatronCheckedInToFSG_system PatronCheckedInToFSG_PacketID = 3
	PatronCheckedInToFSG_ID     PatronCheckedInToFSG_PacketID = 509
)

var PatronCheckedInToFSG_PacketID_name = map[int32]string{
	3:   "system",
	509: "ID",
}
var PatronCheckedInToFSG_PacketID_value = map[string]int32{
	"system": 3,
	"ID":     509,
}

func (x PatronCheckedInToFSG_PacketID) Enum() *PatronCheckedInToFSG_PacketID {
	p := new(PatronCheckedInToFSG_PacketID)
	*p = x
	return p
}
func (x PatronCheckedInToFSG_PacketID) String() string {
	return proto.EnumName(PatronCheckedInToFSG_PacketID_name, int32(x))
}
func (x PatronCheckedInToFSG_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *PatronCheckedInToFSG_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PatronCheckedInToFSG_PacketID_value, data, "PatronCheckedInToFSG_PacketID")
	if err != nil {
		return err
	}
	*x = PatronCheckedInToFSG_PacketID(value)
	return nil
}

// ref: PegasusFSG.PatronCheckedOutOfFSG/PacketID
type PatronCheckedOutOfFSG_PacketID int32

const (
	PatronCheckedOutOfFSG_system PatronCheckedOutOfFSG_PacketID = 3
	PatronCheckedOutOfFSG_ID     PatronCheckedOutOfFSG_PacketID = 510
)

var PatronCheckedOutOfFSG_PacketID_name = map[int32]string{
	3:   "system",
	510: "ID",
}
var PatronCheckedOutOfFSG_PacketID_value = map[string]int32{
	"system": 3,
	"ID":     510,
}

func (x PatronCheckedOutOfFSG_PacketID) Enum() *PatronCheckedOutOfFSG_PacketID {
	p := new(PatronCheckedOutOfFSG_PacketID)
	*p = x
	return p
}
func (x PatronCheckedOutOfFSG_PacketID) String() string {
	return proto.EnumName(PatronCheckedOutOfFSG_PacketID_name, int32(x))
}
func (x PatronCheckedOutOfFSG_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *PatronCheckedOutOfFSG_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PatronCheckedOutOfFSG_PacketID_value, data, "PatronCheckedOutOfFSG_PacketID")
	if err != nil {
		return err
	}
	*x = PatronCheckedOutOfFSG_PacketID(value)
	return nil
}

// ref: PegasusFSG.RequestNearbyFSGs/PacketID
type RequestNearbyFSGs_PacketID int32

const (
	RequestNearbyFSGs_system RequestNearbyFSGs_PacketID = 3
	RequestNearbyFSGs_ID     RequestNearbyFSGs_PacketID = 501
)

var RequestNearbyFSGs_PacketID_name = map[int32]string{
	3:   "system",
	501: "ID",
}
var RequestNearbyFSGs_PacketID_value = map[string]int32{
	"system": 3,
	"ID":     501,
}

func (x RequestNearbyFSGs_PacketID) Enum() *RequestNearbyFSGs_PacketID {
	p := new(RequestNearbyFSGs_PacketID)
	*p = x
	return p
}
func (x RequestNearbyFSGs_PacketID) String() string {
	return proto.EnumName(RequestNearbyFSGs_PacketID_name, int32(x))
}
func (x RequestNearbyFSGs_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *RequestNearbyFSGs_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RequestNearbyFSGs_PacketID_value, data, "RequestNearbyFSGs_PacketID")
	if err != nil {
		return err
	}
	*x = RequestNearbyFSGs_PacketID(value)
	return nil
}

// ref: PegasusFSG.RequestNearbyFSGsResponse/PacketID
type RequestNearbyFSGsResponse_PacketID int32

const (
	RequestNearbyFSGsResponse_ID RequestNearbyFSGsResponse_PacketID = 504
)

var RequestNearbyFSGsResponse_PacketID_name = map[int32]string{
	504: "ID",
}
var RequestNearbyFSGsResponse_PacketID_value = map[string]int32{
	"ID": 504,
}

func (x RequestNearbyFSGsResponse_PacketID) Enum() *RequestNearbyFSGsResponse_PacketID {
	p := new(RequestNearbyFSGsResponse_PacketID)
	*p = x
	return p
}
func (x RequestNearbyFSGsResponse_PacketID) String() string {
	return proto.EnumName(RequestNearbyFSGsResponse_PacketID_name, int32(x))
}
func (x RequestNearbyFSGsResponse_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *RequestNearbyFSGsResponse_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RequestNearbyFSGsResponse_PacketID_value, data, "RequestNearbyFSGsResponse_PacketID")
	if err != nil {
		return err
	}
	*x = RequestNearbyFSGsResponse_PacketID(value)
	return nil
}

// ref: PegasusFSG.CheckInToFSG
type CheckInToFSG struct {
	FsgId            *int64                   `protobuf:"varint,1,req,name=fsg_id" json:"fsg_id,omitempty"`
	Location         *pegasusshared.GPSCoords `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	Bssids           []string                 `protobuf:"bytes,3,rep,name=bssids" json:"bssids,omitempty"`
	Platform         *pegasusshared.Platform  `protobuf:"bytes,4,opt,name=platform" json:"platform,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *CheckInToFSG) Reset()         { *m = CheckInToFSG{} }
func (m *CheckInToFSG) String() string { return proto.CompactTextString(m) }
func (*CheckInToFSG) ProtoMessage()    {}

func (m *CheckInToFSG) GetFsgId() int64 {
	if m != nil && m.FsgId != nil {
		return *m.FsgId
	}
	return 0
}

func (m *CheckInToFSG) GetLocation() *pegasusshared.GPSCoords {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *CheckInToFSG) GetBssids() []string {
	if m != nil {
		return m.Bssids
	}
	return nil
}

func (m *CheckInToFSG) GetPlatform() *pegasusshared.Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

// ref: PegasusFSG.CheckInToFSGResponse
type CheckInToFSGResponse struct {
	ErrorCode          *pegasusshared.ErrorCode               `protobuf:"varint,1,req,name=error_code,enum=pegasusshared.ErrorCode" json:"error_code,omitempty"`
	FsgId              *int64                                 `protobuf:"varint,2,req,name=fsg_id" json:"fsg_id,omitempty"`
	FsgAttendees       []*pegasusshared.FSGPatron             `protobuf:"bytes,3,rep,name=fsg_attendees" json:"fsg_attendees,omitempty"`
	PlayerRecord       *pegasusshared.TavernBrawlPlayerRecord `protobuf:"bytes,4,opt,name=player_record" json:"player_record,omitempty"`
	FsgSharedSecretKey []byte                                 `protobuf:"bytes,5,opt,name=fsg_shared_secret_key" json:"fsg_shared_secret_key,omitempty"`
	XXX_unrecognized   []byte                                 `json:"-"`
}

func (m *CheckInToFSGResponse) Reset()         { *m = CheckInToFSGResponse{} }
func (m *CheckInToFSGResponse) String() string { return proto.CompactTextString(m) }
func (*CheckInToFSGResponse) ProtoMessage()    {}

func (m *CheckInToFSGResponse) GetErrorCode() pegasusshared.ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *CheckInToFSGResponse) GetFsgId() int64 {
	if m != nil && m.FsgId != nil {
		return *m.FsgId
	}
	return 0
}

func (m *CheckInToFSGResponse) GetFsgAttendees() []*pegasusshared.FSGPatron {
	if m != nil {
		return m.FsgAttendees
	}
	return nil
}

func (m *CheckInToFSGResponse) GetPlayerRecord() *pegasusshared.TavernBrawlPlayerRecord {
	if m != nil {
		return m.PlayerRecord
	}
	return nil
}

func (m *CheckInToFSGResponse) GetFsgSharedSecretKey() []byte {
	if m != nil {
		return m.FsgSharedSecretKey
	}
	return nil
}

// ref: PegasusFSG.CheckOutOfFSG
type CheckOutOfFSG struct {
	FsgId            *int64                  `protobuf:"varint,1,req,name=fsg_id" json:"fsg_id,omitempty"`
	Platform         *pegasusshared.Platform `protobuf:"bytes,2,opt,name=platform" json:"platform,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *CheckOutOfFSG) Reset()         { *m = CheckOutOfFSG{} }
func (m *CheckOutOfFSG) String() string { return proto.CompactTextString(m) }
func (*CheckOutOfFSG) ProtoMessage()    {}

func (m *CheckOutOfFSG) GetFsgId() int64 {
	if m != nil && m.FsgId != nil {
		return *m.FsgId
	}
	return 0
}

func (m *CheckOutOfFSG) GetPlatform() *pegasusshared.Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

// ref: PegasusFSG.CheckOutOfFSGResponse
type CheckOutOfFSGResponse struct {
	ErrorCode        *pegasusshared.ErrorCode `protobuf:"varint,1,req,name=error_code,enum=pegasusshared.ErrorCode" json:"error_code,omitempty"`
	FsgId            *int64                   `protobuf:"varint,2,req,name=fsg_id" json:"fsg_id,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *CheckOutOfFSGResponse) Reset()         { *m = CheckOutOfFSGResponse{} }
func (m *CheckOutOfFSGResponse) String() string { return proto.CompactTextString(m) }
func (*CheckOutOfFSGResponse) ProtoMessage()    {}

func (m *CheckOutOfFSGResponse) GetErrorCode() pegasusshared.ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *CheckOutOfFSGResponse) GetFsgId() int64 {
	if m != nil && m.FsgId != nil {
		return *m.FsgId
	}
	return 0
}

// ref: PegasusFSG.DeckValidity
type DeckValidity struct {
	ValidStandardDeck      *bool                              `protobuf:"varint,1,req,name=valid_standard_deck" json:"valid_standard_deck,omitempty"`
	ValidWildDeck          *bool                              `protobuf:"varint,2,req,name=valid_wild_deck" json:"valid_wild_deck,omitempty"`
	ValidTavernBrawlDeck   []*pegasusshared.BrawlDeckValidity `protobuf:"bytes,3,rep,name=valid_tavern_brawl_deck" json:"valid_tavern_brawl_deck,omitempty"`
	ValidFiresideBrawlDeck []*pegasusshared.BrawlDeckValidity `protobuf:"bytes,4,rep,name=valid_fireside_brawl_deck" json:"valid_fireside_brawl_deck,omitempty"`
	XXX_unrecognized       []byte                             `json:"-"`
}

func (m *DeckValidity) Reset()         { *m = DeckValidity{} }
func (m *DeckValidity) String() string { return proto.CompactTextString(m) }
func (*DeckValidity) ProtoMessage()    {}

func (m *DeckValidity) GetValidStandardDeck() bool {
	if m != nil && m.ValidStandardDeck != nil {
		return *m.ValidStandardDeck
	}
	return false
}

func (m *DeckValidity) GetValidWildDeck() bool {
	if m != nil && m.ValidWildDeck != nil {
		return *m.ValidWildDeck
	}
	return false
}

func (m *DeckValidity) GetValidTavernBrawlDeck() []*pegasusshared.BrawlDeckValidity {
	if m != nil {
		return m.ValidTavernBrawlDeck
	}
	return nil
}

func (m *DeckValidity) GetValidFiresideBrawlDeck() []*pegasusshared.BrawlDeckValidity {
	if m != nil {
		return m.ValidFiresideBrawlDeck
	}
	return nil
}

// ref: PegasusFSG.FSGFeatureConfig
type FSGFeatureConfig struct {
	Gps              *bool  `protobuf:"varint,1,opt,name=gps,def=1" json:"gps,omitempty"`
	Wifi             *bool  `protobuf:"varint,2,opt,name=wifi,def=1" json:"wifi,omitempty"`
	AutoCheckin      *bool  `protobuf:"varint,3,opt,name=auto_checkin,def=1" json:"auto_checkin,omitempty"`
	MaxAccuracy      *int64 `protobuf:"varint,4,opt,name=max_accuracy,def=200" json:"max_accuracy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FSGFeatureConfig) Reset()         { *m = FSGFeatureConfig{} }
func (m *FSGFeatureConfig) String() string { return proto.CompactTextString(m) }
func (*FSGFeatureConfig) ProtoMessage()    {}

const Default_FSGFeatureConfig_Gps bool = true
const Default_FSGFeatureConfig_Wifi bool = true
const Default_FSGFeatureConfig_AutoCheckin bool = true
const Default_FSGFeatureConfig_MaxAccuracy int64 = 200

func (m *FSGFeatureConfig) GetGps() bool {
	if m != nil && m.Gps != nil {
		return *m.Gps
	}
	return Default_FSGFeatureConfig_Gps
}

func (m *FSGFeatureConfig) GetWifi() bool {
	if m != nil && m.Wifi != nil {
		return *m.Wifi
	}
	return Default_FSGFeatureConfig_Wifi
}

func (m *FSGFeatureConfig) GetAutoCheckin() bool {
	if m != nil && m.AutoCheckin != nil {
		return *m.AutoCheckin
	}
	return Default_FSGFeatureConfig_AutoCheckin
}

func (m *FSGFeatureConfig) GetMaxAccuracy() int64 {
	if m != nil && m.MaxAccuracy != nil {
		return *m.MaxAccuracy
	}
	return Default_FSGFeatureConfig_MaxAccuracy
}

// ref: PegasusFSG.FSGPatronListUpdate
type FSGPatronListUpdate struct {
	AddedPatrons     []*pegasusshared.FSGPatron `protobuf:"bytes,1,rep,name=added_patrons" json:"added_patrons,omitempty"`
	RemovedPatrons   []*pegasusshared.FSGPatron `protobuf:"bytes,2,rep,name=removed_patrons" json:"removed_patrons,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *FSGPatronListUpdate) Reset()         { *m = FSGPatronListUpdate{} }
func (m *FSGPatronListUpdate) String() string { return proto.CompactTextString(m) }
func (*FSGPatronListUpdate) ProtoMessage()    {}

func (m *FSGPatronListUpdate) GetAddedPatrons() []*pegasusshared.FSGPatron {
	if m != nil {
		return m.AddedPatrons
	}
	return nil
}

func (m *FSGPatronListUpdate) GetRemovedPatrons() []*pegasusshared.FSGPatron {
	if m != nil {
		return m.RemovedPatrons
	}
	return nil
}

// ref: PegasusFSG.InnkeeperSetupGathering
type InnkeeperSetupGathering struct {
	Location         *pegasusshared.GPSCoords `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Bssids           []string                 `protobuf:"bytes,2,rep,name=bssids" json:"bssids,omitempty"`
	FsgId            *int64                   `protobuf:"varint,3,req,name=fsg_id" json:"fsg_id,omitempty"`
	Platform         *pegasusshared.Platform  `protobuf:"bytes,4,opt,name=platform" json:"platform,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *InnkeeperSetupGathering) Reset()         { *m = InnkeeperSetupGathering{} }
func (m *InnkeeperSetupGathering) String() string { return proto.CompactTextString(m) }
func (*InnkeeperSetupGathering) ProtoMessage()    {}

func (m *InnkeeperSetupGathering) GetLocation() *pegasusshared.GPSCoords {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *InnkeeperSetupGathering) GetBssids() []string {
	if m != nil {
		return m.Bssids
	}
	return nil
}

func (m *InnkeeperSetupGathering) GetFsgId() int64 {
	if m != nil && m.FsgId != nil {
		return *m.FsgId
	}
	return 0
}

func (m *InnkeeperSetupGathering) GetPlatform() *pegasusshared.Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

// ref: PegasusFSG.InnkeeperSetupGatheringResponse
type InnkeeperSetupGatheringResponse struct {
	ErrorCode        *pegasusshared.ErrorCode `protobuf:"varint,1,req,name=error_code,enum=pegasusshared.ErrorCode" json:"error_code,omitempty"`
	FsgId            *int64                   `protobuf:"varint,2,req,name=fsg_id" json:"fsg_id,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *InnkeeperSetupGatheringResponse) Reset()         { *m = InnkeeperSetupGatheringResponse{} }
func (m *InnkeeperSetupGatheringResponse) String() string { return proto.CompactTextString(m) }
func (*InnkeeperSetupGatheringResponse) ProtoMessage()    {}

func (m *InnkeeperSetupGatheringResponse) GetErrorCode() pegasusshared.ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *InnkeeperSetupGatheringResponse) GetFsgId() int64 {
	if m != nil && m.FsgId != nil {
		return *m.FsgId
	}
	return 0
}

// ref: PegasusFSG.PatronCheckedInToFSG
type PatronCheckedInToFSG struct {
	Patron           *pegasusshared.FSGPatron `protobuf:"bytes,1,req,name=patron" json:"patron,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *PatronCheckedInToFSG) Reset()         { *m = PatronCheckedInToFSG{} }
func (m *PatronCheckedInToFSG) String() string { return proto.CompactTextString(m) }
func (*PatronCheckedInToFSG) ProtoMessage()    {}

func (m *PatronCheckedInToFSG) GetPatron() *pegasusshared.FSGPatron {
	if m != nil {
		return m.Patron
	}
	return nil
}

// ref: PegasusFSG.PatronCheckedOutOfFSG
type PatronCheckedOutOfFSG struct {
	Patron           *pegasusshared.FSGPatron `protobuf:"bytes,1,req,name=patron" json:"patron,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *PatronCheckedOutOfFSG) Reset()         { *m = PatronCheckedOutOfFSG{} }
func (m *PatronCheckedOutOfFSG) String() string { return proto.CompactTextString(m) }
func (*PatronCheckedOutOfFSG) ProtoMessage()    {}

func (m *PatronCheckedOutOfFSG) GetPatron() *pegasusshared.FSGPatron {
	if m != nil {
		return m.Patron
	}
	return nil
}

// ref: PegasusFSG.RequestNearbyFSGs
type RequestNearbyFSGs struct {
	Location         *pegasusshared.GPSCoords `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Bssids           []string                 `protobuf:"bytes,2,rep,name=bssids" json:"bssids,omitempty"`
	Platform         *pegasusshared.Platform  `protobuf:"bytes,3,opt,name=platform" json:"platform,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *RequestNearbyFSGs) Reset()         { *m = RequestNearbyFSGs{} }
func (m *RequestNearbyFSGs) String() string { return proto.CompactTextString(m) }
func (*RequestNearbyFSGs) ProtoMessage()    {}

func (m *RequestNearbyFSGs) GetLocation() *pegasusshared.GPSCoords {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RequestNearbyFSGs) GetBssids() []string {
	if m != nil {
		return m.Bssids
	}
	return nil
}

func (m *RequestNearbyFSGs) GetPlatform() *pegasusshared.Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

// ref: PegasusFSG.RequestNearbyFSGsResponse
type RequestNearbyFSGsResponse struct {
	ErrorCode          *pegasusshared.ErrorCode   `protobuf:"varint,1,req,name=error_code,enum=pegasusshared.ErrorCode" json:"error_code,omitempty"`
	Fsgs               []*pegasusshared.FSGConfig `protobuf:"bytes,2,rep,name=fsgs" json:"fsgs,omitempty"`
	CheckedInFsgId     *int64                     `protobuf:"varint,3,opt,name=checked_in_fsg_id" json:"checked_in_fsg_id,omitempty"`
	FsgAttendees       []*pegasusshared.FSGPatron `protobuf:"bytes,4,rep,name=fsg_attendees" json:"fsg_attendees,omitempty"`
	FsgSharedSecretKey []byte                     `protobuf:"bytes,5,opt,name=fsg_shared_secret_key" json:"fsg_shared_secret_key,omitempty"`
	XXX_unrecognized   []byte                     `json:"-"`
}

func (m *RequestNearbyFSGsResponse) Reset()         { *m = RequestNearbyFSGsResponse{} }
func (m *RequestNearbyFSGsResponse) String() string { return proto.CompactTextString(m) }
func (*RequestNearbyFSGsResponse) ProtoMessage()    {}

func (m *RequestNearbyFSGsResponse) GetErrorCode() pegasusshared.ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *RequestNearbyFSGsResponse) GetFsgs() []*pegasusshared.FSGConfig {
	if m != nil {
		return m.Fsgs
	}
	return nil
}

func (m *RequestNearbyFSGsResponse) GetCheckedInFsgId() int64 {
	if m != nil && m.CheckedInFsgId != nil {
		return *m.CheckedInFsgId
	}
	return 0
}

func (m *RequestNearbyFSGsResponse) GetFsgAttendees() []*pegasusshared.FSGPatron {
	if m != nil {
		return m.FsgAttendees
	}
	return nil
}

func (m *RequestNearbyFSGsResponse) GetFsgSharedSecretKey() []byte {
	if m != nil {
		return m.FsgSharedSecretKey
	}
	return nil
}

func init() {
	proto.RegisterEnum("pegasusfsg.CheckInToFSG_PacketID", CheckInToFSG_PacketID_name, CheckInToFSG_PacketID_value)
	proto.RegisterEnum("pegasusfsg.CheckInToFSGResponse_PacketID", CheckInToFSGResponse_PacketID_name, CheckInToFSGResponse_PacketID_value)
	proto.RegisterEnum("pegasusfsg.CheckOutOfFSG_PacketID", CheckOutOfFSG_PacketID_name, CheckOutOfFSG_PacketID_value)
	proto.RegisterEnum("pegasusfsg.CheckOutOfFSGResponse_PacketID", CheckOutOfFSGResponse_PacketID_name, CheckOutOfFSGResponse_PacketID_value)
	proto.RegisterEnum("pegasusfsg.DeckValidity_PacketID", DeckValidity_PacketID_name, DeckValidity_PacketID_value)
	proto.RegisterEnum("pegasusfsg.FSGFeatureConfig_PacketID", FSGFeatureConfig_PacketID_name, FSGFeatureConfig_PacketID_value)
	proto.RegisterEnum("pegasusfsg.FSGPatronListUpdate_PacketID", FSGPatronListUpdate_PacketID_name, FSGPatronListUpdate_PacketID_value)
	proto.RegisterEnum("pegasusfsg.InnkeeperSetupGathering_PacketID", InnkeeperSetupGathering_PacketID_name, InnkeeperSetupGathering_PacketID_value)
	proto.RegisterEnum("pegasusfsg.InnkeeperSetupGatheringResponse_PacketID", InnkeeperSetupGatheringResponse_PacketID_name, InnkeeperSetupGatheringResponse_PacketID_value)
	proto.RegisterEnum("pegasusfsg.PatronCheckedInToFSG_PacketID", PatronCheckedInToFSG_PacketID_name, PatronCheckedInToFSG_PacketID_value)
	proto.RegisterEnum("pegasusfsg.PatronCheckedOutOfFSG_PacketID", PatronCheckedOutOfFSG_PacketID_name, PatronCheckedOutOfFSG_PacketID_value)
	proto.RegisterEnum("pegasusfsg.RequestNearbyFSGs_PacketID", RequestNearbyFSGs_PacketID_name, RequestNearbyFSGs_PacketID_value)
	proto.RegisterEnum("pegasusfsg.RequestNearbyFSGsResponse_PacketID", RequestNearbyFSGsResponse_PacketID_name, RequestNearbyFSGsResponse_PacketID_value)
}
