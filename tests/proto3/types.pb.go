// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package proto3

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

type BlockHeaderV1 struct {
	Magic                uint32   `protobuf:"varint,1,opt,name=magic,proto3" json:"magic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeaderV1) Reset()         { *m = BlockHeaderV1{} }
func (m *BlockHeaderV1) String() string { return proto.CompactTextString(m) }
func (*BlockHeaderV1) ProtoMessage()    {}
func (*BlockHeaderV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *BlockHeaderV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeaderV1.Unmarshal(m, b)
}
func (m *BlockHeaderV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeaderV1.Marshal(b, m, deterministic)
}
func (m *BlockHeaderV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeaderV1.Merge(m, src)
}
func (m *BlockHeaderV1) XXX_Size() int {
	return xxx_messageInfo_BlockHeaderV1.Size(m)
}
func (m *BlockHeaderV1) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeaderV1.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeaderV1 proto.InternalMessageInfo

func (m *BlockHeaderV1) GetMagic() uint32 {
	if m != nil {
		return m.Magic
	}
	return 0
}

type PersonRecordV1 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Gender               bool     `protobuf:"varint,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Age                  uint32   `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonRecordV1) Reset()         { *m = PersonRecordV1{} }
func (m *PersonRecordV1) String() string { return proto.CompactTextString(m) }
func (*PersonRecordV1) ProtoMessage()    {}
func (*PersonRecordV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *PersonRecordV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonRecordV1.Unmarshal(m, b)
}
func (m *PersonRecordV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonRecordV1.Marshal(b, m, deterministic)
}
func (m *PersonRecordV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonRecordV1.Merge(m, src)
}
func (m *PersonRecordV1) XXX_Size() int {
	return xxx_messageInfo_PersonRecordV1.Size(m)
}
func (m *PersonRecordV1) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonRecordV1.DiscardUnknown(m)
}

var xxx_messageInfo_PersonRecordV1 proto.InternalMessageInfo

func (m *PersonRecordV1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonRecordV1) GetGender() bool {
	if m != nil {
		return m.Gender
	}
	return false
}

func (m *PersonRecordV1) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type BookRecordV1 struct {
	Isdn                 string   `protobuf:"bytes,1,opt,name=isdn,proto3" json:"isdn,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookRecordV1) Reset()         { *m = BookRecordV1{} }
func (m *BookRecordV1) String() string { return proto.CompactTextString(m) }
func (*BookRecordV1) ProtoMessage()    {}
func (*BookRecordV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

func (m *BookRecordV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookRecordV1.Unmarshal(m, b)
}
func (m *BookRecordV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookRecordV1.Marshal(b, m, deterministic)
}
func (m *BookRecordV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookRecordV1.Merge(m, src)
}
func (m *BookRecordV1) XXX_Size() int {
	return xxx_messageInfo_BookRecordV1.Size(m)
}
func (m *BookRecordV1) XXX_DiscardUnknown() {
	xxx_messageInfo_BookRecordV1.DiscardUnknown(m)
}

var xxx_messageInfo_BookRecordV1 proto.InternalMessageInfo

func (m *BookRecordV1) GetIsdn() string {
	if m != nil {
		return m.Isdn
	}
	return ""
}

func (m *BookRecordV1) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BookRecordV1) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type BlockRecordV1 struct {
	// Types that are valid to be assigned to Record:
	//	*BlockRecordV1_Person
	//	*BlockRecordV1_Book
	Record               isBlockRecordV1_Record `protobuf_oneof:"record"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BlockRecordV1) Reset()         { *m = BlockRecordV1{} }
