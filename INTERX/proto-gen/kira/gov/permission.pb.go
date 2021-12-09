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
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set ClaimValidatorPermission to other actors.
	PermValue_PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION PermValue = 30
	// PERMISSION_CREATE_SET_PROPOSAL_DURATION_PROPOSAL defines the permission needed to create a proposal to set proposal duration.
	PermValue_PERMISSION_CREATE_SET_PROPOSAL_DURATION_PROPOSAL PermValue = 31
	// PERMISSION_VOTE_SET_PROPOSAL_DURATION_PROPOSAL defines the permission needed to vote a proposal to set proposal duration.
	PermValue_PERMISSION_VOTE_SET_PROPOSAL_DURATION_PROPOSAL PermValue = 32
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
	30: "PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION",
	31: "PERMISSION_CREATE_SET_PROPOSAL_DURATION_PROPOSAL",
	32: "PERMISSION_VOTE_SET_PROPOSAL_DURATION_PROPOSAL",
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
	"PERMISSION_SET_CLAIM_VALIDATOR_PERMISSION":             30,
	"PERMISSION_CREATE_SET_PROPOSAL_DURATION_PROPOSAL":      31,
	"PERMISSION_VOTE_SET_PROPOSAL_DURATION_PROPOSAL":        32,
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
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x52, 0xdb, 0x46,
	0x14, 0xc7, 0xd3, 0x36, 0x4d, 0xc9, 0x36, 0x0d, 0xaa, 0x20, 0x84, 0x6c, 0x03, 0x59, 0x1a, 0x42,
	0x80, 0x04, 0xbb, 0x2d, 0x4d, 0x67, 0x3a, 0xbd, 0x12, 0xf6, 0x02, 0xaa, 0x3f, 0xe4, 0x59, 0xc9,
	0x38, 0xa1, 0x93, 0xd1, 0x2c, 0xb0, 0x91, 0x55, 0x1b, 0xaf, 0x47, 0x12, 0x49, 0xfb, 0x06, 0x99,
	0x7d, 0x87, 0xbd, 0xea, 0x8b, 0x76, 0x64, 0x6c, 0xad, 0x24, 0xcb, 0x86, 0x2b, 0xb0, 0xa4, 0xf3,
	0x3b, 0xff, 0xfd, 0x9f, 0x8f, 0x59, 0xf0, 0xa4, 0xe7, 0x07, 0xb4, 0xec, 0xf1, 0x8f, 0xe5, 0x21,
	0x0b, 0x2e, 0xfd, 0x30, 0xf4, 0xf9, 0xa0, 0x34, 0x0c, 0x78, 0xc4, 0xf5, 0x85, 0xf8, 0x55, 0xc9,
	0xe3, 0x1f, 0xe1, 0xb2, 0xc7, 0x3d, 0x3e, 0x7a, 0x58, 0x8e, 0xff, 0xbb, 0x7e, 0xbf, 0xfb, 0xf9,
	0x11, 0xb8, 0xdf, 0x62, 0xc1, 0xe5, 0x09, 0xed, 0x5f, 0x31, 0x7d, 0x03, 0x2c, 0xb6, 0x30, 0x69,
	0x98, 0xb6, 0x6d, 0x5a, 0x4d, 0xf7, 0x14, 0x13, 0x4b, 0xbb, 0x03, 0x1f, 0x08, 0x89, 0x16, 0xe2,
	0x6f, 0x4e, 0x59, 0xc0, 0xf5, 0xdf, 0x00, 0x4c, 0x7d, 0x62, 0x63, 0xc7, 0x55, 0x3f, 0x6d, 0xed,
	0x0b, 0xb8, 0x22, 0x24, 0xd2, 0xe3, 0xaf, 0x6d, 0x16, 0xb5, 0x12, 0x35, 0x61, 0x2e, 0xae, 0x52,
	0x37, 0xcc, 0x86, 0x7b, 0x62, 0xd4, 0xcd, 0xaa, 0xe1, 0x58, 0x44, 0xfb, 0x52, 0xc5, 0x55, 0xfa,
	0xd4, 0x8f, 0xe5, 0xf8, 0x17, 0x34, 0xe2, 0x41, 0x61, 0x5c, 0xc5, 0x6a, 0x37, 0x2b, 0x66, 0xdd,
	0x22, 0xda, 0x57, 0xb9, 0xb8, 0x0a, 0xbf, 0x1a, 0x9c, 0xfb, 0x7d, 0x1e, 0xe8, 0x0e, 0xd8, 0x4d,
	0xc7, 0x11, 0x6c, 0x38, 0x38, 0x2f, 0xd7, 0x6d, 0x11, 0xab, 0x65, 0xd9, 0x46, 0x5d, 0xbb, 0x0b,
	0x37, 0x85, 0x44, 0x68, 0xc4, 0x09, 0x18, 0x8d, 0x58, 0x56, 0x7d, 0x2b, 0xe0, 0x43, 0x1e, 0xd2,
	0xbe, 0x6e, 0x81, 0xed, 0x14, 0xf5, 0xc4, 0x9a, 0xc7, 0xfc, 0x1a, 0x6e, 0x08, 0x89, 0xd6, 0x46,
	0xee, 0xf2, 0x1c, 0x31, 0x01, 0xfe, 0x01, 0xd6, 0x52, 0xc0, 0x76, 0xcb, 0xc6, 0xc4, 0x71, 0x1d,
	0xab, 0x86, 0x9b, 0xae, 0x51, 0x37, 0x0d, 0x5b, 0xbb, 0x07, 0x57, 0x85, 0x44, 0xcb, 0x71, 0x68,
	0x7b, 0x18, 0xb2, 0x20, 0x72, 0x78, 0x8f, 0x0d, 0x8c, 0xbe, 0x4f, 0x43, 0xfd, 0x67, 0xb0, 0x9a,
	0x3e, 0xe3, 0xb1, 0xd1, 0x3c, 0xc2, 0xae, 0xf3, 0xd6, 0x3d, 0xc4, 0x58, 0xfb, 0x06, 0x2e, 0x09,
	0x89, 0x16, 0x47, 0x27, 0xea, 0xd2, 0x81, 0xc7, 0x9c, 0x7f, 0x0e, 0x19, 0xd3, 0x7f, 0x07, 0x4f,
	0x67, 0xe5, 0x23, 0x86, 0x83, 0xb5, 0x05, 0xf8, 0x58, 0x48, 0xb4, 0x94, 0x4b, 0x47, 0x68, 0xc4,
	0xf4, 0x12, 0x58, 0x99, 0x0e, 0x25, 0x56, 0x1d, 0x6b, 0xf7, 0xa1, 0x2e, 0x24, 0x7a, 0xa8, 0x82,
	0x08, 0xef, 0x33, 0xfd, 0x3d, 0x28, 0x4f, 0x57, 0x60, 0x1c, 0x56, 0x35, 0x1c, 0xc3, 0x25, 0xf8,
	0xc8, 0xb4, 0x1d, 0xf2, 0x4e, 0x59, 0x06, 0xe0, 0xb6, 0x90, 0x68, 0x53, 0x95, 0xe1, 0x1a, 0x57,
	0xa5, 0x11, 0x25, 0xcc, 0xf3, 0xc3, 0x28, 0xf8, 0x37, 0x71, 0xee, 0x1d, 0xd8, 0xcb, 0x97, 0x62,
	0x3e, 0xfc, 0x5b, 0xb8, 0x25, 0x24, 0xfa, 0x71, 0x52, 0x8f, 0x39, 0xe8, 0x42, 0xe5, 0x71, 0x9d,
	0x9b, 0xd8, 0xe9, 0x58, 0xa4, 0x36, 0x62, 0x62, 0xe2, 0xa4, 0xe0, 0x0f, 0xf2, 0xca, 0x6d, 0x16,
	0x35, 0x59, 0xf4, 0x89, 0x07, 0xbd, 0x18, 0xcb, 0x82, 0x68, 0xae, 0xf2, 0xf9, 0xf0, 0xef, 0xb2,
	0xca, 0x6f, 0x8d, 0xce, 0x7a, 0x9e, 0xea, 0x2a, 0x85, 0x7e, 0xa8, 0xd0, 0x69, 0xc7, 0x55, 0x93,
	0x25, 0xe8, 0x36, 0x78, 0x35, 0xc3, 0xef, 0x42, 0xf0, 0xa2, 0x9a, 0x28, 0xe5, 0x76, 0x01, 0xf6,
	0x7d, 0x06, 0x9b, 0x9e, 0x53, 0xcb, 0x22, 0x89, 0x27, 0x0d, 0x6c, 0xdb, 0xc6, 0x11, 0xb6, 0x35,
	0x0d, 0xbe, 0x16, 0x12, 0x6d, 0x67, 0x07, 0x95, 0xf3, 0x60, 0x6c, 0x48, 0x83, 0x85, 0x21, 0xf5,
	0x98, 0xc2, 0x9f, 0x81, 0x5f, 0x0a, 0x07, 0xb6, 0x08, 0xae, 0xc4, 0x7f, 0x0f, 0x77, 0x85, 0x44,
	0x5b, 0xe9, 0xd1, 0x9d, 0x93, 0xa3, 0x03, 0x5e, 0xdf, 0x60, 0x7a, 0x3c, 0x5a, 0x8a, 0xae, 0xc3,
	0x17, 0x42, 0xa2, 0x8d, 0x42, 0xcf, 0xe3, 0x49, 0x4b, 0xc0, 0x76, 0x66, 0x87, 0x4d, 0x5b, 0x9e,
	0xc5, 0x2e, 0xc1, 0xe7, 0x42, 0xa2, 0x67, 0x05, 0x8e, 0x67, 0xa0, 0x27, 0x45, 0x86, 0xb7, 0x9b,
	0x7f, 0x1a, 0x66, 0x5d, 0x2d, 0x64, 0x45, 0x5d, 0x9e, 0x12, 0x3b, 0xf8, 0x9b, 0xfa, 0xfd, 0x64,
	0x41, 0x27, 0x5c, 0x02, 0x76, 0xa6, 0xc4, 0xce, 0xa4, 0x3e, 0xca, 0x69, 0x9d, 0xc1, 0x3c, 0x04,
	0x5b, 0xd3, 0x5a, 0xc7, 0x7f, 0xe2, 0xcd, 0xa3, 0x80, 0x2b, 0x10, 0x0a, 0x89, 0x56, 0x94, 0xcc,
	0x78, 0x05, 0x25, 0x9c, 0x63, 0xb0, 0x99, 0xd7, 0x56, 0x48, 0x79, 0x0c, 0xd7, 0x85, 0x44, 0x70,
	0x22, 0xab, 0x80, 0xf4, 0x01, 0xfc, 0x3a, 0xad, 0x68, 0x54, 0x0d, 0xdb, 0xed, 0x1c, 0x9b, 0x0e,
	0x76, 0x0f, 0xea, 0x46, 0xa5, 0x36, 0x59, 0xc6, 0x09, 0x79, 0x35, 0xdf, 0xb7, 0xa3, 0xc2, 0x84,
	0x9d, 0xae, 0x1f, 0xb1, 0x83, 0x3e, 0x3d, 0xef, 0x5d, 0x2f, 0xe9, 0x79, 0x7d, 0x7b, 0x8b, 0x2c,
	0x4f, 0xb2, 0x7d, 0x7b, 0x43, 0x8e, 0x2e, 0x78, 0x33, 0x7d, 0x16, 0x82, 0xe3, 0xf9, 0xe8, 0x1c,
	0xc7, 0xbe, 0xa8, 0xc2, 0x11, 0xa3, 0x59, 0x53, 0x69, 0x20, 0xdc, 0x13, 0x12, 0xed, 0xa4, 0xcc,
	0x66, 0x21, 0x8b, 0x3a, 0x5d, 0xde, 0x67, 0x49, 0x0d, 0x09, 0x1d, 0xf4, 0x92, 0x4c, 0x17, 0x60,
	0x3f, 0x7f, 0x9a, 0xdb, 0xe4, 0xf9, 0x01, 0xbe, 0x12, 0x12, 0xbd, 0x9c, 0x1c, 0xe7, 0xa6, 0x2c,
	0x85, 0x9d, 0x6d, 0x5b, 0x87, 0x4e, 0xc7, 0x20, 0xf1, 0xe4, 0x1c, 0x11, 0xa3, 0x9a, 0x32, 0xeb,
	0x69, 0xbe, 0xb3, 0x6d, 0xfe, 0x21, 0xfa, 0x44, 0x03, 0xd6, 0x1e, 0x7a, 0x01, 0xbd, 0x50, 0x3e,
	0x35, 0x32, 0xdd, 0x33, 0x1b, 0xb8, 0x96, 0x6d, 0xea, 0x59, 0xb8, 0xec, 0xa0, 0xc4, 0x2e, 0xe4,
	0x6e, 0x43, 0xa9, 0xeb, 0x84, 0xb6, 0xae, 0x98, 0x36, 0x8b, 0xb2, 0x77, 0x23, 0x75, 0x9d, 0xd0,
	0x5d, 0xf0, 0xd3, 0x8c, 0x2d, 0x3a, 0x16, 0xe7, 0x56, 0xdb, 0xc4, 0x70, 0xe2, 0x97, 0x89, 0xdc,
	0x67, 0x70, 0x47, 0x48, 0xf4, 0x22, 0xbb, 0x4a, 0xc7, 0x22, 0xab, 0x57, 0x01, 0x8d, 0xd2, 0xf7,
	0x94, 0xbf, 0x40, 0xa9, 0x70, 0x8f, 0xce, 0xc6, 0x23, 0xf8, 0x52, 0x48, 0xf4, 0x3c, 0xbd, 0x43,
	0x67, 0xc0, 0xe1, 0xdd, 0xcf, 0xff, 0xad, 0xdf, 0x39, 0x78, 0x73, 0xba, 0xef, 0xf9, 0x51, 0xf7,
	0xea, 0xac, 0x74, 0xce, 0x2f, 0xcb, 0x35, 0x3f, 0xa0, 0x15, 0x1e, 0xb0, 0x72, 0xc8, 0x7a, 0xd4,
	0x2f, 0x9b, 0x4d, 0x07, 0x93, 0xb7, 0xe5, 0xd1, 0xa5, 0x75, 0xcf, 0x63, 0x83, 0xf2, 0xe4, 0xca,
	0x7b, 0x76, 0x6f, 0xf4, 0x6c, 0xff, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0xc8, 0x70, 0x21,
	0x05, 0x0b, 0x00, 0x00,
}
