// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: location-processing.proto

package proto

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

// LocationProcessingClient is the client API for LocationProcessing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationProcessingClient interface {
	// Sends a greeting
	StoreUserLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*LocationResponse, error)
}

type locationProcessingClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationProcessingClient(cc grpc.ClientConnInterface) LocationProcessingClient {
	return &locationProcessingClient{cc}
}

func (c *locationProcessingClient) StoreUserLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*LocationResponse, error) {
	out := new(LocationResponse)
	err := c.cc.Invoke(ctx, "/proto.LocationProcessing/StoreUserLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationProcessingServer is the server API for LocationProcessing service.
// All implementations must embed UnimplementedLocationProcessingServer
// for forward compatibility
type LocationProcessingServer interface {
	// Sends a greeting
	StoreUserLocation(context.Context, *LocationRequest) (*LocationResponse, error)
	mustEmbedUnimplementedLocationProcessingServer()
}

// UnimplementedLocationProcessingServer must be embedded to have forward compatible implementations.
type UnimplementedLocationProcessingServer struct {
}

func (UnimplementedLocationProcessingServer) StoreUserLocation(context.Context, *LocationRequest) (*LocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreUserLocation not implemented")
}
func (UnimplementedLocationProcessingServer) mustEmbedUnimplementedLocationProcessingServer() {}

// UnsafeLocationProcessingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationProcessingServer will
// result in compilation errors.
type UnsafeLocationProcessingServer interface {
	mustEmbedUnimplementedLocationProcessingServer()
}

func RegisterLocationProcessingServer(s grpc.ServiceRegistrar, srv LocationProcessingServer) {
	s.RegisterService(&LocationProcessing_ServiceDesc, srv)
}

func _LocationProcessing_StoreUserLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationProcessingServer).StoreUserLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LocationProcessing/StoreUserLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationProcessingServer).StoreUserLocation(ctx, req.(*LocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationProcessing_ServiceDesc is the grpc.ServiceDesc for LocationProcessing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationProcessing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LocationProcessing",
	HandlerType: (*LocationProcessingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreUserLocation",
			Handler:    _LocationProcessing_StoreUserLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "location-processing.proto",
}