func (m *BlockRecordV1) String() string { return proto.CompactTextString(m) }
func (*BlockRecordV1) ProtoMessage()    {}
func (*BlockRecordV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

func (m *BlockRecordV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRecordV1.Unmarshal(m, b)
}
func (m *BlockRecordV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRecordV1.Marshal(b, m, deterministic)
}
func (m *BlockRecordV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRecordV1.Merge(m, src)
}
func (m *BlockRecordV1) XXX_Size() int {
	return xxx_messageInfo_BlockRecordV1.Size(m)
}
func (m *BlockRecordV1) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRecordV1.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRecordV1 proto.InternalMessageInfo

type isBlockRecordV1_Record interface {
	isBlockRecordV1_Record()
}

type BlockRecordV1_Person struct {
	Person *PersonRecordV1 `protobuf:"bytes,1,opt,name=person,proto3,oneof"`
}

type BlockRecordV1_Book struct {
	Book *BookRecordV1 `protobuf:"bytes,2,opt,name=book,proto3,oneof"`
}

func (*BlockRecordV1_Person) isBlockRecordV1_Record() {}

func (*BlockRecordV1_Book) isBlockRecordV1_Record() {}

func (m *BlockRecordV1) GetRecord() isBlockRecordV1_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *BlockRecordV1) GetPerson() *PersonRecordV1 {
	if x, ok := m.GetRecord().(*BlockRecordV1_Person); ok {
		return x.Person
	}
	return nil
}

func (m *BlockRecordV1) GetBook() *BookRecordV1 {
	if x, ok := m.GetRecord().(*BlockRecordV1_Book); ok {
		return x.Book
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BlockRecordV1) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BlockRecordV1_Person)(nil),
		(*BlockRecordV1_Book)(nil),
	}
}

type BlockV1 struct {
	Header               *BlockHeaderV1   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Records              []*BlockRecordV1 `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlockV1) Reset()         { *m = BlockV1{} }
func (m *BlockV1) String() string { return proto.CompactTextString(m) }
func (*BlockV1) ProtoMessage()    {}
func (*BlockV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}

func (m *BlockV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockV1.Unmarshal(m, b)
}
func (m *BlockV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockV1.Marshal(b, m, deterministic)
}
func (m *BlockV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockV1.Merge(m, src)
}
func (m *BlockV1) XXX_Size() int {
	return xxx_messageInfo_BlockV1.Size(m)
}
func (m *BlockV1) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockV1.DiscardUnknown(m)
}

var xxx_messageInfo_BlockV1 proto.InternalMessageInfo

func (m *BlockV1) GetHeader() *BlockHeaderV1 {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BlockV1) GetRecords() []*BlockRecordV1 {
	if m != nil {
		return m.Records
	}
	return nil
}

type BlockHeaderV2 struct {
	Magic                uint32   `protobuf:"varint,1,opt,name=magic,proto3" json:"magic,omitempty"`
	Timestamp            uint64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeaderV2) Reset()         { *m = BlockHeaderV2{} }
func (m *BlockHeaderV2) String() string { return proto.CompactTextString(m) }
func (*BlockHeaderV2) ProtoMessage()    {}
func (*BlockHeaderV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{5}
}

func (m *BlockHeaderV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeaderV2.Unmarshal(m, b)
}
func (m *BlockHeaderV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeaderV2.Marshal(b, m, deterministic)
}
func (m *BlockHeaderV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeaderV2.Merge(m, src)
}
func (m *BlockHeaderV2) XXX_Size() int {
	return xxx_messageInfo_BlockHeaderV2.Size(m)
}
func (m *BlockHeaderV2) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeaderV2.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeaderV2 proto.InternalMessageInfo

func (m *BlockHeaderV2) GetMagic() uint32 {
	if m != nil {
		return m.Magic
	}
	return 0
}

func (m *BlockHeaderV2) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PersonRecordV2 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Gender               bool     `protobuf:"varint,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Age                  uint32   `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonRecordV2) Reset()         { *m = PersonRecordV2{} }
func (m *PersonRecordV2) String() string { return proto.CompactTextString(m) }
func (*PersonRecordV2) ProtoMessage()    {}
func (*PersonRecordV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{6}
}

func (m *PersonRecordV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonRecordV2.Unmarshal(m, b)
}
func (m *PersonRecordV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonRecordV2.Marshal(b, m, deterministic)
}
func (m *PersonRecordV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonRecordV2.Merge(m, src)
}
func (m *PersonRecordV2) XXX_Size() int {
	return xxx_messageInfo_PersonRecordV2.Size(m)
}
func (m *PersonRecordV2) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonRecordV2.DiscardUnknown(m)
}

var xxx_messageInfo_PersonRecordV2 proto.InternalMessageInfo

func (m *PersonRecordV2) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonRecordV2) GetGender() bool {
	if m != nil {
		return m.Gender
	}
	return false
}

