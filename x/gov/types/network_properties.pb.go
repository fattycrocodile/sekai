// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network_properties.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type NetworkProperty int32

const (
	MinTxFee                 NetworkProperty = 0
	MaxTxFee                 NetworkProperty = 1
	VoteQuorum               NetworkProperty = 2
	ProposalEndTime          NetworkProperty = 3
	ProposalEnactmentTime    NetworkProperty = 4
	EnableForeignFeePayments NetworkProperty = 5
	PoorNetworkMaxBankSend   NetworkProperty = 6
)

var NetworkProperty_name = map[int32]string{
	0: "MIN_TX_FEE",
	1: "MAX_TX_FEE",
	2: "VOTE_QUORUM",
	3: "PROPOSAL_END_TIME",
	4: "PROPOSAL_ENACTMENT_TIME",
	5: "ENABLE_FOREIGN_TX_FEE_PAYMENTS",
	6: "POOR_NETWORK_MAX_BANK_SEND",
}

var NetworkProperty_value = map[string]int32{
	"MIN_TX_FEE":                     0,
	"MAX_TX_FEE":                     1,
	"VOTE_QUORUM":                    2,
	"PROPOSAL_END_TIME":              3,
	"PROPOSAL_ENACTMENT_TIME":        4,
	"ENABLE_FOREIGN_TX_FEE_PAYMENTS": 5,
	"POOR_NETWORK_MAX_BANK_SEND":     6,
}

func (x NetworkProperty) String() string {
	return proto.EnumName(NetworkProperty_name, int32(x))
}

func (NetworkProperty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_afa35a4ab1e9e2c2, []int{0}
}

type MsgSetNetworkProperties struct {
	NetworkProperties *NetworkProperties                            `protobuf:"bytes,1,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	Proposer          github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty"`
}

func (m *MsgSetNetworkProperties) Reset()         { *m = MsgSetNetworkProperties{} }
func (m *MsgSetNetworkProperties) String() string { return proto.CompactTextString(m) }
func (*MsgSetNetworkProperties) ProtoMessage()    {}
func (*MsgSetNetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa35a4ab1e9e2c2, []int{0}
}
func (m *MsgSetNetworkProperties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetNetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetNetworkProperties.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetNetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetNetworkProperties.Merge(m, src)
}
func (m *MsgSetNetworkProperties) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetNetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetNetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetNetworkProperties proto.InternalMessageInfo

func (m *MsgSetNetworkProperties) GetNetworkProperties() *NetworkProperties {
	if m != nil {
		return m.NetworkProperties
	}
	return nil
}

