// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ambassador.proto

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19084e700d1da46, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ListTopicsResponse struct {
	Topics               []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopicsResponse) Reset()         { *m = ListTopicsResponse{} }
func (m *ListTopicsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTopicsResponse) ProtoMessage()    {}
func (*ListTopicsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19084e700d1da46, []int{1}
}

func (m *ListTopicsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsResponse.Unmarshal(m, b)
}
func (m *ListTopicsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsResponse.Marshal(b, m, deterministic)
}
func (m *ListTopicsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsResponse.Merge(m, src)
}
func (m *ListTopicsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTopicsResponse.Size(m)
}
func (m *ListTopicsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsResponse proto.InternalMessageInfo

func (m *ListTopicsResponse) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type ProdRq struct {
	DeploymentOwnerUserName string   `protobuf:"bytes,1,opt,name=deploymentOwnerUserName,proto3" json:"deploymentOwnerUserName,omitempty"`
	DeploymentName          string   `protobuf:"bytes,2,opt,name=deploymentName,proto3" json:"deploymentName,omitempty"`
	EndpointOutput          string   `protobuf:"bytes,3,opt,name=endpointOutput,proto3" json:"endpointOutput,omitempty"`
	Message                 []byte   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	StreamOffset            uint64   `protobuf:"varint,5,opt,name=streamOffset,proto3" json:"streamOffset,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ProdRq) Reset()         { *m = ProdRq{} }
func (m *ProdRq) String() string { return proto.CompactTextString(m) }
func (*ProdRq) ProtoMessage()    {}
func (*ProdRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19084e700d1da46, []int{2}
}

func (m *ProdRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRq.Unmarshal(m, b)
}
func (m *ProdRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRq.Marshal(b, m, deterministic)
}
func (m *ProdRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRq.Merge(m, src)
}
func (m *ProdRq) XXX_Size() int {
	return xxx_messageInfo_ProdRq.Size(m)
}
func (m *ProdRq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRq.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRq proto.InternalMessageInfo

func (m *ProdRq) GetDeploymentOwnerUserName() string {
	if m != nil {
		return m.DeploymentOwnerUserName
	}
	return ""
}

func (m *ProdRq) GetDeploymentName() string {
	if m != nil {
		return m.DeploymentName
	}
	return ""
}

func (m *ProdRq) GetEndpointOutput() string {
	if m != nil {
		return m.EndpointOutput
	}
	return ""
}

func (m *ProdRq) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ProdRq) GetStreamOffset() uint64 {
	if m != nil {
		return m.StreamOffset
	}
	return 0
}

type ProdRs struct {
	StreamOffset         uint64   `protobuf:"varint,3,opt,name=streamOffset,proto3" json:"streamOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdRs) Reset()         { *m = ProdRs{} }
func (m *ProdRs) String() string { return proto.CompactTextString(m) }
func (*ProdRs) ProtoMessage()    {}
func (*ProdRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19084e700d1da46, []int{3}
}

func (m *ProdRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRs.Unmarshal(m, b)
}
func (m *ProdRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRs.Marshal(b, m, deterministic)
}
func (m *ProdRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRs.Merge(m, src)
}
func (m *ProdRs) XXX_Size() int {
	return xxx_messageInfo_ProdRs.Size(m)
}
func (m *ProdRs) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRs.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRs proto.InternalMessageInfo

func (m *ProdRs) GetStreamOffset() uint64 {
	if m != nil {
		return m.StreamOffset
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*ListTopicsResponse)(nil), "ListTopicsResponse")
	proto.RegisterType((*ProdRq)(nil), "ProdRq")
	proto.RegisterType((*ProdRs)(nil), "ProdRs")
}

func init() { proto.RegisterFile("ambassador.proto", fileDescriptor_c19084e700d1da46) }

var fileDescriptor_c19084e700d1da46 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x97, 0xfd, 0x68, 0xd9, 0x63, 0xa8, 0x44, 0xd0, 0xb0, 0x53, 0xcd, 0x41, 0x0a, 0x8e,
	0x22, 0x7a, 0xf1, 0xaa, 0xe0, 0x49, 0x71, 0x12, 0xf4, 0xe2, 0x2d, 0x5b, 0x5f, 0xa5, 0x68, 0x9a,
	0x98, 0x97, 0x22, 0xfb, 0x2f, 0xfd, 0x93, 0x64, 0x5d, 0x75, 0xcc, 0xe2, 0x2d, 0xef, 0xfb, 0x3e,
	0xdf, 0x40, 0x3e, 0x81, 0x03, 0x6d, 0x16, 0x9a, 0x48, 0xe7, 0xd6, 0x67, 0xce, 0xdb, 0x60, 0x65,
	0x0c, 0xa3, 0x5b, 0xe3, 0xc2, 0x4a, 0xce, 0x80, 0xdf, 0x97, 0x14, 0x9e, 0xac, 0x2b, 0x97, 0xa4,
	0x90, 0x9c, 0xad, 0x08, 0xf9, 0x11, 0x44, 0xa1, 0x49, 0x04, 0x4b, 0x06, 0xe9, 0x58, 0xb5, 0x93,
	0xfc, 0x62, 0x10, 0x3d, 0x7a, 0x9b, 0xab, 0x0f, 0x7e, 0x05, 0xc7, 0x39, 0xba, 0x77, 0xbb, 0x32,
	0x58, 0x85, 0xf9, 0x67, 0x85, 0xfe, 0x99, 0xd0, 0x3f, 0x68, 0x83, 0x82, 0x25, 0x2c, 0x1d, 0xab,
	0xff, 0xd6, 0xfc, 0x14, 0xf6, 0xb6, 0xab, 0xa6, 0xd0, 0x6f, 0x0a, 0x7f, 0xd2, 0x35, 0x87, 0x55,
	0xee, 0x6c, 0x59, 0x85, 0x79, 0x1d, 0x5c, 0x1d, 0xc4, 0x60, 0xc3, 0xed, 0xa6, 0x5c, 0x40, 0x6c,
	0x90, 0x48, 0xbf, 0xa2, 0x18, 0x26, 0x2c, 0x9d, 0xa8, 0x9f, 0x91, 0x4b, 0x98, 0x50, 0xf0, 0xa8,
	0xcd, 0xbc, 0x28, 0x08, 0x83, 0x18, 0x25, 0x2c, 0x1d, 0xaa, 0x9d, 0x4c, 0xce, 0xda, 0x17, 0x51,
	0x87, 0x1e, 0x74, 0xe9, 0x0b, 0x0d, 0xfb, 0x77, 0xba, 0x78, 0xd3, 0xd7, 0xbf, 0x42, 0xf9, 0x09,
	0xc4, 0xeb, 0x0b, 0xea, 0x25, 0xf2, 0x38, 0xdb, 0xc8, 0x99, 0xb6, 0x07, 0x92, 0xbd, 0x94, 0x9d,
	0x33, 0x7e, 0x06, 0xb0, 0x95, 0xcc, 0xa3, 0xac, 0x51, 0x3f, 0x3d, 0xcc, 0xba, 0xe6, 0x65, 0xef,
	0x66, 0xf8, 0xd2, 0x77, 0x8b, 0x45, 0xd4, 0xfc, 0xd3, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x33, 0x65, 0xb3, 0xc7, 0xbb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KafkaAmbassadorClient is the client API for KafkaAmbassador service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KafkaAmbassadorClient interface {
	Produce(ctx context.Context, opts ...grpc.CallOption) (KafkaAmbassador_ProduceClient, error)
	ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTopicsResponse, error)
}

