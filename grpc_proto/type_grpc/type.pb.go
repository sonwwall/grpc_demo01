// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: type.proto

package type_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A1            float64                `protobuf:"fixed64,1,opt,name=a1,proto3" json:"a1,omitempty"`
	A2            float32                `protobuf:"fixed32,2,opt,name=a2,proto3" json:"a2,omitempty"`
	A3            int32                  `protobuf:"varint,3,opt,name=a3,proto3" json:"a3,omitempty"`
	A4            uint32                 `protobuf:"varint,4,opt,name=a4,proto3" json:"a4,omitempty"`
	A5            uint64                 `protobuf:"varint,5,opt,name=a5,proto3" json:"a5,omitempty"`
	A6            int32                  `protobuf:"zigzag32,6,opt,name=a6,proto3" json:"a6,omitempty"`
	A7            int64                  `protobuf:"zigzag64,7,opt,name=a7,proto3" json:"a7,omitempty"`
	A8            uint32                 `protobuf:"fixed32,8,opt,name=a8,proto3" json:"a8,omitempty"`
	A9            uint64                 `protobuf:"fixed64,9,opt,name=a9,proto3" json:"a9,omitempty"`
	A10           int32                  `protobuf:"fixed32,10,opt,name=a10,proto3" json:"a10,omitempty"`
	A11           int64                  `protobuf:"fixed64,11,opt,name=a11,proto3" json:"a11,omitempty"`
	A12           bool                   `protobuf:"varint,12,opt,name=a12,proto3" json:"a12,omitempty"`
	A13           string                 `protobuf:"bytes,13,opt,name=a13,proto3" json:"a13,omitempty"`
	A14           []byte                 `protobuf:"bytes,14,opt,name=a14,proto3" json:"a14,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetA1() float64 {
	if x != nil {
		return x.A1
	}
	return 0
}

func (x *Request) GetA2() float32 {
	if x != nil {
		return x.A2
	}
	return 0
}

func (x *Request) GetA3() int32 {
	if x != nil {
		return x.A3
	}
	return 0
}

func (x *Request) GetA4() uint32 {
	if x != nil {
		return x.A4
	}
	return 0
}

func (x *Request) GetA5() uint64 {
	if x != nil {
		return x.A5
	}
	return 0
}

func (x *Request) GetA6() int32 {
	if x != nil {
		return x.A6
	}
	return 0
}

func (x *Request) GetA7() int64 {
	if x != nil {
		return x.A7
	}
	return 0
}

func (x *Request) GetA8() uint32 {
	if x != nil {
		return x.A8
	}
	return 0
}

func (x *Request) GetA9() uint64 {
	if x != nil {
		return x.A9
	}
	return 0
}

func (x *Request) GetA10() int32 {
	if x != nil {
		return x.A10
	}
	return 0
}

func (x *Request) GetA11() int64 {
	if x != nil {
		return x.A11
	}
	return 0
}

func (x *Request) GetA12() bool {
	if x != nil {
		return x.A12
	}
	return false
}

func (x *Request) GetA13() string {
	if x != nil {
		return x.A13
	}
	return ""
}

func (x *Request) GetA14() []byte {
	if x != nil {
		return x.A14
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{1}
}

type Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code          uint32                 `protobuf:"fixed32,2,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_type_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

