// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: authzed/api/v1/permission_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PermissionsService_ReadRelationships_FullMethodName    = "/authzed.api.v1.PermissionsService/ReadRelationships"
	PermissionsService_WriteRelationships_FullMethodName   = "/authzed.api.v1.PermissionsService/WriteRelationships"
	PermissionsService_DeleteRelationships_FullMethodName  = "/authzed.api.v1.PermissionsService/DeleteRelationships"
	PermissionsService_CheckPermission_FullMethodName      = "/authzed.api.v1.PermissionsService/CheckPermission"
	PermissionsService_CheckBulkPermissions_FullMethodName = "/authzed.api.v1.PermissionsService/CheckBulkPermissions"
	PermissionsService_ExpandPermissionTree_FullMethodName = "/authzed.api.v1.PermissionsService/ExpandPermissionTree"
	PermissionsService_LookupResources_FullMethodName      = "/authzed.api.v1.PermissionsService/LookupResources"
	PermissionsService_LookupSubjects_FullMethodName       = "/authzed.api.v1.PermissionsService/LookupSubjects"
)

// PermissionsServiceClient is the client API for PermissionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PermissionsService implements a set of RPCs that perform operations on
// relationships and permissions.
type PermissionsServiceClient interface {
	// ReadRelationships reads a set of the relationships matching one or more
	// filters.
	ReadRelationships(ctx context.Context, in *ReadRelationshipsRequest, opts ...grpc.CallOption) (PermissionsService_ReadRelationshipsClient, error)
	// WriteRelationships atomically writes and/or deletes a set of specified
	// relationships. An optional set of preconditions can be provided that must
	// be satisfied for the operation to commit.
	WriteRelationships(ctx context.Context, in *WriteRelationshipsRequest, opts ...grpc.CallOption) (*WriteRelationshipsResponse, error)
	// DeleteRelationships atomically bulk deletes all relationships matching the
	// provided filter. If no relationships match, none will be deleted and the
	// operation will succeed. An optional set of preconditions can be provided that must
	// be satisfied for the operation to commit.
	DeleteRelationships(ctx context.Context, in *DeleteRelationshipsRequest, opts ...grpc.CallOption) (*DeleteRelationshipsResponse, error)
	// CheckPermission determines for a given resource whether a subject computes
	// to having a permission or is a direct member of a particular relation.
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error)
	// CheckBulkPermissions evaluates the given list of permission checks
	// and returns the list of results.
	CheckBulkPermissions(ctx context.Context, in *CheckBulkPermissionsRequest, opts ...grpc.CallOption) (*CheckBulkPermissionsResponse, error)
	// ExpandPermissionTree reveals the graph structure for a resource's
	// permission or relation. This RPC does not recurse infinitely deep and may
	// require multiple calls to fully unnest a deeply nested graph.
	ExpandPermissionTree(ctx context.Context, in *ExpandPermissionTreeRequest, opts ...grpc.CallOption) (*ExpandPermissionTreeResponse, error)
	// LookupResources returns all the resources of a given type that a subject
	// can access whether via a computed permission or relation membership.
	LookupResources(ctx context.Context, in *LookupResourcesRequest, opts ...grpc.CallOption) (PermissionsService_LookupResourcesClient, error)
	// LookupSubjects returns all the subjects of a given type that
	// have access whether via a computed permission or relation membership.
	LookupSubjects(ctx context.Context, in *LookupSubjectsRequest, opts ...grpc.CallOption) (PermissionsService_LookupSubjectsClient, error)
}

type permissionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionsServiceClient(cc grpc.ClientConnInterface) PermissionsServiceClient {
	return &permissionsServiceClient{cc}
}

