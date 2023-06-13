// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: incubus/fantoken/v1beta1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type EventIssue struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *EventIssue) Reset()         { *m = EventIssue{} }
func (m *EventIssue) String() string { return proto.CompactTextString(m) }
func (*EventIssue) ProtoMessage()    {}
func (*EventIssue) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e6474600913d18, []int{0}
}
func (m *EventIssue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventIssue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventIssue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventIssue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventIssue.Merge(m, src)
}
func (m *EventIssue) XXX_Size() int {
	return m.Size()
}
func (m *EventIssue) XXX_DiscardUnknown() {
	xxx_messageInfo_EventIssue.DiscardUnknown(m)
}

var xxx_messageInfo_EventIssue proto.InternalMessageInfo

func (m *EventIssue) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type EventDisableMint struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *EventDisableMint) Reset()         { *m = EventDisableMint{} }
func (m *EventDisableMint) String() string { return proto.CompactTextString(m) }
func (*EventDisableMint) ProtoMessage()    {}
func (*EventDisableMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e6474600913d18, []int{1}
}
func (m *EventDisableMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDisableMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDisableMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDisableMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDisableMint.Merge(m, src)
}
func (m *EventDisableMint) XXX_Size() int {
	return m.Size()
}
func (m *EventDisableMint) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDisableMint.DiscardUnknown(m)
}

var xxx_messageInfo_EventDisableMint proto.InternalMessageInfo

func (m *EventDisableMint) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type EventMint struct {
	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Coin      string `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin,omitempty"`
}

func (m *EventMint) Reset()         { *m = EventMint{} }
func (m *EventMint) String() string { return proto.CompactTextString(m) }
func (*EventMint) ProtoMessage()    {}
func (*EventMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e6474600913d18, []int{2}
}
func (m *EventMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMint.Merge(m, src)
}
func (m *EventMint) XXX_Size() int {
	return m.Size()
}
func (m *EventMint) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMint.DiscardUnknown(m)
}

var xxx_messageInfo_EventMint proto.InternalMessageInfo

func (m *EventMint) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *EventMint) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

type EventBurn struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Coin   string `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin,omitempty"`
}

func (m *EventBurn) Reset()         { *m = EventBurn{} }
func (m *EventBurn) String() string { return proto.CompactTextString(m) }
func (*EventBurn) ProtoMessage()    {}
func (*EventBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e6474600913d18, []int{3}
}
func (m *EventBurn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBurn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBurn.Merge(m, src)
}
func (m *EventBurn) XXX_Size() int {
	return m.Size()
}
func (m *EventBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBurn.DiscardUnknown(m)
}

var xxx_messageInfo_EventBurn proto.InternalMessageInfo

func (m *EventBurn) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventBurn) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

type EventSetAuthority struct {
	Denom        string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	OldAuthority string `protobuf:"bytes,2,opt,name=old_authority,json=oldAuthority,proto3" json:"old_authority,omitempty" yaml:"old_authority"`
	NewAuthority string `protobuf:"bytes,3,opt,name=new_authority,json=newAuthority,proto3" json:"new_authority,omitempty" yaml:"new_authority"`
}

func (m *EventSetAuthority) Reset()         { *m = EventSetAuthority{} }
func (m *EventSetAuthority) String() string { return proto.CompactTextString(m) }
func (*EventSetAuthority) ProtoMessage()    {}
func (*EventSetAuthority) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e6474600913d18, []int{4}
}
func (m *EventSetAuthority) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetAuthority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetAuthority.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetAuthority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetAuthority.Merge(m, src)
}
func (m *EventSetAuthority) XXX_Size() int {
	return m.Size()
}
func (m *EventSetAuthority) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetAuthority.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetAuthority proto.InternalMessageInfo

func (m *EventSetAuthority) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventSetAuthority) GetOldAuthority() string {
	if m != nil {
		return m.OldAuthority
	}
	return ""
}

func (m *EventSetAuthority) GetNewAuthority() string {
	if m != nil {
		return m.NewAuthority
	}
	return ""
}

type EventSetMinter struct {
	Denom     string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	OldMinter string `protobuf:"bytes,2,opt,name=old_minter,json=oldMinter,proto3" json:"old_minter,omitempty" yaml:"old_minter"`
	NewMinter string `protobuf:"bytes,3,opt,name=new_minter,json=newMinter,proto3" json:"new_minter,omitempty" yaml:"new_minter"`
}

