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

type RunRequest struct {
	DeploymentOwner      string            `protobuf:"bytes,1,opt,name=deploymentOwner,proto3" json:"deploymentOwner,omitempty"`
	DeploymentName       string            `protobuf:"bytes,2,opt,name=deploymentName,proto3" json:"deploymentName,omitempty"`
	EndpointPath         string            `protobuf:"bytes,3,opt,name=endpointPath,proto3" json:"endpointPath,omitempty"`
	TraceID              string            `protobuf:"bytes,5,opt,name=traceID,proto3" json:"traceID,omitempty"`
	ContentType          string            `protobuf:"bytes,6,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Message              []byte            `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	Parameters           map[string]string `protobuf:"bytes,8,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StreamOffset         uint64            `protobuf:"varint,9,opt,name=streamOffset,proto3" json:"streamOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19084e700d1da46, []int{2}
}

func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (m *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(m, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

func (m *RunRequest) GetDeploymentOwner() string {
	if m != nil {
		return m.DeploymentOwner
	}
	return ""
}

func (m *RunRequest) GetDeploymentName() string {
	if m != nil {
		return m.DeploymentName
	}
	return ""
}

func (m *RunRequest) GetEndpointPath() string {
	if m != nil {
		return m.EndpointPath
	}
	return ""
}

func (m *RunRequest) GetTraceID() string {
	if m != nil {
		return m.TraceID
	}
	return ""
}

func (m *RunRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *RunRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *RunRequest) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *RunRequest) GetStreamOffset() uint64 {
	if m != nil {
		return m.StreamOffset
	}
	return 0
}

type RunResponse struct {
	StreamOffset         uint64   `protobuf:"varint,3,opt,name=streamOffset,proto3" json:"streamOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunResponse) Reset()         { *m = RunResponse{} }
func (m *RunResponse) String() string { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()    {}
func (*RunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19084e700d1da46, []int{3}
}

func (m *RunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunResponse.Unmarshal(m, b)
}
func (m *RunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunResponse.Marshal(b, m, deterministic)
}
func (m *RunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunResponse.Merge(m, src)
}
func (m *RunResponse) XXX_Size() int {
	return xxx_messageInfo_RunResponse.Size(m)
}
func (m *RunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunResponse proto.InternalMessageInfo

func (m *RunResponse) GetStreamOffset() uint64 {
	if m != nil {
		return m.StreamOffset
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*ListTopicsResponse)(nil), "ListTopicsResponse")
	proto.RegisterType((*RunRequest)(nil), "RunRequest")
	proto.RegisterMapType((map[string]string)(nil), "RunRequest.ParametersEntry")
	proto.RegisterType((*RunResponse)(nil), "RunResponse")
}

func init() { proto.RegisterFile("ambassador.proto", fileDescriptor_c19084e700d1da46) }

var fileDescriptor_c19084e700d1da46 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8f, 0xd3, 0x30,
	0x10, 0x85, 0xd7, 0xc9, 0x36, 0xa5, 0x93, 0x8a, 0xad, 0x0c, 0x42, 0x56, 0xb9, 0x44, 0x39, 0xac,
	0x82, 0x40, 0x11, 0x94, 0x0b, 0x02, 0x71, 0x41, 0xf4, 0x80, 0x84, 0x68, 0x65, 0x95, 0x0b, 0x37,
	0xb7, 0x99, 0x42, 0x44, 0x62, 0x1b, 0xdb, 0x01, 0xe5, 0x6f, 0xf0, 0x6b, 0x39, 0xa2, 0x9a, 0x94,
	0xa6, 0xe9, 0x61, 0x6f, 0x79, 0xdf, 0xbc, 0xd8, 0x33, 0xe3, 0x07, 0x33, 0x51, 0x6f, 0x85, 0xb5,
	0xa2, 0x50, 0x26, 0xd7, 0x46, 0x39, 0x95, 0x8e, 0x61, 0xb4, 0xac, 0xb5, 0x6b, 0xd3, 0x67, 0x40,
	0x3f, 0x96, 0xd6, 0x6d, 0x94, 0x2e, 0x77, 0x96, 0xa3, 0xd5, 0x4a, 0x5a, 0xa4, 0x8f, 0x20, 0x72,
	0x9e, 0x30, 0x92, 0x84, 0xd9, 0x84, 0x77, 0x2a, 0xfd, 0x13, 0x00, 0xf0, 0x46, 0x72, 0xfc, 0xd1,
	0xa0, 0x75, 0x34, 0x83, 0x9b, 0x02, 0x75, 0xa5, 0xda, 0x1a, 0xa5, 0x5b, 0xfd, 0x92, 0x68, 0x18,
	0x49, 0x48, 0x36, 0xe1, 0x43, 0x4c, 0x6f, 0xe1, 0xfe, 0x09, 0x7d, 0x12, 0x35, 0xb2, 0xc0, 0x1b,
	0x07, 0x94, 0xa6, 0x30, 0x45, 0x59, 0x68, 0x55, 0x4a, 0xb7, 0x16, 0xee, 0x1b, 0x0b, 0xbd, 0xeb,
	0x8c, 0x51, 0x06, 0x63, 0x67, 0xc4, 0x0e, 0x3f, 0xbc, 0x67, 0x23, 0x5f, 0x3e, 0x4a, 0x9a, 0x40,
	0xbc, 0x53, 0xd2, 0xa1, 0x74, 0x9b, 0x56, 0x23, 0x8b, 0x7c, 0xb5, 0x8f, 0x0e, 0xff, 0xd6, 0x68,
	0xad, 0xf8, 0x8a, 0x6c, 0x9c, 0x90, 0x6c, 0xca, 0x8f, 0x92, 0xbe, 0x01, 0xd0, 0xc2, 0x88, 0x1a,
	0x1d, 0x1a, 0xcb, 0xee, 0x25, 0x61, 0x16, 0x2f, 0x1e, 0xe7, 0xa7, 0x61, 0xf3, 0xf5, 0xff, 0xea,
	0x52, 0x3a, 0xd3, 0xf2, 0x9e, 0xfd, 0xd0, 0xb6, 0x75, 0x06, 0x45, 0xbd, 0xda, 0xef, 0x2d, 0x3a,
	0x36, 0x49, 0x48, 0x76, 0xcd, 0xcf, 0xd8, 0xfc, 0x2d, 0xdc, 0x0c, 0x8e, 0xa0, 0x33, 0x08, 0xbf,
	0x63, 0xdb, 0xed, 0xec, 0xf0, 0x49, 0x1f, 0xc2, 0xe8, 0xa7, 0xa8, 0x9a, 0xe3, 0x7a, 0xfe, 0x89,
	0xd7, 0xc1, 0x2b, 0x92, 0xbe, 0x80, 0xd8, 0x37, 0xd3, 0xbd, 0xd0, 0xf0, 0xc6, 0xf0, 0xf2, 0xc6,
	0xc5, 0x6f, 0x02, 0xb3, 0x75, 0xa9, 0xb1, 0x2a, 0x25, 0x2e, 0xbb, 0x0d, 0xd2, 0x5b, 0x08, 0x79,
	0x23, 0x69, 0xdc, 0x1b, 0x6d, 0x3e, 0xcd, 0x7b, 0x47, 0xa7, 0x57, 0x19, 0x79, 0x4e, 0xe8, 0x13,
	0x88, 0x3e, 0xeb, 0x4a, 0x89, 0xe2, 0x6e, 0xeb, 0x53, 0x80, 0x53, 0x86, 0x68, 0x94, 0xfb, 0x64,
	0xcd, 0x1f, 0xe4, 0x97, 0xc1, 0x4a, 0xaf, 0xde, 0x5d, 0x7f, 0x09, 0xf4, 0x76, 0x1b, 0xf9, 0x18,
	0xbe, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x25, 0x40, 0x99, 0xcc, 0x9a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PipelineEndpointClient is the client API for PipelineEndpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PipelineEndpointClient interface {
	Run(ctx context.Context, opts ...grpc.CallOption) (PipelineEndpoint_RunClient, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (PipelineEndpoint_UploadClient, error)
	ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTopicsResponse, error)
}

type pipelineEndpointClient struct {
	cc *grpc.ClientConn
}

func NewPipelineEndpointClient(cc *grpc.ClientConn) PipelineEndpointClient {
	return &pipelineEndpointClient{cc}
}

func (c *pipelineEndpointClient) Run(ctx context.Context, opts ...grpc.CallOption) (PipelineEndpoint_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PipelineEndpoint_serviceDesc.Streams[0], "/PipelineEndpoint/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineEndpointRunClient{stream}
	return x, nil
}

type PipelineEndpoint_RunClient interface {
	Send(*RunRequest) error
	Recv() (*RunResponse, error)
	grpc.ClientStream
}

type pipelineEndpointRunClient struct {
	grpc.ClientStream
}

func (x *pipelineEndpointRunClient) Send(m *RunRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pipelineEndpointRunClient) Recv() (*RunResponse, error) {
	m := new(RunResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pipelineEndpointClient) Upload(ctx context.Context, opts ...grpc.CallOption) (PipelineEndpoint_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PipelineEndpoint_serviceDesc.Streams[1], "/PipelineEndpoint/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineEndpointUploadClient{stream}
	return x, nil
}

type PipelineEndpoint_UploadClient interface {
	Send(*RunRequest) error
	Recv() (*RunResponse, error)
	grpc.ClientStream
}

type pipelineEndpointUploadClient struct {
	grpc.ClientStream
}

func (x *pipelineEndpointUploadClient) Send(m *RunRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pipelineEndpointUploadClient) Recv() (*RunResponse, error) {
	m := new(RunResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pipelineEndpointClient) ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, "/PipelineEndpoint/ListTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineEndpointServer is the server API for PipelineEndpoint service.
type PipelineEndpointServer interface {
	Run(PipelineEndpoint_RunServer) error
	Upload(PipelineEndpoint_UploadServer) error
	ListTopics(context.Context, *Empty) (*ListTopicsResponse, error)
}

// UnimplementedPipelineEndpointServer can be embedded to have forward compatible implementations.
type UnimplementedPipelineEndpointServer struct {
}

func (*UnimplementedPipelineEndpointServer) Run(srv PipelineEndpoint_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (*UnimplementedPipelineEndpointServer) Upload(srv PipelineEndpoint_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedPipelineEndpointServer) ListTopics(ctx context.Context, req *Empty) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}

func RegisterPipelineEndpointServer(s *grpc.Server, srv PipelineEndpointServer) {
	s.RegisterService(&_PipelineEndpoint_serviceDesc, srv)
}

func _PipelineEndpoint_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PipelineEndpointServer).Run(&pipelineEndpointRunServer{stream})
}

type PipelineEndpoint_RunServer interface {
	Send(*RunResponse) error
	Recv() (*RunRequest, error)
	grpc.ServerStream
}

type pipelineEndpointRunServer struct {
	grpc.ServerStream
}

func (x *pipelineEndpointRunServer) Send(m *RunResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pipelineEndpointRunServer) Recv() (*RunRequest, error) {
	m := new(RunRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PipelineEndpoint_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PipelineEndpointServer).Upload(&pipelineEndpointUploadServer{stream})
}

type PipelineEndpoint_UploadServer interface {
	Send(*RunResponse) error
	Recv() (*RunRequest, error)
	grpc.ServerStream
}

type pipelineEndpointUploadServer struct {
	grpc.ServerStream
}

func (x *pipelineEndpointUploadServer) Send(m *RunResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pipelineEndpointUploadServer) Recv() (*RunRequest, error) {
	m := new(RunRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PipelineEndpoint_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineEndpointServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PipelineEndpoint/ListTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineEndpointServer).ListTopics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PipelineEndpoint_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PipelineEndpoint",
	HandlerType: (*PipelineEndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTopics",
			Handler:    _PipelineEndpoint_ListTopics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _PipelineEndpoint_Run_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Upload",
			Handler:       _PipelineEndpoint_Upload_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ambassador.proto",
}
