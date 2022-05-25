// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: airequest/types/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("airequest/types/tx.proto", fileDescriptor_97af7e8a5031f004) }

var fileDescriptor_97af7e8a5031f004 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcc, 0x2c, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0xa9, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xc8, 0x2f, 0x4a, 0xcc, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3,
	0x03, 0xb1, 0xf4, 0xe0, 0x0a, 0xa5, 0x94, 0x31, 0xf5, 0xc4, 0x27, 0x66, 0xc6, 0x43, 0xc5, 0x20,
	0xda, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3,
	0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xb2, 0x46, 0xfd, 0x8c, 0x5c, 0xcc,
	0xbe, 0xc5, 0xe9, 0x42, 0xed, 0x8c, 0x5c, 0xfc, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x8e, 0x9e,
	0x41, 0x10, 0xfd, 0x42, 0x9a, 0x7a, 0xb8, 0x6c, 0xd6, 0xf3, 0x2d, 0x4e, 0x0f, 0x4e, 0x2d, 0x81,
	0x2b, 0x95, 0xd2, 0x21, 0x5a, 0x69, 0x50, 0x6a, 0xb1, 0x92, 0x74, 0xd3, 0xe5, 0x27, 0x93, 0x99,
	0x44, 0x95, 0x04, 0xf4, 0x11, 0x1e, 0x00, 0xb3, 0xac, 0x18, 0xb5, 0x9c, 0x5c, 0x4f, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3b, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0x1f, 0x6e, 0x1d, 0x98, 0xa5, 0x5f, 0xa1, 0x8f, 0x16, 0x14, 0x49, 0x6c, 0x60,
	0xff, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xab, 0x90, 0xc9, 0x58, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Create a new AI request
	CreateAIRequest(ctx context.Context, in *MsgSetAIRequest, opts ...grpc.CallOption) (*MsgSetAIRequestRes, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateAIRequest(ctx context.Context, in *MsgSetAIRequest, opts ...grpc.CallOption) (*MsgSetAIRequestRes, error) {
	out := new(MsgSetAIRequestRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.airequest.Msg/CreateAIRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Create a new AI request
	CreateAIRequest(context.Context, *MsgSetAIRequest) (*MsgSetAIRequestRes, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateAIRequest(ctx context.Context, req *MsgSetAIRequest) (*MsgSetAIRequestRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAIRequest not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateAIRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetAIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAIRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.airequest.Msg/CreateAIRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAIRequest(ctx, req.(*MsgSetAIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oraichain.orai.airequest.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAIRequest",
			Handler:    _Msg_CreateAIRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "airequest/types/tx.proto",
}
