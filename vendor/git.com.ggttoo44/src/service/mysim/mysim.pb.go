// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mysim.proto

package mysim

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type IntoData struct {
	Data                 []string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntoData) Reset()         { *m = IntoData{} }
func (m *IntoData) String() string { return proto.CompactTextString(m) }
func (*IntoData) ProtoMessage()    {}
func (*IntoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_630a31d6d5a159d1, []int{0}
}

func (m *IntoData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntoData.Unmarshal(m, b)
}
func (m *IntoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntoData.Marshal(b, m, deterministic)
}
func (m *IntoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntoData.Merge(m, src)
}
func (m *IntoData) XXX_Size() int {
	return xxx_messageInfo_IntoData.Size(m)
}
func (m *IntoData) XXX_DiscardUnknown() {
	xxx_messageInfo_IntoData.DiscardUnknown(m)
}

var xxx_messageInfo_IntoData proto.InternalMessageInfo

func (m *IntoData) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

// The response message containing the greetings
type OutData struct {
	CardPoint            []string `protobuf:"bytes,1,rep,name=CardPoint,proto3" json:"CardPoint,omitempty"`
	Type7                string   `protobuf:"bytes,2,opt,name=Type7,proto3" json:"Type7,omitempty"`
	Rank                 int32    `protobuf:"varint,3,opt,name=Rank,proto3" json:"Rank,omitempty"`
	Type5Ch              string   `protobuf:"bytes,4,opt,name=Type5Ch,proto3" json:"Type5Ch,omitempty"`
	Type5En              string   `protobuf:"bytes,5,opt,name=Type5En,proto3" json:"Type5En,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutData) Reset()         { *m = OutData{} }
func (m *OutData) String() string { return proto.CompactTextString(m) }
func (*OutData) ProtoMessage()    {}
func (*OutData) Descriptor() ([]byte, []int) {
	return fileDescriptor_630a31d6d5a159d1, []int{1}
}

func (m *OutData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutData.Unmarshal(m, b)
}
func (m *OutData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutData.Marshal(b, m, deterministic)
}
func (m *OutData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutData.Merge(m, src)
}
func (m *OutData) XXX_Size() int {
	return xxx_messageInfo_OutData.Size(m)
}
func (m *OutData) XXX_DiscardUnknown() {
	xxx_messageInfo_OutData.DiscardUnknown(m)
}

var xxx_messageInfo_OutData proto.InternalMessageInfo

func (m *OutData) GetCardPoint() []string {
	if m != nil {
		return m.CardPoint
	}
	return nil
}

func (m *OutData) GetType7() string {
	if m != nil {
		return m.Type7
	}
	return ""
}

func (m *OutData) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *OutData) GetType5Ch() string {
	if m != nil {
		return m.Type5Ch
	}
	return ""
}

func (m *OutData) GetType5En() string {
	if m != nil {
		return m.Type5En
	}
	return ""
}

func init() {
	proto.RegisterType((*IntoData)(nil), "mysim.IntoData")
	proto.RegisterType((*OutData)(nil), "mysim.OutData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MysimClient is the client API for Mysim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MysimClient interface {
	// Sends a greeting
	PokerCalculator(ctx context.Context, in *IntoData, opts ...grpc.CallOption) (*OutData, error)
}

type mysimClient struct {
	cc *grpc.ClientConn
}

func NewMysimClient(cc *grpc.ClientConn) MysimClient {
	return &mysimClient{cc}
}

func (c *mysimClient) PokerCalculator(ctx context.Context, in *IntoData, opts ...grpc.CallOption) (*OutData, error) {
	out := new(OutData)
	err := c.cc.Invoke(ctx, "/mysim.Mysim/PokerCalculator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MysimServer is the server API for Mysim service.
type MysimServer interface {
	// Sends a greeting
	PokerCalculator(context.Context, *IntoData) (*OutData, error)
}

func RegisterMysimServer(s *grpc.Server, srv MysimServer) {
	s.RegisterService(&_Mysim_serviceDesc, srv)
}

func _Mysim_PokerCalculator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntoData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysimServer).PokerCalculator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysim.Mysim/PokerCalculator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysimServer).PokerCalculator(ctx, req.(*IntoData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mysim_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mysim.Mysim",
	HandlerType: (*MysimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PokerCalculator",
			Handler:    _Mysim_PokerCalculator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mysim.proto",
}

func init() { proto.RegisterFile("mysim.proto", fileDescriptor_630a31d6d5a159d1) }

var fileDescriptor_630a31d6d5a159d1 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x5d, 0xdb, 0xb5, 0x66, 0x04, 0x03, 0x8b, 0x87, 0x45, 0x45, 0x42, 0x4e, 0x39, 0x2d,
	0xe2, 0x1f, 0x3c, 0x79, 0x69, 0x2c, 0xe8, 0x41, 0x0c, 0x41, 0xf0, 0x3c, 0x36, 0x8b, 0x0d, 0xdd,
	0xec, 0x84, 0xed, 0x16, 0xed, 0x03, 0xf8, 0xde, 0xb2, 0xdb, 0x54, 0xbd, 0x7d, 0xbf, 0xf9, 0x86,
	0x61, 0x66, 0xe0, 0xa8, 0xdb, 0xac, 0xda, 0x4e, 0xf5, 0x8e, 0x3c, 0x09, 0x1e, 0x43, 0x7e, 0x01,
	0x87, 0x4f, 0xd6, 0xd3, 0x03, 0x7a, 0x14, 0x02, 0xc6, 0x0d, 0x7a, 0x94, 0x2c, 0x1b, 0x15, 0x49,
	0x1d, 0x39, 0xff, 0x66, 0x30, 0x79, 0x59, 0xfb, 0xe8, 0xcf, 0x21, 0x29, 0xd1, 0x35, 0x15, 0xb5,
	0xd6, 0x0f, 0x4d, 0x7f, 0x05, 0x71, 0x02, 0xfc, 0x75, 0xd3, 0xeb, 0x3b, 0xb9, 0x9f, 0xb1, 0x22,
	0xa9, 0xb7, 0x21, 0xcc, 0xac, 0xd1, 0x2e, 0xe5, 0x28, 0x63, 0x05, 0xaf, 0x23, 0x0b, 0x09, 0x93,
	0x20, 0x6f, 0xcb, 0x85, 0x1c, 0xc7, 0xde, 0x5d, 0xfc, 0x35, 0x33, 0x2b, 0xf9, 0x3f, 0x33, 0xb3,
	0x57, 0xf7, 0xc0, 0x9f, 0xc3, 0xc2, 0xe2, 0x06, 0xd2, 0x8a, 0x96, 0xda, 0x95, 0x68, 0xe6, 0x6b,
	0x83, 0x9e, 0x9c, 0x48, 0xd5, 0xf6, 0xb0, 0xdd, 0x21, 0xa7, 0xc7, 0x43, 0x61, 0x58, 0x3c, 0xdf,
	0x9b, 0x5e, 0xc2, 0x59, 0x4b, 0xea, 0xc3, 0xf5, 0x73, 0xa5, 0xbf, 0xb0, 0xeb, 0x8d, 0x5e, 0xa9,
	0x85, 0x36, 0x86, 0x3e, 0xc9, 0x99, 0x66, 0x9a, 0x3e, 0x06, 0x7e, 0x0b, 0x5c, 0x85, 0xef, 0x54,
	0xec, 0xfd, 0x20, 0xbe, 0xe9, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x37, 0x32, 0xa1, 0x35,
	0x01, 0x00, 0x00,
}
