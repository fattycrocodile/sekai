// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/gov/permission.proto

package gov

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PermValue int32

const (
	// PERMISSION_ZERO is a no-op permission.
	PermValue_PERMISSION_ZERO PermValue = 0
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set Permissions to other actors.
	PermValue_PERMISSION_SET_PERMISSIONS PermValue = 1
	// PERMISSION_CLAIM_VALIDATOR defines the permission that allows to Claim a validator Seat.
	PermValue_PERMISSION_CLAIM_VALIDATOR PermValue = 2
	// PERMISSION_CLAIM_COUNCILOR defines the permission that allows to Claim a Councilor Seat.
	PermValue_PERMISSION_CLAIM_COUNCILOR PermValue = 3
	// PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL defines the permission needed to create proposals for setting permissions.
	PermValue_PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL PermValue = 4
	// PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set permissions.
	PermValue_PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL PermValue = 5
	//  PERMISSION_UPSERT_TOKEN_ALIAS
	PermValue_PERMISSION_UPSERT_TOKEN_ALIAS PermValue = 6
	// PERMISSION_CHANGE_TX_FEE
	PermValue_PERMISSION_CHANGE_TX_FEE PermValue = 7
	// PERMISSION_UPSERT_TOKEN_RATE
	PermValue_PERMISSION_UPSERT_TOKEN_RATE PermValue = 8
	// PERMISSION_UPSERT_ROLE makes possible to add, modify and assign roles.
	PermValue_PERMISSION_UPSERT_ROLE PermValue = 9
	// PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermValue_PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL PermValue = 10
	// PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermValue_PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL PermValue = 11
	// PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission needed to create proposals for setting network property.
	PermValue_PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL PermValue = 12
	// PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set network property.
	PermValue_PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL PermValue = 13
	// PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to create proposals for upsert token Alias.
	PermValue_PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL PermValue = 14
	// PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to vote proposals for upsert token.
	PermValue_PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL PermValue = 15
	// PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES defines the permission needed to create proposals for setting poor network messages
	PermValue_PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES PermValue = 16
	// PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL defines the permission needed to vote proposals to set poor network messages
	PermValue_PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL PermValue = 17
	// PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to create proposals for upsert token rate.
	PermValue_PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL PermValue = 18
	// PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to vote proposals for upsert token rate.
	PermValue_PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL PermValue = 19
	// PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to create a proposal to unjail a validator.
	PermValue_PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL PermValue = 20
	// PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to vote a proposal to unjail a validator.
	PermValue_PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL PermValue = 21
	// PERMISSION_CREATE_CREATE_ROLE_PROPOSAL defines the permission needed to create a proposal to create a role.
	PermValue_PERMISSION_CREATE_CREATE_ROLE_PROPOSAL PermValue = 22
	// PERMISSION_VOTE_CREATE_ROLE_PROPOSAL defines the permission needed to vote a proposal to create a role.
	PermValue_PERMISSION_VOTE_CREATE_ROLE_PROPOSAL PermValue = 23
	// PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to create a proposal to blacklist/whitelisted tokens
	PermValue_PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL PermValue = 24
	// PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to vote on blacklist/whitelisted tokens proposal
	PermValue_PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL PermValue = 25
	// PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to create a proposal to reset whole validator rank
	PermValue_PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL PermValue = 26
	// PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to vote on reset whole validator rank proposal
	PermValue_PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL PermValue = 27
	// PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL defines the permission needed to create a proposal for software upgrade
	PermValue_PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL PermValue = 28
	// PERMISSION_SOFTWARE_UPGRADE_PROPOSAL defines the permission needed to vote on software upgrade proposal
	PermValue_PERMISSION_SOFTWARE_UPGRADE_PROPOSAL PermValue = 29
)

