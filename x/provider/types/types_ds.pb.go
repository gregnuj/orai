// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provider/types/types_ds.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type AIDataSource struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	// Owner is the address who is allowed to make further changes to the data source.
	Owner       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty"`
	Description string                                        `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Fees        github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,5,rep,name=fees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fees" json:"transaction_fee"`
}

func (m *AIDataSource) Reset()         { *m = AIDataSource{} }
func (m *AIDataSource) String() string { return proto.CompactTextString(m) }
func (*AIDataSource) ProtoMessage()    {}
func (*AIDataSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_31f42dc7d5268e98, []int{0}
}
func (m *AIDataSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AIDataSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AIDataSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AIDataSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AIDataSource.Merge(m, src)
}
func (m *AIDataSource) XXX_Size() int {
	return m.Size()
}
func (m *AIDataSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AIDataSource.DiscardUnknown(m)
}

var xxx_messageInfo_AIDataSource proto.InternalMessageInfo

func (m *AIDataSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AIDataSource) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *AIDataSource) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *AIDataSource) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AIDataSource) GetFees() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fees
	}
	return nil
}

func init() {
	proto.RegisterType((*AIDataSource)(nil), "oraichain.orai.provider.AIDataSource")
}

func init() { proto.RegisterFile("provider/types/types_ds.proto", fileDescriptor_31f42dc7d5268e98) }

var fileDescriptor_31f42dc7d5268e98 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x6e, 0xea, 0x30,
	0x14, 0x86, 0x13, 0x08, 0x57, 0xf7, 0x06, 0xa6, 0xe8, 0xea, 0xde, 0x14, 0xa9, 0x49, 0xc4, 0x14,
	0x55, 0xc2, 0x16, 0xed, 0xc6, 0x06, 0x45, 0xaa, 0xda, 0x91, 0x6e, 0x5d, 0x90, 0xe3, 0x18, 0x70,
	0x2b, 0x7c, 0x22, 0xdb, 0xd0, 0xf2, 0x16, 0x1d, 0x3b, 0x76, 0xe6, 0x49, 0x18, 0x19, 0x3b, 0xd1,
	0x0a, 0xde, 0xa0, 0x63, 0xa7, 0x2a, 0x31, 0x45, 0x74, 0xeb, 0x92, 0x7c, 0xfe, 0xe5, 0xf3, 0x9f,
	0xf3, 0xfb, 0xb8, 0xc7, 0x99, 0x84, 0x19, 0x4f, 0x99, 0xc4, 0x7a, 0x9e, 0x31, 0x65, 0xbe, 0x83,
	0x54, 0xa1, 0x4c, 0x82, 0x06, 0xef, 0x3f, 0x48, 0xc2, 0xe9, 0x98, 0x70, 0x81, 0x72, 0x42, 0x5f,
	0xb7, 0xeb, 0x01, 0x05, 0x35, 0x01, 0x85, 0x13, 0xa2, 0x18, 0x9e, 0xb5, 0x12, 0xa6, 0x49, 0x0b,
	0x53, 0xe0, 0xc2, 0x14, 0xd6, 0xff, 0x8e, 0x60, 0x04, 0x05, 0xe2, 0x9c, 0x8c, 0xda, 0x58, 0x94,
	0xdc, 0x5a, 0xe7, 0xb2, 0x47, 0x34, 0xb9, 0x86, 0xa9, 0xa4, 0xcc, 0xf3, 0x5c, 0x47, 0x90, 0x09,
	0xf3, 0xed, 0xc8, 0x8e, 0xff, 0xf4, 0x0b, 0xf6, 0xea, 0xee, 0x6f, 0x0a, 0x42, 0x4b, 0x42, 0xb5,
	0x5f, 0x2a, 0xf4, 0xfd, 0xd9, 0xbb, 0x70, 0x2b, 0x70, 0x2f, 0x98, 0xf4, 0xcb, 0x91, 0x1d, 0xd7,
	0xba, 0xad, 0x8f, 0x75, 0xd8, 0x1c, 0x71, 0x3d, 0x9e, 0x26, 0x88, 0xc2, 0x04, 0xef, 0x86, 0x32,
	0xbf, 0xa6, 0x4a, 0xef, 0x4c, 0x1e, 0xd4, 0xa1, 0xb4, 0x93, 0xa6, 0x92, 0x29, 0xd5, 0x37, 0xf5,
	0x5e, 0xe4, 0x56, 0x53, 0xa6, 0xa8, 0xe4, 0x99, 0xe6, 0x20, 0x7c, 0xa7, 0xe8, 0x73, 0x28, 0x79,
	0x73, 0xd7, 0x19, 0x32, 0xa6, 0xfc, 0x4a, 0x54, 0x8e, 0xab, 0xa7, 0x47, 0xc8, 0x98, 0xa2, 0x3c,
	0x30, 0xda, 0x05, 0x46, 0xe7, 0xc0, 0x45, 0xf7, 0x6a, 0xb9, 0x0e, 0xad, 0xf7, 0x75, 0xf8, 0xef,
	0x56, 0x81, 0x68, 0x37, 0xb4, 0x24, 0x42, 0x11, 0x9a, 0x7b, 0x0c, 0x86, 0x8c, 0x35, 0x16, 0xaf,
	0x61, 0xfc, 0x83, 0x11, 0x73, 0x2b, 0xd5, 0x2f, 0x5a, 0xb6, 0x9d, 0xa7, 0xe7, 0xd0, 0xee, 0xf6,
	0x96, 0x9b, 0xc0, 0x5e, 0x6d, 0x02, 0xfb, 0x6d, 0x13, 0xd8, 0x8f, 0xdb, 0xc0, 0x5a, 0x6d, 0x03,
	0xeb, 0x65, 0x1b, 0x58, 0x37, 0x27, 0x07, 0x7e, 0xfb, 0x05, 0x15, 0x84, 0x1f, 0xf0, 0xf7, 0x85,
	0x26, 0xbf, 0x8a, 0x97, 0x3f, 0xfb, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xd6, 0xed, 0xfe, 0xe9,
	0x01, 0x00, 0x00,
}

func (m *AIDataSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AIDataSource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AIDataSource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesDs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTypesDs(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypesDs(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintTypesDs(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypesDs(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypesDs(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypesDs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AIDataSource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypesDs(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovTypesDs(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypesDs(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTypesDs(uint64(l))
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovTypesDs(uint64(l))
		}
	}
	return n
}

func sovTypesDs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypesDs(x uint64) (n int) {
	return sovTypesDs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AIDataSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesDs
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
			return fmt.Errorf("proto: AIDataSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AIDataSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
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
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
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
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
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
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
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
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesDs
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
				return ErrInvalidLengthTypesDs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesDs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, types.Coin{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesDs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypesDs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypesDs
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
func skipTypesDs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypesDs
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
					return 0, ErrIntOverflowTypesDs
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
					return 0, ErrIntOverflowTypesDs
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
				return 0, ErrInvalidLengthTypesDs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypesDs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypesDs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypesDs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypesDs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypesDs = fmt.Errorf("proto: unexpected end of group")
)
