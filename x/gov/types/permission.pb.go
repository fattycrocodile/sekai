// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/gov/permission.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PermValue int32

const (
	// PERMISSION_ZERO is a no-op permission.
	PermZero PermValue = 0
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set Permissions to other actors.
	PermSetPermissions PermValue = 1
	// PERMISSION_CLAIM_VALIDATOR defines the permission that allows to Claim a validator Seat.
	PermClaimValidator PermValue = 2
	// PERMISSION_CLAIM_COUNCILOR defines the permission that allows to Claim a Councilor Seat.
	PermClaimCouncilor PermValue = 3
	// PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL defines the permission needed to create proposals for setting permissions.
	PermCreateSetPermissionsProposal PermValue = 4
	// PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set permissions.
	PermVoteSetPermissionProposal PermValue = 5
	//  PERMISSION_UPSERT_TOKEN_ALIAS
	PermUpsertTokenAlias PermValue = 6
	// PERMISSION_CHANGE_TX_FEE
	PermChangeTxFee PermValue = 7
	// PERMISSION_UPSERT_TOKEN_RATE
	PermUpsertTokenRate PermValue = 8
	// PERMISSION_UPSERT_ROLE makes possible to add, modify and assign roles.
	PermUpsertRole PermValue = 9
	// PERMISSION_CREATE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermCreateUpsertDataRegistryProposal PermValue = 10
	// PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermVoteUpsertDataRegistryProposal PermValue = 11
	// PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission needed to create proposals for setting network property.
	PermCreateSetNetworkPropertyProposal PermValue = 12
	// PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set network property.
	PermVoteSetNetworkPropertyProposal PermValue = 13
	// PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to create proposals for upsert token Alias.
	PermCreateUpsertTokenAliasProposal PermValue = 14
	// PERMISSION_VOTE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to vote proposals for upsert token.
	PermVoteUpsertTokenAliasProposal PermValue = 15
	// PERMISSION_CREATE_SET_POOR_NETWORK_MESSAGES defines the permission needed to create proposals for setting poor network messages
	PermCreateSetPoorNetworkMessagesProposal PermValue = 16
	// PERMISSION_VOTE_SET_POOR_NETWORK_MESSAGES_PROPOSAL defines the permission needed to vote proposals to set poor network messages
	PermVoteSetPoorNetworkMessagesProposal PermValue = 17
	// PERMISSION_CREATE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to create proposals for upsert token rate.
	PermCreateUpsertTokenRateProposal PermValue = 18
	// PERMISSION_VOTE_UPSERT_TOKEN_RATE_PROPOSAL defines the permission needed to vote proposals for upsert token rate.
	PermVoteUpsertTokenRateProposal PermValue = 19
	// PERMISSION_CREATE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to create a proposal to unjail a validator.
	PermCreateUnjailValidatorProposal PermValue = 20
	// PERMISSION_VOTE_UNJAIL_VALIDATOR_PROPOSAL defines the permission needed to vote a proposal to unjail a validator.
	PermVoteUnjailValidatorProposal PermValue = 21
	// PERMISSION_CREATE_CREATE_ROLE_PROPOSAL defines the permission needed to create a proposal to create a role.
	PermCreateRoleProposal PermValue = 22
	// PERMISSION_VOTE_CREATE_ROLE_PROPOSAL defines the permission needed to vote a proposal to create a role.
	PermVoteCreateRoleProposal PermValue = 23
	// PERMISSION_CREATE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to create a proposal to blacklist/whitelisted tokens
	PermCreateTokensWhiteBlackChangeProposal PermValue = 24
	// PERMISSION_VOTE_TOKENS_WHITE_BLACK_CHANGE_PROPOSAL defines the permission needed to vote on blacklist/whitelisted tokens proposal
	PermVoteTokensWhiteBlackChangeProposal PermValue = 25
	// PERMISSION_CREATE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to create a proposal to reset whole validator rank
	PermCreateResetWholeValidatorRankProposal PermValue = 26
	// PERMISSION_VOTE_RESET_WHOLE_VALIDATOR_RANK_PROPOSAL defines the permission needed to vote on reset whole validator rank proposal
	PermVoteResetWholeValidatorRankProposal PermValue = 27
	// PERMISSION_CREATE_SOFTWARE_UPGRADE_PROPOSAL defines the permission needed to create a proposal for software upgrade
	PermCreateSoftwareUpgradeProposal PermValue = 28
	// PERMISSION_SOFTWARE_UPGRADE_PROPOSAL defines the permission needed to vote on software upgrade proposal
	PermVoteSoftwareUpgradeProposal PermValue = 29
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

func init() { proto.RegisterFile("kira/gov/permission.proto", fileDescriptor_214168f8815c1062) }

var fileDescriptor_214168f8815c1062 = []byte{
	// 939 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0x4d, 0x53, 0xdb, 0x46,
	0x18, 0xc7, 0x4d, 0x9b, 0xa6, 0x64, 0x9b, 0x06, 0x55, 0x50, 0x42, 0xb6, 0xc1, 0x5d, 0x5a, 0x42,
	0x09, 0x49, 0xf0, 0xb4, 0x69, 0x3b, 0xd3, 0xe9, 0xa1, 0x23, 0xcc, 0x02, 0xae, 0x5f, 0xe4, 0x59,
	0xc9, 0xb8, 0x61, 0x26, 0xa3, 0x59, 0x60, 0x63, 0xab, 0x36, 0x5e, 0xcf, 0x4a, 0xe4, 0xe5, 0x1b,
	0x74, 0xf6, 0xd4, 0x2f, 0xb0, 0xa7, 0x7e, 0x99, 0x1e, 0x73, 0xec, 0xb1, 0x03, 0xdf, 0xa1, 0xe7,
	0x8e, 0x0c, 0x68, 0x25, 0x59, 0x36, 0x9c, 0xfc, 0xa2, 0x7d, 0x7e, 0xcf, 0x7f, 0xff, 0xcf, 0xcb,
	0x08, 0x3c, 0xe8, 0xf9, 0x82, 0x96, 0x3a, 0xfc, 0x75, 0x69, 0xc8, 0xc4, 0x89, 0x1f, 0x04, 0x3e,
	0x1f, 0x6c, 0x0e, 0x05, 0x0f, 0xb9, 0x39, 0x1b, 0x3d, 0xda, 0xec, 0xf0, 0xd7, 0x70, 0xa1, 0xc3,
	0x3b, 0x7c, 0xf4, 0x67, 0x29, 0xfa, 0x76, 0xf1, 0x7c, 0xe3, 0x3f, 0x13, 0xdc, 0x69, 0x32, 0x71,
	0xb2, 0x4f, 0xfb, 0xa7, 0xcc, 0x5c, 0x01, 0x73, 0x4d, 0x4c, 0xea, 0x15, 0xc7, 0xa9, 0xd8, 0x0d,
	0xef, 0x00, 0x13, 0xdb, 0x28, 0xc0, 0xbb, 0x52, 0xa1, 0xd9, 0xe8, 0xcc, 0x01, 0x13, 0xdc, 0xfc,
	0x11, 0xc0, 0xc4, 0x11, 0x07, 0xbb, 0x9e, 0xfe, 0xe9, 0x18, 0x33, 0x70, 0x51, 0x2a, 0x64, 0x46,
	0xa7, 0x1d, 0x16, 0x36, 0x63, 0x35, 0x41, 0x26, 0xae, 0x5c, 0xb3, 0x2a, 0x75, 0x6f, 0xdf, 0xaa,
	0x55, 0xb6, 0x2d, 0xd7, 0x26, 0xc6, 0x07, 0x3a, 0xae, 0xdc, 0xa7, 0x7e, 0x24, 0xc7, 0x3f, 0xa6,
	0x21, 0x17, 0xb9, 0x71, 0x65, 0xbb, 0xd5, 0x28, 0x57, 0x6a, 0x36, 0x31, 0x3e, 0xcc, 0xc4, 0x95,
	0xf9, 0xe9, 0xe0, 0xc8, 0xef, 0x73, 0x61, 0xba, 0x60, 0x23, 0x19, 0x47, 0xb0, 0xe5, 0xe2, 0xac,
	0x5c, 0xaf, 0x49, 0xec, 0xa6, 0xed, 0x58, 0x35, 0xe3, 0x16, 0x5c, 0x95, 0x0a, 0xa1, 0x11, 0x47,
	0x30, 0x1a, 0xb2, 0xb4, 0xfa, 0xa6, 0xe0, 0x43, 0x1e, 0xd0, 0xbe, 0x69, 0x83, 0xf5, 0x04, 0x75,
	0xdf, 0x9e, 0xc6, 0xfc, 0x08, 0xae, 0x48, 0x85, 0x96, 0x47, 0xee, 0xf2, 0x0c, 0x31, 0x06, 0xfe,
	0x0c, 0x96, 0x13, 0xc0, 0x56, 0xd3, 0xc1, 0xc4, 0xf5, 0x5c, 0xbb, 0x8a, 0x1b, 0x9e, 0x55, 0xab,
	0x58, 0x8e, 0x71, 0x1b, 0x2e, 0x49, 0x85, 0x16, 0xa2, 0xd0, 0xd6, 0x30, 0x60, 0x22, 0x74, 0x79,
	0x8f, 0x0d, 0xac, 0xbe, 0x4f, 0x03, 0xf3, 0x5b, 0xb0, 0x94, 0xbc, 0xe3, 0x9e, 0xd5, 0xd8, 0xc5,
	0x9e, 0xfb, 0x9b, 0xb7, 0x83, 0xb1, 0xf1, 0x31, 0x9c, 0x97, 0x0a, 0xcd, 0x8d, 0x6e, 0xd4, 0xa5,
	0x83, 0x0e, 0x73, 0xdf, 0xee, 0x30, 0x66, 0xfe, 0x04, 0x1e, 0x4e, 0xca, 0x47, 0x2c, 0x17, 0x1b,
	0xb3, 0xf0, 0xbe, 0x54, 0x68, 0x3e, 0x93, 0x8e, 0xd0, 0x90, 0x99, 0x9b, 0x60, 0x71, 0x3c, 0x94,
	0xd8, 0x35, 0x6c, 0xdc, 0x81, 0xa6, 0x54, 0xe8, 0x9e, 0x0e, 0x22, 0xbc, 0xcf, 0xcc, 0x97, 0xa0,
	0x34, 0x5e, 0x81, 0xcb, 0xb0, 0x6d, 0xcb, 0xb5, 0x3c, 0x82, 0x77, 0x2b, 0x8e, 0x4b, 0x5e, 0x68,
	0xcb, 0x00, 0x5c, 0x97, 0x0a, 0xad, 0xea, 0x32, 0x5c, 0xe0, 0xb6, 0x69, 0x48, 0x09, 0xeb, 0xf8,
	0x41, 0x28, 0xde, 0xc5, 0xce, 0xbd, 0x00, 0xcf, 0xb2, 0xa5, 0x98, 0x0e, 0xff, 0x04, 0xae, 0x49,
	0x85, 0xbe, 0xba, 0xaa, 0xc7, 0x14, 0x74, 0xae, 0xf2, 0xa8, 0xce, 0x0d, 0xec, 0xb6, 0x6d, 0x52,
	0x1d, 0x31, 0x31, 0x71, 0x13, 0xf0, 0xbb, 0x59, 0xe5, 0x0e, 0x0b, 0x1b, 0x2c, 0x7c, 0xc3, 0x45,
	0x2f, 0xc2, 0x32, 0x11, 0x4e, 0x55, 0x3e, 0x1d, 0xfe, 0x69, 0x5a, 0xf9, 0x8d, 0xd1, 0x69, 0xcf,
	0x13, 0x5d, 0xa5, 0xd1, 0xf7, 0x34, 0x3a, 0xe9, 0xb8, 0x6e, 0xb2, 0x18, 0xdd, 0x02, 0x4f, 0x26,
	0xf8, 0x9d, 0x0b, 0x9e, 0xd3, 0x13, 0xa5, 0xdd, 0xce, 0xc1, 0xbe, 0x4c, 0x61, 0x93, 0x73, 0x6a,
	0xdb, 0x24, 0xf6, 0xa4, 0x8e, 0x1d, 0xc7, 0xda, 0xc5, 0x8e, 0x61, 0xc0, 0xa7, 0x52, 0xa1, 0xf5,
	0xf4, 0xa0, 0x72, 0x2e, 0x2e, 0x0d, 0xa9, 0xb3, 0x20, 0xa0, 0x1d, 0xa6, 0xf1, 0x87, 0xe0, 0xbb,
	0xdc, 0x81, 0xcd, 0x83, 0x6b, 0xf1, 0x9f, 0xc1, 0x0d, 0xa9, 0xd0, 0x5a, 0x72, 0x74, 0xa7, 0xe4,
	0x68, 0x83, 0xa7, 0xd7, 0x98, 0x1e, 0x8d, 0x96, 0xa6, 0x9b, 0xf0, 0x91, 0x54, 0x68, 0x25, 0xd7,
	0xf3, 0x68, 0xd2, 0x62, 0xb0, 0x93, 0xda, 0x61, 0xe3, 0x96, 0xa7, 0xb1, 0xf3, 0xf0, 0x6b, 0xa9,
	0xd0, 0x97, 0x39, 0x8e, 0xa7, 0xa0, 0xfb, 0x79, 0x86, 0xb7, 0x1a, 0xbf, 0x5a, 0x95, 0x9a, 0x5e,
	0xc8, 0x9a, 0xba, 0x30, 0x26, 0x76, 0xf0, 0x3b, 0xf5, 0xfb, 0xf1, 0x82, 0x8e, 0xb9, 0x04, 0x3c,
	0x1e, 0x13, 0x3b, 0x91, 0xfa, 0x79, 0x46, 0xeb, 0x04, 0xe6, 0x0e, 0x58, 0x1b, 0xd7, 0x7a, 0xf9,
	0x11, 0x6d, 0x1e, 0x0d, 0x5c, 0x84, 0x50, 0x2a, 0xb4, 0xa8, 0x65, 0x46, 0x2b, 0x28, 0xe6, 0xec,
	0x81, 0xd5, 0xac, 0xb6, 0x5c, 0xca, 0x7d, 0x58, 0x94, 0x0a, 0xc1, 0x2b, 0x59, 0x39, 0xa4, 0x57,
	0xe0, 0xfb, 0x71, 0x45, 0xa3, 0x6a, 0x38, 0x5e, 0x7b, 0xaf, 0xe2, 0x62, 0x6f, 0xab, 0x66, 0x95,
	0xab, 0x57, 0xcb, 0x38, 0x26, 0x2f, 0x65, 0xfb, 0x76, 0x54, 0x98, 0xa0, 0xdd, 0xf5, 0x43, 0xb6,
	0xd5, 0xa7, 0x47, 0xbd, 0x8b, 0x25, 0x3d, 0xad, 0x6f, 0x6f, 0x90, 0xe5, 0x41, 0xba, 0x6f, 0xaf,
	0xc9, 0xd1, 0x05, 0x3f, 0x8c, 0xdf, 0x85, 0xe0, 0x68, 0x3e, 0xda, 0x7b, 0x91, 0x2f, 0xba, 0x70,
	0xc4, 0x6a, 0x54, 0x75, 0x1a, 0x08, 0x9f, 0x49, 0x85, 0x1e, 0x27, 0xcc, 0x66, 0x01, 0x0b, 0xdb,
	0x5d, 0xde, 0x67, 0x71, 0x0d, 0x09, 0x1d, 0xf4, 0xe2, 0x4c, 0xc7, 0xe0, 0x79, 0xf6, 0x36, 0x37,
	0xc9, 0xf3, 0x05, 0x7c, 0x22, 0x15, 0xfa, 0xe6, 0xea, 0x3a, 0xd7, 0x65, 0xc9, 0xed, 0x6c, 0xc7,
	0xde, 0x71, 0xdb, 0x16, 0x89, 0x26, 0x67, 0x97, 0x58, 0xdb, 0x09, 0xb3, 0x1e, 0x66, 0x3b, 0xdb,
	0xe1, 0xaf, 0xc2, 0x37, 0x54, 0xb0, 0xd6, 0xb0, 0x23, 0xe8, 0xb1, 0xf6, 0xa9, 0x9e, 0xea, 0x9e,
	0xc9, 0xc0, 0xe5, 0x74, 0x53, 0x4f, 0xc0, 0xc1, 0x5b, 0x7f, 0xfc, 0x55, 0x2c, 0x6c, 0xfd, 0xf2,
	0xf7, 0x59, 0x71, 0xe6, 0xfd, 0x59, 0x71, 0xe6, 0xdf, 0xb3, 0xe2, 0xcc, 0x9f, 0xe7, 0xc5, 0xc2,
	0xfb, 0xf3, 0x62, 0xe1, 0x9f, 0xf3, 0x62, 0xe1, 0xe0, 0x51, 0xc7, 0x0f, 0xbb, 0xa7, 0x87, 0x9b,
	0x47, 0xfc, 0xa4, 0x54, 0xf5, 0x05, 0x2d, 0x73, 0xc1, 0x4a, 0x01, 0xeb, 0x51, 0xbf, 0xf4, 0x76,
	0xf4, 0x92, 0x17, 0xbe, 0x1b, 0xb2, 0xe0, 0xf0, 0xf6, 0xe8, 0x05, 0xee, 0xf9, 0xff, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x88, 0x7c, 0xfc, 0xbb, 0xfd, 0x09, 0x00, 0x00,
}
