// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: airesult/types/types_result.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/oraichain/orai/x/websocket/types"
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

type AIRequestResult struct {
	RequestID string            `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Results   []types.ValResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results"`
	Status    string            `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty" json:"request_status"`
}

func (m *AIRequestResult) Reset()         { *m = AIRequestResult{} }
func (m *AIRequestResult) String() string { return proto.CompactTextString(m) }
func (*AIRequestResult) ProtoMessage()    {}
func (*AIRequestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ecdf8589e162aaf, []int{0}
}
func (m *AIRequestResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AIRequestResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AIRequestResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AIRequestResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AIRequestResult.Merge(m, src)
}
func (m *AIRequestResult) XXX_Size() int {
	return m.Size()
}
func (m *AIRequestResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AIRequestResult.DiscardUnknown(m)
}

var xxx_messageInfo_AIRequestResult proto.InternalMessageInfo

func (m *AIRequestResult) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *AIRequestResult) GetResults() []types.ValResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *AIRequestResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*AIRequestResult)(nil), "oraichain.orai.airesult.AIRequestResult")
}

func init() { proto.RegisterFile("airesult/types/types_result.proto", fileDescriptor_9ecdf8589e162aaf) }

var fileDescriptor_9ecdf8589e162aaf = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xcc, 0x2c, 0x4a,
	0x2d, 0x2e, 0xcd, 0x29, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x90, 0xf1, 0x10, 0x21, 0xbd,
	0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xf1, 0xfc, 0xa2, 0xc4, 0xcc, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c,
	0x3d, 0x10, 0x4b, 0x0f, 0xa6, 0x43, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac, 0x46, 0x1f, 0xc4,
	0x82, 0x28, 0x97, 0x52, 0x2d, 0x4f, 0x4d, 0x2a, 0xce, 0x4f, 0xce, 0x4e, 0x45, 0x35, 0xb2, 0x2c,
	0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0xbf, 0x08, 0xa2, 0x4c, 0x69, 0x2f, 0x23, 0x17, 0xbf, 0xa3,
	0x67, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x49, 0x10, 0xd8, 0x40, 0x21, 0x1d, 0x2e, 0xae, 0x22,
	0x88, 0x40, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x13, 0xef, 0xa3, 0x7b, 0xf2,
	0x9c, 0x50, 0x65, 0x9e, 0x2e, 0x41, 0x9c, 0x50, 0x05, 0x9e, 0x29, 0x42, 0xce, 0x5c, 0xec, 0x10,
	0x87, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x1b, 0x29, 0xeb, 0xa1, 0xb9, 0x14, 0xee, 0x12,
	0xbd, 0xb0, 0xc4, 0x1c, 0x88, 0x1d, 0x4e, 0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0xc1, 0x74, 0x0a,
	0x19, 0x72, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x4b, 0x30, 0x83, 0xad, 0x93, 0xfc, 0x74,
	0x4f, 0x5e, 0x34, 0xab, 0x38, 0x3f, 0xcf, 0x4a, 0x09, 0xe6, 0x14, 0x88, 0xbc, 0x52, 0x10, 0x54,
	0xa1, 0x15, 0xcb, 0x8c, 0x05, 0xf2, 0x8c, 0x4e, 0x2e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0xa5, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f,
	0x77, 0x10, 0x98, 0xa5, 0x5f, 0xa1, 0x8f, 0x1a, 0xdc, 0x49, 0x6c, 0xe0, 0xc0, 0x30, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0x1a, 0xc0, 0x55, 0x87, 0x01, 0x00, 0x00,
}

func (m *AIRequestResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AIRequestResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AIRequestResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintTypesResult(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Results[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesResult(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RequestID) > 0 {
		i -= len(m.RequestID)
		copy(dAtA[i:], m.RequestID)
		i = encodeVarintTypesResult(dAtA, i, uint64(len(m.RequestID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypesResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypesResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AIRequestResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestID)
	if l > 0 {
		n += 1 + l + sovTypesResult(uint64(l))
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovTypesResult(uint64(l))
		}
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovTypesResult(uint64(l))
	}
	return n
}

func sovTypesResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypesResult(x uint64) (n int) {
	return sovTypesResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AIRequestResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesResult
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
			return fmt.Errorf("proto: AIRequestResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AIRequestResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesResult
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
				return ErrInvalidLengthTypesResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesResult
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
				return ErrInvalidLengthTypesResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, types.ValResult{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesResult
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
				return ErrInvalidLengthTypesResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypesResult
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypesResult
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
func skipTypesResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypesResult
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
					return 0, ErrIntOverflowTypesResult
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
					return 0, ErrIntOverflowTypesResult
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
				return 0, ErrInvalidLengthTypesResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypesResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypesResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypesResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypesResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypesResult = fmt.Errorf("proto: unexpected end of group")
)
