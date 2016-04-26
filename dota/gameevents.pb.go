// Code generated by protoc-gen-go.
// source: gameevents.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EBaseGameEvents int32

const (
	EBaseGameEvents_GE_VDebugGameSessionIDEvent            EBaseGameEvents = 200
	EBaseGameEvents_GE_PlaceDecalEvent                     EBaseGameEvents = 201
	EBaseGameEvents_GE_ClearWorldDecalsEvent               EBaseGameEvents = 202
	EBaseGameEvents_GE_ClearEntityDecalsEvent              EBaseGameEvents = 203
	EBaseGameEvents_GE_ClearDecalsForSkeletonInstanceEvent EBaseGameEvents = 204
	EBaseGameEvents_GE_Source1LegacyGameEventList          EBaseGameEvents = 205
	EBaseGameEvents_GE_Source1LegacyListenEvents           EBaseGameEvents = 206
	EBaseGameEvents_GE_Source1LegacyGameEvent              EBaseGameEvents = 207
	EBaseGameEvents_GE_SosStartSoundEvent                  EBaseGameEvents = 208
	EBaseGameEvents_GE_SosStopSoundEvent                   EBaseGameEvents = 209
	EBaseGameEvents_GE_SosSetSoundEventParams              EBaseGameEvents = 210
	EBaseGameEvents_GE_SosSetLibraryStackFields            EBaseGameEvents = 211
	EBaseGameEvents_GE_SosStopSoundEventHash               EBaseGameEvents = 212
)

var EBaseGameEvents_name = map[int32]string{
	200: "GE_VDebugGameSessionIDEvent",
	201: "GE_PlaceDecalEvent",
	202: "GE_ClearWorldDecalsEvent",
	203: "GE_ClearEntityDecalsEvent",
	204: "GE_ClearDecalsForSkeletonInstanceEvent",
	205: "GE_Source1LegacyGameEventList",
	206: "GE_Source1LegacyListenEvents",
	207: "GE_Source1LegacyGameEvent",
	208: "GE_SosStartSoundEvent",
	209: "GE_SosStopSoundEvent",
	210: "GE_SosSetSoundEventParams",
	211: "GE_SosSetLibraryStackFields",
	212: "GE_SosStopSoundEventHash",
}
var EBaseGameEvents_value = map[string]int32{
	"GE_VDebugGameSessionIDEvent":            200,
	"GE_PlaceDecalEvent":                     201,
	"GE_ClearWorldDecalsEvent":               202,
	"GE_ClearEntityDecalsEvent":              203,
	"GE_ClearDecalsForSkeletonInstanceEvent": 204,
	"GE_Source1LegacyGameEventList":          205,
	"GE_Source1LegacyListenEvents":           206,
	"GE_Source1LegacyGameEvent":              207,
	"GE_SosStartSoundEvent":                  208,
	"GE_SosStopSoundEvent":                   209,
	"GE_SosSetSoundEventParams":              210,
	"GE_SosSetLibraryStackFields":            211,
	"GE_SosStopSoundEventHash":               212,
}

func (x EBaseGameEvents) Enum() *EBaseGameEvents {
	p := new(EBaseGameEvents)
	*p = x
	return p
}
func (x EBaseGameEvents) String() string {
	return proto.EnumName(EBaseGameEvents_name, int32(x))
}
func (x *EBaseGameEvents) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EBaseGameEvents_value, data, "EBaseGameEvents")
	if err != nil {
		return err
	}
	*x = EBaseGameEvents(value)
	return nil
}
func (EBaseGameEvents) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

type CMsgVDebugGameSessionIDEvent struct {
	Clientid         *int32  `protobuf:"varint,1,opt,name=clientid" json:"clientid,omitempty"`
	Gamesessionid    *string `protobuf:"bytes,2,opt,name=gamesessionid" json:"gamesessionid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgVDebugGameSessionIDEvent) Reset()                    { *m = CMsgVDebugGameSessionIDEvent{} }
func (m *CMsgVDebugGameSessionIDEvent) String() string            { return proto.CompactTextString(m) }
func (*CMsgVDebugGameSessionIDEvent) ProtoMessage()               {}
func (*CMsgVDebugGameSessionIDEvent) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *CMsgVDebugGameSessionIDEvent) GetClientid() int32 {
	if m != nil && m.Clientid != nil {
		return *m.Clientid
	}
	return 0
}

func (m *CMsgVDebugGameSessionIDEvent) GetGamesessionid() string {
	if m != nil && m.Gamesessionid != nil {
		return *m.Gamesessionid
	}
	return ""
}

type CMsgPlaceDecalEvent struct {
	Position             *CMsgVector `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	Normal               *CMsgVector `protobuf:"bytes,2,opt,name=normal" json:"normal,omitempty"`
	Saxis                *CMsgVector `protobuf:"bytes,3,opt,name=saxis" json:"saxis,omitempty"`
	Decalmaterialindex   *uint32     `protobuf:"varint,4,opt,name=decalmaterialindex" json:"decalmaterialindex,omitempty"`
	Flags                *uint32     `protobuf:"varint,5,opt,name=flags" json:"flags,omitempty"`
	Color                *uint32     `protobuf:"fixed32,6,opt,name=color" json:"color,omitempty"`
	Width                *float32    `protobuf:"fixed32,7,opt,name=width" json:"width,omitempty"`
	Height               *float32    `protobuf:"fixed32,8,opt,name=height" json:"height,omitempty"`
	Depth                *float32    `protobuf:"fixed32,9,opt,name=depth" json:"depth,omitempty"`
	Entityhandleindex    *uint32     `protobuf:"varint,10,opt,name=entityhandleindex" json:"entityhandleindex,omitempty"`
	Skeletoninstancehash *uint32     `protobuf:"fixed32,11,opt,name=skeletoninstancehash" json:"skeletoninstancehash,omitempty"`
	Boneindex            *int32      `protobuf:"varint,12,opt,name=boneindex" json:"boneindex,omitempty"`
	Translucenthit       *bool       `protobuf:"varint,13,opt,name=translucenthit" json:"translucenthit,omitempty"`
	XXX_unrecognized     []byte      `json:"-"`
}

