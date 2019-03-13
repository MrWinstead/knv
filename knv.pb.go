// Code generated by protoc-gen-go. DO NOT EDIT.
// source: knv.proto

package knv

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DatabaseItem_State int32

const (
	DatabaseItem_UNKNOWN   DatabaseItem_State = 0
	DatabaseItem_NOT_FOUND DatabaseItem_State = 1
)

var DatabaseItem_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "NOT_FOUND",
}

var DatabaseItem_State_value = map[string]int32{
	"UNKNOWN":   0,
	"NOT_FOUND": 1,
}

func (x DatabaseItem_State) String() string {
	return proto.EnumName(DatabaseItem_State_name, int32(x))
}

func (DatabaseItem_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{1, 0}
}

type DatabaseKey struct {
	Table                string   `protobuf:"bytes,1,opt,name=Table,proto3" json:"Table,omitempty"`
	Index                string   `protobuf:"bytes,2,opt,name=Index,proto3" json:"Index,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseKey) Reset()         { *m = DatabaseKey{} }
func (m *DatabaseKey) String() string { return proto.CompactTextString(m) }
func (*DatabaseKey) ProtoMessage()    {}
func (*DatabaseKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{0}
}

func (m *DatabaseKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseKey.Unmarshal(m, b)
}
func (m *DatabaseKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseKey.Marshal(b, m, deterministic)
}
func (m *DatabaseKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseKey.Merge(m, src)
}
func (m *DatabaseKey) XXX_Size() int {
	return xxx_messageInfo_DatabaseKey.Size(m)
}
func (m *DatabaseKey) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseKey.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseKey proto.InternalMessageInfo

func (m *DatabaseKey) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *DatabaseKey) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *DatabaseKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DatabaseItem struct {
	State                DatabaseItem_State `protobuf:"varint,1,opt,name=state,proto3,enum=knv.DatabaseItem_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DatabaseItem) Reset()         { *m = DatabaseItem{} }
func (m *DatabaseItem) String() string { return proto.CompactTextString(m) }
func (*DatabaseItem) ProtoMessage()    {}
func (*DatabaseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{1}
}

func (m *DatabaseItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseItem.Unmarshal(m, b)
}
func (m *DatabaseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseItem.Marshal(b, m, deterministic)
}
func (m *DatabaseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseItem.Merge(m, src)
}
func (m *DatabaseItem) XXX_Size() int {
	return xxx_messageInfo_DatabaseItem.Size(m)
}
func (m *DatabaseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseItem.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseItem proto.InternalMessageInfo

func (m *DatabaseItem) GetState() DatabaseItem_State {
	if m != nil {
		return m.State
	}
	return DatabaseItem_UNKNOWN
}

type RequestHeader struct {
	RequestID            []byte   `protobuf:"bytes,1,opt,name=RequestID,proto3" json:"RequestID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{2}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetRequestID() []byte {
	if m != nil {
		return m.RequestID
	}
	return nil
}

type ResponseHeader struct {
	TransactionID        []byte   `protobuf:"bytes,2,opt,name=TransactionID,proto3" json:"TransactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{3}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetTransactionID() []byte {
	if m != nil {
		return m.TransactionID
	}
	return nil
}

type DatabaseGetRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Key                  *DatabaseKey   `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DatabaseGetRequest) Reset()         { *m = DatabaseGetRequest{} }
func (m *DatabaseGetRequest) String() string { return proto.CompactTextString(m) }
func (*DatabaseGetRequest) ProtoMessage()    {}
func (*DatabaseGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{4}
}

func (m *DatabaseGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseGetRequest.Unmarshal(m, b)
}
func (m *DatabaseGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseGetRequest.Marshal(b, m, deterministic)
}
func (m *DatabaseGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseGetRequest.Merge(m, src)
}
func (m *DatabaseGetRequest) XXX_Size() int {
	return xxx_messageInfo_DatabaseGetRequest.Size(m)
}
func (m *DatabaseGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseGetRequest proto.InternalMessageInfo

func (m *DatabaseGetRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DatabaseGetRequest) GetKey() *DatabaseKey {
	if m != nil {
		return m.Key
	}
	return nil
}

type DatabaseGetResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Item                 *DatabaseItem   `protobuf:"bytes,2,opt,name=Item,proto3" json:"Item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DatabaseGetResponse) Reset()         { *m = DatabaseGetResponse{} }
func (m *DatabaseGetResponse) String() string { return proto.CompactTextString(m) }
func (*DatabaseGetResponse) ProtoMessage()    {}
func (*DatabaseGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{5}
}

func (m *DatabaseGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseGetResponse.Unmarshal(m, b)
}
func (m *DatabaseGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseGetResponse.Marshal(b, m, deterministic)
}
func (m *DatabaseGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseGetResponse.Merge(m, src)
}
func (m *DatabaseGetResponse) XXX_Size() int {
	return xxx_messageInfo_DatabaseGetResponse.Size(m)
}
func (m *DatabaseGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseGetResponse proto.InternalMessageInfo

func (m *DatabaseGetResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DatabaseGetResponse) GetItem() *DatabaseItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type DatabaseDeleteRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DatabaseDeleteRequest) Reset()         { *m = DatabaseDeleteRequest{} }
func (m *DatabaseDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DatabaseDeleteRequest) ProtoMessage()    {}
func (*DatabaseDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{6}
}

func (m *DatabaseDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseDeleteRequest.Unmarshal(m, b)
}
func (m *DatabaseDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseDeleteRequest.Marshal(b, m, deterministic)
}
func (m *DatabaseDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseDeleteRequest.Merge(m, src)
}
func (m *DatabaseDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DatabaseDeleteRequest.Size(m)
}
func (m *DatabaseDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseDeleteRequest proto.InternalMessageInfo

func (m *DatabaseDeleteRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type DatabaseDeleteResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DatabaseDeleteResponse) Reset()         { *m = DatabaseDeleteResponse{} }
func (m *DatabaseDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DatabaseDeleteResponse) ProtoMessage()    {}
func (*DatabaseDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{7}
}

func (m *DatabaseDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseDeleteResponse.Unmarshal(m, b)
}
func (m *DatabaseDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseDeleteResponse.Marshal(b, m, deterministic)
}
func (m *DatabaseDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseDeleteResponse.Merge(m, src)
}
func (m *DatabaseDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DatabaseDeleteResponse.Size(m)
}
func (m *DatabaseDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseDeleteResponse proto.InternalMessageInfo

func (m *DatabaseDeleteResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type DatabasePutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabasePutRequest) Reset()         { *m = DatabasePutRequest{} }
func (m *DatabasePutRequest) String() string { return proto.CompactTextString(m) }
func (*DatabasePutRequest) ProtoMessage()    {}
func (*DatabasePutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{8}
}

func (m *DatabasePutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabasePutRequest.Unmarshal(m, b)
}
func (m *DatabasePutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabasePutRequest.Marshal(b, m, deterministic)
}
func (m *DatabasePutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabasePutRequest.Merge(m, src)
}
func (m *DatabasePutRequest) XXX_Size() int {
	return xxx_messageInfo_DatabasePutRequest.Size(m)
}
func (m *DatabasePutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabasePutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatabasePutRequest proto.InternalMessageInfo

type DatabasePutResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DatabasePutResponse) Reset()         { *m = DatabasePutResponse{} }
func (m *DatabasePutResponse) String() string { return proto.CompactTextString(m) }
func (*DatabasePutResponse) ProtoMessage()    {}
func (*DatabasePutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e323961b2790e61, []int{9}
}

func (m *DatabasePutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabasePutResponse.Unmarshal(m, b)
}
func (m *DatabasePutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabasePutResponse.Marshal(b, m, deterministic)
}
func (m *DatabasePutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabasePutResponse.Merge(m, src)
}
func (m *DatabasePutResponse) XXX_Size() int {
	return xxx_messageInfo_DatabasePutResponse.Size(m)
}
func (m *DatabasePutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabasePutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatabasePutResponse proto.InternalMessageInfo

func (m *DatabasePutResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterEnum("knv.DatabaseItem_State", DatabaseItem_State_name, DatabaseItem_State_value)
	proto.RegisterType((*DatabaseKey)(nil), "knv.DatabaseKey")
	proto.RegisterType((*DatabaseItem)(nil), "knv.DatabaseItem")
	proto.RegisterType((*RequestHeader)(nil), "knv.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "knv.ResponseHeader")
	proto.RegisterType((*DatabaseGetRequest)(nil), "knv.DatabaseGetRequest")
	proto.RegisterType((*DatabaseGetResponse)(nil), "knv.DatabaseGetResponse")
	proto.RegisterType((*DatabaseDeleteRequest)(nil), "knv.DatabaseDeleteRequest")
	proto.RegisterType((*DatabaseDeleteResponse)(nil), "knv.DatabaseDeleteResponse")
	proto.RegisterType((*DatabasePutRequest)(nil), "knv.DatabasePutRequest")
	proto.RegisterType((*DatabasePutResponse)(nil), "knv.DatabasePutResponse")
}

func init() { proto.RegisterFile("knv.proto", fileDescriptor_3e323961b2790e61) }

var fileDescriptor_3e323961b2790e61 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x14, 0xac, 0x71, 0xa1, 0xf5, 0xe3, 0x43, 0xee, 0x42, 0x5b, 0x8b, 0xf6, 0x50, 0x6d, 0x5b, 0xa9,
	0x6a, 0x05, 0x07, 0x2a, 0x55, 0x5c, 0x9b, 0x38, 0x21, 0x96, 0x25, 0x83, 0x36, 0xa0, 0x1c, 0xa3,
	0x75, 0x78, 0x07, 0x04, 0xb1, 0x09, 0x5e, 0x50, 0xf8, 0x79, 0xf9, 0x67, 0x91, 0x77, 0xd7, 0x31,
	0x36, 0x39, 0x71, 0xf3, 0x9b, 0x37, 0x9a, 0x79, 0x33, 0xb6, 0xc1, 0x5a, 0x46, 0xbb, 0xfe, 0x7a,
	0x13, 0x8b, 0x98, 0x98, 0xcb, 0x68, 0x47, 0x7d, 0xa8, 0xbb, 0x5c, 0xf0, 0x90, 0x27, 0xe8, 0xe3,
	0x9e, 0x74, 0xa0, 0x3a, 0xe5, 0xe1, 0x0a, 0x1d, 0xe3, 0x9b, 0xf1, 0xcb, 0x62, 0x6a, 0x48, 0x51,
	0x2f, 0x9a, 0xe3, 0xa3, 0x53, 0x51, 0xa8, 0x1c, 0x88, 0x0d, 0xa6, 0x8f, 0x7b, 0xc7, 0x94, 0x58,
	0xfa, 0x48, 0x43, 0x68, 0x64, 0x62, 0x9e, 0xc0, 0x7b, 0xd2, 0x83, 0x6a, 0x22, 0xb8, 0x50, 0x6a,
	0xad, 0xc1, 0xe7, 0x7e, 0x6a, 0x7e, 0xc8, 0xe8, 0x5f, 0xa7, 0x6b, 0xa6, 0x58, 0xf4, 0x3b, 0x54,
	0xe5, 0x4c, 0xea, 0xf0, 0x6e, 0x16, 0xf8, 0xc1, 0xf8, 0x26, 0xb0, 0xdf, 0x90, 0x26, 0x58, 0xc1,
	0x78, 0x7a, 0x7b, 0x39, 0x9e, 0x05, 0xae, 0x6d, 0xd0, 0x1e, 0x34, 0x19, 0x3e, 0x6c, 0x31, 0x11,
	0x57, 0xc8, 0xe7, 0xb8, 0x21, 0x5f, 0xc1, 0xd2, 0x80, 0xe7, 0x4a, 0xa3, 0x06, 0xcb, 0x01, 0xfa,
	0x0f, 0x5a, 0x0c, 0x93, 0x75, 0x1c, 0x25, 0xa8, 0xf9, 0x3f, 0xa0, 0x39, 0xdd, 0xf0, 0x28, 0xe1,
	0x77, 0x62, 0x11, 0x47, 0x9e, 0x2b, 0x43, 0x35, 0x58, 0x11, 0xa4, 0x73, 0x20, 0xd9, 0xa1, 0x23,
	0x14, 0x5a, 0x8f, 0xfc, 0x86, 0x9a, 0x52, 0x91, 0x46, 0xf5, 0x01, 0x91, 0x89, 0x0a, 0xf7, 0x30,
	0xcd, 0x20, 0x54, 0xd5, 0x53, 0x91, 0x44, 0xbb, 0x10, 0xdd, 0xc7, 0xbd, 0x2a, 0x6c, 0x01, 0xed,
	0x82, 0x8b, 0x3a, 0x94, 0xfc, 0x29, 0xd9, 0xb4, 0xb5, 0xcd, 0x61, 0x8e, 0x17, 0x9f, 0x9f, 0xf0,
	0x36, 0xad, 0x52, 0x1b, 0x7d, 0x38, 0xea, 0x98, 0xc9, 0x35, 0x3d, 0x87, 0x8f, 0x19, 0xea, 0xe2,
	0x0a, 0x05, 0x9e, 0x90, 0x89, 0x5e, 0xc0, 0xa7, 0xb2, 0xc8, 0x09, 0x27, 0xd3, 0x4e, 0x5e, 0xee,
	0x64, 0x9b, 0x95, 0x4b, 0xcf, 0xf2, 0x32, 0x24, 0x7a, 0x82, 0xf2, 0xe0, 0xc9, 0x80, 0xf7, 0x99,
	0x08, 0x19, 0x82, 0x39, 0x42, 0x41, 0x8a, 0x9f, 0x5d, 0xfe, 0x36, 0xbb, 0xce, 0xf1, 0x42, 0x7b,
	0xfe, 0x87, 0x9a, 0xca, 0x47, 0xba, 0x05, 0x4e, 0xa1, 0xb9, 0xee, 0x97, 0x57, 0x77, 0x5a, 0x62,
	0x08, 0xe6, 0x64, 0x5b, 0x36, 0xcf, 0xd3, 0x96, 0xcc, 0x0f, 0x02, 0x87, 0x35, 0xf9, 0x7b, 0xfe,
	0x7d, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xc5, 0x30, 0x33, 0xab, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseClient interface {
	Get(ctx context.Context, in *DatabaseGetRequest, opts ...grpc.CallOption) (*DatabaseGetResponse, error)
	Delete(ctx context.Context, in *DatabaseDeleteRequest, opts ...grpc.CallOption) (*DatabaseDeleteResponse, error)
	Put(ctx context.Context, in *DatabasePutRequest, opts ...grpc.CallOption) (*DatabasePutResponse, error)
}

type databaseClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseClient(cc *grpc.ClientConn) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Get(ctx context.Context, in *DatabaseGetRequest, opts ...grpc.CallOption) (*DatabaseGetResponse, error) {
	out := new(DatabaseGetResponse)
	err := c.cc.Invoke(ctx, "/knv.Database/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Delete(ctx context.Context, in *DatabaseDeleteRequest, opts ...grpc.CallOption) (*DatabaseDeleteResponse, error) {
	out := new(DatabaseDeleteResponse)
	err := c.cc.Invoke(ctx, "/knv.Database/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Put(ctx context.Context, in *DatabasePutRequest, opts ...grpc.CallOption) (*DatabasePutResponse, error) {
	out := new(DatabasePutResponse)
	err := c.cc.Invoke(ctx, "/knv.Database/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
type DatabaseServer interface {
	Get(context.Context, *DatabaseGetRequest) (*DatabaseGetResponse, error)
	Delete(context.Context, *DatabaseDeleteRequest) (*DatabaseDeleteResponse, error)
	Put(context.Context, *DatabasePutRequest) (*DatabasePutResponse, error)
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

func _Database_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knv.Database/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Get(ctx, req.(*DatabaseGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knv.Database/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Delete(ctx, req.(*DatabaseDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabasePutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knv.Database/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Put(ctx, req.(*DatabasePutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "knv.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Database_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Database_Delete_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Database_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "knv.proto",
}
