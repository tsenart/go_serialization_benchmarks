// Code generated by protoc-gen-go.
// source: structdef.proto
// DO NOT EDIT!

/*
Package goserbench is a generated protocol buffer package.

It is generated from these files:
	structdef.proto

It has these top-level messages:
	ProtoBufA
*/
package goserbench

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ProtoBufA struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	BirthDay         *int64   `protobuf:"varint,2,req,name=birthDay" json:"birthDay,omitempty"`
	Phone            *string  `protobuf:"bytes,3,req,name=phone" json:"phone,omitempty"`
	Siblings         *int32   `protobuf:"varint,4,req,name=siblings" json:"siblings,omitempty"`
	Spouse           *bool    `protobuf:"varint,5,req,name=spouse" json:"spouse,omitempty"`
	Money            *float64 `protobuf:"fixed64,6,req,name=money" json:"money,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ProtoBufA) Reset()                    { *m = ProtoBufA{} }
func (m *ProtoBufA) String() string            { return proto.CompactTextString(m) }
func (*ProtoBufA) ProtoMessage()               {}
func (*ProtoBufA) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoBufA) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProtoBufA) GetBirthDay() int64 {
	if m != nil && m.BirthDay != nil {
		return *m.BirthDay
	}
	return 0
}

func (m *ProtoBufA) GetPhone() string {
	if m != nil && m.Phone != nil {
		return *m.Phone
	}
	return ""
}

func (m *ProtoBufA) GetSiblings() int32 {
	if m != nil && m.Siblings != nil {
		return *m.Siblings
	}
	return 0
}

func (m *ProtoBufA) GetSpouse() bool {
	if m != nil && m.Spouse != nil {
		return *m.Spouse
	}
	return false
}

func (m *ProtoBufA) GetMoney() float64 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtoBufA)(nil), "goserbench.ProtoBufA")
}

var fileDescriptor0 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x34, 0xcb, 0x39, 0x0e, 0xc2, 0x30,
	0x10, 0x85, 0x61, 0x65, 0x55, 0x32, 0x62, 0x93, 0xab, 0x29, 0x11, 0x15, 0x15, 0x77, 0x00, 0x71,
	0x00, 0xae, 0x10, 0x87, 0x49, 0x6c, 0x41, 0x6c, 0xcb, 0x63, 0x17, 0xb9, 0x3d, 0x93, 0x82, 0xf2,
	0x7d, 0xfa, 0x1f, 0x1c, 0x39, 0xc5, 0x3c, 0xa6, 0x37, 0x4d, 0xb7, 0x10, 0x7d, 0xf2, 0x0a, 0x66,
	0xcf, 0x14, 0x35, 0xb9, 0xd1, 0x5c, 0x3e, 0xd0, 0xbf, 0x36, 0x7c, 0xe4, 0xe9, 0xae, 0x76, 0x50,
	0xbb, 0x61, 0x21, 0x2c, 0xce, 0xe5, 0xb5, 0x57, 0x27, 0xe8, 0xb4, 0x8d, 0xc9, 0x3c, 0x87, 0x15,
	0x4b, 0x91, 0x4a, 0xed, 0xa1, 0x09, 0xc6, 0x3b, 0xc2, 0xea, 0x1f, 0xb0, 0xd5, 0x5f, 0xeb, 0x66,
	0xc6, 0x5a, 0xa4, 0x51, 0x07, 0x68, 0x39, 0xf8, 0xcc, 0x84, 0x8d, 0xec, 0x6e, 0x3b, 0x2c, 0xd2,
	0xaf, 0xd8, 0xca, 0x2c, 0x7e, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xa4, 0xa8, 0x5c, 0x8a, 0x00,
	0x00, 0x00,
}