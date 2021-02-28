// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heimdall/base/v1beta1/divident.proto

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

type DividendAccount struct {
	User      string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	FeeAmount string `protobuf:"bytes,2,opt,name=fee_amount,json=feeAmount,proto3" json:"fee_amount,omitempty" yaml:"fee_amount"`
}

func (m *DividendAccount) Reset()      { *m = DividendAccount{} }
func (*DividendAccount) ProtoMessage() {}
func (*DividendAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_231b3e3510b72567, []int{0}
}
func (m *DividendAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DividendAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DividendAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DividendAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DividendAccount.Merge(m, src)
}
func (m *DividendAccount) XXX_Size() int {
	return m.Size()
}
func (m *DividendAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_DividendAccount.DiscardUnknown(m)
}

var xxx_messageInfo_DividendAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DividendAccount)(nil), "heimdall.types.DividendAccount")
}

func init() {
	proto.RegisterFile("heimdall/base/v1beta1/divident.proto", fileDescriptor_231b3e3510b72567)
}

var fileDescriptor_231b3e3510b72567 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x48, 0xcd, 0xcc,
	0x4d, 0x49, 0xcc, 0xc9, 0xd1, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0x4f, 0xc9, 0x2c, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x83, 0xa9, 0xd2, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x4b, 0xe9, 0x83, 0x58, 0x10, 0x55, 0x4a, 0xa9, 0x5c, 0xfc, 0x2e, 0x10, 0x7d, 0x29, 0x8e,
	0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0xa5, 0xc5, 0xa9, 0x45, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x09, 0x17, 0x57, 0x5a, 0x6a, 0x6a, 0x7c, 0x62,
	0x2e, 0x48, 0x85, 0x04, 0x13, 0x48, 0xc6, 0x49, 0xf4, 0xd3, 0x3d, 0x79, 0xc1, 0xca, 0xc4, 0xdc,
	0x1c, 0x2b, 0x25, 0x84, 0x9c, 0x52, 0x10, 0x67, 0x5a, 0x6a, 0xaa, 0x23, 0x98, 0x6d, 0xc5, 0xd1,
	0xb1, 0x40, 0x9e, 0x61, 0xc6, 0x02, 0x79, 0x06, 0x27, 0x87, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e,
	0x3c, 0x96, 0x63, 0x88, 0x52, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0xcf, 0x4d, 0x2c, 0xc9, 0x4c, 0xce, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x87, 0x7b, 0x12,
	0xec, 0xfc, 0x24, 0x36, 0xb0, 0x7b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xcb, 0xca,
	0x76, 0xfd, 0x00, 0x00, 0x00,
}

func (m *DividendAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DividendAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DividendAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeAmount) > 0 {
		i -= len(m.FeeAmount)
		copy(dAtA[i:], m.FeeAmount)
		i = encodeVarintDivident(dAtA, i, uint64(len(m.FeeAmount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintDivident(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDivident(dAtA []byte, offset int, v uint64) int {
	offset -= sovDivident(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DividendAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovDivident(uint64(l))
	}
	l = len(m.FeeAmount)
	if l > 0 {
		n += 1 + l + sovDivident(uint64(l))
	}
	return n
}

func sovDivident(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDivident(x uint64) (n int) {
	return sovDivident(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DividendAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDivident
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
			return fmt.Errorf("proto: DividendAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DividendAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDivident
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
				return ErrInvalidLengthDivident
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDivident
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDivident
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
				return ErrInvalidLengthDivident
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDivident
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDivident(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDivident
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDivident
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
func skipDivident(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDivident
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
					return 0, ErrIntOverflowDivident
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
					return 0, ErrIntOverflowDivident
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
				return 0, ErrInvalidLengthDivident
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDivident
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDivident
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDivident        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDivident          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDivident = fmt.Errorf("proto: unexpected end of group")
)