func (m *EventSetMinter) Reset()         { *m = EventSetMinter{} }
func (m *EventSetMinter) String() string { return proto.CompactTextString(m) }
func (*EventSetMinter) ProtoMessage()    {}
func (*EventSetMinter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e6474600913d18, []int{5}
}
func (m *EventSetMinter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetMinter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetMinter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetMinter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetMinter.Merge(m, src)
}
func (m *EventSetMinter) XXX_Size() int {
	return m.Size()
}
func (m *EventSetMinter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetMinter.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetMinter proto.InternalMessageInfo

func (m *EventSetMinter) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventSetMinter) GetOldMinter() string {
	if m != nil {
		return m.OldMinter
	}
	return ""
}

func (m *EventSetMinter) GetNewMinter() string {
	if m != nil {
		return m.NewMinter
	}
	return ""
}

type EventSetUri struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *EventSetUri) Reset()         { *m = EventSetUri{} }
func (m *EventSetUri) String() string { return proto.CompactTextString(m) }
func (*EventSetUri) ProtoMessage()    {}
func (*EventSetUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e6474600913d18, []int{6}
}
func (m *EventSetUri) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetUri.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetUri.Merge(m, src)
}
func (m *EventSetUri) XXX_Size() int {
	return m.Size()
}
func (m *EventSetUri) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetUri.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetUri proto.InternalMessageInfo

func (m *EventSetUri) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*EventIssue)(nil), "incubus.fantoken.v1beta1.EventIssue")
	proto.RegisterType((*EventDisableMint)(nil), "incubus.fantoken.v1beta1.EventDisableMint")
	proto.RegisterType((*EventMint)(nil), "incubus.fantoken.v1beta1.EventMint")
	proto.RegisterType((*EventBurn)(nil), "incubus.fantoken.v1beta1.EventBurn")
	proto.RegisterType((*EventSetAuthority)(nil), "incubus.fantoken.v1beta1.EventSetAuthority")
	proto.RegisterType((*EventSetMinter)(nil), "incubus.fantoken.v1beta1.EventSetMinter")
	proto.RegisterType((*EventSetUri)(nil), "incubus.fantoken.v1beta1.EventSetUri")
}

func init() {
	proto.RegisterFile("incubus/fantoken/v1beta1/events.proto", fileDescriptor_c4e6474600913d18)
}