var PermValue_name = map[int32]string{
	0:  "PERMISSION_ZERO",
	1:  "PERMISSION_SET_PERMISSIONS",
	2:  "PERMISSION_CLAIM_VALIDATOR",
	3:  "PERMISSION_CLAIM_COUNCILOR",
	4:  "PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL",
	5:  "PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL",
	6:  "PERMISSION_UPSERT_TOKEN_ALIAS",
	7:  "PERMISSION_CHANGE_TX_FEE",
	8:  "PERMISSION_UPSERT_TOKEN_RATE",
	9:  "PERMISSION_UPSERT_ROLE",
	10: "PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL",
	11: "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL",
	12: "PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL",
	13: "PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL",
	14: "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	15: "PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL",
	16: "PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES",
	17: "PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL",
	18: "PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL",
	19: "PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL",
	20: "PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL",
	21: "PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL",
	22: "PERMISSION_CREATE_CREATE_ROLE_PROPOSAL",
	23: "PERMISSION_VOTE_CREATE_ROLE_PROPOSAL",
	24: "PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
	25: "PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL",
	26: "PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
	27: "PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL",
	28: "PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL",
	29: "PERMISSION_SOFTWARE_UPGRADE_PROPOSAL",
}

var PermValue_value = map[string]int32{
	"PERMISSION_ZERO":                                       0,
	"PERMISSION_SET_PERMISSIONS":                            1,
	"PERMISSION_CLAIM_VALIDATOR":                            2,
	"PERMISSION_CLAIM_COUNCILOR":                            3,
	"PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL":            4,
	"PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL":              5,
	"PERMISSION_UPSERT_TOKEN_ALIAS":                         6,
	"PERMISSION_CHANGE_TX_FEE":                              7,
	"PERMISSION_UPSERT_TOKEN_RATE":                          8,
	"PERMISSION_UPSERT_ROLE":                                9,
	"PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL":       10,
	"PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL":         11,
	"PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL":       12,
	"PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL":         13,
	"PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL":         14,
	"PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL":           15,
	"PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES":           16,
	"PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL":    17,
	"PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL":          18,
	"PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL":            19,
	"PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL":           20,
	"PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL":             21,
	"PERMISSION_CREATE_CREATE_ROLE_PROPOSAL":                22,
	"PERMISSION_VOTE_CREATE_ROLE_PROPOSAL":                  23,
	"PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL":  24,
	"PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL":    25,
	"PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL": 26,
	"PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL":   27,
	"PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL":           28,
	"PERMISSION_SOFTWARE_UPGRADE_PROPOSAL":                  29,
}

func (x PermValue) String() string {
	return proto.EnumName(PermValue_name, int32(x))
}

func (PermValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_214168f8815c1062, []int{0}
}

func init() {
	proto.RegisterEnum("kira.gov.PermValue", PermValue_name, PermValue_value)
}

func init() {
	proto.RegisterFile("kira/gov/permission.proto", fileDescriptor_214168f8815c1062)
}

