// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0
// source: publish/service/v1/publish.proto

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PublishService_GetPublishList_FullMethodName         = "/publish.service.v1.PublishService/GetPublishList"
	PublishService_PublishAction_FullMethodName          = "/publish.service.v1.PublishService/PublishAction"
	PublishService_GetVideoList_FullMethodName           = "/publish.service.v1.PublishService/GetVideoList"
	PublishService_GetVideoListByVideoIds_FullMethodName = "/publish.service.v1.PublishService/GetVideoListByVideoIds"
	PublishService_UpdateFavorite_FullMethodName         = "/publish.service.v1.PublishService/UpdateFavorite"
	PublishService_UpdateComment_FullMethodName          = "/publish.service.v1.PublishService/UpdateComment"
)

// PublishServiceClient is the client API for PublishService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublishServiceClient interface {
	// 获取用户投稿视频列表
	GetPublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListReply, error)
	// 用户上传视频
	PublishAction(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionReply, error)
	// feed服务请求根据最新投稿时间及需求数量获取视频列表
	GetVideoList(ctx context.Context, in *VideoListRequest, opts ...grpc.CallOption) (*VideoListReply, error)
	// favorite相关服务请求根据视频id列表获取视频列表
	GetVideoListByVideoIds(ctx context.Context, in *VideoListByVideoIdsRequest, opts ...grpc.CallOption) (*VideoListReply, error)
	// relation相关服务请求更新某一视频点赞数量
	UpdateFavorite(ctx context.Context, in *UpdateFavoriteCountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// comment相关服务请求更新某一视频评论数量
	UpdateComment(ctx context.Context, in *UpdateCommentCountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type publishServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublishServiceClient(cc grpc.ClientConnInterface) PublishServiceClient {
	return &publishServiceClient{cc}
}

func (c *publishServiceClient) GetPublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListReply, error) {
	out := new(PublishListReply)
	err := c.cc.Invoke(ctx, PublishService_GetPublishList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishServiceClient) PublishAction(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionReply, error) {
	out := new(PublishActionReply)
	err := c.cc.Invoke(ctx, PublishService_PublishAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishServiceClient) GetVideoList(ctx context.Context, in *VideoListRequest, opts ...grpc.CallOption) (*VideoListReply, error) {
	out := new(VideoListReply)
	err := c.cc.Invoke(ctx, PublishService_GetVideoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishServiceClient) GetVideoListByVideoIds(ctx context.Context, in *VideoListByVideoIdsRequest, opts ...grpc.CallOption) (*VideoListReply, error) {
	out := new(VideoListReply)
	err := c.cc.Invoke(ctx, PublishService_GetVideoListByVideoIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishServiceClient) UpdateFavorite(ctx context.Context, in *UpdateFavoriteCountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PublishService_UpdateFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentCountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PublishService_UpdateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublishServiceServer is the server API for PublishService service.
// All implementations must embed UnimplementedPublishServiceServer
// for forward compatibility
type PublishServiceServer interface {
	// 获取用户投稿视频列表
	GetPublishList(context.Context, *PublishListRequest) (*PublishListReply, error)
	// 用户上传视频
	PublishAction(context.Context, *PublishActionRequest) (*PublishActionReply, error)
	// feed服务请求根据最新投稿时间及需求数量获取视频列表
	GetVideoList(context.Context, *VideoListRequest) (*VideoListReply, error)
	// favorite相关服务请求根据视频id列表获取视频列表
	GetVideoListByVideoIds(context.Context, *VideoListByVideoIdsRequest) (*VideoListReply, error)
	// relation相关服务请求更新某一视频点赞数量
	UpdateFavorite(context.Context, *UpdateFavoriteCountRequest) (*emptypb.Empty, error)
	// comment相关服务请求更新某一视频评论数量
	UpdateComment(context.Context, *UpdateCommentCountRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPublishServiceServer()
}

// UnimplementedPublishServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublishServiceServer struct {
}

func (UnimplementedPublishServiceServer) GetPublishList(context.Context, *PublishListRequest) (*PublishListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishList not implemented")
}
func (UnimplementedPublishServiceServer) PublishAction(context.Context, *PublishActionRequest) (*PublishActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAction not implemented")
}
func (UnimplementedPublishServiceServer) GetVideoList(context.Context, *VideoListRequest) (*VideoListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoList not implemented")
}
func (UnimplementedPublishServiceServer) GetVideoListByVideoIds(context.Context, *VideoListByVideoIdsRequest) (*VideoListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoListByVideoIds not implemented")
}
func (UnimplementedPublishServiceServer) UpdateFavorite(context.Context, *UpdateFavoriteCountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFavorite not implemented")
}
func (UnimplementedPublishServiceServer) UpdateComment(context.Context, *UpdateCommentCountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedPublishServiceServer) mustEmbedUnimplementedPublishServiceServer() {}

// UnsafePublishServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublishServiceServer will
// result in compilation errors.
type UnsafePublishServiceServer interface {
	mustEmbedUnimplementedPublishServiceServer()
}

func RegisterPublishServiceServer(s grpc.ServiceRegistrar, srv PublishServiceServer) {
	s.RegisterService(&PublishService_ServiceDesc, srv)
}

func _PublishService_GetPublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).GetPublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublishService_GetPublishList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).GetPublishList(ctx, req.(*PublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishService_PublishAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).PublishAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublishService_PublishAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).PublishAction(ctx, req.(*PublishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishService_GetVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).GetVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublishService_GetVideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).GetVideoList(ctx, req.(*VideoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishService_GetVideoListByVideoIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoListByVideoIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).GetVideoListByVideoIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublishService_GetVideoListByVideoIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).GetVideoListByVideoIds(ctx, req.(*VideoListByVideoIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishService_UpdateFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFavoriteCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).UpdateFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublishService_UpdateFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).UpdateFavorite(ctx, req.(*UpdateFavoriteCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublishService_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).UpdateComment(ctx, req.(*UpdateCommentCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PublishService_ServiceDesc is the grpc.ServiceDesc for PublishService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublishService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "publish.service.v1.PublishService",
	HandlerType: (*PublishServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublishList",
			Handler:    _PublishService_GetPublishList_Handler,
		},
		{
			MethodName: "PublishAction",
			Handler:    _PublishService_PublishAction_Handler,
		},
		{
			MethodName: "GetVideoList",
			Handler:    _PublishService_GetVideoList_Handler,
		},
		{
			MethodName: "GetVideoListByVideoIds",
			Handler:    _PublishService_GetVideoListByVideoIds_Handler,
		},
		{
			MethodName: "UpdateFavorite",
			Handler:    _PublishService_UpdateFavorite_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _PublishService_UpdateComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publish/service/v1/publish.proto",
}
