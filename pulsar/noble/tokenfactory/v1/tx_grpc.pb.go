// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: noble/tokenfactory/v1/tx.proto

package tokenfactoryv1

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

const (
	Msg_UpdateMasterMinter_FullMethodName        = "/noble.tokenfactory.v1.Msg/UpdateMasterMinter"
	Msg_UpdatePauser_FullMethodName              = "/noble.tokenfactory.v1.Msg/UpdatePauser"
	Msg_UpdateBlacklister_FullMethodName         = "/noble.tokenfactory.v1.Msg/UpdateBlacklister"
	Msg_UpdateOwner_FullMethodName               = "/noble.tokenfactory.v1.Msg/UpdateOwner"
	Msg_AcceptOwner_FullMethodName               = "/noble.tokenfactory.v1.Msg/AcceptOwner"
	Msg_ConfigureMinter_FullMethodName           = "/noble.tokenfactory.v1.Msg/ConfigureMinter"
	Msg_RemoveMinter_FullMethodName              = "/noble.tokenfactory.v1.Msg/RemoveMinter"
	Msg_Mint_FullMethodName                      = "/noble.tokenfactory.v1.Msg/Mint"
	Msg_Burn_FullMethodName                      = "/noble.tokenfactory.v1.Msg/Burn"
	Msg_Blacklist_FullMethodName                 = "/noble.tokenfactory.v1.Msg/Blacklist"
	Msg_Unblacklist_FullMethodName               = "/noble.tokenfactory.v1.Msg/Unblacklist"
	Msg_Pause_FullMethodName                     = "/noble.tokenfactory.v1.Msg/Pause"
	Msg_Unpause_FullMethodName                   = "/noble.tokenfactory.v1.Msg/Unpause"
	Msg_ConfigureMinterController_FullMethodName = "/noble.tokenfactory.v1.Msg/ConfigureMinterController"
	Msg_RemoveMinterController_FullMethodName    = "/noble.tokenfactory.v1.Msg/RemoveMinterController"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	UpdateMasterMinter(ctx context.Context, in *MsgUpdateMasterMinter, opts ...grpc.CallOption) (*MsgUpdateMasterMinterResponse, error)
	UpdatePauser(ctx context.Context, in *MsgUpdatePauser, opts ...grpc.CallOption) (*MsgUpdatePauserResponse, error)
	UpdateBlacklister(ctx context.Context, in *MsgUpdateBlacklister, opts ...grpc.CallOption) (*MsgUpdateBlacklisterResponse, error)
	UpdateOwner(ctx context.Context, in *MsgUpdateOwner, opts ...grpc.CallOption) (*MsgUpdateOwnerResponse, error)
	AcceptOwner(ctx context.Context, in *MsgAcceptOwner, opts ...grpc.CallOption) (*MsgAcceptOwnerResponse, error)
	ConfigureMinter(ctx context.Context, in *MsgConfigureMinter, opts ...grpc.CallOption) (*MsgConfigureMinterResponse, error)
	RemoveMinter(ctx context.Context, in *MsgRemoveMinter, opts ...grpc.CallOption) (*MsgRemoveMinterResponse, error)
	Mint(ctx context.Context, in *MsgMint, opts ...grpc.CallOption) (*MsgMintResponse, error)
	Burn(ctx context.Context, in *MsgBurn, opts ...grpc.CallOption) (*MsgBurnResponse, error)
	Blacklist(ctx context.Context, in *MsgBlacklist, opts ...grpc.CallOption) (*MsgBlacklistResponse, error)
	Unblacklist(ctx context.Context, in *MsgUnblacklist, opts ...grpc.CallOption) (*MsgUnblacklistResponse, error)
	Pause(ctx context.Context, in *MsgPause, opts ...grpc.CallOption) (*MsgPauseResponse, error)
	Unpause(ctx context.Context, in *MsgUnpause, opts ...grpc.CallOption) (*MsgUnpauseResponse, error)
	ConfigureMinterController(ctx context.Context, in *MsgConfigureMinterController, opts ...grpc.CallOption) (*MsgConfigureMinterControllerResponse, error)
	RemoveMinterController(ctx context.Context, in *MsgRemoveMinterController, opts ...grpc.CallOption) (*MsgRemoveMinterControllerResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateMasterMinter(ctx context.Context, in *MsgUpdateMasterMinter, opts ...grpc.CallOption) (*MsgUpdateMasterMinterResponse, error) {
	out := new(MsgUpdateMasterMinterResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateMasterMinter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePauser(ctx context.Context, in *MsgUpdatePauser, opts ...grpc.CallOption) (*MsgUpdatePauserResponse, error) {
	out := new(MsgUpdatePauserResponse)
	err := c.cc.Invoke(ctx, Msg_UpdatePauser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateBlacklister(ctx context.Context, in *MsgUpdateBlacklister, opts ...grpc.CallOption) (*MsgUpdateBlacklisterResponse, error) {
	out := new(MsgUpdateBlacklisterResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateBlacklister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateOwner(ctx context.Context, in *MsgUpdateOwner, opts ...grpc.CallOption) (*MsgUpdateOwnerResponse, error) {
	out := new(MsgUpdateOwnerResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AcceptOwner(ctx context.Context, in *MsgAcceptOwner, opts ...grpc.CallOption) (*MsgAcceptOwnerResponse, error) {
	out := new(MsgAcceptOwnerResponse)
	err := c.cc.Invoke(ctx, Msg_AcceptOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConfigureMinter(ctx context.Context, in *MsgConfigureMinter, opts ...grpc.CallOption) (*MsgConfigureMinterResponse, error) {
	out := new(MsgConfigureMinterResponse)
	err := c.cc.Invoke(ctx, Msg_ConfigureMinter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveMinter(ctx context.Context, in *MsgRemoveMinter, opts ...grpc.CallOption) (*MsgRemoveMinterResponse, error) {
	out := new(MsgRemoveMinterResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveMinter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Mint(ctx context.Context, in *MsgMint, opts ...grpc.CallOption) (*MsgMintResponse, error) {
	out := new(MsgMintResponse)
	err := c.cc.Invoke(ctx, Msg_Mint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Burn(ctx context.Context, in *MsgBurn, opts ...grpc.CallOption) (*MsgBurnResponse, error) {
	out := new(MsgBurnResponse)
	err := c.cc.Invoke(ctx, Msg_Burn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Blacklist(ctx context.Context, in *MsgBlacklist, opts ...grpc.CallOption) (*MsgBlacklistResponse, error) {
	out := new(MsgBlacklistResponse)
	err := c.cc.Invoke(ctx, Msg_Blacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Unblacklist(ctx context.Context, in *MsgUnblacklist, opts ...grpc.CallOption) (*MsgUnblacklistResponse, error) {
	out := new(MsgUnblacklistResponse)
	err := c.cc.Invoke(ctx, Msg_Unblacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Pause(ctx context.Context, in *MsgPause, opts ...grpc.CallOption) (*MsgPauseResponse, error) {
	out := new(MsgPauseResponse)
	err := c.cc.Invoke(ctx, Msg_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Unpause(ctx context.Context, in *MsgUnpause, opts ...grpc.CallOption) (*MsgUnpauseResponse, error) {
	out := new(MsgUnpauseResponse)
	err := c.cc.Invoke(ctx, Msg_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConfigureMinterController(ctx context.Context, in *MsgConfigureMinterController, opts ...grpc.CallOption) (*MsgConfigureMinterControllerResponse, error) {
	out := new(MsgConfigureMinterControllerResponse)
	err := c.cc.Invoke(ctx, Msg_ConfigureMinterController_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveMinterController(ctx context.Context, in *MsgRemoveMinterController, opts ...grpc.CallOption) (*MsgRemoveMinterControllerResponse, error) {
	out := new(MsgRemoveMinterControllerResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveMinterController_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	UpdateMasterMinter(context.Context, *MsgUpdateMasterMinter) (*MsgUpdateMasterMinterResponse, error)
	UpdatePauser(context.Context, *MsgUpdatePauser) (*MsgUpdatePauserResponse, error)
	UpdateBlacklister(context.Context, *MsgUpdateBlacklister) (*MsgUpdateBlacklisterResponse, error)
	UpdateOwner(context.Context, *MsgUpdateOwner) (*MsgUpdateOwnerResponse, error)
	AcceptOwner(context.Context, *MsgAcceptOwner) (*MsgAcceptOwnerResponse, error)
	ConfigureMinter(context.Context, *MsgConfigureMinter) (*MsgConfigureMinterResponse, error)
	RemoveMinter(context.Context, *MsgRemoveMinter) (*MsgRemoveMinterResponse, error)
	Mint(context.Context, *MsgMint) (*MsgMintResponse, error)
	Burn(context.Context, *MsgBurn) (*MsgBurnResponse, error)
	Blacklist(context.Context, *MsgBlacklist) (*MsgBlacklistResponse, error)
	Unblacklist(context.Context, *MsgUnblacklist) (*MsgUnblacklistResponse, error)
	Pause(context.Context, *MsgPause) (*MsgPauseResponse, error)
	Unpause(context.Context, *MsgUnpause) (*MsgUnpauseResponse, error)
	ConfigureMinterController(context.Context, *MsgConfigureMinterController) (*MsgConfigureMinterControllerResponse, error)
	RemoveMinterController(context.Context, *MsgRemoveMinterController) (*MsgRemoveMinterControllerResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateMasterMinter(context.Context, *MsgUpdateMasterMinter) (*MsgUpdateMasterMinterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMasterMinter not implemented")
}
func (UnimplementedMsgServer) UpdatePauser(context.Context, *MsgUpdatePauser) (*MsgUpdatePauserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePauser not implemented")
}
func (UnimplementedMsgServer) UpdateBlacklister(context.Context, *MsgUpdateBlacklister) (*MsgUpdateBlacklisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlacklister not implemented")
}
func (UnimplementedMsgServer) UpdateOwner(context.Context, *MsgUpdateOwner) (*MsgUpdateOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOwner not implemented")
}
func (UnimplementedMsgServer) AcceptOwner(context.Context, *MsgAcceptOwner) (*MsgAcceptOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptOwner not implemented")
}
func (UnimplementedMsgServer) ConfigureMinter(context.Context, *MsgConfigureMinter) (*MsgConfigureMinterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureMinter not implemented")
}
func (UnimplementedMsgServer) RemoveMinter(context.Context, *MsgRemoveMinter) (*MsgRemoveMinterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMinter not implemented")
}
func (UnimplementedMsgServer) Mint(context.Context, *MsgMint) (*MsgMintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mint not implemented")
}
func (UnimplementedMsgServer) Burn(context.Context, *MsgBurn) (*MsgBurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Burn not implemented")
}
func (UnimplementedMsgServer) Blacklist(context.Context, *MsgBlacklist) (*MsgBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Blacklist not implemented")
}
func (UnimplementedMsgServer) Unblacklist(context.Context, *MsgUnblacklist) (*MsgUnblacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unblacklist not implemented")
}
func (UnimplementedMsgServer) Pause(context.Context, *MsgPause) (*MsgPauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedMsgServer) Unpause(context.Context, *MsgUnpause) (*MsgUnpauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedMsgServer) ConfigureMinterController(context.Context, *MsgConfigureMinterController) (*MsgConfigureMinterControllerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureMinterController not implemented")
}
func (UnimplementedMsgServer) RemoveMinterController(context.Context, *MsgRemoveMinterController) (*MsgRemoveMinterControllerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMinterController not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateMasterMinter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateMasterMinter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateMasterMinter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateMasterMinter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateMasterMinter(ctx, req.(*MsgUpdateMasterMinter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePauser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePauser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePauser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdatePauser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePauser(ctx, req.(*MsgUpdatePauser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateBlacklister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateBlacklister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateBlacklister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateBlacklister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateBlacklister(ctx, req.(*MsgUpdateBlacklister))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateOwner(ctx, req.(*MsgUpdateOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AcceptOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAcceptOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AcceptOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AcceptOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptOwner(ctx, req.(*MsgAcceptOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConfigureMinter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConfigureMinter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConfigureMinter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ConfigureMinter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConfigureMinter(ctx, req.(*MsgConfigureMinter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveMinter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveMinter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveMinter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveMinter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveMinter(ctx, req.(*MsgRemoveMinter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Mint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Mint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Mint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Mint(ctx, req.(*MsgMint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Burn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Burn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Burn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Burn(ctx, req.(*MsgBurn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Blacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBlacklist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Blacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Blacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Blacklist(ctx, req.(*MsgBlacklist))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Unblacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnblacklist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Unblacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Unblacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Unblacklist(ctx, req.(*MsgUnblacklist))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPause)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Pause(ctx, req.(*MsgPause))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnpause)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Unpause(ctx, req.(*MsgUnpause))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConfigureMinterController_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConfigureMinterController)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConfigureMinterController(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ConfigureMinterController_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConfigureMinterController(ctx, req.(*MsgConfigureMinterController))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveMinterController_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveMinterController)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveMinterController(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveMinterController_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveMinterController(ctx, req.(*MsgRemoveMinterController))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "noble.tokenfactory.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMasterMinter",
			Handler:    _Msg_UpdateMasterMinter_Handler,
		},
		{
			MethodName: "UpdatePauser",
			Handler:    _Msg_UpdatePauser_Handler,
		},
		{
			MethodName: "UpdateBlacklister",
			Handler:    _Msg_UpdateBlacklister_Handler,
		},
		{
			MethodName: "UpdateOwner",
			Handler:    _Msg_UpdateOwner_Handler,
		},
		{
			MethodName: "AcceptOwner",
			Handler:    _Msg_AcceptOwner_Handler,
		},
		{
			MethodName: "ConfigureMinter",
			Handler:    _Msg_ConfigureMinter_Handler,
		},
		{
			MethodName: "RemoveMinter",
			Handler:    _Msg_RemoveMinter_Handler,
		},
		{
			MethodName: "Mint",
			Handler:    _Msg_Mint_Handler,
		},
		{
			MethodName: "Burn",
			Handler:    _Msg_Burn_Handler,
		},
		{
			MethodName: "Blacklist",
			Handler:    _Msg_Blacklist_Handler,
		},
		{
			MethodName: "Unblacklist",
			Handler:    _Msg_Unblacklist_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _Msg_Pause_Handler,
		},
		{
			MethodName: "Unpause",
			Handler:    _Msg_Unpause_Handler,
		},
		{
			MethodName: "ConfigureMinterController",
			Handler:    _Msg_ConfigureMinterController_Handler,
		},
		{
			MethodName: "RemoveMinterController",
			Handler:    _Msg_RemoveMinterController_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noble/tokenfactory/v1/tx.proto",
}
