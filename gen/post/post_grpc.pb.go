// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/post/post.proto

package post

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
	Post_GetAllPosts_FullMethodName         = "/Post/GetAllPosts"
	Post_GetPostById_FullMethodName         = "/Post/GetPostById"
	Post_GetPostByUser_FullMethodName       = "/Post/GetPostByUser"
	Post_FindPost_FullMethodName            = "/Post/FindPost"
	Post_AddPost_FullMethodName             = "/Post/AddPost"
	Post_ChangePostRating_FullMethodName    = "/Post/ChangePostRating"
	Post_DeletePost_FullMethodName          = "/Post/DeletePost"
	Post_GetCommentsByPostId_FullMethodName = "/Post/GetCommentsByPostId"
	Post_AddComment_FullMethodName          = "/Post/AddComment"
	Post_ChangeCommentRating_FullMethodName = "/Post/ChangeCommentRating"
	Post_PinComment_FullMethodName          = "/Post/PinComment"
	Post_DeleteComment_FullMethodName       = "/Post/DeleteComment"
)

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostClient interface {
	GetAllPosts(ctx context.Context, in *GetAllPostsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetAllPostsResponse], error)
	GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error)
	GetPostByUser(ctx context.Context, in *GetPostByUserRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetPostByUserResponse], error)
	FindPost(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FindPostRequest, FindPostResponse], error)
	AddPost(ctx context.Context, in *AddPostRequest, opts ...grpc.CallOption) (*AddPostResponse, error)
	ChangePostRating(ctx context.Context, in *ChangePostRatingRequest, opts ...grpc.CallOption) (*ChangePostRatingResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	GetCommentsByPostId(ctx context.Context, in *GetCommentsByPostIdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetCommentsByPostIdResponse], error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error)
	ChangeCommentRating(ctx context.Context, in *ChangeCommentRatingRequest, opts ...grpc.CallOption) (*ChangeCommentRatingResponse, error)
	PinComment(ctx context.Context, in *PinCommentRequest, opts ...grpc.CallOption) (*PinCommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
}

type postClient struct {
	cc grpc.ClientConnInterface
}

func NewPostClient(cc grpc.ClientConnInterface) PostClient {
	return &postClient{cc}
}

