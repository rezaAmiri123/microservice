// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: search_api.proto

package searchpb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string  `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	StoreId     string  `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ProductName string  `protobuf:"bytes,3,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	StoreName   string  `protobuf:"bytes,4,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	Price       float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Quantity    int64   `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_search_api_proto_msgTypes[0]
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
	return file_search_api_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Item) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Item) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Item) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *Item) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  string  `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId   string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string  `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Items    []*Item `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	Total    float64 `protobuf:"fixed64,5,opt,name=total,proto3" json:"total,omitempty"`
	Status   string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_search_api_proto_msgTypes[1]
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
	return file_search_api_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type SearchOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters *SearchOrdersRequest_Filters `protobuf:"bytes,1,opt,name=filters,proto3" json:"filters,omitempty"`
	Next    string                       `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
	Limit   int32                        `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SearchOrdersRequest) Reset() {
	*x = SearchOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOrdersRequest) ProtoMessage() {}

func (x *SearchOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOrdersRequest.ProtoReflect.Descriptor instead.
func (*SearchOrdersRequest) Descriptor() ([]byte, []int) {
	return file_search_api_proto_rawDescGZIP(), []int{2}
}

func (x *SearchOrdersRequest) GetFilters() *SearchOrdersRequest_Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *SearchOrdersRequest) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

func (x *SearchOrdersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	Next   string   `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
}

func (x *SearchOrdersResponse) Reset() {
	*x = SearchOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOrdersResponse) ProtoMessage() {}

func (x *SearchOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOrdersResponse.ProtoReflect.Descriptor instead.
func (*SearchOrdersResponse) Descriptor() ([]byte, []int) {
	return file_search_api_proto_rawDescGZIP(), []int{3}
}

func (x *SearchOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *SearchOrdersResponse) GetNext() string {
	if x != nil {
		return x.Next
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
		mi := &file_search_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_api_proto_msgTypes[4]
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
	return file_search_api_proto_rawDescGZIP(), []int{4}
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
		mi := &file_search_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_api_proto_msgTypes[5]
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
	return file_search_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type SearchOrdersRequest_Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	After      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
	Before     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=before,proto3" json:"before,omitempty"`
	StoreIds   []string               `protobuf:"bytes,4,rep,name=store_ids,json=storeIds,proto3" json:"store_ids,omitempty"`
	ProductIds []string               `protobuf:"bytes,5,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	MinTotal   float64                `protobuf:"fixed64,6,opt,name=min_total,json=minTotal,proto3" json:"min_total,omitempty"`
	MaxTotal   float64                `protobuf:"fixed64,7,opt,name=max_total,json=maxTotal,proto3" json:"max_total,omitempty"`
	Status     string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SearchOrdersRequest_Filters) Reset() {
	*x = SearchOrdersRequest_Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOrdersRequest_Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOrdersRequest_Filters) ProtoMessage() {}

func (x *SearchOrdersRequest_Filters) ProtoReflect() protoreflect.Message {
	mi := &file_search_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOrdersRequest_Filters.ProtoReflect.Descriptor instead.
func (*SearchOrdersRequest_Filters) Descriptor() ([]byte, []int) {
	return file_search_api_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SearchOrdersRequest_Filters) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SearchOrdersRequest_Filters) GetAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.After
	}
	return nil
}

func (x *SearchOrdersRequest_Filters) GetBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *SearchOrdersRequest_Filters) GetStoreIds() []string {
	if x != nil {
		return x.StoreIds
	}
	return nil
}

func (x *SearchOrdersRequest_Filters) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

func (x *SearchOrdersRequest_Filters) GetMinTotal() float64 {
	if x != nil {
		return x.MinTotal
	}
	return 0
}

func (x *SearchOrdersRequest_Filters) GetMaxTotal() float64 {
	if x != nil {
		return x.MaxTotal
	}
	return 0
}

func (x *SearchOrdersRequest_Filters) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_search_api_proto protoreflect.FileDescriptor

var file_search_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a,
	0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0xac, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x9c, 0x03, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x98, 0x02, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x54, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0xcf, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa9, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x92, 0x41, 0x2e, 0x12,
	0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x1d,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x91, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x1a, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x92, 0x41, 0x26, 0x12,
	0x09, 0x47, 0x65, 0x74, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x19, 0x55, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x67,
	0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x94, 0x01, 0x92, 0x41, 0x5b, 0x12, 0x59,
	0x0a, 0x12, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x20, 0x41, 0x50, 0x49, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x7a, 0x61, 0x41,
	0x6d, 0x69, 0x72, 0x69, 0x31, 0x32, 0x33, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x31, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x7a, 0x61, 0x41, 0x6d, 0x69, 0x72, 0x69, 0x31,
	0x32, 0x33, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_api_proto_rawDescOnce sync.Once
	file_search_api_proto_rawDescData = file_search_api_proto_rawDesc
)

func file_search_api_proto_rawDescGZIP() []byte {
	file_search_api_proto_rawDescOnce.Do(func() {
		file_search_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_api_proto_rawDescData)
	})
	return file_search_api_proto_rawDescData
}

var file_search_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_search_api_proto_goTypes = []interface{}{
	(*Item)(nil),                        // 0: basketspb.Item
	(*Order)(nil),                       // 1: basketspb.Order
	(*SearchOrdersRequest)(nil),         // 2: basketspb.SearchOrdersRequest
	(*SearchOrdersResponse)(nil),        // 3: basketspb.SearchOrdersResponse
	(*GetOrderRequest)(nil),             // 4: basketspb.GetOrderRequest
	(*GetOrderResponse)(nil),            // 5: basketspb.GetOrderResponse
	(*SearchOrdersRequest_Filters)(nil), // 6: basketspb.SearchOrdersRequest.Filters
	(*timestamppb.Timestamp)(nil),       // 7: google.protobuf.Timestamp
}
var file_search_api_proto_depIdxs = []int32{
	0, // 0: basketspb.Order.items:type_name -> basketspb.Item
	6, // 1: basketspb.SearchOrdersRequest.filters:type_name -> basketspb.SearchOrdersRequest.Filters
	1, // 2: basketspb.SearchOrdersResponse.orders:type_name -> basketspb.Order
	1, // 3: basketspb.GetOrderResponse.order:type_name -> basketspb.Order
	7, // 4: basketspb.SearchOrdersRequest.Filters.after:type_name -> google.protobuf.Timestamp
	7, // 5: basketspb.SearchOrdersRequest.Filters.before:type_name -> google.protobuf.Timestamp
	2, // 6: basketspb.SearchService.SearchOrders:input_type -> basketspb.SearchOrdersRequest
	4, // 7: basketspb.SearchService.GetOrder:input_type -> basketspb.GetOrderRequest
	3, // 8: basketspb.SearchService.SearchOrders:output_type -> basketspb.SearchOrdersResponse
	5, // 9: basketspb.SearchService.GetOrder:output_type -> basketspb.GetOrderResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_search_api_proto_init() }
func file_search_api_proto_init() {
	if File_search_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_search_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_search_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOrdersRequest); i {
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
		file_search_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOrdersResponse); i {
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
		file_search_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_search_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_search_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOrdersRequest_Filters); i {
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
			RawDescriptor: file_search_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_api_proto_goTypes,
		DependencyIndexes: file_search_api_proto_depIdxs,
		MessageInfos:      file_search_api_proto_msgTypes,
	}.Build()
	File_search_api_proto = out.File
	file_search_api_proto_rawDesc = nil
	file_search_api_proto_goTypes = nil
	file_search_api_proto_depIdxs = nil
}
