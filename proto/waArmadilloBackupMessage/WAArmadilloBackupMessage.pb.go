// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: waArmadilloBackupMessage/WAArmadilloBackupMessage.proto

package waArmadilloBackupMessage

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BackupMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*BackupMessage_EncryptedTransportMessage
	//	*BackupMessage_EncryptedTransportEvent
	//	*BackupMessage_EncryptedTransportLocallyTransformedMessage
	Payload  isBackupMessage_Payload `protobuf_oneof:"payload"`
	Metadata *BackupMessage_Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
}

func (x *BackupMessage) Reset() {
	*x = BackupMessage{}
	mi := &file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupMessage) ProtoMessage() {}

func (x *BackupMessage) ProtoReflect() protoreflect.Message {
	mi := &file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupMessage.ProtoReflect.Descriptor instead.
func (*BackupMessage) Descriptor() ([]byte, []int) {
	return file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescGZIP(), []int{0}
}

func (m *BackupMessage) GetPayload() isBackupMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *BackupMessage) GetEncryptedTransportMessage() []byte {
	if x, ok := x.GetPayload().(*BackupMessage_EncryptedTransportMessage); ok {
		return x.EncryptedTransportMessage
	}
	return nil
}

func (x *BackupMessage) GetEncryptedTransportEvent() *BackupMessage_Subprotocol {
	if x, ok := x.GetPayload().(*BackupMessage_EncryptedTransportEvent); ok {
		return x.EncryptedTransportEvent
	}
	return nil
}

func (x *BackupMessage) GetEncryptedTransportLocallyTransformedMessage() *BackupMessage_Subprotocol {
	if x, ok := x.GetPayload().(*BackupMessage_EncryptedTransportLocallyTransformedMessage); ok {
		return x.EncryptedTransportLocallyTransformedMessage
	}
	return nil
}

func (x *BackupMessage) GetMetadata() *BackupMessage_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type isBackupMessage_Payload interface {
	isBackupMessage_Payload()
}

type BackupMessage_EncryptedTransportMessage struct {
	EncryptedTransportMessage []byte `protobuf:"bytes,2,opt,name=encryptedTransportMessage,oneof"`
}

type BackupMessage_EncryptedTransportEvent struct {
	EncryptedTransportEvent *BackupMessage_Subprotocol `protobuf:"bytes,5,opt,name=encryptedTransportEvent,oneof"`
}

type BackupMessage_EncryptedTransportLocallyTransformedMessage struct {
	EncryptedTransportLocallyTransformedMessage *BackupMessage_Subprotocol `protobuf:"bytes,6,opt,name=encryptedTransportLocallyTransformedMessage,oneof"`
}

func (*BackupMessage_EncryptedTransportMessage) isBackupMessage_Payload() {}

func (*BackupMessage_EncryptedTransportEvent) isBackupMessage_Payload() {}

func (*BackupMessage_EncryptedTransportLocallyTransformedMessage) isBackupMessage_Payload() {}

type BackupMessage_Subprotocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	Version *int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (x *BackupMessage_Subprotocol) Reset() {
	*x = BackupMessage_Subprotocol{}
	mi := &file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupMessage_Subprotocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupMessage_Subprotocol) ProtoMessage() {}

func (x *BackupMessage_Subprotocol) ProtoReflect() protoreflect.Message {
	mi := &file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupMessage_Subprotocol.ProtoReflect.Descriptor instead.
func (*BackupMessage_Subprotocol) Descriptor() ([]byte, []int) {
	return file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BackupMessage_Subprotocol) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *BackupMessage_Subprotocol) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

type BackupMessage_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderID            *string                                  `protobuf:"bytes,1,opt,name=senderID" json:"senderID,omitempty"`
	MessageID           *string                                  `protobuf:"bytes,2,opt,name=messageID" json:"messageID,omitempty"`
	TimestampMS         *int64                                   `protobuf:"varint,3,opt,name=timestampMS" json:"timestampMS,omitempty"`
	FrankingMetadata    *BackupMessage_Metadata_FrankingMetadata `protobuf:"bytes,4,opt,name=frankingMetadata" json:"frankingMetadata,omitempty"`
	PayloadVersion      *int32                                   `protobuf:"varint,5,opt,name=payloadVersion" json:"payloadVersion,omitempty"`
	FutureProofBehavior *int32                                   `protobuf:"varint,6,opt,name=futureProofBehavior" json:"futureProofBehavior,omitempty"`
	ThreadTypeTag       *int32                                   `protobuf:"varint,7,opt,name=threadTypeTag" json:"threadTypeTag,omitempty"`
}

func (x *BackupMessage_Metadata) Reset() {
	*x = BackupMessage_Metadata{}
	mi := &file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupMessage_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupMessage_Metadata) ProtoMessage() {}

