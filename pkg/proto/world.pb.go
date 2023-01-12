// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: world.proto

package proto

import (
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

type World struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the world
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the world
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *World) Reset() {
	*x = World{}
	if protoimpl.UnsafeEnabled {
		mi := &file_world_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *World) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*World) ProtoMessage() {}

func (x *World) ProtoReflect() protoreflect.Message {
	mi := &file_world_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use World.ProtoReflect.Descriptor instead.
func (*World) Descriptor() ([]byte, []int) {
	return file_world_proto_rawDescGZIP(), []int{0}
}

func (x *World) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *World) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the world to be returned
	// If this field is not provided, all created worlds will be returned
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WorldRequest) Reset() {
	*x = WorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_world_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRequest) ProtoMessage() {}

func (x *WorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_world_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRequest.ProtoReflect.Descriptor instead.
func (*WorldRequest) Descriptor() ([]byte, []int) {
	return file_world_proto_rawDescGZIP(), []int{1}
}

func (x *WorldRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type WorldSelectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the world to be selected
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WorldSelectionRequest) Reset() {
	*x = WorldSelectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_world_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldSelectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldSelectionRequest) ProtoMessage() {}

func (x *WorldSelectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_world_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldSelectionRequest.ProtoReflect.Descriptor instead.
func (*WorldSelectionRequest) Descriptor() ([]byte, []int) {
	return file_world_proto_rawDescGZIP(), []int{2}
}

func (x *WorldSelectionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type WorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The selected or retrieved world
	World *World `protobuf:"bytes,1,opt,name=world,proto3" json:"world,omitempty"`
	// If an error occurs, this field will be set with the error message
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *WorldResponse) Reset() {
	*x = WorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_world_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldResponse) ProtoMessage() {}

func (x *WorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_world_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldResponse.ProtoReflect.Descriptor instead.
func (*WorldResponse) Descriptor() ([]byte, []int) {
	return file_world_proto_rawDescGZIP(), []int{3}
}

func (x *WorldResponse) GetWorld() *World {
	if x != nil {
		return x.World
	}
	return nil
}

func (x *WorldResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_world_proto protoreflect.FileDescriptor

var file_world_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a,
	0x05, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x05, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x9d, 0x01, 0x0a, 0x0c, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x06, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x1a, 0x0e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x12, 0x16, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x0d, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x6b, 0x65, 0x67, 0x72, 0x69, 0x66, 0x66,
	0x69, 0x74, 0x68, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_world_proto_rawDescOnce sync.Once
	file_world_proto_rawDescData = file_world_proto_rawDesc
)

func file_world_proto_rawDescGZIP() []byte {
	file_world_proto_rawDescOnce.Do(func() {
		file_world_proto_rawDescData = protoimpl.X.CompressGZIP(file_world_proto_rawDescData)
	})
	return file_world_proto_rawDescData
}

var file_world_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_world_proto_goTypes = []interface{}{
	(*World)(nil),                 // 0: World
	(*WorldRequest)(nil),          // 1: WorldRequest
	(*WorldSelectionRequest)(nil), // 2: WorldSelectionRequest
	(*WorldResponse)(nil),         // 3: WorldResponse
}
var file_world_proto_depIdxs = []int32{
	0, // 0: WorldResponse.world:type_name -> World
	0, // 1: WorldService.CreateWorld:input_type -> World
	2, // 2: WorldService.SelectWorld:input_type -> WorldSelectionRequest
	1, // 3: WorldService.GetWorld:input_type -> WorldRequest
	3, // 4: WorldService.CreateWorld:output_type -> WorldResponse
	3, // 5: WorldService.SelectWorld:output_type -> WorldResponse
	3, // 6: WorldService.GetWorld:output_type -> WorldResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_world_proto_init() }
func file_world_proto_init() {
	if File_world_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_world_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*World); i {
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
		file_world_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRequest); i {
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
		file_world_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldSelectionRequest); i {
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
		file_world_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldResponse); i {
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
			RawDescriptor: file_world_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_world_proto_goTypes,
		DependencyIndexes: file_world_proto_depIdxs,
		MessageInfos:      file_world_proto_msgTypes,
	}.Build()
	File_world_proto = out.File
	file_world_proto_rawDesc = nil
	file_world_proto_goTypes = nil
	file_world_proto_depIdxs = nil
}