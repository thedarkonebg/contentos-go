// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_following.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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

type SoFollowing struct {
	FollowingInfo        *prototype.FollowingRelation `protobuf:"bytes,1,opt,name=following_info,json=followingInfo,proto3" json:"following_info,omitempty"`
	CreatedTime          *prototype.TimePointSec      `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoFollowing) Reset()         { *m = SoFollowing{} }
func (m *SoFollowing) String() string { return proto.CompactTextString(m) }
func (*SoFollowing) ProtoMessage()    {}
func (*SoFollowing) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f46122cc3ea918f, []int{0}
}

func (m *SoFollowing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoFollowing.Unmarshal(m, b)
}
func (m *SoFollowing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoFollowing.Marshal(b, m, deterministic)
}
func (m *SoFollowing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoFollowing.Merge(m, src)
}
func (m *SoFollowing) XXX_Size() int {
	return xxx_messageInfo_SoFollowing.Size(m)
}
func (m *SoFollowing) XXX_DiscardUnknown() {
	xxx_messageInfo_SoFollowing.DiscardUnknown(m)
}

var xxx_messageInfo_SoFollowing proto.InternalMessageInfo

func (m *SoFollowing) GetFollowingInfo() *prototype.FollowingRelation {
	if m != nil {
		return m.FollowingInfo
	}
	return nil
}

func (m *SoFollowing) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

type SoUniqueFollowingByFollowingInfo struct {
	FollowingInfo        *prototype.FollowingRelation `protobuf:"bytes,1,opt,name=following_info,json=followingInfo,proto3" json:"following_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoUniqueFollowingByFollowingInfo) Reset()         { *m = SoUniqueFollowingByFollowingInfo{} }
func (m *SoUniqueFollowingByFollowingInfo) String() string { return proto.CompactTextString(m) }
func (*SoUniqueFollowingByFollowingInfo) ProtoMessage()    {}
func (*SoUniqueFollowingByFollowingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f46122cc3ea918f, []int{1}
}

func (m *SoUniqueFollowingByFollowingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueFollowingByFollowingInfo.Unmarshal(m, b)
}
func (m *SoUniqueFollowingByFollowingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueFollowingByFollowingInfo.Marshal(b, m, deterministic)
}
func (m *SoUniqueFollowingByFollowingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueFollowingByFollowingInfo.Merge(m, src)
}
func (m *SoUniqueFollowingByFollowingInfo) XXX_Size() int {
	return xxx_messageInfo_SoUniqueFollowingByFollowingInfo.Size(m)
}
func (m *SoUniqueFollowingByFollowingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueFollowingByFollowingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueFollowingByFollowingInfo proto.InternalMessageInfo

func (m *SoUniqueFollowingByFollowingInfo) GetFollowingInfo() *prototype.FollowingRelation {
	if m != nil {
		return m.FollowingInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*SoFollowing)(nil), "table.so_following")
	proto.RegisterType((*SoUniqueFollowingByFollowingInfo)(nil), "table.so_unique_following_by_following_info")
}

func init() { proto.RegisterFile("app/table/so_following.proto", fileDescriptor_0f46122cc3ea918f) }

var fileDescriptor_0f46122cc3ea918f = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x50, 0xbd, 0x4a, 0x03, 0x41,
	0x10, 0xe6, 0x04, 0x2d, 0x36, 0xd1, 0xe2, 0xb0, 0x88, 0xa2, 0x20, 0x01, 0x25, 0x8d, 0xb7, 0xa0,
	0xad, 0x95, 0xd8, 0xd8, 0x06, 0x2b, 0x9b, 0x61, 0x6f, 0x9d, 0xbb, 0x0c, 0xdc, 0xce, 0xac, 0xb7,
	0x73, 0x48, 0x5e, 0xc3, 0x27, 0x96, 0x6c, 0xe4, 0x72, 0xf6, 0x36, 0x03, 0xdf, 0x2f, 0x1f, 0x63,
	0xae, 0x5c, 0x8c, 0x56, 0x5d, 0xdd, 0xa1, 0x4d, 0x02, 0x8d, 0x74, 0x9d, 0x7c, 0x11, 0xb7, 0x55,
	0xec, 0x45, 0xa5, 0x3c, 0xce, 0xca, 0xe5, 0x79, 0x46, 0xba, 0x8d, 0x68, 0x77, 0x67, 0x2f, 0x2e,
	0xbf, 0x0b, 0x33, 0x9f, 0x66, 0xca, 0x17, 0x73, 0x36, 0x02, 0x20, 0x6e, 0x64, 0x51, 0xdc, 0x14,
	0xab, 0xd9, 0xc3, 0x75, 0x35, 0xe6, 0xab, 0x83, 0xa1, 0xc7, 0xce, 0x29, 0x09, 0xaf, 0x4f, 0x47,
	0xee, 0x95, 0x1b, 0x29, 0x9f, 0xcc, 0xdc, 0xf7, 0xe8, 0x14, 0x3f, 0x40, 0x29, 0xe0, 0xe2, 0x28,
	0x77, 0x5c, 0x4c, 0x3a, 0x76, 0x34, 0x44, 0x21, 0x56, 0x48, 0xe8, 0xd7, 0xb3, 0x5f, 0xfb, 0x1b,
	0x05, 0x5c, 0x06, 0x73, 0x9b, 0x04, 0x06, 0xa6, 0xcf, 0x01, 0x0f, 0xd3, 0xa0, 0xde, 0xc2, 0xdf,
	0x69, 0xff, 0x33, 0xf6, 0x79, 0xf5, 0x7e, 0xd7, 0x92, 0x6e, 0x86, 0xba, 0xf2, 0x12, 0xac, 0x97,
	0xe4, 0x37, 0x8e, 0xd8, 0x7a, 0x61, 0x45, 0x56, 0x49, 0xf7, 0xad, 0xec, 0xbf, 0x5b, 0x9f, 0xe4,
	0xda, 0xc7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x85, 0xda, 0xe9, 0x71, 0x01, 0x00, 0x00,
}