func (m *CMsgPlaceDecalEvent) Reset()                    { *m = CMsgPlaceDecalEvent{} }
func (m *CMsgPlaceDecalEvent) String() string            { return proto.CompactTextString(m) }
func (*CMsgPlaceDecalEvent) ProtoMessage()               {}
func (*CMsgPlaceDecalEvent) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

func (m *CMsgPlaceDecalEvent) GetPosition() *CMsgVector {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *CMsgPlaceDecalEvent) GetNormal() *CMsgVector {
	if m != nil {
		return m.Normal
	}
	return nil
}

func (m *CMsgPlaceDecalEvent) GetSaxis() *CMsgVector {
	if m != nil {
		return m.Saxis
	}
	return nil
}

func (m *CMsgPlaceDecalEvent) GetDecalmaterialindex() uint32 {
	if m != nil && m.Decalmaterialindex != nil {
		return *m.Decalmaterialindex
	}
	return 0
}

func (m *CMsgPlaceDecalEvent) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *CMsgPlaceDecalEvent) GetColor() uint32 {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return 0
}

func (m *CMsgPlaceDecalEvent) GetWidth() float32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *CMsgPlaceDecalEvent) GetHeight() float32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *CMsgPlaceDecalEvent) GetDepth() float32 {
	if m != nil && m.Depth != nil {
		return *m.Depth
	}
	return 0
}

func (m *CMsgPlaceDecalEvent) GetEntityhandleindex() uint32 {
	if m != nil && m.Entityhandleindex != nil {
		return *m.Entityhandleindex
	}
	return 0
}

func (m *CMsgPlaceDecalEvent) GetSkeletoninstancehash() uint32 {
	if m != nil && m.Skeletoninstancehash != nil {
		return *m.Skeletoninstancehash
	}
	return 0
}

func (m *CMsgPlaceDecalEvent) GetBoneindex() int32 {
	if m != nil && m.Boneindex != nil {
		return *m.Boneindex
	}
	return 0
}

func (m *CMsgPlaceDecalEvent) GetTranslucenthit() bool {
	if m != nil && m.Translucenthit != nil {
		return *m.Translucenthit
	}
	return false
}

