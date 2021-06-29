// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: resolver/aws/v1/aws.proto

package awsv1

import (
	_ "github.com/lyft/clutch/backend/api/resolver/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InstanceID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *InstanceID) Reset() {
	*x = InstanceID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_aws_v1_aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceID) ProtoMessage() {}

func (x *InstanceID) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_aws_v1_aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceID.ProtoReflect.Descriptor instead.
func (*InstanceID) Descriptor() ([]byte, []int) {
	return file_resolver_aws_v1_aws_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InstanceID) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type AutoscalingGroupName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *AutoscalingGroupName) Reset() {
	*x = AutoscalingGroupName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_aws_v1_aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoscalingGroupName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoscalingGroupName) ProtoMessage() {}

func (x *AutoscalingGroupName) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_aws_v1_aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoscalingGroupName.ProtoReflect.Descriptor instead.
func (*AutoscalingGroupName) Descriptor() ([]byte, []int) {
	return file_resolver_aws_v1_aws_proto_rawDescGZIP(), []int{1}
}

func (x *AutoscalingGroupName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AutoscalingGroupName) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type KinesisStreamName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *KinesisStreamName) Reset() {
	*x = KinesisStreamName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_aws_v1_aws_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KinesisStreamName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KinesisStreamName) ProtoMessage() {}

func (x *KinesisStreamName) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_aws_v1_aws_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KinesisStreamName.ProtoReflect.Descriptor instead.
func (*KinesisStreamName) Descriptor() ([]byte, []int) {
	return file_resolver_aws_v1_aws_proto_rawDescGZIP(), []int{2}
}

func (x *KinesisStreamName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KinesisStreamName) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

var File_resolver_aws_v1_aws_proto protoreflect.FileDescriptor

var file_resolver_aws_v1_aws_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1a, 0xea, 0x9f, 0x1d, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x01, 0x1a, 0x0e, 0x0a, 0x0c, 0x69,
	0x2d, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x31, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x19, 0xea, 0x9f, 0x1d, 0x15, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x0b, 0x08,
	0x01, 0x12, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x3a, 0x15, 0xea, 0x9f, 0x1d, 0x11, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x20, 0x49, 0x44, 0x1a, 0x02, 0x08, 0x01, 0x22, 0x98, 0x01, 0x0a, 0x14, 0x41, 0x75,
	0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x29, 0xea, 0x9f, 0x1d, 0x25, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x1a, 0x1b,
	0x0a, 0x19, 0x6d, 0x79, 0x2d, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67,
	0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x19, 0xea, 0x9f, 0x1d, 0x15, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22,
	0x0b, 0x08, 0x01, 0x12, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x3a, 0x0e, 0xea, 0x9f, 0x1d, 0x0a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x1a, 0x02, 0x08, 0x01, 0x22, 0x92, 0x01, 0x0a, 0x11, 0x4b, 0x69, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xea, 0x9f, 0x1d, 0x22, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x1a, 0x18, 0x0a, 0x16, 0x6d, 0x79, 0x2d, 0x6b, 0x69, 0x6e,
	0x65, 0x73, 0x69, 0x73, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xea, 0x9f, 0x1d, 0x15, 0x0a, 0x06, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x22, 0x0b, 0x08, 0x01, 0x12, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x3a, 0x0e, 0xea, 0x9f, 0x1d, 0x0a, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x02, 0x08, 0x01, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x77, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resolver_aws_v1_aws_proto_rawDescOnce sync.Once
	file_resolver_aws_v1_aws_proto_rawDescData = file_resolver_aws_v1_aws_proto_rawDesc
)

func file_resolver_aws_v1_aws_proto_rawDescGZIP() []byte {
	file_resolver_aws_v1_aws_proto_rawDescOnce.Do(func() {
		file_resolver_aws_v1_aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_resolver_aws_v1_aws_proto_rawDescData)
	})
	return file_resolver_aws_v1_aws_proto_rawDescData
}

var file_resolver_aws_v1_aws_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resolver_aws_v1_aws_proto_goTypes = []interface{}{
	(*InstanceID)(nil),           // 0: clutch.resolver.aws.v1.InstanceID
	(*AutoscalingGroupName)(nil), // 1: clutch.resolver.aws.v1.AutoscalingGroupName
	(*KinesisStreamName)(nil),    // 2: clutch.resolver.aws.v1.KinesisStreamName
}
var file_resolver_aws_v1_aws_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resolver_aws_v1_aws_proto_init() }
func file_resolver_aws_v1_aws_proto_init() {
	if File_resolver_aws_v1_aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resolver_aws_v1_aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceID); i {
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
		file_resolver_aws_v1_aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoscalingGroupName); i {
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
		file_resolver_aws_v1_aws_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KinesisStreamName); i {
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
			RawDescriptor: file_resolver_aws_v1_aws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resolver_aws_v1_aws_proto_goTypes,
		DependencyIndexes: file_resolver_aws_v1_aws_proto_depIdxs,
		MessageInfos:      file_resolver_aws_v1_aws_proto_msgTypes,
	}.Build()
	File_resolver_aws_v1_aws_proto = out.File
	file_resolver_aws_v1_aws_proto_rawDesc = nil
	file_resolver_aws_v1_aws_proto_goTypes = nil
	file_resolver_aws_v1_aws_proto_depIdxs = nil
}
