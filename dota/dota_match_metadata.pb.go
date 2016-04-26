// Code generated by protoc-gen-go.
// source: dota_match_metadata.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CDOTAMatchMetadataFile struct {
	Version          *int32              `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	MatchId          *uint64             `protobuf:"varint,2,req,name=match_id" json:"match_id,omitempty"`
	Metadata         *CDOTAMatchMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	PrivateMetadata  []byte              `protobuf:"bytes,4,opt,name=private_metadata" json:"private_metadata,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *CDOTAMatchMetadataFile) Reset()                    { *m = CDOTAMatchMetadataFile{} }
func (m *CDOTAMatchMetadataFile) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadataFile) ProtoMessage()               {}
func (*CDOTAMatchMetadataFile) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *CDOTAMatchMetadataFile) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *CDOTAMatchMetadataFile) GetMatchId() uint64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CDOTAMatchMetadataFile) GetMetadata() *CDOTAMatchMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CDOTAMatchMetadataFile) GetPrivateMetadata() []byte {
	if m != nil {
		return m.PrivateMetadata
	}
	return nil
}

type CDOTAMatchMetadata struct {
	Teams            []*CDOTAMatchMetadata_Team  `protobuf:"bytes,1,rep,name=teams" json:"teams,omitempty"`
	ItemRewards      []*CLobbyTimedRewardDetails `protobuf:"bytes,2,rep,name=item_rewards" json:"item_rewards,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *CDOTAMatchMetadata) Reset()                    { *m = CDOTAMatchMetadata{} }
func (m *CDOTAMatchMetadata) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata) ProtoMessage()               {}
func (*CDOTAMatchMetadata) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *CDOTAMatchMetadata) GetTeams() []*CDOTAMatchMetadata_Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *CDOTAMatchMetadata) GetItemRewards() []*CLobbyTimedRewardDetails {
	if m != nil {
		return m.ItemRewards
	}
	return nil
}

type CDOTAMatchMetadata_Team struct {
	DotaTeam          *uint32                           `protobuf:"varint,1,opt,name=dota_team" json:"dota_team,omitempty"`
	Players           []*CDOTAMatchMetadata_Team_Player `protobuf:"bytes,2,rep,name=players" json:"players,omitempty"`
	GraphExperience   []float32                         `protobuf:"fixed32,3,rep,name=graph_experience" json:"graph_experience,omitempty"`
	GraphGoldEarned   []float32                         `protobuf:"fixed32,4,rep,name=graph_gold_earned" json:"graph_gold_earned,omitempty"`
	GraphNetWorth     []float32                         `protobuf:"fixed32,5,rep,name=graph_net_worth" json:"graph_net_worth,omitempty"`
	CmFirstPick       *bool                             `protobuf:"varint,6,opt,name=cm_first_pick" json:"cm_first_pick,omitempty"`
	CmCaptainPlayerId *uint32                           `protobuf:"varint,7,opt,name=cm_captain_player_id" json:"cm_captain_player_id,omitempty"`
	CmBans            []uint32                          `protobuf:"varint,8,rep,name=cm_bans" json:"cm_bans,omitempty"`
	CmPicks           []uint32                          `protobuf:"varint,9,rep,name=cm_picks" json:"cm_picks,omitempty"`
	CmPenalty         *uint32                           `protobuf:"varint,10,opt,name=cm_penalty" json:"cm_penalty,omitempty"`
	XXX_unrecognized  []byte                            `json:"-"`
}

func (m *CDOTAMatchMetadata_Team) Reset()                    { *m = CDOTAMatchMetadata_Team{} }
func (m *CDOTAMatchMetadata_Team) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team) ProtoMessage()               {}
func (*CDOTAMatchMetadata_Team) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1, 0} }

func (m *CDOTAMatchMetadata_Team) GetDotaTeam() uint32 {
	if m != nil && m.DotaTeam != nil {
		return *m.DotaTeam
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team) GetPlayers() []*CDOTAMatchMetadata_Team_Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphExperience() []float32 {
	if m != nil {
		return m.GraphExperience
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphGoldEarned() []float32 {
	if m != nil {
		return m.GraphGoldEarned
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphNetWorth() []float32 {
	if m != nil {
		return m.GraphNetWorth
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmFirstPick() bool {
	if m != nil && m.CmFirstPick != nil {
		return *m.CmFirstPick
	}
	return false
}

func (m *CDOTAMatchMetadata_Team) GetCmCaptainPlayerId() uint32 {
	if m != nil && m.CmCaptainPlayerId != nil {
		return *m.CmCaptainPlayerId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team) GetCmBans() []uint32 {
	if m != nil {
		return m.CmBans
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmPicks() []uint32 {
	if m != nil {
		return m.CmPicks
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmPenalty() uint32 {
	if m != nil && m.CmPenalty != nil {
		return *m.CmPenalty
	}
	return 0
}

type CDOTAMatchMetadata_Team_PlayerKill struct {
	VictimSlot       *uint32 `protobuf:"varint,1,opt,name=victim_slot" json:"victim_slot,omitempty"`
	Count            *uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) Reset()         { *m = CDOTAMatchMetadata_Team_PlayerKill{} }
func (m *CDOTAMatchMetadata_Team_PlayerKill) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_PlayerKill) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_PlayerKill) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{1, 0, 0}
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) GetVictimSlot() uint32 {
	if m != nil && m.VictimSlot != nil {
		return *m.VictimSlot
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type CDOTAMatchMetadata_Team_Player struct {
	AccountId         *uint32                               `protobuf:"varint,1,opt,name=account_id" json:"account_id,omitempty"`
	AbilityUpgrades   []uint32                              `protobuf:"varint,2,rep,name=ability_upgrades" json:"ability_upgrades,omitempty"`
	PlayerSlot        *uint32                               `protobuf:"varint,3,opt,name=player_slot" json:"player_slot,omitempty"`
	EquippedEconItems []*CSOEconItem                        `protobuf:"bytes,4,rep,name=equipped_econ_items" json:"equipped_econ_items,omitempty"`
	Kills             []*CDOTAMatchMetadata_Team_PlayerKill `protobuf:"bytes,5,rep,name=kills" json:"kills,omitempty"`
	XXX_unrecognized  []byte                                `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_Player) Reset()         { *m = CDOTAMatchMetadata_Team_Player{} }