func (m *PersonRecordV2) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *PersonRecordV2) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type BookRecordV2 struct {
	Isdn                 string   `protobuf:"bytes,1,opt,name=isdn,proto3" json:"isdn,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookRecordV2) Reset()         { *m = BookRecordV2{} }
func (m *BookRecordV2) String() string { return proto.CompactTextString(m) }
func (*BookRecordV2) ProtoMessage()    {}
func (*BookRecordV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{7}
}

func (m *BookRecordV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookRecordV2.Unmarshal(m, b)
}
func (m *BookRecordV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookRecordV2.Marshal(b, m, deterministic)
}
func (m *BookRecordV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookRecordV2.Merge(m, src)
}
func (m *BookRecordV2) XXX_Size() int {
	return xxx_messageInfo_BookRecordV2.Size(m)
}
func (m *BookRecordV2) XXX_DiscardUnknown() {
	xxx_messageInfo_BookRecordV2.DiscardUnknown(m)
}

var xxx_messageInfo_BookRecordV2 proto.InternalMessageInfo

func (m *BookRecordV2) GetIsdn() string {
	if m != nil {
		return m.Isdn
	}
	return ""
}

func (m *BookRecordV2) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BookRecordV2) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type CarRecordV2 struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Color                string   `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	Power                uint32   `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CarRecordV2) Reset()         { *m = CarRecordV2{} }
func (m *CarRecordV2) String() string { return proto.CompactTextString(m) }
func (*CarRecordV2) ProtoMessage()    {}
func (*CarRecordV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{8}
}

func (m *CarRecordV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarRecordV2.Unmarshal(m, b)
}
func (m *CarRecordV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarRecordV2.Marshal(b, m, deterministic)
}
func (m *CarRecordV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarRecordV2.Merge(m, src)
}
func (m *CarRecordV2) XXX_Size() int {
	return xxx_messageInfo_CarRecordV2.Size(m)
}
func (m *CarRecordV2) XXX_DiscardUnknown() {
	xxx_messageInfo_CarRecordV2.DiscardUnknown(m)
}

var xxx_messageInfo_CarRecordV2 proto.InternalMessageInfo

func (m *CarRecordV2) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CarRecordV2) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *CarRecordV2) GetPower() uint32 {
	if m != nil {
		return m.Power
	}
	return 0
}

// new in v2
type BlockRecordV2 struct {
	// Types that are valid to be assigned to Record:
	//	*BlockRecordV2_Person
	//	*BlockRecordV2_Book
	//	*BlockRecordV2_Car
	Record               isBlockRecordV2_Record `protobuf_oneof:"record"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BlockRecordV2) Reset()         { *m = BlockRecordV2{} }
func (m *BlockRecordV2) String() string { return proto.CompactTextString(m) }
func (*BlockRecordV2) ProtoMessage()    {}
func (*BlockRecordV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{9}
}

func (m *BlockRecordV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRecordV2.Unmarshal(m, b)
}
func (m *BlockRecordV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRecordV2.Marshal(b, m, deterministic)
}
func (m *BlockRecordV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRecordV2.Merge(m, src)
}
func (m *BlockRecordV2) XXX_Size() int {
	return xxx_messageInfo_BlockRecordV2.Size(m)
}
func (m *BlockRecordV2) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRecordV2.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRecordV2 proto.InternalMessageInfo

type isBlockRecordV2_Record interface {
	isBlockRecordV2_Record()
}

type BlockRecordV2_Person struct {
	Person *PersonRecordV2 `protobuf:"bytes,1,opt,name=person,proto3,oneof"`
}

type BlockRecordV2_Book struct {
	Book *BookRecordV2 `protobuf:"bytes,2,opt,name=book,proto3,oneof"`
}

type BlockRecordV2_Car struct {
	Car *CarRecordV2 `protobuf:"bytes,3,opt,name=car,proto3,oneof"`
}

func (*BlockRecordV2_Person) isBlockRecordV2_Record() {}

func (*BlockRecordV2_Book) isBlockRecordV2_Record() {}

func (*BlockRecordV2_Car) isBlockRecordV2_Record() {}

func (m *BlockRecordV2) GetRecord() isBlockRecordV2_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *BlockRecordV2) GetPerson() *PersonRecordV2 {
	if x, ok := m.GetRecord().(*BlockRecordV2_Person); ok {
		return x.Person
	}
	return nil
}

