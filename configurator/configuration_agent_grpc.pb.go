// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: configurator/configuration_agent.proto

package configurator

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

// ConfigurationAgentServiceClient is the client API for ConfigurationAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigurationAgentServiceClient interface {
	AssignIPToInterface(ctx context.Context, in *IPAssignment, opts ...grpc.CallOption) (*Result, error)
}

type configurationAgentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationAgentServiceClient(cc grpc.ClientConnInterface) ConfigurationAgentServiceClient {
	return &configurationAgentServiceClient{cc}
}

func (c *configurationAgentServiceClient) AssignIPToInterface(ctx context.Context, in *IPAssignment, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/configurator.ConfigurationAgentService/AssignIPToInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationAgentServiceServer is the server API for ConfigurationAgentService service.
// All implementations must embed UnimplementedConfigurationAgentServiceServer
// for forward compatibility
type ConfigurationAgentServiceServer interface {
	AssignIPToInterface(context.Context, *IPAssignment) (*Result, error)
	mustEmbedUnimplementedConfigurationAgentServiceServer()
}

// UnimplementedConfigurationAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigurationAgentServiceServer struct {
}

func (UnimplementedConfigurationAgentServiceServer) AssignIPToInterface(context.Context, *IPAssignment) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignIPToInterface not implemented")
}
func (UnimplementedConfigurationAgentServiceServer) mustEmbedUnimplementedConfigurationAgentServiceServer() {
}

// UnsafeConfigurationAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigurationAgentServiceServer will
// result in compilation errors.
type UnsafeConfigurationAgentServiceServer interface {
	mustEmbedUnimplementedConfigurationAgentServiceServer()
}

func RegisterConfigurationAgentServiceServer(s grpc.ServiceRegistrar, srv ConfigurationAgentServiceServer) {
	s.RegisterService(&ConfigurationAgentService_ServiceDesc, srv)
}

func _ConfigurationAgentService_AssignIPToInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPAssignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationAgentServiceServer).AssignIPToInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configurator.ConfigurationAgentService/AssignIPToInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationAgentServiceServer).AssignIPToInterface(ctx, req.(*IPAssignment))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigurationAgentService_ServiceDesc is the grpc.ServiceDesc for ConfigurationAgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigurationAgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "configurator.ConfigurationAgentService",
	HandlerType: (*ConfigurationAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignIPToInterface",
			Handler:    _ConfigurationAgentService_AssignIPToInterface_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configurator/configuration_agent.proto",
}
