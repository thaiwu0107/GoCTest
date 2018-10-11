// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flopplayer2.proto

package flopplayer2

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
	Data                 []*IntoData_Record `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	PublicPoker          []int32            `protobuf:"varint,2,rep,packed,name=PublicPoker,proto3" json:"PublicPoker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IntoData) Reset()         { *m = IntoData{} }
func (m *IntoData) String() string { return proto.CompactTextString(m) }
func (*IntoData) ProtoMessage()    {}
func (*IntoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7fd05b5f85ad55, []int{0}
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

func (m *IntoData) GetData() []*IntoData_Record {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *IntoData) GetPublicPoker() []int32 {
	if m != nil {
		return m.PublicPoker
	}
	return nil
}

type IntoData_Record struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Pokers               []int32  `protobuf:"varint,2,rep,packed,name=Pokers,proto3" json:"Pokers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntoData_Record) Reset()         { *m = IntoData_Record{} }
func (m *IntoData_Record) String() string { return proto.CompactTextString(m) }
func (*IntoData_Record) ProtoMessage()    {}
func (*IntoData_Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7fd05b5f85ad55, []int{0, 0}
}

func (m *IntoData_Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntoData_Record.Unmarshal(m, b)
}
func (m *IntoData_Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntoData_Record.Marshal(b, m, deterministic)
}
func (m *IntoData_Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntoData_Record.Merge(m, src)
}
func (m *IntoData_Record) XXX_Size() int {
	return xxx_messageInfo_IntoData_Record.Size(m)
}
func (m *IntoData_Record) XXX_DiscardUnknown() {
	xxx_messageInfo_IntoData_Record.DiscardUnknown(m)
}

var xxx_messageInfo_IntoData_Record proto.InternalMessageInfo

func (m *IntoData_Record) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IntoData_Record) GetPokers() []int32 {
	if m != nil {
		return m.Pokers
	}
	return nil
}

// The response message containing the greetings
type OutData struct {
	OverCardNumber       int32    `protobuf:"varint,1,opt,name=OverCardNumber,proto3" json:"OverCardNumber"`
	LeaderId             string   `protobuf:"bytes,2,opt,name=LeaderId,proto3" json:"LeaderId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutData) Reset()         { *m = OutData{} }
func (m *OutData) String() string { return proto.CompactTextString(m) }
func (*OutData) ProtoMessage()    {}
func (*OutData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7fd05b5f85ad55, []int{1}
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

func (m *OutData) GetOverCardNumber() int32 {
	if m != nil {
		return m.OverCardNumber
	}
	return 0
}

func (m *OutData) GetLeaderId() string {
	if m != nil {
		return m.LeaderId
	}
	return ""
}

func init() {
	proto.RegisterType((*IntoData)(nil), "flopplayer2.IntoData")
	proto.RegisterType((*IntoData_Record)(nil), "flopplayer2.IntoData.Record")
	proto.RegisterType((*OutData)(nil), "flopplayer2.OutData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Flopplayer2Client is the client API for Flopplayer2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Flopplayer2Client interface {
	// Sends a greeting
	CalculatorFlop(ctx context.Context, in *IntoData, opts ...grpc.CallOption) (*OutData, error)
}

type flopplayer2Client struct {
	cc *grpc.ClientConn
}

func NewFlopplayer2Client(cc *grpc.ClientConn) Flopplayer2Client {
	return &flopplayer2Client{cc}
}

func (c *flopplayer2Client) CalculatorFlop(ctx context.Context, in *IntoData, opts ...grpc.CallOption) (*OutData, error) {
	out := new(OutData)
	err := c.cc.Invoke(ctx, "/flopplayer2.Flopplayer2/CalculatorFlop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Flopplayer2Server is the server API for Flopplayer2 service.
type Flopplayer2Server interface {
	// Sends a greeting
	CalculatorFlop(context.Context, *IntoData) (*OutData, error)
}

func RegisterFlopplayer2Server(s *grpc.Server, srv Flopplayer2Server) {
	s.RegisterService(&_Flopplayer2_serviceDesc, srv)
}

func _Flopplayer2_CalculatorFlop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntoData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Flopplayer2Server).CalculatorFlop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flopplayer2.Flopplayer2/CalculatorFlop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Flopplayer2Server).CalculatorFlop(ctx, req.(*IntoData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Flopplayer2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flopplayer2.Flopplayer2",
	HandlerType: (*Flopplayer2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculatorFlop",
			Handler:    _Flopplayer2_CalculatorFlop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flopplayer2.proto",
}

func init() { proto.RegisterFile("flopplayer2.proto", fileDescriptor_1e7fd05b5f85ad55) }

var fileDescriptor_1e7fd05b5f85ad55 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x4d, 0xc6, 0xa9, 0xe3, 0x2b, 0x14, 0x7c, 0xa8, 0x94, 0xe2, 0xa2, 0x74, 0x21, 0x5d,
	0x95, 0xa1, 0x1e, 0x40, 0x70, 0x44, 0x28, 0xe8, 0x4c, 0xc9, 0x0d, 0xd2, 0x26, 0x82, 0x18, 0x4d,
	0x79, 0x26, 0x82, 0xb7, 0xf0, 0xc8, 0xd2, 0x58, 0xb5, 0x8a, 0xab, 0xf0, 0xbe, 0xfc, 0x5f, 0x92,
	0x3f, 0x70, 0x74, 0x6f, 0xec, 0x30, 0x18, 0xf9, 0xa6, 0xa9, 0xae, 0x06, 0xb2, 0xce, 0x62, 0x3c,
	0x43, 0xc5, 0x3b, 0x83, 0x55, 0xf3, 0xec, 0xec, 0xb5, 0x74, 0x12, 0xd7, 0xb0, 0x3f, 0xae, 0x29,
	0xcb, 0x17, 0x65, 0x5c, 0x9f, 0x55, 0x73, 0xf7, 0x2b, 0x54, 0x09, 0xdd, 0x5b, 0x52, 0x22, 0x24,
	0x31, 0x87, 0xb8, 0xf5, 0x9d, 0x79, 0xe8, 0x5b, 0xfb, 0xa8, 0x29, 0xe5, 0xf9, 0xa2, 0x5c, 0x8a,
	0x39, 0xca, 0xd6, 0x10, 0x7d, 0x1a, 0x98, 0x00, 0x6f, 0x54, 0xca, 0x72, 0x56, 0x1e, 0x0a, 0xde,
	0x28, 0x3c, 0x85, 0x28, 0x44, 0x5e, 0x26, 0x6d, 0x9a, 0x8a, 0x3b, 0x38, 0xd8, 0x79, 0x17, 0x8e,
	0x3f, 0x87, 0x64, 0xf7, 0xaa, 0x69, 0x23, 0x49, 0x6d, 0xfd, 0x53, 0xa7, 0x29, 0xe8, 0x4b, 0xf1,
	0x87, 0x62, 0x06, 0xab, 0x5b, 0x2d, 0x95, 0xa6, 0x46, 0xa5, 0x3c, 0x5c, 0xf0, 0x3d, 0xd7, 0x5b,
	0x88, 0x6f, 0x7e, 0x7a, 0xe0, 0x25, 0x24, 0x1b, 0x69, 0x7a, 0x6f, 0xa4, 0xb3, 0x34, 0x6e, 0xe0,
	0xc9, 0xbf, 0x3d, 0xb3, 0xe3, 0x5f, 0x78, 0x7a, 0x51, 0xb1, 0x77, 0xc5, 0x5b, 0xd6, 0x45, 0xe1,
	0x27, 0x2f, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x5d, 0x1c, 0x0a, 0x5e, 0x01, 0x00, 0x00,
}