func (x *BackupMessage_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupMessage_Metadata.ProtoReflect.Descriptor instead.
func (*BackupMessage_Metadata) Descriptor() ([]byte, []int) {
	return file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescGZIP(), []int{0, 1}
}

func (x *BackupMessage_Metadata) GetSenderID() string {
	if x != nil && x.SenderID != nil {
		return *x.SenderID
	}
	return ""
}

func (x *BackupMessage_Metadata) GetMessageID() string {
	if x != nil && x.MessageID != nil {
		return *x.MessageID
	}
	return ""
}

func (x *BackupMessage_Metadata) GetTimestampMS() int64 {
	if x != nil && x.TimestampMS != nil {
		return *x.TimestampMS
	}
	return 0
}

func (x *BackupMessage_Metadata) GetFrankingMetadata() *BackupMessage_Metadata_FrankingMetadata {
	if x != nil {
		return x.FrankingMetadata
	}
	return nil
}

func (x *BackupMessage_Metadata) GetPayloadVersion() int32 {
	if x != nil && x.PayloadVersion != nil {
		return *x.PayloadVersion
	}
	return 0
}

func (x *BackupMessage_Metadata) GetFutureProofBehavior() int32 {
	if x != nil && x.FutureProofBehavior != nil {
		return *x.FutureProofBehavior
	}
	return 0
}

func (x *BackupMessage_Metadata) GetThreadTypeTag() int32 {
	if x != nil && x.ThreadTypeTag != nil {
		return *x.ThreadTypeTag
	}
	return 0
}

type BackupMessage_Metadata_FrankingMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrankingTag  []byte `protobuf:"bytes,3,opt,name=frankingTag" json:"frankingTag,omitempty"`
	ReportingTag []byte `protobuf:"bytes,4,opt,name=reportingTag" json:"reportingTag,omitempty"`
}

func (x *BackupMessage_Metadata_FrankingMetadata) Reset() {
	*x = BackupMessage_Metadata_FrankingMetadata{}
	mi := &file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupMessage_Metadata_FrankingMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupMessage_Metadata_FrankingMetadata) ProtoMessage() {}

func (x *BackupMessage_Metadata_FrankingMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupMessage_Metadata_FrankingMetadata.ProtoReflect.Descriptor instead.
func (*BackupMessage_Metadata_FrankingMetadata) Descriptor() ([]byte, []int) {
	return file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *BackupMessage_Metadata_FrankingMetadata) GetFrankingTag() []byte {
	if x != nil {
		return x.FrankingTag
	}
	return nil
}

func (x *BackupMessage_Metadata_FrankingMetadata) GetReportingTag() []byte {
	if x != nil {
		return x.ReportingTag
	}
	return nil
}

var File_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto protoreflect.FileDescriptor

//go:embed WAArmadilloBackupMessage.pb.raw
var file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDesc []byte

var (
	file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescOnce sync.Once
	file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescData = file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDesc
)

func file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescGZIP() []byte {
	file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescOnce.Do(func() {
		file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescData)
	})
	return file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDescData
}

var file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_goTypes = []any{
	(*BackupMessage)(nil),                           // 0: WAArmadilloBackupMessage.BackupMessage
	(*BackupMessage_Subprotocol)(nil),               // 1: WAArmadilloBackupMessage.BackupMessage.Subprotocol
	(*BackupMessage_Metadata)(nil),                  // 2: WAArmadilloBackupMessage.BackupMessage.Metadata
	(*BackupMessage_Metadata_FrankingMetadata)(nil), // 3: WAArmadilloBackupMessage.BackupMessage.Metadata.FrankingMetadata
}
var file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_depIdxs = []int32{
	1, // 0: WAArmadilloBackupMessage.BackupMessage.encryptedTransportEvent:type_name -> WAArmadilloBackupMessage.BackupMessage.Subprotocol
	1, // 1: WAArmadilloBackupMessage.BackupMessage.encryptedTransportLocallyTransformedMessage:type_name -> WAArmadilloBackupMessage.BackupMessage.Subprotocol
	2, // 2: WAArmadilloBackupMessage.BackupMessage.metadata:type_name -> WAArmadilloBackupMessage.BackupMessage.Metadata
	3, // 3: WAArmadilloBackupMessage.BackupMessage.Metadata.frankingMetadata:type_name -> WAArmadilloBackupMessage.BackupMessage.Metadata.FrankingMetadata
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_init() }
func file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_init() {
	if File_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto != nil {
		return
	}
	file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes[0].OneofWrappers = []any{
		(*BackupMessage_EncryptedTransportMessage)(nil),
		(*BackupMessage_EncryptedTransportEvent)(nil),
		(*BackupMessage_EncryptedTransportLocallyTransformedMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_goTypes,
		DependencyIndexes: file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_depIdxs,
		MessageInfos:      file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_msgTypes,
	}.Build()
	File_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto = out.File
	file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_rawDesc = nil
	file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_goTypes = nil
	file_waArmadilloBackupMessage_WAArmadilloBackupMessage_proto_depIdxs = nil
}