type CMsgClearWorldDecalsEvent struct {
	Flagstoclear     *uint32 `protobuf:"varint,1,opt,name=flagstoclear" json:"flagstoclear,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgClearWorldDecalsEvent) Reset()                    { *m = CMsgClearWorldDecalsEvent{} }
func (m *CMsgClearWorldDecalsEvent) String() string            { return proto.CompactTextString(m) }
func (*CMsgClearWorldDecalsEvent) ProtoMessage()               {}
func (*CMsgClearWorldDecalsEvent) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func (m *CMsgClearWorldDecalsEvent) GetFlagstoclear() uint32 {
	if m != nil && m.Flagstoclear != nil {
		return *m.Flagstoclear
	}
	return 0
}

type CMsgClearEntityDecalsEvent struct {
	Flagstoclear     *uint32 `protobuf:"varint,1,opt,name=flagstoclear" json:"flagstoclear,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgClearEntityDecalsEvent) Reset()                    { *m = CMsgClearEntityDecalsEvent{} }
func (m *CMsgClearEntityDecalsEvent) String() string            { return proto.CompactTextString(m) }
func (*CMsgClearEntityDecalsEvent) ProtoMessage()               {}
func (*CMsgClearEntityDecalsEvent) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

func (m *CMsgClearEntityDecalsEvent) GetFlagstoclear() uint32 {
	if m != nil && m.Flagstoclear != nil {
		return *m.Flagstoclear
	}
	return 0
}

type CMsgClearDecalsForSkeletonInstanceEvent struct {
	Flagstoclear         *uint32 `protobuf:"varint,1,opt,name=flagstoclear" json:"flagstoclear,omitempty"`
	Entityhandleindex    *uint32 `protobuf:"varint,2,opt,name=entityhandleindex" json:"entityhandleindex,omitempty"`
	Skeletoninstancehash *uint32 `protobuf:"varint,3,opt,name=skeletoninstancehash" json:"skeletoninstancehash,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *CMsgClearDecalsForSkeletonInstanceEvent) Reset() {
	*m = CMsgClearDecalsForSkeletonInstanceEvent{}
}
func (m *CMsgClearDecalsForSkeletonInstanceEvent) String() string { return proto.CompactTextString(m) }
func (*CMsgClearDecalsForSkeletonInstanceEvent) ProtoMessage()    {}
func (*CMsgClearDecalsForSkeletonInstanceEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor16, []int{4}
}

func (m *CMsgClearDecalsForSkeletonInstanceEvent) GetFlagstoclear() uint32 {
	if m != nil && m.Flagstoclear != nil {
		return *m.Flagstoclear
	}
	return 0
}

func (m *CMsgClearDecalsForSkeletonInstanceEvent) GetEntityhandleindex() uint32 {
	if m != nil && m.Entityhandleindex != nil {
		return *m.Entityhandleindex
	}
	return 0
}

func (m *CMsgClearDecalsForSkeletonInstanceEvent) GetSkeletoninstancehash() uint32 {
	if m != nil && m.Skeletoninstancehash != nil {
		return *m.Skeletoninstancehash
	}
	return 0
}

type CMsgSource1LegacyGameEventList struct {
	Descriptors      []*CMsgSource1LegacyGameEventListDescriptorT `protobuf:"bytes,1,rep,name=descriptors" json:"descriptors,omitempty"`
	XXX_unrecognized []byte                                       `json:"-"`
}

func (m *CMsgSource1LegacyGameEventList) Reset()                    { *m = CMsgSource1LegacyGameEventList{} }
func (m *CMsgSource1LegacyGameEventList) String() string            { return proto.CompactTextString(m) }
func (*CMsgSource1LegacyGameEventList) ProtoMessage()               {}
func (*CMsgSource1LegacyGameEventList) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{5} }

func (m *CMsgSource1LegacyGameEventList) GetDescriptors() []*CMsgSource1LegacyGameEventListDescriptorT {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

type CMsgSource1LegacyGameEventListKeyT struct {
	Type             *int32  `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgSource1LegacyGameEventListKeyT) Reset()         { *m = CMsgSource1LegacyGameEventListKeyT{} }
func (m *CMsgSource1LegacyGameEventListKeyT) String() string { return proto.CompactTextString(m) }
func (*CMsgSource1LegacyGameEventListKeyT) ProtoMessage()    {}
func (*CMsgSource1LegacyGameEventListKeyT) Descriptor() ([]byte, []int) {
	return fileDescriptor16, []int{5, 0}
}

func (m *CMsgSource1LegacyGameEventListKeyT) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *CMsgSource1LegacyGameEventListKeyT) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type CMsgSource1LegacyGameEventListDescriptorT struct {
	Eventid          *int32                                `protobuf:"varint,1,opt,name=eventid" json:"eventid,omitempty"`
	Name             *string                               `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Keys             []*CMsgSource1LegacyGameEventListKeyT `protobuf:"bytes,3,rep,name=keys" json:"keys,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *CMsgSource1LegacyGameEventListDescriptorT) Reset() {
	*m = CMsgSource1LegacyGameEventListDescriptorT{}
}
func (m *CMsgSource1LegacyGameEventListDescriptorT) String() string { return proto.CompactTextString(m) }
func (*CMsgSource1LegacyGameEventListDescriptorT) ProtoMessage()    {}
func (*CMsgSource1LegacyGameEventListDescriptorT) Descriptor() ([]byte, []int) {
	return fileDescriptor16, []int{5, 1}
}

func (m *CMsgSource1LegacyGameEventListDescriptorT) GetEventid() int32 {
	if m != nil && m.Eventid != nil {
		return *m.Eventid
	}
	return 0
}

func (m *CMsgSource1LegacyGameEventListDescriptorT) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CMsgSource1LegacyGameEventListDescriptorT) GetKeys() []*CMsgSource1LegacyGameEventListKeyT {
	if m != nil {
		return m.Keys
	}
	return nil
}

