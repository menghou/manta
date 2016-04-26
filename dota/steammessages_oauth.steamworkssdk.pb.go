// Code generated by protoc-gen-go.
// source: steammessages_oauth.steamworkssdk.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type COAuthToken_ImplicitGrantNoPrompt_Request struct {
	Clientid         *string `protobuf:"bytes,1,opt,name=clientid" json:"clientid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *COAuthToken_ImplicitGrantNoPrompt_Request) Reset() {
	*m = COAuthToken_ImplicitGrantNoPrompt_Request{}
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Request) String() string { return proto.CompactTextString(m) }
func (*COAuthToken_ImplicitGrantNoPrompt_Request) ProtoMessage()    {}
func (*COAuthToken_ImplicitGrantNoPrompt_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor27, []int{0}
}

func (m *COAuthToken_ImplicitGrantNoPrompt_Request) GetClientid() string {
	if m != nil && m.Clientid != nil {
		return *m.Clientid
	}
	return ""
}

type COAuthToken_ImplicitGrantNoPrompt_Response struct {
	AccessToken      *string `protobuf:"bytes,1,opt,name=access_token" json:"access_token,omitempty"`
	RedirectUri      *string `protobuf:"bytes,2,opt,name=redirect_uri" json:"redirect_uri,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *COAuthToken_ImplicitGrantNoPrompt_Response) Reset() {
	*m = COAuthToken_ImplicitGrantNoPrompt_Response{}
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Response) String() string {
	return proto.CompactTextString(m)
}
func (*COAuthToken_ImplicitGrantNoPrompt_Response) ProtoMessage() {}
func (*COAuthToken_ImplicitGrantNoPrompt_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor27, []int{1}
}

func (m *COAuthToken_ImplicitGrantNoPrompt_Response) GetAccessToken() string {
	if m != nil && m.AccessToken != nil {
		return *m.AccessToken
	}
	return ""
}

func (m *COAuthToken_ImplicitGrantNoPrompt_Response) GetRedirectUri() string {
	if m != nil && m.RedirectUri != nil {
		return *m.RedirectUri
	}
	return ""
}

func init() {
	proto.RegisterType((*COAuthToken_ImplicitGrantNoPrompt_Request)(nil), "dota.COAuthToken_ImplicitGrantNoPrompt_Request")
	proto.RegisterType((*COAuthToken_ImplicitGrantNoPrompt_Response)(nil), "dota.COAuthToken_ImplicitGrantNoPrompt_Response")
}

