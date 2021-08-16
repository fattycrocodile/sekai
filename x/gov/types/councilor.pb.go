// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/gov/councilor.proto

package types

import (
	bytes "bytes"
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

type MsgClaimCouncilor struct {
	Moniker string                                        `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
}

func (m *MsgClaimCouncilor) Reset()         { *m = MsgClaimCouncilor{} }
func (m *MsgClaimCouncilor) String() string { return proto.CompactTextString(m) }
func (*MsgClaimCouncilor) ProtoMessage()    {}
func (*MsgClaimCouncilor) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac0b13545050495, []int{0}
}
func (m *MsgClaimCouncilor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimCouncilor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimCouncilor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimCouncilor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimCouncilor.Merge(m, src)
}
func (m *MsgClaimCouncilor) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimCouncilor) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimCouncilor.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimCouncilor proto.InternalMessageInfo

func (m *MsgClaimCouncilor) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *MsgClaimCouncilor) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

type Councilor struct {
	Moniker string                                        `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
}

func (m *Councilor) Reset()         { *m = Councilor{} }
func (m *Councilor) String() string { return proto.CompactTextString(m) }
func (*Councilor) ProtoMessage()    {}
func (*Councilor) Descriptor() ([]byte, []int) {
	return fileDescriptor_cac0b13545050495, []int{1}
}
func (m *Councilor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Councilor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Councilor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Councilor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Councilor.Merge(m, src)
}
func (m *Councilor) XXX_Size() int {
	return m.Size()
}
func (m *Councilor) XXX_DiscardUnknown() {
	xxx_messageInfo_Councilor.DiscardUnknown(m)
}

var xxx_messageInfo_Councilor proto.InternalMessageInfo

func (m *Councilor) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Councilor) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgClaimCouncilor)(nil), "kira.gov.MsgClaimCouncilor")
	proto.RegisterType((*Councilor)(nil), "kira.gov.Councilor")
}

func init() { proto.RegisterFile("kira/gov/councilor.proto", fileDescriptor_cac0b13545050495) }

var fileDescriptor_cac0b13545050495 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xce, 0x2c, 0x4a,
	0xd4, 0x4f, 0xcf, 0x2f, 0xd3, 0x4f, 0xce, 0x2f, 0xcd, 0x4b, 0xce, 0xcc, 0xc9, 0x2f, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xc9, 0xe8, 0xa5, 0xe7, 0x97, 0x49, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0x05, 0xf5, 0x41, 0x2c, 0x88, 0xbc, 0xd2, 0x24, 0x46, 0x2e, 0x41, 0xdf, 0xe2,
	0x74, 0xe7, 0x9c, 0xc4, 0xcc, 0x5c, 0x67, 0x98, 0x5e, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xfc, 0xbc,
	0xcc, 0xec, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x28, 0x96, 0x8b,
	0x3d, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x49, 0x81, 0x51, 0x83, 0xc7, 0xc9, 0xf9,
	0xd3, 0x3d, 0x79, 0xbe, 0xca, 0xc4, 0xdc, 0x1c, 0x2b, 0x25, 0xa8, 0x84, 0xd2, 0xaf, 0x7b, 0xf2,
	0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc9, 0xf9, 0xc5, 0xb9,
	0xf9, 0xc5, 0x50, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x58, 0xcf, 0x31,
	0x39, 0xd9, 0x11, 0xa2, 0x23, 0x08, 0x66, 0xa6, 0x15, 0xcb, 0x8b, 0x05, 0xf2, 0x8c, 0x4a, 0x2d,
	0x8c, 0x5c, 0x9c, 0x03, 0xef, 0x18, 0x27, 0xfb, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0x52, 0x45, 0x32, 0xd1, 0x3b, 0xb3, 0x28, 0xd1, 0x39, 0xbf, 0x28, 0x55, 0xbf, 0x38,
	0x35, 0x3b, 0x31, 0x53, 0xbf, 0x02, 0x1c, 0x0d, 0x60, 0x43, 0x93, 0xd8, 0xc0, 0x61, 0x6c, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xae, 0xd6, 0xea, 0xc6, 0x9f, 0x01, 0x00, 0x00,
}

func (this *MsgClaimCouncilor) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgClaimCouncilor)
	if !ok {
		that2, ok := that.(MsgClaimCouncilor)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Moniker != that1.Moniker {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	return true
}
func (m *MsgClaimCouncilor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimCouncilor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimCouncilor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Councilor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Councilor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Councilor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintCouncilor(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCouncilor(dAtA []byte, offset int, v uint64) int {
	offset -= sovCouncilor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgClaimCouncilor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	return n
}

func (m *Councilor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCouncilor(uint64(l))
	}
	return n
}

func sovCouncilor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCouncilor(x uint64) (n int) {
	return sovCouncilor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgClaimCouncilor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCouncilor
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
			return fmt.Errorf("proto: MsgClaimCouncilor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimCouncilor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCouncilor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCouncilor
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
func (m *Councilor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCouncilor
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
			return fmt.Errorf("proto: Councilor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Councilor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCouncilor
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
				return ErrInvalidLengthCouncilor
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCouncilor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCouncilor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCouncilor
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
func skipCouncilor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCouncilor
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
					return 0, ErrIntOverflowCouncilor
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
					return 0, ErrIntOverflowCouncilor
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
				return 0, ErrInvalidLengthCouncilor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCouncilor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCouncilor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCouncilor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCouncilor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCouncilor = fmt.Errorf("proto: unexpected end of group")
)
