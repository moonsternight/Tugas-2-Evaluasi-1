// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/payment.proto

package payment

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

type ProcessPaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Amount        float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessPaymentRequest) Reset() {
	*x = ProcessPaymentRequest{}
	mi := &file_proto_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentRequest) ProtoMessage() {}

func (x *ProcessPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentRequest.ProtoReflect.Descriptor instead.
func (*ProcessPaymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessPaymentRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ProcessPaymentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ProcessPaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // SUCCESS, FAILED
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessPaymentResponse) Reset() {
	*x = ProcessPaymentResponse{}
	mi := &file_proto_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentResponse) ProtoMessage() {}

func (x *ProcessPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentResponse.ProtoReflect.Descriptor instead.
func (*ProcessPaymentResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessPaymentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type RefundPaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefundPaymentRequest) Reset() {
	*x = RefundPaymentRequest{}
	mi := &file_proto_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefundPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundPaymentRequest) ProtoMessage() {}

func (x *RefundPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundPaymentRequest.ProtoReflect.Descriptor instead.
func (*RefundPaymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{2}
}

func (x *RefundPaymentRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type RefundPaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // REFUNDED
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefundPaymentResponse) Reset() {
	*x = RefundPaymentResponse{}
	mi := &file_proto_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefundPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundPaymentResponse) ProtoMessage() {}

func (x *RefundPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundPaymentResponse.ProtoReflect.Descriptor instead.
func (*RefundPaymentResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{3}
}

func (x *RefundPaymentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_payment_proto protoreflect.FileDescriptor

const file_proto_payment_proto_rawDesc = "" +
	"\n" +
	"\x13proto/payment.proto\x12\apayment\"J\n" +
	"\x15ProcessPaymentRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x01R\x06amount\"0\n" +
	"\x16ProcessPaymentResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\"1\n" +
	"\x14RefundPaymentRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"/\n" +
	"\x15RefundPaymentResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2\xb3\x01\n" +
	"\x0ePaymentService\x12Q\n" +
	"\x0eProcessPayment\x12\x1e.payment.ProcessPaymentRequest\x1a\x1f.payment.ProcessPaymentResponse\x12N\n" +
	"\rRefundPayment\x12\x1d.payment.RefundPaymentRequest\x1a\x1e.payment.RefundPaymentResponseB!Z\x1fordersaga/proto/payment;paymentb\x06proto3"

var (
	file_proto_payment_proto_rawDescOnce sync.Once
	file_proto_payment_proto_rawDescData []byte
)

func file_proto_payment_proto_rawDescGZIP() []byte {
	file_proto_payment_proto_rawDescOnce.Do(func() {
		file_proto_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_payment_proto_rawDesc), len(file_proto_payment_proto_rawDesc)))
	})
	return file_proto_payment_proto_rawDescData
}

var file_proto_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_payment_proto_goTypes = []any{
	(*ProcessPaymentRequest)(nil),  // 0: payment.ProcessPaymentRequest
	(*ProcessPaymentResponse)(nil), // 1: payment.ProcessPaymentResponse
	(*RefundPaymentRequest)(nil),   // 2: payment.RefundPaymentRequest
	(*RefundPaymentResponse)(nil),  // 3: payment.RefundPaymentResponse
}
var file_proto_payment_proto_depIdxs = []int32{
	0, // 0: payment.PaymentService.ProcessPayment:input_type -> payment.ProcessPaymentRequest
	2, // 1: payment.PaymentService.RefundPayment:input_type -> payment.RefundPaymentRequest
	1, // 2: payment.PaymentService.ProcessPayment:output_type -> payment.ProcessPaymentResponse
	3, // 3: payment.PaymentService.RefundPayment:output_type -> payment.RefundPaymentResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_payment_proto_init() }
func file_proto_payment_proto_init() {
	if File_proto_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_payment_proto_rawDesc), len(file_proto_payment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_payment_proto_goTypes,
		DependencyIndexes: file_proto_payment_proto_depIdxs,
		MessageInfos:      file_proto_payment_proto_msgTypes,
	}.Build()
	File_proto_payment_proto = out.File
	file_proto_payment_proto_goTypes = nil
	file_proto_payment_proto_depIdxs = nil
}