var fileDescriptor27 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x8b, 0x13, 0x41,
	0x10, 0x25, 0x8b, 0x07, 0x1d, 0xf6, 0x34, 0x20, 0x84, 0x5c, 0x2c, 0xe2, 0x21, 0xbb, 0x22, 0xe3,
	0x07, 0x0a, 0x2b, 0x78, 0xd1, 0x15, 0x24, 0x97, 0x55, 0xa2, 0x9e, 0x87, 0x4e, 0x77, 0x65, 0xa6,
	0x48, 0xa6, 0x7b, 0xec, 0xaa, 0xde, 0xb0, 0x78, 0xf3, 0x3f, 0x89, 0x07, 0x7f, 0x97, 0x77, 0x6b,
	0x3a, 0x11, 0x5d, 0xc8, 0x21, 0xd7, 0xd7, 0xef, 0x75, 0xbd, 0x7a, 0xaf, 0x8a, 0x19, 0x0b, 0x9a,
	0xae, 0x43, 0x66, 0xd3, 0x20, 0xd7, 0xc1, 0x24, 0x69, 0xab, 0x8c, 0x6d, 0x43, 0x5c, 0x33, 0xbb,
	0x75, 0xd5, 0xc7, 0x20, 0xa1, 0xbc, 0xe3, 0x82, 0x98, 0x49, 0x75, 0x9b, 0x9e, 0x3c, 0xad, 0x08,
	0x5d, 0xbd, 0x34, 0x8c, 0x87, 0x54, 0xd3, 0x6f, 0xc5, 0xf9, 0xe5, 0x87, 0x37, 0xfa, 0xe7, 0xe7,
	0xb0, 0x46, 0x5f, 0xcf, 0xbb, 0x7e, 0x43, 0x96, 0xe4, 0x7d, 0x34, 0x5e, 0xae, 0xc2, 0xc7, 0x18,
	0xba, 0x5e, 0xea, 0x05, 0x7e, 0x4d, 0xc8, 0x52, 0x5e, 0x15, 0x77, 0xed, 0x86, 0xd0, 0x0b, 0xb9,
	0xf1, 0x08, 0x46, 0x67, 0xf7, 0xde, 0xbe, 0xfe, 0xfe, 0x63, 0x7c, 0x71, 0x99, 0x31, 0x98, 0xbf,
	0x83, 0x55, 0x88, 0xb0, 0x6d, 0xc9, 0xb6, 0x20, 0x01, 0x6c, 0x48, 0x8a, 0x4a, 0x8b, 0xe0, 0x53,
	0xb7, 0xc4, 0x08, 0x61, 0x05, 0xc4, 0x9c, 0xd0, 0xe9, 0xab, 0x0e, 0xe3, 0xe9, 0xcf, 0x51, 0xf1,
	0xe8, 0x98, 0xe9, 0xdc, 0x07, 0xcf, 0x58, 0xbe, 0x2a, 0x4e, 0x8d, 0xb5, 0xba, 0x5a, 0x9d, 0xf5,
	0x7b, 0x0b, 0x0f, 0xd5, 0xc2, 0x83, 0xfc, 0x09, 0xe4, 0x5f, 0x1e, 0x43, 0x33, 0xc8, 0x75, 0x4a,
	0xf0, 0xc0, 0x29, 0x0b, 0xd4, 0xf9, 0x69, 0x44, 0x47, 0x11, 0xad, 0xd4, 0x29, 0xd2, 0xf8, 0x24,
	0x4b, 0x2f, 0x54, 0xfa, 0x62, 0xb1, 0xc7, 0x49, 0xe9, 0x5f, 0x16, 0x73, 0xd0, 0x5c, 0xae, 0xc9,
	0xa9, 0xdc, 0x29, 0xd1, 0x37, 0xb0, 0xdb, 0x18, 0x22, 0x36, 0xc4, 0x12, 0xcd, 0xc0, 0xab, 0x9e,
	0xff, 0x3a, 0x29, 0x8a, 0x7f, 0xc6, 0xcb, 0xdf, 0xa3, 0xe2, 0xfe, 0x41, 0xf3, 0xe5, 0x93, 0x6a,
	0xa8, 0xa5, 0x3a, 0x3a, 0xe3, 0xc9, 0xd3, 0xe3, 0x05, 0xbb, 0x58, 0xa6, 0x37, 0xba, 0x45, 0xca,
	0x6f, 0x0c, 0xc6, 0x03, 0xed, 0xd9, 0xb0, 0xcb, 0x24, 0x47, 0x05, 0x67, 0x39, 0x13, 0x90, 0x9b,
	0x1e, 0x61, 0x96, 0xa1, 0xd9, 0x79, 0x6e, 0x6b, 0xa8, 0x87, 0x7b, 0xb4, 0xf9, 0x4c, 0xfe, 0xee,
	0xaa, 0x4d, 0x6a, 0x18, 0x4b, 0x6c, 0xcd, 0x66, 0x35, 0xb4, 0x66, 0x20, 0xb1, 0xf6, 0xb7, 0x25,
	0x69, 0x43, 0x92, 0x21, 0x21, 0x35, 0xa0, 0xd9, 0x4c, 0x5e, 0xea, 0xe8, 0x67, 0x9f, 0x30, 0x5e,
	0x93, 0x45, 0xed, 0xdb, 0x8b, 0x21, 0x3f, 0x64, 0xd6, 0xa1, 0x32, 0x1d, 0x0f, 0x57, 0xd0, 0x19,
	0xaf, 0x97, 0xf8, 0xbf, 0x19, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x7f, 0xd0, 0x7e, 0xd4,
	0x02, 0x00, 0x00,
}
