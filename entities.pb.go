// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: agenda/entities.proto

package grpc_agenda_go

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Person message defining an entry of the agenda.
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id with the person identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name with the full name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Email to receive notifications.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agenda_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_agenda_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_agenda_entities_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// PersonList message with a collection of persons.
type PersonList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Persons array.
	Persons []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
}

func (x *PersonList) Reset() {
	*x = PersonList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agenda_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonList) ProtoMessage() {}

func (x *PersonList) ProtoReflect() protoreflect.Message {
	mi := &file_agenda_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonList.ProtoReflect.Descriptor instead.
func (*PersonList) Descriptor() ([]byte, []int) {
	return file_agenda_entities_proto_rawDescGZIP(), []int{1}
}

func (x *PersonList) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

// AddPersonRequest message to add a new entry to the agenda.
type AddPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the person.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Email to receive notifications.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AddPersonRequest) Reset() {
	*x = AddPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agenda_entities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPersonRequest) ProtoMessage() {}

func (x *AddPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agenda_entities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPersonRequest.ProtoReflect.Descriptor instead.
func (*AddPersonRequest) Descriptor() ([]byte, []int) {
	return file_agenda_entities_proto_rawDescGZIP(), []int{2}
}

func (x *AddPersonRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddPersonRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// ListPersonRequest message.
type ListPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPersonRequest) Reset() {
	*x = ListPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agenda_entities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonRequest) ProtoMessage() {}

func (x *ListPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agenda_entities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonRequest.ProtoReflect.Descriptor instead.
func (*ListPersonRequest) Descriptor() ([]byte, []int) {
	return file_agenda_entities_proto_rawDescGZIP(), []int{3}
}

var File_agenda_entities_proto protoreflect.FileDescriptor

var file_agenda_entities_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x36, 0x0a, 0x0a,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x64, 0x61, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x22, 0x45, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x13, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x70, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x61, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x2d, 0x67, 0x6f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x67, 0x65, 0x6e, 0x64, 0x61, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agenda_entities_proto_rawDescOnce sync.Once
	file_agenda_entities_proto_rawDescData = file_agenda_entities_proto_rawDesc
)

func file_agenda_entities_proto_rawDescGZIP() []byte {
	file_agenda_entities_proto_rawDescOnce.Do(func() {
		file_agenda_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_agenda_entities_proto_rawDescData)
	})
	return file_agenda_entities_proto_rawDescData
}

var file_agenda_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_agenda_entities_proto_goTypes = []interface{}{
	(*Person)(nil),            // 0: agenda.Person
	(*PersonList)(nil),        // 1: agenda.PersonList
	(*AddPersonRequest)(nil),  // 2: agenda.AddPersonRequest
	(*ListPersonRequest)(nil), // 3: agenda.ListPersonRequest
}
var file_agenda_entities_proto_depIdxs = []int32{
	0, // 0: agenda.PersonList.persons:type_name -> agenda.Person
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_agenda_entities_proto_init() }
func file_agenda_entities_proto_init() {
	if File_agenda_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agenda_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_agenda_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonList); i {
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
		file_agenda_entities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPersonRequest); i {
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
		file_agenda_entities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonRequest); i {
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
			RawDescriptor: file_agenda_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_agenda_entities_proto_goTypes,
		DependencyIndexes: file_agenda_entities_proto_depIdxs,
		MessageInfos:      file_agenda_entities_proto_msgTypes,
	}.Build()
	File_agenda_entities_proto = out.File
	file_agenda_entities_proto_rawDesc = nil
	file_agenda_entities_proto_goTypes = nil
	file_agenda_entities_proto_depIdxs = nil
}