var fileDescriptor_214168f8815c1062 = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x52, 0xdb, 0x46,
	0x1b, 0xc7, 0xf3, 0xbe, 0x4d, 0x53, 0xb2, 0x4d, 0x83, 0x2a, 0x28, 0x21, 0xdb, 0x90, 0x2e, 0x2d,
	0xa5, 0x84, 0x04, 0x3c, 0x2d, 0x4d, 0x67, 0x3a, 0x3d, 0x12, 0x66, 0x01, 0xd5, 0x1f, 0xf2, 0xac,
	0x64, 0x9c, 0x30, 0x93, 0xd1, 0x2c, 0xb0, 0x91, 0x55, 0x1b, 0xaf, 0x67, 0x25, 0x92, 0xf6, 0x0e,
	0x3a, 0x7b, 0x0f, 0x7b, 0xd4, 0x7b, 0xec, 0x71, 0x47, 0x02, 0x6b, 0x25, 0x59, 0x36, 0x1c, 0x81,
	0xa5, 0x7d, 0x7e, 0xcf, 0x7f, 0xff, 0xcf, 0xc7, 0x08, 0x3c, 0x1d, 0x84, 0x82, 0xd6, 0x02, 0xfe,
	0xa1, 0x36, 0x66, 0xe2, 0x32, 0x8c, 0xa2, 0x90, 0x8f, 0x76, 0xc7, 0x82, 0xc7, 0xdc, 0x5c, 0x48,
	0x5e, 0xed, 0x06, 0xfc, 0x03, 0x5c, 0x0e, 0x78, 0xc0, 0xd3, 0x87, 0xb5, 0xe4, 0xbf, 0xeb, 0xf7,
	0xdb, 0xff, 0x9a, 0xe0, 0x61, 0x87, 0x89, 0xcb, 0x13, 0x3a, 0xbc, 0x62, 0xe6, 0x3a, 0x58, 0xec,
	0x60, 0xd2, 0xb2, 0x5d, 0xd7, 0x76, 0xda, 0xfe, 0x29, 0x26, 0x8e, 0x71, 0x0f, 0x3e, 0x92, 0x0a,
	0x2d, 0x24, 0x67, 0x4e, 0x99, 0xe0, 0xe6, 0x2f, 0x00, 0xe6, 0x8e, 0xb8, 0xd8, 0xf3, 0xf5, 0x4f,
	0xd7, 0xf8, 0x1f, 0x5c, 0x91, 0x0a, 0x99, 0xc9, 0x69, 0x97, 0xc5, 0x9d, 0x4c, 0x4d, 0x54, 0x8a,
	0xab, 0x37, 0x2d, 0xbb, 0xe5, 0x9f, 0x58, 0x4d, 0xfb, 0xc0, 0xf2, 0x1c, 0x62, 0xfc, 0x5f, 0xc7,
	0xd5, 0x87, 0x34, 0x4c, 0xe4, 0x84, 0x17, 0x34, 0xe6, 0xa2, 0x32, 0xae, 0xee, 0x74, 0xdb, 0x75,
	0xbb, 0xe9, 0x10, 0xe3, 0x93, 0x52, 0x5c, 0x9d, 0x5f, 0x8d, 0xce, 0xc3, 0x21, 0x17, 0xa6, 0x07,
	0xb6, 0xf3, 0x71, 0x04, 0x5b, 0x1e, 0x2e, 0xcb, 0xf5, 0x3b, 0xc4, 0xe9, 0x38, 0xae, 0xd5, 0x34,
	0xee, 0xc3, 0x0d, 0xa9, 0x10, 0x4a, 0x39, 0x82, 0xd1, 0x98, 0x15, 0xd5, 0x77, 0x04, 0x1f, 0xf3,
	0x88, 0x0e, 0x4d, 0x07, 0x6c, 0xe5, 0xa8, 0x27, 0xce, 0x3c, 0xe6, 0xa7, 0x70, 0x5d, 0x2a, 0xb4,
	0x96, 0xba, 0xcb, 0x4b, 0xc4, 0x0c, 0xf8, 0x1b, 0x58, 0xcb, 0x01, 0xbb, 0x1d, 0x17, 0x13, 0xcf,
	0xf7, 0x9c, 0x06, 0x6e, 0xfb, 0x56, 0xd3, 0xb6, 0x5c, 0xe3, 0x01, 0x5c, 0x95, 0x0a, 0x2d, 0x27,
	0xa1, 0xdd, 0x71, 0xc4, 0x44, 0xec, 0xf1, 0x01, 0x1b, 0x59, 0xc3, 0x90, 0x46, 0xe6, 0x8f, 0x60,
	0x35, 0x7f, 0xc7, 0x63, 0xab, 0x7d, 0x84, 0x7d, 0xef, 0x8d, 0x7f, 0x88, 0xb1, 0xf1, 0x19, 0x5c,
	0x92, 0x0a, 0x2d, 0xa6, 0x37, 0xea, 0xd3, 0x51, 0xc0, 0xbc, 0x3f, 0x0f, 0x19, 0x33, 0x7f, 0x05,
	0xcf, 0x66, 0xe5, 0x23, 0x96, 0x87, 0x8d, 0x05, 0xf8, 0x44, 0x2a, 0xb4, 0x54, 0x4a, 0x47, 0x68,
	0xcc, 0xcc, 0x5d, 0xb0, 0x32, 0x1d, 0x4a, 0x9c, 0x26, 0x36, 0x1e, 0x42, 0x53, 0x2a, 0xf4, 0x58,
	0x07, 0x11, 0x3e, 0x64, 0xe6, 0x3b, 0x50, 0x9b, 0xae, 0xc0, 0x4d, 0xd8, 0x81, 0xe5, 0x59, 0x3e,
	0xc1, 0x47, 0xb6, 0xeb, 0x91, 0xb7, 0xda, 0x32, 0x00, 0xb7, 0xa4, 0x42, 0x1b, 0xba, 0x0c, 0xd7,
	0xb8, 0x03, 0x1a, 0x53, 0xc2, 0x82, 0x30, 0x8a, 0xc5, 0x5f, 0x99, 0x73, 0x6f, 0xc1, 0x4e, 0xb9,
	0x14, 0xf3, 0xe1, 0x9f, 0xc3, 0x4d, 0xa9, 0xd0, 0xb7, 0x93, 0x7a, 0xcc, 0x41, 0x57, 0x2a, 0x4f,
	0xea, 0xdc, 0xc6, 0x5e, 0xcf, 0x21, 0x8d, 0x94, 0x89, 0x89, 0x97, 0x83, 0x3f, 0x2a, 0x2b, 0x77,
	0x59, 0xdc, 0x66, 0xf1, 0x47, 0x2e, 0x06, 0x09, 0x96, 0x89, 0x78, 0xae, 0xf2, 0xf9, 0xf0, 0x2f,
	0x8a, 0xca, 0xef, 0x8c, 0x2e, 0x7a, 0x9e, 0xeb, 0x2a, 0x8d, 0x7e, 0xac, 0xd1, 0x79, 0xc7, 0x75,
	0x93, 0x65, 0xe8, 0x2e, 0x78, 0x39, 0xc3, 0xef, 0x4a, 0xf0, 0xa2, 0x9e, 0x28, 0xed, 0x76, 0x05,
	0xf6, 0x5d, 0x01, 0x9b, 0x9f, 0x53, 0xc7, 0x21, 0x99, 0x27, 0x2d, 0xec, 0xba, 0xd6, 0x11, 0x76,
	0x0d, 0x03, 0xbe, 0x92, 0x0a, 0x6d, 0x15, 0x07, 0x95, 0x73, 0x71, 0x63, 0x48, 0x8b, 0x45, 0x11,
	0x0d, 0x98, 0xc6, 0x9f, 0x81, 0x9f, 0x2a, 0x07, 0xb6, 0x0a, 0xae, 0xc5, 0x7f, 0x09, 0xb7, 0xa5,
	0x42, 0x9b, 0xf9, 0xd1, 0x9d, 0x93, 0xa3, 0x07, 0x5e, 0xdd, 0x62, 0x7a, 0x32, 0x5a, 0x9a, 0x6e,
	0xc2, 0xef, 0xa5, 0x42, 0xeb, 0x95, 0x9e, 0x27, 0x93, 0x96, 0x81, 0xdd, 0xc2, 0x0e, 0x9b, 0xb6,
	0xbc, 0x88, 0x5d, 0x82, 0xdf, 0x49, 0x85, 0xbe, 0xa9, 0x70, 0xbc, 0x00, 0x3d, 0xa9, 0x32, 0xbc,
	0xdb, 0xfe, 0xdd, 0xb2, 0x9b, 0x7a, 0x21, 0x6b, 0xea, 0xf2, 0x94, 0xd8, 0xd1, 0x1f, 0x34, 0x1c,
	0x66, 0x0b, 0x3a, 0xe3, 0x12, 0xf0, 0x62, 0x4a, 0xec, 0x4c, 0xea, 0x57, 0x25, 0xad, 0x33, 0x98,
	0x87, 0x60, 0x73, 0x5a, 0xeb, 0xcd, 0x9f, 0x64, 0xf3, 0x68, 0xe0, 0x0a, 0x84, 0x52, 0xa1, 0x15,
	0x2d, 0x33, 0x59, 0x41, 0x19, 0xe7, 0x18, 0x6c, 0x94, 0xb5, 0x55, 0x52, 0x9e, 0xc0, 0xe7, 0x52,
	0x21, 0x38, 0x91, 0x55, 0x41, 0x7a, 0x0f, 0x7e, 0x9e, 0x56, 0x94, 0x56, 0xc3, 0xf5, 0x7b, 0xc7,
	0xb6, 0x87, 0xfd, 0xfd, 0xa6, 0x55, 0x6f, 0x4c, 0x96, 0x71, 0x46, 0x5e, 0x2d, 0xf7, 0x6d, 0x5a,
	0x98, 0xa8, 0xd7, 0x0f, 0x63, 0xb6, 0x3f, 0xa4, 0xe7, 0x83, 0xeb, 0x25, 0x3d, 0xaf, 0x6f, 0xef,
	0x90, 0xe5, 0x69, 0xb1, 0x6f, 0x6f, 0xc9, 0xd1, 0x07, 0xaf, 0xa7, 0xef, 0x42, 0x70, 0x32, 0x1f,
	0xbd, 0xe3, 0xc4, 0x17, 0x5d, 0x38, 0x62, 0xb5, 0x1b, 0x3a, 0x0d, 0x84, 0x3b, 0x52, 0xa1, 0x17,
	0x39, 0xb3, 0x59, 0xc4, 0xe2, 0x5e, 0x9f, 0x0f, 0x59, 0x56, 0x43, 0x42, 0x47, 0x83, 0x2c, 0xd3,
	0x05, 0xd8, 0x2b, 0xdf, 0xe6, 0x2e, 0x79, 0xbe, 0x86, 0x2f, 0xa5, 0x42, 0x3f, 0x4c, 0xae, 0x73,
	0x5b, 0x96, 0xca, 0xce, 0x76, 0x9d, 0x43, 0xaf, 0x67, 0x91, 0x64, 0x72, 0x8e, 0x88, 0x75, 0x90,
	0x33, 0xeb, 0x59, 0xb9, 0xb3, 0x5d, 0xfe, 0x3e, 0xfe, 0x48, 0x05, 0xeb, 0x8e, 0x03, 0x41, 0x2f,
	0xb4, 0x4f, 0xad, 0x42, 0xf7, 0xcc, 0x06, 0xae, 0x15, 0x9b, 0x7a, 0x06, 0x0e, 0xde, 0xff, 0xfb,
	0x9f, 0xe7, 0xf7, 0xf6, 0x5f, 0x9f, 0xee, 0x05, 0x61, 0xdc, 0xbf, 0x3a, 0xdb, 0x3d, 0xe7, 0x97,
	0xb5, 0x46, 0x28, 0x68, 0x9d, 0x0b, 0x56, 0x8b, 0xd8, 0x80, 0x86, 0x35, 0xbb, 0xed, 0x61, 0xf2,
	0xa6, 0x96, 0x7e, 0xa2, 0xed, 0x04, 0x6c, 0x54, 0x9b, 0x7c, 0xe0, 0x9d, 0x3d, 0x48, 0x9f, 0xed,
	0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x56, 0xa0, 0x27, 0xf3, 0x09, 0x00, 0x00,
}
