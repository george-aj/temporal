// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/compound/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53c556dce20a4edf, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53c556dce20a4edf, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetCompoundSettingRequest struct {
	Index123 string `protobuf:"bytes,1,opt,name=index123,proto3" json:"index123,omitempty"`
}

func (m *QueryGetCompoundSettingRequest) Reset()         { *m = QueryGetCompoundSettingRequest{} }
func (m *QueryGetCompoundSettingRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCompoundSettingRequest) ProtoMessage()    {}
func (*QueryGetCompoundSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53c556dce20a4edf, []int{2}
}
func (m *QueryGetCompoundSettingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCompoundSettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCompoundSettingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCompoundSettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCompoundSettingRequest.Merge(m, src)
}
func (m *QueryGetCompoundSettingRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCompoundSettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCompoundSettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCompoundSettingRequest proto.InternalMessageInfo

func (m *QueryGetCompoundSettingRequest) GetIndex123() string {
	if m != nil {
		return m.Index123
	}
	return ""
}

type QueryGetCompoundSettingResponse struct {
	CompoundSetting CompoundSetting `protobuf:"bytes,1,opt,name=compoundSetting,proto3" json:"compoundSetting"`
}

func (m *QueryGetCompoundSettingResponse) Reset()         { *m = QueryGetCompoundSettingResponse{} }
func (m *QueryGetCompoundSettingResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCompoundSettingResponse) ProtoMessage()    {}
func (*QueryGetCompoundSettingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53c556dce20a4edf, []int{3}
}
func (m *QueryGetCompoundSettingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCompoundSettingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCompoundSettingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCompoundSettingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCompoundSettingResponse.Merge(m, src)
}
func (m *QueryGetCompoundSettingResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCompoundSettingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCompoundSettingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCompoundSettingResponse proto.InternalMessageInfo

func (m *QueryGetCompoundSettingResponse) GetCompoundSetting() CompoundSetting {
	if m != nil {
		return m.CompoundSetting
	}
	return CompoundSetting{}
}

type QueryAllCompoundSettingRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCompoundSettingRequest) Reset()         { *m = QueryAllCompoundSettingRequest{} }
func (m *QueryAllCompoundSettingRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCompoundSettingRequest) ProtoMessage()    {}
func (*QueryAllCompoundSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53c556dce20a4edf, []int{4}
}
func (m *QueryAllCompoundSettingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCompoundSettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCompoundSettingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCompoundSettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCompoundSettingRequest.Merge(m, src)
}
func (m *QueryAllCompoundSettingRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCompoundSettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCompoundSettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCompoundSettingRequest proto.InternalMessageInfo

func (m *QueryAllCompoundSettingRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCompoundSettingResponse struct {
	CompoundSetting []CompoundSetting   `protobuf:"bytes,1,rep,name=compoundSetting,proto3" json:"compoundSetting"`
	Pagination      *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCompoundSettingResponse) Reset()         { *m = QueryAllCompoundSettingResponse{} }
func (m *QueryAllCompoundSettingResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCompoundSettingResponse) ProtoMessage()    {}
func (*QueryAllCompoundSettingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53c556dce20a4edf, []int{5}
}
func (m *QueryAllCompoundSettingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCompoundSettingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCompoundSettingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCompoundSettingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCompoundSettingResponse.Merge(m, src)
}
func (m *QueryAllCompoundSettingResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCompoundSettingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCompoundSettingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCompoundSettingResponse proto.InternalMessageInfo

func (m *QueryAllCompoundSettingResponse) GetCompoundSetting() []CompoundSetting {
	if m != nil {
		return m.CompoundSetting
	}
	return nil
}

func (m *QueryAllCompoundSettingResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "temporal.compound.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "temporal.compound.QueryParamsResponse")
	proto.RegisterType((*QueryGetCompoundSettingRequest)(nil), "temporal.compound.QueryGetCompoundSettingRequest")
	proto.RegisterType((*QueryGetCompoundSettingResponse)(nil), "temporal.compound.QueryGetCompoundSettingResponse")
	proto.RegisterType((*QueryAllCompoundSettingRequest)(nil), "temporal.compound.QueryAllCompoundSettingRequest")
	proto.RegisterType((*QueryAllCompoundSettingResponse)(nil), "temporal.compound.QueryAllCompoundSettingResponse")
}

func init() { proto.RegisterFile("temporal/compound/query.proto", fileDescriptor_53c556dce20a4edf) }

var fileDescriptor_53c556dce20a4edf = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x0d, 0x2a, 0x30, 0x87, 0x09, 0xb3, 0x03, 0x0d, 0xe0, 0x81, 0xd1, 0xc6, 0x00,
	0xc9, 0xa6, 0xad, 0x04, 0x17, 0x2e, 0x1b, 0x12, 0xbb, 0xa1, 0x11, 0x6e, 0x5c, 0x90, 0xdb, 0x5a,
	0x21, 0x52, 0x1a, 0x7b, 0xb1, 0x33, 0x6d, 0x20, 0x2e, 0x7c, 0x02, 0x24, 0xbe, 0x04, 0x47, 0x3e,
	0x01, 0xe7, 0x1d, 0x27, 0x71, 0xe1, 0x84, 0xa0, 0xe5, 0x83, 0xa0, 0xda, 0xce, 0xda, 0x26, 0x35,
	0x05, 0x71, 0x73, 0xf2, 0xfe, 0xef, 0xff, 0x7e, 0xef, 0xe5, 0x39, 0xe0, 0x86, 0xe6, 0x43, 0x29,
	0x72, 0x96, 0xd2, 0xbe, 0x18, 0x4a, 0x51, 0x64, 0x03, 0x7a, 0x50, 0xf0, 0xfc, 0x98, 0xc8, 0x5c,
	0x68, 0x01, 0x2f, 0x97, 0x61, 0x52, 0x86, 0xc3, 0xf5, 0x58, 0xc4, 0xc2, 0x44, 0xe9, 0xe4, 0x64,
	0x85, 0xe1, 0xf5, 0x58, 0x88, 0x38, 0xe5, 0x94, 0xc9, 0x84, 0xb2, 0x2c, 0x13, 0x9a, 0xe9, 0x44,
	0x64, 0xca, 0x45, 0xef, 0xf5, 0x85, 0x1a, 0x0a, 0x45, 0x7b, 0x4c, 0x71, 0xeb, 0x4f, 0x0f, 0xdb,
	0x3d, 0xae, 0x59, 0x9b, 0x4a, 0x16, 0x27, 0x99, 0x11, 0x3b, 0x2d, 0xaa, 0x13, 0x49, 0x96, 0xb3,
	0x61, 0xe9, 0xb5, 0x5d, 0x8f, 0x97, 0x87, 0x57, 0x8a, 0x6b, 0x9d, 0x64, 0xb1, 0x53, 0xde, 0xad,
	0x2b, 0x0f, 0x59, 0x9a, 0x0c, 0x98, 0x16, 0x79, 0x45, 0x8a, 0x66, 0x01, 0x4b, 0xb4, 0xbe, 0x48,
	0x1c, 0x14, 0x5e, 0x07, 0xf0, 0xf9, 0x04, 0x7b, 0xdf, 0x90, 0x44, 0xfc, 0xa0, 0xe0, 0x4a, 0xe3,
	0x67, 0xe0, 0xca, 0xdc, 0x5b, 0x25, 0x45, 0xa6, 0x38, 0x7c, 0x04, 0x9a, 0x96, 0xf8, 0x6a, 0x70,
	0x33, 0xd8, 0xbe, 0xd4, 0x69, 0x91, 0xda, 0x14, 0x89, 0x4d, 0xd9, 0x3d, 0x77, 0xf2, 0x7d, 0xa3,
	0x11, 0x39, 0x39, 0x7e, 0x0c, 0x90, 0xf1, 0xdb, 0xe3, 0xfa, 0x89, 0x13, 0xbe, 0xb0, 0x98, 0xae,
	0x22, 0x0c, 0xc1, 0x85, 0x24, 0x1b, 0xf0, 0xa3, 0x76, 0xa7, 0x6b, 0xcc, 0x2f, 0x46, 0x67, 0xcf,
	0xb8, 0x00, 0x1b, 0xde, 0x6c, 0x47, 0x16, 0x81, 0xb5, 0xfe, 0x7c, 0xc8, 0x21, 0xe2, 0x05, 0x88,
	0x15, 0x13, 0xc7, 0x5a, 0x35, 0xc0, 0xaf, 0x1d, 0xf4, 0x4e, 0x9a, 0x7a, 0xa0, 0x9f, 0x02, 0x30,
	0xfd, 0xca, 0xae, 0xe0, 0x16, 0xb1, 0x13, 0x27, 0x93, 0x89, 0x13, 0xbb, 0x72, 0x6e, 0xee, 0x64,
	0x9f, 0xc5, 0xdc, 0xe5, 0x46, 0x33, 0x99, 0xf8, 0x4b, 0xe0, 0x3a, 0x5c, 0x54, 0xea, 0x4f, 0x1d,
	0xae, 0xfe, 0x57, 0x87, 0x70, 0x6f, 0x8e, 0x7f, 0xc5, 0xf0, 0xdf, 0x59, 0xca, 0x6f, 0x81, 0x66,
	0x1b, 0xe8, 0xfc, 0x5c, 0x05, 0xe7, 0x4d, 0x03, 0xf0, 0x0d, 0x68, 0xda, 0x0d, 0x80, 0x9b, 0x0b,
	0xb8, 0xea, 0xab, 0x16, 0x6e, 0x2d, 0x93, 0xd9, 0x72, 0xf8, 0xd6, 0xfb, 0xaf, 0xbf, 0x3e, 0xae,
	0x5c, 0x83, 0x2d, 0xea, 0xbb, 0x46, 0xf0, 0x73, 0x00, 0xd6, 0x2a, 0x9d, 0xc3, 0xb6, 0xcf, 0xde,
	0xbb, 0x8a, 0x61, 0xe7, 0x5f, 0x52, 0x1c, 0xdd, 0x43, 0x43, 0xf7, 0x00, 0x12, 0xba, 0xfc, 0x12,
	0xd3, 0xb7, 0xe5, 0x66, 0xbf, 0x83, 0x9f, 0x02, 0x00, 0x2b, 0x9e, 0x3b, 0x69, 0xea, 0xa7, 0xf6,
	0xee, 0xa2, 0x9f, 0xda, 0xbf, 0x53, 0xf8, 0xbe, 0xa1, 0xde, 0x84, 0xb7, 0xff, 0x82, 0x7a, 0xb7,
	0x7b, 0x32, 0x42, 0xc1, 0xe9, 0x08, 0x05, 0x3f, 0x46, 0x28, 0xf8, 0x30, 0x46, 0x8d, 0xd3, 0x31,
	0x6a, 0x7c, 0x1b, 0xa3, 0xc6, 0xcb, 0xd6, 0x59, 0xf6, 0xd1, 0x34, 0x5f, 0x1f, 0x4b, 0xae, 0x7a,
	0x4d, 0xf3, 0x97, 0xe9, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x98, 0x40, 0x50, 0x52, 0x8e, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of CompoundSetting items.
	CompoundSetting(ctx context.Context, in *QueryGetCompoundSettingRequest, opts ...grpc.CallOption) (*QueryGetCompoundSettingResponse, error)
	CompoundSettingAll(ctx context.Context, in *QueryAllCompoundSettingRequest, opts ...grpc.CallOption) (*QueryAllCompoundSettingResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/temporal.compound.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CompoundSetting(ctx context.Context, in *QueryGetCompoundSettingRequest, opts ...grpc.CallOption) (*QueryGetCompoundSettingResponse, error) {
	out := new(QueryGetCompoundSettingResponse)
	err := c.cc.Invoke(ctx, "/temporal.compound.Query/CompoundSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CompoundSettingAll(ctx context.Context, in *QueryAllCompoundSettingRequest, opts ...grpc.CallOption) (*QueryAllCompoundSettingResponse, error) {
	out := new(QueryAllCompoundSettingResponse)
	err := c.cc.Invoke(ctx, "/temporal.compound.Query/CompoundSettingAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of CompoundSetting items.
	CompoundSetting(context.Context, *QueryGetCompoundSettingRequest) (*QueryGetCompoundSettingResponse, error)
	CompoundSettingAll(context.Context, *QueryAllCompoundSettingRequest) (*QueryAllCompoundSettingResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) CompoundSetting(ctx context.Context, req *QueryGetCompoundSettingRequest) (*QueryGetCompoundSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompoundSetting not implemented")
}
func (*UnimplementedQueryServer) CompoundSettingAll(ctx context.Context, req *QueryAllCompoundSettingRequest) (*QueryAllCompoundSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompoundSettingAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.compound.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CompoundSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCompoundSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CompoundSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.compound.Query/CompoundSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CompoundSetting(ctx, req.(*QueryGetCompoundSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CompoundSettingAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCompoundSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CompoundSettingAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.compound.Query/CompoundSettingAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CompoundSettingAll(ctx, req.(*QueryAllCompoundSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.compound.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "CompoundSetting",
			Handler:    _Query_CompoundSetting_Handler,
		},
		{
			MethodName: "CompoundSettingAll",
			Handler:    _Query_CompoundSettingAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temporal/compound/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetCompoundSettingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCompoundSettingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCompoundSettingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index123) > 0 {
		i -= len(m.Index123)
		copy(dAtA[i:], m.Index123)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index123)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCompoundSettingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCompoundSettingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCompoundSettingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CompoundSetting.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllCompoundSettingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCompoundSettingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCompoundSettingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCompoundSettingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCompoundSettingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCompoundSettingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CompoundSetting) > 0 {
		for iNdEx := len(m.CompoundSetting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CompoundSetting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetCompoundSettingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index123)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCompoundSettingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CompoundSetting.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCompoundSettingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCompoundSettingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CompoundSetting) > 0 {
		for _, e := range m.CompoundSetting {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetCompoundSettingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCompoundSettingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCompoundSettingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index123", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index123 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetCompoundSettingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetCompoundSettingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCompoundSettingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompoundSetting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CompoundSetting.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCompoundSettingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllCompoundSettingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCompoundSettingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCompoundSettingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllCompoundSettingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCompoundSettingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompoundSetting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompoundSetting = append(m.CompoundSetting, CompoundSetting{})
			if err := m.CompoundSetting[len(m.CompoundSetting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
