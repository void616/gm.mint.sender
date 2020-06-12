// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.10.1
// source: mintsender_event.proto

package sender

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

// Sent is an event from the service notifying about a wallet sending transaction result
type Sent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`        // Success is true in case of success
	Error       string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`             // Error contains error descrition in case of failure
	Service     string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`         // Service name (to differentiate multiple requestors): 1..64
	Id          string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`                   // Unique request ID: 1..64
	PublicKey   string `protobuf:"bytes,5,opt,name=publicKey,proto3" json:"publicKey,omitempty"`     // Destination wallet address in Base58 (empty on failure)
	Token       string `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`             // GOLD or MNT (empty on failure)
	Amount      string `protobuf:"bytes,7,opt,name=amount,proto3" json:"amount,omitempty"`           // Token amount in major units: 1.234 (18 decimal places, empty on failure)
	Transaction string `protobuf:"bytes,8,opt,name=transaction,proto3" json:"transaction,omitempty"` // Transaction digest in Base58 (empty on failure)
}

func (x *Sent) Reset() {
	*x = Sent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintsender_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sent) ProtoMessage() {}

func (x *Sent) ProtoReflect() protoreflect.Message {
	mi := &file_mintsender_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sent.ProtoReflect.Descriptor instead.
func (*Sent) Descriptor() ([]byte, []int) {
	return file_mintsender_event_proto_rawDescGZIP(), []int{0}
}

func (x *Sent) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Sent) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Sent) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Sent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sent) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Sent) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Sent) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Sent) GetTransaction() string {
	if x != nil {
		return x.Transaction
	}
	return ""
}

// SentAck is a reply for Sent
type SentAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Success is true in case of success
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`      // Error contains error descrition in case of failure
}

func (x *SentAck) Reset() {
	*x = SentAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintsender_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SentAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentAck) ProtoMessage() {}

func (x *SentAck) ProtoReflect() protoreflect.Message {
	mi := &file_mintsender_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentAck.ProtoReflect.Descriptor instead.
func (*SentAck) Descriptor() ([]byte, []int) {
	return file_mintsender_event_proto_rawDescGZIP(), []int{1}
}

func (x *SentAck) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SentAck) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Approved is an event from the service notifying about a wallet approvement result
type Approved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`        // Success is true in case of success
	Error       string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`             // Error contains error descrition in case of failure
	Service     string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`         // Service name (to differentiate multiple requestors): 1..64
	Id          string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`                   // Unique request ID: 1..64
	PublicKey   string `protobuf:"bytes,5,opt,name=publicKey,proto3" json:"publicKey,omitempty"`     // Destination wallet address in Base58 (empty on failure)
	Transaction string `protobuf:"bytes,6,opt,name=transaction,proto3" json:"transaction,omitempty"` // Transaction digest in Base58 (empty on failure)
}

func (x *Approved) Reset() {
	*x = Approved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintsender_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Approved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Approved) ProtoMessage() {}

func (x *Approved) ProtoReflect() protoreflect.Message {
	mi := &file_mintsender_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Approved.ProtoReflect.Descriptor instead.
func (*Approved) Descriptor() ([]byte, []int) {
	return file_mintsender_event_proto_rawDescGZIP(), []int{2}
}

func (x *Approved) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Approved) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Approved) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Approved) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Approved) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Approved) GetTransaction() string {
	if x != nil {
		return x.Transaction
	}
	return ""
}

// ApprovedAck is a reply for Approved
type ApprovedAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Success is true in case of success
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`      // Error contains error descrition in case of failure
}

func (x *ApprovedAck) Reset() {
	*x = ApprovedAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mintsender_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovedAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovedAck) ProtoMessage() {}

func (x *ApprovedAck) ProtoReflect() protoreflect.Message {
	mi := &file_mintsender_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovedAck.ProtoReflect.Descriptor instead.
func (*ApprovedAck) Descriptor() ([]byte, []int) {
	return file_mintsender_event_proto_rawDescGZIP(), []int{3}
}

func (x *ApprovedAck) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ApprovedAck) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_mintsender_event_proto protoreflect.FileDescriptor

var file_mintsender_event_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x69, 0x6e, 0x74, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0xce, 0x01, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x39, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xa4, 0x01, 0x0a, 0x08,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x41, 0x63,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x42, 0x24, 0x5a, 0x08, 0x2e, 0x3b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0xaa, 0x02, 0x17,
	0x4d, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mintsender_event_proto_rawDescOnce sync.Once
	file_mintsender_event_proto_rawDescData = file_mintsender_event_proto_rawDesc
)

func file_mintsender_event_proto_rawDescGZIP() []byte {
	file_mintsender_event_proto_rawDescOnce.Do(func() {
		file_mintsender_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_mintsender_event_proto_rawDescData)
	})
	return file_mintsender_event_proto_rawDescData
}

var file_mintsender_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mintsender_event_proto_goTypes = []interface{}{
	(*Sent)(nil),        // 0: event.Sent
	(*SentAck)(nil),     // 1: event.SentAck
	(*Approved)(nil),    // 2: event.Approved
	(*ApprovedAck)(nil), // 3: event.ApprovedAck
}
var file_mintsender_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mintsender_event_proto_init() }
func file_mintsender_event_proto_init() {
	if File_mintsender_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mintsender_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sent); i {
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
		file_mintsender_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SentAck); i {
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
		file_mintsender_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Approved); i {
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
		file_mintsender_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovedAck); i {
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
			RawDescriptor: file_mintsender_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mintsender_event_proto_goTypes,
		DependencyIndexes: file_mintsender_event_proto_depIdxs,
		MessageInfos:      file_mintsender_event_proto_msgTypes,
	}.Build()
	File_mintsender_event_proto = out.File
	file_mintsender_event_proto_rawDesc = nil
	file_mintsender_event_proto_goTypes = nil
	file_mintsender_event_proto_depIdxs = nil
}