func (m *BlockRecordV2) GetBook() *BookRecordV2 {
	if x, ok := m.GetRecord().(*BlockRecordV2_Book); ok {
		return x.Book
	}
	return nil
}

func (m *BlockRecordV2) GetCar() *CarRecordV2 {
	if x, ok := m.GetRecord().(*BlockRecordV2_Car); ok {
		return x.Car
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BlockRecordV2) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BlockRecordV2_Person)(nil),
		(*BlockRecordV2_Book)(nil),
		(*BlockRecordV2_Car)(nil),
	}
}

type BlockV2 struct {
	Header               *BlockHeaderV2   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Records              []*BlockRecordV2 `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	Signature            []byte           `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlockV2) Reset()         { *m = BlockV2{} }
func (m *BlockV2) String() string { return proto.CompactTextString(m) }
func (*BlockV2) ProtoMessage()    {}
func (*BlockV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{10}
}

func (m *BlockV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockV2.Unmarshal(m, b)
}
func (m *BlockV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockV2.Marshal(b, m, deterministic)
}
func (m *BlockV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockV2.Merge(m, src)
}
func (m *BlockV2) XXX_Size() int {
	return xxx_messageInfo_BlockV2.Size(m)
}
func (m *BlockV2) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockV2.DiscardUnknown(m)
}

var xxx_messageInfo_BlockV2 proto.InternalMessageInfo

