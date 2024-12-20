// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: internal/proto/storage/storage.proto

package storage

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Storage_SavePassword_FullMethodName   = "/models.Storage/SavePassword"
	Storage_GetPassword_FullMethodName    = "/models.Storage/GetPassword"
	Storage_DeletePassword_FullMethodName = "/models.Storage/DeletePassword"
	Storage_SaveText_FullMethodName       = "/models.Storage/SaveText"
	Storage_GetText_FullMethodName        = "/models.Storage/GetText"
	Storage_DeleteText_FullMethodName     = "/models.Storage/DeleteText"
	Storage_SaveBinary_FullMethodName     = "/models.Storage/SaveBinary"
	Storage_GetBinary_FullMethodName      = "/models.Storage/GetBinary"
	Storage_DeleteBinary_FullMethodName   = "/models.Storage/DeleteBinary"
	Storage_SaveBankCard_FullMethodName   = "/models.Storage/SaveBankCard"
	Storage_GetBankCard_FullMethodName    = "/models.Storage/GetBankCard"
	Storage_DeleteBankCard_FullMethodName = "/models.Storage/DeleteBankCard"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	SavePassword(ctx context.Context, in *SavePasswordRequest, opts ...grpc.CallOption) (*SavePasswordResponse, error)
	GetPassword(ctx context.Context, in *GetPasswordRequest, opts ...grpc.CallOption) (*GetPasswordResponse, error)
	DeletePassword(ctx context.Context, in *DeletePasswordRequest, opts ...grpc.CallOption) (*DeletePasswordResponse, error)
	SaveText(ctx context.Context, in *SaveTextRequest, opts ...grpc.CallOption) (*SaveTextResponse, error)
	GetText(ctx context.Context, in *GetTextRequest, opts ...grpc.CallOption) (*GetTextResponse, error)
	DeleteText(ctx context.Context, in *DeleteTextRequest, opts ...grpc.CallOption) (*DeleteTextResponse, error)
	SaveBinary(ctx context.Context, in *SaveBinaryRequest, opts ...grpc.CallOption) (*SaveBinaryResponse, error)
	GetBinary(ctx context.Context, in *GetBinaryRequest, opts ...grpc.CallOption) (*GetBinaryResponse, error)
	DeleteBinary(ctx context.Context, in *DeleteBinaryRequest, opts ...grpc.CallOption) (*DeleteBinaryResponse, error)
	SaveBankCard(ctx context.Context, in *SaveBankCardRequest, opts ...grpc.CallOption) (*SaveBankCardResponse, error)
	GetBankCard(ctx context.Context, in *GetBankCardRequest, opts ...grpc.CallOption) (*GetBankCardResponse, error)
	DeleteBankCard(ctx context.Context, in *DeleteBankCardRequest, opts ...grpc.CallOption) (*DeleteBankCardResponse, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) SavePassword(ctx context.Context, in *SavePasswordRequest, opts ...grpc.CallOption) (*SavePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SavePasswordResponse)
	err := c.cc.Invoke(ctx, Storage_SavePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetPassword(ctx context.Context, in *GetPasswordRequest, opts ...grpc.CallOption) (*GetPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPasswordResponse)
	err := c.cc.Invoke(ctx, Storage_GetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeletePassword(ctx context.Context, in *DeletePasswordRequest, opts ...grpc.CallOption) (*DeletePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePasswordResponse)
	err := c.cc.Invoke(ctx, Storage_DeletePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) SaveText(ctx context.Context, in *SaveTextRequest, opts ...grpc.CallOption) (*SaveTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveTextResponse)
	err := c.cc.Invoke(ctx, Storage_SaveText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetText(ctx context.Context, in *GetTextRequest, opts ...grpc.CallOption) (*GetTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTextResponse)
	err := c.cc.Invoke(ctx, Storage_GetText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeleteText(ctx context.Context, in *DeleteTextRequest, opts ...grpc.CallOption) (*DeleteTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTextResponse)
	err := c.cc.Invoke(ctx, Storage_DeleteText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) SaveBinary(ctx context.Context, in *SaveBinaryRequest, opts ...grpc.CallOption) (*SaveBinaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveBinaryResponse)
	err := c.cc.Invoke(ctx, Storage_SaveBinary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetBinary(ctx context.Context, in *GetBinaryRequest, opts ...grpc.CallOption) (*GetBinaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBinaryResponse)
	err := c.cc.Invoke(ctx, Storage_GetBinary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeleteBinary(ctx context.Context, in *DeleteBinaryRequest, opts ...grpc.CallOption) (*DeleteBinaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBinaryResponse)
	err := c.cc.Invoke(ctx, Storage_DeleteBinary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) SaveBankCard(ctx context.Context, in *SaveBankCardRequest, opts ...grpc.CallOption) (*SaveBankCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveBankCardResponse)
	err := c.cc.Invoke(ctx, Storage_SaveBankCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetBankCard(ctx context.Context, in *GetBankCardRequest, opts ...grpc.CallOption) (*GetBankCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBankCardResponse)
	err := c.cc.Invoke(ctx, Storage_GetBankCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeleteBankCard(ctx context.Context, in *DeleteBankCardRequest, opts ...grpc.CallOption) (*DeleteBankCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBankCardResponse)
	err := c.cc.Invoke(ctx, Storage_DeleteBankCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility.
type StorageServer interface {
	SavePassword(context.Context, *SavePasswordRequest) (*SavePasswordResponse, error)
	GetPassword(context.Context, *GetPasswordRequest) (*GetPasswordResponse, error)
	DeletePassword(context.Context, *DeletePasswordRequest) (*DeletePasswordResponse, error)
	SaveText(context.Context, *SaveTextRequest) (*SaveTextResponse, error)
	GetText(context.Context, *GetTextRequest) (*GetTextResponse, error)
	DeleteText(context.Context, *DeleteTextRequest) (*DeleteTextResponse, error)
	SaveBinary(context.Context, *SaveBinaryRequest) (*SaveBinaryResponse, error)
	GetBinary(context.Context, *GetBinaryRequest) (*GetBinaryResponse, error)
	DeleteBinary(context.Context, *DeleteBinaryRequest) (*DeleteBinaryResponse, error)
	SaveBankCard(context.Context, *SaveBankCardRequest) (*SaveBankCardResponse, error)
	GetBankCard(context.Context, *GetBankCardRequest) (*GetBankCardResponse, error)
	DeleteBankCard(context.Context, *DeleteBankCardRequest) (*DeleteBankCardResponse, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStorageServer struct{}

func (UnimplementedStorageServer) SavePassword(context.Context, *SavePasswordRequest) (*SavePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePassword not implemented")
}
func (UnimplementedStorageServer) GetPassword(context.Context, *GetPasswordRequest) (*GetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}
func (UnimplementedStorageServer) DeletePassword(context.Context, *DeletePasswordRequest) (*DeletePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePassword not implemented")
}
func (UnimplementedStorageServer) SaveText(context.Context, *SaveTextRequest) (*SaveTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveText not implemented")
}
func (UnimplementedStorageServer) GetText(context.Context, *GetTextRequest) (*GetTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetText not implemented")
}
func (UnimplementedStorageServer) DeleteText(context.Context, *DeleteTextRequest) (*DeleteTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteText not implemented")
}
func (UnimplementedStorageServer) SaveBinary(context.Context, *SaveBinaryRequest) (*SaveBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBinary not implemented")
}
func (UnimplementedStorageServer) GetBinary(context.Context, *GetBinaryRequest) (*GetBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBinary not implemented")
}
func (UnimplementedStorageServer) DeleteBinary(context.Context, *DeleteBinaryRequest) (*DeleteBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBinary not implemented")
}
func (UnimplementedStorageServer) SaveBankCard(context.Context, *SaveBankCardRequest) (*SaveBankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBankCard not implemented")
}
func (UnimplementedStorageServer) GetBankCard(context.Context, *GetBankCardRequest) (*GetBankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankCard not implemented")
}
func (UnimplementedStorageServer) DeleteBankCard(context.Context, *DeleteBankCardRequest) (*DeleteBankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBankCard not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}
func (UnimplementedStorageServer) testEmbeddedByValue()                 {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	// If the following call pancis, it indicates UnimplementedStorageServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_SavePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SavePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SavePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SavePassword(ctx, req.(*SavePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetPassword(ctx, req.(*GetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeletePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeletePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_DeletePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeletePassword(ctx, req.(*DeletePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_SaveText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SaveText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveText(ctx, req.(*SaveTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GetText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetText(ctx, req.(*GetTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeleteText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeleteText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_DeleteText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeleteText(ctx, req.(*DeleteTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_SaveBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SaveBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveBinary(ctx, req.(*SaveBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GetBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetBinary(ctx, req.(*GetBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeleteBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeleteBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_DeleteBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeleteBinary(ctx, req.(*DeleteBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_SaveBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBankCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SaveBankCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveBankCard(ctx, req.(*SaveBankCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GetBankCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetBankCard(ctx, req.(*GetBankCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeleteBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBankCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeleteBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_DeleteBankCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeleteBankCard(ctx, req.(*DeleteBankCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SavePassword",
			Handler:    _Storage_SavePassword_Handler,
		},
		{
			MethodName: "GetPassword",
			Handler:    _Storage_GetPassword_Handler,
		},
		{
			MethodName: "DeletePassword",
			Handler:    _Storage_DeletePassword_Handler,
		},
		{
			MethodName: "SaveText",
			Handler:    _Storage_SaveText_Handler,
		},
		{
			MethodName: "GetText",
			Handler:    _Storage_GetText_Handler,
		},
		{
			MethodName: "DeleteText",
			Handler:    _Storage_DeleteText_Handler,
		},
		{
			MethodName: "SaveBinary",
			Handler:    _Storage_SaveBinary_Handler,
		},
		{
			MethodName: "GetBinary",
			Handler:    _Storage_GetBinary_Handler,
		},
		{
			MethodName: "DeleteBinary",
			Handler:    _Storage_DeleteBinary_Handler,
		},
		{
			MethodName: "SaveBankCard",
			Handler:    _Storage_SaveBankCard_Handler,
		},
		{
			MethodName: "GetBankCard",
			Handler:    _Storage_GetBankCard_Handler,
		},
		{
			MethodName: "DeleteBankCard",
			Handler:    _Storage_DeleteBankCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/storage/storage.proto",
}
