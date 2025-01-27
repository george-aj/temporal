// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/record/query.proto

package types

import (
	context "context"
	fmt "fmt"
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
	return fileDescriptor_b0728550ca9868d5, []int{0}
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
	return fileDescriptor_b0728550ca9868d5, []int{1}
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

type QueryGetDelegationHistoryRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryGetDelegationHistoryRequest) Reset()         { *m = QueryGetDelegationHistoryRequest{} }
func (m *QueryGetDelegationHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetDelegationHistoryRequest) ProtoMessage()    {}
func (*QueryGetDelegationHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0728550ca9868d5, []int{2}
}
func (m *QueryGetDelegationHistoryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDelegationHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDelegationHistoryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDelegationHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDelegationHistoryRequest.Merge(m, src)
}
func (m *QueryGetDelegationHistoryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDelegationHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDelegationHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDelegationHistoryRequest proto.InternalMessageInfo

func (m *QueryGetDelegationHistoryRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryGetDelegationHistoryResponse struct {
	DelegationHistory DelegationHistory `protobuf:"bytes,1,opt,name=delegationHistory,proto3" json:"delegationHistory"`
}

func (m *QueryGetDelegationHistoryResponse) Reset()         { *m = QueryGetDelegationHistoryResponse{} }
func (m *QueryGetDelegationHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetDelegationHistoryResponse) ProtoMessage()    {}
func (*QueryGetDelegationHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0728550ca9868d5, []int{3}
}
func (m *QueryGetDelegationHistoryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDelegationHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDelegationHistoryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDelegationHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDelegationHistoryResponse.Merge(m, src)
}
func (m *QueryGetDelegationHistoryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDelegationHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDelegationHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDelegationHistoryResponse proto.InternalMessageInfo

func (m *QueryGetDelegationHistoryResponse) GetDelegationHistory() DelegationHistory {
	if m != nil {
		return m.DelegationHistory
	}
	return DelegationHistory{}
}

type QueryAllDelegationHistoryRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllDelegationHistoryRequest) Reset()         { *m = QueryAllDelegationHistoryRequest{} }
func (m *QueryAllDelegationHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllDelegationHistoryRequest) ProtoMessage()    {}
func (*QueryAllDelegationHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0728550ca9868d5, []int{4}
}
func (m *QueryAllDelegationHistoryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDelegationHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDelegationHistoryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDelegationHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDelegationHistoryRequest.Merge(m, src)
}
func (m *QueryAllDelegationHistoryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDelegationHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDelegationHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDelegationHistoryRequest proto.InternalMessageInfo

func (m *QueryAllDelegationHistoryRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllDelegationHistoryResponse struct {
	DelegationHistory []DelegationHistory `protobuf:"bytes,1,rep,name=delegationHistory,proto3" json:"delegationHistory"`
	Pagination        *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllDelegationHistoryResponse) Reset()         { *m = QueryAllDelegationHistoryResponse{} }
func (m *QueryAllDelegationHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllDelegationHistoryResponse) ProtoMessage()    {}
func (*QueryAllDelegationHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0728550ca9868d5, []int{5}
}
func (m *QueryAllDelegationHistoryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDelegationHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDelegationHistoryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDelegationHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDelegationHistoryResponse.Merge(m, src)
}
func (m *QueryAllDelegationHistoryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDelegationHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDelegationHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDelegationHistoryResponse proto.InternalMessageInfo

func (m *QueryAllDelegationHistoryResponse) GetDelegationHistory() []DelegationHistory {
	if m != nil {
		return m.DelegationHistory
	}
	return nil
}

func (m *QueryAllDelegationHistoryResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "temporal.record.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "temporal.record.QueryParamsResponse")
	proto.RegisterType((*QueryGetDelegationHistoryRequest)(nil), "temporal.record.QueryGetDelegationHistoryRequest")
	proto.RegisterType((*QueryGetDelegationHistoryResponse)(nil), "temporal.record.QueryGetDelegationHistoryResponse")
	proto.RegisterType((*QueryAllDelegationHistoryRequest)(nil), "temporal.record.QueryAllDelegationHistoryRequest")
	proto.RegisterType((*QueryAllDelegationHistoryResponse)(nil), "temporal.record.QueryAllDelegationHistoryResponse")
}

func init() { proto.RegisterFile("temporal/record/query.proto", fileDescriptor_b0728550ca9868d5) }

var fileDescriptor_b0728550ca9868d5 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0x16, 0x82, 0x58, 0x0e, 0xa8, 0x4b, 0xa4, 0x16, 0x83, 0xdc, 0xb2, 0xe5, 0x4f,
	0x55, 0xd4, 0x5d, 0x12, 0xd4, 0x1b, 0x97, 0x56, 0x88, 0x22, 0xc4, 0xa1, 0xe4, 0xc0, 0x81, 0x0b,
	0xda, 0xc4, 0x23, 0xd7, 0xc8, 0xf6, 0x6e, 0xbd, 0x1b, 0x44, 0xa9, 0xb8, 0xf0, 0x04, 0x48, 0xbc,
	0x05, 0x37, 0xde, 0x81, 0x43, 0x8f, 0x95, 0xb8, 0x70, 0xaa, 0x50, 0xc2, 0x83, 0xa0, 0xec, 0xae,
	0xfb, 0xc7, 0x4e, 0x5c, 0x50, 0x6f, 0x49, 0x66, 0xe6, 0x9b, 0xdf, 0x97, 0xf9, 0x6c, 0x74, 0x4b,
	0x43, 0x2a, 0x45, 0xce, 0x13, 0x96, 0x43, 0x5f, 0xe4, 0x21, 0xdb, 0x1d, 0x40, 0xbe, 0x47, 0x65,
	0x2e, 0xb4, 0xc0, 0xd7, 0x8b, 0x22, 0xb5, 0x45, 0xbf, 0x15, 0x89, 0x48, 0x98, 0x1a, 0x1b, 0x7f,
	0xb2, 0x6d, 0xfe, 0xed, 0x48, 0x88, 0x28, 0x01, 0xc6, 0x65, 0xcc, 0x78, 0x96, 0x09, 0xcd, 0x75,
	0x2c, 0x32, 0xe5, 0xaa, 0xab, 0x7d, 0xa1, 0x52, 0xa1, 0x58, 0x8f, 0x2b, 0xb0, 0xea, 0xec, 0x7d,
	0xbb, 0x07, 0x9a, 0xb7, 0x99, 0xe4, 0x51, 0x9c, 0x99, 0xe6, 0x42, 0xa9, 0x4c, 0x23, 0x79, 0xce,
	0xd3, 0x42, 0x69, 0xa5, 0x5c, 0x0d, 0x21, 0x81, 0xc8, 0xcc, 0xbf, 0xdd, 0x89, 0x95, 0x16, 0x05,
	0xb8, 0xbf, 0x5a, 0xd3, 0xa9, 0xe3, 0x14, 0x94, 0xe6, 0xa9, 0xb4, 0xbd, 0xa4, 0x85, 0xf0, 0xab,
	0x31, 0xd5, 0xb6, 0x59, 0xd5, 0x85, 0xdd, 0x01, 0x28, 0x4d, 0x5e, 0xa2, 0x1b, 0x67, 0x7e, 0x55,
	0x52, 0x64, 0x0a, 0xf0, 0x3a, 0x6a, 0x5a, 0xa4, 0x05, 0x6f, 0xc9, 0x5b, 0xb9, 0xd6, 0x99, 0xa7,
	0xa5, 0xbf, 0x88, 0xda, 0x81, 0xcd, 0x4b, 0x07, 0x47, 0x8b, 0x8d, 0xae, 0x6b, 0x26, 0x4f, 0xd0,
	0x92, 0x51, 0xdb, 0x02, 0xfd, 0xf4, 0x98, 0xe4, 0xb9, 0x45, 0x76, 0x1b, 0xf1, 0x02, 0xba, 0xc2,
	0xc3, 0x30, 0x07, 0x65, 0xb5, 0xaf, 0x76, 0x8b, 0xaf, 0x64, 0x1f, 0xdd, 0xa9, 0x99, 0x76, 0x64,
	0xaf, 0xd1, 0x5c, 0x58, 0x2e, 0x3a, 0x48, 0x52, 0x81, 0xac, 0xc8, 0x38, 0xde, 0xaa, 0x04, 0x79,
	0xe7, 0xd0, 0x37, 0x92, 0x64, 0x2a, 0xfa, 0x33, 0x84, 0x4e, 0x4e, 0xe9, 0x96, 0xde, 0xa7, 0xf6,
	0xee, 0x74, 0x7c, 0x77, 0x6a, 0x53, 0xe5, 0xee, 0x4e, 0xb7, 0x79, 0x04, 0x6e, 0xb6, 0x7b, 0x6a,
	0x92, 0xfc, 0xf0, 0x9c, 0xd3, 0xc9, 0xcb, 0xea, 0x9d, 0xce, 0x5e, 0xd0, 0x29, 0xde, 0x3a, 0xe3,
	0x62, 0xc6, 0xb8, 0x78, 0x70, 0xae, 0x0b, 0x0b, 0x75, 0xda, 0x46, 0xe7, 0x68, 0x16, 0x5d, 0x36,
	0x36, 0xb0, 0x46, 0x4d, 0x9b, 0x07, 0xbc, 0x5c, 0x21, 0xab, 0x86, 0xce, 0xbf, 0x5b, 0xdf, 0x64,
	0x57, 0x91, 0xc5, 0xcf, 0x3f, 0xff, 0x7c, 0x9d, 0xb9, 0x89, 0xe7, 0xd9, 0xe4, 0xa7, 0x05, 0x7f,
	0xf7, 0xd0, 0x5c, 0xc5, 0x37, 0x6e, 0x4f, 0x16, 0xaf, 0x89, 0xa4, 0xdf, 0xf9, 0x9f, 0x11, 0x47,
	0xb7, 0x6e, 0xe8, 0x18, 0x5e, 0x63, 0xe7, 0x3f, 0xad, 0x6c, 0xdf, 0x45, 0xfc, 0x13, 0xfe, 0xe6,
	0xa1, 0x56, 0x45, 0x74, 0x23, 0x49, 0xa6, 0x61, 0xd7, 0xc4, 0x71, 0x1a, 0x76, 0x5d, 0xa8, 0xc8,
	0x43, 0x83, 0x7d, 0x0f, 0x2f, 0xff, 0x03, 0xf6, 0xe6, 0x8b, 0x83, 0x61, 0xe0, 0x1d, 0x0e, 0x03,
	0xef, 0xf7, 0x30, 0xf0, 0xbe, 0x8c, 0x82, 0xc6, 0xe1, 0x28, 0x68, 0xfc, 0x1a, 0x05, 0x8d, 0x37,
	0x8f, 0xa2, 0x58, 0xef, 0x0c, 0x7a, 0xb4, 0x2f, 0xd2, 0x63, 0xa1, 0xb5, 0x8f, 0x22, 0x83, 0x13,
	0xd9, 0x0f, 0x85, 0xb0, 0xde, 0x93, 0xa0, 0x7a, 0x4d, 0xf3, 0x16, 0x7a, 0xfc, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x37, 0x3e, 0xf5, 0x01, 0x89, 0x05, 0x00, 0x00,
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
	// Queries a list of DelegationHistory items.
	DelegationHistory(ctx context.Context, in *QueryGetDelegationHistoryRequest, opts ...grpc.CallOption) (*QueryGetDelegationHistoryResponse, error)
	DelegationHistoryAll(ctx context.Context, in *QueryAllDelegationHistoryRequest, opts ...grpc.CallOption) (*QueryAllDelegationHistoryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/temporal.record.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegationHistory(ctx context.Context, in *QueryGetDelegationHistoryRequest, opts ...grpc.CallOption) (*QueryGetDelegationHistoryResponse, error) {
	out := new(QueryGetDelegationHistoryResponse)
	err := c.cc.Invoke(ctx, "/temporal.record.Query/DelegationHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegationHistoryAll(ctx context.Context, in *QueryAllDelegationHistoryRequest, opts ...grpc.CallOption) (*QueryAllDelegationHistoryResponse, error) {
	out := new(QueryAllDelegationHistoryResponse)
	err := c.cc.Invoke(ctx, "/temporal.record.Query/DelegationHistoryAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of DelegationHistory items.
	DelegationHistory(context.Context, *QueryGetDelegationHistoryRequest) (*QueryGetDelegationHistoryResponse, error)
	DelegationHistoryAll(context.Context, *QueryAllDelegationHistoryRequest) (*QueryAllDelegationHistoryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DelegationHistory(ctx context.Context, req *QueryGetDelegationHistoryRequest) (*QueryGetDelegationHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationHistory not implemented")
}
func (*UnimplementedQueryServer) DelegationHistoryAll(ctx context.Context, req *QueryAllDelegationHistoryRequest) (*QueryAllDelegationHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationHistoryAll not implemented")
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
		FullMethod: "/temporal.record.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegationHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDelegationHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.record.Query/DelegationHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationHistory(ctx, req.(*QueryGetDelegationHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegationHistoryAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllDelegationHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationHistoryAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.record.Query/DelegationHistoryAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationHistoryAll(ctx, req.(*QueryAllDelegationHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.record.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DelegationHistory",
			Handler:    _Query_DelegationHistory_Handler,
		},
		{
			MethodName: "DelegationHistoryAll",
			Handler:    _Query_DelegationHistoryAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temporal/record/query.proto",
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

func (m *QueryGetDelegationHistoryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDelegationHistoryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDelegationHistoryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDelegationHistoryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDelegationHistoryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDelegationHistoryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DelegationHistory.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllDelegationHistoryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDelegationHistoryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDelegationHistoryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllDelegationHistoryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDelegationHistoryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDelegationHistoryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.DelegationHistory) > 0 {
		for iNdEx := len(m.DelegationHistory) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationHistory[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryGetDelegationHistoryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetDelegationHistoryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DelegationHistory.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllDelegationHistoryRequest) Size() (n int) {
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

func (m *QueryAllDelegationHistoryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DelegationHistory) > 0 {
		for _, e := range m.DelegationHistory {
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
func (m *QueryGetDelegationHistoryRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetDelegationHistoryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDelegationHistoryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetDelegationHistoryResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetDelegationHistoryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDelegationHistoryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationHistory", wireType)
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
			if err := m.DelegationHistory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllDelegationHistoryRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllDelegationHistoryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDelegationHistoryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllDelegationHistoryResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllDelegationHistoryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDelegationHistoryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationHistory", wireType)
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
			m.DelegationHistory = append(m.DelegationHistory, DelegationHistory{})
			if err := m.DelegationHistory[len(m.DelegationHistory)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
