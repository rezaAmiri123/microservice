// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: rpc_create_transfer.proto

package grpc

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

type CreateTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccountId string `protobuf:"bytes,1,opt,name=from_account_id,json=fromAccountId,proto3" json:"from_account_id,omitempty"`
	ToAccountId   string `protobuf:"bytes,2,opt,name=to_account_id,json=toAccountId,proto3" json:"to_account_id,omitempty"`
	Amount        int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreateTransferRequest) Reset() {
	*x = CreateTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransferRequest) ProtoMessage() {}

func (x *CreateTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransferRequest.ProtoReflect.Descriptor instead.
func (*CreateTransferRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTransferRequest) GetFromAccountId() string {
	if x != nil {
		return x.FromAccountId
	}
	return ""
}

func (x *CreateTransferRequest) GetToAccountId() string {
	if x != nil {
		return x.ToAccountId
	}
	return ""
}

func (x *CreateTransferRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transfer *Transfer `protobuf:"bytes,1,opt,name=transfer,proto3" json:"transfer,omitempty"`
}

func (x *CreateTransferResponse) Reset() {
	*x = CreateTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransferResponse) ProtoMessage() {}

func (x *CreateTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransferResponse.ProtoReflect.Descriptor instead.
func (*CreateTransferResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTransferResponse) GetTransfer() *Transfer {
	if x != nil {
		return x.Transfer
	}
	return nil
}

var File_rpc_create_transfer_proto protoreflect.FileDescriptor

var file_rpc_create_transfer_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x1a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x7a, 0x61, 0x41, 0x6d, 0x69, 0x72, 0x69, 0x31, 0x32, 0x33, 0x2f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_create_transfer_proto_rawDescOnce sync.Once
	file_rpc_create_transfer_proto_rawDescData = file_rpc_create_transfer_proto_rawDesc
)

func file_rpc_create_transfer_proto_rawDescGZIP() []byte {
	file_rpc_create_transfer_proto_rawDescOnce.Do(func() {
		file_rpc_create_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_transfer_proto_rawDescData)
	})
	return file_rpc_create_transfer_proto_rawDescData
}

var file_rpc_create_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_transfer_proto_goTypes = []interface{}{
	(*CreateTransferRequest)(nil),  // 0: grpc.CreateTransferRequest
	(*CreateTransferResponse)(nil), // 1: grpc.CreateTransferResponse
	(*Transfer)(nil),               // 2: grpc.Transfer
}
var file_rpc_create_transfer_proto_depIdxs = []int32{
	2, // 0: grpc.CreateTransferResponse.transfer:type_name -> grpc.Transfer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_create_transfer_proto_init() }
func file_rpc_create_transfer_proto_init() {
	if File_rpc_create_transfer_proto != nil {
		return
	}
	file_transfer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransferRequest); i {
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
		file_rpc_create_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransferResponse); i {
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
			RawDescriptor: file_rpc_create_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_transfer_proto_goTypes,
		DependencyIndexes: file_rpc_create_transfer_proto_depIdxs,
		MessageInfos:      file_rpc_create_transfer_proto_msgTypes,
	}.Build()
	File_rpc_create_transfer_proto = out.File
	file_rpc_create_transfer_proto_rawDesc = nil
	file_rpc_create_transfer_proto_goTypes = nil
	file_rpc_create_transfer_proto_depIdxs = nil
}
