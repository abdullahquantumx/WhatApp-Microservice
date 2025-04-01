// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.20.3
// source: proto/whatapp.proto

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

// SendTemplateMessageRequest contains parameters for sending a template message
type SendTemplateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string            `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`                                                                    // Phone number of the recipient (with or without WhatsApp prefix)
	TemplateId  string            `protobuf:"bytes,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`                                                                       // ID of the template to use
	Parameters  map[string]string `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Template parameters
	OrderId     string            `protobuf:"bytes,4,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                                                                                // Optional: Order ID for tracking
	CustomerId  string            `protobuf:"bytes,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`                                                                       // Optional: Customer ID for tracking
}

func (x *SendTemplateMessageRequest) Reset() {
	*x = SendTemplateMessageRequest{}
	mi := &file_proto_whatapp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendTemplateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTemplateMessageRequest) ProtoMessage() {}

func (x *SendTemplateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatapp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTemplateMessageRequest.ProtoReflect.Descriptor instead.
func (*SendTemplateMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatapp_proto_rawDescGZIP(), []int{0}
}

func (x *SendTemplateMessageRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *SendTemplateMessageRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendTemplateMessageRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *SendTemplateMessageRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *SendTemplateMessageRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

// SendTemplateMessageResponse contains the result of sending a template message
type SendTemplateMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId  int64  `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`   // Internal message ID
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`                           // Status of the message (queued, sending, sent, delivered, read, failed)
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"` // External ID from the WhatsApp provider (if available)
}

func (x *SendTemplateMessageResponse) Reset() {
	*x = SendTemplateMessageResponse{}
	mi := &file_proto_whatapp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendTemplateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTemplateMessageResponse) ProtoMessage() {}

func (x *SendTemplateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatapp_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTemplateMessageResponse.ProtoReflect.Descriptor instead.
func (*SendTemplateMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_whatapp_proto_rawDescGZIP(), []int{1}
}

func (x *SendTemplateMessageResponse) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *SendTemplateMessageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SendTemplateMessageResponse) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

// GetMessageRequest contains parameters for retrieving a message
type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64 `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"` // Internal message ID
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	mi := &file_proto_whatapp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatapp_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatapp_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessageRequest) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

// MessageResponse contains details of a message
type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                                                                        // Internal message ID
	PhoneNumber  string            `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`                                                                    // Phone number of the recipient
	TemplateId   string            `protobuf:"bytes,3,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`                                                                       // ID of the template used
	Parameters   map[string]string `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Template parameters
	OrderId      string            `protobuf:"bytes,5,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                                                                                // Order ID for tracking
	CustomerId   string            `protobuf:"bytes,6,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`                                                                       // Customer ID for tracking
	Status       string            `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`                                                                                                 // Status of the message
	ErrorMessage string            `protobuf:"bytes,8,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`                                                                 // Error message (if any)
	ExternalId   string            `protobuf:"bytes,9,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`                                                                       // External ID from the WhatsApp provider
	CreatedAt    string            `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                                                                         // Creation timestamp in RFC3339 format
	UpdatedAt    string            `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`                                                                         // Last update timestamp in RFC3339 format
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	mi := &file_proto_whatapp_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatapp_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_whatapp_proto_rawDescGZIP(), []int{3}
}

func (x *MessageResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *MessageResponse) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *MessageResponse) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *MessageResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *MessageResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MessageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MessageResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *MessageResponse) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *MessageResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MessageResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// ListMessagesRequest contains parameters for listing messages
type ListMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`             // Optional: Filter by order ID
	CustomerId  string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`    // Optional: Filter by customer ID
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"` // Optional: Filter by phone number
	Limit       int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`                               // Maximum number of records to return
	Offset      int32  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`                             // Offset for pagination
}

func (x *ListMessagesRequest) Reset() {
	*x = ListMessagesRequest{}
	mi := &file_proto_whatapp_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesRequest) ProtoMessage() {}

func (x *ListMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatapp_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListMessagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatapp_proto_rawDescGZIP(), []int{4}
}

func (x *ListMessagesRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ListMessagesRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *ListMessagesRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *ListMessagesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListMessagesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// ListMessagesResponse contains a list of messages
type ListMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages   []*MessageResponse `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`                        // List of messages
	TotalCount int32              `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"` // Total number of messages matching the filters
}

func (x *ListMessagesResponse) Reset() {
	*x = ListMessagesResponse{}
	mi := &file_proto_whatapp_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesResponse) ProtoMessage() {}

