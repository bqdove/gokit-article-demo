// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arithmetic.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ArithmeticRequest struct {
	RequestType          string   `protobuf:"bytes,1,opt,name=requestType,proto3" json:"requestType,omitempty"`
	A                    int32    `protobuf:"varint,2,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,3,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArithmeticRequest) Reset()         { *m = ArithmeticRequest{} }
func (m *ArithmeticRequest) String() string { return proto.CompactTextString(m) }
func (*ArithmeticRequest) ProtoMessage()    {}
func (*ArithmeticRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dae331f0499ff811, []int{0}
}

func (m *ArithmeticRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArithmeticRequest.Unmarshal(m, b)
}
func (m *ArithmeticRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArithmeticRequest.Marshal(b, m, deterministic)
}
func (m *ArithmeticRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithmeticRequest.Merge(m, src)
}
func (m *ArithmeticRequest) XXX_Size() int {
	return xxx_messageInfo_ArithmeticRequest.Size(m)
}
func (m *ArithmeticRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithmeticRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArithmeticRequest proto.InternalMessageInfo

func (m *ArithmeticRequest) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *ArithmeticRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *ArithmeticRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type ArithmeticResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArithmeticResponse) Reset()         { *m = ArithmeticResponse{} }
func (m *ArithmeticResponse) String() string { return proto.CompactTextString(m) }
func (*ArithmeticResponse) ProtoMessage()    {}
func (*ArithmeticResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dae331f0499ff811, []int{1}
}

func (m *ArithmeticResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArithmeticResponse.Unmarshal(m, b)
}
func (m *ArithmeticResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArithmeticResponse.Marshal(b, m, deterministic)
}
func (m *ArithmeticResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithmeticResponse.Merge(m, src)
}
func (m *ArithmeticResponse) XXX_Size() int {
	return xxx_messageInfo_ArithmeticResponse.Size(m)
}
func (m *ArithmeticResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithmeticResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArithmeticResponse proto.InternalMessageInfo

func (m *ArithmeticResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *ArithmeticResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*ArithmeticRequest)(nil), "pb.ArithmeticRequest")
	proto.RegisterType((*ArithmeticResponse)(nil), "pb.ArithmeticResponse")
}

func init() { proto.RegisterFile("arithmetic.proto", fileDescriptor_dae331f0499ff811) }

var fileDescriptor_dae331f0499ff811 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x41, 0xcb, 0x82, 0x40,
	0x10, 0x86, 0xbf, 0x55, 0x14, 0x9c, 0xaf, 0x83, 0x0d, 0x24, 0xd2, 0x49, 0x3c, 0x79, 0xf2, 0x50,
	0xd7, 0x08, 0xa2, 0x5f, 0xd0, 0xd6, 0x1f, 0xd8, 0x95, 0x81, 0x04, 0xcb, 0x6d, 0x76, 0x0d, 0xfa,
	0xf7, 0xa1, 0x46, 0x58, 0xdd, 0xde, 0xf7, 0x39, 0x3c, 0xf3, 0x0e, 0xc4, 0x8a, 0x6b, 0x77, 0xbe,
	0x90, 0xab, 0xab, 0xd2, 0x70, 0xeb, 0x5a, 0xf4, 0x8c, 0xce, 0x0f, 0x30, 0xdf, 0xbd, 0xb9, 0xa4,
	0x5b, 0x47, 0xd6, 0x61, 0x06, 0xff, 0x3c, 0xc6, 0xd3, 0xc3, 0x50, 0x2a, 0x32, 0x51, 0x44, 0x72,
	0x8a, 0x70, 0x06, 0x42, 0xa5, 0x5e, 0x26, 0x8a, 0x40, 0x0a, 0xd5, 0x37, 0x9d, 0xfa, 0x63, 0xd3,
	0xf9, 0x16, 0x70, 0xaa, 0xb4, 0xa6, 0xbd, 0x5a, 0xc2, 0x04, 0x42, 0x26, 0xdb, 0x35, 0x6e, 0xd0,
	0x05, 0xf2, 0xd5, 0x30, 0x06, 0x9f, 0x98, 0x07, 0x57, 0x24, 0xfb, 0xb8, 0xfa, 0x98, 0x74, 0x24,
	0xbe, 0xd7, 0x15, 0xe1, 0x06, 0xa2, 0xbd, 0x6a, 0xaa, 0xae, 0x51, 0x8e, 0x70, 0x51, 0x1a, 0x5d,
	0xfe, 0xcc, 0x5e, 0x26, 0xdf, 0x78, 0x3c, 0x9d, 0xff, 0xe9, 0x70, 0x78, 0x78, 0xfd, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0xa7, 0xa3, 0xd4, 0x0a, 0x04, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArithmeticServiceClient is the client API for ArithmeticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArithmeticServiceClient interface {
	Calculate(ctx context.Context, in *ArithmeticRequest, opts ...grpc.CallOption) (*ArithmeticResponse, error)
}

type arithmeticServiceClient struct {
	cc *grpc.ClientConn
}

func NewArithmeticServiceClient(cc *grpc.ClientConn) ArithmeticServiceClient {
	return &arithmeticServiceClient{cc}
}

func (c *arithmeticServiceClient) Calculate(ctx context.Context, in *ArithmeticRequest, opts ...grpc.CallOption) (*ArithmeticResponse, error) {
	out := new(ArithmeticResponse)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticServiceServer is the server API for ArithmeticService service.
type ArithmeticServiceServer interface {
	Calculate(context.Context, *ArithmeticRequest) (*ArithmeticResponse, error)
}

// UnimplementedArithmeticServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArithmeticServiceServer struct {
}

func (*UnimplementedArithmeticServiceServer) Calculate(ctx context.Context, req *ArithmeticRequest) (*ArithmeticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}

func RegisterArithmeticServiceServer(s *grpc.Server, srv ArithmeticServiceServer) {
	s.RegisterService(&_ArithmeticService_serviceDesc, srv)
}

func _ArithmeticService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithmeticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).Calculate(ctx, req.(*ArithmeticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArithmeticService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArithmeticService",
	HandlerType: (*ArithmeticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _ArithmeticService_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arithmetic.proto",
}