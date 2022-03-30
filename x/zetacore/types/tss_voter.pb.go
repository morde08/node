// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetacore/tss_voter.proto

package types

import (
	fmt "fmt"
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

type TSSVoter struct {
	Creator         string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index           string   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Chain           string   `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Address         string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Pubkey          string   `protobuf:"bytes,5,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signers         []string `protobuf:"bytes,6,rep,name=signers,proto3" json:"signers,omitempty"`
	FinalizedHeight uint64   `protobuf:"varint,7,opt,name=finalizedHeight,proto3" json:"finalizedHeight,omitempty"`
}

func (m *TSSVoter) Reset()         { *m = TSSVoter{} }
func (m *TSSVoter) String() string { return proto.CompactTextString(m) }
func (*TSSVoter) ProtoMessage()    {}
func (*TSSVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e88278440279c17, []int{0}
}
func (m *TSSVoter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TSSVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TSSVoter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TSSVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSSVoter.Merge(m, src)
}
func (m *TSSVoter) XXX_Size() int {
	return m.Size()
}
func (m *TSSVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_TSSVoter.DiscardUnknown(m)
}

var xxx_messageInfo_TSSVoter proto.InternalMessageInfo

func (m *TSSVoter) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TSSVoter) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *TSSVoter) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *TSSVoter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TSSVoter) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *TSSVoter) GetSigners() []string {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *TSSVoter) GetFinalizedHeight() uint64 {
	if m != nil {
		return m.FinalizedHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*TSSVoter)(nil), "zetachain.zetacore.zetacore.TSSVoter")
}

func init() { proto.RegisterFile("zetacore/tss_voter.proto", fileDescriptor_0e88278440279c17) }

var fileDescriptor_0e88278440279c17 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x63, 0xda, 0xa6, 0xe0, 0x05, 0xc9, 0xaa, 0x90, 0x05, 0x92, 0x15, 0x31, 0x65, 0x21,
	0x1e, 0xf8, 0x03, 0x26, 0x58, 0x5b, 0xc4, 0xc0, 0x82, 0x9c, 0xe4, 0xe1, 0x58, 0x40, 0x1c, 0xd9,
	0x2e, 0x6a, 0xfb, 0x15, 0x7c, 0x15, 0x62, 0xec, 0xc8, 0x88, 0x92, 0x1f, 0x41, 0x76, 0xd2, 0x20,
	0xb1, 0xbd, 0x73, 0xdf, 0xbb, 0xd7, 0xd6, 0xc5, 0x74, 0x07, 0x4e, 0x14, 0xda, 0x00, 0x77, 0xd6,
	0x3e, 0xbd, 0x6b, 0x07, 0x26, 0x6b, 0x8c, 0x76, 0x9a, 0x5c, 0x84, 0x4d, 0x25, 0x54, 0x9d, 0x1d,
	0x6e, 0xc6, 0xe1, 0x7c, 0x21, 0xb5, 0xd4, 0xe1, 0x8e, 0xfb, 0xa9, 0xb7, 0x5c, 0x7e, 0x22, 0x7c,
	0x7c, 0xbf, 0x5a, 0x3d, 0xf8, 0x14, 0x42, 0xf1, 0xbc, 0x30, 0x20, 0x9c, 0x36, 0x14, 0x25, 0x28,
	0x3d, 0x59, 0x1e, 0x90, 0x2c, 0xf0, 0x4c, 0xd5, 0x25, 0x6c, 0xe8, 0x51, 0xd0, 0x7b, 0xf0, 0x6a,
	0x78, 0x8d, 0x4e, 0x7a, 0x35, 0x80, 0x4f, 0x11, 0x65, 0x69, 0xc0, 0x5a, 0x3a, 0xed, 0x53, 0x06,
	0x24, 0x67, 0x38, 0x6e, 0xd6, 0xf9, 0x0b, 0x6c, 0xe9, 0x2c, 0x2c, 0x06, 0xf2, 0x0e, 0xab, 0x64,
	0x0d, 0xc6, 0xd2, 0x38, 0x99, 0x78, 0xc7, 0x80, 0x24, 0xc5, 0xa7, 0xcf, 0xaa, 0x16, 0xaf, 0x6a,
	0x07, 0xe5, 0x2d, 0x28, 0x59, 0x39, 0x3a, 0x4f, 0x50, 0x3a, 0x5d, 0xfe, 0x97, 0x6f, 0xee, 0xbe,
	0x5a, 0x86, 0xf6, 0x2d, 0x43, 0x3f, 0x2d, 0x43, 0x1f, 0x1d, 0x8b, 0xf6, 0x1d, 0x8b, 0xbe, 0x3b,
	0x16, 0x3d, 0x72, 0xa9, 0x5c, 0xb5, 0xce, 0xb3, 0x42, 0xbf, 0x71, 0xdf, 0xc6, 0x55, 0xf8, 0x26,
	0x1f, 0x5b, 0xdc, 0xfc, 0x8d, 0x6e, 0xdb, 0x80, 0xcd, 0xe3, 0x50, 0xcd, 0xf5, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0x8f, 0x82, 0xb1, 0x69, 0x01, 0x00, 0x00,
}

func (m *TSSVoter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TSSVoter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TSSVoter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalizedHeight != 0 {
		i = encodeVarintTssVoter(dAtA, i, uint64(m.FinalizedHeight))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTssVoter(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintTssVoter(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTssVoter(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTssVoter(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTssVoter(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTssVoter(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTssVoter(dAtA []byte, offset int, v uint64) int {
	offset -= sovTssVoter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TSSVoter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTssVoter(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTssVoter(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTssVoter(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTssVoter(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovTssVoter(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTssVoter(uint64(l))
		}
	}
	if m.FinalizedHeight != 0 {
		n += 1 + sovTssVoter(uint64(m.FinalizedHeight))
	}
	return n
}

func sovTssVoter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTssVoter(x uint64) (n int) {
	return sovTssVoter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TSSVoter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTssVoter
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
			return fmt.Errorf("proto: TSSVoter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TSSVoter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTssVoter
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
				return ErrInvalidLengthTssVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTssVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTssVoter
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
				return ErrInvalidLengthTssVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTssVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTssVoter
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
				return ErrInvalidLengthTssVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTssVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTssVoter
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
				return ErrInvalidLengthTssVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTssVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTssVoter
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
				return ErrInvalidLengthTssVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTssVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTssVoter
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
				return ErrInvalidLengthTssVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTssVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedHeight", wireType)
			}
			m.FinalizedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTssVoter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalizedHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTssVoter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTssVoter
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
func skipTssVoter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTssVoter
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
					return 0, ErrIntOverflowTssVoter
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
					return 0, ErrIntOverflowTssVoter
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
				return 0, ErrInvalidLengthTssVoter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTssVoter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTssVoter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTssVoter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTssVoter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTssVoter = fmt.Errorf("proto: unexpected end of group")
)