func (c *postClient) GetAllPosts(ctx context.Context, in *GetAllPostsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetAllPostsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Post_ServiceDesc.Streams[0], Post_GetAllPosts_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetAllPostsRequest, GetAllPostsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Post_GetAllPostsClient = grpc.ServerStreamingClient[GetAllPostsResponse]

func (c *postClient) GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostByIdResponse)
	err := c.cc.Invoke(ctx, Post_GetPostById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) GetPostByUser(ctx context.Context, in *GetPostByUserRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetPostByUserResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Post_ServiceDesc.Streams[1], Post_GetPostByUser_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetPostByUserRequest, GetPostByUserResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Post_GetPostByUserClient = grpc.ServerStreamingClient[GetPostByUserResponse]

func (c *postClient) FindPost(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FindPostRequest, FindPostResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Post_ServiceDesc.Streams[2], Post_FindPost_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FindPostRequest, FindPostResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Post_FindPostClient = grpc.BidiStreamingClient[FindPostRequest, FindPostResponse]

func (c *postClient) AddPost(ctx context.Context, in *AddPostRequest, opts ...grpc.CallOption) (*AddPostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPostResponse)
	err := c.cc.Invoke(ctx, Post_AddPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) ChangePostRating(ctx context.Context, in *ChangePostRatingRequest, opts ...grpc.CallOption) (*ChangePostRatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePostRatingResponse)
	err := c.cc.Invoke(ctx, Post_ChangePostRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, Post_DeletePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) GetCommentsByPostId(ctx context.Context, in *GetCommentsByPostIdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetCommentsByPostIdResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Post_ServiceDesc.Streams[3], Post_GetCommentsByPostId_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetCommentsByPostIdRequest, GetCommentsByPostIdResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Post_GetCommentsByPostIdClient = grpc.ServerStreamingClient[GetCommentsByPostIdResponse]

func (c *postClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCommentResponse)
	err := c.cc.Invoke(ctx, Post_AddComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) ChangeCommentRating(ctx context.Context, in *ChangeCommentRatingRequest, opts ...grpc.CallOption) (*ChangeCommentRatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeCommentRatingResponse)
	err := c.cc.Invoke(ctx, Post_ChangeCommentRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) PinComment(ctx context.Context, in *PinCommentRequest, opts ...grpc.CallOption) (*PinCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PinCommentResponse)
	err := c.cc.Invoke(ctx, Post_PinComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, Post_DeleteComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServer is the server API for Post service.
// All implementations must embed UnimplementedPostServer
// for forward compatibility.
type PostServer interface {
	GetAllPosts(*GetAllPostsRequest, grpc.ServerStreamingServer[GetAllPostsResponse]) error
	GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error)
	GetPostByUser(*GetPostByUserRequest, grpc.ServerStreamingServer[GetPostByUserResponse]) error
	FindPost(grpc.BidiStreamingServer[FindPostRequest, FindPostResponse]) error
	AddPost(context.Context, *AddPostRequest) (*AddPostResponse, error)
	ChangePostRating(context.Context, *ChangePostRatingRequest) (*ChangePostRatingResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	GetCommentsByPostId(*GetCommentsByPostIdRequest, grpc.ServerStreamingServer[GetCommentsByPostIdResponse]) error
	AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error)
	ChangeCommentRating(context.Context, *ChangeCommentRatingRequest) (*ChangeCommentRatingResponse, error)
	PinComment(context.Context, *PinCommentRequest) (*PinCommentResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	mustEmbedUnimplementedPostServer()
}

// UnimplementedPostServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostServer struct{}

func (UnimplementedPostServer) GetAllPosts(*GetAllPostsRequest, grpc.ServerStreamingServer[GetAllPostsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetAllPosts not implemented")
}
func (UnimplementedPostServer) GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (UnimplementedPostServer) GetPostByUser(*GetPostByUserRequest, grpc.ServerStreamingServer[GetPostByUserResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetPostByUser not implemented")
}
func (UnimplementedPostServer) FindPost(grpc.BidiStreamingServer[FindPostRequest, FindPostResponse]) error {
	return status.Errorf(codes.Unimplemented, "method FindPost not implemented")
}
func (UnimplementedPostServer) AddPost(context.Context, *AddPostRequest) (*AddPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPost not implemented")
}
func (UnimplementedPostServer) ChangePostRating(context.Context, *ChangePostRatingRequest) (*ChangePostRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePostRating not implemented")
}
func (UnimplementedPostServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServer) GetCommentsByPostId(*GetCommentsByPostIdRequest, grpc.ServerStreamingServer[GetCommentsByPostIdResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetCommentsByPostId not implemented")
}
func (UnimplementedPostServer) AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedPostServer) ChangeCommentRating(context.Context, *ChangeCommentRatingRequest) (*ChangeCommentRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCommentRating not implemented")
}
func (UnimplementedPostServer) PinComment(context.Context, *PinCommentRequest) (*PinCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PinComment not implemented")
}
func (UnimplementedPostServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedPostServer) mustEmbedUnimplementedPostServer() {}
func (UnimplementedPostServer) testEmbeddedByValue()              {}

// UnsafePostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServer will
// result in compilation errors.
type UnsafePostServer interface {
	mustEmbedUnimplementedPostServer()
}

func RegisterPostServer(s grpc.ServiceRegistrar, srv PostServer) {
	// If the following call pancis, it indicates UnimplementedPostServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Post_ServiceDesc, srv)
}

func _Post_GetAllPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllPostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostServer).GetAllPosts(m, &grpc.GenericServerStream[GetAllPostsRequest, GetAllPostsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Post_GetAllPostsServer = grpc.ServerStreamingServer[GetAllPostsResponse]

func _Post_GetPostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).GetPostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_GetPostById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).GetPostById(ctx, req.(*GetPostByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_GetPostByUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPostByUserRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostServer).GetPostByUser(m, &grpc.GenericServerStream[GetPostByUserRequest, GetPostByUserResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Post_GetPostByUserServer = grpc.ServerStreamingServer[GetPostByUserResponse]

func _Post_FindPost_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PostServer).FindPost(&grpc.GenericServerStream[FindPostRequest, FindPostResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Post_FindPostServer = grpc.BidiStreamingServer[FindPostRequest, FindPostResponse]

func _Post_AddPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).AddPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_AddPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).AddPost(ctx, req.(*AddPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_ChangePostRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePostRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).ChangePostRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_ChangePostRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).ChangePostRating(ctx, req.(*ChangePostRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_GetCommentsByPostId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCommentsByPostIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostServer).GetCommentsByPostId(m, &grpc.GenericServerStream[GetCommentsByPostIdRequest, GetCommentsByPostIdResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Post_GetCommentsByPostIdServer = grpc.ServerStreamingServer[GetCommentsByPostIdResponse]

func _Post_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_ChangeCommentRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCommentRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).ChangeCommentRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_ChangeCommentRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).ChangeCommentRating(ctx, req.(*ChangeCommentRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_PinComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PinCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).PinComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_PinComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).PinComment(ctx, req.(*PinCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Post_ServiceDesc is the grpc.ServiceDesc for Post service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Post_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPostById",
			Handler:    _Post_GetPostById_Handler,
		},
		{
			MethodName: "AddPost",
			Handler:    _Post_AddPost_Handler,
		},
		{
			MethodName: "ChangePostRating",
			Handler:    _Post_ChangePostRating_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Post_DeletePost_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _Post_AddComment_Handler,
		},
		{
			MethodName: "ChangeCommentRating",
			Handler:    _Post_ChangeCommentRating_Handler,
		},
		{
			MethodName: "PinComment",
			Handler:    _Post_PinComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Post_DeleteComment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllPosts",
			Handler:       _Post_GetAllPosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPostByUser",
			Handler:       _Post_GetPostByUser_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindPost",
			Handler:       _Post_FindPost_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetCommentsByPostId",
			Handler:       _Post_GetCommentsByPostId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/post/post.proto",
}
