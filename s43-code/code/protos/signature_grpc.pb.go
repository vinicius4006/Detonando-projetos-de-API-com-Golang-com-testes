// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: signature.proto

package protos

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

// SignVerifyClient is the client API for SignVerify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignVerifyClient interface {
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
}

type signVerifyClient struct {
	cc grpc.ClientConnInterface
}

func NewSignVerifyClient(cc grpc.ClientConnInterface) SignVerifyClient {
	return &signVerifyClient{cc}
}

func (c *signVerifyClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/signverify.SignVerify/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignVerifyServer is the server API for SignVerify service.
// All implementations must embed UnimplementedSignVerifyServer
// for forward compatibility
type SignVerifyServer interface {
	Sign(context.Context, *SignRequest) (*SignResponse, error)
	mustEmbedUnimplementedSignVerifyServer()
}

// UnimplementedSignVerifyServer must be embedded to have forward compatible implementations.
type UnimplementedSignVerifyServer struct {
}

func (UnimplementedSignVerifyServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedSignVerifyServer) mustEmbedUnimplementedSignVerifyServer() {}

// UnsafeSignVerifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignVerifyServer will
// result in compilation errors.
type UnsafeSignVerifyServer interface {
	mustEmbedUnimplementedSignVerifyServer()
}

func RegisterSignVerifyServer(s grpc.ServiceRegistrar, srv SignVerifyServer) {
	s.RegisterService(&SignVerify_ServiceDesc, srv)
}

func _SignVerify_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignVerifyServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signverify.SignVerify/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignVerifyServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignVerify_ServiceDesc is the grpc.ServiceDesc for SignVerify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignVerify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "signverify.SignVerify",
	HandlerType: (*SignVerifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _SignVerify_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signature.proto",
}