func (m *BlockV2) GetHeader() *BlockHeaderV2 {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BlockV2) GetRecords() []*BlockRecordV2 {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *BlockV2) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeaderV1)(nil), "block_header_v1")
	proto.RegisterType((*PersonRecordV1)(nil), "person_record_v1")
	proto.RegisterType((*BookRecordV1)(nil), "book_record_v1")
	proto.RegisterType((*BlockRecordV1)(nil), "block_record_v1")
	proto.RegisterType((*BlockV1)(nil), "block_v1")
	proto.RegisterType((*BlockHeaderV2)(nil), "block_header_v2")
	proto.RegisterType((*PersonRecordV2)(nil), "person_record_v2")
	proto.RegisterType((*BookRecordV2)(nil), "book_record_v2")
	proto.RegisterType((*CarRecordV2)(nil), "car_record_v2")
	proto.RegisterType((*BlockRecordV2)(nil), "block_record_v2")
	proto.RegisterType((*BlockV2)(nil), "block_v2")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x4d, 0x9a, 0xd4, 0x4d, 0x6e, 0xd6, 0x36, 0x13, 0x65, 0xf8, 0x61, 0x0f, 0xc1, 0x30, 0x16,
	0x36, 0x6a, 0x53, 0xed, 0x1f, 0x14, 0x06, 0x79, 0xdc, 0xf4, 0xb8, 0x97, 0x4c, 0x96, 0x85, 0xe3,
	0x35, 0xd6, 0x35, 0x92, 0xd2, 0xb1, 0x97, 0xfd, 0x89, 0xfd, 0xe1, 0x21, 0xc9, 0xce, 0x57, 0x07,
	0x0d, 0xac, 0x4f, 0xd1, 0xb9, 0xba, 0x39, 0xe7, 0x5e, 0x9d, 0x83, 0x61, 0x62, 0x7f, 0x35, 0xd2,
	0xa4, 0x8d, 0x46, 0x8b, 0xc9, 0x7b, 0xb8, 0xce, 0xd7, 0x28, 0x1e, 0x96, 0x2b, 0xc9, 0x0b, 0xa9,
	0x97, 0x8f, 0x77, 0xe4, 0x06, 0xce, 0x6b, 0x5e, 0x56, 0x22, 0xee, 0xcf, 0xfa, 0xf3, 0x4b, 0x16,
	0x40, 0xf2, 0x05, 0xa6, 0x8d, 0xd4, 0x06, 0xd5, 0x52, 0x4b, 0x81, 0xba, 0x70, 0x9d, 0x04, 0x86,
	0x8a, 0xd7, 0xd2, 0x37, 0x8e, 0x99, 0x3f, 0x93, 0x37, 0x10, 0x95, 0x52, 0x15, 0x52, 0xc7, 0x67,
	0xb3, 0xfe, 0x7c, 0xc4, 0x5a, 0x44, 0xa6, 0x30, 0xe0, 0xa5, 0x8c, 0x07, 0x9e, 0xd3, 0x1d, 0x13,
	0x06, 0x57, 0x39, 0xe2, 0xc3, 0x21, 0x5f, 0x65, 0x0a, 0xd5, 0xf1, 0xb9, 0xb3, 0x9b, 0xc6, 0x56,
	0x76, 0x2d, 0x3d, 0xdd, 0x98, 0x05, 0xe0, 0x54, 0xf8, 0xc6, 0xae, 0x50, 0x7b, 0xc2, 0x31, 0x6b,
	0x51, 0x82, 0xdd, 0x3a, 0x3b, 0xd2, 0x8f, 0x10, 0x85, 0xc1, 0x3d, 0xed, 0x84, 0xbe, 0x4e, 0x8f,
	0xf7, 0x58, 0xf4, 0x58, 0xdb, 0x42, 0xde, 0xc1, 0xd0, 0xcd, 0xe4, 0xc5, 0x26, 0xf4, 0x3a, 0x3d,
	0x1c, 0x70, 0xd1, 0x63, 0xfe, 0xfa, 0x7e, 0x04, 0x51, 0x28, 0x26, 0xdf, 0x61, 0x14, 0x04, 0x1f,
	0xef, 0xc8, 0x1c, 0xa2, 0xf0, 0x8a, 0xad, 0xd2, 0x34, 0x3d, 0x7a, 0x5a, 0xd6, 0xde, 0x93, 0x0f,
	0x70, 0x11, 0xfe, 0x6f, 0xe2, 0xb3, 0xd9, 0x60, 0xaf, 0x75, 0x2b, 0xc5, 0xba, 0x86, 0xe4, 0xf3,
	0xb1, 0x43, 0xf4, 0xdf, 0x0e, 0x91, 0xb7, 0x30, 0xb6, 0x55, 0x2d, 0x8d, 0xe5, 0x75, 0xe3, 0x17,
	0x18, 0xb2, 0x5d, 0x21, 0xf9, 0xf1, 0xc4, 0x3f, 0xfa, 0x7f, 0xfe, 0x91, 0x18, 0x2e, 0x78, 0x51,
	0x68, 0x69, 0x4c, 0x3c, 0xf4, 0x04, 0x1d, 0x7c, 0xe2, 0x2c, 0x7d, 0x01, 0x67, 0xbf, 0xc2, 0xa5,
	0xe0, 0x7a, 0x8f, 0xf2, 0x06, 0xce, 0x73, 0xcd, 0x55, 0xd1, 0x72, 0x06, 0xe0, 0xaa, 0x02, 0xd7,
	0xa8, 0x3b, 0x52, 0x0f, 0x5c, 0xb5, 0xc1, 0x9f, 0x52, 0xb7, 0xe3, 0x07, 0x90, 0xfc, 0xe9, 0x1f,
	0xa7, 0x85, 0x3e, 0x9f, 0x16, 0x7a, 0x6a, 0x5a, 0x68, 0x97, 0x16, 0x92, 0xc0, 0x40, 0xf0, 0xa0,
	0x3d, 0xa1, 0x57, 0xe9, 0xc1, 0x1a, 0x8b, 0x1e, 0x73, 0x97, 0x7b, 0x89, 0xfa, 0xbd, 0x4d, 0x14,
	0x7d, 0x36, 0x51, 0xf4, 0xf4, 0x44, 0xd1, 0x6d, 0xa2, 0x5c, 0x50, 0x4c, 0x55, 0x2a, 0x6e, 0x37,
	0x3a, 0x18, 0xfa, 0x8a, 0xed, 0x0a, 0xf7, 0xd9, 0xb7, 0xdb, 0xb2, 0xb2, 0xab, 0x4d, 0x9e, 0x0a,
	0xac, 0x33, 0x81, 0x46, 0xac, 0x78, 0xa5, 0x32, 0x81, 0xca, 0x4a, 0x65, 0xd1, 0xdc, 0x96, 0x98,
	0x59, 0x69, 0xac, 0xc9, 0xfc, 0x17, 0xe4, 0x53, 0x1e, 0x85, 0xdf, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x83, 0xb3, 0xa3, 0x3e, 0x58, 0x04, 0x00, 0x00,
}