var fileDescriptor_c4e6474600913d18 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x1b, 0xab, 0x85, 0xbc, 0xb6, 0x62, 0x43, 0x95, 0x20, 0x92, 0xca, 0x88, 0xd0, 0x8b,
	0x09, 0xc5, 0xa2, 0x20, 0xf4, 0x60, 0xd1, 0x83, 0x07, 0x41, 0x22, 0x5e, 0xbc, 0x48, 0xfe, 0xbc,
	0x4d, 0x07, 0x93, 0x99, 0x92, 0x4c, 0x5a, 0xfb, 0x2d, 0x96, 0xfd, 0x12, 0xfb, 0x55, 0xf6, 0xd8,
	0xe3, 0x9e, 0xca, 0xd2, 0x7e, 0x83, 0x7e, 0x82, 0x25, 0x93, 0xc9, 0x66, 0x5b, 0xda, 0xdb, 0x4c,
	0x9e, 0xdf, 0xef, 0xcd, 0xc3, 0xf0, 0xc2, 0x3b, 0x9f, 0x8a, 0x8c, 0xb3, 0xc8, 0x99, 0x7a, 0x4c,
	0xf0, 0x7f, 0xc8, 0x9c, 0xc5, 0xd0, 0x47, 0xe1, 0x0d, 0x1d, 0x5c, 0x20, 0x13, 0x99, 0x3d, 0x4f,
	0xb9, 0xe0, 0x86, 0xa9, 0x30, 0xbb, 0xc2, 0x6c, 0x85, 0xbd, 0xea, 0x45, 0x3c, 0xe2, 0x12, 0x72,
	0x8a, 0x53, 0xc9, 0x13, 0x02, 0xf0, 0xad, 0xf0, 0xbf, 0x67, 0x59, 0x8e, 0x46, 0x0f, 0x9e, 0x84,
	0xc8, 0x78, 0x62, 0x6a, 0x6f, 0xb4, 0x81, 0xee, 0x96, 0x17, 0x32, 0x80, 0xe7, 0x92, 0xf9, 0x4a,
	0x33, 0xcf, 0x8f, 0xf1, 0x07, 0x65, 0xe2, 0x0c, 0x39, 0x06, 0x5d, 0x92, 0x12, 0x79, 0x0d, 0x7a,
	0x8a, 0x01, 0x9d, 0x53, 0x64, 0x42, 0x61, 0xf5, 0x07, 0xc3, 0x80, 0xc7, 0x01, 0xa7, 0xcc, 0x7c,
	0x24, 0x03, 0x79, 0x26, 0x9f, 0x94, 0x3e, 0xc9, 0x53, 0x66, 0xbc, 0x84, 0x56, 0x86, 0x2c, 0xc4,
	0x54, 0xb9, 0xea, 0x76, 0x52, 0xbc, 0xd2, 0xa0, 0x2b, 0xcd, 0x5f, 0x28, 0xbe, 0xe4, 0x62, 0xc6,
	0x53, 0x2a, 0x56, 0xa7, 0x3b, 0x1a, 0x63, 0xe8, 0xf0, 0x38, 0xfc, 0xeb, 0x55, 0x58, 0x39, 0x68,
	0x62, 0xee, 0x37, 0xfd, 0xde, 0xca, 0x4b, 0xe2, 0xcf, 0xe4, 0x20, 0x26, 0x6e, 0x9b, 0xc7, 0x61,
	0x3d, 0x74, 0x0c, 0x1d, 0x86, 0xcb, 0x07, 0x7a, 0xf3, 0x58, 0x3f, 0x88, 0x89, 0xdb, 0x66, 0xb8,
	0xbc, 0xd7, 0xc9, 0xa5, 0x06, 0xcf, 0xaa, 0xa6, 0xc5, 0x2b, 0x61, 0x7a, 0xa6, 0xe6, 0x08, 0xa0,
	0xe8, 0x91, 0x48, 0x46, 0x75, 0x7c, 0xb1, 0xdf, 0xf4, 0xbb, 0x75, 0xc7, 0x32, 0x23, 0xae, 0xce,
	0xe3, 0x50, 0xcd, 0x1a, 0x01, 0x14, 0xbf, 0x57, 0x56, 0xf3, 0xd8, 0xaa, 0x33, 0xe2, 0xea, 0x0c,
	0x97, 0xa5, 0x45, 0xde, 0xc2, 0xd3, 0xaa, 0xd3, 0xef, 0x94, 0x9e, 0x2e, 0x34, 0xf9, 0x79, 0xbd,
	0xb5, 0xb4, 0xf5, 0xd6, 0xd2, 0x6e, 0xb7, 0x96, 0x76, 0xb1, 0xb3, 0x1a, 0xeb, 0x9d, 0xd5, 0xb8,
	0xd9, 0x59, 0x8d, 0x3f, 0x1f, 0x23, 0x2a, 0x66, 0xb9, 0x6f, 0x07, 0x3c, 0x71, 0xd4, 0xfa, 0xf1,
	0xe9, 0x94, 0x06, 0xd4, 0x8b, 0x9d, 0x88, 0xbf, 0xaf, 0x16, 0xf7, 0x7f, 0xbd, 0xba, 0x62, 0x35,
	0xc7, 0xcc, 0x6f, 0xc9, 0x15, 0xfc, 0x70, 0x17, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x2e, 0x76, 0x59,
	0xdb, 0x02, 0x00, 0x00,
}

func (m *EventIssue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventIssue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventIssue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventDisableMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDisableMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDisableMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coin) > 0 {
		i -= len(m.Coin)
		copy(dAtA[i:], m.Coin)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Coin)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBurn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBurn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBurn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coin) > 0 {
		i -= len(m.Coin)
		copy(dAtA[i:], m.Coin)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Coin)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSetAuthority) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetAuthority) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetAuthority) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewAuthority) > 0 {
		i -= len(m.NewAuthority)
		copy(dAtA[i:], m.NewAuthority)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewAuthority)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OldAuthority) > 0 {
		i -= len(m.OldAuthority)
		copy(dAtA[i:], m.OldAuthority)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.OldAuthority)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSetMinter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetMinter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetMinter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewMinter) > 0 {
		i -= len(m.NewMinter)
		copy(dAtA[i:], m.NewMinter)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewMinter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OldMinter) > 0 {
		i -= len(m.OldMinter)
		copy(dAtA[i:], m.OldMinter)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.OldMinter)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSetUri) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetUri) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetUri) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventIssue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventDisableMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Coin)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventBurn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Coin)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetAuthority) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.OldAuthority)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewAuthority)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetMinter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.OldMinter)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewMinter)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetUri) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventIssue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventIssue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventIssue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventDisableMint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventDisableMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDisableMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventMint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventBurn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventBurn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBurn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSetAuthority) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetAuthority: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetAuthority: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldAuthority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldAuthority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewAuthority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewAuthority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSetMinter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetMinter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetMinter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldMinter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldMinter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewMinter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewMinter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSetUri) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetUri: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetUri: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)