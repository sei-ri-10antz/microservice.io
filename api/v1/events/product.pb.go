// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: api/v1/events/product.proto

package events

import (
	proto "github.com/golang/protobuf/proto"
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

type ProductCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Price    int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Qty      int64  `protobuf:"varint,5,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *ProductCreated) Reset() {
	*x = ProductCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_events_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreated) ProtoMessage() {}

func (x *ProductCreated) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_events_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreated.ProtoReflect.Descriptor instead.
func (*ProductCreated) Descriptor() ([]byte, []int) {
	return file_api_v1_events_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCreated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductCreated) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductCreated) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *ProductCreated) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductCreated) GetQty() int64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type ProductQtyDeducted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Qty int64  `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *ProductQtyDeducted) Reset() {
	*x = ProductQtyDeducted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_events_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductQtyDeducted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductQtyDeducted) ProtoMessage() {}

func (x *ProductQtyDeducted) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_events_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductQtyDeducted.ProtoReflect.Descriptor instead.
func (*ProductQtyDeducted) Descriptor() ([]byte, []int) {
	return file_api_v1_events_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductQtyDeducted) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductQtyDeducted) GetQty() int64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

var File_api_v1_events_product_proto protoreflect.FileDescriptor

var file_api_v1_events_product_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x79, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x71, 0x74, 0x79, 0x22, 0x36, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x51,
	0x74, 0x79, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x71, 0x74, 0x79, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x69, 0x2d, 0x72,
	0x69, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_events_product_proto_rawDescOnce sync.Once
	file_api_v1_events_product_proto_rawDescData = file_api_v1_events_product_proto_rawDesc
)

func file_api_v1_events_product_proto_rawDescGZIP() []byte {
	file_api_v1_events_product_proto_rawDescOnce.Do(func() {
		file_api_v1_events_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_events_product_proto_rawDescData)
	})
	return file_api_v1_events_product_proto_rawDescData
}

var file_api_v1_events_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_events_product_proto_goTypes = []interface{}{
	(*ProductCreated)(nil),     // 0: microservice.api.v1.events.ProductCreated
	(*ProductQtyDeducted)(nil), // 1: microservice.api.v1.events.ProductQtyDeducted
}
var file_api_v1_events_product_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_events_product_proto_init() }
func file_api_v1_events_product_proto_init() {
	if File_api_v1_events_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_events_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreated); i {
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
		file_api_v1_events_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductQtyDeducted); i {
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
			RawDescriptor: file_api_v1_events_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_events_product_proto_goTypes,
		DependencyIndexes: file_api_v1_events_product_proto_depIdxs,
		MessageInfos:      file_api_v1_events_product_proto_msgTypes,
	}.Build()
	File_api_v1_events_product_proto = out.File
	file_api_v1_events_product_proto_rawDesc = nil
	file_api_v1_events_product_proto_goTypes = nil
	file_api_v1_events_product_proto_depIdxs = nil
}
