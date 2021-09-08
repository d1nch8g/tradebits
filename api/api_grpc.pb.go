// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// SyncTreeClient is the client API for SyncTree service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncTreeClient interface {
	InfoUser(ctx context.Context, in *InfoUserRequest, opts ...grpc.CallOption) (*InfoUserResponse, error)
	InfoMarket(ctx context.Context, in *InfoMarketRequest, opts ...grpc.CallOption) (*InfoMarketResponse, error)
	InfoSearch(ctx context.Context, in *InfoSearchRequest, opts ...grpc.CallOption) (*InfoSearchResponse, error)
	InfoHasTrades(ctx context.Context, in *InfoHasTradesRequest, opts ...grpc.CallOption) (*Response, error)
	InfoMessages(ctx context.Context, in *InfoMessagesRequest, opts ...grpc.CallOption) (*Messages, error)
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Response, error)
	UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*Response, error)
	UserSend(ctx context.Context, in *UserSendRequest, opts ...grpc.CallOption) (*Response, error)
	UserDeposit(ctx context.Context, in *UserDepositRequest, opts ...grpc.CallOption) (*Response, error)
	UserWithdrawal(ctx context.Context, in *UserWithdrawalRequest, opts ...grpc.CallOption) (*Response, error)
	UserSendMessage(ctx context.Context, in *UserSendMessageRequest, opts ...grpc.CallOption) (*Response, error)
	UserBuy(ctx context.Context, in *UserBuyRequest, opts ...grpc.CallOption) (*Response, error)
	UserSell(ctx context.Context, in *UserSellRequest, opts ...grpc.CallOption) (*Response, error)
	UserCancelTrade(ctx context.Context, in *UserCancelTradeRequest, opts ...grpc.CallOption) (*Response, error)
	MarketCraete(ctx context.Context, in *MarketCreateRequest, opts ...grpc.CallOption) (*Response, error)
	MarketUpdate(ctx context.Context, in *MarketUpdateRequest, opts ...grpc.CallOption) (*Response, error)
	MarketDeposit(ctx context.Context, in *MarketDepositRequest, opts ...grpc.CallOption) (*Response, error)
	MarketWithdrawal(ctx context.Context, in *MarketWithdrawalRequest, opts ...grpc.CallOption) (*Response, error)
	MarketSendMessage(ctx context.Context, in *MarketSendMessageRequest, opts ...grpc.CallOption) (*Response, error)
}

type syncTreeClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncTreeClient(cc grpc.ClientConnInterface) SyncTreeClient {
	return &syncTreeClient{cc}
}

