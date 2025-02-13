// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package snapshots

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SnapshotsClient is the client API for Snapshots service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnapshotsClient interface {
	Prepare(ctx context.Context, in *PrepareSnapshotRequest, opts ...grpc.CallOption) (*PrepareSnapshotResponse, error)
	View(ctx context.Context, in *ViewSnapshotRequest, opts ...grpc.CallOption) (*ViewSnapshotResponse, error)
	Mounts(ctx context.Context, in *MountsRequest, opts ...grpc.CallOption) (*MountsResponse, error)
	Commit(ctx context.Context, in *CommitSnapshotRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Remove(ctx context.Context, in *RemoveSnapshotRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Stat(ctx context.Context, in *StatSnapshotRequest, opts ...grpc.CallOption) (*StatSnapshotResponse, error)
	Update(ctx context.Context, in *UpdateSnapshotRequest, opts ...grpc.CallOption) (*UpdateSnapshotResponse, error)
	List(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (Snapshots_ListClient, error)
	Usage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*UsageResponse, error)
	Cleanup(ctx context.Context, in *CleanupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type snapshotsClient struct {
	cc grpc.ClientConnInterface
}

func NewSnapshotsClient(cc grpc.ClientConnInterface) SnapshotsClient {
	return &snapshotsClient{cc}
}

func (c *snapshotsClient) Prepare(ctx context.Context, in *PrepareSnapshotRequest, opts ...grpc.CallOption) (*PrepareSnapshotResponse, error) {
	out := new(PrepareSnapshotResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.snapshots.v1.Snapshots/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) View(ctx context.Context, in *ViewSnapshotRequest, opts ...grpc.CallOption) (*ViewSnapshotResponse, error) {
	out := new(ViewSnapshotResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.snapshots.v1.Snapshots/View", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) Mounts(ctx context.Context, in *MountsRequest, opts ...grpc.CallOption) (*MountsResponse, error) {
	out := new(MountsResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.snapshots.v1.Snapshots/Mounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) Commit(ctx context.Context, in *CommitSnapshotRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.snapshots.v1.Snapshots/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) Remove(ctx context.Context, in *RemoveSnapshotRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.snapshots.v1.Snapshots/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) Stat(ctx context.Context, in *StatSnapshotRequest, opts ...grpc.CallOption) (*StatSnapshotResponse, error) {
	out := new(StatSnapshotResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.snapshots.v1.Snapshots/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) Update(ctx context.Context, in *UpdateSnapshotRequest, opts ...grpc.CallOption) (*UpdateSnapshotResponse, error) {
	out := new(UpdateSnapshotResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.snapshots.v1.Snapshots/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) List(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (Snapshots_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Snapshots_ServiceDesc.Streams[0], "/containerd.services.snapshots.v1.Snapshots/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &snapshotsListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Snapshots_ListClient interface {
	Recv() (*ListSnapshotsResponse, error)
	grpc.ClientStream
}

type snapshotsListClient struct {
	grpc.ClientStream
}

func (x *snapshotsListClient) Recv() (*ListSnapshotsResponse, error) {
	m := new(ListSnapshotsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *snapshotsClient) Usage(ctx context.Context, in *UsageRequest, opts ...grpc.CallOption) (*UsageResponse, error) {
	out := new(UsageResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.snapshots.v1.Snapshots/Usage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) Cleanup(ctx context.Context, in *CleanupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.snapshots.v1.Snapshots/Cleanup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnapshotsServer is the server API for Snapshots service.
// All implementations must embed UnimplementedSnapshotsServer
// for forward compatibility
type SnapshotsServer interface {
	Prepare(context.Context, *PrepareSnapshotRequest) (*PrepareSnapshotResponse, error)
	View(context.Context, *ViewSnapshotRequest) (*ViewSnapshotResponse, error)
	Mounts(context.Context, *MountsRequest) (*MountsResponse, error)
	Commit(context.Context, *CommitSnapshotRequest) (*empty.Empty, error)
	Remove(context.Context, *RemoveSnapshotRequest) (*empty.Empty, error)
	Stat(context.Context, *StatSnapshotRequest) (*StatSnapshotResponse, error)
	Update(context.Context, *UpdateSnapshotRequest) (*UpdateSnapshotResponse, error)
	List(*ListSnapshotsRequest, Snapshots_ListServer) error
	Usage(context.Context, *UsageRequest) (*UsageResponse, error)
	Cleanup(context.Context, *CleanupRequest) (*empty.Empty, error)
	mustEmbedUnimplementedSnapshotsServer()
}

// UnimplementedSnapshotsServer must be embedded to have forward compatible implementations.
type UnimplementedSnapshotsServer struct {
}

func (UnimplementedSnapshotsServer) Prepare(context.Context, *PrepareSnapshotRequest) (*PrepareSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedSnapshotsServer) View(context.Context, *ViewSnapshotRequest) (*ViewSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedSnapshotsServer) Mounts(context.Context, *MountsRequest) (*MountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mounts not implemented")
}
func (UnimplementedSnapshotsServer) Commit(context.Context, *CommitSnapshotRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedSnapshotsServer) Remove(context.Context, *RemoveSnapshotRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedSnapshotsServer) Stat(context.Context, *StatSnapshotRequest) (*StatSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedSnapshotsServer) Update(context.Context, *UpdateSnapshotRequest) (*UpdateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSnapshotsServer) List(*ListSnapshotsRequest, Snapshots_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSnapshotsServer) Usage(context.Context, *UsageRequest) (*UsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Usage not implemented")
}
func (UnimplementedSnapshotsServer) Cleanup(context.Context, *CleanupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cleanup not implemented")
}
func (UnimplementedSnapshotsServer) mustEmbedUnimplementedSnapshotsServer() {}

// UnsafeSnapshotsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnapshotsServer will
// result in compilation errors.
type UnsafeSnapshotsServer interface {
	mustEmbedUnimplementedSnapshotsServer()
}

func RegisterSnapshotsServer(s grpc.ServiceRegistrar, srv SnapshotsServer) {
	s.RegisterService(&Snapshots_ServiceDesc, srv)
}

func _Snapshots_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.snapshots.v1.Snapshots/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Prepare(ctx, req.(*PrepareSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.snapshots.v1.Snapshots/View",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).View(ctx, req.(*ViewSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_Mounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Mounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.snapshots.v1.Snapshots/Mounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Mounts(ctx, req.(*MountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.snapshots.v1.Snapshots/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Commit(ctx, req.(*CommitSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.snapshots.v1.Snapshots/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Remove(ctx, req.(*RemoveSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.snapshots.v1.Snapshots/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Stat(ctx, req.(*StatSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.snapshots.v1.Snapshots/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Update(ctx, req.(*UpdateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListSnapshotsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SnapshotsServer).List(m, &snapshotsListServer{stream})
}

type Snapshots_ListServer interface {
	Send(*ListSnapshotsResponse) error
	grpc.ServerStream
}

type snapshotsListServer struct {
	grpc.ServerStream
}

func (x *snapshotsListServer) Send(m *ListSnapshotsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Snapshots_Usage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Usage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.snapshots.v1.Snapshots/Usage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Usage(ctx, req.(*UsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_Cleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Cleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.snapshots.v1.Snapshots/Cleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Cleanup(ctx, req.(*CleanupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Snapshots_ServiceDesc is the grpc.ServiceDesc for Snapshots service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Snapshots_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.snapshots.v1.Snapshots",
	HandlerType: (*SnapshotsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _Snapshots_Prepare_Handler,
		},
		{
			MethodName: "View",
			Handler:    _Snapshots_View_Handler,
		},
		{
			MethodName: "Mounts",
			Handler:    _Snapshots_Mounts_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Snapshots_Commit_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Snapshots_Remove_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _Snapshots_Stat_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Snapshots_Update_Handler,
		},
		{
			MethodName: "Usage",
			Handler:    _Snapshots_Usage_Handler,
		},
		{
			MethodName: "Cleanup",
			Handler:    _Snapshots_Cleanup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Snapshots_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/containerd/containerd/api/services/snapshots/v1/snapshots.proto",
}
