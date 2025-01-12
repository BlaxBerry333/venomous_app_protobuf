// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: notes.proto

package notes

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

type SelectableNoteType int32

const (
	SelectableNoteType_RAFT SelectableNoteType = 0
)

// Enum value maps for SelectableNoteType.
var (
	SelectableNoteType_name = map[int32]string{
		0: "RAFT",
	}
	SelectableNoteType_value = map[string]int32{
		"RAFT": 0,
	}
)

func (x SelectableNoteType) Enum() *SelectableNoteType {
	p := new(SelectableNoteType)
	*p = x
	return p
}

func (x SelectableNoteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SelectableNoteType) Descriptor() protoreflect.EnumDescriptor {
	return file_notes_proto_enumTypes[0].Descriptor()
}

func (SelectableNoteType) Type() protoreflect.EnumType {
	return &file_notes_proto_enumTypes[0]
}

func (x SelectableNoteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SelectableNoteType.Descriptor instead.
func (SelectableNoteType) EnumDescriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{0}
}

type NoteDataType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId       string             `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Type      SelectableNoteType `protobuf:"varint,2,opt,name=type,proto3,enum=notes.SelectableNoteType" json:"type,omitempty"`
	Title     string             `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Message   string             `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt string             `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string             `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *NoteDataType) Reset() {
	*x = NoteDataType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteDataType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteDataType) ProtoMessage() {}

func (x *NoteDataType) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteDataType.ProtoReflect.Descriptor instead.
func (*NoteDataType) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{0}
}

func (x *NoteDataType) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *NoteDataType) GetType() SelectableNoteType {
	if x != nil {
		return x.Type
	}
	return SelectableNoteType_RAFT
}

func (x *NoteDataType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NoteDataType) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NoteDataType) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NoteDataType) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_notes_proto protoreflect.FileDescriptor

var file_notes_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x2a, 0x1e, 0x0a, 0x12, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x41, 0x46,
	0x54, 0x10, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x67, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notes_proto_rawDescOnce sync.Once
	file_notes_proto_rawDescData = file_notes_proto_rawDesc
)

func file_notes_proto_rawDescGZIP() []byte {
	file_notes_proto_rawDescOnce.Do(func() {
		file_notes_proto_rawDescData = protoimpl.X.CompressGZIP(file_notes_proto_rawDescData)
	})
	return file_notes_proto_rawDescData
}

var file_notes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_notes_proto_goTypes = []interface{}{
	(SelectableNoteType)(0), // 0: notes.SelectableNoteType
	(*NoteDataType)(nil),    // 1: notes.NoteDataType
}
var file_notes_proto_depIdxs = []int32{
	0, // 0: notes.NoteDataType.type:type_name -> notes.SelectableNoteType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_notes_proto_init() }
func file_notes_proto_init() {
	if File_notes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteDataType); i {
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
			RawDescriptor: file_notes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notes_proto_goTypes,
		DependencyIndexes: file_notes_proto_depIdxs,
		EnumInfos:         file_notes_proto_enumTypes,
		MessageInfos:      file_notes_proto_msgTypes,
	}.Build()
	File_notes_proto = out.File
	file_notes_proto_rawDesc = nil
	file_notes_proto_goTypes = nil
	file_notes_proto_depIdxs = nil
}