func (c *syncTreeClient) InfoUser(ctx context.Context, in *InfoUserRequest, opts ...grpc.CallOption) (*InfoUserResponse, error) {
	out := new(InfoUserResponse)
	err := c.cc.Invoke(ctx, "/api.SyncTree/InfoUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) InfoMarket(ctx context.Context, in *InfoMarketRequest, opts ...grpc.CallOption) (*InfoMarketResponse, error) {
	out := new(InfoMarketResponse)
	err := c.cc.Invoke(ctx, "/api.SyncTree/InfoMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) InfoSearch(ctx context.Context, in *InfoSearchRequest, opts ...grpc.CallOption) (*InfoSearchResponse, error) {
	out := new(InfoSearchResponse)
	err := c.cc.Invoke(ctx, "/api.SyncTree/InfoSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) InfoHasTrades(ctx context.Context, in *InfoHasTradesRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/InfoHasTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) InfoMessages(ctx context.Context, in *InfoMessagesRequest, opts ...grpc.CallOption) (*Messages, error) {
	out := new(Messages)
	err := c.cc.Invoke(ctx, "/api.SyncTree/InfoMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserSend(ctx context.Context, in *UserSendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserDeposit(ctx context.Context, in *UserDepositRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserWithdrawal(ctx context.Context, in *UserWithdrawalRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserSendMessage(ctx context.Context, in *UserSendMessageRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserSendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserBuy(ctx context.Context, in *UserBuyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserBuy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserSell(ctx context.Context, in *UserSellRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) UserCancelTrade(ctx context.Context, in *UserCancelTradeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/UserCancelTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketCraete(ctx context.Context, in *MarketCreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketCraete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketUpdate(ctx context.Context, in *MarketUpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketDeposit(ctx context.Context, in *MarketDepositRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketWithdrawal(ctx context.Context, in *MarketWithdrawalRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncTreeClient) MarketSendMessage(ctx context.Context, in *MarketSendMessageRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SyncTree/MarketSendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncTreeServer is the server API for SyncTree service.
// All implementations must embed UnimplementedSyncTreeServer
// for forward compatibility
type SyncTreeServer interface {
	InfoUser(context.Context, *InfoUserRequest) (*InfoUserResponse, error)
	InfoMarket(context.Context, *InfoMarketRequest) (*InfoMarketResponse, error)
	InfoSearch(context.Context, *InfoSearchRequest) (*InfoSearchResponse, error)
	InfoHasTrades(context.Context, *InfoHasTradesRequest) (*Response, error)
	InfoMessages(context.Context, *InfoMessagesRequest) (*Messages, error)
	UserCreate(context.Context, *UserCreateRequest) (*Response, error)
	UserUpdate(context.Context, *UserUpdateRequest) (*Response, error)
	UserSend(context.Context, *UserSendRequest) (*Response, error)
	UserDeposit(context.Context, *UserDepositRequest) (*Response, error)
	UserWithdrawal(context.Context, *UserWithdrawalRequest) (*Response, error)
	UserSendMessage(context.Context, *UserSendMessageRequest) (*Response, error)
	UserBuy(context.Context, *UserBuyRequest) (*Response, error)
	UserSell(context.Context, *UserSellRequest) (*Response, error)
	UserCancelTrade(context.Context, *UserCancelTradeRequest) (*Response, error)
	MarketCraete(context.Context, *MarketCreateRequest) (*Response, error)
	MarketUpdate(context.Context, *MarketUpdateRequest) (*Response, error)
	MarketDeposit(context.Context, *MarketDepositRequest) (*Response, error)
	MarketWithdrawal(context.Context, *MarketWithdrawalRequest) (*Response, error)
	MarketSendMessage(context.Context, *MarketSendMessageRequest) (*Response, error)
	mustEmbedUnimplementedSyncTreeServer()
}

// UnimplementedSyncTreeServer must be embedded to have forward compatible implementations.
type UnimplementedSyncTreeServer struct {
}

func (UnimplementedSyncTreeServer) InfoUser(context.Context, *InfoUserRequest) (*InfoUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoUser not implemented")
}
func (UnimplementedSyncTreeServer) InfoMarket(context.Context, *InfoMarketRequest) (*InfoMarketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoMarket not implemented")
}
func (UnimplementedSyncTreeServer) InfoSearch(context.Context, *InfoSearchRequest) (*InfoSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoSearch not implemented")
}
func (UnimplementedSyncTreeServer) InfoHasTrades(context.Context, *InfoHasTradesRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoHasTrades not implemented")
}
func (UnimplementedSyncTreeServer) InfoMessages(context.Context, *InfoMessagesRequest) (*Messages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoMessages not implemented")
}
func (UnimplementedSyncTreeServer) UserCreate(context.Context, *UserCreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedSyncTreeServer) UserUpdate(context.Context, *UserUpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (UnimplementedSyncTreeServer) UserSend(context.Context, *UserSendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSend not implemented")
}
func (UnimplementedSyncTreeServer) UserDeposit(context.Context, *UserDepositRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDeposit not implemented")
}
func (UnimplementedSyncTreeServer) UserWithdrawal(context.Context, *UserWithdrawalRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserWithdrawal not implemented")
}
func (UnimplementedSyncTreeServer) UserSendMessage(context.Context, *UserSendMessageRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSendMessage not implemented")
}
func (UnimplementedSyncTreeServer) UserBuy(context.Context, *UserBuyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserBuy not implemented")
}
func (UnimplementedSyncTreeServer) UserSell(context.Context, *UserSellRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSell not implemented")
}
func (UnimplementedSyncTreeServer) UserCancelTrade(context.Context, *UserCancelTradeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCancelTrade not implemented")
}
func (UnimplementedSyncTreeServer) MarketCraete(context.Context, *MarketCreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketCraete not implemented")
}
func (UnimplementedSyncTreeServer) MarketUpdate(context.Context, *MarketUpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketUpdate not implemented")
}
func (UnimplementedSyncTreeServer) MarketDeposit(context.Context, *MarketDepositRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketDeposit not implemented")
}
func (UnimplementedSyncTreeServer) MarketWithdrawal(context.Context, *MarketWithdrawalRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketWithdrawal not implemented")
}
func (UnimplementedSyncTreeServer) MarketSendMessage(context.Context, *MarketSendMessageRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketSendMessage not implemented")
}
func (UnimplementedSyncTreeServer) mustEmbedUnimplementedSyncTreeServer() {}

// UnsafeSyncTreeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncTreeServer will
// result in compilation errors.
type UnsafeSyncTreeServer interface {
	mustEmbedUnimplementedSyncTreeServer()
}

func RegisterSyncTreeServer(s grpc.ServiceRegistrar, srv SyncTreeServer) {
	s.RegisterService(&SyncTree_ServiceDesc, srv)
}

func _SyncTree_InfoUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).InfoUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/InfoUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).InfoUser(ctx, req.(*InfoUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_InfoMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).InfoMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/InfoMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).InfoMarket(ctx, req.(*InfoMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_InfoSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).InfoSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/InfoSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).InfoSearch(ctx, req.(*InfoSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_InfoHasTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoHasTradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).InfoHasTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/InfoHasTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).InfoHasTrades(ctx, req.(*InfoHasTradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_InfoMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).InfoMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/InfoMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).InfoMessages(ctx, req.(*InfoMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserUpdate(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserSend(ctx, req.(*UserSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserDeposit(ctx, req.(*UserDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserWithdrawal(ctx, req.(*UserWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserSendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserSendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserSendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserSendMessage(ctx, req.(*UserSendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserBuy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBuyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserBuy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserBuy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserBuy(ctx, req.(*UserBuyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserSell(ctx, req.(*UserSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_UserCancelTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCancelTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).UserCancelTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/UserCancelTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).UserCancelTrade(ctx, req.(*UserCancelTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketCraete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketCraete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketCraete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketCraete(ctx, req.(*MarketCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketUpdate(ctx, req.(*MarketUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketDeposit(ctx, req.(*MarketDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketWithdrawal(ctx, req.(*MarketWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncTree_MarketSendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketSendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncTreeServer).MarketSendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncTree/MarketSendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncTreeServer).MarketSendMessage(ctx, req.(*MarketSendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncTree_ServiceDesc is the grpc.ServiceDesc for SyncTree service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncTree_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SyncTree",
	HandlerType: (*SyncTreeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InfoUser",
			Handler:    _SyncTree_InfoUser_Handler,
		},
		{
			MethodName: "InfoMarket",
			Handler:    _SyncTree_InfoMarket_Handler,
		},
		{
			MethodName: "InfoSearch",
			Handler:    _SyncTree_InfoSearch_Handler,
		},
		{
			MethodName: "InfoHasTrades",
			Handler:    _SyncTree_InfoHasTrades_Handler,
		},
		{
			MethodName: "InfoMessages",
			Handler:    _SyncTree_InfoMessages_Handler,
		},
		{
			MethodName: "UserCreate",
			Handler:    _SyncTree_UserCreate_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _SyncTree_UserUpdate_Handler,
		},
		{
			MethodName: "UserSend",
			Handler:    _SyncTree_UserSend_Handler,
		},
		{
			MethodName: "UserDeposit",
			Handler:    _SyncTree_UserDeposit_Handler,
		},
		{
			MethodName: "UserWithdrawal",
			Handler:    _SyncTree_UserWithdrawal_Handler,
		},
		{
			MethodName: "UserSendMessage",
			Handler:    _SyncTree_UserSendMessage_Handler,
		},
		{
			MethodName: "UserBuy",
			Handler:    _SyncTree_UserBuy_Handler,
		},
		{
			MethodName: "UserSell",
			Handler:    _SyncTree_UserSell_Handler,
		},
		{
			MethodName: "UserCancelTrade",
			Handler:    _SyncTree_UserCancelTrade_Handler,
		},
		{
			MethodName: "MarketCraete",
			Handler:    _SyncTree_MarketCraete_Handler,
		},
		{
			MethodName: "MarketUpdate",
			Handler:    _SyncTree_MarketUpdate_Handler,
		},
		{
			MethodName: "MarketDeposit",
			Handler:    _SyncTree_MarketDeposit_Handler,
		},
		{
			MethodName: "MarketWithdrawal",
			Handler:    _SyncTree_MarketWithdrawal_Handler,
		},
		{
			MethodName: "MarketSendMessage",
			Handler:    _SyncTree_MarketSendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
