// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: myproto.proto

package Mygrpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//客户端发送给服务端
type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myproto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_myproto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_myproto_proto_rawDescGZIP(), []int{0}
}

func (x *HelloReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//服务端返回给客户端
type HelloRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloRsp) Reset() {
	*x = HelloRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myproto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRsp) ProtoMessage() {}

func (x *HelloRsp) ProtoReflect() protoreflect.Message {
	mi := &file_myproto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRsp.ProtoReflect.Descriptor instead.
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return file_myproto_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

//客户端发送给服务端
type NameRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NameRep) Reset() {
	*x = NameRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myproto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameRep) ProtoMessage() {}

func (x *NameRep) ProtoReflect() protoreflect.Message {
	mi := &file_myproto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameRep.ProtoReflect.Descriptor instead.
func (*NameRep) Descriptor() ([]byte, []int) {
	return file_myproto_proto_rawDescGZIP(), []int{2}
}

func (x *NameRep) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//服务端返回给客户端
type NameRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NameRsp) Reset() {
	*x = NameRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myproto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameRsp) ProtoMessage() {}

func (x *NameRsp) ProtoReflect() protoreflect.Message {
	mi := &file_myproto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameRsp.ProtoReflect.Descriptor instead.
func (*NameRsp) Descriptor() ([]byte, []int) {
	return file_myproto_proto_rawDescGZIP(), []int{3}
}

func (x *NameRsp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_myproto_proto protoreflect.FileDescriptor

var file_myproto_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x4d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x22, 0x1e, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1d, 0x0a, 0x07, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0x6e, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x10,
	0x2e, 0x4d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x4d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x07, 0x53, 0x61, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0f, 0x2e, 0x4d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70,
	0x1a, 0x0f, 0x2e, 0x4d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x73,
	0x70, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_myproto_proto_rawDescOnce sync.Once
	file_myproto_proto_rawDescData = file_myproto_proto_rawDesc
)

func file_myproto_proto_rawDescGZIP() []byte {
	file_myproto_proto_rawDescOnce.Do(func() {
		file_myproto_proto_rawDescData = protoimpl.X.CompressGZIP(file_myproto_proto_rawDescData)
	})
	return file_myproto_proto_rawDescData
}

var file_myproto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_myproto_proto_goTypes = []interface{}{
	(*HelloReq)(nil), // 0: Mygrpc.HelloReq
	(*HelloRsp)(nil), // 1: Mygrpc.HelloRsp
	(*NameRep)(nil),  // 2: Mygrpc.NameRep
	(*NameRsp)(nil),  // 3: Mygrpc.NameRsp
}
var file_myproto_proto_depIdxs = []int32{
	0, // 0: Mygrpc.Helloserver.Sayhello:input_type -> Mygrpc.HelloReq
	2, // 1: Mygrpc.Helloserver.Sayname:input_type -> Mygrpc.NameRep
	1, // 2: Mygrpc.Helloserver.Sayhello:output_type -> Mygrpc.HelloRsp
	3, // 3: Mygrpc.Helloserver.Sayname:output_type -> Mygrpc.NameRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_myproto_proto_init() }
func file_myproto_proto_init() {
	if File_myproto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_myproto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_myproto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_myproto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameRep); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_myproto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_myproto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_myproto_proto_goTypes,
		DependencyIndexes: file_myproto_proto_depIdxs,
		MessageInfos:      file_myproto_proto_msgTypes,
	}.Build()
	File_myproto_proto = out.File
	file_myproto_proto_rawDesc = nil
	file_myproto_proto_goTypes = nil
	file_myproto_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloserverClient is the client API for Helloserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloserverClient interface {
	//一个打招呼的函书服务
	Sayhello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error)
	//一个说名字的服务
	Sayname(ctx context.Context, in *NameRep, opts ...grpc.CallOption) (*NameRsp, error)
}

type helloserverClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloserverClient(cc grpc.ClientConnInterface) HelloserverClient {
	return &helloserverClient{cc}
}

func (c *helloserverClient) Sayhello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error) {
	out := new(HelloRsp)
	err := c.cc.Invoke(ctx, "/Mygrpc.Helloserver/Sayhello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloserverClient) Sayname(ctx context.Context, in *NameRep, opts ...grpc.CallOption) (*NameRsp, error) {
	out := new(NameRsp)
	err := c.cc.Invoke(ctx, "/Mygrpc.Helloserver/Sayname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloserverServer is the server API for Helloserver service.
type HelloserverServer interface {
	//一个打招呼的函书服务
	Sayhello(context.Context, *HelloReq) (*HelloRsp, error)
	//一个说名字的服务
	Sayname(context.Context, *NameRep) (*NameRsp, error)
}

// UnimplementedHelloserverServer can be embedded to have forward compatible implementations.
type UnimplementedHelloserverServer struct {
}

func (*UnimplementedHelloserverServer) Sayhello(context.Context, *HelloReq) (*HelloRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sayhello not implemented")
}
func (*UnimplementedHelloserverServer) Sayname(context.Context, *NameRep) (*NameRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sayname not implemented")
}

func RegisterHelloserverServer(s *grpc.Server, srv HelloserverServer) {
	s.RegisterService(&_Helloserver_serviceDesc, srv)
}

func _Helloserver_Sayhello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloserverServer).Sayhello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mygrpc.Helloserver/Sayhello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloserverServer).Sayhello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloserver_Sayname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloserverServer).Sayname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mygrpc.Helloserver/Sayname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloserverServer).Sayname(ctx, req.(*NameRep))
	}
	return interceptor(ctx, in, info, handler)
}

var _Helloserver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Mygrpc.Helloserver",
	HandlerType: (*HelloserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sayhello",
			Handler:    _Helloserver_Sayhello_Handler,
		},
		{
			MethodName: "Sayname",
			Handler:    _Helloserver_Sayname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myproto.proto",
}
