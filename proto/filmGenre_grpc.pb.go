// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: proto/filmGenre.proto

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

// FilmGenreServiceClient is the client API for FilmGenreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilmGenreServiceClient interface {
	CreateFilm(ctx context.Context, in *FilmData, opts ...grpc.CallOption) (*FilmResponse, error)
	GetAllFilms(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*GetFilmResponse, error)
	GetFilmById(ctx context.Context, in *ID, opts ...grpc.CallOption) (*FilmResponse, error)
	UpDateFilm(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OperationStatus, error)
	DeleteFilm(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OperationStatus, error)
	CreateGenre(ctx context.Context, in *GenreData, opts ...grpc.CallOption) (*GenreResponse, error)
	GetGenreIdsByName(ctx context.Context, in *GenreData, opts ...grpc.CallOption) (*GenreIds, error)
	GetAllGenres(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*GetGenresResponse, error)
	DeleteGenre(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OperationStatus, error)
}

type filmGenreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilmGenreServiceClient(cc grpc.ClientConnInterface) FilmGenreServiceClient {
	return &filmGenreServiceClient{cc}
}

func (c *filmGenreServiceClient) CreateFilm(ctx context.Context, in *FilmData, opts ...grpc.CallOption) (*FilmResponse, error) {
	out := new(FilmResponse)
	err := c.cc.Invoke(ctx, "/filmGenre_service.FilmGenreService/CreateFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmGenreServiceClient) GetAllFilms(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*GetFilmResponse, error) {
	out := new(GetFilmResponse)
	err := c.cc.Invoke(ctx, "/filmGenre_service.FilmGenreService/GetAllFilms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmGenreServiceClient) GetFilmById(ctx context.Context, in *ID, opts ...grpc.CallOption) (*FilmResponse, error) {
	out := new(FilmResponse)
	err := c.cc.Invoke(ctx, "/filmGenre_service.FilmGenreService/GetFilmById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmGenreServiceClient) UpDateFilm(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/filmGenre_service.FilmGenreService/UpDateFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmGenreServiceClient) DeleteFilm(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/filmGenre_service.FilmGenreService/DeleteFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmGenreServiceClient) CreateGenre(ctx context.Context, in *GenreData, opts ...grpc.CallOption) (*GenreResponse, error) {
	out := new(GenreResponse)
	err := c.cc.Invoke(ctx, "/filmGenre_service.FilmGenreService/CreateGenre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmGenreServiceClient) GetGenreIdsByName(ctx context.Context, in *GenreData, opts ...grpc.CallOption) (*GenreIds, error) {
	out := new(GenreIds)
	err := c.cc.Invoke(ctx, "/filmGenre_service.FilmGenreService/GetGenreIdsByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmGenreServiceClient) GetAllGenres(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (*GetGenresResponse, error) {
	out := new(GetGenresResponse)
	err := c.cc.Invoke(ctx, "/filmGenre_service.FilmGenreService/GetAllGenres", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmGenreServiceClient) DeleteGenre(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/filmGenre_service.FilmGenreService/DeleteGenre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilmGenreServiceServer is the server API for FilmGenreService service.
// All implementations must embed UnimplementedFilmGenreServiceServer
// for forward compatibility
type FilmGenreServiceServer interface {
	CreateFilm(context.Context, *FilmData) (*FilmResponse, error)
	GetAllFilms(context.Context, *NoParams) (*GetFilmResponse, error)
	GetFilmById(context.Context, *ID) (*FilmResponse, error)
	UpDateFilm(context.Context, *ID) (*OperationStatus, error)
	DeleteFilm(context.Context, *ID) (*OperationStatus, error)
	CreateGenre(context.Context, *GenreData) (*GenreResponse, error)
	GetGenreIdsByName(context.Context, *GenreData) (*GenreIds, error)
	GetAllGenres(context.Context, *NoParams) (*GetGenresResponse, error)
	DeleteGenre(context.Context, *ID) (*OperationStatus, error)
	mustEmbedUnimplementedFilmGenreServiceServer()
}

// UnimplementedFilmGenreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFilmGenreServiceServer struct {
}

func (UnimplementedFilmGenreServiceServer) CreateFilm(context.Context, *FilmData) (*FilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFilm not implemented")
}
func (UnimplementedFilmGenreServiceServer) GetAllFilms(context.Context, *NoParams) (*GetFilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFilms not implemented")
}
func (UnimplementedFilmGenreServiceServer) GetFilmById(context.Context, *ID) (*FilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilmById not implemented")
}
func (UnimplementedFilmGenreServiceServer) UpDateFilm(context.Context, *ID) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpDateFilm not implemented")
}
func (UnimplementedFilmGenreServiceServer) DeleteFilm(context.Context, *ID) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFilm not implemented")
}
func (UnimplementedFilmGenreServiceServer) CreateGenre(context.Context, *GenreData) (*GenreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGenre not implemented")
}
func (UnimplementedFilmGenreServiceServer) GetGenreIdsByName(context.Context, *GenreData) (*GenreIds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenreIdsByName not implemented")
}
func (UnimplementedFilmGenreServiceServer) GetAllGenres(context.Context, *NoParams) (*GetGenresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGenres not implemented")
}
func (UnimplementedFilmGenreServiceServer) DeleteGenre(context.Context, *ID) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGenre not implemented")
}
func (UnimplementedFilmGenreServiceServer) mustEmbedUnimplementedFilmGenreServiceServer() {}

// UnsafeFilmGenreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilmGenreServiceServer will
// result in compilation errors.
type UnsafeFilmGenreServiceServer interface {
	mustEmbedUnimplementedFilmGenreServiceServer()
}

func RegisterFilmGenreServiceServer(s grpc.ServiceRegistrar, srv FilmGenreServiceServer) {
	s.RegisterService(&FilmGenreService_ServiceDesc, srv)
}

func _FilmGenreService_CreateFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmGenreServiceServer).CreateFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filmGenre_service.FilmGenreService/CreateFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmGenreServiceServer).CreateFilm(ctx, req.(*FilmData))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmGenreService_GetAllFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmGenreServiceServer).GetAllFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filmGenre_service.FilmGenreService/GetAllFilms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmGenreServiceServer).GetAllFilms(ctx, req.(*NoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmGenreService_GetFilmById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmGenreServiceServer).GetFilmById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filmGenre_service.FilmGenreService/GetFilmById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmGenreServiceServer).GetFilmById(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmGenreService_UpDateFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmGenreServiceServer).UpDateFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filmGenre_service.FilmGenreService/UpDateFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmGenreServiceServer).UpDateFilm(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmGenreService_DeleteFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmGenreServiceServer).DeleteFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filmGenre_service.FilmGenreService/DeleteFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmGenreServiceServer).DeleteFilm(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmGenreService_CreateGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenreData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmGenreServiceServer).CreateGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filmGenre_service.FilmGenreService/CreateGenre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmGenreServiceServer).CreateGenre(ctx, req.(*GenreData))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmGenreService_GetGenreIdsByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenreData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmGenreServiceServer).GetGenreIdsByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filmGenre_service.FilmGenreService/GetGenreIdsByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmGenreServiceServer).GetGenreIdsByName(ctx, req.(*GenreData))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmGenreService_GetAllGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmGenreServiceServer).GetAllGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filmGenre_service.FilmGenreService/GetAllGenres",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmGenreServiceServer).GetAllGenres(ctx, req.(*NoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmGenreService_DeleteGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmGenreServiceServer).DeleteGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filmGenre_service.FilmGenreService/DeleteGenre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmGenreServiceServer).DeleteGenre(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// FilmGenreService_ServiceDesc is the grpc.ServiceDesc for FilmGenreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilmGenreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filmGenre_service.FilmGenreService",
	HandlerType: (*FilmGenreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFilm",
			Handler:    _FilmGenreService_CreateFilm_Handler,
		},
		{
			MethodName: "GetAllFilms",
			Handler:    _FilmGenreService_GetAllFilms_Handler,
		},
		{
			MethodName: "GetFilmById",
			Handler:    _FilmGenreService_GetFilmById_Handler,
		},
		{
			MethodName: "UpDateFilm",
			Handler:    _FilmGenreService_UpDateFilm_Handler,
		},
		{
			MethodName: "DeleteFilm",
			Handler:    _FilmGenreService_DeleteFilm_Handler,
		},
		{
			MethodName: "CreateGenre",
			Handler:    _FilmGenreService_CreateGenre_Handler,
		},
		{
			MethodName: "GetGenreIdsByName",
			Handler:    _FilmGenreService_GetGenreIdsByName_Handler,
		},
		{
			MethodName: "GetAllGenres",
			Handler:    _FilmGenreService_GetAllGenres_Handler,
		},
		{
			MethodName: "DeleteGenre",
			Handler:    _FilmGenreService_DeleteGenre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/filmGenre.proto",
}
