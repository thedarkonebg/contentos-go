// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extDailyTrx.proto

package table

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoExtDailyTrx struct {
	Date                 int64    `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	Count                uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoExtDailyTrx) Reset()         { *m = SoExtDailyTrx{} }
func (m *SoExtDailyTrx) String() string { return proto.CompactTextString(m) }
func (*SoExtDailyTrx) ProtoMessage()    {}
func (*SoExtDailyTrx) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{0}
}

func (m *SoExtDailyTrx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtDailyTrx.Unmarshal(m, b)
}
func (m *SoExtDailyTrx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtDailyTrx.Marshal(b, m, deterministic)
}
func (m *SoExtDailyTrx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtDailyTrx.Merge(m, src)
}
func (m *SoExtDailyTrx) XXX_Size() int {
	return xxx_messageInfo_SoExtDailyTrx.Size(m)
}
func (m *SoExtDailyTrx) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtDailyTrx.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtDailyTrx proto.InternalMessageInfo

func (m *SoExtDailyTrx) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *SoExtDailyTrx) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type SoMemExtDailyTrxByDate struct {
	Date                 int64    `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemExtDailyTrxByDate) Reset()         { *m = SoMemExtDailyTrxByDate{} }
func (m *SoMemExtDailyTrxByDate) String() string { return proto.CompactTextString(m) }
func (*SoMemExtDailyTrxByDate) ProtoMessage()    {}
func (*SoMemExtDailyTrxByDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{1}
}

func (m *SoMemExtDailyTrxByDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtDailyTrxByDate.Unmarshal(m, b)
}
func (m *SoMemExtDailyTrxByDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtDailyTrxByDate.Marshal(b, m, deterministic)
}
func (m *SoMemExtDailyTrxByDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtDailyTrxByDate.Merge(m, src)
}
func (m *SoMemExtDailyTrxByDate) XXX_Size() int {
	return xxx_messageInfo_SoMemExtDailyTrxByDate.Size(m)
}
func (m *SoMemExtDailyTrxByDate) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtDailyTrxByDate.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtDailyTrxByDate proto.InternalMessageInfo

func (m *SoMemExtDailyTrxByDate) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type SoMemExtDailyTrxByCount struct {
	Count                uint64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemExtDailyTrxByCount) Reset()         { *m = SoMemExtDailyTrxByCount{} }
func (m *SoMemExtDailyTrxByCount) String() string { return proto.CompactTextString(m) }
func (*SoMemExtDailyTrxByCount) ProtoMessage()    {}
func (*SoMemExtDailyTrxByCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{2}
}

func (m *SoMemExtDailyTrxByCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtDailyTrxByCount.Unmarshal(m, b)
}
func (m *SoMemExtDailyTrxByCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtDailyTrxByCount.Marshal(b, m, deterministic)
}
func (m *SoMemExtDailyTrxByCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtDailyTrxByCount.Merge(m, src)
}
func (m *SoMemExtDailyTrxByCount) XXX_Size() int {
	return xxx_messageInfo_SoMemExtDailyTrxByCount.Size(m)
}
func (m *SoMemExtDailyTrxByCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtDailyTrxByCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtDailyTrxByCount proto.InternalMessageInfo

func (m *SoMemExtDailyTrxByCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type SoListExtDailyTrxByDate struct {
	Date                 int64    `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoListExtDailyTrxByDate) Reset()         { *m = SoListExtDailyTrxByDate{} }
func (m *SoListExtDailyTrxByDate) String() string { return proto.CompactTextString(m) }
func (*SoListExtDailyTrxByDate) ProtoMessage()    {}
func (*SoListExtDailyTrxByDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{3}
}

func (m *SoListExtDailyTrxByDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtDailyTrxByDate.Unmarshal(m, b)
}
func (m *SoListExtDailyTrxByDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtDailyTrxByDate.Marshal(b, m, deterministic)
}
func (m *SoListExtDailyTrxByDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtDailyTrxByDate.Merge(m, src)
}
func (m *SoListExtDailyTrxByDate) XXX_Size() int {
	return xxx_messageInfo_SoListExtDailyTrxByDate.Size(m)
}
func (m *SoListExtDailyTrxByDate) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtDailyTrxByDate.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtDailyTrxByDate proto.InternalMessageInfo

func (m *SoListExtDailyTrxByDate) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type SoListExtDailyTrxByCount struct {
	Count                uint64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Date                 int64    `protobuf:"varint,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoListExtDailyTrxByCount) Reset()         { *m = SoListExtDailyTrxByCount{} }
func (m *SoListExtDailyTrxByCount) String() string { return proto.CompactTextString(m) }
func (*SoListExtDailyTrxByCount) ProtoMessage()    {}
func (*SoListExtDailyTrxByCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{4}
}

func (m *SoListExtDailyTrxByCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtDailyTrxByCount.Unmarshal(m, b)
}
func (m *SoListExtDailyTrxByCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtDailyTrxByCount.Marshal(b, m, deterministic)
}
func (m *SoListExtDailyTrxByCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtDailyTrxByCount.Merge(m, src)
}
func (m *SoListExtDailyTrxByCount) XXX_Size() int {
	return xxx_messageInfo_SoListExtDailyTrxByCount.Size(m)
}
func (m *SoListExtDailyTrxByCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtDailyTrxByCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtDailyTrxByCount proto.InternalMessageInfo

func (m *SoListExtDailyTrxByCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SoListExtDailyTrxByCount) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type SoUniqueExtDailyTrxByDate struct {
	Date                 int64    `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueExtDailyTrxByDate) Reset()         { *m = SoUniqueExtDailyTrxByDate{} }
func (m *SoUniqueExtDailyTrxByDate) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtDailyTrxByDate) ProtoMessage()    {}
func (*SoUniqueExtDailyTrxByDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{5}
}

func (m *SoUniqueExtDailyTrxByDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtDailyTrxByDate.Unmarshal(m, b)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtDailyTrxByDate.Marshal(b, m, deterministic)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtDailyTrxByDate.Merge(m, src)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtDailyTrxByDate.Size(m)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtDailyTrxByDate.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtDailyTrxByDate proto.InternalMessageInfo

func (m *SoUniqueExtDailyTrxByDate) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func init() {
	proto.RegisterType((*SoExtDailyTrx)(nil), "table.so_extDailyTrx")
	proto.RegisterType((*SoMemExtDailyTrxByDate)(nil), "table.so_mem_extDailyTrx_by_date")
	proto.RegisterType((*SoMemExtDailyTrxByCount)(nil), "table.so_mem_extDailyTrx_by_count")
	proto.RegisterType((*SoListExtDailyTrxByDate)(nil), "table.so_list_extDailyTrx_by_date")
	proto.RegisterType((*SoListExtDailyTrxByCount)(nil), "table.so_list_extDailyTrx_by_count")
	proto.RegisterType((*SoUniqueExtDailyTrxByDate)(nil), "table.so_unique_extDailyTrx_by_date")
}

func init() { proto.RegisterFile("app/table/so_extDailyTrx.proto", fileDescriptor_a71b07e7e9721030) }

var fileDescriptor_a71b07e7e9721030 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2c, 0x28, 0xd0,
	0x2f, 0x49, 0x4c, 0xca, 0x49, 0xd5, 0x2f, 0xce, 0x8f, 0x4f, 0xad, 0x28, 0x71, 0x49, 0xcc, 0xcc,
	0xa9, 0x0c, 0x29, 0xaa, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xcb, 0x29, 0x59,
	0x71, 0xf1, 0xa1, 0x4a, 0x0b, 0x09, 0x71, 0xb1, 0xa4, 0x24, 0x96, 0xa4, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x30, 0x07, 0x81, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x10, 0x8e, 0x92, 0x01, 0x97, 0x54, 0x71, 0x7e, 0x7c, 0x6e, 0x6a,
	0x2e, 0xb2, 0xfe, 0xf8, 0xa4, 0xca, 0x78, 0xb0, 0x1e, 0x2c, 0xe6, 0x28, 0x19, 0x73, 0x49, 0x63,
	0xd7, 0x01, 0x36, 0x10, 0x61, 0x0d, 0x23, 0xb2, 0x35, 0x86, 0x60, 0x4d, 0x39, 0x99, 0xc5, 0x25,
	0x44, 0xdb, 0xe3, 0xc1, 0x25, 0x83, 0x43, 0x0b, 0x1e, 0x8b, 0xe0, 0x26, 0x31, 0xa1, 0xb8, 0x58,
	0xb6, 0x38, 0x3f, 0xbe, 0x34, 0x2f, 0xb3, 0xb0, 0x34, 0x95, 0x58, 0xeb, 0x9d, 0x34, 0xa2, 0xd4,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x8b, 0x93, 0x33,
	0x12, 0x33, 0xf3, 0xf4, 0x93, 0xf3, 0xf3, 0x4a, 0x52, 0xf3, 0x4a, 0xf2, 0x8b, 0x75, 0xd3, 0xf3,
	0x21, 0x51, 0x93, 0xc4, 0x06, 0x8e, 0x0c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0xdb,
	0xf6, 0xb8, 0xae, 0x01, 0x00, 0x00,
}
