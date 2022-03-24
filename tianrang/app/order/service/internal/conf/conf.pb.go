// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: service/internal/conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProductType 饮品类型
type ProductType int32

const (
	ProductType_PRODUCT_TYPE_ZERO     ProductType = 0 // 默认
	ProductType_PRODUCT_TYPE_COFFEE   ProductType = 1
	ProductType_PRODUCT_TYPE_MILL_TEA ProductType = 2
)

// Enum value maps for ProductType.
var (
	ProductType_name = map[int32]string{
		0: "PRODUCT_TYPE_ZERO",
		1: "PRODUCT_TYPE_COFFEE",
		2: "PRODUCT_TYPE_MILL_TEA",
	}
	ProductType_value = map[string]int32{
		"PRODUCT_TYPE_ZERO":     0,
		"PRODUCT_TYPE_COFFEE":   1,
		"PRODUCT_TYPE_MILL_TEA": 2,
	}
)

func (x ProductType) Enum() *ProductType {
	p := new(ProductType)
	*p = x
	return p
}

func (x ProductType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_internal_conf_conf_proto_enumTypes[0].Descriptor()
}

func (ProductType) Type() protoreflect.EnumType {
	return &file_service_internal_conf_conf_proto_enumTypes[0]
}

func (x ProductType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductType.Descriptor instead.
func (ProductType) EnumDescriptor() ([]byte, []int) {
	return file_service_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

// ProductTypeRules 饮品类型
type ProductTypeRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule        string      `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	ProductType ProductType `protobuf:"varint,2,opt,name=product_type,json=productType,proto3,enum=ProductType" json:"product_type,omitempty"`
}

func (x *ProductTypeRule) Reset() {
	*x = ProductTypeRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_internal_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductTypeRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductTypeRule) ProtoMessage() {}

func (x *ProductTypeRule) ProtoReflect() protoreflect.Message {
	mi := &file_service_internal_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductTypeRule.ProtoReflect.Descriptor instead.
func (*ProductTypeRule) Descriptor() ([]byte, []int) {
	return file_service_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *ProductTypeRule) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *ProductTypeRule) GetProductType() ProductType {
	if x != nil {
		return x.ProductType
	}
	return ProductType_PRODUCT_TYPE_ZERO
}

type Rules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 订单最多饮品数量
	OrderMaxCnt int32 `protobuf:"varint,1,opt,name=order_max_cnt,json=orderMaxCnt,proto3" json:"order_max_cnt,omitempty"`
	// 单个饮品品种的限制
	ProductTypeRules map[string]*ProductTypeRule `protobuf:"bytes,2,rep,name=product_type_rules,json=productTypeRules,proto3" json:"product_type_rules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Rules) Reset() {
	*x = Rules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_internal_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rules) ProtoMessage() {}

func (x *Rules) ProtoReflect() protoreflect.Message {
	mi := &file_service_internal_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rules.ProtoReflect.Descriptor instead.
func (*Rules) Descriptor() ([]byte, []int) {
	return file_service_internal_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Rules) GetOrderMaxCnt() int32 {
	if x != nil {
		return x.OrderMaxCnt
	}
	return 0
}

func (x *Rules) GetProductTypeRules() map[string]*ProductTypeRule {
	if x != nil {
		return x.ProductTypeRules
	}
	return nil
}

// TODO 生产环境可以保存在 DB、加载到缓存中
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 饮品列表
	Products string `protobuf:"bytes,1,opt,name=products,proto3" json:"products,omitempty"`
	// 配料列表
	Ingredients string `protobuf:"bytes,2,opt,name=ingredients,proto3" json:"ingredients,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_internal_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_service_internal_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_service_internal_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetProducts() string {
	if x != nil {
		return x.Products
	}
	return ""
}

func (x *Data) GetIngredients() string {
	if x != nil {
		return x.Ingredients
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grpc *Server_Grpc `protobuf:"bytes,1,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_internal_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_service_internal_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_service_internal_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Server) GetGrpc() *Server_Grpc {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   *Data   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Rules  *Rules  `protobuf:"bytes,2,opt,name=rules,proto3" json:"rules,omitempty"`
	Server *Server `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_internal_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_service_internal_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_service_internal_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Config) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Config) GetRules() *Rules {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Config) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type Server_Grpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr    string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_Grpc) Reset() {
	*x = Server_Grpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_internal_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_Grpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_Grpc) ProtoMessage() {}

func (x *Server_Grpc) ProtoReflect() protoreflect.Message {
	mi := &file_service_internal_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_Grpc.ProtoReflect.Descriptor instead.
func (*Server_Grpc) Descriptor() ([]byte, []int) {
	return file_service_internal_conf_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Server_Grpc) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_Grpc) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_service_internal_conf_conf_proto protoreflect.FileDescriptor

var file_service_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x05, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61,
	0x78, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4d, 0x61, 0x78, 0x43, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x1a, 0x55, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x44, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x7b, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x4f, 0x0a,
	0x04, 0x47, 0x72, 0x70, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x62,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2a, 0x58, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x44,
	0x55, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x46, 0x46, 0x45, 0x45, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x5f, 0x54, 0x45, 0x41, 0x10, 0x02, 0x42, 0x26, 0x5a, 0x24,
	0x61, 0x70, 0x70, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b,
	0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_internal_conf_conf_proto_rawDescOnce sync.Once
	file_service_internal_conf_conf_proto_rawDescData = file_service_internal_conf_conf_proto_rawDesc
)

func file_service_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_service_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_service_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_internal_conf_conf_proto_rawDescData)
	})
	return file_service_internal_conf_conf_proto_rawDescData
}

var file_service_internal_conf_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_internal_conf_conf_proto_goTypes = []interface{}{
	(ProductType)(0),            // 0: ProductType
	(*ProductTypeRule)(nil),     // 1: ProductTypeRule
	(*Rules)(nil),               // 2: Rules
	(*Data)(nil),                // 3: Data
	(*Server)(nil),              // 4: Server
	(*Config)(nil),              // 5: Config
	nil,                         // 6: Rules.ProductTypeRulesEntry
	(*Server_Grpc)(nil),         // 7: Server.Grpc
	(*durationpb.Duration)(nil), // 8: google.protobuf.Duration
}
var file_service_internal_conf_conf_proto_depIdxs = []int32{
	0, // 0: ProductTypeRule.product_type:type_name -> ProductType
	6, // 1: Rules.product_type_rules:type_name -> Rules.ProductTypeRulesEntry
	7, // 2: Server.grpc:type_name -> Server.Grpc
	3, // 3: Config.data:type_name -> Data
	2, // 4: Config.rules:type_name -> Rules
	4, // 5: Config.server:type_name -> Server
	1, // 6: Rules.ProductTypeRulesEntry.value:type_name -> ProductTypeRule
	8, // 7: Server.Grpc.timeout:type_name -> google.protobuf.Duration
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_service_internal_conf_conf_proto_init() }
func file_service_internal_conf_conf_proto_init() {
	if File_service_internal_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_internal_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductTypeRule); i {
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
		file_service_internal_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rules); i {
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
		file_service_internal_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_service_internal_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_service_internal_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_service_internal_conf_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_Grpc); i {
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
			RawDescriptor: file_service_internal_conf_conf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_service_internal_conf_conf_proto_depIdxs,
		EnumInfos:         file_service_internal_conf_conf_proto_enumTypes,
		MessageInfos:      file_service_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_service_internal_conf_conf_proto = out.File
	file_service_internal_conf_conf_proto_rawDesc = nil
	file_service_internal_conf_conf_proto_goTypes = nil
	file_service_internal_conf_conf_proto_depIdxs = nil
}