// Code generated by protoc-gen-go. DO NOT EDIT.
// source: onlineVideoDl.proto

package onlineVideoDl

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Timestamp from public import google/protobuf/timestamp.proto
type Timestamp = timestamp.Timestamp

// 查询有那些任务正在进行
type TasksInProgressRequest struct {
	MaxQueries           int32    `protobuf:"varint,1,opt,name=maxQueries,proto3" json:"maxQueries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TasksInProgressRequest) Reset()         { *m = TasksInProgressRequest{} }
func (m *TasksInProgressRequest) String() string { return proto.CompactTextString(m) }
func (*TasksInProgressRequest) ProtoMessage()    {}
func (*TasksInProgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4cde1866015bf44, []int{0}
}

func (m *TasksInProgressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksInProgressRequest.Unmarshal(m, b)
}
func (m *TasksInProgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksInProgressRequest.Marshal(b, m, deterministic)
}
func (m *TasksInProgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksInProgressRequest.Merge(m, src)
}
func (m *TasksInProgressRequest) XXX_Size() int {
	return xxx_messageInfo_TasksInProgressRequest.Size(m)
}
func (m *TasksInProgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksInProgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TasksInProgressRequest proto.InternalMessageInfo

func (m *TasksInProgressRequest) GetMaxQueries() int32 {
	if m != nil {
		return m.MaxQueries
	}
	return 0
}

type TasksInProgressResponse struct {
	OnlineVideoTasks     []*OnlineVideoTask `protobuf:"bytes,1,rep,name=OnlineVideoTasks,proto3" json:"OnlineVideoTasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TasksInProgressResponse) Reset()         { *m = TasksInProgressResponse{} }
func (m *TasksInProgressResponse) String() string { return proto.CompactTextString(m) }
func (*TasksInProgressResponse) ProtoMessage()    {}
func (*TasksInProgressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4cde1866015bf44, []int{1}
}

func (m *TasksInProgressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksInProgressResponse.Unmarshal(m, b)
}
func (m *TasksInProgressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksInProgressResponse.Marshal(b, m, deterministic)
}
func (m *TasksInProgressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksInProgressResponse.Merge(m, src)
}
func (m *TasksInProgressResponse) XXX_Size() int {
	return xxx_messageInfo_TasksInProgressResponse.Size(m)
}
func (m *TasksInProgressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksInProgressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TasksInProgressResponse proto.InternalMessageInfo

func (m *TasksInProgressResponse) GetOnlineVideoTasks() []*OnlineVideoTask {
	if m != nil {
		return m.OnlineVideoTasks
	}
	return nil
}

type OnlineVideoTask struct {
	StartTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	ID                   int32                `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	URL                  string               `protobuf:"bytes,4,opt,name=URL,proto3" json:"URL,omitempty"`
	SavePath             string               `protobuf:"bytes,5,opt,name=SavePath,proto3" json:"SavePath,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OnlineVideoTask) Reset()         { *m = OnlineVideoTask{} }
func (m *OnlineVideoTask) String() string { return proto.CompactTextString(m) }
func (*OnlineVideoTask) ProtoMessage()    {}
func (*OnlineVideoTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4cde1866015bf44, []int{2}
}

func (m *OnlineVideoTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineVideoTask.Unmarshal(m, b)
}
func (m *OnlineVideoTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineVideoTask.Marshal(b, m, deterministic)
}
func (m *OnlineVideoTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineVideoTask.Merge(m, src)
}
func (m *OnlineVideoTask) XXX_Size() int {
	return xxx_messageInfo_OnlineVideoTask.Size(m)
}
func (m *OnlineVideoTask) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineVideoTask.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineVideoTask proto.InternalMessageInfo

func (m *OnlineVideoTask) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OnlineVideoTask) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *OnlineVideoTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OnlineVideoTask) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *OnlineVideoTask) GetSavePath() string {
	if m != nil {
		return m.SavePath
	}
	return ""
}

func init() {
	proto.RegisterType((*TasksInProgressRequest)(nil), "onlineVideoDl.TasksInProgressRequest")
	proto.RegisterType((*TasksInProgressResponse)(nil), "onlineVideoDl.TasksInProgressResponse")
	proto.RegisterType((*OnlineVideoTask)(nil), "onlineVideoDl.OnlineVideoTask")
}

func init() { proto.RegisterFile("onlineVideoDl.proto", fileDescriptor_c4cde1866015bf44) }

