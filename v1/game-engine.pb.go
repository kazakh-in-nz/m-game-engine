// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: game-engine.proto

package gameengine

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

type GetSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSizeRequest) Reset() {
	*x = GetSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSizeRequest) ProtoMessage() {}

func (x *GetSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSizeRequest.ProtoReflect.Descriptor instead.
func (*GetSizeRequest) Descriptor() ([]byte, []int) {
	return file_game_engine_proto_rawDescGZIP(), []int{0}
}

type GetSizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size float64 `protobuf:"fixed64,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *GetSizeResponse) Reset() {
	*x = GetSizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSizeResponse) ProtoMessage() {}

func (x *GetSizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSizeResponse.ProtoReflect.Descriptor instead.
func (*GetSizeResponse) Descriptor() ([]byte, []int) {
	return file_game_engine_proto_rawDescGZIP(), []int{1}
}

func (x *GetSizeResponse) GetSize() float64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SetScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score float64 `protobuf:"fixed64,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *SetScoreRequest) Reset() {
	*x = SetScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScoreRequest) ProtoMessage() {}

func (x *SetScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScoreRequest.ProtoReflect.Descriptor instead.
func (*SetScoreRequest) Descriptor() ([]byte, []int) {
	return file_game_engine_proto_rawDescGZIP(), []int{2}
}

func (x *SetScoreRequest) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type SetScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set bool `protobuf:"varint,1,opt,name=set,proto3" json:"set,omitempty"` // boolean will let knwo if setting the score was successful
}

func (x *SetScoreResponse) Reset() {
	*x = SetScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScoreResponse) ProtoMessage() {}

func (x *SetScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScoreResponse.ProtoReflect.Descriptor instead.
func (*SetScoreResponse) Descriptor() ([]byte, []int) {
	return file_game_engine_proto_rawDescGZIP(), []int{3}
}

func (x *SetScoreResponse) GetSet() bool {
	if x != nil {
		return x.Set
	}
	return false
}

var File_game_engine_proto protoreflect.FileDescriptor

var file_game_engine_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x27,
	0x0a, 0x0f, 0x53, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x24, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x65, 0x74, 0x32, 0xaf, 0x01,
	0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x4e, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x2e, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x73, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x08,
	0x53, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x6d, 0x73, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x73,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x67, 0x61, 0x6d, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_engine_proto_rawDescOnce sync.Once
	file_game_engine_proto_rawDescData = file_game_engine_proto_rawDesc
)

func file_game_engine_proto_rawDescGZIP() []byte {
	file_game_engine_proto_rawDescOnce.Do(func() {
		file_game_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_engine_proto_rawDescData)
	})
	return file_game_engine_proto_rawDescData
}

var file_game_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_game_engine_proto_goTypes = []interface{}{
	(*GetSizeRequest)(nil),   // 0: ms.gameengine.v1.GetSizeRequest
	(*GetSizeResponse)(nil),  // 1: ms.gameengine.v1.GetSizeResponse
	(*SetScoreRequest)(nil),  // 2: ms.gameengine.v1.SetScoreRequest
	(*SetScoreResponse)(nil), // 3: ms.gameengine.v1.SetScoreResponse
}
var file_game_engine_proto_depIdxs = []int32{
	0, // 0: ms.gameengine.v1.GameEngine.GetSize:input_type -> ms.gameengine.v1.GetSizeRequest
	2, // 1: ms.gameengine.v1.GameEngine.SetScore:input_type -> ms.gameengine.v1.SetScoreRequest
	1, // 2: ms.gameengine.v1.GameEngine.GetSize:output_type -> ms.gameengine.v1.GetSizeResponse
	3, // 3: ms.gameengine.v1.GameEngine.SetScore:output_type -> ms.gameengine.v1.SetScoreResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_engine_proto_init() }
func file_game_engine_proto_init() {
	if File_game_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSizeRequest); i {
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
		file_game_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSizeResponse); i {
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
		file_game_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScoreRequest); i {
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
		file_game_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScoreResponse); i {
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
			RawDescriptor: file_game_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_engine_proto_goTypes,
		DependencyIndexes: file_game_engine_proto_depIdxs,
		MessageInfos:      file_game_engine_proto_msgTypes,
	}.Build()
	File_game_engine_proto = out.File
	file_game_engine_proto_rawDesc = nil
	file_game_engine_proto_goTypes = nil
	file_game_engine_proto_depIdxs = nil
}