func (m *CDOTAMatchMetadata_Team_Player) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_Player) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{1, 0, 1}
}

func (m *CDOTAMatchMetadata_Team_Player) GetAccountId() uint32 {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAbilityUpgrades() []uint32 {
	if m != nil {
		return m.AbilityUpgrades
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetPlayerSlot() uint32 {
	if m != nil && m.PlayerSlot != nil {
		return *m.PlayerSlot
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetEquippedEconItems() []*CSOEconItem {
	if m != nil {
		return m.EquippedEconItems
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetKills() []*CDOTAMatchMetadata_Team_PlayerKill {
	if m != nil {
		return m.Kills
	}
	return nil
}

type CDOTAMatchPrivateMetadata struct {
	Kills            []*CDOTAMatchPrivateMetadata_Kill `protobuf:"bytes,1,rep,name=kills" json:"kills,omitempty"`
	TestString       *string                           `protobuf:"bytes,100,opt,name=test_string" json:"test_string,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata) Reset()                    { *m = CDOTAMatchPrivateMetadata{} }
func (m *CDOTAMatchPrivateMetadata) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata) ProtoMessage()               {}
func (*CDOTAMatchPrivateMetadata) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *CDOTAMatchPrivateMetadata) GetKills() []*CDOTAMatchPrivateMetadata_Kill {
	if m != nil {
		return m.Kills
	}
	return nil
}

func (m *CDOTAMatchPrivateMetadata) GetTestString() string {
	if m != nil && m.TestString != nil {
		return *m.TestString
	}
	return ""
}

type CDOTAMatchPrivateMetadata_Kill struct {
	Timestamp        *int32   `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	KillerHeroId     *uint32  `protobuf:"varint,2,opt,name=killer_hero_id" json:"killer_hero_id,omitempty"`
	VictimHeroId     *uint32  `protobuf:"varint,3,opt,name=victim_hero_id" json:"victim_hero_id,omitempty"`
	AssistHeroIds    []uint32 `protobuf:"varint,4,rep,name=assist_hero_ids" json:"assist_hero_ids,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata_Kill) Reset()         { *m = CDOTAMatchPrivateMetadata_Kill{} }
func (m *CDOTAMatchPrivateMetadata_Kill) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata_Kill) ProtoMessage()    {}
func (*CDOTAMatchPrivateMetadata_Kill) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{2, 0}
}

func (m *CDOTAMatchPrivateMetadata_Kill) GetTimestamp() int32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Kill) GetKillerHeroId() uint32 {
	if m != nil && m.KillerHeroId != nil {
		return *m.KillerHeroId
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Kill) GetVictimHeroId() uint32 {
	if m != nil && m.VictimHeroId != nil {
		return *m.VictimHeroId
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Kill) GetAssistHeroIds() []uint32 {
	if m != nil {
		return m.AssistHeroIds
	}
	return nil
}

func init() {
	proto.RegisterType((*CDOTAMatchMetadataFile)(nil), "dota.CDOTAMatchMetadataFile")
	proto.RegisterType((*CDOTAMatchMetadata)(nil), "dota.CDOTAMatchMetadata")
	proto.RegisterType((*CDOTAMatchMetadata_Team)(nil), "dota.CDOTAMatchMetadata.Team")
	proto.RegisterType((*CDOTAMatchMetadata_Team_PlayerKill)(nil), "dota.CDOTAMatchMetadata.Team.PlayerKill")
	proto.RegisterType((*CDOTAMatchMetadata_Team_Player)(nil), "dota.CDOTAMatchMetadata.Team.Player")
	proto.RegisterType((*CDOTAMatchPrivateMetadata)(nil), "dota.CDOTAMatchPrivateMetadata")
	proto.RegisterType((*CDOTAMatchPrivateMetadata_Kill)(nil), "dota.CDOTAMatchPrivateMetadata.Kill")
}

var fileDescriptor12 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x4e, 0x14, 0x4d,
	0x10, 0xfd, 0x86, 0xdd, 0xe5, 0xa7, 0x96, 0xf9, 0x80, 0x46, 0xb0, 0xd9, 0xa0, 0x21, 0xc4, 0x8b,
	0x8d, 0x31, 0x13, 0x83, 0x1a, 0xaf, 0x55, 0x34, 0x31, 0x4a, 0x20, 0xca, 0x7d, 0xa7, 0x67, 0xa6,
	0xdc, 0xed, 0x30, 0x33, 0x3d, 0x76, 0x37, 0xe0, 0xde, 0xe9, 0x3b, 0xf9, 0x12, 0xbe, 0x8c, 0x37,
	0xbe, 0x80, 0xd5, 0xdd, 0xb3, 0x90, 0x28, 0x1a, 0x6f, 0xcf, 0xa9, 0x73, 0xaa, 0xea, 0x54, 0x37,
	0xec, 0x94, 0xda, 0x49, 0x51, 0x4b, 0x57, 0x4c, 0x45, 0x8d, 0x4e, 0x96, 0xd2, 0xc9, 0xac, 0x35,
	0xda, 0x69, 0xd6, 0xf7, 0xd4, 0x68, 0x2b, 0x97, 0x16, 0xc5, 0xa4, 0xa8, 0xd1, 0x5a, 0x39, 0x41,
	0x1b, 0xc9, 0xd1, 0x6e, 0xd0, 0x5d, 0xc3, 0xa2, 0xd0, 0x75, 0xad, 0x9b, 0xc8, 0xee, 0x7f, 0x49,
	0x60, 0xfb, 0xc5, 0xe1, 0xf1, 0xe9, 0xb3, 0x23, 0x6f, 0x7c, 0xd4, 0xf9, 0xbe, 0x52, 0x15, 0xb2,
	0x35, 0x58, 0xba, 0x40, 0x63, 0x95, 0x6e, 0x78, 0xb2, 0xb7, 0x30, 0x1e, 0xb0, 0x75, 0x58, 0x8e,
	0xed, 0x55, 0xc9, 0x17, 0x08, 0xe9, 0xb3, 0xfb, 0x84, 0x74, 0x12, 0xde, 0xdb, 0x4b, 0xc6, 0xc3,
	0x03, 0x9e, 0xf9, 0x76, 0xd9, 0xef, 0x96, 0x8c, 0xc3, 0x7a, 0x6b, 0xd4, 0x85, 0x74, 0x78, 0x35,
	0x3e, 0xef, 0x93, 0x66, 0x75, 0xff, 0x7b, 0x1f, 0xd8, 0x0d, 0x82, 0x07, 0x30, 0x70, 0x28, 0x6b,
	0x4b, 0xdd, 0x7b, 0xe4, 0x7c, 0xe7, 0x4f, 0xce, 0xd9, 0x29, 0x55, 0xb1, 0xc7, 0xb0, 0xaa, 0x1c,
	0xd6, 0xc2, 0xe0, 0xa5, 0x34, 0xa5, 0xa5, 0x01, 0xbd, 0xe8, 0x6e, 0x27, 0x7a, 0xab, 0xf3, 0x7c,
	0x76, 0xaa, 0x6a, 0x2c, 0xdf, 0x05, 0xfe, 0x90, 0xb4, 0xaa, 0xb2, 0xa3, 0x1f, 0x3d, 0xe8, 0x07,
	0xf9, 0x06, 0xac, 0x84, 0x9c, 0x7c, 0x47, 0x6a, 0x98, 0x8c, 0x53, 0xf6, 0x04, 0x96, 0xda, 0x4a,
	0xce, 0x28, 0x82, 0xce, 0xec, 0xde, 0x5f, 0x27, 0xc8, 0x4e, 0x42, 0xb1, 0xdf, 0x73, 0x62, 0x64,
	0x3b, 0x15, 0xf8, 0xa9, 0x45, 0xa3, 0xb0, 0x29, 0x90, 0xb2, 0xe9, 0x8d, 0x17, 0xd8, 0x0e, 0x6c,
	0x44, 0x66, 0xa2, 0xab, 0x52, 0xa0, 0x34, 0x0d, 0x96, 0x14, 0x81, 0xa7, 0x6e, 0xc3, 0x5a, 0xa4,
	0x1a, 0x74, 0xe2, 0x52, 0x1b, 0x37, 0xe5, 0x83, 0x40, 0x6c, 0x41, 0x5a, 0xd4, 0xe2, 0x83, 0x32,
	0xd6, 0x89, 0x56, 0x15, 0x67, 0x7c, 0x91, 0x66, 0x5b, 0x66, 0xbb, 0x70, 0x8b, 0xe0, 0x42, 0xb6,
	0xb4, 0x46, 0x23, 0xe2, 0x98, 0xfe, 0x2c, 0x4b, 0x61, 0x72, 0xba, 0x1c, 0xb1, 0xb9, 0x6c, 0x2c,
	0x5f, 0x26, 0x97, 0xd4, 0x5f, 0x8e, 0x00, 0xaf, 0xb7, 0x7c, 0x25, 0x20, 0x0c, 0xc0, 0x23, 0xd8,
	0xc8, 0xca, 0xcd, 0x38, 0x78, 0xd9, 0xe8, 0x21, 0x40, 0xdc, 0xe1, 0x8d, 0xaa, 0x2a, 0xb6, 0x09,
	0xc3, 0x0b, 0x55, 0x38, 0x55, 0x0b, 0x5b, 0x69, 0xd7, 0x65, 0x92, 0xc2, 0xa0, 0xd0, 0xe7, 0x8d,
	0xa3, 0x44, 0xbc, 0xe2, 0x6b, 0x02, 0x8b, 0xdd, 0xda, 0x64, 0x28, 0x8b, 0xc0, 0xf9, 0x39, 0x62,
	0x35, 0x45, 0x21, 0x73, 0x55, 0x29, 0x37, 0x13, 0xe7, 0x2d, 0xed, 0x57, 0x62, 0x8c, 0x32, 0xf5,
	0xe6, 0xdd, 0xd0, 0xc1, 0xbc, 0x17, 0xca, 0x33, 0xd8, 0xc4, 0x8f, 0xe7, 0xaa, 0x6d, 0x91, 0xd2,
	0x29, 0x74, 0x23, 0xfc, 0x41, 0x6d, 0x48, 0x68, 0x78, 0xb0, 0xd1, 0x85, 0xff, 0xfe, 0xf8, 0x25,
	0x51, 0xaf, 0x89, 0x61, 0x4f, 0x61, 0x70, 0x46, 0x93, 0xda, 0x10, 0xd5, 0xf0, 0x60, 0xfc, 0x2f,
	0xe7, 0xf1, 0xab, 0xed, 0x7f, 0x4b, 0x60, 0xe7, 0xba, 0xec, 0x24, 0xbe, 0xca, 0xab, 0x77, 0xf7,
	0x68, 0x6e, 0x9b, 0xdc, 0x7c, 0xf5, 0x5f, 0xea, 0xb3, 0x79, 0x5a, 0x0e, 0xe9, 0x46, 0xd6, 0x19,
	0xd5, 0x4c, 0x78, 0x49, 0x0b, 0xad, 0x8c, 0x72, 0xe8, 0x07, 0x92, 0x1e, 0x17, 0xe5, 0x48, 0xb4,
	0xac, 0xdb, 0x10, 0xcd, 0x80, 0x6d, 0xc3, 0xff, 0xbe, 0x09, 0x05, 0x30, 0x45, 0xa3, 0xe3, 0x8f,
	0xf2, 0x19, 0x10, 0xde, 0xa5, 0x3e, 0xc7, 0x63, 0x36, 0xf4, 0x40, 0xa4, 0xb5, 0x8a, 0x3a, 0x74,
	0x78, 0xcc, 0x25, 0x7d, 0xde, 0xfb, 0x9c, 0xfc, 0xf7, 0x33, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x23,
	0xb4, 0x6b, 0x1c, 0x04, 0x00, 0x00,
}