type CMsgSource1LegacyListenEvents struct {
	Playerslot       *int32   `protobuf:"varint,1,opt,name=playerslot" json:"playerslot,omitempty"`
	Eventarraybits   []uint32 `protobuf:"varint,2,rep,name=eventarraybits" json:"eventarraybits,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CMsgSource1LegacyListenEvents) Reset()                    { *m = CMsgSource1LegacyListenEvents{} }
func (m *CMsgSource1LegacyListenEvents) String() string            { return proto.CompactTextString(m) }
func (*CMsgSource1LegacyListenEvents) ProtoMessage()               {}
func (*CMsgSource1LegacyListenEvents) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{6} }

func (m *CMsgSource1LegacyListenEvents) GetPlayerslot() int32 {
	if m != nil && m.Playerslot != nil {
		return *m.Playerslot
	}
	return 0
}

func (m *CMsgSource1LegacyListenEvents) GetEventarraybits() []uint32 {
	if m != nil {
		return m.Eventarraybits
	}
	return nil
}

type CMsgSource1LegacyGameEvent struct {
	EventName        *string                           `protobuf:"bytes,1,opt,name=event_name" json:"event_name,omitempty"`
	Eventid          *int32                            `protobuf:"varint,2,opt,name=eventid" json:"eventid,omitempty"`
	Keys             []*CMsgSource1LegacyGameEventKeyT `protobuf:"bytes,3,rep,name=keys" json:"keys,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *CMsgSource1LegacyGameEvent) Reset()                    { *m = CMsgSource1LegacyGameEvent{} }
func (m *CMsgSource1LegacyGameEvent) String() string            { return proto.CompactTextString(m) }
func (*CMsgSource1LegacyGameEvent) ProtoMessage()               {}
func (*CMsgSource1LegacyGameEvent) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{7} }

func (m *CMsgSource1LegacyGameEvent) GetEventName() string {
	if m != nil && m.EventName != nil {
		return *m.EventName
	}
	return ""
}

func (m *CMsgSource1LegacyGameEvent) GetEventid() int32 {
	if m != nil && m.Eventid != nil {
		return *m.Eventid
	}
	return 0
}

func (m *CMsgSource1LegacyGameEvent) GetKeys() []*CMsgSource1LegacyGameEventKeyT {
	if m != nil {
		return m.Keys
	}
	return nil
}

type CMsgSource1LegacyGameEventKeyT struct {
	Type             *int32   `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	ValString        *string  `protobuf:"bytes,2,opt,name=val_string" json:"val_string,omitempty"`
	ValFloat         *float32 `protobuf:"fixed32,3,opt,name=val_float" json:"val_float,omitempty"`
	ValLong          *int32   `protobuf:"varint,4,opt,name=val_long" json:"val_long,omitempty"`
	ValShort         *int32   `protobuf:"varint,5,opt,name=val_short" json:"val_short,omitempty"`
	ValByte          *int32   `protobuf:"varint,6,opt,name=val_byte" json:"val_byte,omitempty"`
	ValBool          *bool    `protobuf:"varint,7,opt,name=val_bool" json:"val_bool,omitempty"`
	ValUint64        *uint64  `protobuf:"varint,8,opt,name=val_uint64" json:"val_uint64,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CMsgSource1LegacyGameEventKeyT) Reset()         { *m = CMsgSource1LegacyGameEventKeyT{} }
func (m *CMsgSource1LegacyGameEventKeyT) String() string { return proto.CompactTextString(m) }
func (*CMsgSource1LegacyGameEventKeyT) ProtoMessage()    {}
func (*CMsgSource1LegacyGameEventKeyT) Descriptor() ([]byte, []int) {
	return fileDescriptor16, []int{7, 0}
}

func (m *CMsgSource1LegacyGameEventKeyT) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *CMsgSource1LegacyGameEventKeyT) GetValString() string {
	if m != nil && m.ValString != nil {
		return *m.ValString
	}
	return ""
}

func (m *CMsgSource1LegacyGameEventKeyT) GetValFloat() float32 {
	if m != nil && m.ValFloat != nil {
		return *m.ValFloat
	}
	return 0
}

func (m *CMsgSource1LegacyGameEventKeyT) GetValLong() int32 {
	if m != nil && m.ValLong != nil {
		return *m.ValLong
	}
	return 0
}

func (m *CMsgSource1LegacyGameEventKeyT) GetValShort() int32 {
	if m != nil && m.ValShort != nil {
		return *m.ValShort
	}
	return 0
}

func (m *CMsgSource1LegacyGameEventKeyT) GetValByte() int32 {
	if m != nil && m.ValByte != nil {
		return *m.ValByte
	}
	return 0
}

func (m *CMsgSource1LegacyGameEventKeyT) GetValBool() bool {
	if m != nil && m.ValBool != nil {
		return *m.ValBool
	}
	return false
}

func (m *CMsgSource1LegacyGameEventKeyT) GetValUint64() uint64 {
	if m != nil && m.ValUint64 != nil {
		return *m.ValUint64
	}
	return 0
}

