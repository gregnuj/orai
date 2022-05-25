// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: airesult/types/query.proto

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

func init() { proto.RegisterFile("airesult/types/query.proto", fileDescriptor_9a84ee9582adbca2) }

var fileDescriptor_9a84ee9582adbca2 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xcc, 0x2c, 0x4a,
	0x2d, 0x2e, 0xcd, 0x29, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcf, 0x2f, 0x4a, 0xcc, 0x4c, 0xce, 0x48, 0xcc,
	0xcc, 0xd3, 0x03, 0xb1, 0xf4, 0x60, 0x4a, 0xa5, 0x94, 0xb1, 0x69, 0x8a, 0x4f, 0x2b, 0xcd, 0xc9,
	0x89, 0x2f, 0x4a, 0x2d, 0x84, 0xe8, 0x96, 0x52, 0xc4, 0xaa, 0xa8, 0x28, 0xb5, 0x3c, 0xb1, 0x28,
	0x05, 0xaa, 0x44, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31,
	0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x18, 0x22, 0x6b, 0xb4, 0x9a, 0x89, 0x8b,
	0x35, 0x10, 0xa4, 0x49, 0x68, 0x1a, 0x23, 0x97, 0x00, 0x98, 0xe5, 0x56, 0x9a, 0x93, 0x13, 0x94,
	0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xa3, 0x87, 0xc3, 0x79, 0x7a, 0xe8, 0x4a, 0x83, 0x52,
	0x0b, 0xa5, 0x48, 0x51, 0x5d, 0xac, 0xa4, 0xd6, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x05, 0x21, 0x39,
	0x7d, 0xb8, 0x27, 0x40, 0xbe, 0x2b, 0x4a, 0x2d, 0xd4, 0xaf, 0x2e, 0x82, 0x28, 0x8b, 0xcf, 0x4c,
	0xa9, 0x15, 0xea, 0x60, 0xe4, 0xe2, 0x06, 0xeb, 0x0f, 0x02, 0x7b, 0x4b, 0x48, 0x1d, 0xbf, 0x2d,
	0x10, 0x55, 0x20, 0xe7, 0x10, 0xa9, 0xb0, 0x58, 0x49, 0x1d, 0xec, 0x12, 0x45, 0x21, 0x79, 0x84,
	0x4b, 0x20, 0x41, 0xa8, 0x5f, 0x9d, 0x94, 0x93, 0x9f, 0x9c, 0x1d, 0x9f, 0x91, 0x9a, 0x99, 0x9e,
	0x51, 0x52, 0xeb, 0xe4, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x5a,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x70, 0x5b, 0xc1, 0x2c, 0xfd,
	0x0a, 0x7d, 0xd4, 0x48, 0x4a, 0x62, 0x03, 0x07, 0xbd, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0x9c, 0x1d, 0x1c, 0x17, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// QueryFullRequest gets the ai result meta data
	QueryFullRequest(ctx context.Context, in *QueryFullRequestReq, opts ...grpc.CallOption) (*QueryFullRequestRes, error)
	// QueryReward gets the reward given a block height
	QueryReward(ctx context.Context, in *QueryRewardReq, opts ...grpc.CallOption) (*QueryRewardRes, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) QueryFullRequest(ctx context.Context, in *QueryFullRequestReq, opts ...grpc.CallOption) (*QueryFullRequestRes, error) {
	out := new(QueryFullRequestRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.airesult.Query/QueryFullRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryReward(ctx context.Context, in *QueryRewardReq, opts ...grpc.CallOption) (*QueryRewardRes, error) {
	out := new(QueryRewardRes)
	err := c.cc.Invoke(ctx, "/oraichain.orai.airesult.Query/QueryReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// QueryFullRequest gets the ai result meta data
	QueryFullRequest(context.Context, *QueryFullRequestReq) (*QueryFullRequestRes, error)
	// QueryReward gets the reward given a block height
	QueryReward(context.Context, *QueryRewardReq) (*QueryRewardRes, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) QueryFullRequest(ctx context.Context, req *QueryFullRequestReq) (*QueryFullRequestRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFullRequest not implemented")
}
func (*UnimplementedQueryServer) QueryReward(ctx context.Context, req *QueryRewardReq) (*QueryRewardRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryReward not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_QueryFullRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFullRequestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryFullRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.airesult.Query/QueryFullRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryFullRequest(ctx, req.(*QueryFullRequestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRewardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oraichain.orai.airesult.Query/QueryReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryReward(ctx, req.(*QueryRewardReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oraichain.orai.airesult.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryFullRequest",
			Handler:    _Query_QueryFullRequest_Handler,
		},
		{
			MethodName: "QueryReward",
			Handler:    _Query_QueryReward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "airesult/types/query.proto",
}