// list
type ArraryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	I6List        []int64                `protobuf:"varint,1,rep,packed,name=i6_list,json=i6List,proto3" json:"i6_list,omitempty"`
	SList         []string               `protobuf:"bytes,2,rep,name=s_list,json=sList,proto3" json:"s_list,omitempty"`
	ItemList      []*Item                `protobuf:"bytes,3,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ArraryRequest) Reset() {
	*x = ArraryRequest{}
	mi := &file_type_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ArraryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArraryRequest) ProtoMessage() {}

func (x *ArraryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArraryRequest.ProtoReflect.Descriptor instead.
func (*ArraryRequest) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{3}
}

func (x *ArraryRequest) GetI6List() []int64 {
	if x != nil {
		return x.I6List
	}
	return nil
}

func (x *ArraryRequest) GetSList() []string {
	if x != nil {
		return x.SList
	}
	return nil
}

func (x *ArraryRequest) GetItemList() []*Item {
	if x != nil {
		return x.ItemList
	}
	return nil
}

// map
type MapRequset struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IS            map[int32]string       `protobuf:"bytes,1,rep,name=i_s,json=iS,proto3" json:"i_s,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SS            map[string]string      `protobuf:"bytes,2,rep,name=s_s,json=sS,proto3" json:"s_s,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SItem         map[string]*Item       `protobuf:"bytes,3,rep,name=s_item,json=sItem,proto3" json:"s_item,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapRequset) Reset() {
	*x = MapRequset{}
	mi := &file_type_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapRequset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRequset) ProtoMessage() {}

func (x *MapRequset) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRequset.ProtoReflect.Descriptor instead.
func (*MapRequset) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{4}
}

func (x *MapRequset) GetIS() map[int32]string {
	if x != nil {
		return x.IS
	}
	return nil
}

func (x *MapRequset) GetSS() map[string]string {
	if x != nil {
		return x.SS
	}
	return nil
}

func (x *MapRequset) GetSItem() map[string]*Item {
	if x != nil {
		return x.SItem
	}
	return nil
}

// 嵌套
type Q1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name1         string                 `protobuf:"bytes,1,opt,name=name1,proto3" json:"name1,omitempty"`
	Q2            *Q1_Q2                 `protobuf:"bytes,2,opt,name=q2,proto3" json:"q2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Q1) Reset() {
	*x = Q1{}
	mi := &file_type_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Q1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Q1) ProtoMessage() {}

func (x *Q1) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Q1.ProtoReflect.Descriptor instead.
func (*Q1) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{5}
}

func (x *Q1) GetName1() string {
	if x != nil {
		return x.Name1
	}
	return ""
}

func (x *Q1) GetQ2() *Q1_Q2 {
	if x != nil {
		return x.Q2
	}
	return nil
}

