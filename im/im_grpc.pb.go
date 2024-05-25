// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: im/im.proto

package im

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ImService_OfflinePushMsg_FullMethodName = "/im.ImService/OfflinePushMsg"
)

// ImServiceClient is the client API for ImService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImServiceClient interface {
	OfflinePushMsg(ctx context.Context, in *OfflinePushMsgReq, opts ...grpc.CallOption) (*OfflinePushMsgResp, error)
}

type imServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImServiceClient(cc grpc.ClientConnInterface) ImServiceClient {
	return &imServiceClient{cc}
}

func (c *imServiceClient) OfflinePushMsg(ctx context.Context, in *OfflinePushMsgReq, opts ...grpc.CallOption) (*OfflinePushMsgResp, error) {
	out := new(OfflinePushMsgResp)
	err := c.cc.Invoke(ctx, ImService_OfflinePushMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImServiceServer is the server API for ImService service.
// All implementations must embed UnimplementedImServiceServer
// for forward compatibility
type ImServiceServer interface {
	OfflinePushMsg(context.Context, *OfflinePushMsgReq) (*OfflinePushMsgResp, error)
	mustEmbedUnimplementedImServiceServer()
}

// UnimplementedImServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImServiceServer struct {
}

func (UnimplementedImServiceServer) OfflinePushMsg(context.Context, *OfflinePushMsgReq) (*OfflinePushMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OfflinePushMsg not implemented")
}
func (UnimplementedImServiceServer) mustEmbedUnimplementedImServiceServer() {}

// UnsafeImServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImServiceServer will
// result in compilation errors.
type UnsafeImServiceServer interface {
	mustEmbedUnimplementedImServiceServer()
}

func RegisterImServiceServer(s grpc.ServiceRegistrar, srv ImServiceServer) {
	s.RegisterService(&ImService_ServiceDesc, srv)
}

func _ImService_OfflinePushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfflinePushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImServiceServer).OfflinePushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImService_OfflinePushMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImServiceServer).OfflinePushMsg(ctx, req.(*OfflinePushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ImService_ServiceDesc is the grpc.ServiceDesc for ImService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "im.ImService",
	HandlerType: (*ImServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OfflinePushMsg",
			Handler:    _ImService_OfflinePushMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im/im.proto",
}
