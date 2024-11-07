// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: proto/orderProcessing.proto

package processingProto

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

type OrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqNumber int32  `protobuf:"varint,1,opt,name=uniqNumber,proto3" json:"uniqNumber,omitempty"`
	OrderName  string `protobuf:"bytes,2,opt,name=orderName,proto3" json:"orderName,omitempty"`
}

func (x *OrderInfo) Reset() {
	*x = OrderInfo{}
	mi := &file_proto_orderProcessing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfo) ProtoMessage() {}

func (x *OrderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orderProcessing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfo.ProtoReflect.Descriptor instead.
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return file_proto_orderProcessing_proto_rawDescGZIP(), []int{0}
}

func (x *OrderInfo) GetUniqNumber() int32 {
	if x != nil {
		return x.UniqNumber
	}
	return 0
}

func (x *OrderInfo) GetOrderName() string {
	if x != nil {
		return x.OrderName
	}
	return ""
}

type OrderStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqNumber int32  `protobuf:"varint,1,opt,name=uniqNumber,proto3" json:"uniqNumber,omitempty"`
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderStatus) Reset() {
	*x = OrderStatus{}
	mi := &file_proto_orderProcessing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatus) ProtoMessage() {}

func (x *OrderStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orderProcessing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatus.ProtoReflect.Descriptor instead.
func (*OrderStatus) Descriptor() ([]byte, []int) {
	return file_proto_orderProcessing_proto_rawDescGZIP(), []int{1}
}

func (x *OrderStatus) GetUniqNumber() int32 {
	if x != nil {
		return x.UniqNumber
	}
	return 0
}

func (x *OrderStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_orderProcessing_proto protoreflect.FileDescriptor

var file_proto_orderProcessing_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x22, 0x49,
	0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x6e, 0x69, 0x71, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x75, 0x6e, 0x69, 0x71, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x0b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x71,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x6e,
	0x69, 0x71, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x32, 0xab, 0x01, 0x0a, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1c,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_orderProcessing_proto_rawDescOnce sync.Once
	file_proto_orderProcessing_proto_rawDescData = file_proto_orderProcessing_proto_rawDesc
)

func file_proto_orderProcessing_proto_rawDescGZIP() []byte {
	file_proto_orderProcessing_proto_rawDescOnce.Do(func() {
		file_proto_orderProcessing_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_orderProcessing_proto_rawDescData)
	})
	return file_proto_orderProcessing_proto_rawDescData
}

var file_proto_orderProcessing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_orderProcessing_proto_goTypes = []any{
	(*OrderInfo)(nil),   // 0: orderProcessing.OrderInfo
	(*OrderStatus)(nil), // 1: orderProcessing.OrderStatus
}
var file_proto_orderProcessing_proto_depIdxs = []int32{
	0, // 0: orderProcessing.orderProcessingService.AddOrder:input_type -> orderProcessing.OrderInfo
	0, // 1: orderProcessing.orderProcessingService.CheckStatus:input_type -> orderProcessing.OrderInfo
	1, // 2: orderProcessing.orderProcessingService.AddOrder:output_type -> orderProcessing.OrderStatus
	1, // 3: orderProcessing.orderProcessingService.CheckStatus:output_type -> orderProcessing.OrderStatus
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_orderProcessing_proto_init() }
func file_proto_orderProcessing_proto_init() {
	if File_proto_orderProcessing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_orderProcessing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_orderProcessing_proto_goTypes,
		DependencyIndexes: file_proto_orderProcessing_proto_depIdxs,
		MessageInfos:      file_proto_orderProcessing_proto_msgTypes,
	}.Build()
	File_proto_orderProcessing_proto = out.File
	file_proto_orderProcessing_proto_rawDesc = nil
	file_proto_orderProcessing_proto_goTypes = nil
	file_proto_orderProcessing_proto_depIdxs = nil
}