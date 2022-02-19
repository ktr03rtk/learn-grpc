// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: matching.proto

package pb

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

type JoinRoomResponse_Status int32

const (
	JoinRoomResponse_UNKNOWN JoinRoomResponse_Status = 0
	JoinRoomResponse_WAITING JoinRoomResponse_Status = 1
	JoinRoomResponse_MATCHED JoinRoomResponse_Status = 2
)

// Enum value maps for JoinRoomResponse_Status.
var (
	JoinRoomResponse_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "WAITING",
		2: "MATCHED",
	}
	JoinRoomResponse_Status_value = map[string]int32{
		"UNKNOWN": 0,
		"WAITING": 1,
		"MATCHED": 2,
	}
)

func (x JoinRoomResponse_Status) Enum() *JoinRoomResponse_Status {
	p := new(JoinRoomResponse_Status)
	*p = x
	return p
}

func (x JoinRoomResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JoinRoomResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_matching_proto_enumTypes[0].Descriptor()
}

func (JoinRoomResponse_Status) Type() protoreflect.EnumType {
	return &file_matching_proto_enumTypes[0]
}

func (x JoinRoomResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JoinRoomResponse_Status.Descriptor instead.
func (JoinRoomResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_matching_proto_rawDescGZIP(), []int{1, 0}
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matching_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matching_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_matching_proto_rawDescGZIP(), []int{0}
}

type JoinRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room   *Room                   `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Me     *Player                 `protobuf:"bytes,2,opt,name=me,proto3" json:"me,omitempty"`
	Status JoinRoomResponse_Status `protobuf:"varint,3,opt,name=status,proto3,enum=game.JoinRoomResponse_Status" json:"status,omitempty"`
}

func (x *JoinRoomResponse) Reset() {
	*x = JoinRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matching_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomResponse) ProtoMessage() {}

func (x *JoinRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_matching_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomResponse.ProtoReflect.Descriptor instead.
func (*JoinRoomResponse) Descriptor() ([]byte, []int) {
	return file_matching_proto_rawDescGZIP(), []int{1}
}

func (x *JoinRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *JoinRoomResponse) GetMe() *Player {
	if x != nil {
		return x.Me
	}
	return nil
}

func (x *JoinRoomResponse) GetStatus() JoinRoomResponse_Status {
	if x != nil {
		return x.Status
	}
	return JoinRoomResponse_UNKNOWN
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Host  *Player `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Guest *Player `protobuf:"bytes,3,opt,name=guest,proto3" json:"guest,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matching_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_matching_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_matching_proto_rawDescGZIP(), []int{2}
}

func (x *Room) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Room) GetHost() *Player {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *Room) GetGuest() *Player {
	if x != nil {
		return x.Guest
	}
	return nil
}

var File_matching_proto protoreflect.FileDescriptor

var file_matching_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x1a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x02,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x02, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x2f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x45, 0x44,
	0x10, 0x02, 0x22, 0x5c, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05,
	0x67, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x05, 0x67, 0x75, 0x65, 0x73, 0x74,
	0x32, 0x50, 0x0a, 0x0f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12,
	0x15, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_matching_proto_rawDescOnce sync.Once
	file_matching_proto_rawDescData = file_matching_proto_rawDesc
)

func file_matching_proto_rawDescGZIP() []byte {
	file_matching_proto_rawDescOnce.Do(func() {
		file_matching_proto_rawDescData = protoimpl.X.CompressGZIP(file_matching_proto_rawDescData)
	})
	return file_matching_proto_rawDescData
}

var file_matching_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_matching_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_matching_proto_goTypes = []interface{}{
	(JoinRoomResponse_Status)(0), // 0: game.JoinRoomResponse.Status
	(*JoinRoomRequest)(nil),      // 1: game.JoinRoomRequest
	(*JoinRoomResponse)(nil),     // 2: game.JoinRoomResponse
	(*Room)(nil),                 // 3: game.Room
	(*Player)(nil),               // 4: game.Player
}
var file_matching_proto_depIdxs = []int32{
	3, // 0: game.JoinRoomResponse.room:type_name -> game.Room
	4, // 1: game.JoinRoomResponse.me:type_name -> game.Player
	0, // 2: game.JoinRoomResponse.status:type_name -> game.JoinRoomResponse.Status
	4, // 3: game.Room.host:type_name -> game.Player
	4, // 4: game.Room.guest:type_name -> game.Player
	1, // 5: game.MatchingService.JoinRoom:input_type -> game.JoinRoomRequest
	2, // 6: game.MatchingService.JoinRoom:output_type -> game.JoinRoomResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_matching_proto_init() }
func file_matching_proto_init() {
	if File_matching_proto != nil {
		return
	}
	file_player_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_matching_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomRequest); i {
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
		file_matching_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomResponse); i {
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
		file_matching_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
			RawDescriptor: file_matching_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_matching_proto_goTypes,
		DependencyIndexes: file_matching_proto_depIdxs,
		EnumInfos:         file_matching_proto_enumTypes,
		MessageInfos:      file_matching_proto_msgTypes,
	}.Build()
	File_matching_proto = out.File
	file_matching_proto_rawDesc = nil
	file_matching_proto_goTypes = nil
	file_matching_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MatchingServiceClient is the client API for MatchingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MatchingServiceClient interface {
	JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (MatchingService_JoinRoomClient, error)
}

type matchingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchingServiceClient(cc grpc.ClientConnInterface) MatchingServiceClient {
	return &matchingServiceClient{cc}
}

func (c *matchingServiceClient) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (MatchingService_JoinRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MatchingService_serviceDesc.Streams[0], "/game.MatchingService/JoinRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &matchingServiceJoinRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MatchingService_JoinRoomClient interface {
	Recv() (*JoinRoomResponse, error)
	grpc.ClientStream
}

type matchingServiceJoinRoomClient struct {
	grpc.ClientStream
}

func (x *matchingServiceJoinRoomClient) Recv() (*JoinRoomResponse, error) {
	m := new(JoinRoomResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MatchingServiceServer is the server API for MatchingService service.
type MatchingServiceServer interface {
	JoinRoom(*JoinRoomRequest, MatchingService_JoinRoomServer) error
}

// UnimplementedMatchingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMatchingServiceServer struct {
}

func (*UnimplementedMatchingServiceServer) JoinRoom(*JoinRoomRequest, MatchingService_JoinRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}

func RegisterMatchingServiceServer(s *grpc.Server, srv MatchingServiceServer) {
	s.RegisterService(&_MatchingService_serviceDesc, srv)
}

func _MatchingService_JoinRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinRoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatchingServiceServer).JoinRoom(m, &matchingServiceJoinRoomServer{stream})
}

type MatchingService_JoinRoomServer interface {
	Send(*JoinRoomResponse) error
	grpc.ServerStream
}

type matchingServiceJoinRoomServer struct {
	grpc.ServerStream
}

func (x *matchingServiceJoinRoomServer) Send(m *JoinRoomResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MatchingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game.MatchingService",
	HandlerType: (*MatchingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinRoom",
			Handler:       _MatchingService_JoinRoom_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "matching.proto",
}