type Q1_Q2 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name2         string                 `protobuf:"bytes,2,opt,name=name2,proto3" json:"name2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Q1_Q2) Reset() {
	*x = Q1_Q2{}
	mi := &file_type_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Q1_Q2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Q1_Q2) ProtoMessage() {}

func (x *Q1_Q2) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Q1_Q2.ProtoReflect.Descriptor instead.
func (*Q1_Q2) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Q1_Q2) GetName2() string {
	if x != nil {
		return x.Name2
	}
	return ""
}

var File_type_proto protoreflect.FileDescriptor

const file_type_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"type.proto\"\xf3\x01\n" +
	"\aRequest\x12\x0e\n" +
	"\x02a1\x18\x01 \x01(\x01R\x02a1\x12\x0e\n" +
	"\x02a2\x18\x02 \x01(\x02R\x02a2\x12\x0e\n" +
	"\x02a3\x18\x03 \x01(\x05R\x02a3\x12\x0e\n" +
	"\x02a4\x18\x04 \x01(\rR\x02a4\x12\x0e\n" +
	"\x02a5\x18\x05 \x01(\x04R\x02a5\x12\x0e\n" +
	"\x02a6\x18\x06 \x01(\x11R\x02a6\x12\x0e\n" +
	"\x02a7\x18\a \x01(\x12R\x02a7\x12\x0e\n" +
	"\x02a8\x18\b \x01(\aR\x02a8\x12\x0e\n" +
	"\x02a9\x18\t \x01(\x06R\x02a9\x12\x10\n" +
	"\x03a10\x18\n" +
	" \x01(\x0fR\x03a10\x12\x10\n" +
	"\x03a11\x18\v \x01(\x10R\x03a11\x12\x10\n" +
	"\x03a12\x18\f \x01(\bR\x03a12\x12\x10\n" +
	"\x03a13\x18\r \x01(\tR\x03a13\x12\x10\n" +
	"\x03a14\x18\x0e \x01(\fR\x03a14\"\n" +
	"\n" +
	"\bResponse\".\n" +
	"\x04Item\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04code\x18\x02 \x01(\aR\x04code\"c\n" +
	"\rArraryRequest\x12\x17\n" +
	"\ai6_list\x18\x01 \x03(\x03R\x06i6List\x12\x15\n" +
	"\x06s_list\x18\x02 \x03(\tR\x05sList\x12\"\n" +
	"\titem_list\x18\x03 \x03(\v2\x05.ItemR\bitemList\"\xb6\x02\n" +
	"\n" +
	"MapRequset\x12$\n" +
	"\x03i_s\x18\x01 \x03(\v2\x13.MapRequset.ISEntryR\x02iS\x12$\n" +
	"\x03s_s\x18\x02 \x03(\v2\x13.MapRequset.SSEntryR\x02sS\x12-\n" +
	"\x06s_item\x18\x03 \x03(\v2\x16.MapRequset.SItemEntryR\x05sItem\x1a5\n" +
	"\aISEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a5\n" +
	"\aSSEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a?\n" +
	"\n" +
	"SItemEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x1b\n" +
	"\x05value\x18\x02 \x01(\v2\x05.ItemR\x05value:\x028\x01\"N\n" +
	"\x02Q1\x12\x14\n" +
	"\x05name1\x18\x01 \x01(\tR\x05name1\x12\x16\n" +
	"\x02q2\x18\x02 \x01(\v2\x06.Q1.Q2R\x02q2\x1a\x1a\n" +
	"\x02Q2\x12\x14\n" +
	"\x05name2\x18\x02 \x01(\tR\x05name22+\n" +
	"\vTypeService\x12\x1c\n" +
	"\x03Say\x12\b.Request\x1a\t.Response\"\x00B\fZ\n" +
	"/type_grpcb\x06proto3"

var (
	file_type_proto_rawDescOnce sync.Once
	file_type_proto_rawDescData []byte
)

func file_type_proto_rawDescGZIP() []byte {
	file_type_proto_rawDescOnce.Do(func() {
		file_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_type_proto_rawDesc), len(file_type_proto_rawDesc)))
	})
	return file_type_proto_rawDescData
}

var file_type_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_type_proto_goTypes = []any{
	(*Request)(nil),       // 0: Request
	(*Response)(nil),      // 1: Response
	(*Item)(nil),          // 2: Item
	(*ArraryRequest)(nil), // 3: ArraryRequest
	(*MapRequset)(nil),    // 4: MapRequset
	(*Q1)(nil),            // 5: Q1
	nil,                   // 6: MapRequset.ISEntry
	nil,                   // 7: MapRequset.SSEntry
	nil,                   // 8: MapRequset.SItemEntry
	(*Q1_Q2)(nil),         // 9: Q1.Q2
}
var file_type_proto_depIdxs = []int32{
	2, // 0: ArraryRequest.item_list:type_name -> Item
	6, // 1: MapRequset.i_s:type_name -> MapRequset.ISEntry
	7, // 2: MapRequset.s_s:type_name -> MapRequset.SSEntry
	8, // 3: MapRequset.s_item:type_name -> MapRequset.SItemEntry
	9, // 4: Q1.q2:type_name -> Q1.Q2
	2, // 5: MapRequset.SItemEntry.value:type_name -> Item
	0, // 6: TypeService.Say:input_type -> Request
	1, // 7: TypeService.Say:output_type -> Response
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_type_proto_init() }
func file_type_proto_init() {
	if File_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_type_proto_rawDesc), len(file_type_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_type_proto_goTypes,
		DependencyIndexes: file_type_proto_depIdxs,
		MessageInfos:      file_type_proto_msgTypes,
	}.Build()
	File_type_proto = out.File
	file_type_proto_goTypes = nil
	file_type_proto_depIdxs = nil
}
