// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: incubus/merkledrop/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type QueryMerkledropRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryMerkledropRequest) Reset()         { *m = QueryMerkledropRequest{} }
func (m *QueryMerkledropRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMerkledropRequest) ProtoMessage()    {}
func (*QueryMerkledropRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bc458987e39e5e, []int{0}
}
func (m *QueryMerkledropRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMerkledropRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMerkledropRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMerkledropRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMerkledropRequest.Merge(m, src)
}
func (m *QueryMerkledropRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryMerkledropRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMerkledropRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMerkledropRequest proto.InternalMessageInfo

type QueryMerkledropResponse struct {
	Merkledrop Merkledrop `protobuf:"bytes,1,opt,name=merkledrop,proto3" json:"merkledrop"`
}

func (m *QueryMerkledropResponse) Reset()         { *m = QueryMerkledropResponse{} }
func (m *QueryMerkledropResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMerkledropResponse) ProtoMessage()    {}
func (*QueryMerkledropResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bc458987e39e5e, []int{1}
}
func (m *QueryMerkledropResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMerkledropResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMerkledropResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMerkledropResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMerkledropResponse.Merge(m, src)
}
func (m *QueryMerkledropResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMerkledropResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMerkledropResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMerkledropResponse proto.InternalMessageInfo

type QueryIndexClaimedRequest struct {
	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Index uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryIndexClaimedRequest) Reset()         { *m = QueryIndexClaimedRequest{} }
func (m *QueryIndexClaimedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIndexClaimedRequest) ProtoMessage()    {}
func (*QueryIndexClaimedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bc458987e39e5e, []int{2}
}
func (m *QueryIndexClaimedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIndexClaimedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIndexClaimedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIndexClaimedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIndexClaimedRequest.Merge(m, src)
}
func (m *QueryIndexClaimedRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIndexClaimedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIndexClaimedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIndexClaimedRequest proto.InternalMessageInfo

type QueryIndexClaimedResponse struct {
	IsClaimed bool `protobuf:"varint,1,opt,name=is_claimed,json=isClaimed,proto3" json:"is_claimed,omitempty" yaml:"is_claimed"`
}

func (m *QueryIndexClaimedResponse) Reset()         { *m = QueryIndexClaimedResponse{} }
func (m *QueryIndexClaimedResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIndexClaimedResponse) ProtoMessage()    {}
func (*QueryIndexClaimedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bc458987e39e5e, []int{3}
}
func (m *QueryIndexClaimedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIndexClaimedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIndexClaimedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIndexClaimedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIndexClaimedResponse.Merge(m, src)
}
func (m *QueryIndexClaimedResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIndexClaimedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIndexClaimedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIndexClaimedResponse proto.InternalMessageInfo

// QueryParametersRequest is request type for the Query/Parameters RPC method
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bc458987e39e5e, []int{4}
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

// QueryParametersResponse is response type for the Query/Parameters RPC method
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bc458987e39e5e, []int{5}
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

func init() {
	proto.RegisterType((*QueryMerkledropRequest)(nil), "incubus.merkledrop.v1beta1.QueryMerkledropRequest")
	proto.RegisterType((*QueryMerkledropResponse)(nil), "incubus.merkledrop.v1beta1.QueryMerkledropResponse")
	proto.RegisterType((*QueryIndexClaimedRequest)(nil), "incubus.merkledrop.v1beta1.QueryIndexClaimedRequest")
	proto.RegisterType((*QueryIndexClaimedResponse)(nil), "incubus.merkledrop.v1beta1.QueryIndexClaimedResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "incubus.merkledrop.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "incubus.merkledrop.v1beta1.QueryParamsResponse")
}

func init() {
	proto.RegisterFile("incubus/merkledrop/v1beta1/query.proto", fileDescriptor_34bc458987e39e5e)
}

var fileDescriptor_34bc458987e39e5e = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0x4f, 0x4a, 0xb7, 0xe8, 0xa7, 0x08, 0x8e, 0x55, 0x6b, 0x90, 0xac, 0x0c, 0xb2, 0x2e, 0xfe,
	0xc9, 0xb0, 0xdd, 0x0a, 0xe2, 0x69, 0xa9, 0x5e, 0x14, 0x05, 0xb7, 0x07, 0x05, 0x2f, 0x32, 0x49,
	0x66, 0xe3, 0x60, 0x92, 0xc9, 0x66, 0x52, 0xd9, 0xb2, 0xec, 0xc5, 0x27, 0x10, 0x44, 0x7c, 0x08,
	0x1f, 0xc4, 0x1e, 0x17, 0xbc, 0x78, 0x5a, 0xb4, 0xf5, 0x09, 0x7c, 0x02, 0xe9, 0xcc, 0x74, 0xd3,
	0xc5, 0x6e, 0x6a, 0x6f, 0xcd, 0x37, 0xbf, 0x7f, 0xf9, 0x7d, 0xd3, 0xc0, 0x9a, 0xcf, 0x0b, 0x29,
	0xd2, 0x88, 0x24, 0x2c, 0x7f, 0x17, 0xb3, 0x30, 0x17, 0x19, 0x79, 0xbf, 0xe1, 0xb3, 0x82, 0x6e,
	0x90, 0xdd, 0x3e, 0xcb, 0x07, 0x5e, 0x96, 0x8b, 0x42, 0x20, 0xc7, 0xe0, 0xbc, 0x12, 0xe7, 0x19,
	0x9c, 0xd3, 0x8c, 0x44, 0x24, 0x14, 0x8c, 0x4c, 0x7e, 0x69, 0x86, 0x73, 0x3d, 0x12, 0x22, 0x8a,
	0x19, 0xa1, 0x19, 0x27, 0x34, 0x4d, 0x45, 0x41, 0x0b, 0x2e, 0x52, 0x69, 0x4e, 0xdd, 0x40, 0xc8,
	0x44, 0x48, 0xe2, 0x53, 0xc9, 0x8e, 0x0d, 0x03, 0xc1, 0x53, 0x73, 0x7e, 0xa7, 0x22, 0xd7, 0x4c,
	0x04, 0x0d, 0xbe, 0x55, 0x01, 0xce, 0x68, 0x4e, 0x13, 0xe3, 0x8a, 0xd7, 0xe1, 0xca, 0xf6, 0xe4,
	0xa5, 0x9e, 0x1f, 0xe3, 0x7a, 0x6c, 0xb7, 0xcf, 0x64, 0x81, 0x2e, 0x40, 0x8d, 0x87, 0x2d, 0xfb,
	0x86, 0xbd, 0x5e, 0xef, 0xd5, 0x78, 0x88, 0x23, 0xb8, 0xfa, 0x0f, 0x52, 0x66, 0x22, 0x95, 0x0c,
	0x3d, 0x03, 0x28, 0x7d, 0x14, 0xe5, 0x5c, 0x7b, 0xcd, 0x3b, 0xbd, 0x1f, 0xaf, 0xd4, 0xe8, 0xd6,
	0x87, 0x47, 0xab, 0x56, 0x6f, 0x86, 0x8f, 0xb7, 0xa0, 0xa5, 0x8c, 0x9e, 0xa4, 0x21, 0xdb, 0x7b,
	0x14, 0x53, 0x9e, 0xb0, 0xf0, 0x94, 0x50, 0xa8, 0x09, 0x2b, 0x7c, 0x02, 0x6b, 0xd5, 0xd4, 0x48,
	0x3f, 0xe0, 0x6d, 0xb8, 0x36, 0x47, 0xc1, 0x84, 0xed, 0x00, 0x70, 0xf9, 0x26, 0xd0, 0x53, 0x25,
	0x75, 0xa6, 0x7b, 0xf9, 0xcf, 0xd1, 0xea, 0xc5, 0x01, 0x4d, 0xe2, 0x87, 0xb8, 0x3c, 0xc3, 0xbd,
	0xb3, 0x5c, 0x1a, 0x36, 0x6e, 0x02, 0x52, 0x92, 0x2f, 0x54, 0x79, 0x26, 0x0e, 0x7e, 0x05, 0x97,
	0x4e, 0x4c, 0x8d, 0xc5, 0x16, 0x34, 0x74, 0xc9, 0xa6, 0x0b, 0x5c, 0xd5, 0x85, 0xe6, 0x9a, 0x1e,
	0x0c, 0xaf, 0xfd, 0xa5, 0x0e, 0x2b, 0x4a, 0x19, 0x7d, 0xb5, 0x01, 0xca, 0xba, 0x50, 0xbb, 0x4a,
	0x6a, 0xfe, 0x26, 0x9d, 0xcd, 0xa5, 0x38, 0xfa, 0x1d, 0x70, 0xe7, 0xc3, 0xf7, 0xdf, 0x9f, 0x6a,
	0x1e, 0xba, 0x4b, 0xaa, 0xee, 0x1d, 0x9d, 0x8e, 0x24, 0xd9, 0xe7, 0xe1, 0x01, 0xfa, 0x66, 0xc3,
	0xf9, 0xd9, 0xd6, 0x51, 0x67, 0xa1, 0xf7, 0x9c, 0x35, 0x3b, 0xf7, 0x97, 0x64, 0x99, 0xcc, 0x4f,
	0x55, 0xe6, 0xc7, 0xa8, 0xbb, 0x4c, 0x66, 0xa2, 0xee, 0xcc, 0x74, 0xe9, 0x64, 0x5f, 0x3d, 0x1e,
	0xa0, 0xcf, 0x36, 0x34, 0xf4, 0x6a, 0x90, 0xb7, 0x30, 0xcd, 0x89, 0x5b, 0xe1, 0x90, 0xff, 0xc6,
	0x9b, 0xdc, 0xb7, 0x55, 0xee, 0x9b, 0x08, 0x93, 0x85, 0x7f, 0xdb, 0xee, 0xcb, 0xe1, 0x2f, 0xd7,
	0x1a, 0x8e, 0x5c, 0xfb, 0x70, 0xe4, 0xda, 0x3f, 0x47, 0xae, 0xfd, 0x71, 0xec, 0x5a, 0x87, 0x63,
	0xd7, 0xfa, 0x31, 0x76, 0xad, 0xd7, 0x0f, 0x22, 0x5e, 0xbc, 0xed, 0xfb, 0x5e, 0x20, 0x92, 0xa9,
	0x96, 0xd8, 0xd9, 0xe1, 0x01, 0xa7, 0x31, 0x89, 0xc4, 0xbd, 0xa9, 0xfc, 0xde, 0xac, 0x41, 0x31,
	0xc8, 0x98, 0xf4, 0x1b, 0xea, 0x7b, 0xb0, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xeb, 0x44,
	0xaf, 0xff, 0x04, 0x00, 0x00,
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
	Merkledrop(ctx context.Context, in *QueryMerkledropRequest, opts ...grpc.CallOption) (*QueryMerkledropResponse, error)
	IndexClaimed(ctx context.Context, in *QueryIndexClaimedRequest, opts ...grpc.CallOption) (*QueryIndexClaimedResponse, error)
	// Params queries the fantoken parameters
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Merkledrop(ctx context.Context, in *QueryMerkledropRequest, opts ...grpc.CallOption) (*QueryMerkledropResponse, error) {
	out := new(QueryMerkledropResponse)
	err := c.cc.Invoke(ctx, "/incubus.merkledrop.v1beta1.Query/Merkledrop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IndexClaimed(ctx context.Context, in *QueryIndexClaimedRequest, opts ...grpc.CallOption) (*QueryIndexClaimedResponse, error) {
	out := new(QueryIndexClaimedResponse)
	err := c.cc.Invoke(ctx, "/incubus.merkledrop.v1beta1.Query/IndexClaimed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/incubus.merkledrop.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Merkledrop(context.Context, *QueryMerkledropRequest) (*QueryMerkledropResponse, error)
	IndexClaimed(context.Context, *QueryIndexClaimedRequest) (*QueryIndexClaimedResponse, error)
	// Params queries the fantoken parameters
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Merkledrop(ctx context.Context, req *QueryMerkledropRequest) (*QueryMerkledropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Merkledrop not implemented")
}
func (*UnimplementedQueryServer) IndexClaimed(ctx context.Context, req *QueryIndexClaimedRequest) (*QueryIndexClaimedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndexClaimed not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Merkledrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMerkledropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Merkledrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/incubus.merkledrop.v1beta1.Query/Merkledrop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Merkledrop(ctx, req.(*QueryMerkledropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IndexClaimed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIndexClaimedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IndexClaimed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/incubus.merkledrop.v1beta1.Query/IndexClaimed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IndexClaimed(ctx, req.(*QueryIndexClaimedRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/incubus.merkledrop.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "incubus.merkledrop.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Merkledrop",
			Handler:    _Query_Merkledrop_Handler,
		},
		{
			MethodName: "IndexClaimed",
			Handler:    _Query_IndexClaimed_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "incubus/merkledrop/v1beta1/query.proto",
}

func (m *QueryMerkledropRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMerkledropRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMerkledropRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryMerkledropResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMerkledropResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMerkledropResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Merkledrop.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryIndexClaimedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIndexClaimedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIndexClaimedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryIndexClaimedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIndexClaimedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIndexClaimedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsClaimed {
		i--
		if m.IsClaimed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
func (m *QueryMerkledropRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryMerkledropResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Merkledrop.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryIndexClaimedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	if m.Index != 0 {
		n += 1 + sovQuery(uint64(m.Index))
	}
	return n
}

func (m *QueryIndexClaimedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsClaimed {
		n += 2
	}
	return n
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

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryMerkledropRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMerkledropRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMerkledropRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryMerkledropResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMerkledropResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMerkledropResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Merkledrop", wireType)
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
			if err := m.Merkledrop.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryIndexClaimedRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIndexClaimedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIndexClaimedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryIndexClaimedResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIndexClaimedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIndexClaimedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsClaimed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsClaimed = bool(v != 0)
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