var fileDescriptor_c4cde1866015bf44 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xdf, 0xf2, 0xe7, 0x8d, 0x0c, 0x41, 0xc9, 0x98, 0xe8, 0xa6, 0x07, 0x24, 0x24, 0x1a,
	0x4e, 0x25, 0xc1, 0x0b, 0x1f, 0x80, 0x4b, 0x8d, 0x91, 0xba, 0x54, 0xef, 0x0b, 0x8e, 0x75, 0x63,
	0xdb, 0xad, 0xbb, 0x5b, 0xe3, 0x97, 0xf1, 0xbb, 0x1a, 0xb7, 0x41, 0xa0, 0x98, 0x78, 0x9b, 0x7d,
	0xf6, 0x37, 0xc9, 0xf3, 0x1b, 0x38, 0x55, 0x79, 0x2a, 0x73, 0x7a, 0x94, 0x4f, 0xa4, 0xe6, 0x69,
	0x50, 0x68, 0x65, 0x15, 0xf6, 0xf6, 0x42, 0xff, 0x22, 0x51, 0x2a, 0x49, 0x69, 0xe2, 0x3e, 0x57,
	0xe5, 0xf3, 0xc4, 0xca, 0x8c, 0x8c, 0x15, 0x59, 0x51, 0xf1, 0xa3, 0x19, 0x9c, 0xc5, 0xc2, 0xbc,
	0x9a, 0x30, 0x8f, 0xb4, 0x4a, 0x34, 0x19, 0xc3, 0xe9, 0xad, 0x24, 0x63, 0x71, 0x00, 0x90, 0x89,
	0x8f, 0xfb, 0x92, 0xb4, 0x24, 0xc3, 0xbc, 0xa1, 0x37, 0x6e, 0xf3, 0x9d, 0x64, 0x44, 0x70, 0x7e,
	0xb0, 0x69, 0x0a, 0x95, 0x1b, 0xc2, 0x1b, 0xe8, 0x2f, 0xb6, 0x35, 0x1c, 0xc5, 0xbc, 0x61, 0x73,
	0xdc, 0x9d, 0x0e, 0x82, 0xfd, 0xd2, 0x35, 0x8c, 0x1f, 0xec, 0x8d, 0x3e, 0x3d, 0x38, 0xa9, 0x85,
	0x38, 0x83, 0xce, 0xd2, 0x0a, 0x6d, 0x63, 0x99, 0x91, 0x6b, 0xd6, 0x9d, 0xfa, 0x41, 0x65, 0x1a,
	0x6c, 0x4c, 0x83, 0x78, 0x63, 0xca, 0xb7, 0x30, 0x1e, 0x43, 0x23, 0x9c, 0xb3, 0x86, 0x93, 0x69,
	0x84, 0x73, 0x44, 0x68, 0xdd, 0x89, 0x8c, 0x58, 0x73, 0xe8, 0x8d, 0x3b, 0xdc, 0xcd, 0xd8, 0x87,
	0xe6, 0x03, 0xbf, 0x65, 0x2d, 0x17, 0x7d, 0x8f, 0xe8, 0xc3, 0xd1, 0x52, 0xbc, 0x53, 0x24, 0xec,
	0x0b, 0x6b, 0xbb, 0xf8, 0xe7, 0x3d, 0xb5, 0xd0, 0x5b, 0xec, 0x2a, 0xe1, 0x1a, 0x30, 0x21, 0x1b,
	0x69, 0xb5, 0x26, 0x63, 0x64, 0x9e, 0x38, 0x0d, 0xbc, 0xac, 0x89, 0xff, 0x7e, 0x74, 0xff, 0xea,
	0x2f, 0xac, 0xba, 0x70, 0xf4, 0x6f, 0xf5, 0xdf, 0x89, 0x5e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0x51, 0x45, 0x5f, 0x06, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OnlineVideoDlClient is the client API for OnlineVideoDl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnlineVideoDlClient interface {
	GetProcessingTasks(ctx context.Context, in *TasksInProgressRequest, opts ...grpc.CallOption) (*TasksInProgressResponse, error)
}

type onlineVideoDlClient struct {
	cc *grpc.ClientConn
}

func NewOnlineVideoDlClient(cc *grpc.ClientConn) OnlineVideoDlClient {
	return &onlineVideoDlClient{cc}
}

func (c *onlineVideoDlClient) GetProcessingTasks(ctx context.Context, in *TasksInProgressRequest, opts ...grpc.CallOption) (*TasksInProgressResponse, error) {
	out := new(TasksInProgressResponse)
	err := c.cc.Invoke(ctx, "/onlineVideoDl.OnlineVideoDl/getProcessingTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnlineVideoDlServer is the server API for OnlineVideoDl service.
type OnlineVideoDlServer interface {
	GetProcessingTasks(context.Context, *TasksInProgressRequest) (*TasksInProgressResponse, error)
}

// UnimplementedOnlineVideoDlServer can be embedded to have forward compatible implementations.
type UnimplementedOnlineVideoDlServer struct {
}

func (*UnimplementedOnlineVideoDlServer) GetProcessingTasks(ctx context.Context, req *TasksInProgressRequest) (*TasksInProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProcessingTasks not implemented")
}

func RegisterOnlineVideoDlServer(s *grpc.Server, srv OnlineVideoDlServer) {
	s.RegisterService(&_OnlineVideoDl_serviceDesc, srv)
}

func _OnlineVideoDl_GetProcessingTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TasksInProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineVideoDlServer).GetProcessingTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onlineVideoDl.OnlineVideoDl/GetProcessingTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineVideoDlServer).GetProcessingTasks(ctx, req.(*TasksInProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnlineVideoDl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onlineVideoDl.OnlineVideoDl",
	HandlerType: (*OnlineVideoDlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getProcessingTasks",
			Handler:    _OnlineVideoDl_GetProcessingTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onlineVideoDl.proto",
}