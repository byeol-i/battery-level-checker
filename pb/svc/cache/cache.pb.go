// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pb/svc/cache/cache.proto

package cache

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

type WriteMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Value    []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *WriteMsgReq) Reset() {
	*x = WriteMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_cache_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteMsgReq) ProtoMessage() {}

func (x *WriteMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_cache_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteMsgReq.ProtoReflect.Descriptor instead.
func (*WriteMsgReq) Descriptor() ([]byte, []int) {
	return file_pb_svc_cache_cache_proto_rawDescGZIP(), []int{0}
}

func (x *WriteMsgReq) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *WriteMsgReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *WriteMsgReq) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type WriteMsgRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteMsgRes) Reset() {
	*x = WriteMsgRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_cache_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteMsgRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteMsgRes) ProtoMessage() {}

func (x *WriteMsgRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_cache_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteMsgRes.ProtoReflect.Descriptor instead.
func (*WriteMsgRes) Descriptor() ([]byte, []int) {
	return file_pb_svc_cache_cache_proto_rawDescGZIP(), []int{1}
}

type GetCurrentMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetCurrentMsgReq) Reset() {
	*x = GetCurrentMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_cache_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentMsgReq) ProtoMessage() {}

func (x *GetCurrentMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_cache_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentMsgReq.ProtoReflect.Descriptor instead.
func (*GetCurrentMsgReq) Descriptor() ([]byte, []int) {
	return file_pb_svc_cache_cache_proto_rawDescGZIP(), []int{2}
}

func (x *GetCurrentMsgReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetCurrentMsgRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []string `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetCurrentMsgRes) Reset() {
	*x = GetCurrentMsgRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_svc_cache_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentMsgRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentMsgRes) ProtoMessage() {}

func (x *GetCurrentMsgRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_svc_cache_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentMsgRes.ProtoReflect.Descriptor instead.
func (*GetCurrentMsgRes) Descriptor() ([]byte, []int) {
	return file_pb_svc_cache_cache_proto_rawDescGZIP(), []int{3}
}

func (x *GetCurrentMsgRes) GetResult() []string {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_pb_svc_cache_cache_proto protoreflect.FileDescriptor

var file_pb_svc_cache_cache_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x62, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x22, 0x57, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x9a, 0x01, 0x0a, 0x05, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x19,
	0x2e, 0x70, 0x62, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x73, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x65, 0x6f, 0x6c, 0x2d, 0x69, 0x2f, 0x62, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x2d, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_svc_cache_cache_proto_rawDescOnce sync.Once
	file_pb_svc_cache_cache_proto_rawDescData = file_pb_svc_cache_cache_proto_rawDesc
)

func file_pb_svc_cache_cache_proto_rawDescGZIP() []byte {
	file_pb_svc_cache_cache_proto_rawDescOnce.Do(func() {
		file_pb_svc_cache_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_svc_cache_cache_proto_rawDescData)
	})
	return file_pb_svc_cache_cache_proto_rawDescData
}

var file_pb_svc_cache_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_svc_cache_cache_proto_goTypes = []interface{}{
	(*WriteMsgReq)(nil),      // 0: pb.svc.cache.WriteMsgReq
	(*WriteMsgRes)(nil),      // 1: pb.svc.cache.WriteMsgRes
	(*GetCurrentMsgReq)(nil), // 2: pb.svc.cache.GetCurrentMsgReq
	(*GetCurrentMsgRes)(nil), // 3: pb.svc.cache.GetCurrentMsgRes
}
var file_pb_svc_cache_cache_proto_depIdxs = []int32{
	0, // 0: pb.svc.cache.Cache.WriteMsg:input_type -> pb.svc.cache.WriteMsgReq
	2, // 1: pb.svc.cache.Cache.GetCurrentMsg:input_type -> pb.svc.cache.GetCurrentMsgReq
	1, // 2: pb.svc.cache.Cache.WriteMsg:output_type -> pb.svc.cache.WriteMsgRes
	3, // 3: pb.svc.cache.Cache.GetCurrentMsg:output_type -> pb.svc.cache.GetCurrentMsgRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_svc_cache_cache_proto_init() }
func file_pb_svc_cache_cache_proto_init() {
	if File_pb_svc_cache_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_svc_cache_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteMsgReq); i {
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
		file_pb_svc_cache_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteMsgRes); i {
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
		file_pb_svc_cache_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentMsgReq); i {
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
		file_pb_svc_cache_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentMsgRes); i {
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
			RawDescriptor: file_pb_svc_cache_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_svc_cache_cache_proto_goTypes,
		DependencyIndexes: file_pb_svc_cache_cache_proto_depIdxs,
		MessageInfos:      file_pb_svc_cache_cache_proto_msgTypes,
	}.Build()
	File_pb_svc_cache_cache_proto = out.File
	file_pb_svc_cache_cache_proto_rawDesc = nil
	file_pb_svc_cache_cache_proto_goTypes = nil
	file_pb_svc_cache_cache_proto_depIdxs = nil
}
