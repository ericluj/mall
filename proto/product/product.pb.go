// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package product

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ProductRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Num                  int64    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *ProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRequest.Unmarshal(m, b)
}
func (m *ProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRequest.Marshal(b, m, deterministic)
}
func (m *ProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRequest.Merge(m, src)
}
func (m *ProductRequest) XXX_Size() int {
	return xxx_messageInfo_ProductRequest.Size(m)
}
func (m *ProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRequest proto.InternalMessageInfo

func (m *ProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductRequest) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ProductResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Num                  int64    `protobuf:"varint,5,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductResponse) Reset()         { *m = ProductResponse{} }
func (m *ProductResponse) String() string { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()    {}
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *ProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResponse.Unmarshal(m, b)
}
func (m *ProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResponse.Marshal(b, m, deterministic)
}
func (m *ProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResponse.Merge(m, src)
}
func (m *ProductResponse) XXX_Size() int {
	return xxx_messageInfo_ProductResponse.Size(m)
}
func (m *ProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResponse proto.InternalMessageInfo

func (m *ProductResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *ProductResponse) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *ProductResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductResponse) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "ProductRequest")
	proto.RegisterType((*ProductResponse)(nil), "ProductResponse")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe3, 0xe2, 0x0b, 0x80, 0x08,
	0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0x79, 0xa5, 0xb9, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x20, 0xa6, 0x52, 0x33, 0x23, 0x17, 0x3f, 0x5c, 0x63, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x0a, 0x58, 0x1f, 0x73, 0x10, 0x53,
	0x66, 0x8a, 0x90, 0x0c, 0x17, 0x67, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x6a, 0x8a, 0x63, 0x09, 0x54,
	0x2f, 0x42, 0x00, 0x24, 0x5b, 0x5a, 0x90, 0x02, 0x95, 0x65, 0x86, 0xc8, 0xc2, 0x05, 0xe0, 0xae,
	0x60, 0xc1, 0x74, 0x05, 0x2b, 0xdc, 0x15, 0x46, 0xab, 0x18, 0xb9, 0xd8, 0xa1, 0xae, 0x10, 0x32,
	0xe1, 0xe2, 0x75, 0x06, 0x1b, 0x0e, 0x13, 0xe0, 0xd7, 0x43, 0xf5, 0x99, 0x94, 0x80, 0x1e, 0x9a,
	0x8b, 0x95, 0x18, 0x84, 0x8c, 0xb9, 0x78, 0x02, 0x4b, 0x53, 0x8b, 0x2a, 0x49, 0xd2, 0x64, 0xc2,
	0xc5, 0x1b, 0x0a, 0x76, 0x29, 0x29, 0xba, 0x92, 0xd8, 0xc0, 0x21, 0x6e, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x14, 0x8e, 0x94, 0xbf, 0x82, 0x01, 0x00, 0x00,
}
