// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: proto/todo_service.proto

package proto

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

// The request message containing the user's name.
type CreateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *NewTodo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTodoRequest) GetTodo() *NewTodo {
	if x != nil {
		return x.Todo
	}
	return nil
}

// The response message containing the greetings
type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_service_proto_rawDescGZIP(), []int{1}
}

func (x *TodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

// The request message containing the user's name.
type TodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *TodosRequest) Reset() {
	*x = TodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodosRequest) ProtoMessage() {}

func (x *TodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodosRequest.ProtoReflect.Descriptor instead.
func (*TodosRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_service_proto_rawDescGZIP(), []int{2}
}

func (x *TodosRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// The request message containing the user's name.
type TodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *TodosResponse) Reset() {
	*x = TodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodosResponse) ProtoMessage() {}

func (x *TodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodosResponse.ProtoReflect.Descriptor instead.
func (*TodosResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_service_proto_rawDescGZIP(), []int{3}
}

func (x *TodosResponse) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type NewTodo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Done   bool   `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	UserId int32  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *NewTodo) Reset() {
	*x = NewTodo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTodo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTodo) ProtoMessage() {}

func (x *NewTodo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTodo.ProtoReflect.Descriptor instead.
func (*NewTodo) Descriptor() ([]byte, []int) {
	return file_proto_todo_service_proto_rawDescGZIP(), []int{4}
}

func (x *NewTodo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewTodo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *NewTodo) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Done   bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	UserId int32  `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_proto_todo_service_proto_rawDescGZIP(), []int{5}
}

func (x *Todo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *Todo) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_proto_todo_service_proto protoreflect.FileDescriptor

var file_proto_todo_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x2f, 0x0a, 0x0c, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x6f,
	0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x26, 0x0a, 0x0c, 0x54,
	0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0d, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x4b, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54, 0x6f,
	0x64, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x82,
	0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x05, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todo_service_proto_rawDescOnce sync.Once
	file_proto_todo_service_proto_rawDescData = file_proto_todo_service_proto_rawDesc
)

func file_proto_todo_service_proto_rawDescGZIP() []byte {
	file_proto_todo_service_proto_rawDescOnce.Do(func() {
		file_proto_todo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todo_service_proto_rawDescData)
	})
	return file_proto_todo_service_proto_rawDescData
}

var file_proto_todo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_todo_service_proto_goTypes = []interface{}{
	(*CreateTodoRequest)(nil), // 0: proto.CreateTodoRequest
	(*TodoResponse)(nil),      // 1: proto.TodoResponse
	(*TodosRequest)(nil),      // 2: proto.TodosRequest
	(*TodosResponse)(nil),     // 3: proto.TodosResponse
	(*NewTodo)(nil),           // 4: proto.NewTodo
	(*Todo)(nil),              // 5: proto.Todo
}
var file_proto_todo_service_proto_depIdxs = []int32{
	4, // 0: proto.CreateTodoRequest.todo:type_name -> proto.NewTodo
	5, // 1: proto.TodoResponse.todo:type_name -> proto.Todo
	5, // 2: proto.TodosResponse.todos:type_name -> proto.Todo
	0, // 3: proto.TodoService.CreateTodo:input_type -> proto.CreateTodoRequest
	2, // 4: proto.TodoService.Todos:input_type -> proto.TodosRequest
	1, // 5: proto.TodoService.CreateTodo:output_type -> proto.TodoResponse
	3, // 6: proto.TodoService.Todos:output_type -> proto.TodosResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_todo_service_proto_init() }
func file_proto_todo_service_proto_init() {
	if File_proto_todo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequest); i {
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
		file_proto_todo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResponse); i {
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
		file_proto_todo_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodosRequest); i {
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
		file_proto_todo_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodosResponse); i {
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
		file_proto_todo_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTodo); i {
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
		file_proto_todo_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
			RawDescriptor: file_proto_todo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todo_service_proto_goTypes,
		DependencyIndexes: file_proto_todo_service_proto_depIdxs,
		MessageInfos:      file_proto_todo_service_proto_msgTypes,
	}.Build()
	File_proto_todo_service_proto = out.File
	file_proto_todo_service_proto_rawDesc = nil
	file_proto_todo_service_proto_goTypes = nil
	file_proto_todo_service_proto_depIdxs = nil
}
