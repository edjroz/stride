// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/gov.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type AddValidatorProposal struct {
	Title            string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description      string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	HostZone         string `protobuf:"bytes,3,opt,name=host_zone,json=hostZone,proto3" json:"host_zone,omitempty"`
	ValidatorName    string `protobuf:"bytes,4,opt,name=validator_name,json=validatorName,proto3" json:"validator_name,omitempty"`
	ValidatorAddress string `protobuf:"bytes,5,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Deposit          string `protobuf:"bytes,6,opt,name=deposit,proto3" json:"deposit,omitempty" yaml:"deposit"`
}

func (m *AddValidatorProposal) Reset()      { *m = AddValidatorProposal{} }
func (*AddValidatorProposal) ProtoMessage() {}
func (*AddValidatorProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8204317b384c5680, []int{0}
}
func (m *AddValidatorProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddValidatorProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddValidatorProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddValidatorProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddValidatorProposal.Merge(m, src)
}
func (m *AddValidatorProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddValidatorProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddValidatorProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddValidatorProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddValidatorProposal)(nil), "stride.stakeibc.AddValidatorProposal")
}

func init() { proto.RegisterFile("stride/stakeibc/gov.proto", fileDescriptor_8204317b384c5680) }

var fileDescriptor_8204317b384c5680 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xb1, 0x4e, 0x2a, 0x41,
	0x14, 0x86, 0x77, 0xb9, 0x17, 0xee, 0x65, 0xee, 0x15, 0x75, 0xb2, 0xc5, 0x82, 0xc9, 0x2e, 0x21,
	0x31, 0xb1, 0x10, 0x36, 0x86, 0x8e, 0x0e, 0x12, 0x2b, 0x8d, 0x31, 0x90, 0x58, 0xd0, 0x90, 0xd9,
	0x9d, 0xc9, 0x32, 0x71, 0x77, 0xcf, 0x66, 0x67, 0x24, 0xe2, 0x13, 0x58, 0xda, 0x98, 0x58, 0xf2,
	0x10, 0x3e, 0x84, 0x25, 0xb1, 0xb2, 0x32, 0x06, 0x1a, 0x6b, 0x9f, 0xc0, 0x30, 0xb3, 0x80, 0xdd,
	0x9c, 0xef, 0xfb, 0x4f, 0x26, 0xe7, 0x1c, 0x54, 0x15, 0x32, 0xe3, 0x94, 0x79, 0x42, 0x92, 0x6b,
	0xc6, 0xfd, 0xc0, 0x0b, 0x61, 0xd2, 0x4a, 0x33, 0x90, 0x80, 0x77, 0xb5, 0x6a, 0xad, 0x55, 0xad,
	0x1a, 0x80, 0x88, 0x41, 0x8c, 0x94, 0xf6, 0x74, 0xa1, 0xb3, 0x35, 0x2b, 0x84, 0x10, 0x34, 0x5f,
	0xbd, 0x34, 0x6d, 0x3c, 0x16, 0x90, 0xd5, 0xa5, 0xf4, 0x8a, 0x44, 0x9c, 0x12, 0x09, 0xd9, 0x65,
	0x06, 0x29, 0x08, 0x12, 0x61, 0x0b, 0x15, 0x25, 0x97, 0x11, 0xb3, 0xcd, 0xba, 0x79, 0x54, 0xee,
	0xeb, 0x02, 0xd7, 0xd1, 0x3f, 0xca, 0x44, 0x90, 0xf1, 0x54, 0x72, 0x48, 0xec, 0x82, 0x72, 0x3f,
	0x11, 0x3e, 0x40, 0xe5, 0x31, 0x08, 0x39, 0xba, 0x83, 0x84, 0xd9, 0xbf, 0x94, 0xff, 0xbb, 0x02,
	0x43, 0x48, 0x18, 0x3e, 0x44, 0x95, 0xc9, 0xfa, 0xa7, 0x51, 0x42, 0x62, 0x66, 0xff, 0x56, 0x89,
	0x9d, 0x0d, 0xbd, 0x20, 0x31, 0xc3, 0xa7, 0x68, 0x7f, 0x1b, 0x23, 0x94, 0x66, 0x4c, 0x08, 0xbb,
	0xb8, 0x4a, 0xf6, 0xec, 0xd7, 0xe7, 0xa6, 0x95, 0xcf, 0xd5, 0xd5, 0x66, 0x20, 0x33, 0x9e, 0x84,
	0xfd, 0xbd, 0x4d, 0x4b, 0xce, 0xf1, 0x31, 0xfa, 0x43, 0x59, 0x0a, 0x82, 0x4b, 0xbb, 0xa4, 0x9a,
	0xf1, 0xd7, 0xbb, 0x5b, 0x99, 0x92, 0x38, 0xea, 0x34, 0x72, 0xd1, 0xe8, 0xaf, 0x23, 0x9d, 0xff,
	0xf7, 0x33, 0xd7, 0x78, 0x9a, 0xb9, 0xc6, 0xe7, 0xcc, 0x35, 0x7b, 0x67, 0x2f, 0x0b, 0xc7, 0x9c,
	0x2f, 0x1c, 0xf3, 0x63, 0xe1, 0x98, 0x0f, 0x4b, 0xc7, 0x98, 0x2f, 0x1d, 0xe3, 0x6d, 0xe9, 0x18,
	0xc3, 0x93, 0x90, 0xcb, 0xf1, 0x8d, 0xdf, 0x0a, 0x20, 0xf6, 0x06, 0x6a, 0xfd, 0xcd, 0x73, 0xe2,
	0x0b, 0x2f, 0xbf, 0xd2, 0xa4, 0xed, 0xdd, 0x6e, 0x4f, 0x25, 0xa7, 0x29, 0x13, 0x7e, 0x49, 0xed,
	0xba, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x8b, 0xe3, 0x25, 0xca, 0x01, 0x00, 0x00,
}

func (this *AddValidatorProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AddValidatorProposal)
	if !ok {
		that2, ok := that.(AddValidatorProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.HostZone != that1.HostZone {
		return false
	}
	if this.ValidatorName != that1.ValidatorName {
		return false
	}
	if this.ValidatorAddress != that1.ValidatorAddress {
		return false
	}
	if this.Deposit != that1.Deposit {
		return false
	}
	return true
}
func (m *AddValidatorProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddValidatorProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddValidatorProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGov(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ValidatorName) > 0 {
		i -= len(m.ValidatorName)
		copy(dAtA[i:], m.ValidatorName)
		i = encodeVarintGov(dAtA, i, uint64(len(m.ValidatorName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.HostZone) > 0 {
		i -= len(m.HostZone)
		copy(dAtA[i:], m.HostZone)
		i = encodeVarintGov(dAtA, i, uint64(len(m.HostZone)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddValidatorProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.HostZone)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.ValidatorName)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddValidatorProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: AddValidatorProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddValidatorProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostZone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)