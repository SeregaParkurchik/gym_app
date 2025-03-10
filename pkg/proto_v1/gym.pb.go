// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: gym.proto

package proto_v1

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

type Administrator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Role string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Administrator) Reset() {
	*x = Administrator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gym_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Administrator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Administrator) ProtoMessage() {}

func (x *Administrator) ProtoReflect() protoreflect.Message {
	mi := &file_gym_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Administrator.ProtoReflect.Descriptor instead.
func (*Administrator) Descriptor() ([]byte, []int) {
	return file_gym_proto_rawDescGZIP(), []int{0}
}

func (x *Administrator) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Administrator) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Administrator) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Administrator *Administrator `protobuf:"bytes,1,opt,name=administrator,proto3" json:"administrator,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gym_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gym_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_gym_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetAdministrator() *Administrator {
	if x != nil {
		return x.Administrator
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Administrator *Administrator `protobuf:"bytes,1,opt,name=administrator,proto3" json:"administrator,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gym_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gym_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_gym_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetAdministrator() *Administrator {
	if x != nil {
		return x.Administrator
	}
	return nil
}

var File_gym_proto protoreflect.FileDescriptor

var file_gym_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x4b, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0d,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x4c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x0a, 0x0f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x12, 0x3f, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x65, 0x67, 0x61, 0x50,
	0x61, 0x72, 0x6b, 0x75, 0x72, 0x63, 0x68, 0x69, 0x6b, 0x2f, 0x67, 0x79, 0x6d, 0x5f, 0x61, 0x70,
	0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gym_proto_rawDescOnce sync.Once
	file_gym_proto_rawDescData = file_gym_proto_rawDesc
)

func file_gym_proto_rawDescGZIP() []byte {
	file_gym_proto_rawDescOnce.Do(func() {
		file_gym_proto_rawDescData = protoimpl.X.CompressGZIP(file_gym_proto_rawDescData)
	})
	return file_gym_proto_rawDescData
}

var file_gym_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gym_proto_goTypes = []interface{}{
	(*Administrator)(nil),  // 0: proto.Administrator
	(*CreateRequest)(nil),  // 1: proto.CreateRequest
	(*CreateResponse)(nil), // 2: proto.CreateResponse
}
var file_gym_proto_depIdxs = []int32{
	0, // 0: proto.CreateRequest.administrator:type_name -> proto.Administrator
	0, // 1: proto.CreateResponse.administrator:type_name -> proto.Administrator
	1, // 2: proto.AdministratorV1.CreateMembership:input_type -> proto.CreateRequest
	2, // 3: proto.AdministratorV1.CreateMembership:output_type -> proto.CreateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gym_proto_init() }
func file_gym_proto_init() {
	if File_gym_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gym_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Administrator); i {
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
		file_gym_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_gym_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_gym_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gym_proto_goTypes,
		DependencyIndexes: file_gym_proto_depIdxs,
		MessageInfos:      file_gym_proto_msgTypes,
	}.Build()
	File_gym_proto = out.File
	file_gym_proto_rawDesc = nil
	file_gym_proto_goTypes = nil
	file_gym_proto_depIdxs = nil
}