type CMsgSosStartSoundEvent struct {
	SoundeventGuid    *int32  `protobuf:"varint,1,opt,name=soundevent_guid" json:"soundevent_guid,omitempty"`
	SoundeventHash    *uint32 `protobuf:"fixed32,2,opt,name=soundevent_hash" json:"soundevent_hash,omitempty"`
	SourceEntityIndex *int32  `protobuf:"varint,3,opt,name=source_entity_index" json:"source_entity_index,omitempty"`
	Seed              *int32  `protobuf:"varint,4,opt,name=seed" json:"seed,omitempty"`
	PackedParams      []byte  `protobuf:"bytes,5,opt,name=packed_params" json:"packed_params,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *CMsgSosStartSoundEvent) Reset()                    { *m = CMsgSosStartSoundEvent{} }
func (m *CMsgSosStartSoundEvent) String() string            { return proto.CompactTextString(m) }
func (*CMsgSosStartSoundEvent) ProtoMessage()               {}
func (*CMsgSosStartSoundEvent) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{8} }

func (m *CMsgSosStartSoundEvent) GetSoundeventGuid() int32 {
	if m != nil && m.SoundeventGuid != nil {
		return *m.SoundeventGuid
	}
	return 0
}

func (m *CMsgSosStartSoundEvent) GetSoundeventHash() uint32 {
	if m != nil && m.SoundeventHash != nil {
		return *m.SoundeventHash
	}
	return 0
}

func (m *CMsgSosStartSoundEvent) GetSourceEntityIndex() int32 {
	if m != nil && m.SourceEntityIndex != nil {
		return *m.SourceEntityIndex
	}
	return 0
}

func (m *CMsgSosStartSoundEvent) GetSeed() int32 {
	if m != nil && m.Seed != nil {
		return *m.Seed
	}
	return 0
}

func (m *CMsgSosStartSoundEvent) GetPackedParams() []byte {
	if m != nil {
		return m.PackedParams
	}
	return nil
}

type CMsgSosStopSoundEvent struct {
	SoundeventGuid   *int32 `protobuf:"varint,1,opt,name=soundevent_guid" json:"soundevent_guid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMsgSosStopSoundEvent) Reset()                    { *m = CMsgSosStopSoundEvent{} }
func (m *CMsgSosStopSoundEvent) String() string            { return proto.CompactTextString(m) }
func (*CMsgSosStopSoundEvent) ProtoMessage()               {}
func (*CMsgSosStopSoundEvent) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{9} }

func (m *CMsgSosStopSoundEvent) GetSoundeventGuid() int32 {
	if m != nil && m.SoundeventGuid != nil {
		return *m.SoundeventGuid
	}
	return 0
}

type CMsgSosStopSoundEventHash struct {
	SoundeventHash    *uint32 `protobuf:"fixed32,1,opt,name=soundevent_hash" json:"soundevent_hash,omitempty"`
	SourceEntityIndex *int32  `protobuf:"varint,2,opt,name=source_entity_index" json:"source_entity_index,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *CMsgSosStopSoundEventHash) Reset()                    { *m = CMsgSosStopSoundEventHash{} }
func (m *CMsgSosStopSoundEventHash) String() string            { return proto.CompactTextString(m) }
func (*CMsgSosStopSoundEventHash) ProtoMessage()               {}
func (*CMsgSosStopSoundEventHash) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{10} }

func (m *CMsgSosStopSoundEventHash) GetSoundeventHash() uint32 {
	if m != nil && m.SoundeventHash != nil {
		return *m.SoundeventHash
	}
	return 0
}

func (m *CMsgSosStopSoundEventHash) GetSourceEntityIndex() int32 {
	if m != nil && m.SourceEntityIndex != nil {
		return *m.SourceEntityIndex
	}
	return 0
}

type CMsgSosSetSoundEventParams struct {
	SoundeventGuid   *int32 `protobuf:"varint,1,opt,name=soundevent_guid" json:"soundevent_guid,omitempty"`
	PackedParams     []byte `protobuf:"bytes,5,opt,name=packed_params" json:"packed_params,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMsgSosSetSoundEventParams) Reset()                    { *m = CMsgSosSetSoundEventParams{} }
func (m *CMsgSosSetSoundEventParams) String() string            { return proto.CompactTextString(m) }
func (*CMsgSosSetSoundEventParams) ProtoMessage()               {}
func (*CMsgSosSetSoundEventParams) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{11} }

func (m *CMsgSosSetSoundEventParams) GetSoundeventGuid() int32 {
	if m != nil && m.SoundeventGuid != nil {
		return *m.SoundeventGuid
	}
	return 0
}

func (m *CMsgSosSetSoundEventParams) GetPackedParams() []byte {
	if m != nil {
		return m.PackedParams
	}
	return nil
}