type kafkaAmbassadorClient struct {
	cc *grpc.ClientConn
}

func NewKafkaAmbassadorClient(cc *grpc.ClientConn) KafkaAmbassadorClient {
	return &kafkaAmbassadorClient{cc}
}

func (c *kafkaAmbassadorClient) Produce(ctx context.Context, opts ...grpc.CallOption) (KafkaAmbassador_ProduceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KafkaAmbassador_serviceDesc.Streams[0], "/KafkaAmbassador/Produce", opts...)
	if err != nil {
		return nil, err
	}
	x := &kafkaAmbassadorProduceClient{stream}
	return x, nil
}

type KafkaAmbassador_ProduceClient interface {
	Send(*ProdRq) error
	Recv() (*ProdRs, error)
	grpc.ClientStream
}

type kafkaAmbassadorProduceClient struct {
	grpc.ClientStream
}

func (x *kafkaAmbassadorProduceClient) Send(m *ProdRq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kafkaAmbassadorProduceClient) Recv() (*ProdRs, error) {
	m := new(ProdRs)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kafkaAmbassadorClient) ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, "/KafkaAmbassador/ListTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaAmbassadorServer is the server API for KafkaAmbassador service.
type KafkaAmbassadorServer interface {
	Produce(KafkaAmbassador_ProduceServer) error
	ListTopics(context.Context, *Empty) (*ListTopicsResponse, error)
}

// UnimplementedKafkaAmbassadorServer can be embedded to have forward compatible implementations.
type UnimplementedKafkaAmbassadorServer struct {
}

func (*UnimplementedKafkaAmbassadorServer) Produce(srv KafkaAmbassador_ProduceServer) error {
	return status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (*UnimplementedKafkaAmbassadorServer) ListTopics(ctx context.Context, req *Empty) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}

func RegisterKafkaAmbassadorServer(s *grpc.Server, srv KafkaAmbassadorServer) {
	s.RegisterService(&_KafkaAmbassador_serviceDesc, srv)
}

func _KafkaAmbassador_Produce_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KafkaAmbassadorServer).Produce(&kafkaAmbassadorProduceServer{stream})
}

type KafkaAmbassador_ProduceServer interface {
	Send(*ProdRs) error
	Recv() (*ProdRq, error)
	grpc.ServerStream
}

type kafkaAmbassadorProduceServer struct {
	grpc.ServerStream
}

func (x *kafkaAmbassadorProduceServer) Send(m *ProdRs) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kafkaAmbassadorProduceServer) Recv() (*ProdRq, error) {
	m := new(ProdRq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KafkaAmbassador_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaAmbassadorServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KafkaAmbassador/ListTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaAmbassadorServer).ListTopics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _KafkaAmbassador_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KafkaAmbassador",
	HandlerType: (*KafkaAmbassadorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTopics",
			Handler:    _KafkaAmbassador_ListTopics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Produce",
			Handler:       _KafkaAmbassador_Produce_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ambassador.proto",
}
