// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: airequest/types/types_ai_request.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/oraichain/orai/x/provider/types"
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

type AIRequest struct {
	RequestID        string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	OracleScriptName string `protobuf:"bytes,2,opt,name=oracle_script_name,json=oracleScriptName,proto3" json:"oracle_script_name,omitempty" json:"oscript_name"`
	// Owner is the address who is allowed to make further changes to the data source.
	Creator        github_com_cosmos_cosmos_sdk_types.AccAddress   `protobuf:"bytes,3,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty" json:"request_creator"`
	Validators     []github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,4,rep,name=validators,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"validators,omitempty" json:"validator_addr"`
	BlockHeight    int64                                           `protobuf:"varint,5,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	AiDataSources  []types.AIDataSource                            `protobuf:"bytes,6,rep,name=ai_data_sources,json=aiDataSources,proto3" json:"ai_data_sources" json:"data_sources"`
	TestCases      []types.TestCase                                `protobuf:"bytes,7,rep,name=test_cases,json=testCases,proto3" json:"test_cases" json:"test_cases"`
	Fees           github_com_cosmos_cosmos_sdk_types.Coins        `protobuf:"bytes,8,rep,name=fees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fees" json:"transaction_fee"`
	Input          []byte                                          `protobuf:"bytes,9,opt,name=input,proto3" json:"input,omitempty" json:"request_input"`
	ExpectedOutput []byte                                          `protobuf:"bytes,10,opt,name=expected_output,json=expectedOutput,proto3" json:"expected_output,omitempty" json:"expected_output"`
}

func (m *AIRequest) Reset()         { *m = AIRequest{} }
func (m *AIRequest) String() string { return proto.CompactTextString(m) }
func (*AIRequest) ProtoMessage()    {}
func (*AIRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffd315b5790d6393, []int{0}
}
func (m *AIRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AIRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AIRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AIRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AIRequest.Merge(m, src)
}
func (m *AIRequest) XXX_Size() int {
	return m.Size()
}
func (m *AIRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AIRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AIRequest proto.InternalMessageInfo

func (m *AIRequest) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *AIRequest) GetOracleScriptName() string {
	if m != nil {
		return m.OracleScriptName
	}
	return ""
}

func (m *AIRequest) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *AIRequest) GetValidators() []github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *AIRequest) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *AIRequest) GetAiDataSources() []types.AIDataSource {
	if m != nil {
		return m.AiDataSources
	}
	return nil
}

func (m *AIRequest) GetTestCases() []types.TestCase {
	if m != nil {
		return m.TestCases
	}
	return nil
}

func (m *AIRequest) GetFees() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fees
	}
	return nil
}

func (m *AIRequest) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *AIRequest) GetExpectedOutput() []byte {
	if m != nil {
		return m.ExpectedOutput
	}
	return nil
}