type CMsgSosSetLibraryStackFields struct {
	StackHash        *uint32 `protobuf:"fixed32,1,opt,name=stack_hash" json:"stack_hash,omitempty"`
	PackedFields     []byte  `protobuf:"bytes,5,opt,name=packed_fields" json:"packed_fields,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgSosSetLibraryStackFields) Reset()                    { *m = CMsgSosSetLibraryStackFields{} }
func (m *CMsgSosSetLibraryStackFields) String() string            { return proto.CompactTextString(m) }
func (*CMsgSosSetLibraryStackFields) ProtoMessage()               {}
func (*CMsgSosSetLibraryStackFields) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{12} }

func (m *CMsgSosSetLibraryStackFields) GetStackHash() uint32 {
	if m != nil && m.StackHash != nil {
		return *m.StackHash
	}
	return 0
}

func (m *CMsgSosSetLibraryStackFields) GetPackedFields() []byte {
	if m != nil {
		return m.PackedFields
	}
	return nil
}

func init() {
	proto.RegisterType((*CMsgVDebugGameSessionIDEvent)(nil), "dota.CMsgVDebugGameSessionIDEvent")
	proto.RegisterType((*CMsgPlaceDecalEvent)(nil), "dota.CMsgPlaceDecalEvent")
	proto.RegisterType((*CMsgClearWorldDecalsEvent)(nil), "dota.CMsgClearWorldDecalsEvent")
	proto.RegisterType((*CMsgClearEntityDecalsEvent)(nil), "dota.CMsgClearEntityDecalsEvent")
	proto.RegisterType((*CMsgClearDecalsForSkeletonInstanceEvent)(nil), "dota.CMsgClearDecalsForSkeletonInstanceEvent")
	proto.RegisterType((*CMsgSource1LegacyGameEventList)(nil), "dota.CMsgSource1LegacyGameEventList")
	proto.RegisterType((*CMsgSource1LegacyGameEventListKeyT)(nil), "dota.CMsgSource1LegacyGameEventList.key_t")
	proto.RegisterType((*CMsgSource1LegacyGameEventListDescriptorT)(nil), "dota.CMsgSource1LegacyGameEventList.descriptor_t")
	proto.RegisterType((*CMsgSource1LegacyListenEvents)(nil), "dota.CMsgSource1LegacyListenEvents")
	proto.RegisterType((*CMsgSource1LegacyGameEvent)(nil), "dota.CMsgSource1LegacyGameEvent")
	proto.RegisterType((*CMsgSource1LegacyGameEventKeyT)(nil), "dota.CMsgSource1LegacyGameEvent.key_t")
	proto.RegisterType((*CMsgSosStartSoundEvent)(nil), "dota.CMsgSosStartSoundEvent")
	proto.RegisterType((*CMsgSosStopSoundEvent)(nil), "dota.CMsgSosStopSoundEvent")
	proto.RegisterType((*CMsgSosStopSoundEventHash)(nil), "dota.CMsgSosStopSoundEventHash")
	proto.RegisterType((*CMsgSosSetSoundEventParams)(nil), "dota.CMsgSosSetSoundEventParams")
	proto.RegisterType((*CMsgSosSetLibraryStackFields)(nil), "dota.CMsgSosSetLibraryStackFields")
	proto.RegisterEnum("dota.EBaseGameEvents", EBaseGameEvents_name, EBaseGameEvents_value)
}

var fileDescriptor16 = []byte{
	// 930 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xcb, 0x6e, 0x1b, 0x37,
	0x17, 0xfe, 0xa9, 0x8b, 0x2d, 0x1f, 0x4b, 0xf1, 0x64, 0x62, 0xfb, 0x1f, 0x2b, 0x76, 0x9a, 0x4e,
	0x81, 0xb4, 0x48, 0x01, 0xa3, 0x0e, 0x82, 0xa2, 0xeb, 0xc4, 0x97, 0x18, 0x75, 0x81, 0x14, 0x02,
	0xda, 0xa5, 0x40, 0xcd, 0xd0, 0xd2, 0x40, 0xf4, 0x50, 0x20, 0xa9, 0xc4, 0xda, 0xf5, 0x01, 0xba,
	0xeb, 0xaa, 0x4f, 0xd2, 0x57, 0x48, 0xef, 0xd7, 0x67, 0xe9, 0xba, 0x87, 0x87, 0x33, 0x92, 0x6c,
	0x4b, 0x4e, 0x76, 0xc3, 0x39, 0xb7, 0xef, 0x7c, 0xe7, 0x3b, 0x24, 0x04, 0x7d, 0x7e, 0x21, 0xc4,
	0x2b, 0x91, 0x5b, 0xb3, 0x3f, 0xd2, 0xca, 0xaa, 0xb0, 0x96, 0x2a, 0xcb, 0xdb, 0xdb, 0xb9, 0xb0,
	0xaf, 0x95, 0x1e, 0xf6, 0xb8, 0x11, 0x76, 0x32, 0x12, 0x85, 0x35, 0x3e, 0x81, 0xdd, 0xe7, 0x5f,
	0x98, 0xfe, 0x57, 0x87, 0xa2, 0x37, 0xee, 0x9f, 0x60, 0x6c, 0x47, 0x18, 0x93, 0xa9, 0xfc, 0xf4,
	0xf0, 0xc8, 0x25, 0x09, 0x03, 0x68, 0x24, 0x32, 0xc3, 0xaf, 0x2c, 0x8d, 0xd8, 0x43, 0xf6, 0x51,
	0x3d, 0xdc, 0x82, 0x96, 0xab, 0x61, 0xbc, 0x1f, 0xfe, 0xae, 0xe0, 0xef, 0xb5, 0xf8, 0x4d, 0x05,
	0xee, 0xb9, 0x4c, 0x2f, 0x25, 0x4f, 0xc4, 0xa1, 0x48, 0xb8, 0xf4, 0x09, 0x62, 0x68, 0x8c, 0x94,
	0xc9, 0x2c, 0xfa, 0x52, 0x82, 0xf5, 0x27, 0xc1, 0xbe, 0x43, 0xb4, 0x4f, 0x65, 0x45, 0x62, 0x95,
	0x0e, 0x1f, 0xc2, 0x4a, 0xae, 0xf4, 0x05, 0x97, 0x94, 0x6b, 0x91, 0xc7, 0x7b, 0x50, 0x37, 0xfc,
	0x32, 0x33, 0x51, 0x75, 0x89, 0x43, 0x1b, 0xc2, 0xd4, 0x15, 0xbd, 0xe0, 0x56, 0xe8, 0x8c, 0xcb,
	0x2c, 0x4f, 0xc5, 0x65, 0x54, 0x43, 0xef, 0x56, 0xd8, 0x82, 0xfa, 0xb9, 0xe4, 0x7d, 0x13, 0xd5,
	0xcb, 0x63, 0xa2, 0xa4, 0xd2, 0xd1, 0x0a, 0x1e, 0x57, 0xdd, 0xf1, 0x75, 0x96, 0xda, 0x41, 0xb4,
	0x8a, 0xc7, 0x4a, 0x78, 0x07, 0x56, 0x06, 0x22, 0xeb, 0x0f, 0x6c, 0xd4, 0xa0, 0x33, 0x9a, 0x53,
	0x31, 0x42, 0xf3, 0x1a, 0x1d, 0x77, 0xe0, 0xae, 0x23, 0xc3, 0x4e, 0x06, 0x3c, 0x4f, 0xa5, 0xf0,
	0x65, 0x80, 0xf2, 0xee, 0xc2, 0xa6, 0x19, 0x0a, 0x29, 0x2c, 0xb2, 0x92, 0x1b, 0xcb, 0xf3, 0x44,
	0x0c, 0xb8, 0x19, 0x44, 0xeb, 0x54, 0xe6, 0x2e, 0xac, 0xf5, 0x54, 0x5e, 0x04, 0x34, 0x89, 0xc9,
	0x6d, 0xb8, 0x63, 0x35, 0xcf, 0x8d, 0x1c, 0x27, 0x98, 0x73, 0x90, 0xd9, 0xa8, 0x85, 0xff, 0x1b,
	0xf1, 0x01, 0xec, 0xb8, 0xce, 0x9e, 0x4b, 0xc1, 0xf5, 0xd7, 0x4a, 0xcb, 0x94, 0xe8, 0x34, 0x9e,
	0xcf, 0x4d, 0x68, 0x52, 0x33, 0x56, 0x25, 0xce, 0x4e, 0x9c, 0xb6, 0xe2, 0x27, 0xd0, 0x9e, 0x86,
	0x1c, 0x11, 0xbe, 0xb7, 0xc7, 0x5c, 0xc2, 0x87, 0xd3, 0x18, 0xef, 0x7d, 0xac, 0x74, 0xa7, 0xe8,
	0xe0, 0xb4, 0xe8, 0xe0, 0x96, 0x04, 0x8b, 0xb9, 0xa8, 0xdc, 0xca, 0x45, 0x95, 0x2a, 0xff, 0xcb,
	0xe0, 0x81, 0x2b, 0xdd, 0x51, 0x63, 0x9d, 0x88, 0x83, 0x33, 0xd1, 0xe7, 0xc9, 0xc4, 0x89, 0x8f,
	0xaa, 0x9d, 0x65, 0xc6, 0x86, 0xc7, 0xb0, 0x9e, 0x0a, 0x93, 0xe8, 0x6c, 0x84, 0xd3, 0x35, 0x58,
	0xb0, 0x8a, 0x63, 0x3f, 0x98, 0x8d, 0x7d, 0x79, 0xe8, 0xfe, 0x2c, 0xae, 0x6b, 0xdb, 0x1f, 0x40,
	0x7d, 0x28, 0x26, 0x5d, 0x1b, 0x36, 0xa1, 0xe6, 0x74, 0x5f, 0x88, 0x18, 0x4f, 0x39, 0x06, 0x79,
	0xed, 0xb6, 0xfb, 0xd0, 0x9c, 0x0f, 0x0a, 0x37, 0x60, 0x95, 0x56, 0x68, 0xaa, 0xf9, 0x2b, 0xee,
	0xe1, 0x67, 0x50, 0xc3, 0x9c, 0x4e, 0x8b, 0x0e, 0xd4, 0xe3, 0x77, 0x02, 0x45, 0x20, 0xe2, 0xcf,
	0x61, 0xef, 0x86, 0x9f, 0x33, 0x8b, 0x9c, 0x3c, 0x4d, 0x18, 0x02, 0x8c, 0x24, 0x9f, 0x08, 0x6d,
	0xa4, 0xb2, 0x45, 0x71, 0x94, 0x09, 0xa1, 0xe1, 0x5a, 0xf3, 0x49, 0x2f, 0xb3, 0x06, 0x61, 0x54,
	0x91, 0xc5, 0xef, 0x2a, 0x7e, 0xe8, 0x8b, 0xab, 0xba, 0x54, 0x14, 0xd6, 0x25, 0xe4, 0x8c, 0x90,
	0xcf, 0x35, 0x56, 0xa1, 0xdc, 0x4f, 0xaf, 0xb4, 0xf2, 0xe8, 0x6d, 0xad, 0xf8, 0x36, 0xda, 0xdf,
	0xb3, 0xc5, 0xac, 0x62, 0xc9, 0x57, 0x5c, 0x76, 0x8d, 0xd5, 0x59, 0xde, 0x2f, 0xc8, 0x42, 0xdd,
	0xbb, 0x7f, 0xe7, 0x52, 0x71, 0x4b, 0xe3, 0xaf, 0xb8, 0x3b, 0xc5, 0xfd, 0x92, 0x0a, 0x9d, 0x6a,
	0x14, 0x58, 0x38, 0x99, 0x81, 0xd2, 0x96, 0xb6, 0xb4, 0x5e, 0x3a, 0xf5, 0x26, 0x56, 0xd0, 0xa2,
	0xce, 0xfe, 0x28, 0x25, 0x69, 0x57, 0x1b, 0x65, 0xbd, 0x71, 0x96, 0xdb, 0x4f, 0x9f, 0xd2, 0xbe,
	0xd6, 0xe2, 0x6f, 0x19, 0x6c, 0xfb, 0x06, 0x4c, 0x07, 0x29, 0xb3, 0xd8, 0x48, 0x9e, 0x7a, 0x46,
	0xfe, 0x0f, 0x1b, 0xc6, 0x9d, 0x3c, 0x2d, 0xfd, 0xf1, 0x74, 0xbc, 0x57, 0x0d, 0x24, 0xd4, 0x0a,
	0x2d, 0xed, 0x7d, 0xb8, 0x67, 0x88, 0x88, 0xae, 0x17, 0x7a, 0xd7, 0x6b, 0xbc, 0x5a, 0x8a, 0xc2,
	0x08, 0x91, 0x16, 0x2d, 0xe0, 0xb5, 0x38, 0xe2, 0xc9, 0x50, 0xa4, 0xdd, 0x11, 0xd7, 0xfc, 0xc2,
	0x5f, 0x36, 0xcd, 0xf8, 0x13, 0xd8, 0x9a, 0xa2, 0x51, 0xa3, 0x77, 0x00, 0x13, 0x7f, 0xe9, 0xb7,
	0xff, 0x46, 0xc4, 0x0b, 0x84, 0xb5, 0x08, 0x29, 0xbb, 0x0d, 0x29, 0x4d, 0x39, 0x3e, 0x2b, 0x85,
	0x62, 0x3a, 0x62, 0x8e, 0x90, 0x97, 0x04, 0x74, 0x39, 0x2d, 0x4b, 0x5a, 0x3a, 0xf5, 0x4f, 0x86,
	0xcf, 0x76, 0x96, 0xf5, 0x34, 0xd7, 0x13, 0xa4, 0x3a, 0x19, 0x1e, 0x67, 0x42, 0xa6, 0xa4, 0x61,
	0xe3, 0x8e, 0xf3, 0xf0, 0x66, 0xa9, 0xce, 0xc9, 0xc9, 0xa7, 0x7a, 0xfc, 0x43, 0x15, 0x36, 0x8e,
	0x9e, 0xe1, 0x93, 0x34, 0x55, 0x98, 0xc1, 0xc7, 0xe0, 0xfe, 0xc9, 0x51, 0x77, 0xd9, 0x83, 0x14,
	0xbc, 0x61, 0x08, 0x38, 0x44, 0x8f, 0x6b, 0x0f, 0x4d, 0xf0, 0x23, 0x0b, 0xf7, 0x20, 0x42, 0xc3,
	0xc2, 0x7b, 0x33, 0xf8, 0x89, 0x85, 0x0f, 0x60, 0xa7, 0x34, 0xdf, 0xb8, 0x23, 0x83, 0x9f, 0x59,
	0xf8, 0x31, 0x3c, 0x2a, 0xed, 0xb7, 0xdf, 0x87, 0xc1, 0x2f, 0x0c, 0xdf, 0xb5, 0x3d, 0x74, 0x5e,
	0xbe, 0xf1, 0xc1, 0xaf, 0x2c, 0x7c, 0x1f, 0x76, 0xaf, 0xfb, 0xcc, 0x6f, 0x7b, 0xf0, 0x5b, 0x89,
	0x69, 0x71, 0x9a, 0xe0, 0x77, 0x86, 0xef, 0xda, 0x16, 0xd9, 0xaf, 0x8b, 0x39, 0xf8, 0x83, 0xe1,
	0xfd, 0xbb, 0x59, 0xda, 0xe6, 0x85, 0x12, 0xfc, 0x39, 0x4b, 0xbb, 0x68, 0xe0, 0xc1, 0x5f, 0xac,
	0x20, 0x79, 0xd9, 0x08, 0x83, 0xbf, 0x4b, 0x2e, 0x17, 0xaa, 0x30, 0xf8, 0x87, 0x3d, 0xab, 0xbf,
	0x60, 0xdf, 0xb0, 0xff, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x3c, 0x5f, 0x79, 0x6f, 0x08,
	0x00, 0x00,
}
