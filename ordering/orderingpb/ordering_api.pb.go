// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ordering_api.proto

package orderingpb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PaymentId string  `protobuf:"bytes,3,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Items     []*Item `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	Status    string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId     string  `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ProductId   string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	StoreName   string  `protobuf:"bytes,3,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	ProductName string  `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Price       float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Quantity    int32   `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_ordering_api_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Item) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *Item) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Item) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items     []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	UserId    string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PaymentId string  `protobuf:"bytes,3,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type CancelOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelOrderRequest) Reset() {
	*x = CancelOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRequest) ProtoMessage() {}

func (x *CancelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{6}
}

func (x *CancelOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelOrderResponse) Reset() {
	*x = CancelOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderResponse) ProtoMessage() {}

func (x *CancelOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderResponse.ProtoReflect.Descriptor instead.
func (*CancelOrderResponse) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{7}
}

type ReadyOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadyOrderRequest) Reset() {
	*x = ReadyOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyOrderRequest) ProtoMessage() {}

func (x *ReadyOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyOrderRequest.ProtoReflect.Descriptor instead.
func (*ReadyOrderRequest) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{8}
}

func (x *ReadyOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadyOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadyOrderResponse) Reset() {
	*x = ReadyOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyOrderResponse) ProtoMessage() {}

func (x *ReadyOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyOrderResponse.ProtoReflect.Descriptor instead.
func (*ReadyOrderResponse) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{9}
}

type CompleteOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InvoiceId string `protobuf:"bytes,2,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
}

func (x *CompleteOrderRequest) Reset() {
	*x = CompleteOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteOrderRequest) ProtoMessage() {}

func (x *CompleteOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteOrderRequest.ProtoReflect.Descriptor instead.
func (*CompleteOrderRequest) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{10}
}

func (x *CompleteOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompleteOrderRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

type CompleteOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompleteOrderResponse) Reset() {
	*x = CompleteOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteOrderResponse) ProtoMessage() {}

func (x *CompleteOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteOrderResponse.ProtoReflect.Descriptor instead.
func (*CompleteOrderResponse) Descriptor() ([]byte, []int) {
	return file_ordering_api_proto_rawDescGZIP(), []int{11}
}

var File_ordering_api_proto protoreflect.FileDescriptor

var file_ordering_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f,
	0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xb4, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x74, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x25, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x14,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfa, 0x03, 0x0a,
	0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xa4, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x54, 0x92, 0x41, 0x36, 0x12, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e,
	0x65, 0x77, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x22, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x1e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x9c, 0x01, 0x92, 0x41, 0x5f, 0x12,
	0x5d, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x20, 0x41, 0x50, 0x49, 0x22, 0x40, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x65, 0x7a, 0x61, 0x41, 0x6d, 0x69, 0x72, 0x69, 0x31, 0x32, 0x33, 0x2f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x7a, 0x61, 0x41,
	0x6d, 0x69, 0x72, 0x69, 0x31, 0x32, 0x33, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ordering_api_proto_rawDescOnce sync.Once
	file_ordering_api_proto_rawDescData = file_ordering_api_proto_rawDesc
)

func file_ordering_api_proto_rawDescGZIP() []byte {
	file_ordering_api_proto_rawDescOnce.Do(func() {
		file_ordering_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ordering_api_proto_rawDescData)
	})
	return file_ordering_api_proto_rawDescData
}

var file_ordering_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_ordering_api_proto_goTypes = []interface{}{
	(*Order)(nil),                 // 0: orderingpb.Order
	(*Item)(nil),                  // 1: orderingpb.Item
	(*CreateOrderRequest)(nil),    // 2: orderingpb.CreateOrderRequest
	(*CreateOrderResponse)(nil),   // 3: orderingpb.CreateOrderResponse
	(*GetOrderRequest)(nil),       // 4: orderingpb.GetOrderRequest
	(*GetOrderResponse)(nil),      // 5: orderingpb.GetOrderResponse
	(*CancelOrderRequest)(nil),    // 6: orderingpb.CancelOrderRequest
	(*CancelOrderResponse)(nil),   // 7: orderingpb.CancelOrderResponse
	(*ReadyOrderRequest)(nil),     // 8: orderingpb.ReadyOrderRequest
	(*ReadyOrderResponse)(nil),    // 9: orderingpb.ReadyOrderResponse
	(*CompleteOrderRequest)(nil),  // 10: orderingpb.CompleteOrderRequest
	(*CompleteOrderResponse)(nil), // 11: orderingpb.CompleteOrderResponse
}
var file_ordering_api_proto_depIdxs = []int32{
	1,  // 0: orderingpb.Order.items:type_name -> orderingpb.Item
	1,  // 1: orderingpb.CreateOrderRequest.items:type_name -> orderingpb.Item
	0,  // 2: orderingpb.GetOrderResponse.order:type_name -> orderingpb.Order
	2,  // 3: orderingpb.OrderingService.CreateOrder:input_type -> orderingpb.CreateOrderRequest
	4,  // 4: orderingpb.OrderingService.GetOrder:input_type -> orderingpb.GetOrderRequest
	6,  // 5: orderingpb.OrderingService.CancelOrder:input_type -> orderingpb.CancelOrderRequest
	8,  // 6: orderingpb.OrderingService.ReadyOrder:input_type -> orderingpb.ReadyOrderRequest
	10, // 7: orderingpb.OrderingService.CompleteOrder:input_type -> orderingpb.CompleteOrderRequest
	3,  // 8: orderingpb.OrderingService.CreateOrder:output_type -> orderingpb.CreateOrderResponse
	5,  // 9: orderingpb.OrderingService.GetOrder:output_type -> orderingpb.GetOrderResponse
	7,  // 10: orderingpb.OrderingService.CancelOrder:output_type -> orderingpb.CancelOrderResponse
	9,  // 11: orderingpb.OrderingService.ReadyOrder:output_type -> orderingpb.ReadyOrderResponse
	11, // 12: orderingpb.OrderingService.CompleteOrder:output_type -> orderingpb.CompleteOrderResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_ordering_api_proto_init() }
func file_ordering_api_proto_init() {
	if File_ordering_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ordering_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_ordering_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_ordering_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_ordering_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_ordering_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderRequest); i {
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
		file_ordering_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderResponse); i {
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
		file_ordering_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderRequest); i {
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
		file_ordering_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderResponse); i {
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
		file_ordering_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyOrderRequest); i {
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
		file_ordering_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyOrderResponse); i {
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
		file_ordering_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteOrderRequest); i {
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
		file_ordering_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteOrderResponse); i {
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
			RawDescriptor: file_ordering_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ordering_api_proto_goTypes,
		DependencyIndexes: file_ordering_api_proto_depIdxs,
		MessageInfos:      file_ordering_api_proto_msgTypes,
	}.Build()
	File_ordering_api_proto = out.File
	file_ordering_api_proto_rawDesc = nil
	file_ordering_api_proto_goTypes = nil
	file_ordering_api_proto_depIdxs = nil
}
