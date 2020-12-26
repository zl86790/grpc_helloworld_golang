// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: Hi.proto

package services

import (
	context "context"
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

type HiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Say string `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
}

func (x *HiRequest) Reset() {
	*x = HiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Hi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiRequest) ProtoMessage() {}

func (x *HiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Hi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiRequest.ProtoReflect.Descriptor instead.
func (*HiRequest) Descriptor() ([]byte, []int) {
	return file_Hi_proto_rawDescGZIP(), []int{0}
}

func (x *HiRequest) GetSay() string {
	if x != nil {
		return x.Say
	}
	return ""
}

type HiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responed string `protobuf:"bytes,1,opt,name=responed,proto3" json:"responed,omitempty"`
}

func (x *HiResponse) Reset() {
	*x = HiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Hi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiResponse) ProtoMessage() {}

func (x *HiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Hi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiResponse.ProtoReflect.Descriptor instead.
func (*HiResponse) Descriptor() ([]byte, []int) {
	return file_Hi_proto_rawDescGZIP(), []int{1}
}

func (x *HiResponse) GetResponed() string {
	if x != nil {
		return x.Responed
	}
	return ""
}

var File_Hi_proto protoreflect.FileDescriptor

var file_Hi_proto_rawDesc = []byte{
	0x0a, 0x08, 0x48, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x1d, 0x0a, 0x09, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x61, 0x79, 0x22, 0x28, 0x0a, 0x0a, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x64, 0x32, 0x47, 0x0a,
	0x09, 0x48, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Hi_proto_rawDescOnce sync.Once
	file_Hi_proto_rawDescData = file_Hi_proto_rawDesc
)

func file_Hi_proto_rawDescGZIP() []byte {
	file_Hi_proto_rawDescOnce.Do(func() {
		file_Hi_proto_rawDescData = protoimpl.X.CompressGZIP(file_Hi_proto_rawDescData)
	})
	return file_Hi_proto_rawDescData
}

var file_Hi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Hi_proto_goTypes = []interface{}{
	(*HiRequest)(nil),  // 0: services.HiRequest
	(*HiResponse)(nil), // 1: services.HiResponse
}
var file_Hi_proto_depIdxs = []int32{
	0, // 0: services.HiService.GetHiResponed:input_type -> services.HiRequest
	1, // 1: services.HiService.GetHiResponed:output_type -> services.HiResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Hi_proto_init() }
func file_Hi_proto_init() {
	if File_Hi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Hi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiRequest); i {
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
		file_Hi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiResponse); i {
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
			RawDescriptor: file_Hi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Hi_proto_goTypes,
		DependencyIndexes: file_Hi_proto_depIdxs,
		MessageInfos:      file_Hi_proto_msgTypes,
	}.Build()
	File_Hi_proto = out.File
	file_Hi_proto_rawDesc = nil
	file_Hi_proto_goTypes = nil
	file_Hi_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HiServiceClient is the client API for HiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HiServiceClient interface {
	GetHiResponed(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error)
}

type hiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHiServiceClient(cc grpc.ClientConnInterface) HiServiceClient {
	return &hiServiceClient{cc}
}

func (c *hiServiceClient) GetHiResponed(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error) {
	out := new(HiResponse)
	err := c.cc.Invoke(ctx, "/services.HiService/GetHiResponed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HiServiceServer is the server API for HiService service.
type HiServiceServer interface {
	GetHiResponed(context.Context, *HiRequest) (*HiResponse, error)
}

// UnimplementedHiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHiServiceServer struct {
}

func (*UnimplementedHiServiceServer) GetHiResponed(context.Context, *HiRequest) (*HiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHiResponed not implemented")
}

func RegisterHiServiceServer(s *grpc.Server, srv HiServiceServer) {
	s.RegisterService(&_HiService_serviceDesc, srv)
}

func _HiService_GetHiResponed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HiServiceServer).GetHiResponed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.HiService/GetHiResponed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HiServiceServer).GetHiResponed(ctx, req.(*HiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.HiService",
	HandlerType: (*HiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHiResponed",
			Handler:    _HiService_GetHiResponed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Hi.proto",
}