func (m *MsgSetNetworkProperties) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type NetworkProperties struct {
	MinTxFee                 uint64 `protobuf:"varint,1,opt,name=min_tx_fee,json=minTxFee,proto3" json:"min_tx_fee,omitempty"`
	MaxTxFee                 uint64 `protobuf:"varint,2,opt,name=max_tx_fee,json=maxTxFee,proto3" json:"max_tx_fee,omitempty"`
	VoteQuorum               uint64 `protobuf:"varint,3,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	ProposalEndTime          uint64 `protobuf:"varint,4,opt,name=proposal_end_time,json=proposalEndTime,proto3" json:"proposal_end_time,omitempty"`
	ProposalEnactmentTime    uint64 `protobuf:"varint,5,opt,name=proposal_enactment_time,json=proposalEnactmentTime,proto3" json:"proposal_enactment_time,omitempty"`
	EnableForeignFeePayments bool   `protobuf:"varint,6,opt,name=enable_foreign_fee_payments,json=enableForeignFeePayments,proto3" json:"enable_foreign_fee_payments,omitempty"`
	// The rank property is a long term statistics implying the "longest" streak that validator ever achieved,
	// it can be expressed as rank = MAX(rank, streak).
	// Under certain circumstances we should however decrease the rank of the validator.
	// If the mischance property is incremented, the rank should be decremented by X (default 10), that is rank = MAX(rank - X, 0).
	// Every time node status changes to inactive the rank should be divided by 2, that is rank = FLOOR(rank / 2)
	// The streak and rank will enable governance to judge real life performance of validators on the mainnet or testnet, and potentially propose eviction of the weakest and least reliable operators.
	MischanceRankDecreaseAmount uint64 `protobuf:"varint,7,opt,name=mischance_rank_decrease_amount,json=mischanceRankDecreaseAmount,proto3" json:"mischance_rank_decrease_amount,omitempty"`
	InactiveRankDecreasePercent uint64 `protobuf:"varint,8,opt,name=inactive_rank_decrease_percent,json=inactiveRankDecreasePercent,proto3" json:"inactive_rank_decrease_percent,omitempty"`
	MinValidators               uint64 `protobuf:"varint,9,opt,name=min_validators,json=minValidators,proto3" json:"min_validators,omitempty"`
	PoorNetworkMaxBankSend      uint64 `protobuf:"varint,10,opt,name=poor_network_max_bank_send,json=poorNetworkMaxBankSend,proto3" json:"poor_network_max_bank_send,omitempty"`
}

func (m *NetworkProperties) Reset()         { *m = NetworkProperties{} }
func (m *NetworkProperties) String() string { return proto.CompactTextString(m) }
func (*NetworkProperties) ProtoMessage()    {}
func (*NetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa35a4ab1e9e2c2, []int{1}
}
func (m *NetworkProperties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkProperties.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkProperties.Merge(m, src)
}
func (m *NetworkProperties) XXX_Size() int {
	return m.Size()
}
func (m *NetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkProperties proto.InternalMessageInfo

func (m *NetworkProperties) GetMinTxFee() uint64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

func (m *NetworkProperties) GetMaxTxFee() uint64 {
	if m != nil {
		return m.MaxTxFee
	}
	return 0
}

func (m *NetworkProperties) GetVoteQuorum() uint64 {
	if m != nil {
		return m.VoteQuorum
	}
	return 0
}

func (m *NetworkProperties) GetProposalEndTime() uint64 {
	if m != nil {
		return m.ProposalEndTime
	}
	return 0
}

func (m *NetworkProperties) GetProposalEnactmentTime() uint64 {
	if m != nil {
		return m.ProposalEnactmentTime
	}
	return 0
}

func (m *NetworkProperties) GetEnableForeignFeePayments() bool {
	if m != nil {
		return m.EnableForeignFeePayments
	}
	return false
}

func (m *NetworkProperties) GetMischanceRankDecreaseAmount() uint64 {
	if m != nil {
		return m.MischanceRankDecreaseAmount
	}
	return 0
}

func (m *NetworkProperties) GetInactiveRankDecreasePercent() uint64 {
	if m != nil {
		return m.InactiveRankDecreasePercent
	}
	return 0
}

func (m *NetworkProperties) GetMinValidators() uint64 {
	if m != nil {
		return m.MinValidators
	}
	return 0
}

func (m *NetworkProperties) GetPoorNetworkMaxBankSend() uint64 {
	if m != nil {
		return m.PoorNetworkMaxBankSend
	}
	return 0
}

func init() {
	proto.RegisterEnum("kira.gov.NetworkProperty", NetworkProperty_name, NetworkProperty_value)
	proto.RegisterType((*MsgSetNetworkProperties)(nil), "kira.gov.MsgSetNetworkProperties")
	proto.RegisterType((*NetworkProperties)(nil), "kira.gov.NetworkProperties")
}

func init() { proto.RegisterFile("network_properties.proto", fileDescriptor_afa35a4ab1e9e2c2) }

var fileDescriptor_afa35a4ab1e9e2c2 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x4f, 0x4f, 0xdb, 0x48,
	0x14, 0x8f, 0x43, 0x60, 0xb3, 0x03, 0x0b, 0xc1, 0xbb, 0x80, 0xd7, 0x20, 0x63, 0x21, 0x21, 0x21,
	0x24, 0x12, 0xed, 0xae, 0xb4, 0x07, 0xa4, 0xd5, 0xae, 0x03, 0xce, 0x8a, 0x82, 0xff, 0xe0, 0x18,
	0x4a, 0x7b, 0x19, 0x4d, 0xec, 0x47, 0xb0, 0x82, 0x67, 0xd2, 0xb1, 0x93, 0x86, 0x5b, 0x8f, 0x55,
	0x4e, 0xfd, 0x02, 0x39, 0xf5, 0x2b, 0xf4, 0x43, 0xf4, 0xc8, 0xa9, 0xea, 0xa9, 0xaa, 0xe0, 0x5b,
	0xf4, 0x54, 0xd9, 0x4e, 0xd2, 0xd0, 0xd0, 0x93, 0xad, 0xf7, 0xfb, 0xf3, 0xde, 0xf8, 0xfd, 0xc6,
	0x48, 0xa2, 0x10, 0xbf, 0x64, 0xbc, 0x85, 0xdb, 0x9c, 0xb5, 0x81, 0xc7, 0x01, 0x44, 0xe5, 0x36,
	0x67, 0x31, 0x13, 0x8b, 0xad, 0x80, 0x93, 0x72, 0x93, 0x75, 0xe5, 0xdf, 0x9a, 0xac, 0xc9, 0xd2,
	0x62, 0x25, 0x79, 0xcb, 0xf0, 0xad, 0x77, 0x02, 0x5a, 0x33, 0xa2, 0x66, 0x1d, 0x62, 0x33, 0xb3,
	0xb0, 0xc7, 0x0e, 0xe2, 0x13, 0x24, 0x4e, 0xfb, 0x4a, 0x82, 0x2a, 0xec, 0xcc, 0xff, 0xb9, 0x5e,
	0x1e, 0x19, 0x97, 0xa7, 0x84, 0xce, 0x32, 0x9d, 0xf2, 0x32, 0x50, 0x31, 0xf1, 0x60, 0x11, 0x70,
	0x29, 0xaf, 0x0a, 0x3b, 0x0b, 0xd5, 0x3f, 0xbe, 0x7c, 0xda, 0xdc, 0x6b, 0x06, 0xf1, 0x55, 0xa7,
	0x51, 0xf6, 0x58, 0x58, 0xf1, 0x58, 0x14, 0xb2, 0x68, 0xf8, 0xd8, 0x8b, 0xfc, 0x56, 0x25, 0xbe,
	0x69, 0x43, 0x54, 0xd6, 0x3c, 0x4f, 0xf3, 0x7d, 0x0e, 0x51, 0xe4, 0x8c, 0x2d, 0xb6, 0x5e, 0x15,
	0xd0, 0xf2, 0xf4, 0xc0, 0x1b, 0x08, 0x85, 0x01, 0xc5, 0x71, 0x0f, 0x5f, 0x02, 0xa4, 0x83, 0x16,
	0x9c, 0x62, 0x18, 0x50, 0xb7, 0x57, 0x03, 0x48, 0x51, 0xd2, 0x1b, 0xa1, 0xf9, 0x21, 0x4a, 0x7a,
	0x19, 0xba, 0x89, 0xe6, 0xbb, 0x2c, 0x06, 0xfc, 0xa2, 0xc3, 0x78, 0x27, 0x94, 0x66, 0x52, 0x18,
	0x25, 0xa5, 0xd3, 0xb4, 0x22, 0xee, 0xa2, 0xe5, 0xac, 0x3d, 0xb9, 0xc6, 0x40, 0x7d, 0x1c, 0x07,
	0x21, 0x48, 0x85, 0x94, 0xb6, 0x34, 0x02, 0x74, 0xea, 0xbb, 0x41, 0x08, 0xe2, 0xdf, 0x68, 0x6d,
	0x82, 0x4b, 0xbc, 0x38, 0x04, 0x1a, 0x67, 0x8a, 0xd9, 0x54, 0xb1, 0xf2, 0x4d, 0x31, 0x44, 0x53,
	0xdd, 0x3f, 0x68, 0x1d, 0x28, 0x69, 0x5c, 0x03, 0xbe, 0x64, 0x1c, 0x82, 0x26, 0x4d, 0x46, 0xc5,
	0x6d, 0x72, 0x93, 0x30, 0x22, 0x69, 0x4e, 0x15, 0x76, 0x8a, 0x8e, 0x94, 0x51, 0x6a, 0x19, 0xa3,
	0x06, 0x60, 0x0f, 0x71, 0xf1, 0x00, 0x29, 0x61, 0x10, 0x79, 0x57, 0x84, 0x7a, 0x80, 0x39, 0xa1,
	0x2d, 0xec, 0x83, 0xc7, 0x81, 0x44, 0x80, 0x49, 0xc8, 0x3a, 0x34, 0x96, 0x7e, 0x4a, 0xbb, 0xaf,
	0x8f, 0x59, 0x0e, 0xa1, 0xad, 0xc3, 0x21, 0x47, 0x4b, 0x29, 0x89, 0x49, 0x90, 0x0c, 0x15, 0x74,
	0xbf, 0xf7, 0x68, 0x03, 0xf7, 0x80, 0xc6, 0x52, 0x31, 0x33, 0x19, 0xb1, 0x26, 0x3d, 0xec, 0x8c,
	0x22, 0x6e, 0xa3, 0xc5, 0x64, 0x13, 0x5d, 0x72, 0x1d, 0xf8, 0x24, 0x66, 0x3c, 0x92, 0x7e, 0x4e,
	0x45, 0xbf, 0x84, 0x01, 0x3d, 0x1f, 0x17, 0xc5, 0x7d, 0x24, 0xb7, 0x19, 0xe3, 0x78, 0x14, 0xb3,
	0x64, 0x3f, 0x8d, 0xa4, 0x67, 0x04, 0xd4, 0x97, 0x50, 0x2a, 0x59, 0x4d, 0x18, 0xc3, 0x5d, 0x1b,
	0xa4, 0x57, 0x25, 0xb4, 0x55, 0x07, 0xea, 0xef, 0x7e, 0xc8, 0xa3, 0xa5, 0x87, 0x11, 0xb8, 0x49,
	0x56, 0x6c, 0x1c, 0x99, 0xd8, 0xbd, 0xc0, 0x35, 0x5d, 0x2f, 0xe5, 0xe4, 0x85, 0xfe, 0x40, 0x2d,
	0x1a, 0x13, 0x01, 0x30, 0xb4, 0x8b, 0x11, 0x2a, 0x0c, 0xd1, 0x89, 0x00, 0x9c, 0x5b, 0xae, 0x8e,
	0x4f, 0xcf, 0x2c, 0xe7, 0xcc, 0x28, 0xe5, 0xe5, 0xc5, 0xfe, 0x40, 0x45, 0xe7, 0x0f, 0x02, 0x60,
	0x3b, 0x96, 0x6d, 0xd5, 0xb5, 0x13, 0xac, 0x9b, 0x87, 0xd8, 0x3d, 0x32, 0xf4, 0xd2, 0x8c, 0xfc,
	0x6b, 0x7f, 0xa0, 0x2e, 0xd9, 0xd3, 0x01, 0x98, 0xe0, 0x6a, 0x07, 0xae, 0xa1, 0x9b, 0x6e, 0xa6,
	0x28, 0xc8, 0xbf, 0xf7, 0x07, 0xea, 0x8a, 0xfd, 0x68, 0x00, 0xfe, 0x43, 0x8a, 0x6e, 0x6a, 0xd5,
	0x13, 0x1d, 0xd7, 0x2c, 0x47, 0x3f, 0xfa, 0x7f, 0x74, 0x16, 0x6c, 0x6b, 0xcf, 0x12, 0x8b, 0x7a,
	0x69, 0x56, 0xde, 0xe8, 0x0f, 0x54, 0x49, 0xff, 0x51, 0x06, 0xf6, 0x91, 0x6c, 0x5b, 0x96, 0x83,
	0x4d, 0xdd, 0x7d, 0x6a, 0x39, 0xc7, 0x38, 0x39, 0x71, 0x55, 0x33, 0x8f, 0x71, 0x5d, 0x37, 0x0f,
	0x4b, 0x73, 0xb2, 0xdc, 0x1f, 0xa8, 0xab, 0xf6, 0xa3, 0x9f, 0x54, 0x2e, 0xbc, 0x7e, 0xab, 0xe4,
	0xaa, 0xff, 0xbe, 0xbf, 0x53, 0x84, 0xdb, 0x3b, 0x45, 0xf8, 0x7c, 0xa7, 0x08, 0x6f, 0xee, 0x95,
	0xdc, 0xed, 0xbd, 0x92, 0xfb, 0x78, 0xaf, 0xe4, 0x9e, 0x6f, 0x4f, 0x5c, 0xd7, 0xe3, 0x80, 0x93,
	0x03, 0xc6, 0xa1, 0x12, 0x41, 0x8b, 0x04, 0x95, 0x5e, 0xa5, 0xc9, 0xba, 0xd9, 0x8d, 0x6d, 0xcc,
	0xa5, 0xbf, 0x96, 0xbf, 0xbe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xa7, 0xf6, 0xf6, 0x96, 0x04,
	0x00, 0x00,
}

func (m *MsgSetNetworkProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetNetworkProperties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetNetworkProperties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintNetworkProperties(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if m.NetworkProperties != nil {
		{
			size, err := m.NetworkProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNetworkProperties(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NetworkProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkProperties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkProperties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoorNetworkMaxBankSend != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.PoorNetworkMaxBankSend))
		i--
		dAtA[i] = 0x50
	}
	if m.MinValidators != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MinValidators))
		i--
		dAtA[i] = 0x48
	}
	if m.InactiveRankDecreasePercent != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.InactiveRankDecreasePercent))
		i--
		dAtA[i] = 0x40
	}
	if m.MischanceRankDecreaseAmount != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MischanceRankDecreaseAmount))
		i--
		dAtA[i] = 0x38
	}
	if m.EnableForeignFeePayments {
		i--
		if m.EnableForeignFeePayments {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.ProposalEnactmentTime != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.ProposalEnactmentTime))
		i--
		dAtA[i] = 0x28
	}
	if m.ProposalEndTime != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.ProposalEndTime))
		i--
		dAtA[i] = 0x20
	}
	if m.VoteQuorum != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.VoteQuorum))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxTxFee != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MaxTxFee))
		i--
		dAtA[i] = 0x10
	}
	if m.MinTxFee != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MinTxFee))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNetworkProperties(dAtA []byte, offset int, v uint64) int {
	offset -= sovNetworkProperties(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetNetworkProperties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NetworkProperties != nil {
		l = m.NetworkProperties.Size()
		n += 1 + l + sovNetworkProperties(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovNetworkProperties(uint64(l))
	}
	return n
}

func (m *NetworkProperties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinTxFee != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MinTxFee))
	}
	if m.MaxTxFee != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MaxTxFee))
	}
	if m.VoteQuorum != 0 {
		n += 1 + sovNetworkProperties(uint64(m.VoteQuorum))
	}
	if m.ProposalEndTime != 0 {
		n += 1 + sovNetworkProperties(uint64(m.ProposalEndTime))
	}
	if m.ProposalEnactmentTime != 0 {
		n += 1 + sovNetworkProperties(uint64(m.ProposalEnactmentTime))
	}
	if m.EnableForeignFeePayments {
		n += 2
	}
	if m.MischanceRankDecreaseAmount != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MischanceRankDecreaseAmount))
	}
	if m.InactiveRankDecreasePercent != 0 {
		n += 1 + sovNetworkProperties(uint64(m.InactiveRankDecreasePercent))
	}
	if m.MinValidators != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MinValidators))
	}
	if m.PoorNetworkMaxBankSend != 0 {
		n += 1 + sovNetworkProperties(uint64(m.PoorNetworkMaxBankSend))
	}
	return n
}

func sovNetworkProperties(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNetworkProperties(x uint64) (n int) {
	return sovNetworkProperties(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetNetworkProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkProperties
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetNetworkProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetNetworkProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkProperties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NetworkProperties == nil {
				m.NetworkProperties = &NetworkProperties{}
			}
			if err := m.NetworkProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkProperties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NetworkProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkProperties
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTxFee", wireType)
			}
			m.MinTxFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTxFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTxFee", wireType)
			}
			m.MaxTxFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTxFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteQuorum", wireType)
			}
			m.VoteQuorum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteQuorum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalEndTime", wireType)
			}
			m.ProposalEndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalEndTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalEnactmentTime", wireType)
			}
			m.ProposalEnactmentTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalEnactmentTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableForeignFeePayments", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableForeignFeePayments = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MischanceRankDecreaseAmount", wireType)
			}
			m.MischanceRankDecreaseAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MischanceRankDecreaseAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveRankDecreasePercent", wireType)
			}
			m.InactiveRankDecreasePercent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InactiveRankDecreasePercent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinValidators", wireType)
			}
			m.MinValidators = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinValidators |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoorNetworkMaxBankSend", wireType)
			}
			m.PoorNetworkMaxBankSend = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoorNetworkMaxBankSend |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkProperties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNetworkProperties(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkProperties
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNetworkProperties
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNetworkProperties
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNetworkProperties
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNetworkProperties        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkProperties          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNetworkProperties = fmt.Errorf("proto: unexpected end of group")
)