// Params defines the set of wasm parameters.
type Params struct {
	MaximumRequestBytes uint64 `protobuf:"varint,1,opt,name=maximum_request_bytes,json=maximumRequestBytes,proto3" json:"maximum_request_bytes,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffd315b5790d6393, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaximumRequestBytes() uint64 {
	if m != nil {
		return m.MaximumRequestBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*AIRequest)(nil), "oraichain.orai.airequest.AIRequest")
	proto.RegisterType((*Params)(nil), "oraichain.orai.airequest.Params")
}

func init() {
	proto.RegisterFile("airequest/types/types_ai_request.proto", fileDescriptor_ffd315b5790d6393)
}

var fileDescriptor_ffd315b5790d6393 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x53, 0xd3, 0x40,
	0x14, 0x6f, 0x6c, 0x29, 0x76, 0x01, 0xd1, 0x00, 0x1a, 0x70, 0x4c, 0x4a, 0x66, 0x74, 0x3a, 0xa3,
	0x24, 0x03, 0xde, 0xb8, 0x35, 0xc0, 0x8c, 0xc5, 0x19, 0x75, 0x82, 0xe3, 0x41, 0x0f, 0x99, 0xd7,
	0xcd, 0xd2, 0x2e, 0x34, 0xd9, 0xba, 0xbb, 0x65, 0xe0, 0x5b, 0x78, 0xf4, 0xe8, 0xd9, 0xbb, 0xdf,
	0x81, 0x23, 0x47, 0x4f, 0xd1, 0x29, 0xdf, 0xa0, 0x47, 0x4f, 0x4e, 0x36, 0x5b, 0x5a, 0x98, 0xd1,
	0xe1, 0xd2, 0x6e, 0xde, 0xfb, 0xfd, 0x79, 0x79, 0xfb, 0x5e, 0xd0, 0x33, 0xa0, 0x9c, 0x7c, 0x1e,
	0x10, 0x21, 0x7d, 0x79, 0xd6, 0x27, 0xa2, 0xf8, 0x8d, 0x80, 0x46, 0x3a, 0xec, 0xf5, 0x39, 0x93,
	0xcc, 0xb4, 0x18, 0x07, 0x8a, 0xbb, 0x40, 0x53, 0x2f, 0x3f, 0x79, 0x57, 0xb4, 0x35, 0x1b, 0x33,
	0x91, 0x30, 0xe1, 0xb7, 0x41, 0x10, 0xff, 0x64, 0xb3, 0x4d, 0x24, 0x6c, 0xfa, 0x98, 0xd1, 0xb4,
	0x60, 0xae, 0x2d, 0x77, 0x58, 0x87, 0xa9, 0xa3, 0x9f, 0x9f, 0x74, 0xf4, 0x49, 0x9f, 0xb3, 0x13,
	0x1a, 0x13, 0x7e, 0xcd, 0x36, 0x16, 0xff, 0x4d, 0x4b, 0x5c, 0xa4, 0xdd, 0x1f, 0x55, 0x54, 0x6b,
	0xb6, 0xc2, 0xa2, 0x02, 0xf3, 0x05, 0x42, 0xba, 0x98, 0x88, 0xc6, 0x96, 0x51, 0x37, 0x1a, 0xb5,
	0x60, 0x61, 0x98, 0x39, 0x35, 0x0d, 0x68, 0xed, 0x86, 0x35, 0x0d, 0x68, 0xc5, 0xe6, 0x1e, 0x32,
	0x19, 0x07, 0xdc, 0x23, 0x91, 0xc0, 0x9c, 0xf6, 0x65, 0x94, 0x42, 0x42, 0xac, 0x3b, 0x8a, 0xf5,
	0x68, 0x94, 0x39, 0x4b, 0x47, 0x82, 0xa5, 0xdb, 0x2e, 0x9b, 0xca, 0xba, 0xe1, 0xfd, 0x82, 0x72,
	0xa0, 0x62, 0x6f, 0x20, 0x21, 0x26, 0x41, 0xb3, 0x98, 0x13, 0x90, 0x8c, 0x5b, 0xe5, 0xba, 0xd1,
	0x98, 0x0f, 0x5e, 0x8f, 0x32, 0xe7, 0x61, 0xc1, 0x1d, 0x57, 0xa3, 0x01, 0xee, 0x9f, 0xcc, 0xd9,
	0xe8, 0x50, 0xd9, 0x1d, 0xb4, 0x3d, 0xcc, 0x12, 0x5f, 0x37, 0xac, 0xf8, 0xdb, 0x10, 0xf1, 0x71,
	0xf1, 0x7a, 0x5e, 0x13, 0xe3, 0x66, 0x1c, 0x73, 0x22, 0x44, 0x38, 0xd6, 0x36, 0x8f, 0x10, 0x3a,
	0x81, 0x1e, 0x8d, 0xf3, 0x07, 0x61, 0x55, 0xea, 0xe5, 0xc6, 0x7c, 0xb0, 0x3f, 0xca, 0x9c, 0x95,
	0xc2, 0xe9, 0x2a, 0x17, 0x41, 0x1c, 0xdf, 0xda, 0xe8, 0x03, 0xf4, 0xc6, 0x46, 0x53, 0xea, 0xe6,
	0x3a, 0x9a, 0x6f, 0xf7, 0x18, 0x3e, 0x8e, 0xba, 0x84, 0x76, 0xba, 0xd2, 0x9a, 0xa9, 0x1b, 0x8d,
	0x72, 0x38, 0xa7, 0x62, 0xaf, 0x54, 0xc8, 0x3c, 0x42, 0x8b, 0x40, 0xa3, 0x18, 0x24, 0x44, 0x82,
	0x0d, 0x38, 0x26, 0xc2, 0xaa, 0xd6, 0xcb, 0x8d, 0xb9, 0xad, 0xa7, 0xde, 0x8d, 0x01, 0x19, 0x5f,
	0xa0, 0xd7, 0x6c, 0xed, 0x82, 0x84, 0x03, 0x85, 0x0e, 0x1e, 0x9f, 0x67, 0x4e, 0x69, 0xd2, 0xe4,
	0x69, 0x21, 0x37, 0x5c, 0x00, 0x3a, 0x81, 0x0a, 0xf3, 0x13, 0x42, 0x52, 0x75, 0x11, 0x04, 0x11,
	0xd6, 0xac, 0xb2, 0x59, 0xff, 0xa7, 0xcd, 0x7b, 0x22, 0xe4, 0x0e, 0x08, 0x12, 0xac, 0x6a, 0x8b,
	0x07, 0x85, 0xc5, 0x44, 0xc2, 0x0d, 0x6b, 0x52, 0x83, 0x84, 0x79, 0x86, 0x2a, 0x87, 0x84, 0x08,
	0xeb, 0xae, 0x92, 0x5d, 0xf5, 0x8a, 0x1e, 0x79, 0xf9, 0x10, 0x7b, 0x7a, 0x88, 0xbd, 0x1d, 0x46,
	0xd3, 0x60, 0x5f, 0xcb, 0xe9, 0xab, 0x95, 0x1c, 0x52, 0x01, 0x58, 0x52, 0x96, 0x46, 0x87, 0x84,
	0xb8, 0xdf, 0x7f, 0x39, 0x8d, 0x5b, 0x74, 0x3c, 0x97, 0x12, 0xa1, 0xb2, 0x34, 0x3d, 0x34, 0x43,
	0xd3, 0xfe, 0x40, 0x5a, 0x35, 0x35, 0x37, 0xd6, 0x28, 0x73, 0x96, 0xaf, 0xcf, 0x8d, 0x4a, 0xbb,
	0x61, 0x01, 0x33, 0x77, 0xd0, 0x22, 0x39, 0xed, 0x13, 0x2c, 0x49, 0x1c, 0xb1, 0x81, 0xcc, 0x99,
	0x48, 0x31, 0xd7, 0x26, 0x65, 0xdd, 0x00, 0xb8, 0xe1, 0xbd, 0x71, 0xe4, 0xad, 0x0a, 0x6c, 0x57,
	0xbe, 0x7e, 0x73, 0x0c, 0x37, 0x40, 0xd5, 0x77, 0xc0, 0x21, 0x11, 0xe6, 0x16, 0x5a, 0x49, 0xe0,
	0x94, 0x26, 0x83, 0x64, 0xbc, 0xe8, 0x51, 0xfb, 0x4c, 0x12, 0xa1, 0xd6, 0xa7, 0x12, 0x2e, 0xe9,
	0xa4, 0xde, 0xa0, 0x20, 0x4f, 0x29, 0x8d, 0x52, 0xb0, 0x77, 0x3e, 0xb4, 0x8d, 0x8b, 0xa1, 0x6d,
	0xfc, 0x1e, 0xda, 0xc6, 0x97, 0x4b, 0xbb, 0x74, 0x71, 0x69, 0x97, 0x7e, 0x5e, 0xda, 0xa5, 0x8f,
	0xcf, 0xa7, 0x1a, 0x71, 0x75, 0x4d, 0xea, 0xe4, 0x9f, 0xfa, 0x37, 0xbe, 0x33, 0xed, 0xaa, 0xda,
	0xe4, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xf9, 0x5b, 0x6a, 0x81, 0x04, 0x00, 0x00,
}

func (m *AIRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AIRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AIRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExpectedOutput) > 0 {
		i -= len(m.ExpectedOutput)
		copy(dAtA[i:], m.ExpectedOutput)
		i = encodeVarintTypesAiRequest(dAtA, i, uint64(len(m.ExpectedOutput)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Input) > 0 {
		i -= len(m.Input)
		copy(dAtA[i:], m.Input)
		i = encodeVarintTypesAiRequest(dAtA, i, uint64(len(m.Input)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesAiRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.TestCases) > 0 {
		for iNdEx := len(m.TestCases) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TestCases[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesAiRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AiDataSources) > 0 {
		for iNdEx := len(m.AiDataSources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AiDataSources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesAiRequest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTypesAiRequest(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Validators[iNdEx])
			copy(dAtA[i:], m.Validators[iNdEx])
			i = encodeVarintTypesAiRequest(dAtA, i, uint64(len(m.Validators[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTypesAiRequest(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OracleScriptName) > 0 {
		i -= len(m.OracleScriptName)
		copy(dAtA[i:], m.OracleScriptName)
		i = encodeVarintTypesAiRequest(dAtA, i, uint64(len(m.OracleScriptName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RequestID) > 0 {
		i -= len(m.RequestID)
		copy(dAtA[i:], m.RequestID)
		i = encodeVarintTypesAiRequest(dAtA, i, uint64(len(m.RequestID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaximumRequestBytes != 0 {
		i = encodeVarintTypesAiRequest(dAtA, i, uint64(m.MaximumRequestBytes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypesAiRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypesAiRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AIRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestID)
	if l > 0 {
		n += 1 + l + sovTypesAiRequest(uint64(l))
	}
	l = len(m.OracleScriptName)
	if l > 0 {
		n += 1 + l + sovTypesAiRequest(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTypesAiRequest(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, b := range m.Validators {
			l = len(b)
			n += 1 + l + sovTypesAiRequest(uint64(l))
		}
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTypesAiRequest(uint64(m.BlockHeight))
	}
	if len(m.AiDataSources) > 0 {
		for _, e := range m.AiDataSources {
			l = e.Size()
			n += 1 + l + sovTypesAiRequest(uint64(l))
		}
	}
	if len(m.TestCases) > 0 {
		for _, e := range m.TestCases {
			l = e.Size()
			n += 1 + l + sovTypesAiRequest(uint64(l))
		}
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovTypesAiRequest(uint64(l))
		}
	}
	l = len(m.Input)
	if l > 0 {
		n += 1 + l + sovTypesAiRequest(uint64(l))
	}
	l = len(m.ExpectedOutput)
	if l > 0 {
		n += 1 + l + sovTypesAiRequest(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaximumRequestBytes != 0 {
		n += 1 + sovTypesAiRequest(uint64(m.MaximumRequestBytes))
	}
	return n
}

func sovTypesAiRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypesAiRequest(x uint64) (n int) {
	return sovTypesAiRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AIRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesAiRequest
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
			return fmt.Errorf("proto: AIRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AIRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
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
				return ErrInvalidLengthTypesAiRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleScriptName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
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
				return ErrInvalidLengthTypesAiRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleScriptName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
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
				return ErrInvalidLengthTypesAiRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
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
				return ErrInvalidLengthTypesAiRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, make([]byte, postIndex-iNdEx))
			copy(m.Validators[len(m.Validators)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AiDataSources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
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
				return ErrInvalidLengthTypesAiRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AiDataSources = append(m.AiDataSources, types.AIDataSource{})
			if err := m.AiDataSources[len(m.AiDataSources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestCases", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
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
				return ErrInvalidLengthTypesAiRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TestCases = append(m.TestCases, types.TestCase{})
			if err := m.TestCases[len(m.TestCases)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
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
				return ErrInvalidLengthTypesAiRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, types1.Coin{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
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
				return ErrInvalidLengthTypesAiRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Input = append(m.Input[:0], dAtA[iNdEx:postIndex]...)
			if m.Input == nil {
				m.Input = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedOutput", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
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
				return ErrInvalidLengthTypesAiRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpectedOutput = append(m.ExpectedOutput[:0], dAtA[iNdEx:postIndex]...)
			if m.ExpectedOutput == nil {
				m.ExpectedOutput = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesAiRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypesAiRequest
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesAiRequest
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaximumRequestBytes", wireType)
			}
			m.MaximumRequestBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesAiRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaximumRequestBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypesAiRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypesAiRequest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypesAiRequest
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
func skipTypesAiRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypesAiRequest
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
					return 0, ErrIntOverflowTypesAiRequest
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
					return 0, ErrIntOverflowTypesAiRequest
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
				return 0, ErrInvalidLengthTypesAiRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypesAiRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypesAiRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypesAiRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypesAiRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypesAiRequest = fmt.Errorf("proto: unexpected end of group")
)
