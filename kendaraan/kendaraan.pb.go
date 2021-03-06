// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: kendaraan/kendaraan.proto

package kendaraan

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kendaraan_kendaraan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_kendaraan_kendaraan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_kendaraan_kendaraan_proto_rawDescGZIP(), []int{0}
}

type InfoMobil struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Warna      string      `protobuf:"bytes,2,opt,name=warna,proto3" json:"warna,omitempty"`
	Tahun      int32       `protobuf:"varint,3,opt,name=tahun,proto3" json:"tahun,omitempty"`
	Manufaktur *Manufaktur `protobuf:"bytes,4,opt,name=manufaktur,proto3" json:"manufaktur,omitempty"`
}

func (x *InfoMobil) Reset() {
	*x = InfoMobil{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kendaraan_kendaraan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoMobil) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoMobil) ProtoMessage() {}

func (x *InfoMobil) ProtoReflect() protoreflect.Message {
	mi := &file_kendaraan_kendaraan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoMobil.ProtoReflect.Descriptor instead.
func (*InfoMobil) Descriptor() ([]byte, []int) {
	return file_kendaraan_kendaraan_proto_rawDescGZIP(), []int{1}
}

func (x *InfoMobil) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InfoMobil) GetWarna() string {
	if x != nil {
		return x.Warna
	}
	return ""
}

func (x *InfoMobil) GetTahun() int32 {
	if x != nil {
		return x.Tahun
	}
	return 0
}

func (x *InfoMobil) GetManufaktur() *Manufaktur {
	if x != nil {
		return x.Manufaktur
	}
	return nil
}

type Manufaktur struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Merek string `protobuf:"bytes,1,opt,name=merek,proto3" json:"merek,omitempty"`
	Jenis string `protobuf:"bytes,2,opt,name=jenis,proto3" json:"jenis,omitempty"`
}

func (x *Manufaktur) Reset() {
	*x = Manufaktur{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kendaraan_kendaraan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manufaktur) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manufaktur) ProtoMessage() {}

func (x *Manufaktur) ProtoReflect() protoreflect.Message {
	mi := &file_kendaraan_kendaraan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manufaktur.ProtoReflect.Descriptor instead.
func (*Manufaktur) Descriptor() ([]byte, []int) {
	return file_kendaraan_kendaraan_proto_rawDescGZIP(), []int{2}
}

func (x *Manufaktur) GetMerek() string {
	if x != nil {
		return x.Merek
	}
	return ""
}

func (x *Manufaktur) GetJenis() string {
	if x != nil {
		return x.Jenis
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kendaraan_kendaraan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_kendaraan_kendaraan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_kendaraan_kendaraan_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kendaraan_kendaraan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_kendaraan_kendaraan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_kendaraan_kendaraan_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_kendaraan_kendaraan_proto protoreflect.FileDescriptor

var file_kendaraan_kendaraan_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x2f, 0x6b, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x61, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6b, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x7e, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x61, 0x72, 0x6e, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x61, 0x72,
	0x6e, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x68, 0x75, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x61, 0x68, 0x75, 0x6e, 0x12, 0x35, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x6b, 0x74, 0x75, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x6b,
	0x74, 0x75, 0x72, 0x52, 0x0a, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x6b, 0x74, 0x75, 0x72, 0x22,
	0x38, 0x0a, 0x0a, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x6b, 0x74, 0x75, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x65, 0x72, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x65,
	0x72, 0x65, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x65, 0x6e, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6a, 0x65, 0x6e, 0x69, 0x73, 0x22, 0x1a, 0x0a, 0x02, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x8e, 0x02, 0x0a, 0x05, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x12,
	0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x12, 0x10,
	0x2e, 0x6b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x14, 0x2e, 0x6b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x12, 0x0d, 0x2e, 0x6b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e,
	0x2e, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x6b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x6b, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x61, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x1a, 0x0d,
	0x2e, 0x6b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x2e, 0x49, 0x64, 0x12, 0x36, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x6b,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x1a, 0x11, 0x2e, 0x6b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x12, 0x0d, 0x2e, 0x6b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e,
	0x2e, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x6b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x1d, 0x5a, 0x1b, 0x62, 0x65, 0x6c, 0x61, 0x6a, 0x61,
	0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x6b, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x61, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kendaraan_kendaraan_proto_rawDescOnce sync.Once
	file_kendaraan_kendaraan_proto_rawDescData = file_kendaraan_kendaraan_proto_rawDesc
)

func file_kendaraan_kendaraan_proto_rawDescGZIP() []byte {
	file_kendaraan_kendaraan_proto_rawDescOnce.Do(func() {
		file_kendaraan_kendaraan_proto_rawDescData = protoimpl.X.CompressGZIP(file_kendaraan_kendaraan_proto_rawDescData)
	})
	return file_kendaraan_kendaraan_proto_rawDescData
}

var file_kendaraan_kendaraan_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_kendaraan_kendaraan_proto_goTypes = []interface{}{
	(*Empty)(nil),      // 0: kendaraan.Empty
	(*InfoMobil)(nil),  // 1: kendaraan.InfoMobil
	(*Manufaktur)(nil), // 2: kendaraan.Manufaktur
	(*Id)(nil),         // 3: kendaraan.Id
	(*Status)(nil),     // 4: kendaraan.Status
}
var file_kendaraan_kendaraan_proto_depIdxs = []int32{
	2, // 0: kendaraan.InfoMobil.manufaktur:type_name -> kendaraan.Manufaktur
	0, // 1: kendaraan.Mobil.GetAllMobil:input_type -> kendaraan.Empty
	3, // 2: kendaraan.Mobil.GetMobil:input_type -> kendaraan.Id
	1, // 3: kendaraan.Mobil.CreateMobil:input_type -> kendaraan.InfoMobil
	1, // 4: kendaraan.Mobil.UpdateMobil:input_type -> kendaraan.InfoMobil
	3, // 5: kendaraan.Mobil.DeleteMobil:input_type -> kendaraan.Id
	1, // 6: kendaraan.Mobil.GetAllMobil:output_type -> kendaraan.InfoMobil
	1, // 7: kendaraan.Mobil.GetMobil:output_type -> kendaraan.InfoMobil
	3, // 8: kendaraan.Mobil.CreateMobil:output_type -> kendaraan.Id
	4, // 9: kendaraan.Mobil.UpdateMobil:output_type -> kendaraan.Status
	4, // 10: kendaraan.Mobil.DeleteMobil:output_type -> kendaraan.Status
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kendaraan_kendaraan_proto_init() }
func file_kendaraan_kendaraan_proto_init() {
	if File_kendaraan_kendaraan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kendaraan_kendaraan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_kendaraan_kendaraan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoMobil); i {
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
		file_kendaraan_kendaraan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manufaktur); i {
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
		file_kendaraan_kendaraan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_kendaraan_kendaraan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_kendaraan_kendaraan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kendaraan_kendaraan_proto_goTypes,
		DependencyIndexes: file_kendaraan_kendaraan_proto_depIdxs,
		MessageInfos:      file_kendaraan_kendaraan_proto_msgTypes,
	}.Build()
	File_kendaraan_kendaraan_proto = out.File
	file_kendaraan_kendaraan_proto_rawDesc = nil
	file_kendaraan_kendaraan_proto_goTypes = nil
	file_kendaraan_kendaraan_proto_depIdxs = nil
}