func (c *permissionsServiceClient) ReadRelationships(ctx context.Context, in *ReadRelationshipsRequest, opts ...grpc.CallOption) (PermissionsService_ReadRelationshipsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PermissionsService_ServiceDesc.Streams[0], PermissionsService_ReadRelationships_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &permissionsServiceReadRelationshipsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PermissionsService_ReadRelationshipsClient interface {
	Recv() (*ReadRelationshipsResponse, error)
	grpc.ClientStream
}

type permissionsServiceReadRelationshipsClient struct {
	grpc.ClientStream
}

func (x *permissionsServiceReadRelationshipsClient) Recv() (*ReadRelationshipsResponse, error) {
	m := new(ReadRelationshipsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *permissionsServiceClient) WriteRelationships(ctx context.Context, in *WriteRelationshipsRequest, opts ...grpc.CallOption) (*WriteRelationshipsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WriteRelationshipsResponse)
	err := c.cc.Invoke(ctx, PermissionsService_WriteRelationships_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) DeleteRelationships(ctx context.Context, in *DeleteRelationshipsRequest, opts ...grpc.CallOption) (*DeleteRelationshipsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRelationshipsResponse)
	err := c.cc.Invoke(ctx, PermissionsService_DeleteRelationships_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckPermissionResponse)
	err := c.cc.Invoke(ctx, PermissionsService_CheckPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) CheckBulkPermissions(ctx context.Context, in *CheckBulkPermissionsRequest, opts ...grpc.CallOption) (*CheckBulkPermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckBulkPermissionsResponse)
	err := c.cc.Invoke(ctx, PermissionsService_CheckBulkPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) ExpandPermissionTree(ctx context.Context, in *ExpandPermissionTreeRequest, opts ...grpc.CallOption) (*ExpandPermissionTreeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpandPermissionTreeResponse)
	err := c.cc.Invoke(ctx, PermissionsService_ExpandPermissionTree_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) LookupResources(ctx context.Context, in *LookupResourcesRequest, opts ...grpc.CallOption) (PermissionsService_LookupResourcesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PermissionsService_ServiceDesc.Streams[1], PermissionsService_LookupResources_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &permissionsServiceLookupResourcesClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PermissionsService_LookupResourcesClient interface {
	Recv() (*LookupResourcesResponse, error)
	grpc.ClientStream
}

type permissionsServiceLookupResourcesClient struct {
	grpc.ClientStream
}

func (x *permissionsServiceLookupResourcesClient) Recv() (*LookupResourcesResponse, error) {
	m := new(LookupResourcesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *permissionsServiceClient) LookupSubjects(ctx context.Context, in *LookupSubjectsRequest, opts ...grpc.CallOption) (PermissionsService_LookupSubjectsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PermissionsService_ServiceDesc.Streams[2], PermissionsService_LookupSubjects_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &permissionsServiceLookupSubjectsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PermissionsService_LookupSubjectsClient interface {
	Recv() (*LookupSubjectsResponse, error)
	grpc.ClientStream
}

type permissionsServiceLookupSubjectsClient struct {
	grpc.ClientStream
}

func (x *permissionsServiceLookupSubjectsClient) Recv() (*LookupSubjectsResponse, error) {
	m := new(LookupSubjectsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PermissionsServiceServer is the server API for PermissionsService service.
// All implementations must embed UnimplementedPermissionsServiceServer
// for forward compatibility
//
// PermissionsService implements a set of RPCs that perform operations on
// relationships and permissions.
type PermissionsServiceServer interface {
	// ReadRelationships reads a set of the relationships matching one or more
	// filters.
	ReadRelationships(*ReadRelationshipsRequest, PermissionsService_ReadRelationshipsServer) error
	// WriteRelationships atomically writes and/or deletes a set of specified
	// relationships. An optional set of preconditions can be provided that must
	// be satisfied for the operation to commit.
	WriteRelationships(context.Context, *WriteRelationshipsRequest) (*WriteRelationshipsResponse, error)
	// DeleteRelationships atomically bulk deletes all relationships matching the
	// provided filter. If no relationships match, none will be deleted and the
	// operation will succeed. An optional set of preconditions can be provided that must
	// be satisfied for the operation to commit.
	DeleteRelationships(context.Context, *DeleteRelationshipsRequest) (*DeleteRelationshipsResponse, error)
	// CheckPermission determines for a given resource whether a subject computes
	// to having a permission or is a direct member of a particular relation.
	CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error)
	// CheckBulkPermissions evaluates the given list of permission checks
	// and returns the list of results.
	CheckBulkPermissions(context.Context, *CheckBulkPermissionsRequest) (*CheckBulkPermissionsResponse, error)
	// ExpandPermissionTree reveals the graph structure for a resource's
	// permission or relation. This RPC does not recurse infinitely deep and may
	// require multiple calls to fully unnest a deeply nested graph.
	ExpandPermissionTree(context.Context, *ExpandPermissionTreeRequest) (*ExpandPermissionTreeResponse, error)
	// LookupResources returns all the resources of a given type that a subject
	// can access whether via a computed permission or relation membership.
	LookupResources(*LookupResourcesRequest, PermissionsService_LookupResourcesServer) error
	// LookupSubjects returns all the subjects of a given type that
	// have access whether via a computed permission or relation membership.
	LookupSubjects(*LookupSubjectsRequest, PermissionsService_LookupSubjectsServer) error
	mustEmbedUnimplementedPermissionsServiceServer()
}

// UnimplementedPermissionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionsServiceServer struct {
}

func (UnimplementedPermissionsServiceServer) ReadRelationships(*ReadRelationshipsRequest, PermissionsService_ReadRelationshipsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadRelationships not implemented")
}
func (UnimplementedPermissionsServiceServer) WriteRelationships(context.Context, *WriteRelationshipsRequest) (*WriteRelationshipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteRelationships not implemented")
}
func (UnimplementedPermissionsServiceServer) DeleteRelationships(context.Context, *DeleteRelationshipsRequest) (*DeleteRelationshipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelationships not implemented")
}
func (UnimplementedPermissionsServiceServer) CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedPermissionsServiceServer) CheckBulkPermissions(context.Context, *CheckBulkPermissionsRequest) (*CheckBulkPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBulkPermissions not implemented")
}
func (UnimplementedPermissionsServiceServer) ExpandPermissionTree(context.Context, *ExpandPermissionTreeRequest) (*ExpandPermissionTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpandPermissionTree not implemented")
}
func (UnimplementedPermissionsServiceServer) LookupResources(*LookupResourcesRequest, PermissionsService_LookupResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method LookupResources not implemented")
}
func (UnimplementedPermissionsServiceServer) LookupSubjects(*LookupSubjectsRequest, PermissionsService_LookupSubjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method LookupSubjects not implemented")
}
func (UnimplementedPermissionsServiceServer) mustEmbedUnimplementedPermissionsServiceServer() {}

// UnsafePermissionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionsServiceServer will
// result in compilation errors.
type UnsafePermissionsServiceServer interface {
	mustEmbedUnimplementedPermissionsServiceServer()
}

func RegisterPermissionsServiceServer(s grpc.ServiceRegistrar, srv PermissionsServiceServer) {
	s.RegisterService(&PermissionsService_ServiceDesc, srv)
}

func _PermissionsService_ReadRelationships_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRelationshipsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PermissionsServiceServer).ReadRelationships(m, &permissionsServiceReadRelationshipsServer{ServerStream: stream})
}

type PermissionsService_ReadRelationshipsServer interface {
	Send(*ReadRelationshipsResponse) error
	grpc.ServerStream
}

type permissionsServiceReadRelationshipsServer struct {
	grpc.ServerStream
}

func (x *permissionsServiceReadRelationshipsServer) Send(m *ReadRelationshipsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PermissionsService_WriteRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRelationshipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).WriteRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_WriteRelationships_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).WriteRelationships(ctx, req.(*WriteRelationshipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_DeleteRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelationshipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).DeleteRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_DeleteRelationships_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).DeleteRelationships(ctx, req.(*DeleteRelationshipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_CheckPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_CheckBulkPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckBulkPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).CheckBulkPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_CheckBulkPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).CheckBulkPermissions(ctx, req.(*CheckBulkPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_ExpandPermissionTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandPermissionTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).ExpandPermissionTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_ExpandPermissionTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).ExpandPermissionTree(ctx, req.(*ExpandPermissionTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_LookupResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LookupResourcesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PermissionsServiceServer).LookupResources(m, &permissionsServiceLookupResourcesServer{ServerStream: stream})
}

type PermissionsService_LookupResourcesServer interface {
	Send(*LookupResourcesResponse) error
	grpc.ServerStream
}

type permissionsServiceLookupResourcesServer struct {
	grpc.ServerStream
}

func (x *permissionsServiceLookupResourcesServer) Send(m *LookupResourcesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PermissionsService_LookupSubjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LookupSubjectsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PermissionsServiceServer).LookupSubjects(m, &permissionsServiceLookupSubjectsServer{ServerStream: stream})
}

type PermissionsService_LookupSubjectsServer interface {
	Send(*LookupSubjectsResponse) error
	grpc.ServerStream
}

type permissionsServiceLookupSubjectsServer struct {
	grpc.ServerStream
}

func (x *permissionsServiceLookupSubjectsServer) Send(m *LookupSubjectsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PermissionsService_ServiceDesc is the grpc.ServiceDesc for PermissionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authzed.api.v1.PermissionsService",
	HandlerType: (*PermissionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteRelationships",
			Handler:    _PermissionsService_WriteRelationships_Handler,
		},
		{
			MethodName: "DeleteRelationships",
			Handler:    _PermissionsService_DeleteRelationships_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _PermissionsService_CheckPermission_Handler,
		},
		{
			MethodName: "CheckBulkPermissions",
			Handler:    _PermissionsService_CheckBulkPermissions_Handler,
		},
		{
			MethodName: "ExpandPermissionTree",
			Handler:    _PermissionsService_ExpandPermissionTree_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadRelationships",
			Handler:       _PermissionsService_ReadRelationships_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LookupResources",
			Handler:       _PermissionsService_LookupResources_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LookupSubjects",
			Handler:       _PermissionsService_LookupSubjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "authzed/api/v1/permission_service.proto",
}