func (x *ListMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatapp_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesResponse.ProtoReflect.Descriptor instead.
func (*ListMessagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_whatapp_proto_rawDescGZIP(), []int{5}
}

func (x *ListMessagesResponse) GetMessages() []*MessageResponse {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ListMessagesResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

// WebhookRequest contains data about a webhook event from WhatsApp provider
type WebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalId   string `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`       // External message ID
	Status       string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`                                 // New status of the message
	ErrorCode    string `protobuf:"bytes,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`          // Error code (if applicable)
	ErrorMessage string `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // Error message (if applicable)
}

func (x *WebhookRequest) Reset() {
	*x = WebhookRequest{}
	mi := &file_proto_whatapp_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookRequest) ProtoMessage() {}

func (x *WebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatapp_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookRequest.ProtoReflect.Descriptor instead.
func (*WebhookRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatapp_proto_rawDescGZIP(), []int{6}
}

func (x *WebhookRequest) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *WebhookRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WebhookRequest) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *WebhookRequest) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// WebhookResponse contains the result of processing a webhook
type WebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Whether the webhook was processed successfully
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // Additional information
}

func (x *WebhookResponse) Reset() {
	*x = WebhookResponse{}
	mi := &file_proto_whatapp_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookResponse) ProtoMessage() {}

func (x *WebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatapp_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookResponse.ProtoReflect.Descriptor instead.
func (*WebhookResponse) Descriptor() ([]byte, []int) {
	return file_proto_whatapp_proto_rawDescGZIP(), []int{7}
}

func (x *WebhookResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *WebhookResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_whatapp_proto protoreflect.FileDescriptor

var file_proto_whatapp_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x68, 0x61, 0x74, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x22,
	0xb1, 0x02, 0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x54, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70,
	0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x75, 0x0a, 0x1b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xc7,
	0x03, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x77, 0x68, 0x61,
	0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa2, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x6e, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61,
	0x70, 0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8d, 0x01,
	0x0a, 0x0e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a,
	0x0f, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x90, 0x02, 0x0a, 0x0f, 0x57, 0x68, 0x61, 0x74, 0x73, 0x41, 0x70,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x24, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x77,
	0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x68, 0x61, 0x74,
	0x73, 0x61, 0x70, 0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70,
	0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_whatapp_proto_rawDescOnce sync.Once
	file_proto_whatapp_proto_rawDescData = file_proto_whatapp_proto_rawDesc
)

func file_proto_whatapp_proto_rawDescGZIP() []byte {
	file_proto_whatapp_proto_rawDescOnce.Do(func() {
		file_proto_whatapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_whatapp_proto_rawDescData)
	})
	return file_proto_whatapp_proto_rawDescData
}

var file_proto_whatapp_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_whatapp_proto_goTypes = []any{
	(*SendTemplateMessageRequest)(nil),  // 0: whatsapp.SendTemplateMessageRequest
	(*SendTemplateMessageResponse)(nil), // 1: whatsapp.SendTemplateMessageResponse
	(*GetMessageRequest)(nil),           // 2: whatsapp.GetMessageRequest
	(*MessageResponse)(nil),             // 3: whatsapp.MessageResponse
	(*ListMessagesRequest)(nil),         // 4: whatsapp.ListMessagesRequest
	(*ListMessagesResponse)(nil),        // 5: whatsapp.ListMessagesResponse
	(*WebhookRequest)(nil),              // 6: whatsapp.WebhookRequest
	(*WebhookResponse)(nil),             // 7: whatsapp.WebhookResponse
	nil,                                 // 8: whatsapp.SendTemplateMessageRequest.ParametersEntry
	nil,                                 // 9: whatsapp.MessageResponse.ParametersEntry
}
var file_proto_whatapp_proto_depIdxs = []int32{
	8, // 0: whatsapp.SendTemplateMessageRequest.parameters:type_name -> whatsapp.SendTemplateMessageRequest.ParametersEntry
	9, // 1: whatsapp.MessageResponse.parameters:type_name -> whatsapp.MessageResponse.ParametersEntry
	3, // 2: whatsapp.ListMessagesResponse.messages:type_name -> whatsapp.MessageResponse
	0, // 3: whatsapp.WhatsAppService.SendTemplateMessage:input_type -> whatsapp.SendTemplateMessageRequest
	2, // 4: whatsapp.WhatsAppService.GetMessage:input_type -> whatsapp.GetMessageRequest
	4, // 5: whatsapp.WhatsAppService.ListMessages:input_type -> whatsapp.ListMessagesRequest
	1, // 6: whatsapp.WhatsAppService.SendTemplateMessage:output_type -> whatsapp.SendTemplateMessageResponse
	3, // 7: whatsapp.WhatsAppService.GetMessage:output_type -> whatsapp.MessageResponse
	5, // 8: whatsapp.WhatsAppService.ListMessages:output_type -> whatsapp.ListMessagesResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_whatapp_proto_init() }
func file_proto_whatapp_proto_init() {
	if File_proto_whatapp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_whatapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_whatapp_proto_goTypes,
		DependencyIndexes: file_proto_whatapp_proto_depIdxs,
		MessageInfos:      file_proto_whatapp_proto_msgTypes,
	}.Build()
	File_proto_whatapp_proto = out.File
	file_proto_whatapp_proto_rawDesc = nil
	file_proto_whatapp_proto_goTypes = nil
	file_proto_whatapp_proto_depIdxs = nil
}
