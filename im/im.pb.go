// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: im/im.proto

package im

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

type OfflinePushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDs []string            `protobuf:"bytes,1,rep,name=userIDs,proto3" json:"userIDs,omitempty"`
	MsgData *OfflinePushMsgData `protobuf:"bytes,2,opt,name=msgData,proto3" json:"msgData,omitempty"`
}

func (x *OfflinePushMsgReq) Reset() {
	*x = OfflinePushMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_im_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflinePushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflinePushMsgReq) ProtoMessage() {}

func (x *OfflinePushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_im_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflinePushMsgReq.ProtoReflect.Descriptor instead.
func (*OfflinePushMsgReq) Descriptor() ([]byte, []int) {
	return file_im_im_proto_rawDescGZIP(), []int{0}
}

func (x *OfflinePushMsgReq) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

func (x *OfflinePushMsgReq) GetMsgData() *OfflinePushMsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

type OfflinePushMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OfflinePushMsgResp) Reset() {
	*x = OfflinePushMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_im_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflinePushMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflinePushMsgResp) ProtoMessage() {}

func (x *OfflinePushMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_im_im_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflinePushMsgResp.ProtoReflect.Descriptor instead.
func (*OfflinePushMsgResp) Descriptor() ([]byte, []int) {
	return file_im_im_proto_rawDescGZIP(), []int{1}
}

type OfflinePushMsgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title           string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content         string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UseIOSPushSound string `protobuf:"bytes,3,opt,name=useIOSPushSound,proto3" json:"useIOSPushSound,omitempty"`
	ShowBadgeCount  bool   `protobuf:"varint,4,opt,name=showBadgeCount,proto3" json:"showBadgeCount,omitempty"`
	Ex              string `protobuf:"bytes,5,opt,name=ex,proto3" json:"ex,omitempty"`
}

func (x *OfflinePushMsgData) Reset() {
	*x = OfflinePushMsgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_im_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflinePushMsgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflinePushMsgData) ProtoMessage() {}

func (x *OfflinePushMsgData) ProtoReflect() protoreflect.Message {
	mi := &file_im_im_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflinePushMsgData.ProtoReflect.Descriptor instead.
func (*OfflinePushMsgData) Descriptor() ([]byte, []int) {
	return file_im_im_proto_rawDescGZIP(), []int{2}
}

func (x *OfflinePushMsgData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OfflinePushMsgData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *OfflinePushMsgData) GetUseIOSPushSound() string {
	if x != nil {
		return x.UseIOSPushSound
	}
	return ""
}

func (x *OfflinePushMsgData) GetShowBadgeCount() bool {
	if x != nil {
		return x.ShowBadgeCount
	}
	return false
}

func (x *OfflinePushMsgData) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

var File_im_im_proto protoreflect.FileDescriptor

var file_im_im_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6d, 0x2f, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69,
	0x6d, 0x22, 0x5f, 0x0a, 0x11, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73,
	0x12, 0x30, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6d, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75,
	0x73, 0x68, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x14, 0x0a, 0x12, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73,
	0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0xa6, 0x01, 0x0a, 0x12, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x49, 0x4f, 0x53, 0x50, 0x75, 0x73, 0x68, 0x53, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x49, 0x4f, 0x53,
	0x50, 0x75, 0x73, 0x68, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x68, 0x6f,
	0x77, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x73, 0x68, 0x6f, 0x77, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65,
	0x78, 0x32, 0x4c, 0x0a, 0x09, 0x49, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f,
	0x0a, 0x0e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67,
	0x12, 0x15, 0x2e, 0x69, 0x6d, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73,
	0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x69, 0x6d, 0x2e, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6f,
	0x75, 0x6e, 0x74, 0x79, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_im_proto_rawDescOnce sync.Once
	file_im_im_proto_rawDescData = file_im_im_proto_rawDesc
)

func file_im_im_proto_rawDescGZIP() []byte {
	file_im_im_proto_rawDescOnce.Do(func() {
		file_im_im_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_im_proto_rawDescData)
	})
	return file_im_im_proto_rawDescData
}

var file_im_im_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_im_im_proto_goTypes = []interface{}{
	(*OfflinePushMsgReq)(nil),  // 0: im.OfflinePushMsgReq
	(*OfflinePushMsgResp)(nil), // 1: im.OfflinePushMsgResp
	(*OfflinePushMsgData)(nil), // 2: im.OfflinePushMsgData
}
var file_im_im_proto_depIdxs = []int32{
	2, // 0: im.OfflinePushMsgReq.msgData:type_name -> im.OfflinePushMsgData
	0, // 1: im.ImService.OfflinePushMsg:input_type -> im.OfflinePushMsgReq
	1, // 2: im.ImService.OfflinePushMsg:output_type -> im.OfflinePushMsgResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_im_im_proto_init() }
func file_im_im_proto_init() {
	if File_im_im_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_im_im_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflinePushMsgReq); i {
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
		file_im_im_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflinePushMsgResp); i {
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
		file_im_im_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflinePushMsgData); i {
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
			RawDescriptor: file_im_im_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_im_proto_goTypes,
		DependencyIndexes: file_im_im_proto_depIdxs,
		MessageInfos:      file_im_im_proto_msgTypes,
	}.Build()
	File_im_im_proto = out.File
	file_im_im_proto_rawDesc = nil
	file_im_im_proto_goTypes = nil
	file_im_im_proto_depIdxs = nil
}