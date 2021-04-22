// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package event

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	SaveEvent(ctx context.Context, in *SaveEventRequest, opts ...grpc.CallOption) (*SaveReponse, error)
	QueryEvent(ctx context.Context, in *QueryEventRequest, opts ...grpc.CallOption) (*OperateEventSet, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) SaveEvent(ctx context.Context, in *SaveEventRequest, opts ...grpc.CallOption) (*SaveReponse, error) {
	out := new(SaveReponse)
	err := c.cc.Invoke(ctx, "/eventbox.event.Service/SaveEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryEvent(ctx context.Context, in *QueryEventRequest, opts ...grpc.CallOption) (*OperateEventSet, error) {
	out := new(OperateEventSet)
	err := c.cc.Invoke(ctx, "/eventbox.event.Service/QueryEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	SaveEvent(context.Context, *SaveEventRequest) (*SaveReponse, error)
	QueryEvent(context.Context, *QueryEventRequest) (*OperateEventSet, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) SaveEvent(context.Context, *SaveEventRequest) (*SaveReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveEvent not implemented")
}
func (UnimplementedServiceServer) QueryEvent(context.Context, *QueryEventRequest) (*OperateEventSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryEvent not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_SaveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SaveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventbox.event.Service/SaveEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SaveEvent(ctx, req.(*SaveEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventbox.event.Service/QueryEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryEvent(ctx, req.(*QueryEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventbox.event.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveEvent",
			Handler:    _Service_SaveEvent_Handler,
		},
		{
			MethodName: "QueryEvent",
			Handler:    _Service_QueryEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/event/pb/service.proto",
}