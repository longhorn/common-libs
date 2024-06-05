// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: profilerrpc/profilerrpc.proto

package profilerrpc

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
	Profiler_ProfilerOP_FullMethodName = "/profilerrpc.Profiler/ProfilerOP"
)

// ProfilerClient is the client API for Profiler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilerClient interface {
	ProfilerOP(ctx context.Context, in *ProfilerOPRequest, opts ...grpc.CallOption) (*ProfilerOPResponse, error)
}

type profilerClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilerClient(cc grpc.ClientConnInterface) ProfilerClient {
	return &profilerClient{cc}
}

func (c *profilerClient) ProfilerOP(ctx context.Context, in *ProfilerOPRequest, opts ...grpc.CallOption) (*ProfilerOPResponse, error) {
	out := new(ProfilerOPResponse)
	err := c.cc.Invoke(ctx, Profiler_ProfilerOP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilerServer is the server API for Profiler service.
// All implementations must embed UnimplementedProfilerServer
// for forward compatibility
type ProfilerServer interface {
	ProfilerOP(context.Context, *ProfilerOPRequest) (*ProfilerOPResponse, error)
	mustEmbedUnimplementedProfilerServer()
}

// UnimplementedProfilerServer must be embedded to have forward compatible implementations.
type UnimplementedProfilerServer struct {
}

func (UnimplementedProfilerServer) ProfilerOP(context.Context, *ProfilerOPRequest) (*ProfilerOPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProfilerOP not implemented")
}
func (UnimplementedProfilerServer) mustEmbedUnimplementedProfilerServer() {}

// UnsafeProfilerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilerServer will
// result in compilation errors.
type UnsafeProfilerServer interface {
	mustEmbedUnimplementedProfilerServer()
}

func RegisterProfilerServer(s grpc.ServiceRegistrar, srv ProfilerServer) {
	s.RegisterService(&Profiler_ServiceDesc, srv)
}

func _Profiler_ProfilerOP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfilerOPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerServer).ProfilerOP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profiler_ProfilerOP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerServer).ProfilerOP(ctx, req.(*ProfilerOPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Profiler_ServiceDesc is the grpc.ServiceDesc for Profiler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profiler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profilerrpc.Profiler",
	HandlerType: (*ProfilerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfilerOP",
			Handler:    _Profiler_ProfilerOP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profilerrpc/profilerrpc.proto",
}
